package provider

import (
	"context"
	"fmt"
	"strconv"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/pulumi-zia/provider/internal/zia"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/common"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/firewallpolicies/filteringrules"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/rule_labels"
)

// RuleLabel implements the zia:index:RuleLabel resource.
type RuleLabel struct{}

// RuleLabelArgs are the inputs for RuleLabel.
type RuleLabelArgs struct {
	Name        *string `pulumi:"name,optional"`
	Description *string `pulumi:"description,optional"`
}

// RuleLabelState is the persisted state for RuleLabel.
type RuleLabelState struct {
	RuleLabelArgs
	RuleLabelId *int `pulumi:"ruleLabelId"`
}

// Check validates and normalizes inputs.
func (RuleLabel) Check(ctx context.Context, req infer.CheckRequest) (infer.CheckResponse[RuleLabelArgs], error) {
	inputs, failures, err := infer.DefaultCheck[RuleLabelArgs](ctx, req.NewInputs)
	if err != nil {
		return infer.CheckResponse[RuleLabelArgs]{}, err
	}
	if len(failures) > 0 {
		return infer.CheckResponse[RuleLabelArgs]{Failures: failures}, nil
	}
	if inputs.Name != nil && len(*inputs.Name) > 255 {
		return infer.CheckResponse[RuleLabelArgs]{Failures: []p.CheckFailure{{
			Property: "name",
			Reason:   "name must be at most 255 characters",
		}}}, nil
	}
	if inputs.Description != nil && len(*inputs.Description) > 10240 {
		return infer.CheckResponse[RuleLabelArgs]{Failures: []p.CheckFailure{{
			Property: "description",
			Reason:   "description must be at most 10240 characters",
		}}}, nil
	}
	return infer.CheckResponse[RuleLabelArgs]{Inputs: inputs}, nil
}

// Create creates a new Rule Label.
func (RuleLabel) Create(ctx context.Context, req infer.CreateRequest[RuleLabelArgs]) (infer.CreateResponse[RuleLabelState], error) {
	// Preview/dry-run: return placeholder without calling API
	if req.DryRun {
		state := RuleLabelState{
			RuleLabelArgs: req.Inputs,
			RuleLabelId:   intPtr(0),
		}
		return infer.CreateResponse[RuleLabelState]{ID: "preview", Output: state}, nil
	}

	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[RuleLabelState]{}, fmt.Errorf("ZIA provider not configured; set clientId, clientSecret, vanityDomain (or use env vars)")
	}
	client := cfg.Client()
	svc := client.Service

	in := req.Inputs
	rl := rule_labels.RuleLabels{
		Name:        ptrToString(in.Name),
		Description: ptrToString(in.Description),
	}

	resp, _, err := rule_labels.Create(ctx, svc, &rl)
	if err != nil {
		return infer.CreateResponse[RuleLabelState]{}, fmt.Errorf("creating rule label: %w", err)
	}

	state := RuleLabelState{
		RuleLabelArgs: req.Inputs,
		RuleLabelId:   &resp.ID,
	}
	return infer.CreateResponse[RuleLabelState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

// Read fetches the Rule Label state from ZIA.
// Supports import by ID (e.g. "12345") or by name (e.g. "My Label Name").
func (RuleLabel) Read(ctx context.Context, req infer.ReadRequest[RuleLabelArgs, RuleLabelState]) (infer.ReadResponse[RuleLabelArgs, RuleLabelState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[RuleLabelArgs, RuleLabelState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		resp, lookupErr := rule_labels.GetRuleLabelByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[RuleLabelArgs, RuleLabelState]{}, nil
			}
			return infer.ReadResponse[RuleLabelArgs, RuleLabelState]{}, fmt.Errorf("reading rule label by name: %w", lookupErr)
		}
		id = resp.ID
	}

	resp, err := rule_labels.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[RuleLabelArgs, RuleLabelState]{}, nil
		}
		return infer.ReadResponse[RuleLabelArgs, RuleLabelState]{}, fmt.Errorf("reading rule label: %w", err)
	}

	args := RuleLabelArgs{
		Name:        stringPtr(resp.Name),
		Description: stringPtr(resp.Description),
	}
	state := RuleLabelState{
		RuleLabelArgs: args,
		RuleLabelId:   &resp.ID,
	}
	return infer.ReadResponse[RuleLabelArgs, RuleLabelState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

// Update updates an existing Rule Label.
func (RuleLabel) Update(ctx context.Context, req infer.UpdateRequest[RuleLabelArgs, RuleLabelState]) (infer.UpdateResponse[RuleLabelState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[RuleLabelState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[RuleLabelState]{}, fmt.Errorf("invalid rule label id: %w", err)
	}

	in := req.Inputs
	rl := rule_labels.RuleLabels{
		ID:          id,
		Name:        ptrToString(in.Name),
		Description: ptrToString(in.Description),
	}

	_, _, err = rule_labels.Update(ctx, svc, id, &rl)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.UpdateResponse[RuleLabelState]{}, nil
		}
		return infer.UpdateResponse[RuleLabelState]{}, fmt.Errorf("updating rule label: %w", err)
	}

	state := RuleLabelState{
		RuleLabelArgs: req.Inputs,
		RuleLabelId:   &id,
	}
	return infer.UpdateResponse[RuleLabelState]{Output: state}, nil
}

