// Copyright (c) 2023 Zscaler Technology Alliances, <zscaler-partner-labs@z-bd.com>
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

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/saas_security_api/casb_dlp_rules"
)

const casbDlpResourceType = "casb_dlp_rules"

type CasbDlpRule struct{}

type CasbDlpRuleArgs struct {
	Name                         string   `pulumi:"name"`
	Type                         string   `pulumi:"type"`
	Order                        int      `pulumi:"order"`
	Description                  *string  `pulumi:"description,optional"`
	State                        *string  `pulumi:"state,optional"`
	Rank                         *int     `pulumi:"rank,optional"`
	Action                       *string  `pulumi:"action,optional"`
	Severity                     *string  `pulumi:"severity,optional"`
	ContentLocation              *string  `pulumi:"contentLocation,optional"`
	Recipient                    *string  `pulumi:"recipient,optional"`
	QuarantineLocation           *string  `pulumi:"quarantineLocation,optional"`
	WatermarkDeleteOldVersion    *bool    `pulumi:"watermarkDeleteOldVersion,optional"`
	IncludeCriteriaDomainProfile *bool    `pulumi:"includeCriteriaDomainProfile,optional"`
	IncludeEmailRecipientProfile *bool    `pulumi:"includeEmailRecipientProfile,optional"`
	WithoutContentInspection     *bool    `pulumi:"withoutContentInspection,optional"`
	IncludeEntityGroups          *bool    `pulumi:"includeEntityGroups,optional"`
	ExternalAuditorEmail         *string  `pulumi:"externalAuditorEmail,optional"`
	BucketOwner                  *string  `pulumi:"bucketOwner,optional"`
	Domains                      []string `pulumi:"domains,optional"`
	Components                    []string `pulumi:"components,optional"`
	CollaborationScope           []string `pulumi:"collaborationScope,optional"`
	FileTypes                    []string `pulumi:"fileTypes,optional"`
	Groups                       []int    `pulumi:"groups,optional"`
	Departments                  []int    `pulumi:"departments,optional"`
	Users                        []int    `pulumi:"users,optional"`
	Labels                       []int    `pulumi:"labels,optional"`
	CloudAppTenants              []int    `pulumi:"cloudAppTenants,optional"`
	EntityGroups                 []int    `pulumi:"entityGroups,optional"`
	IncludedDomainProfiles       []int    `pulumi:"includedDomainProfiles,optional"`
	ExcludedDomainProfiles       []int    `pulumi:"excludedDomainProfiles,optional"`
	CriteriaDomainProfiles       []int    `pulumi:"criteriaDomainProfiles,optional"`
	EmailRecipientProfiles       []int    `pulumi:"emailRecipientProfiles,optional"`
	ObjectTypes                  []int    `pulumi:"objectTypes,optional"`
	Buckets                      []int    `pulumi:"buckets,optional"`
	DlpEngines                   []int    `pulumi:"dlpEngines,optional"`
}

type CasbDlpRuleState struct {
	CasbDlpRuleArgs
	RuleID *int `pulumi:"ruleId"`
}

