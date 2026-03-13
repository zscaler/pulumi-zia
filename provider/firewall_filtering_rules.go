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

// Package provider implements the Firewall Filtering Rules resource.
// Uses filteringrules package (firewallpolicies/filteringrules).
// validateFirewallRule blocks delete for predefined one-click and default rules.

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
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/firewallpolicies/filteringrules"
)

const firewallFilteringResourceType = "firewall_filtering_rules"

// Names that block deletion (predefined one-click and default rules).
var firewallFilteringBlockedDeleteNames = map[string]bool{
	"Office 365 One Click Rule":       true,
	"UCaaS One Click Rule":            true,
	"Block All IPv6":                  true,
	"Block malicious IPs and domains": true,
	"Default Firewall Filtering Rule": true,
}

func validateFirewallFilteringRuleDelete(name string) error {
	if firewallFilteringBlockedDeleteNames[name] {
		return fmt.Errorf("deletion of predefined rule '%s' is not allowed", name)
	}
	return nil
}

// FirewallFilteringRule implements the zia:index:FirewallFilteringRule resource.
type FirewallFilteringRule struct{}

// FirewallFilteringRuleArgs are the inputs.
type FirewallFilteringRuleArgs struct {
	Name                string               `pulumi:"name"`
	Order               int                  `pulumi:"order"`
	Description         *string              `pulumi:"description,optional"`
	Rank                *int                 `pulumi:"rank,optional"`
	State               *string              `pulumi:"state,optional"`
	Action              *string              `pulumi:"action,optional"`
	EnableFullLogging   *bool                `pulumi:"enableFullLogging,optional"`
	DefaultRule         *bool                `pulumi:"defaultRule,optional"`
	Predefined          *bool                `pulumi:"predefined,optional"`
	ExcludeSrcCountries *bool                `pulumi:"excludeSrcCountries,optional"`
	SrcIps              []string             `pulumi:"srcIps,optional"`
	DestAddresses       []string             `pulumi:"destAddresses,optional"`
	DestIpCategories    []string             `pulumi:"destIpCategories,optional"`
	DestCountries       []string             `pulumi:"destCountries,optional"`
	SourceCountries     []string             `pulumi:"sourceCountries,optional"`
	NwApplications      []string             `pulumi:"nwApplications,optional"`
	Locations           []int                `pulumi:"locations,optional"`
	LocationGroups      []int                `pulumi:"locationGroups,optional"`
	Departments         []int                `pulumi:"departments,optional"`
	Groups              []int                `pulumi:"groups,optional"`
	Users               []int                `pulumi:"users,optional"`
	TimeWindows         []int                `pulumi:"timeWindows,optional"`
	NwApplicationGroups []int                `pulumi:"nwApplicationGroups,optional"`
	AppServices         []int                `pulumi:"appServices,optional"`
	AppServiceGroups    []int                `pulumi:"appServiceGroups,optional"`
	Labels              []int                `pulumi:"labels,optional"`
	SrcIpGroups         []int                `pulumi:"srcIpGroups,optional"`
	DestIpGroups        []int                `pulumi:"destIpGroups,optional"`
	NwServices          []int                `pulumi:"nwServices,optional"`
	NwServiceGroups     []int                `pulumi:"nwServiceGroups,optional"`
	DeviceGroups        []int                `pulumi:"deviceGroups,optional"`
	Devices             []int                `pulumi:"devices,optional"`
	DeviceTrustLevels   []string             `pulumi:"deviceTrustLevels,optional"`
	WorkloadGroups      []WorkloadGroupInput `pulumi:"workloadGroups,optional"`
	ZpaAppSegments      []ZPAAppSegmentInput `pulumi:"zpaAppSegments,optional"`
}

// FirewallFilteringRuleState is the persisted state.
type FirewallFilteringRuleState struct {
	FirewallFilteringRuleArgs
	RuleID *int `pulumi:"ruleId"`
}

