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
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/forwarding_control_policy/forwarding_rules"
)

const forwardingControlResourceType = "forwarding_control_rule"

type ForwardingControlRule struct{}

type ForwardingControlRuleArgs struct {
	Name                string             `pulumi:"name"`
	Order               int                `pulumi:"order"`
	Description         *string            `pulumi:"description,optional"`
	Type                *string            `pulumi:"type,optional"`
	ForwardMethod       string             `pulumi:"forwardMethod"`
	Rank                *int               `pulumi:"rank,optional"`
	State               *string            `pulumi:"state,optional"`
	SrcIps              []string           `pulumi:"srcIps,optional"`
	DestAddresses       []string           `pulumi:"destAddresses,optional"`
	DestIpCategories    []string           `pulumi:"destIpCategories,optional"`
	ResCategories       []string           `pulumi:"resCategories,optional"`
	DestCountries       []string           `pulumi:"destCountries,optional"`
	Locations           []int              `pulumi:"locations,optional"`
	LocationGroups      []int              `pulumi:"locationGroups,optional"`
	ECGroups            []int              `pulumi:"ecGroups,optional"`
	Departments         []int              `pulumi:"departments,optional"`
	Groups              []int              `pulumi:"groups,optional"`
	Users               []int              `pulumi:"users,optional"`
	SrcIpGroups         []int              `pulumi:"srcIpGroups,optional"`
	SrcIpv6Groups       []int              `pulumi:"srcIpv6Groups,optional"`
	DestIpGroups        []int              `pulumi:"destIpGroups,optional"`
	DestIpv6Groups      []int              `pulumi:"destIpv6Groups,optional"`
	NwServices          []int              `pulumi:"nwServices,optional"`
	NwServiceGroups     []int              `pulumi:"nwServiceGroups,optional"`
	NwApplicationGroups []int              `pulumi:"nwApplicationGroups,optional"`
	AppServiceGroups    []int              `pulumi:"appServiceGroups,optional"`
	Labels              []int              `pulumi:"labels,optional"`
	DeviceGroups        []int              `pulumi:"deviceGroups,optional"`
	ProxyGatewayID      *int               `pulumi:"proxyGatewayId,optional"`
	ZPAGatewayID        *int               `pulumi:"zpaGatewayId,optional"`
	ZpaAppSegments      []ZPAAppSegmentInput `pulumi:"zpaAppSegments,optional"`
}

type ForwardingControlRuleState struct {
	ForwardingControlRuleArgs
	RuleID *int `pulumi:"ruleId"`
}

func validateForwardingControlPredefined(rule *forwarding_rules.ForwardingRules) error {
	if rule.Name == "Client Connector Traffic Direct" || rule.Name == "ZPA Pool For Stray Traffic" {
		return fmt.Errorf("predefined rule '%s' cannot be deleted", rule.Name)
	}
	if rule.Name == "ZIA Inspected ZPA Apps" || rule.Name == "Fallback mode of ZPA Forwarding" {
		return fmt.Errorf("predefined rule '%s' cannot be deleted", rule.Name)
	}
	return nil
}

func forwardingControlRuleArgsToAPI(args *ForwardingControlRuleArgs, id int) forwarding_rules.ForwardingRules {
	order := args.Order
	if order == 0 {
		order = 1
	}
	rank := ptrToIntDefault(args.Rank, 7)
	state := ptrToString(args.State)
	t := ptrToString(args.Type)
	fm := args.ForwardMethod
	proxyID := 0
	if args.ProxyGatewayID != nil {
		proxyID = *args.ProxyGatewayID
	}
	zpaID := 0
	if args.ZPAGatewayID != nil {
		zpaID = *args.ZPAGatewayID
	}
	return forwarding_rules.ForwardingRules{
		ID:                  id,
		Name:                args.Name,
		Description:         ptrToString(args.Description),
		Order:               order,
		Rank:                rank,
		Type:                t,
		State:               state,
		ForwardMethod:       fm,
		SrcIps:              args.SrcIps,
		DestAddresses:       args.DestAddresses,
		DestIpCategories:    args.DestIpCategories,
		ResCategories:       args.ResCategories,
		DestCountries:       processCountries(args.DestCountries),
		Locations:           idsToIDNameExtensions(args.Locations),
		LocationsGroups:     idsToIDNameExtensions(args.LocationGroups),
		ECGroups:            idsToIDNameExtensions(args.ECGroups),
		Departments:         idsToIDNameExtensions(args.Departments),
		Groups:              idsToIDNameExtensions(args.Groups),
		Users:               idsToIDNameExtensions(args.Users),
		SrcIpGroups:         idsToIDNameExtensions(args.SrcIpGroups),
		SrcIpv6Groups:       idsToIDNameExtensions(args.SrcIpv6Groups),
		DestIpGroups:        idsToIDNameExtensions(args.DestIpGroups),
		DestIpv6Groups:      idsToIDNameExtensions(args.DestIpv6Groups),
		NwServices:          idsToIDNameExtensions(args.NwServices),
		NwServiceGroups:     idsToIDNameExtensions(args.NwServiceGroups),
		NwApplicationGroups: idsToIDNameExtensions(args.NwApplicationGroups),
		AppServiceGroups:    idsToIDNameExtensions(args.AppServiceGroups),
		Labels:              idsToIDNameExtensions(args.Labels),
		DeviceGroups:        idsToIDNameExtensions(args.DeviceGroups),
		ProxyGateway:        idToOptionalIDName(proxyID),
		ZPAGateway:          idToOptionalIDName(zpaID),
		ZPAAppSegments:      expandZPAAppSegments(args.ZpaAppSegments),
	}
}

