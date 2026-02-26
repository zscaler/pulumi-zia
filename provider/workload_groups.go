// Copyright (c) 2023 Zscaler Technology Alliances, <devrel@zscaler.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package provider implements the Workload Groups resource and invoke.
// Adopted from terraform-provider-zia resource_zia_workload_groups.go and data_source_zia_workload_groups.go.

package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/workloadgroups"
)

// --- Resource ---

// WorkloadGroup implements the zia:index:WorkloadGroup resource.
type WorkloadGroup struct{}

// WorkloadGroupArgs are the inputs for WorkloadGroup.
type WorkloadGroupArgs struct {
	Name           *string                           `pulumi:"name,optional"`
	Description    *string                           `pulumi:"description,optional"`
	ExpressionJson *WorkloadGroupExpressionJsonInput `pulumi:"expressionJson,optional"`
}

// WorkloadGroupExpressionJsonInput is the expression_json block.
type WorkloadGroupExpressionJsonInput struct {
	ExpressionContainers []WorkloadGroupExpressionContainerInput `pulumi:"expressionContainers,optional"`
}

// WorkloadGroupExpressionContainerInput is an expression container block.
type WorkloadGroupExpressionContainerInput struct {
	TagType      *string                         `pulumi:"tagType,optional"`
	Operator     *string                         `pulumi:"operator,optional"`
	TagContainer *WorkloadGroupTagContainerInput `pulumi:"tagContainer,optional"`
}

// WorkloadGroupTagContainerInput is the tag_container block.
type WorkloadGroupTagContainerInput struct {
	Tags     []WorkloadGroupTagInput `pulumi:"tags,optional"`
	Operator *string                 `pulumi:"operator,optional"`
}

// WorkloadGroupTagInput is a tag key-value pair.
type WorkloadGroupTagInput struct {
	Key   *string `pulumi:"key,optional"`
	Value *string `pulumi:"value,optional"`
}

// WorkloadGroupState is the persisted state.
type WorkloadGroupState struct {
	WorkloadGroupArgs
	GroupId *int `pulumi:"groupId"`
}

func (WorkloadGroup) Check(ctx context.Context, req infer.CheckRequest) (infer.CheckResponse[WorkloadGroupArgs], error) {
	inputs, failures, err := infer.DefaultCheck[WorkloadGroupArgs](ctx, req.NewInputs)
	if err != nil {
		return infer.CheckResponse[WorkloadGroupArgs]{}, err
	}
	if len(failures) > 0 {
		return infer.CheckResponse[WorkloadGroupArgs]{Failures: failures}, nil
	}
	// Validate expression_json tag_type and operator enums if present
	if inputs.ExpressionJson != nil {
		for _, c := range inputs.ExpressionJson.ExpressionContainers {
			if c.TagType != nil {
				valid := map[string]bool{"ANY": true, "VPC": true, "SUBNET": true, "VM": true, "ENI": true, "ATTR": true}
				if !valid[*c.TagType] {
					return infer.CheckResponse[WorkloadGroupArgs]{Failures: []p.CheckFailure{{
						Property: "expressionJson.expressionContainers.tagType",
						Reason:   "tagType must be ANY, VPC, SUBNET, VM, ENI, or ATTR",
					}}}, nil
				}
			}
			if c.Operator != nil {
				valid := map[string]bool{"AND": true, "OR": true, "OPEN_PARENTHESES": true, "CLOSE_PARENTHESES": true}
				if !valid[*c.Operator] {
					return infer.CheckResponse[WorkloadGroupArgs]{Failures: []p.CheckFailure{{
						Property: "expressionJson.expressionContainers.operator",
						Reason:   "operator must be AND, OR, OPEN_PARENTHESES, or CLOSE_PARENTHESES",
					}}}, nil
				}
			}
		}
	}
	return infer.CheckResponse[WorkloadGroupArgs]{Inputs: inputs}, nil
}