func casbDlpRuleArgsToAPI(args *CasbDlpRuleArgs, id int) casb_dlp_rules.CasbDLPRules {
	order := args.Order
	if order == 0 {
		order = 1
	}
	return casb_dlp_rules.CasbDLPRules{
		ID:                           id,
		Name:                         args.Name,
		Type:                         args.Type,
		Description:                  ptrToString(args.Description),
		Order:                        order,
		Rank:                         ptrToIntDefault(args.Rank, 7),
		State:                        ptrToStringDefault(args.State, "ENABLED"),
		Action:                       ptrToString(args.Action),
		Severity:                     ptrToString(args.Severity),
		ContentLocation:              ptrToString(args.ContentLocation),
		Recipient:                    ptrToString(args.Recipient),
		QuarantineLocation:           ptrToString(args.QuarantineLocation),
		WatermarkDeleteOldVersion:    ptrToBool(args.WatermarkDeleteOldVersion),
		IncludeCriteriaDomainProfile: ptrToBool(args.IncludeCriteriaDomainProfile),
		IncludeEmailRecipientProfile: ptrToBool(args.IncludeEmailRecipientProfile),
		WithoutContentInspection:     ptrToBool(args.WithoutContentInspection),
		IncludeEntityGroups:          ptrToBool(args.IncludeEntityGroups),
		ExternalAuditorEmail:         ptrToString(args.ExternalAuditorEmail),
		BucketOwner:                  ptrToString(args.BucketOwner),
		Domains:                      args.Domains,
		Components:                   args.Components,
		CollaborationScope:           args.CollaborationScope,
		FileTypes:                    args.FileTypes,
		Groups:                       idsToIDNameExtensions(args.Groups),
		Departments:                  idsToIDNameExtensions(args.Departments),
		Users:                        idsToIDNameExtensions(args.Users),
		Labels:                       idsToIDNameExtensions(args.Labels),
		CloudAppTenants:              idsToIDNameExtensions(args.CloudAppTenants),
		EntityGroups:                 idsToIDNameExtensions(args.EntityGroups),
		IncludedDomainProfiles:       idsToIDNameExtensions(args.IncludedDomainProfiles),
		ExcludedDomainProfiles:       idsToIDNameExtensions(args.ExcludedDomainProfiles),
		CriteriaDomainProfiles:       idsToIDNameExtensions(args.CriteriaDomainProfiles),
		EmailRecipientProfiles:       idsToIDNameExtensions(args.EmailRecipientProfiles),
		ObjectTypes:                  idsToIDNameExtensions(args.ObjectTypes),
		Buckets:                      idsToIDNameExtensions(args.Buckets),
		DLPEngines:                   idsToIDNameExtensions(args.DlpEngines),
	}
}

func casbDlpRuleAPIToState(api *casb_dlp_rules.CasbDLPRules) CasbDlpRuleState {
	return CasbDlpRuleState{
		CasbDlpRuleArgs: CasbDlpRuleArgs{
			Name:                         api.Name,
			Type:                         api.Type,
			Order:                        api.Order,
			Description:                  stringPtr(api.Description),
			State:                        stringPtr(api.State),
			Rank:                         intPtr(api.Rank),
			Action:                       stringPtr(api.Action),
			Severity:                     stringPtr(api.Severity),
			ContentLocation:              stringPtr(api.ContentLocation),
			Recipient:                    stringPtr(api.Recipient),
			QuarantineLocation:           stringPtr(api.QuarantineLocation),
			WatermarkDeleteOldVersion:    boolPtr(api.WatermarkDeleteOldVersion),
			IncludeCriteriaDomainProfile: boolPtr(api.IncludeCriteriaDomainProfile),
			IncludeEmailRecipientProfile: boolPtr(api.IncludeEmailRecipientProfile),
			WithoutContentInspection:     boolPtr(api.WithoutContentInspection),
			IncludeEntityGroups:          boolPtr(api.IncludeEntityGroups),
			ExternalAuditorEmail:         stringPtr(api.ExternalAuditorEmail),
			BucketOwner:                  stringPtr(api.BucketOwner),
			Domains:                      api.Domains,
			Components:                   api.Components,
			CollaborationScope:           api.CollaborationScope,
			FileTypes:                    api.FileTypes,
			Groups:                       idNameExtensionsToIDs(api.Groups),
			Departments:                  idNameExtensionsToIDs(api.Departments),
			Users:                        idNameExtensionsToIDs(api.Users),
			Labels:                       idNameExtensionsToIDs(api.Labels),
			CloudAppTenants:              idNameExtensionsToIDs(api.CloudAppTenants),
			EntityGroups:                 idNameExtensionsToIDs(api.EntityGroups),
			IncludedDomainProfiles:       idNameExtensionsToIDs(api.IncludedDomainProfiles),
			ExcludedDomainProfiles:       idNameExtensionsToIDs(api.ExcludedDomainProfiles),
			CriteriaDomainProfiles:       idNameExtensionsToIDs(api.CriteriaDomainProfiles),
			EmailRecipientProfiles:       idNameExtensionsToIDs(api.EmailRecipientProfiles),
			ObjectTypes:                  idNameExtensionsToIDs(api.ObjectTypes),
			Buckets:                      idNameExtensionsToIDs(api.Buckets),
			DlpEngines:                   idNameExtensionsToIDs(api.DLPEngines),
		},
		RuleID: intPtr(api.ID),
	}
}