// Delete removes the Rule Label.
func (RuleLabel) Delete(ctx context.Context, req infer.DeleteRequest[RuleLabelState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid rule label id: %w", err)
	}

	// Detach from firewall rules before delete (mirrors Terraform)
	_ = detachRuleLabel(ctx, client, id)

	_, err = rule_labels.Delete(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.DeleteResponse{}, nil
		}
		return infer.DeleteResponse{}, fmt.Errorf("deleting rule label: %w", err)
	}
	return infer.DeleteResponse{}, nil
}

func (RuleLabel) Annotate(a infer.Annotator) {
	describeResource(a, &RuleLabel{}, `The zia_rule_label resource manages rule labels in the Zscaler Internet Access (ZIA) cloud service. Rule labels are used to tag and organize firewall filtering rules, URL filtering rules, and other policy rules.

For more information, see the [ZIA Rule Labels documentation](https://help.zscaler.com/zia/rule-labels).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Rule Label

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.RuleLabel("example", {
    name: "Example Rule Label",
    description: "Label for branch office rules",
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.RuleLabel("example",
    name="Example Rule Label",
    description="Label for branch office rules",
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:RuleLabel
    properties:
      name: Example Rule Label
      description: Label for branch office rules
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing rule label can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:RuleLabel example 12345
`+tripleBacktick("")+`
`)
}

func (a *RuleLabelArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the rule label. Maximum 255 characters.")
	ann.Describe(&a.Description, "Additional information about the rule label. Maximum 10240 characters.")
}

func (s *RuleLabelState) Annotate(a infer.Annotator) {
	a.Describe(&s.RuleLabelId, "The system-generated ID of the rule label.")
}

// detachRuleLabel removes the label from all firewall rules that reference it.
func detachRuleLabel(ctx context.Context, client *zia.Client, labelID int) error {
	svc := client.Service
	rules, err := filteringrules.GetAll(ctx, svc, nil)
	if err != nil {
		return err
	}
	for i := range rules {
		r := &rules[i]
		var newLabels []common.IDNameExtensions
		for _, l := range r.Labels {
			if l.ID != labelID {
				newLabels = append(newLabels, l)
			}
		}
		if len(newLabels) != len(r.Labels) {
			r.Labels = newLabels
			r.LastModifiedTime = 0
			r.LastModifiedBy = nil
			_, _ = filteringrules.Update(ctx, svc, r.ID, r)
		}
	}
	return nil
}