func forwardingControlRuleAPIToState(api *forwarding_rules.ForwardingRules) ForwardingControlRuleState {
	destCountries := make([]string, len(api.DestCountries))
	for i, c := range api.DestCountries {
		destCountries[i] = strings.TrimPrefix(c, "COUNTRY_")
	}
	return ForwardingControlRuleState{
		ForwardingControlRuleArgs: ForwardingControlRuleArgs{
			Name:                api.Name,
			Order:               api.Order,
			Description:         stringPtr(api.Description),
			Type:                stringPtr(api.Type),
			ForwardMethod:       api.ForwardMethod,
			Rank:                intPtr(api.Rank),
			State:               stringPtr(api.State),
			SrcIps:              api.SrcIps,
			DestAddresses:       api.DestAddresses,
			DestIpCategories:    api.DestIpCategories,
			ResCategories:       api.ResCategories,
			DestCountries:       destCountries,
			Locations:           idNameExtensionsToIDs(api.Locations),
			LocationGroups:      idNameExtensionsToIDs(api.LocationsGroups),
			ECGroups:            idNameExtensionsToIDs(api.ECGroups),
			Departments:         idNameExtensionsToIDs(api.Departments),
			Groups:              idNameExtensionsToIDs(api.Groups),
			Users:               idNameExtensionsToIDs(api.Users),
			SrcIpGroups:         idNameExtensionsToIDs(api.SrcIpGroups),
			SrcIpv6Groups:       idNameExtensionsToIDs(api.SrcIpv6Groups),
			DestIpGroups:        idNameExtensionsToIDs(api.DestIpGroups),
			DestIpv6Groups:      idNameExtensionsToIDs(api.DestIpv6Groups),
			NwServices:          idNameExtensionsToIDs(api.NwServices),
			NwServiceGroups:     idNameExtensionsToIDs(api.NwServiceGroups),
			NwApplicationGroups: idNameExtensionsToIDs(api.NwApplicationGroups),
			AppServiceGroups:    idNameExtensionsToIDs(api.AppServiceGroups),
			Labels:              idNameExtensionsToIDs(api.Labels),
			DeviceGroups:        idNameExtensionsToIDs(api.DeviceGroups),
			ProxyGatewayID:      idNameToOptionalID(api.ProxyGateway),
			ZPAGatewayID:        idNameToOptionalID(api.ZPAGateway),
			ZpaAppSegments:      flattenZPAAppSegments(api.ZPAAppSegments),
		},
		RuleID: intPtr(api.ID),
	}
}

