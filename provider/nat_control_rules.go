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

// Package provider implements the NAT Control Rules resource.
// Adopted from terraform-provider-zia resource_zia_nat_control_rules.go.
// Uses shared reorderWithBeforeReorder; updateOrder skips predefined rules.

package provider

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/nat_control_policies"
)

const natControlResourceType = "nat_control_rules"

// NatControlRule implements the zia:index:NatControlRule resource.
type NatControlRule struct{}

// NatControlRuleArgs are the inputs.
type NatControlRuleArgs struct {
	Name              string   `pulumi:"name"`
	Order             int      `pulumi:"order"`
	Description       *string  `pulumi:"description,optional"`
	Rank              *int     `pulumi:"rank,optional"`
	State             *string  `pulumi:"state,optional"`
	EnableFullLogging *bool    `pulumi:"enableFullLogging,optional"`
	RedirectFqdn      *string  `pulumi:"redirectFqdn,optional"`
	RedirectIp        *string  `pulumi:"redirectIp,optional"`
	RedirectPort      *int     `pulumi:"redirectPort,optional"`
	SrcIps            []string `pulumi:"srcIps,optional"`
	DestAddresses     []string `pulumi:"destAddresses,optional"`
	DestIpCategories  []string `pulumi:"destIpCategories,optional"`
	ResCategories     []string `pulumi:"resCategories,optional"`
	DestCountries     []string `pulumi:"destCountries,optional"`
	DefaultRule       *bool    `pulumi:"defaultRule,optional"`
	Predefined        *bool    `pulumi:"predefined,optional"`
	Locations         []int    `pulumi:"locations,optional"`
	LocationGroups    []int    `pulumi:"locationGroups,optional"`
	Departments       []int    `pulumi:"departments,optional"`
	Groups            []int    `pulumi:"groups,optional"`
	Users             []int    `pulumi:"users,optional"`
	TimeWindows       []int    `pulumi:"timeWindows,optional"`
	SrcIpGroups       []int    `pulumi:"srcIpGroups,optional"`
	SrcIpv6Groups     []int    `pulumi:"srcIpv6Groups,optional"`
	DestIpGroups      []int    `pulumi:"destIpGroups,optional"`
	DestIpv6Groups    []int    `pulumi:"destIpv6Groups,optional"`
	NwServices        []int    `pulumi:"nwServices,optional"`
	NwServiceGroups   []int    `pulumi:"nwServiceGroups,optional"`
	Labels            []int    `pulumi:"labels,optional"`
	DeviceGroups      []int    `pulumi:"deviceGroups,optional"`
	Devices           []int    `pulumi:"devices,optional"`
}

// NatControlRuleState is the persisted state.
type NatControlRuleState struct {
	NatControlRuleArgs
	RuleID *int `pulumi:"ruleId"`
}

func natControlRuleArgsToAPI(args *NatControlRuleArgs, id int) nat_control_policies.NatControlPolicies {
	order := args.Order
	if order == 0 {
		order = 1
	}
	rank := ptrToIntDefault(args.Rank, 7)
	state := ptrToString(args.State)
	if state == "" {
		state = "ENABLED"
	}
	redirectPort := 0
	if args.RedirectPort != nil {
		redirectPort = *args.RedirectPort
	}
	return nat_control_policies.NatControlPolicies{
		ID:                id,
		Name:              args.Name,
		Order:             order,
		Rank:              rank,
		Description:       ptrToString(args.Description),
		State:             state,
		RedirectFqdn:      ptrToString(args.RedirectFqdn),
		RedirectIp:        ptrToString(args.RedirectIp),
		RedirectPort:      redirectPort,
		SrcIps:            args.SrcIps,
		DestAddresses:     args.DestAddresses,
		DestIpCategories:  args.DestIpCategories,
		ResCategories:     args.ResCategories,
		DestCountries:     processCountries(args.DestCountries),
		EnableFullLogging: ptrToBool(args.EnableFullLogging),
		DefaultRule:       ptrToBool(args.DefaultRule),
		Predefined:        ptrToBool(args.Predefined),
		Locations:         idsToIDNameExtensions(args.Locations),
		LocationGroups:    idsToIDNameExtensions(args.LocationGroups),
		Departments:       idsToIDNameExtensions(args.Departments),
		Groups:            idsToIDNameExtensions(args.Groups),
		Users:             idsToIDNameExtensions(args.Users),
		TimeWindows:       idsToIDNameExtensions(args.TimeWindows),
		SrcIpGroups:       idsToIDNameExtensions(args.SrcIpGroups),
		SrcIpv6Groups:     idsToIDNameExtensions(args.SrcIpv6Groups),
		DestIpGroups:      idsToIDNameExtensions(args.DestIpGroups),
		DestIpv6Groups:    idsToIDNameExtensions(args.DestIpv6Groups),
		NwServices:        idsToIDNameExtensions(args.NwServices),
		NwServiceGroups:   idsToIDNameExtensions(args.NwServiceGroups),
		Labels:            idsToIDNameExtensions(args.Labels),
		DeviceGroups:      idsToIDNameExtensions(args.DeviceGroups),
		Devices:           idsToIDNameExtensions(args.Devices),
	}
}