func (WorkloadGroup) Create(ctx context.Context, req infer.CreateRequest[WorkloadGroupArgs]) (infer.CreateResponse[WorkloadGroupState], error) {
	if req.DryRun {
		return infer.CreateResponse[WorkloadGroupState]{
			ID: "preview",
			Output: WorkloadGroupState{
				WorkloadGroupArgs: req.Inputs,
				GroupId:           intPtr(0),
			},
		}, nil
	}

	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[WorkloadGroupState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := workloadGroupArgsToAPI(req.Inputs, 0)
	resp, _, err := workloadgroups.Create(ctx, svc, &apiReq)
	if err != nil {
		return infer.CreateResponse[WorkloadGroupState]{}, fmt.Errorf("creating workload group: %w", err)
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.CreateResponse[WorkloadGroupState]{}, activationErr
		}
	}

	rule, err := workloadgroups.Get(ctx, svc, resp.ID)
	if err != nil {
		return infer.CreateResponse[WorkloadGroupState]{
			ID:     strconv.Itoa(resp.ID),
			Output: WorkloadGroupState{WorkloadGroupArgs: req.Inputs, GroupId: &resp.ID},
		}, nil
	}

	state := workloadGroupAPIToState(rule)
	return infer.CreateResponse[WorkloadGroupState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (WorkloadGroup) Read(ctx context.Context, req infer.ReadRequest[WorkloadGroupArgs, WorkloadGroupState]) (infer.ReadResponse[WorkloadGroupArgs, WorkloadGroupState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[WorkloadGroupArgs, WorkloadGroupState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		rule, lookupErr := workloadgroups.GetByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[WorkloadGroupArgs, WorkloadGroupState]{}, fmt.Errorf("workload group not found")
			}
			return infer.ReadResponse[WorkloadGroupArgs, WorkloadGroupState]{}, lookupErr
		}
		id = rule.ID
	}

	rule, err := workloadgroups.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[WorkloadGroupArgs, WorkloadGroupState]{}, fmt.Errorf("workload group not found")
		}
		return infer.ReadResponse[WorkloadGroupArgs, WorkloadGroupState]{}, err
	}

	state := workloadGroupAPIToState(rule)
	args := workloadGroupStateToArgs(rule)
	return infer.ReadResponse[WorkloadGroupArgs, WorkloadGroupState]{
		ID:     strconv.Itoa(rule.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (WorkloadGroup) Update(ctx context.Context, req infer.UpdateRequest[WorkloadGroupArgs, WorkloadGroupState]) (infer.UpdateResponse[WorkloadGroupState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[WorkloadGroupState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[WorkloadGroupState]{}, fmt.Errorf("invalid group id: %w", err)
	}

	apiReq := workloadGroupArgsToAPI(req.Inputs, id)
	if _, _, err := workloadgroups.Update(ctx, svc, id, &apiReq); err != nil {
		return infer.UpdateResponse[WorkloadGroupState]{}, fmt.Errorf("updating workload group: %w", err)
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[WorkloadGroupState]{}, activationErr
		}
	}

	rule, err := workloadgroups.Get(ctx, svc, id)
	if err != nil {
		return infer.UpdateResponse[WorkloadGroupState]{
			Output: WorkloadGroupState{WorkloadGroupArgs: req.Inputs, GroupId: &id},
		}, nil
	}
	return infer.UpdateResponse[WorkloadGroupState]{Output: workloadGroupAPIToState(rule)}, nil
}

func (WorkloadGroup) Delete(ctx context.Context, req infer.DeleteRequest[WorkloadGroupState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid group id: %w", err)
	}

	if _, err := workloadgroups.Delete(ctx, svc, id); err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("deleting workload group: %w", err)
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}
	return infer.DeleteResponse{}, nil
}

func (WorkloadGroup) Annotate(a infer.Annotator) {
	describeResource(a, &WorkloadGroup{}, `The zia.WorkloadGroup resource manages workload groups in the Zscaler Internet Access (ZIA) cloud.
Workload groups define sets of cloud workloads based on tag expressions that can be used in
firewall rules, URL filtering rules, and other policy rules to apply policies to specific
cloud workloads (e.g., VMs, subnets, ENIs).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Workload Group

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.WorkloadGroup("example", {
    name: "Example Workload Group",
    description: "Managed by Pulumi",
    expressionJson: {
        expressionContainers: [{
            tagType: "VM",
            operator: "AND",
            tagContainer: {
                tags: [{
                    key: "environment",
                    value: "production",
                }],
                operator: "OR",
            },
        }],
    },
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.WorkloadGroup("example",
    name="Example Workload Group",
    description="Managed by Pulumi",
    expression_json={
        "expression_containers": [{
            "tag_type": "VM",
            "operator": "AND",
            "tag_container": {
                "tags": [{
                    "key": "environment",
                    "value": "production",
                }],
                "operator": "OR",
            },
        }],
    },
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:WorkloadGroup
    properties:
      name: Example Workload Group
      description: Managed by Pulumi
      expressionJson:
        expressionContainers:
          - tagType: VM
            operator: AND
            tagContainer:
              tags:
                - key: environment
                  value: production
              operator: OR
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing workload group can be imported using its ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:WorkloadGroup example 12345
`+tripleBacktick("")+`
`)
}

func (a *WorkloadGroupArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "Name of the workload group.")
	ann.Describe(&a.Description, "Description of the workload group.")
	ann.Describe(&a.ExpressionJson, "The expression JSON that defines the workload group matching criteria using tag expressions.")
}

