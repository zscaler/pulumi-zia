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

// Package provider implements the Sandbox Rules resource.
// Adopted from terraform-provider-zia resource_zia_sandbox_rules.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/sandbox/sandbox_rules"
)

const sandboxResourceType = "sandbox_rules"

// SandboxRule implements the zia:index:SandboxRule resource.
type SandboxRule struct{}

// SandboxRuleArgs are the inputs.
type SandboxRuleArgs struct {
	Name               string               `pulumi:"name"`
	Order              int                  `pulumi:"order"`
	Description        *string              `pulumi:"description,optional"`
	Rank               *int                 `pulumi:"rank,optional"`
	State              *string              `pulumi:"state,optional"`
	BaRuleAction       *string              `pulumi:"baRuleAction,optional"`
	FirstTimeEnable    *bool                `pulumi:"firstTimeEnable,optional"`
	FirstTimeOperation *string              `pulumi:"firstTimeOperation,optional"`
	MLActionEnabled    *bool                `pulumi:"mlActionEnabled,optional"`
	ByThreatScore      *int                 `pulumi:"byThreatScore,optional"`
	Locations          []int                `pulumi:"locations,optional"`
	LocationGroups     []int                `pulumi:"locationGroups,optional"`
	Departments        []int                `pulumi:"departments,optional"`
	Groups             []int                `pulumi:"groups,optional"`
	Users              []int                `pulumi:"users,optional"`
	Labels             []int                `pulumi:"labels,optional"`
	URLCategories      []string             `pulumi:"urlCategories,optional"`
	BaPolicyCategories []string             `pulumi:"baPolicyCategories,optional"`
	FileTypes          []string             `pulumi:"fileTypes,optional"`
	Protocols          []string             `pulumi:"protocols,optional"`
	ZPAAppSegments     []ZPAAppSegmentInput `pulumi:"zpaAppSegments,optional"`
}

// SandboxRuleState is the persisted state.
type SandboxRuleState struct {
	SandboxRuleArgs
	RuleID *int `pulumi:"ruleId"`
}

func filterOutDefaultSandboxRule(rules []sandbox_rules.SandboxRules) []sandbox_rules.SandboxRules {
	var out []sandbox_rules.SandboxRules
	for _, r := range rules {
		if r.Order != 127 && r.Name != "Default BA Rule" {
			out = append(out, r)
		}
	}
	return out
}

func sandboxRuleArgsToAPI(args *SandboxRuleArgs, id int) sandbox_rules.SandboxRules {
	order := args.Order
	if order == 0 {
		order = 1
	}
	rank := ptrToIntDefault(args.Rank, 7)
	state := ptrToString(args.State)
	if state == "" {
		state = "ENABLED"
	}
	baRuleAction := ptrToString(args.BaRuleAction)
	if baRuleAction == "" {
		baRuleAction = "ALLOW"
	}
	firstTimeOperation := ptrToString(args.FirstTimeOperation)
	if firstTimeOperation == "" {
		firstTimeOperation = "ALLOW_SCAN"
	}
	api := sandbox_rules.SandboxRules{
		ID:                 id,
		Name:               args.Name,
		Order:              order,
		Rank:               rank,
		State:              state,
		Description:        ptrToString(args.Description),
		BaRuleAction:       baRuleAction,
		FirstTimeEnable:    ptrToBool(args.FirstTimeEnable),
		FirstTimeOperation: firstTimeOperation,
		MLActionEnabled:    ptrToBool(args.MLActionEnabled),
		ByThreatScore:      ptrToIntDefault(args.ByThreatScore, 0),
		URLCategories:      args.URLCategories,
		BaPolicyCategories: args.BaPolicyCategories,
		FileTypes:          args.FileTypes,
		Protocols:          args.Protocols,
		Locations:          idsToIDNameExtensions(args.Locations),
		LocationGroups:     idsToIDNameExtensions(args.LocationGroups),
		Groups:             idsToIDNameExtensions(args.Groups),
		Departments:        idsToIDNameExtensions(args.Departments),
		Users:              idsToIDNameExtensions(args.Users),
		Labels:             idsToIDNameExtensions(args.Labels),
	}
	if len(args.ZPAAppSegments) > 0 {
		api.ZPAAppSegments = expandZPAAppSegments(args.ZPAAppSegments)
	}
	return api
}