func firewallFilteringRuleArgsToAPI(args *FirewallFilteringRuleArgs, id int) filteringrules.FirewallFilteringRules {
	order := args.Order
	if order == 0 {
		order = 1
	}
	rank := ptrToIntDefault(args.Rank, 7)
	state := ptrToString(args.State)
	if state == "" {
		state = "ENABLED"
	}
	return filteringrules.FirewallFilteringRules{
		ID:                  id,
		Name:                args.Name,
		Order:               order,
		Rank:                rank,
		State:               state,
		Action:              ptrToString(args.Action),
		Description:         ptrToString(args.Description),
		EnableFullLogging:   ptrToBool(args.EnableFullLogging),
		DefaultRule:         ptrToBool(args.DefaultRule),
		Predefined:          ptrToBool(args.Predefined),
		ExcludeSrcCountries: ptrToBool(args.ExcludeSrcCountries),
		SrcIps:              args.SrcIps,
		DestAddresses:       args.DestAddresses,
		DestIpCategories:    args.DestIpCategories,
		DestCountries:       processCountries(args.DestCountries),
		SourceCountries:     processCountries(args.SourceCountries),
		NwApplications:      args.NwApplications,
		Locations:           idsToIDNameExtensions(args.Locations),
		LocationsGroups:     idsToIDNameExtensions(args.LocationGroups),
		Departments:         idsToIDNameExtensions(args.Departments),
		Groups:              idsToIDNameExtensions(args.Groups),
		Users:               idsToIDNameExtensions(args.Users),
		TimeWindows:         idsToIDNameExtensions(args.TimeWindows),
		NwApplicationGroups: idsToIDNameExtensions(args.NwApplicationGroups),
		AppServices:         idsToIDNameExtensions(args.AppServices),
		AppServiceGroups:    idsToIDNameExtensions(args.AppServiceGroups),
		Labels:              idsToIDNameExtensions(args.Labels),
		SrcIpGroups:         idsToIDNameExtensions(args.SrcIpGroups),
		DestIpGroups:        idsToIDNameExtensions(args.DestIpGroups),
		NwServices:          idsToIDNameExtensions(args.NwServices),
		NwServiceGroups:     idsToIDNameExtensions(args.NwServiceGroups),
		DeviceGroups:        idsToIDNameExtensions(args.DeviceGroups),
		Devices:             idsToIDNameExtensions(args.Devices),
		DeviceTrustLevels:   args.DeviceTrustLevels,
		WorkloadGroups:      expandWorkloadGroups(args.WorkloadGroups),
		ZPAAppSegments:      expandZPAAppSegments(args.ZpaAppSegments),
	}
}

func firewallFilteringRuleAPIToState(api *filteringrules.FirewallFilteringRules) FirewallFilteringRuleState {
	destCountries := processCountriesFromAPI(api.DestCountries)
	srcCountries := processCountriesFromAPI(api.SourceCountries)
	return FirewallFilteringRuleState{
		FirewallFilteringRuleArgs: FirewallFilteringRuleArgs{
			Name:                api.Name,
			Order:               api.Order,
			Description:         stringPtr(api.Description),
			Rank:                intPtr(api.Rank),
			State:               stringPtr(api.State),
			Action:              stringPtr(api.Action),
			EnableFullLogging:   boolPtr(api.EnableFullLogging),
			DefaultRule:         boolPtr(api.DefaultRule),
			Predefined:          boolPtr(api.Predefined),
			ExcludeSrcCountries: boolPtr(api.ExcludeSrcCountries),
			SrcIps:              api.SrcIps,
			DestAddresses:       api.DestAddresses,
			DestIpCategories:    api.DestIpCategories,
			DestCountries:       destCountries,
			SourceCountries:     srcCountries,
			NwApplications:      api.NwApplications,
			Locations:           idNameExtensionsToIDs(api.Locations),
			LocationGroups:      idNameExtensionsToIDs(api.LocationsGroups),
			Departments:         idNameExtensionsToIDs(api.Departments),
			Groups:              idNameExtensionsToIDs(api.Groups),
			Users:               idNameExtensionsToIDs(api.Users),
			TimeWindows:         idNameExtensionsToIDs(api.TimeWindows),
			NwApplicationGroups: idNameExtensionsToIDs(api.NwApplicationGroups),
			AppServices:         idNameExtensionsToIDs(api.AppServices),
			AppServiceGroups:    idNameExtensionsToIDs(api.AppServiceGroups),
			Labels:              idNameExtensionsToIDs(api.Labels),
			SrcIpGroups:         idNameExtensionsToIDs(api.SrcIpGroups),
			DestIpGroups:        idNameExtensionsToIDs(api.DestIpGroups),
			NwServices:          idNameExtensionsToIDs(api.NwServices),
			NwServiceGroups:     idNameExtensionsToIDs(api.NwServiceGroups),
			DeviceGroups:        idNameExtensionsToIDs(api.DeviceGroups),
			Devices:             idNameExtensionsToIDs(api.Devices),
			DeviceTrustLevels:   api.DeviceTrustLevels,
			WorkloadGroups:      workloadGroupOutputsToInputs(flattenWorkloadGroups(api.WorkloadGroups)),
			ZpaAppSegments:      flattenZPAAppSegments(api.ZPAAppSegments),
		},
		RuleID: intPtr(api.ID),
	}
}