func (s *WorkloadGroupState) Annotate(a infer.Annotator) {
	a.Describe(&s.GroupId, "The unique identifier for the workload group assigned by the ZIA cloud.")
}

func workloadGroupArgsToAPI(in WorkloadGroupArgs, existingID int) workloadgroups.WorkloadGroup {
	result := workloadgroups.WorkloadGroup{
		ID:          existingID,
		Name:        ptrToString(in.Name),
		Description: ptrToString(in.Description),
	}
	if in.ExpressionJson != nil {
		result.WorkloadTagExpression = expandWorkloadGroupExpressionJson(in.ExpressionJson)
	}
	return result
}

func expandWorkloadGroupExpressionJson(in *WorkloadGroupExpressionJsonInput) workloadgroups.WorkloadTagExpression {
	if in == nil || len(in.ExpressionContainers) == 0 {
		return workloadgroups.WorkloadTagExpression{}
	}
	containers := make([]workloadgroups.ExpressionContainer, len(in.ExpressionContainers))
	for i, c := range in.ExpressionContainers {
		containers[i] = workloadgroups.ExpressionContainer{
			TagType:  ptrToString(c.TagType),
			Operator: ptrToString(c.Operator),
		}
		if c.TagContainer != nil {
			tc := workloadgroups.TagContainer{
				Operator: ptrToString(c.TagContainer.Operator),
			}
			if len(c.TagContainer.Tags) > 0 {
				tc.Tags = make([]workloadgroups.Tags, len(c.TagContainer.Tags))
				for j, t := range c.TagContainer.Tags {
					tc.Tags[j] = workloadgroups.Tags{
						Key:   ptrToString(t.Key),
						Value: ptrToString(t.Value),
					}
				}
			}
			containers[i].TagContainer = tc
		}
	}
	return workloadgroups.WorkloadTagExpression{ExpressionContainers: containers}
}

func workloadGroupAPIToState(rule *workloadgroups.WorkloadGroup) WorkloadGroupState {
	state := WorkloadGroupState{
		WorkloadGroupArgs: WorkloadGroupArgs{
			Name:        stringPtr(rule.Name),
			Description: stringPtr(rule.Description),
		},
		GroupId: &rule.ID,
	}
	if len(rule.WorkloadTagExpression.ExpressionContainers) > 0 {
		state.ExpressionJson = flattenWorkloadGroupExpressionJson(&rule.WorkloadTagExpression)
	}
	return state
}

func flattenWorkloadGroupExpressionJson(expr *workloadgroups.WorkloadTagExpression) *WorkloadGroupExpressionJsonInput {
	if expr == nil || len(expr.ExpressionContainers) == 0 {
		return nil
	}
	containers := make([]WorkloadGroupExpressionContainerInput, len(expr.ExpressionContainers))
	for i, c := range expr.ExpressionContainers {
		containers[i] = WorkloadGroupExpressionContainerInput{
			TagType:  stringPtr(c.TagType),
			Operator: stringPtr(c.Operator),
		}
		if c.TagContainer.Operator != "" || len(c.TagContainer.Tags) > 0 {
			tc := &WorkloadGroupTagContainerInput{
				Operator: stringPtr(c.TagContainer.Operator),
			}
			if len(c.TagContainer.Tags) > 0 {
				tc.Tags = make([]WorkloadGroupTagInput, len(c.TagContainer.Tags))
				for j, t := range c.TagContainer.Tags {
					tc.Tags[j] = WorkloadGroupTagInput{
						Key:   stringPtr(t.Key),
						Value: stringPtr(t.Value),
					}
				}
			}
			containers[i].TagContainer = tc
		}
	}
	return &WorkloadGroupExpressionJsonInput{ExpressionContainers: containers}
}