func sandboxRuleAPIToState(api *sandbox_rules.SandboxRules) SandboxRuleState {
	state := SandboxRuleState{
		SandboxRuleArgs: SandboxRuleArgs{
			Name:               api.Name,
			Order:              api.Order,
			Description:        stringPtr(api.Description),
			Rank:               intPtr(api.Rank),
			State:              stringPtr(api.State),
			BaRuleAction:       stringPtr(api.BaRuleAction),
			FirstTimeEnable:    boolPtr(api.FirstTimeEnable),
			FirstTimeOperation: stringPtr(api.FirstTimeOperation),
			MLActionEnabled:    boolPtr(api.MLActionEnabled),
			ByThreatScore:      intPtr(api.ByThreatScore),
			URLCategories:      api.URLCategories,
			BaPolicyCategories: api.BaPolicyCategories,
			FileTypes:          api.FileTypes,
			Protocols:          api.Protocols,
			Locations:          idNameExtensionsToIDs(api.Locations),
			LocationGroups:     idNameExtensionsToIDs(api.LocationGroups),
			Groups:             idNameExtensionsToIDs(api.Groups),
			Departments:        idNameExtensionsToIDs(api.Departments),
			Users:              idNameExtensionsToIDs(api.Users),
			Labels:             idNameExtensionsToIDs(api.Labels),
			ZPAAppSegments:     flattenZPAAppSegments(api.ZPAAppSegments),
		},
		RuleID: intPtr(api.ID),
	}
	return state
}