func (CasbDlpRule) Create(ctx context.Context, req infer.CreateRequest[CasbDlpRuleArgs]) (infer.CreateResponse[CasbDlpRuleState], error) {
	if req.DryRun {
		return infer.CreateResponse[CasbDlpRuleState]{ID: "preview", Output: CasbDlpRuleState{CasbDlpRuleArgs: req.Inputs, RuleID: intPtr(0)}}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[CasbDlpRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := casbDlpRuleArgsToAPI(&req.Inputs, 0)
	ruleType := apiReq.Type
	timeout := 60 * time.Minute
	start := time.Now()

	for {
		cloudCasbDlpRuleLock.Lock()
		if cloudCasbDlpRuleStartingOrder == 0 {
			rules, _ := casb_dlp_rules.GetByRuleType(ctx, svc, ruleType)
			for _, r := range rules {
				if r.Order > cloudCasbDlpRuleStartingOrder {
					cloudCasbDlpRuleStartingOrder = r.Order
				}
			}
			if cloudCasbDlpRuleStartingOrder == 0 {
				cloudCasbDlpRuleStartingOrder = 1
			}
		}
		cloudCasbDlpRuleLock.Unlock()

		intendedOrder := apiReq.Order
		apiReq.Order = cloudCasbDlpRuleStartingOrder

		resp, err := casb_dlp_rules.Create(ctx, svc, &apiReq)
		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.CreateResponse[CasbDlpRuleState]{}, customErr
		}
		if err != nil {
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Creating casb dlp rule name: %v, got INVALID_INPUT_ARGUMENT\n", apiReq.Name)
				if time.Since(start) < timeout {
					time.Sleep(10 * time.Second)
					continue
				}
			}
			return infer.CreateResponse[CasbDlpRuleState]{}, fmt.Errorf("creating casb dlp rule: %w", err)
		}

		reorderWithBeforeReorder(
			OrderRule{Order: intendedOrder, Rank: apiReq.Rank},
			resp.ID,
			casbDlpResourceType,
			func() (int, error) {
				rules, err := casb_dlp_rules.GetByRuleType(ctx, svc, ruleType)
				if err != nil {
					return 0, err
				}
				return len(rules), nil
			},
			func(id int, order OrderRule) error {
				rule, err := casb_dlp_rules.GetByRuleID(ctx, svc, ruleType, id)
				if err != nil {
					return err
				}
				if rule.Order == order.Order {
					return nil
				}
				rule.Order = order.Order
				_, err = casb_dlp_rules.Update(ctx, svc, id, rule)
				return err
			},
			nil,
		)
		markOrderRuleAsDone(resp.ID, casbDlpResourceType)
		waitForReorder(casbDlpResourceType)

		if shouldActivate() {
			time.Sleep(2 * time.Second)
			if activationErr := triggerActivation(ctx, client); activationErr != nil {
				return infer.CreateResponse[CasbDlpRuleState]{}, activationErr
			}
		} else {
			log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
		}

		rule, err := casb_dlp_rules.GetByRuleID(ctx, svc, ruleType, resp.ID)
		if err != nil {
			return infer.CreateResponse[CasbDlpRuleState]{ID: strconv.Itoa(resp.ID), Output: CasbDlpRuleState{CasbDlpRuleArgs: req.Inputs, RuleID: intPtr(resp.ID)}}, nil
		}
		return infer.CreateResponse[CasbDlpRuleState]{ID: strconv.Itoa(resp.ID), Output: casbDlpRuleAPIToState(rule)}, nil
	}
}