func workloadGroupStateToArgs(rule *workloadgroups.WorkloadGroup) WorkloadGroupArgs {
	return workloadGroupAPIToState(rule).WorkloadGroupArgs
}

// --- Invoke (data source) ---

// GetWorkloadGroupArgs are the inputs for the GetWorkloadGroup invoke.
type GetWorkloadGroupArgs struct {
	Id   *int    `pulumi:"groupId,optional"` // Pulumi reserves "id" in function I/O
	Name *string `pulumi:"name,optional"`
}

// GetWorkloadGroupResult is the output of the GetWorkloadGroup invoke.
type GetWorkloadGroupResult struct {
	Id               int                               `pulumi:"groupId"` // Pulumi reserves "id" in function outputs
	Name             string                            `pulumi:"name"`
	Description      string                            `pulumi:"description"`
	Expression       string                            `pulumi:"expression"`
	ExpressionJson   *WorkloadGroupExpressionJsonInput `pulumi:"expressionJson,optional"`
	LastModifiedTime int                               `pulumi:"lastModifiedTime"`
}

// GetWorkloadGroup implements the zia:index:GetWorkloadGroup invoke.
type GetWorkloadGroup struct{}

func (f *GetWorkloadGroup) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a workload group by ID or name.")
}

func (a *GetWorkloadGroupArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Id, "The ID of the workload group to look up.")
	ann.Describe(&a.Name, "The name of the workload group to look up.")
}

func (r *GetWorkloadGroupResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the workload group.")
	a.Describe(&r.Name, "The name of the workload group.")
	a.Describe(&r.Description, "The description of the workload group.")
	a.Describe(&r.Expression, "The expression string for the workload group.")
	a.Describe(&r.ExpressionJson, "The expression JSON that defines the workload group matching criteria.")
	a.Describe(&r.LastModifiedTime, "The last modification time of the workload group (epoch).")
}

func (*GetWorkloadGroup) Invoke(ctx context.Context, req infer.FunctionRequest[GetWorkloadGroupArgs]) (infer.FunctionResponse[GetWorkloadGroupResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetWorkloadGroupResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	var resp *workloadgroups.WorkloadGroup
	if req.Input.Id != nil {
		all, err := workloadgroups.GetAll(ctx, svc)
		if err != nil {
			return infer.FunctionResponse[GetWorkloadGroupResult]{}, err
		}
		for i := range all {
			if all[i].ID == *req.Input.Id {
				resp = &all[i]
				break
			}
		}
	}
	if resp == nil && req.Input.Name != nil && *req.Input.Name != "" {
		r, err := workloadgroups.GetByName(ctx, svc, *req.Input.Name)
		if err != nil {
			return infer.FunctionResponse[GetWorkloadGroupResult]{}, err
		}
		resp = r
	}

	if resp == nil {
		return infer.FunctionResponse[GetWorkloadGroupResult]{}, fmt.Errorf("couldn't find any workload group with id %v or name %v", req.Input.Id, ptrToString(req.Input.Name))
	}

	result := GetWorkloadGroupResult{
		Id:               resp.ID,
		Name:             resp.Name,
		Description:      resp.Description,
		Expression:       resp.Expression,
		LastModifiedTime: resp.LastModifiedTime,
	}
	if len(resp.WorkloadTagExpression.ExpressionContainers) > 0 {
		result.ExpressionJson = flattenWorkloadGroupExpressionJson(&resp.WorkloadTagExpression)
	}
	return infer.FunctionResponse[GetWorkloadGroupResult]{Output: result}, nil
}