func (ForwardingControlRule) Create(ctx context.Context, req infer.CreateRequest[ForwardingControlRuleArgs]) (infer.CreateResponse[ForwardingControlRuleState], error) {
	if req.DryRun {
		return infer.CreateResponse[ForwardingControlRuleState]{ID: "preview", Output: ForwardingControlRuleState{ForwardingControlRuleArgs: req.Inputs, RuleID: intPtr(0)}}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[ForwardingControlRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := forwardingControlRuleArgsToAPI(&req.Inputs, 0)
	if err := validateForwardingControlPredefined(&apiReq); err != nil {
		return infer.CreateResponse[ForwardingControlRuleState]{}, err
	}

	if req.Inputs.ForwardMethod == "ZPA" {
		time.Sleep(60 * time.Second)
	}

	for {
		forwardingControlLock.Lock()
		if forwardingControlStartingOrder == 0 {
			list, _ := forwarding_rules.GetAll(ctx, svc)
			for _, r := range list {
				if r.Order > forwardingControlStartingOrder {
					forwardingControlStartingOrder = r.Order
				}
			}
			if forwardingControlStartingOrder == 0 {
				forwardingControlStartingOrder = 1
			}
		}
		forwardingControlLock.Unlock()

		intendedOrder := apiReq.Order
		intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)
		if intendedRank < 7 {
			apiReq.Rank = 7
		}
		apiReq.Order = forwardingControlStartingOrder

		var resp *forwarding_rules.ForwardingRules
		var err error
		for i := 0; i < 3; i++ {
			resp, err = forwarding_rules.Create(ctx, svc, &apiReq)
			if err == nil {
				break
			}
			if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr != nil && respErr.Response != nil && respErr.Response.StatusCode == 400 &&
				strings.Contains(respErr.Message, "is no longer an active Source IP Anchored App Segment") {
				time.Sleep(30 * time.Second)
				continue
			}
			break
		}
		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.CreateResponse[ForwardingControlRuleState]{}, customErr
		}
		if err != nil {
			reg := regexp.MustCompile("Rule with rank [0-9]+ is not allowed at order [0-9]+")
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") && reg.MatchString(err.Error()) {
				return infer.CreateResponse[ForwardingControlRuleState]{}, fmt.Errorf("error creating forwarding control rule: %s, check order %d vs rank %d, err:%s", apiReq.Name, intendedOrder, intendedRank, err)
			}
			return infer.CreateResponse[ForwardingControlRuleState]{}, fmt.Errorf("creating forwarding control rule: %w", err)
		}

		reorderWithBeforeReorder(
			OrderRule{Order: intendedOrder, Rank: intendedRank},
			resp.ID,
			forwardingControlResourceType,
			func() (int, error) {
				allRules, err := forwarding_rules.GetAll(ctx, svc)
				if err != nil {
					return 0, err
				}
				return len(allRules), nil
			},
			func(id int, order OrderRule) error {
				rule, err := forwarding_rules.Get(ctx, svc, id)
				if err != nil {
					return err
				}
				rule.Order = order.Order
				rule.Rank = order.Rank
				_, err = forwarding_rules.Update(ctx, svc, id, rule)
				return err
			},
			nil,
		)
		markOrderRuleAsDone(resp.ID, forwardingControlResourceType)
		waitForReorder(forwardingControlResourceType)

		if shouldActivate() {
			time.Sleep(2 * time.Second)
			if activationErr := triggerActivation(ctx, client); activationErr != nil {
				return infer.CreateResponse[ForwardingControlRuleState]{}, activationErr
			}
		} else {
			log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
		}

		rule, err := forwarding_rules.Get(ctx, svc, resp.ID)
		if err != nil {
			return infer.CreateResponse[ForwardingControlRuleState]{ID: strconv.Itoa(resp.ID), Output: ForwardingControlRuleState{ForwardingControlRuleArgs: req.Inputs, RuleID: intPtr(resp.ID)}}, nil
		}
		return infer.CreateResponse[ForwardingControlRuleState]{ID: strconv.Itoa(resp.ID), Output: forwardingControlRuleAPIToState(rule)}, nil
	}
}

func (ForwardingControlRule) Read(ctx context.Context, req infer.ReadRequest[ForwardingControlRuleArgs, ForwardingControlRuleState]) (infer.ReadResponse[ForwardingControlRuleArgs, ForwardingControlRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[ForwardingControlRuleArgs, ForwardingControlRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		rule, lookupErr := forwarding_rules.GetByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[ForwardingControlRuleArgs, ForwardingControlRuleState]{}, fmt.Errorf("forwarding control rule not found")
			}
			return infer.ReadResponse[ForwardingControlRuleArgs, ForwardingControlRuleState]{}, lookupErr
		}
		id = rule.ID
	}

	rule, err := forwarding_rules.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[ForwardingControlRuleArgs, ForwardingControlRuleState]{}, fmt.Errorf("forwarding control rule not found")
		}
		return infer.ReadResponse[ForwardingControlRuleArgs, ForwardingControlRuleState]{}, err
	}

	state := forwardingControlRuleAPIToState(rule)
	return infer.ReadResponse[ForwardingControlRuleArgs, ForwardingControlRuleState]{ID: req.ID, Inputs: state.ForwardingControlRuleArgs, State: state}, nil
}