func natControlRuleAPIToState(api *nat_control_policies.NatControlPolicies) NatControlRuleState {
	destCountries := make([]string, len(api.DestCountries))
	for i, c := range api.DestCountries {
		destCountries[i] = strings.TrimPrefix(c, "COUNTRY_")
	}
	return NatControlRuleState{
		NatControlRuleArgs: NatControlRuleArgs{
			Name:              api.Name,
			Order:             api.Order,
			Description:       stringPtr(api.Description),
			Rank:              intPtr(api.Rank),
			State:             stringPtr(api.State),
			EnableFullLogging: boolPtr(api.EnableFullLogging),
			RedirectFqdn:      stringPtr(api.RedirectFqdn),
			RedirectIp:        stringPtr(api.RedirectIp),
			RedirectPort:      intPtr(api.RedirectPort),
			SrcIps:            api.SrcIps,
			DestAddresses:     api.DestAddresses,
			DestIpCategories:  api.DestIpCategories,
			ResCategories:     api.ResCategories,
			DestCountries:     destCountries,
			DefaultRule:       boolPtr(api.DefaultRule),
			Predefined:        boolPtr(api.Predefined),
			Locations:         idNameExtensionsToIDs(api.Locations),
			LocationGroups:    idNameExtensionsToIDs(api.LocationGroups),
			Departments:       idNameExtensionsToIDs(api.Departments),
			Groups:            idNameExtensionsToIDs(api.Groups),
			Users:             idNameExtensionsToIDs(api.Users),
			TimeWindows:       idNameExtensionsToIDs(api.TimeWindows),
			SrcIpGroups:       idNameExtensionsToIDs(api.SrcIpGroups),
			SrcIpv6Groups:     idNameExtensionsToIDs(api.SrcIpv6Groups),
			DestIpGroups:      idNameExtensionsToIDs(api.DestIpGroups),
			DestIpv6Groups:    idNameExtensionsToIDs(api.DestIpv6Groups),
			NwServices:        idNameExtensionsToIDs(api.NwServices),
			NwServiceGroups:   idNameExtensionsToIDs(api.NwServiceGroups),
			Labels:            idNameExtensionsToIDs(api.Labels),
			DeviceGroups:      idNameExtensionsToIDs(api.DeviceGroups),
			Devices:           idNameExtensionsToIDs(api.Devices),
		},
		RuleID: intPtr(api.ID),
	}
}