func (SandboxRule) Create(ctx context.Context, req infer.CreateRequest[SandboxRuleArgs]) (infer.CreateResponse[SandboxRuleState], error) {
	if req.DryRun {
		s := SandboxRuleState{SandboxRuleArgs: req.Inputs, RuleID: intPtr(0)}
		return infer.CreateResponse[SandboxRuleState]{ID: "preview", Output: s}, nil
	}
	if req.Inputs.Order < 1 {
		return infer.CreateResponse[SandboxRuleState]{}, fmt.Errorf("order must be a positive whole number (>= 1), got %d", req.Inputs.Order)
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[SandboxRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := sandboxRuleArgsToAPI(&req.Inputs, 0)
	timeout := 60 * time.Minute
	start := time.Now()

	intendedOrder := apiReq.Order
	intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)

	for {
		select {
		case sandboxSem <- struct{}{}:
		case <-ctx.Done():
			return infer.CreateResponse[SandboxRuleState]{}, fmt.Errorf("creating sandbox rule: %w", ctx.Err())
		}

		sandboxOrderMu.Lock()
		if sandboxStartingOrder == 0 {
			list, _ := sandbox_rules.GetAll(ctx, svc)
			for _, r := range list {
				if r.Order == 127 || r.Name == "Default BA Rule" {
					continue
				}
				if r.Order > sandboxStartingOrder {
					sandboxStartingOrder = r.Order
				}
			}
			if sandboxStartingOrder == 0 {
				sandboxStartingOrder = 1
			}
		}

		if intendedRank < 7 {
			apiReq.Rank = 7
		}
		apiReq.Order = sandboxStartingOrder
		sandboxOrderMu.Unlock()

		resp, err := sandbox_rules.Create(ctx, svc, &apiReq)

		if err == nil {
			sandboxOrderMu.Lock()
			sandboxStartingOrder++
			sandboxOrderMu.Unlock()
		}

		<-sandboxSem

		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.CreateResponse[SandboxRuleState]{}, customErr
		}
		if err != nil {
			reg := regexp.MustCompile("Rule with rank [0-9]+ is not allowed at order [0-9]+")
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") && reg.MatchString(err.Error()) {
				return infer.CreateResponse[SandboxRuleState]{}, fmt.Errorf("error creating sandbox rule: %s, check order %d vs rank %d, err:%s",
					apiReq.Name, intendedOrder, intendedRank, err)
			}
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Creating sandbox rule name: %v, got INVALID_INPUT_ARGUMENT\n", apiReq.Name)
				sandboxOrderMu.Lock()
				sandboxStartingOrder = 0
				sandboxOrderMu.Unlock()
				if time.Since(start) < timeout {
					select {
					case <-time.After(10 * time.Second):
					case <-ctx.Done():
						return infer.CreateResponse[SandboxRuleState]{}, fmt.Errorf("creating sandbox rule: %w", ctx.Err())
					}
					continue
				}
			}
			return infer.CreateResponse[SandboxRuleState]{}, fmt.Errorf("creating sandbox rule: %w", err)
		}

		reorderWithBeforeReorder(
			OrderRule{Order: intendedOrder, Rank: intendedRank},
			resp.ID,
			sandboxResourceType,
			func() (int, error) {
				list, err := sandbox_rules.GetAll(ctx, svc)
				if err != nil {
					return 0, err
				}
				return len(filterOutDefaultSandboxRule(list)), nil
			},
			func(id int, order OrderRule) error {
				rule, err := sandbox_rules.Get(ctx, svc, id)
				if err != nil {
					return err
				}
				// to avoid the STALE_CONFIGURATION_ERROR
				rule.LastModifiedTime = 0
				rule.LastModifiedBy = nil
				// Strip read-only fields that cause "Request body is invalid" for predefined rules
				rule.DefaultRule = false
				rule.AccessControl = ""
				rule.Order = order.Order
				rule.Rank = order.Rank
				_, err = sandbox_rules.Update(ctx, svc, id, rule)
				return err
			},
			nil,
		)

		markOrderRuleAsDone(resp.ID, sandboxResourceType)
		waitForReorder(sandboxResourceType)

		if shouldActivate() {
			time.Sleep(2 * time.Second)
			if activationErr := triggerActivation(ctx, client); activationErr != nil {
				return infer.CreateResponse[SandboxRuleState]{}, activationErr
			}
		} else {
			log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
		}

		rule, err := sandbox_rules.Get(ctx, svc, resp.ID)
		if err != nil {
			state := SandboxRuleState{SandboxRuleArgs: req.Inputs, RuleID: intPtr(resp.ID)}
			return infer.CreateResponse[SandboxRuleState]{ID: strconv.Itoa(resp.ID), Output: state}, nil
		}
		return infer.CreateResponse[SandboxRuleState]{
			ID:     strconv.Itoa(resp.ID),
			Output: sandboxRuleAPIToState(rule),
		}, nil
	}
}

func (SandboxRule) Read(ctx context.Context, req infer.ReadRequest[SandboxRuleArgs, SandboxRuleState]) (infer.ReadResponse[SandboxRuleArgs, SandboxRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[SandboxRuleArgs, SandboxRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		rule, lookupErr := sandbox_rules.GetByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[SandboxRuleArgs, SandboxRuleState]{}, nil
			}
			return infer.ReadResponse[SandboxRuleArgs, SandboxRuleState]{}, lookupErr
		}
		id = rule.ID
	}

	rule, err := sandbox_rules.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[SandboxRuleArgs, SandboxRuleState]{}, nil
		}
		return infer.ReadResponse[SandboxRuleArgs, SandboxRuleState]{}, err
	}
	if rule.Order == 127 || rule.Name == "Default BA Rule" {
		return infer.ReadResponse[SandboxRuleArgs, SandboxRuleState]{}, nil
	}

	state := sandboxRuleAPIToState(rule)
	return infer.ReadResponse[SandboxRuleArgs, SandboxRuleState]{
		ID:     req.ID,
		Inputs: state.SandboxRuleArgs,
		State:  state,
	}, nil
}