func (ForwardingControlRule) Update(ctx context.Context, req infer.UpdateRequest[ForwardingControlRuleArgs, ForwardingControlRuleState]) (infer.UpdateResponse[ForwardingControlRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[ForwardingControlRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[ForwardingControlRuleState]{}, fmt.Errorf("invalid forwarding control rule ID: %s", req.ID)
	}
	apiReq := forwardingControlRuleArgsToAPI(&req.Inputs, id)
	if err := validateForwardingControlPredefined(&apiReq); err != nil {
		return infer.UpdateResponse[ForwardingControlRuleState]{}, err
	}

	existingRules, err := forwarding_rules.GetAll(ctx, svc)
	if err == nil {
		sort.Slice(existingRules, func(i, j int) bool {
			return existingRules[i].Rank < existingRules[j].Rank || (existingRules[i].Rank == existingRules[j].Rank && existingRules[i].Order < existingRules[j].Order)
		})
		nextOrder := apiReq.Order
		if len(existingRules) > 0 {
			nextOrder = existingRules[len(existingRules)-1].Order
		}
		apiReq.Rank = 7
		apiReq.Order = nextOrder
	}

	if req.Inputs.ForwardMethod == "ZPA" {
		time.Sleep(60 * time.Second)
	}

	intendedOrder := req.Inputs.Order
	intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)

	for i := 0; i < 3; i++ {
		_, err = forwarding_rules.Update(ctx, svc, id, &apiReq)
		if err == nil {
			break
		}
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr != nil && respErr.Response != nil && respErr.Response.StatusCode == 400 &&
			strings.Contains(respErr.Message, "is no longer an active Source IP Anchored App Segment") {
			time.Sleep(30 * time.Second)
			continue
		}
		break
	}
	if customErr := failFastOnErrorCodes(err); customErr != nil {
		return infer.UpdateResponse[ForwardingControlRuleState]{}, customErr
	}
	if err != nil {
		return infer.UpdateResponse[ForwardingControlRuleState]{}, err
	}

	reorderWithBeforeReorder(
		OrderRule{Order: intendedOrder, Rank: intendedRank},
		id,
		forwardingControlResourceType,
		func() (int, error) {
			allRules, err := forwarding_rules.GetAll(ctx, svc)
			if err != nil {
				return 0, err
			}
			return len(allRules), nil
		},
		func(ruleID int, order OrderRule) error {
			rule, err := forwarding_rules.Get(ctx, svc, ruleID)
			if err != nil {
				return err
			}
			if rule.Order == order.Order && rule.Rank == order.Rank {
				return nil
			}
			rule.Order = order.Order
			rule.Rank = order.Rank
			_, err = forwarding_rules.Update(ctx, svc, ruleID, rule)
			return err
		},
		nil,
	)
	markOrderRuleAsDone(id, forwardingControlResourceType)
	waitForReorder(forwardingControlResourceType)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[ForwardingControlRuleState]{}, activationErr
		}
	}

	updated, err := forwarding_rules.Get(ctx, svc, id)
	if err != nil {
		return infer.UpdateResponse[ForwardingControlRuleState]{Output: ForwardingControlRuleState{ForwardingControlRuleArgs: req.Inputs, RuleID: intPtr(id)}}, nil
	}
	return infer.UpdateResponse[ForwardingControlRuleState]{Output: forwardingControlRuleAPIToState(updated)}, nil
}