func (CasbDlpRule) Read(ctx context.Context, req infer.ReadRequest[CasbDlpRuleArgs, CasbDlpRuleState]) (infer.ReadResponse[CasbDlpRuleArgs, CasbDlpRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[CasbDlpRuleArgs, CasbDlpRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.ReadResponse[CasbDlpRuleArgs, CasbDlpRuleState]{}, fmt.Errorf("casb dlp rule ID must be numeric")
	}
	ruleType := req.State.Type
	if ruleType == "" {
		ruleType = req.Inputs.Type
	}
	if ruleType == "" {
		return infer.ReadResponse[CasbDlpRuleArgs, CasbDlpRuleState]{}, fmt.Errorf("rule type is required")
	}

	rule, err := casb_dlp_rules.GetByRuleID(ctx, svc, ruleType, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[CasbDlpRuleArgs, CasbDlpRuleState]{}, fmt.Errorf("casb dlp rule not found")
		}
		return infer.ReadResponse[CasbDlpRuleArgs, CasbDlpRuleState]{}, err
	}

	state := casbDlpRuleAPIToState(rule)
	return infer.ReadResponse[CasbDlpRuleArgs, CasbDlpRuleState]{ID: req.ID, Inputs: state.CasbDlpRuleArgs, State: state}, nil
}

func (CasbDlpRule) Update(ctx context.Context, req infer.UpdateRequest[CasbDlpRuleArgs, CasbDlpRuleState]) (infer.UpdateResponse[CasbDlpRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[CasbDlpRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[CasbDlpRuleState]{}, fmt.Errorf("invalid casb dlp rule ID: %s", req.ID)
	}
	apiReq := casbDlpRuleArgsToAPI(&req.Inputs, id)
	ruleType := apiReq.Type
	timeout := 60 * time.Minute
	start := time.Now()

	for {
		_, err = casb_dlp_rules.Update(ctx, svc, id, &apiReq)
		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.UpdateResponse[CasbDlpRuleState]{}, customErr
		}
		if err != nil {
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Updating casb dlp rule ID: %v, got INVALID_INPUT_ARGUMENT\n", id)
				if time.Since(start) < timeout {
					time.Sleep(10 * time.Second)
					continue
				}
			}
			return infer.UpdateResponse[CasbDlpRuleState]{}, err
		}

		reorderWithBeforeReorder(
			OrderRule{Order: req.Inputs.Order, Rank: ptrToIntDefault(req.Inputs.Rank, 7)},
			id,
			casbDlpResourceType,
			func() (int, error) {
				rules, err := casb_dlp_rules.GetByRuleType(ctx, svc, ruleType)
				if err != nil {
					return 0, err
				}
				return len(rules), nil
			},
			func(ruleID int, order OrderRule) error {
				rule, err := casb_dlp_rules.GetByRuleID(ctx, svc, ruleType, ruleID)
				if err != nil {
					return err
				}
				if rule.Order == order.Order {
					return nil
				}
				rule.Order = order.Order
				_, err = casb_dlp_rules.Update(ctx, svc, ruleID, rule)
				return err
			},
			nil,
		)
		markOrderRuleAsDone(id, casbDlpResourceType)
		waitForReorder(casbDlpResourceType)
		break
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[CasbDlpRuleState]{}, activationErr
		}
	}

	updated, err := casb_dlp_rules.GetByRuleID(ctx, svc, ruleType, id)
	if err != nil {
		return infer.UpdateResponse[CasbDlpRuleState]{Output: CasbDlpRuleState{CasbDlpRuleArgs: req.Inputs, RuleID: intPtr(id)}}, nil
	}
	return infer.UpdateResponse[CasbDlpRuleState]{Output: casbDlpRuleAPIToState(updated)}, nil
}