func (NatControlRule) Create(ctx context.Context, req infer.CreateRequest[NatControlRuleArgs]) (infer.CreateResponse[NatControlRuleState], error) {
	if req.DryRun {
		s := NatControlRuleState{NatControlRuleArgs: req.Inputs, RuleID: intPtr(0)}
		return infer.CreateResponse[NatControlRuleState]{ID: "preview", Output: s}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[NatControlRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := natControlRuleArgsToAPI(&req.Inputs, 0)
	timeout := 60 * time.Minute
	start := time.Now()

	for {
		natControlRuleLock.Lock()
		if natControlRuleStartingOrder == 0 {
			list, _ := nat_control_policies.GetAll(ctx, svc)
			for _, r := range list {
				if r.Order > natControlRuleStartingOrder {
					natControlRuleStartingOrder = r.Order
				}
			}
			if natControlRuleStartingOrder == 0 {
				natControlRuleStartingOrder = 1
			}
		}
		natControlRuleLock.Unlock()

		intendedOrder := apiReq.Order
		intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)
		if intendedRank < 7 {
			apiReq.Rank = 7
		}
		apiReq.Order = natControlRuleStartingOrder

		resp, err := nat_control_policies.Create(ctx, svc, &apiReq)
		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.CreateResponse[NatControlRuleState]{}, customErr
		}
		if err != nil {
			reg := regexp.MustCompile("Rule with rank [0-9]+ is not allowed at order [0-9]+")
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") && reg.MatchString(err.Error()) {
				return infer.CreateResponse[NatControlRuleState]{}, fmt.Errorf("error creating nat control rule: %s, check order %d vs rank %d, err:%s",
					apiReq.Name, intendedOrder, intendedRank, err)
			}
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Creating nat control rule name: %v, got INVALID_INPUT_ARGUMENT\n", apiReq.Name)
				if time.Since(start) < timeout {
					time.Sleep(10 * time.Second)
					continue
				}
			}
			return infer.CreateResponse[NatControlRuleState]{}, fmt.Errorf("creating nat control rule: %w", err)
		}

		reorderWithBeforeReorder(
			OrderRule{Order: intendedOrder, Rank: intendedRank},
			resp.ID,
			natControlResourceType,
			func() (int, error) {
				allRules, err := nat_control_policies.GetAll(ctx, svc)
				if err != nil {
					return 0, err
				}
				return len(allRules), nil
			},
			func(id int, order OrderRule) error {
				rule, err := nat_control_policies.Get(ctx, svc, id)
				if err != nil {
					return err
				}
				if rule.Predefined {
					log.Printf("[INFO] Skipping reorder update for predefined rule ID %d (order: %d)", id, rule.Order)
					return nil
				}
				// to avoid the STALE_CONFIGURATION_ERROR
				rule.LastModifiedTime = 0
				rule.LastModifiedBy = nil
				rule.Order = order.Order
				rule.Rank = order.Rank
				_, err = nat_control_policies.Update(ctx, svc, id, rule)
				return err
			},
			nil,
		)

		markOrderRuleAsDone(resp.ID, natControlResourceType)
		waitForReorder(natControlResourceType)

		if shouldActivate() {
			time.Sleep(2 * time.Second)
			if activationErr := triggerActivation(ctx, client); activationErr != nil {
				return infer.CreateResponse[NatControlRuleState]{}, activationErr
			}
		} else {
			log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
		}

		rule, err := nat_control_policies.Get(ctx, svc, resp.ID)
		if err != nil {
			state := NatControlRuleState{NatControlRuleArgs: req.Inputs, RuleID: intPtr(resp.ID)}
			return infer.CreateResponse[NatControlRuleState]{ID: strconv.Itoa(resp.ID), Output: state}, nil
		}
		return infer.CreateResponse[NatControlRuleState]{
			ID:     strconv.Itoa(resp.ID),
			Output: natControlRuleAPIToState(rule),
		}, nil
	}
}

func (NatControlRule) Read(ctx context.Context, req infer.ReadRequest[NatControlRuleArgs, NatControlRuleState]) (infer.ReadResponse[NatControlRuleArgs, NatControlRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[NatControlRuleArgs, NatControlRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		rule, lookupErr := nat_control_policies.GetByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[NatControlRuleArgs, NatControlRuleState]{}, fmt.Errorf("nat control rule not found")
			}
			return infer.ReadResponse[NatControlRuleArgs, NatControlRuleState]{}, lookupErr
		}
		id = rule.ID
	}

	rule, err := nat_control_policies.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[NatControlRuleArgs, NatControlRuleState]{}, fmt.Errorf("nat control rule not found")
		}
		return infer.ReadResponse[NatControlRuleArgs, NatControlRuleState]{}, err
	}

	state := natControlRuleAPIToState(rule)
	return infer.ReadResponse[NatControlRuleArgs, NatControlRuleState]{
		ID:     req.ID,
		Inputs: state.NatControlRuleArgs,
		State:  state,
	}, nil
}