func (FirewallFilteringRule) Diff(ctx context.Context, req infer.DiffRequest[FirewallFilteringRuleArgs, FirewallFilteringRuleState]) (infer.DiffResponse, error) {
	diff, hasChanges := diffSkippingUnsetInputs(req.State.FirewallFilteringRuleArgs, req.Inputs)
	return infer.DiffResponse{HasChanges: hasChanges, DetailedDiff: diff}, nil
}

func (FirewallFilteringRule) Create(ctx context.Context, req infer.CreateRequest[FirewallFilteringRuleArgs]) (infer.CreateResponse[FirewallFilteringRuleState], error) {
	if req.DryRun {
		s := FirewallFilteringRuleState{FirewallFilteringRuleArgs: req.Inputs, RuleID: intPtr(0)}
		return infer.CreateResponse[FirewallFilteringRuleState]{ID: "preview", Output: s}, nil
	}
	if req.Inputs.Order < 1 {
		return infer.CreateResponse[FirewallFilteringRuleState]{}, fmt.Errorf("order must be a positive whole number (>= 1), got %d", req.Inputs.Order)
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[FirewallFilteringRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := firewallFilteringRuleArgsToAPI(&req.Inputs, 0)
	timeout := 60 * time.Minute
	start := time.Now()

	intendedOrder := apiReq.Order
	intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)

	for {
		select {
		case firewallFilteringSem <- struct{}{}:
		case <-ctx.Done():
			return infer.CreateResponse[FirewallFilteringRuleState]{}, fmt.Errorf("creating firewall filtering rule: %w", ctx.Err())
		}

		firewallFilteringOrderMu.Lock()
		if firewallFilteringStartingOrder == 0 {
			list, _ := filteringrules.GetAll(ctx, svc, nil)
			for _, r := range list {
				if r.Order > firewallFilteringStartingOrder {
					firewallFilteringStartingOrder = r.Order
				}
			}
			if firewallFilteringStartingOrder == 0 {
				firewallFilteringStartingOrder = 1
			} else {
				firewallFilteringStartingOrder++
			}
		}

		if intendedRank < 7 {
			apiReq.Rank = 7
		}
		apiReq.Order = firewallFilteringStartingOrder
		firewallFilteringOrderMu.Unlock()

		resp, err := filteringrules.Create(ctx, svc, &apiReq)

		if err == nil {
			firewallFilteringOrderMu.Lock()
			firewallFilteringStartingOrder++
			firewallFilteringOrderMu.Unlock()
		}

		<-firewallFilteringSem

		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.CreateResponse[FirewallFilteringRuleState]{}, customErr
		}
		if err != nil {
			reg := regexp.MustCompile("Rule with rank [0-9]+ is not allowed at order [0-9]+")
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") && reg.MatchString(err.Error()) {
				return infer.CreateResponse[FirewallFilteringRuleState]{}, fmt.Errorf("error creating firewall filtering rule: %s, check order %d vs rank %d, err:%s",
					apiReq.Name, intendedOrder, intendedRank, err)
			}
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Creating firewall filtering rule name: %v, got INVALID_INPUT_ARGUMENT\n", apiReq.Name)
				firewallFilteringOrderMu.Lock()
				firewallFilteringStartingOrder = 0
				firewallFilteringOrderMu.Unlock()
				if time.Since(start) < timeout {
					select {
					case <-time.After(10 * time.Second):
					case <-ctx.Done():
						return infer.CreateResponse[FirewallFilteringRuleState]{}, fmt.Errorf("creating firewall filtering rule: %w", ctx.Err())
					}
					continue
				}
			}
			return infer.CreateResponse[FirewallFilteringRuleState]{}, fmt.Errorf("creating firewall filtering rule: %w", err)
		}

		reorderWithBeforeReorder(
			OrderRule{Order: intendedOrder, Rank: intendedRank},
			resp.ID,
			firewallFilteringResourceType,
			func() (int, error) {
				allRules, err := filteringrules.GetAll(ctx, svc, nil)
				if err != nil {
					return 0, err
				}
				return len(allRules), nil
			},
			func(id int, order OrderRule) error {
				rule, err := filteringrules.Get(ctx, svc, id)
				if err != nil {
					return err
				}
				// to avoid the STALE_CONFIGURATION_ERROR
				rule.LastModifiedTime = 0
				rule.LastModifiedBy = nil
				// Strip read-only fields that cause "Request body is invalid" for predefined rules
				rule.Predefined = false
				rule.DefaultRule = false
				rule.AccessControl = ""
				rule.Order = order.Order
				rule.Rank = order.Rank
				_, err = filteringrules.Update(ctx, svc, id, rule)
				return err
			},
			nil,
		)

		markOrderRuleAsDone(resp.ID, firewallFilteringResourceType)
		waitForReorder(firewallFilteringResourceType)

		if shouldActivate() {
			time.Sleep(2 * time.Second)
			if activationErr := triggerActivation(ctx, client); activationErr != nil {
				return infer.CreateResponse[FirewallFilteringRuleState]{}, activationErr
			}
		} else {
			log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
		}

		rule, err := filteringrules.Get(ctx, svc, resp.ID)
		if err != nil {
			state := FirewallFilteringRuleState{FirewallFilteringRuleArgs: req.Inputs, RuleID: intPtr(resp.ID)}
			return infer.CreateResponse[FirewallFilteringRuleState]{ID: strconv.Itoa(resp.ID), Output: state}, nil
		}
		return infer.CreateResponse[FirewallFilteringRuleState]{
			ID:     strconv.Itoa(resp.ID),
			Output: firewallFilteringRuleAPIToState(rule),
		}, nil
	}
}