func (CasbDlpRule) Delete(ctx context.Context, req infer.DeleteRequest[CasbDlpRuleState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid casb dlp rule ID: %s", req.ID)
	}
	ruleType := req.State.Type
	if ruleType == "" {
		return infer.DeleteResponse{}, fmt.Errorf("rule type is required for delete")
	}
	if _, err := casb_dlp_rules.Delete(ctx, svc, ruleType, id); err != nil {
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

func (CasbDlpRule) Annotate(a infer.Annotator) {
	describeResource(a, &CasbDlpRule{}, `The zia_casb_dlp_rules resource manages CASB (Cloud Access Security Broker) DLP rules in the Zscaler Internet Access (ZIA) cloud service. CASB DLP rules define data loss prevention policies for SaaS applications to protect sensitive data from unauthorized access or sharing.

{{% examples %}}
## Example Usage

{{% example %}}
### CASB DLP Rule

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.CasbDlpRule("example", {
    name: "Example CASB DLP Rule",
    type: "CASB_DLP",
    order: 1,
    state: "ENABLED",
    action: "BLOCK",
    severity: "HIGH",
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.CasbDlpRule("example",
    name="Example CASB DLP Rule",
    type="CASB_DLP",
    order=1,
    state="ENABLED",
    action="BLOCK",
    severity="HIGH",
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:CasbDlpRule
    properties:
      name: Example CASB DLP Rule
      type: CASB_DLP
      order: 1
      state: ENABLED
      action: BLOCK
      severity: HIGH
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing CASB DLP Rule can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:CasbDlpRule example 12345
`+tripleBacktick("")+`
`)
}

func (a *CasbDlpRuleArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the CASB DLP rule. Must be unique.")
	ann.Describe(&a.Type, "The rule type (e.g. `CASB_DLP`).")
	ann.Describe(&a.Order, "The order of execution of the rule with respect to other CASB DLP rules.")
	ann.Describe(&a.Description, "Additional information about the CASB DLP rule.")
	ann.Describe(&a.State, "Rule state. Valid values: `ENABLED`, `DISABLED`.")
	ann.Describe(&a.Rank, "Admin rank of the CASB DLP rule. Valid values: 0-7. Default: 7.")
	ann.Describe(&a.Action, "Action taken when the rule is matched (e.g. `BLOCK`, `ALLOW`, `QUARANTINE`).")
	ann.Describe(&a.Severity, "Severity level of the rule (e.g. `HIGH`, `MEDIUM`, `LOW`).")
	ann.Describe(&a.ContentLocation, "Content location scope for the rule.")
	ann.Describe(&a.Recipient, "Notification recipient.")
	ann.Describe(&a.QuarantineLocation, "Quarantine location for matched content.")
	ann.Describe(&a.WatermarkDeleteOldVersion, "Whether to delete old versions when watermarking.")
	ann.Describe(&a.IncludeCriteriaDomainProfile, "Whether to include criteria based on domain profiles.")
	ann.Describe(&a.IncludeEmailRecipientProfile, "Whether to include email recipient profile criteria.")
	ann.Describe(&a.WithoutContentInspection, "Whether the rule applies without content inspection.")
	ann.Describe(&a.IncludeEntityGroups, "Whether to include entity groups in the rule criteria.")
	ann.Describe(&a.ExternalAuditorEmail, "Email address of the external auditor.")
	ann.Describe(&a.BucketOwner, "The bucket owner identifier.")
	ann.Describe(&a.Domains, "List of domains for the rule.")
	ann.Describe(&a.Components, "List of components for the rule.")
	ann.Describe(&a.CollaborationScope, "Collaboration scope for the rule.")
	ann.Describe(&a.FileTypes, "List of file types the rule applies to.")
	ann.Describe(&a.Groups, "IDs of groups for which the rule applies.")
	ann.Describe(&a.Departments, "IDs of departments for which the rule applies.")
	ann.Describe(&a.Users, "IDs of users for which the rule applies.")
	ann.Describe(&a.Labels, "IDs of labels associated with the rule.")
	ann.Describe(&a.CloudAppTenants, "IDs of cloud application tenants.")
	ann.Describe(&a.EntityGroups, "IDs of entity groups.")
	ann.Describe(&a.IncludedDomainProfiles, "IDs of included domain profiles.")
	ann.Describe(&a.ExcludedDomainProfiles, "IDs of excluded domain profiles.")
	ann.Describe(&a.CriteriaDomainProfiles, "IDs of criteria-based domain profiles.")
	ann.Describe(&a.EmailRecipientProfiles, "IDs of email recipient profiles.")
	ann.Describe(&a.ObjectTypes, "IDs of object types.")
	ann.Describe(&a.Buckets, "IDs of buckets.")
	ann.Describe(&a.DlpEngines, "IDs of DLP engines.")
}

func (s *CasbDlpRuleState) Annotate(a infer.Annotator) {
	a.Describe(&s.RuleID, "The system-generated ID of the CASB DLP rule.")
}

var _ infer.CustomResource[CasbDlpRuleArgs, CasbDlpRuleState] = CasbDlpRule{}