func (NatControlRule) Update(ctx context.Context, req infer.UpdateRequest[NatControlRuleArgs, NatControlRuleState]) (infer.UpdateResponse[NatControlRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[NatControlRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[NatControlRuleState]{}, fmt.Errorf("invalid nat control rule ID: %s", req.ID)
	}
	apiReq := natControlRuleArgsToAPI(&req.Inputs, id)

	timeout := 60 * time.Minute
	start := time.Now()
	for {
		if _, err := nat_control_policies.Update(ctx, svc, id, &apiReq); err != nil {
			if customErr := failFastOnErrorCodes(err); customErr != nil {
				return infer.UpdateResponse[NatControlRuleState]{}, customErr
			}
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Updating nat control rule ID: %v, got INVALID_INPUT_ARGUMENT\n", id)
				if time.Since(start) < timeout {
					time.Sleep(10 * time.Second)
					continue
				}
			}
			return infer.UpdateResponse[NatControlRuleState]{}, err
		}

		reorderWithBeforeReorder(
			OrderRule{Order: apiReq.Order, Rank: apiReq.Rank},
			id,
			natControlResourceType,
			func() (int, error) {
				allRules, err := nat_control_policies.GetAll(ctx, svc)
				if err != nil {
					return 0, err
				}
				return len(allRules), nil
			},
			func(ruleID int, order OrderRule) error {
				rule, err := nat_control_policies.Get(ctx, svc, ruleID)
				if err != nil {
					return err
				}
				if rule.Predefined {
					log.Printf("[INFO] Skipping reorder update for predefined rule ID %d (order: %d)", ruleID, rule.Order)
					return nil
				}
				// Optional: avoid unnecessary updates if the current order is already correct
				if rule.Order == order.Order && rule.Rank == order.Rank {
					return nil
				}
				// to avoid the STALE_CONFIGURATION_ERROR
				rule.LastModifiedTime = 0
				rule.LastModifiedBy = nil
				rule.Order = order.Order
				rule.Rank = order.Rank
				_, err = nat_control_policies.Update(ctx, svc, ruleID, rule)
				return err
			},
			nil,
		)

		markOrderRuleAsDone(id, natControlResourceType)
		waitForReorder(natControlResourceType)
		break
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[NatControlRuleState]{}, activationErr
		}
	}

	updated, err := nat_control_policies.Get(ctx, svc, id)
	if err != nil {
		return infer.UpdateResponse[NatControlRuleState]{Output: NatControlRuleState{
			NatControlRuleArgs: req.Inputs,
			RuleID:             intPtr(id),
		}}, nil
	}
	return infer.UpdateResponse[NatControlRuleState]{Output: natControlRuleAPIToState(updated)}, nil
}