func (FirewallFilteringRule) Read(ctx context.Context, req infer.ReadRequest[FirewallFilteringRuleArgs, FirewallFilteringRuleState]) (infer.ReadResponse[FirewallFilteringRuleArgs, FirewallFilteringRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[FirewallFilteringRuleArgs, FirewallFilteringRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		rule, lookupErr := filteringrules.GetByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[FirewallFilteringRuleArgs, FirewallFilteringRuleState]{}, nil
			}
			return infer.ReadResponse[FirewallFilteringRuleArgs, FirewallFilteringRuleState]{}, lookupErr
		}
		id = rule.ID
	}

	rule, err := filteringrules.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[FirewallFilteringRuleArgs, FirewallFilteringRuleState]{}, nil
		}
		return infer.ReadResponse[FirewallFilteringRuleArgs, FirewallFilteringRuleState]{}, err
	}

	state := firewallFilteringRuleAPIToState(rule)
	return infer.ReadResponse[FirewallFilteringRuleArgs, FirewallFilteringRuleState]{
		ID:     req.ID,
		Inputs: state.FirewallFilteringRuleArgs,
		State:  state,
	}, nil
}

func (FirewallFilteringRule) Update(ctx context.Context, req infer.UpdateRequest[FirewallFilteringRuleArgs, FirewallFilteringRuleState]) (infer.UpdateResponse[FirewallFilteringRuleState], error) {
	if req.Inputs.Order < 1 {
		return infer.UpdateResponse[FirewallFilteringRuleState]{}, fmt.Errorf("order must be a positive whole number (>= 1), got %d", req.Inputs.Order)
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[FirewallFilteringRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[FirewallFilteringRuleState]{}, fmt.Errorf("invalid firewall filtering rule ID: %s", req.ID)
	}
	apiReq := firewallFilteringRuleArgsToAPI(&req.Inputs, id)

	// Store intended order/rank from inputs before overwriting (matches Terraform)
	intendedOrder := apiReq.Order
	intendedRank := apiReq.Rank

	existingRules, err := filteringrules.GetAll(ctx, svc, nil)
	if err == nil && len(existingRules) > 0 {
		sort.Slice(existingRules, func(i, j int) bool {
			return existingRules[i].Rank < existingRules[j].Rank || (existingRules[i].Rank == existingRules[j].Rank && existingRules[i].Order < existingRules[j].Order)
		})
		apiReq.Rank = 7
		apiReq.Order = existingRules[len(existingRules)-1].Order
	}

	timeout := 60 * time.Minute
	start := time.Now()
	for {
		if _, err := filteringrules.Update(ctx, svc, id, &apiReq); err != nil {
			if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.UpdateResponse[FirewallFilteringRuleState]{}, nil
			}
			if customErr := failFastOnErrorCodes(err); customErr != nil {
				return infer.UpdateResponse[FirewallFilteringRuleState]{}, customErr
			}
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Updating firewall filtering rule ID: %v, got INVALID_INPUT_ARGUMENT\n", id)
				if time.Since(start) < timeout {
					time.Sleep(10 * time.Second)
					continue
				}
			}
			return infer.UpdateResponse[FirewallFilteringRuleState]{}, err
		}

		reorderWithBeforeReorder(
			OrderRule{Order: intendedOrder, Rank: intendedRank},
			id,
			firewallFilteringResourceType,
			func() (int, error) {
				allRules, err := filteringrules.GetAll(ctx, svc, nil)
				if err != nil {
					return 0, err
				}
				return len(allRules), nil
			},
			func(ruleID int, order OrderRule) error {
				rule, err := filteringrules.Get(ctx, svc, ruleID)
				if err != nil {
					return err
				}
				// to avoid the STALE_CONFIGURATION_ERROR
				rule.LastModifiedTime = 0
				rule.LastModifiedBy = nil
				// Strip read-only fields that cause "Request body is invalid" for predefined rules
				rule.Predefined = false
				rule.DefaultRule = false
				rule.AccessControl = ""
				rule.Order = order.Order
				rule.Rank = order.Rank
				_, err = filteringrules.Update(ctx, svc, ruleID, rule)
				return err
			},
			nil,
		)

		markOrderRuleAsDone(id, firewallFilteringResourceType)
		waitForReorder(firewallFilteringResourceType)
		break
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[FirewallFilteringRuleState]{}, activationErr
		}
	}

	updated, err := filteringrules.Get(ctx, svc, id)
	if err != nil {
		return infer.UpdateResponse[FirewallFilteringRuleState]{Output: FirewallFilteringRuleState{
			FirewallFilteringRuleArgs: req.Inputs,
			RuleID:                    intPtr(id),
		}}, nil
	}
	return infer.UpdateResponse[FirewallFilteringRuleState]{Output: firewallFilteringRuleAPIToState(updated)}, nil
}