func (SandboxRule) Update(ctx context.Context, req infer.UpdateRequest[SandboxRuleArgs, SandboxRuleState]) (infer.UpdateResponse[SandboxRuleState], error) {
	if req.Inputs.Order < 1 {
		return infer.UpdateResponse[SandboxRuleState]{}, fmt.Errorf("order must be a positive whole number (>= 1), got %d", req.Inputs.Order)
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[SandboxRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[SandboxRuleState]{}, fmt.Errorf("invalid sandbox rule ID: %s", req.ID)
	}
	apiReq := sandboxRuleArgsToAPI(&req.Inputs, id)

	existingRules, err := sandbox_rules.GetAll(ctx, svc)
	if err != nil {
		return infer.UpdateResponse[SandboxRuleState]{}, err
	}
	existingRules = filterOutDefaultSandboxRule(existingRules)
	sort.Slice(existingRules, func(i, j int) bool {
		return existingRules[i].Rank < existingRules[j].Rank || (existingRules[i].Rank == existingRules[j].Rank && existingRules[i].Order < existingRules[j].Order)
	})
	intendedOrder := apiReq.Order
	intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)
	nextAvailableOrder := existingRules[len(existingRules)-1].Order
	apiReq.Rank = 7
	apiReq.Order = nextAvailableOrder

	if _, err = sandbox_rules.Update(ctx, svc, id, &apiReq); err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.UpdateResponse[SandboxRuleState]{}, nil
		}
		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.UpdateResponse[SandboxRuleState]{}, customErr
		}
		return infer.UpdateResponse[SandboxRuleState]{}, err
	}

	reorderWithBeforeReorder(
		OrderRule{Order: intendedOrder, Rank: intendedRank},
		id,
		sandboxResourceType,
		func() (int, error) {
			list, err := sandbox_rules.GetAll(ctx, svc)
			if err != nil {
				return 0, err
			}
			return len(filterOutDefaultSandboxRule(list)), nil
		},
		func(ruleID int, order OrderRule) error {
			rule, err := sandbox_rules.Get(ctx, svc, ruleID)
			if err != nil {
				return err
			}
			// to avoid the STALE_CONFIGURATION_ERROR
			rule.LastModifiedTime = 0
			rule.LastModifiedBy = nil
			// Strip read-only fields that cause "Request body is invalid" for predefined rules
			rule.DefaultRule = false
			rule.AccessControl = ""
			rule.Order = order.Order
			rule.Rank = order.Rank
			_, err = sandbox_rules.Update(ctx, svc, ruleID, rule)
			return err
		},
		nil,
	)

	markOrderRuleAsDone(id, sandboxResourceType)
	waitForReorder(sandboxResourceType)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[SandboxRuleState]{}, activationErr
		}
	}

	updated, err := sandbox_rules.Get(ctx, svc, id)
	if err != nil {
		return infer.UpdateResponse[SandboxRuleState]{Output: SandboxRuleState{
			SandboxRuleArgs: req.Inputs,
			RuleID:          intPtr(id),
		}}, nil
	}
	return infer.UpdateResponse[SandboxRuleState]{Output: sandboxRuleAPIToState(updated)}, nil
}

func (SandboxRule) Delete(ctx context.Context, req infer.DeleteRequest[SandboxRuleState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid sandbox rule ID: %s", req.ID)
	}
	if _, err := sandbox_rules.Delete(ctx, svc, id); err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.DeleteResponse{}, nil
		}
		return infer.DeleteResponse{}, err
	}
	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}
	return infer.DeleteResponse{}, nil
}