func (ForwardingControlRule) Delete(ctx context.Context, req infer.DeleteRequest[ForwardingControlRuleState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid forwarding control rule ID: %s", req.ID)
	}
	rule, err := forwarding_rules.Get(ctx, svc, id)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("error retrieving forwarding control rule %d: %w", id, err)
	}
	if err := validateForwardingControlPredefined(rule); err != nil {
		return infer.DeleteResponse{}, err
	}
	if _, err := forwarding_rules.Delete(ctx, svc, id); err != nil {
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

func (ForwardingControlRule) Annotate(a infer.Annotator) {
	describeResource(a, &ForwardingControlRule{}, `The zia_forwarding_control_rule resource manages forwarding control rules in the Zscaler Internet Access (ZIA) cloud service. Forwarding control rules determine how traffic is forwarded — directly to the internet, via an explicit proxy, or through Zscaler Private Access (ZPA).

For more information, see the [ZIA Forwarding Control documentation](https://help.zscaler.com/zia/forwarding-control-policies).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Forwarding Control Rule

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.ForwardingControlRule("example", {
    name: "Example Forwarding Rule",
    description: "Forward traffic directly",
    order: 1,
    state: "ENABLED",
    forwardMethod: "DIRECT",
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.ForwardingControlRule("example",
    name="Example Forwarding Rule",
    description="Forward traffic directly",
    order=1,
    state="ENABLED",
    forward_method="DIRECT",
)
`+tripleBacktick("")+`

`+tripleBacktick("go")+`
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zia "github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := zia.NewForwardingControlRule(ctx, "example", &zia.ForwardingControlRuleArgs{
			Name:          pulumi.String("Example Forwarding Rule"),
			Description:   pulumi.StringRef("Forward traffic directly"),
			Order:         pulumi.Int(1),
			State:         pulumi.StringRef("ENABLED"),
			ForwardMethod: pulumi.String("DIRECT"),
		})
		return err
	})
}
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:ForwardingControlRule
    properties:
      name: Example Forwarding Rule
      description: Forward traffic directly
      order: 1
      state: ENABLED
      forwardMethod: DIRECT
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing Forwarding Control Rule can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:ForwardingControlRule example 12345
`+tripleBacktick("")+`
`)
}

func (a *ForwardingControlRuleArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the forwarding control rule. Must be unique.")
	ann.Describe(&a.Order, "The order of execution of the rule with respect to other forwarding control rules.")
	ann.Describe(&a.Description, "Additional information about the forwarding control rule.")
	ann.Describe(&a.Type, "The rule type. Valid values: `FORWARDING`.")
	ann.Describe(&a.ForwardMethod, "The type of traffic forwarding method. Valid values: `DIRECT`, `PROXYCHAIN`, `ZPA`, `ECZPA`, `DIRECT_NSS`.")
	ann.Describe(&a.Rank, "Admin rank of the forwarding control policy rule. Valid values: 0-7. Default: 7.")
	ann.Describe(&a.State, "Rule state. Valid values: `ENABLED`, `DISABLED`.")
	ann.Describe(&a.SrcIps, "Source IP addresses or CIDR ranges for the rule.")
	ann.Describe(&a.DestAddresses, "Destination IP addresses, FQDNs, or wildcard FQDNs for the rule.")
	ann.Describe(&a.DestIpCategories, "Destination IP address URL categories for the rule.")
	ann.Describe(&a.ResCategories, "URL categories that apply to the response for the rule.")
	ann.Describe(&a.DestCountries, "Destination countries (ISO 3166-1 alpha-2 codes) for the rule.")
	ann.Describe(&a.Locations, "IDs of locations to which the rule must be applied.")
	ann.Describe(&a.LocationGroups, "IDs of location groups to which the rule must be applied.")
	ann.Describe(&a.ECGroups, "IDs of Zscaler Edge Connector groups to which the rule applies.")
	ann.Describe(&a.Departments, "IDs of departments to which the rule must be applied.")
	ann.Describe(&a.Groups, "IDs of groups to which the rule must be applied.")
	ann.Describe(&a.Users, "IDs of users to which the rule must be applied.")
	ann.Describe(&a.SrcIpGroups, "IDs of source IP address groups for the rule.")
	ann.Describe(&a.SrcIpv6Groups, "IDs of source IPv6 address groups for the rule.")
	ann.Describe(&a.DestIpGroups, "IDs of destination IP address groups for the rule.")
	ann.Describe(&a.DestIpv6Groups, "IDs of destination IPv6 address groups for the rule.")
	ann.Describe(&a.NwServices, "IDs of network services to which the rule applies.")
	ann.Describe(&a.NwServiceGroups, "IDs of network service groups to which the rule applies.")
	ann.Describe(&a.NwApplicationGroups, "IDs of network application groups to which the rule applies.")
	ann.Describe(&a.AppServiceGroups, "IDs of application service groups to which the rule applies.")
	ann.Describe(&a.Labels, "IDs of labels associated with the rule.")
	ann.Describe(&a.DeviceGroups, "IDs of device groups for which the rule must be applied. Applicable for devices managed using Zscaler Client Connector.")
	ann.Describe(&a.ProxyGatewayID, "The ID of the proxy gateway. Required when forwardMethod is `PROXYCHAIN`.")
	ann.Describe(&a.ZPAGatewayID, "The ID of the ZPA gateway. Required when forwardMethod is `ZPA`.")
	ann.Describe(&a.ZpaAppSegments, "List of ZPA application segments for which this rule is applicable. This field is applicable only when forwardMethod is `ZPA`.")
}

func (s *ForwardingControlRuleState) Annotate(a infer.Annotator) {
	a.Describe(&s.RuleID, "The system-generated ID of the forwarding control rule.")
}

var _ infer.CustomResource[ForwardingControlRuleArgs, ForwardingControlRuleState] = ForwardingControlRule{}