func (NatControlRule) Delete(ctx context.Context, req infer.DeleteRequest[NatControlRuleState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid nat control rule ID: %s", req.ID)
	}
	rule, err := nat_control_policies.Get(ctx, svc, id)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("error retrieving nat control rule %d: %w", id, err)
	}
	if rule.Predefined {
		return infer.DeleteResponse{}, fmt.Errorf("deletion of predefined rule '%s' is not allowed", rule.Name)
	}
	if _, err := nat_control_policies.Delete(ctx, svc, id); err != nil {
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

func (NatControlRule) Annotate(a infer.Annotator) {
	describeResource(a, &NatControlRule{}, `The zia_nat_control_rule resource manages NAT control rules in the Zscaler Internet Access (ZIA) cloud service. NAT control rules allow you to redirect traffic to specific IP addresses or FQDNs and ports based on various criteria such as source, destination, users, and locations.

For more information, see the [ZIA NAT Control documentation](https://help.zscaler.com/zia/nat-control-policies).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic NAT Control Rule

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.NatControlRule("example", {
    name: "Example NAT Control Rule",
    description: "Redirect traffic",
    order: 1,
    state: "ENABLED",
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.NatControlRule("example",
    name="Example NAT Control Rule",
    description="Redirect traffic",
    order=1,
    state="ENABLED",
)
`+tripleBacktick("")+`

`+tripleBacktick("go")+`
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zia "github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := zia.NewNatControlRule(ctx, "example", &zia.NatControlRuleArgs{
			Name:        pulumi.String("Example NAT Control Rule"),
			Description: pulumi.StringRef("Redirect traffic"),
			Order:       pulumi.Int(1),
			State:       pulumi.StringRef("ENABLED"),
		})
		return err
	})
}
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:NatControlRule
    properties:
      name: Example NAT Control Rule
      description: Redirect traffic
      order: 1
      state: ENABLED
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing NAT Control Rule can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:NatControlRule example 12345
`+tripleBacktick("")+`
`)
}

func (a *NatControlRuleArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the NAT control rule. Must be unique.")
	ann.Describe(&a.Order, "The order of execution of the rule with respect to other NAT control rules.")
	ann.Describe(&a.Description, "Additional information about the NAT control rule.")
	ann.Describe(&a.Rank, "Admin rank of the NAT control policy rule. Valid values: 0-7. Default: 7.")
	ann.Describe(&a.State, "Rule state. Valid values: `ENABLED`, `DISABLED`.")
	ann.Describe(&a.EnableFullLogging, "If set to true, enables full logging for the rule.")
	ann.Describe(&a.RedirectFqdn, "The FQDN to which traffic should be redirected.")
	ann.Describe(&a.RedirectIp, "The IP address to which traffic should be redirected.")
	ann.Describe(&a.RedirectPort, "The port to which traffic should be redirected.")
	ann.Describe(&a.SrcIps, "Source IP addresses or CIDR ranges for the rule.")
	ann.Describe(&a.DestAddresses, "Destination IP addresses, FQDNs, or wildcard FQDNs for the rule.")
	ann.Describe(&a.DestIpCategories, "Destination IP address URL categories for the rule.")
	ann.Describe(&a.ResCategories, "URL categories that apply to the response for the rule.")
	ann.Describe(&a.DestCountries, "Destination countries (ISO 3166-1 alpha-2 codes) for the rule.")
	ann.Describe(&a.DefaultRule, "Indicates whether this is the default NAT control rule.")
	ann.Describe(&a.Predefined, "Indicates whether this is a predefined rule.")
	ann.Describe(&a.Locations, "IDs of locations to which the rule must be applied.")
	ann.Describe(&a.LocationGroups, "IDs of location groups to which the rule must be applied.")
	ann.Describe(&a.Departments, "IDs of departments to which the rule must be applied.")
	ann.Describe(&a.Groups, "IDs of groups to which the rule must be applied.")
	ann.Describe(&a.Users, "IDs of users to which the rule must be applied.")
	ann.Describe(&a.TimeWindows, "IDs of time intervals during which the rule must be enforced.")
	ann.Describe(&a.SrcIpGroups, "IDs of source IP address groups for the rule.")
	ann.Describe(&a.SrcIpv6Groups, "IDs of source IPv6 address groups for the rule.")
	ann.Describe(&a.DestIpGroups, "IDs of destination IP address groups for the rule.")
	ann.Describe(&a.DestIpv6Groups, "IDs of destination IPv6 address groups for the rule.")
	ann.Describe(&a.NwServices, "IDs of network services to which the rule applies.")
	ann.Describe(&a.NwServiceGroups, "IDs of network service groups to which the rule applies.")
	ann.Describe(&a.Labels, "IDs of labels associated with the rule.")
	ann.Describe(&a.DeviceGroups, "IDs of device groups for which the rule must be applied. Applicable for devices managed using Zscaler Client Connector.")
	ann.Describe(&a.Devices, "IDs of devices for which the rule must be applied.")
}

func (s *NatControlRuleState) Annotate(a infer.Annotator) {
	a.Describe(&s.RuleID, "The system-generated ID of the NAT control rule.")
}

var _ infer.CustomResource[NatControlRuleArgs, NatControlRuleState] = NatControlRule{}