func (FirewallFilteringRule) Delete(ctx context.Context, req infer.DeleteRequest[FirewallFilteringRuleState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid firewall filtering rule ID: %s", req.ID)
	}
	rule, err := filteringrules.Get(ctx, svc, id)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("error retrieving firewall filtering rule %d: %w", id, err)
	}
	if err := validateFirewallFilteringRuleDelete(rule.Name); err != nil {
		return infer.DeleteResponse{}, err
	}
	if rule.Predefined {
		return infer.DeleteResponse{}, fmt.Errorf("deletion of predefined rule '%s' is not allowed", rule.Name)
	}
	if _, err := filteringrules.Delete(ctx, svc, id); err != nil {
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

func (FirewallFilteringRule) Annotate(a infer.Annotator) {
	describeResource(a, &FirewallFilteringRule{}, `The zia_firewall_filtering_rule resource manages firewall filtering rules in the Zscaler Internet Access (ZIA) cloud service. Cloud firewall rules control traffic that is forwarded to the Zscaler service for inspection, allowing you to allow, block, or apply specific actions based on source, destination, applications, and other criteria.

For more information, see the [ZIA Cloud Firewall documentation](https://help.zscaler.com/zia/firewall-policies).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Firewall Filtering Rule

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.FirewallFilteringRule("example", {
    name: "Example Firewall Rule",
    description: "Allow outbound traffic",
    order: 1,
    state: "ENABLED",
    action: "ALLOW",
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.FirewallFilteringRule("example",
    name="Example Firewall Rule",
    description="Allow outbound traffic",
    order=1,
    state="ENABLED",
    action="ALLOW",
)
`+tripleBacktick("")+`

`+tripleBacktick("go")+`
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zia "github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := zia.NewFirewallFilteringRule(ctx, "example", &zia.FirewallFilteringRuleArgs{
			Name:        pulumi.String("Example Firewall Rule"),
			Description: pulumi.StringRef("Allow outbound traffic"),
			Order:       pulumi.Int(1),
			State:       pulumi.StringRef("ENABLED"),
			Action:      pulumi.StringRef("ALLOW"),
		})
		return err
	})
}
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:FirewallFilteringRule
    properties:
      name: Example Firewall Rule
      description: Allow outbound traffic
      order: 1
      state: ENABLED
      action: ALLOW
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing Firewall Filtering Rule can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:FirewallFilteringRule example 12345
`+tripleBacktick("")+`
`)
}

func (a *FirewallFilteringRuleArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the firewall filtering rule. Must be unique.")
	ann.Describe(&a.Order, "The order of execution of the rule with respect to other firewall filtering rules.")
	ann.Describe(&a.Description, "Additional information about the firewall filtering rule.")
	ann.Describe(&a.Rank, "Admin rank of the firewall filtering policy rule. Valid values: 0-7. Default: 7.")
	ann.Describe(&a.State, "Rule state. Valid values: `ENABLED`, `DISABLED`.")
	ann.Describe(&a.Action, "The action the rule takes when traffic matches. Valid values: `ALLOW`, `BLOCK_DROP`, `BLOCK_RESET`, `BLOCK_ICMP`, `EVAL_NWAPP`.")
	ann.Describe(&a.EnableFullLogging, "If set to true, enables full logging for the rule.")
	ann.Describe(&a.DefaultRule, "Indicates whether this is the default firewall filtering rule.")
	ann.Describe(&a.Predefined, "Indicates whether this is a predefined rule.")
	ann.Describe(&a.ExcludeSrcCountries, "If set to true, the countries specified in sourceCountries are excluded from the rule.")
	ann.Describe(&a.SrcIps, "Source IP addresses or CIDR ranges for the rule.")
	ann.Describe(&a.DestAddresses, "Destination IP addresses, FQDNs, or wildcard FQDNs for the rule.")
	ann.Describe(&a.DestIpCategories, "Destination IP address URL categories. Allows you to identify destinations based on the URL category of the domain.")
	ann.Describe(&a.DestCountries, "Destination countries (ISO 3166-1 alpha-2 codes) for the rule.")
	ann.Describe(&a.SourceCountries, "Source countries (ISO 3166-1 alpha-2 codes) for the rule.")
	ann.Describe(&a.NwApplications, "Network application values to which the rule applies (e.g., `APNS`, `DNS`, `HTTP`).")
	ann.Describe(&a.Locations, "IDs of locations to which the rule must be applied.")
	ann.Describe(&a.LocationGroups, "IDs of location groups to which the rule must be applied.")
	ann.Describe(&a.Departments, "IDs of departments to which the rule must be applied.")
	ann.Describe(&a.Groups, "IDs of groups to which the rule must be applied.")
	ann.Describe(&a.Users, "IDs of users to which the rule must be applied.")
	ann.Describe(&a.TimeWindows, "IDs of time intervals during which the rule must be enforced.")
	ann.Describe(&a.NwApplicationGroups, "IDs of network application groups to which the rule applies.")
	ann.Describe(&a.AppServices, "IDs of application services to which the rule applies.")
	ann.Describe(&a.AppServiceGroups, "IDs of application service groups to which the rule applies.")
	ann.Describe(&a.Labels, "IDs of labels associated with the rule.")
	ann.Describe(&a.SrcIpGroups, "IDs of source IP address groups for the rule.")
	ann.Describe(&a.DestIpGroups, "IDs of destination IP address groups for the rule.")
	ann.Describe(&a.NwServices, "IDs of network services to which the rule applies.")
	ann.Describe(&a.NwServiceGroups, "IDs of network service groups to which the rule applies.")
	ann.Describe(&a.DeviceGroups, "IDs of device groups for which the rule must be applied. Applicable for devices managed using Zscaler Client Connector.")
	ann.Describe(&a.Devices, "IDs of devices for which the rule must be applied.")
	ann.Describe(&a.DeviceTrustLevels, "Device trust levels for the rule. Valid values: `ANY`, `UNKNOWN_DEVICETRUSTLEVEL`, `LOW_TRUST`, `MEDIUM_TRUST`, `HIGH_TRUST`.")
	ann.Describe(&a.WorkloadGroups, "List of preconfigured workload groups to which the policy must be applied.")
	ann.Describe(&a.ZpaAppSegments, "List of ZPA application segments for which this rule is applicable. This field is applicable only for the ZPA gateway forwarding method.")
}

func (s *FirewallFilteringRuleState) Annotate(a infer.Annotator) {
	a.Describe(&s.RuleID, "The system-generated ID of the firewall filtering rule.")
}

var _ infer.CustomResource[FirewallFilteringRuleArgs, FirewallFilteringRuleState] = FirewallFilteringRule{}