func (SandboxRule) Annotate(a infer.Annotator) {
	describeResource(a, &SandboxRule{}, `The zia_sandbox_rules resource manages sandbox policy rules in the Zscaler Internet Access (ZIA) cloud service. Sandbox rules define actions for file analysis based on criteria such as locations, departments, groups, users, and file types.

For more information, see the [ZIA Cloud Sandbox documentation](https://help.zscaler.com/zia/about-cloud-sandbox-policies).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Sandbox Rule

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.SandboxRule("example", {
    name: "Example Sandbox Rule",
    description: "Block suspicious file types",
    order: 1,
    state: "ENABLED",
    baRuleAction: "ALLOW",
    fileTypes: ["ALL_OUTBOUND"],
    protocols: ["FTP_RULE", "SSL_RULE", "FOHTTP_RULE", "HTTP_PROXY"],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.SandboxRule("example",
    name="Example Sandbox Rule",
    description="Block suspicious file types",
    order=1,
    state="ENABLED",
    ba_rule_action="ALLOW",
    file_types=["ALL_OUTBOUND"],
    protocols=["FTP_RULE", "SSL_RULE", "FOHTTP_RULE", "HTTP_PROXY"],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:SandboxRule
    properties:
      name: Example Sandbox Rule
      description: Block suspicious file types
      order: 1
      state: ENABLED
      baRuleAction: ALLOW
      fileTypes:
        - ALL_OUTBOUND
      protocols:
        - FTP_RULE
        - SSL_RULE
        - FOHTTP_RULE
        - HTTP_PROXY
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing Sandbox Rule can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:SandboxRule example 12345
`+tripleBacktick("")+`
`)
}

func (a *SandboxRuleArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the sandbox rule. Must be unique.")
	ann.Describe(&a.Order, "The order of execution of the rule with respect to other sandbox rules.")
	ann.Describe(&a.Description, "Additional information about the sandbox rule.")
	ann.Describe(&a.Rank, "Admin rank of the sandbox policy rule. Valid values: 0-7. Default: 7.")
	ann.Describe(&a.State, "Rule state. Valid values: `ENABLED`, `DISABLED`.")
	ann.Describe(&a.BaRuleAction, "The action applied when the rule is matched. Valid values: `ALLOW`, `BLOCK`, `QUARANTINE`.")
	ann.Describe(&a.FirstTimeEnable, "If set to true, a first-time action is enabled.")
	ann.Describe(&a.FirstTimeOperation, "The action for first-time file downloads. Valid values: `ALLOW_SCAN`, `QUARANTINE`.")
	ann.Describe(&a.MLActionEnabled, "If set to true, machine learning-based analysis action is enabled.")
	ann.Describe(&a.ByThreatScore, "Threat score threshold for the rule. Files with a score above this value trigger the action.")
	ann.Describe(&a.Locations, "IDs of locations to which the rule applies.")
	ann.Describe(&a.LocationGroups, "IDs of location groups to which the rule applies.")
	ann.Describe(&a.Departments, "IDs of departments to which the rule applies.")
	ann.Describe(&a.Groups, "IDs of groups to which the rule applies.")
	ann.Describe(&a.Users, "IDs of users to which the rule applies.")
	ann.Describe(&a.Labels, "IDs of labels associated with the rule.")
	ann.Describe(&a.URLCategories, "List of URL categories to which the rule applies.")
	ann.Describe(&a.BaPolicyCategories, "List of behavioral analysis policy categories.")
	ann.Describe(&a.FileTypes, "List of file types for which the rule applies (e.g., `ALL_OUTBOUND`, `EXE`, `DLL`).")
	ann.Describe(&a.Protocols, "Protocols to which the rule applies. Valid values: `FTP_RULE`, `SSL_RULE`, `FOHTTP_RULE`, `HTTP_PROXY`.")
	ann.Describe(&a.ZPAAppSegments, "List of ZPA application segments to which the rule applies.")
}

func (s *SandboxRuleState) Annotate(a infer.Annotator) {
	a.Describe(&s.RuleID, "The system-generated ID of the sandbox rule.")
}

var _ infer.CustomResource[SandboxRuleArgs, SandboxRuleState] = SandboxRule{}
