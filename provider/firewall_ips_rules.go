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
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/firewallipscontrolpolicies"
)

const firewallIPSResourceType = "firewall_ips_rule"

type FirewallIPSRule struct{}

type FirewallIPSRuleArgs struct {
	Name              string             `pulumi:"name"`
	Order             int                `pulumi:"order"`
	Description       *string            `pulumi:"description,optional"`
	Rank              *int               `pulumi:"rank,optional"`
	State             *string            `pulumi:"state,optional"`
	Action            *string            `pulumi:"action,optional"`
	EnableFullLogging *bool              `pulumi:"enableFullLogging,optional"`
	CapturePcap       *bool              `pulumi:"capturePcap,optional"`
	DefaultRule       *bool              `pulumi:"defaultRule,optional"`
	Predefined        *bool              `pulumi:"predefined,optional"`
	IsEunEnabled      *bool              `pulumi:"isEunEnabled,optional"`
	EunTemplateId     *int               `pulumi:"eunTemplateId,optional"`
	SrcIps            []string           `pulumi:"srcIps,optional"`
	DestAddresses     []string           `pulumi:"destAddresses,optional"`
	DestIpCategories  []string           `pulumi:"destIpCategories,optional"`
	ResCategories     []string           `pulumi:"resCategories,optional"`
	DestCountries     []string           `pulumi:"destCountries,optional"`
	SourceCountries   []string           `pulumi:"sourceCountries,optional"`
	Locations         []int              `pulumi:"locations,optional"`
	LocationGroups    []int              `pulumi:"locationGroups,optional"`
	Departments       []int              `pulumi:"departments,optional"`
	Groups            []int              `pulumi:"groups,optional"`
	Users             []int              `pulumi:"users,optional"`
	TimeWindows       []int              `pulumi:"timeWindows,optional"`
	SrcIpGroups       []int              `pulumi:"srcIpGroups,optional"`
	SrcIpv6Groups     []int              `pulumi:"srcIpv6Groups,optional"`
	DestIpGroups      []int              `pulumi:"destIpGroups,optional"`
	DestIpv6Groups    []int              `pulumi:"destIpv6Groups,optional"`
	NwServices        []int              `pulumi:"nwServices,optional"`
	NwServiceGroups   []int              `pulumi:"nwServiceGroups,optional"`
	Labels            []int              `pulumi:"labels,optional"`
	DeviceGroups      []int              `pulumi:"deviceGroups,optional"`
	Devices           []int              `pulumi:"devices,optional"`
	ThreatCategories  []int              `pulumi:"threatCategories,optional"`
	ZpaAppSegments   []ZPAAppSegmentInput `pulumi:"zpaAppSegments,optional"`
}

type FirewallIPSRuleState struct {
	FirewallIPSRuleArgs
	RuleID *int `pulumi:"ruleId"`
}

func firewallIPSRuleArgsToAPI(args *FirewallIPSRuleArgs, id int) firewallipscontrolpolicies.FirewallIPSRules {
	order := args.Order
	if order == 0 {
		order = 1
	}
	rank := ptrToIntDefault(args.Rank, 7)
	eunID := 0
	if args.EunTemplateId != nil {
		eunID = *args.EunTemplateId
	}
	return firewallipscontrolpolicies.FirewallIPSRules{
		ID:                id,
		Name:              args.Name,
		Order:             order,
		Rank:              rank,
		Action:            ptrToString(args.Action),
		State:             ptrToString(args.State),
		Description:       ptrToString(args.Description),
		SrcIps:            args.SrcIps,
		DestAddresses:     args.DestAddresses,
		DestIpCategories:  args.DestIpCategories,
		ResCategories:     args.ResCategories,
		DestCountries:     processCountries(args.DestCountries),
		SourceCountries:   processCountries(args.SourceCountries),
		EnableFullLogging: ptrToBool(args.EnableFullLogging),
		CapturePCAP:       ptrToBool(args.CapturePcap),
		DefaultRule:       ptrToBool(args.DefaultRule),
		Predefined:        ptrToBool(args.Predefined),
		IsEUNEnabled:      ptrToBool(args.IsEunEnabled),
		EUNTemplateID:     eunID,
		Locations:         idsToIDNameExtensions(args.Locations),
		LocationsGroups:   idsToIDNameExtensions(args.LocationGroups),
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
		ThreatCategories:  idsToIDNameExtensions(args.ThreatCategories),
		ZPAAppSegments:    expandZPAAppSegments(args.ZpaAppSegments),
	}
}

func firewallIPSRuleAPIToState(api *firewallipscontrolpolicies.FirewallIPSRules) FirewallIPSRuleState {
	return FirewallIPSRuleState{
		FirewallIPSRuleArgs: FirewallIPSRuleArgs{
			Name:              api.Name,
			Order:             api.Order,
			Description:       stringPtr(api.Description),
			Rank:              intPtr(api.Rank),
			State:             stringPtr(api.State),
			Action:            stringPtr(api.Action),
			EnableFullLogging: boolPtr(api.EnableFullLogging),
			CapturePcap:       boolPtr(api.CapturePCAP),
			DefaultRule:       boolPtr(api.DefaultRule),
			Predefined:        boolPtr(api.Predefined),
			IsEunEnabled:      boolPtr(api.IsEUNEnabled),
			EunTemplateId:     intPtr(api.EUNTemplateID),
			SrcIps:            api.SrcIps,
			DestAddresses:     api.DestAddresses,
			DestIpCategories:  api.DestIpCategories,
			ResCategories:     api.ResCategories,
			DestCountries:     processCountriesFromAPI(api.DestCountries),
			SourceCountries:   processCountriesFromAPI(api.SourceCountries),
			Locations:         idNameExtensionsToIDs(api.Locations),
			LocationGroups:    idNameExtensionsToIDs(api.LocationsGroups),
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
			ThreatCategories:  idNameExtensionsToIDs(api.ThreatCategories),
			ZpaAppSegments:    flattenZPAAppSegments(api.ZPAAppSegments),
		},
		RuleID: intPtr(api.ID),
	}
}

func (FirewallIPSRule) Diff(ctx context.Context, req infer.DiffRequest[FirewallIPSRuleArgs, FirewallIPSRuleState]) (infer.DiffResponse, error) {
	diff, hasChanges := diffSkippingUnsetInputs(req.State.FirewallIPSRuleArgs, req.Inputs)
	return infer.DiffResponse{HasChanges: hasChanges, DetailedDiff: diff}, nil
}

func (FirewallIPSRule) Create(ctx context.Context, req infer.CreateRequest[FirewallIPSRuleArgs]) (infer.CreateResponse[FirewallIPSRuleState], error) {
	if req.DryRun {
		return infer.CreateResponse[FirewallIPSRuleState]{ID: "preview", Output: FirewallIPSRuleState{FirewallIPSRuleArgs: req.Inputs, RuleID: intPtr(0)}}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[FirewallIPSRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := firewallIPSRuleArgsToAPI(&req.Inputs, 0)
	timeout := 60 * time.Minute
	start := time.Now()

	for {
		firewallIPSLock.Lock()
		if firewallIPSStartingOrder == 0 {
			list, _ := firewallipscontrolpolicies.GetAll(ctx, svc)
			for _, r := range list {
				if r.Order > firewallIPSStartingOrder {
					firewallIPSStartingOrder = r.Order
				}
			}
			if firewallIPSStartingOrder == 0 {
				firewallIPSStartingOrder = 1
			}
		}
		firewallIPSLock.Unlock()

		intendedOrder := apiReq.Order
		intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)
		if intendedRank < 7 {
			apiReq.Rank = 7
		}
		apiReq.Order = firewallIPSStartingOrder

		resp, err := firewallipscontrolpolicies.Create(ctx, svc, &apiReq)
		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.CreateResponse[FirewallIPSRuleState]{}, customErr
		}
		if err != nil {
			reg := regexp.MustCompile("Rule with rank [0-9]+ is not allowed at order [0-9]+")
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") && reg.MatchString(err.Error()) {
				return infer.CreateResponse[FirewallIPSRuleState]{}, fmt.Errorf("error creating firewall IPS rule: %s, check order %d vs rank %d, err:%s", apiReq.Name, intendedOrder, intendedRank, err)
			}
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Creating firewall ips rule name: %v, got INVALID_INPUT_ARGUMENT\n", apiReq.Name)
				if time.Since(start) < timeout {
					time.Sleep(10 * time.Second)
					continue
				}
			}
			return infer.CreateResponse[FirewallIPSRuleState]{}, fmt.Errorf("creating firewall IPS rule: %w", err)
		}

		reorderWithBeforeReorder(
			OrderRule{Order: intendedOrder, Rank: intendedRank},
			resp.ID,
			firewallIPSResourceType,
			func() (int, error) {
				allRules, err := firewallipscontrolpolicies.GetAll(ctx, svc)
				if err != nil {
					return 0, err
				}
				return len(allRules), nil
			},
			func(id int, order OrderRule) error {
				rule, err := firewallipscontrolpolicies.Get(ctx, svc, id)
				if err != nil {
					return err
				}
				// to avoid the STALE_CONFIGURATION_ERROR
				rule.LastModifiedTime = 0
				rule.LastModifiedBy = nil
				rule.Order = order.Order
				rule.Rank = order.Rank
				_, err = firewallipscontrolpolicies.Update(ctx, svc, id, rule)
				return err
			},
			nil,
		)
		markOrderRuleAsDone(resp.ID, firewallIPSResourceType)
		waitForReorder(firewallIPSResourceType)

		if shouldActivate() {
			time.Sleep(2 * time.Second)
			if activationErr := triggerActivation(ctx, client); activationErr != nil {
				return infer.CreateResponse[FirewallIPSRuleState]{}, activationErr
			}
		} else {
			log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
		}

		rule, err := firewallipscontrolpolicies.Get(ctx, svc, resp.ID)
		if err != nil {
			return infer.CreateResponse[FirewallIPSRuleState]{ID: strconv.Itoa(resp.ID), Output: FirewallIPSRuleState{FirewallIPSRuleArgs: req.Inputs, RuleID: intPtr(resp.ID)}}, nil
		}
		return infer.CreateResponse[FirewallIPSRuleState]{ID: strconv.Itoa(resp.ID), Output: firewallIPSRuleAPIToState(rule)}, nil
	}
}

func (FirewallIPSRule) Read(ctx context.Context, req infer.ReadRequest[FirewallIPSRuleArgs, FirewallIPSRuleState]) (infer.ReadResponse[FirewallIPSRuleArgs, FirewallIPSRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[FirewallIPSRuleArgs, FirewallIPSRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		rule, lookupErr := firewallipscontrolpolicies.GetByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[FirewallIPSRuleArgs, FirewallIPSRuleState]{}, fmt.Errorf("firewall IPS rule not found")
			}
			return infer.ReadResponse[FirewallIPSRuleArgs, FirewallIPSRuleState]{}, lookupErr
		}
		id = rule.ID
	}

	rule, err := firewallipscontrolpolicies.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[FirewallIPSRuleArgs, FirewallIPSRuleState]{}, fmt.Errorf("firewall IPS rule not found")
		}
		return infer.ReadResponse[FirewallIPSRuleArgs, FirewallIPSRuleState]{}, err
	}

	state := firewallIPSRuleAPIToState(rule)
	return infer.ReadResponse[FirewallIPSRuleArgs, FirewallIPSRuleState]{ID: req.ID, Inputs: state.FirewallIPSRuleArgs, State: state}, nil
}

func (FirewallIPSRule) Update(ctx context.Context, req infer.UpdateRequest[FirewallIPSRuleArgs, FirewallIPSRuleState]) (infer.UpdateResponse[FirewallIPSRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[FirewallIPSRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[FirewallIPSRuleState]{}, fmt.Errorf("invalid firewall IPS rule ID: %s", req.ID)
	}
	apiReq := firewallIPSRuleArgsToAPI(&req.Inputs, id)

	existingRules, err := firewallipscontrolpolicies.GetAll(ctx, svc)
	if err == nil && len(existingRules) > 0 {
		sort.Slice(existingRules, func(i, j int) bool {
			return existingRules[i].Rank < existingRules[j].Rank || (existingRules[i].Rank == existingRules[j].Rank && existingRules[i].Order < existingRules[j].Order)
		})
		apiReq.Rank = 7
		apiReq.Order = existingRules[len(existingRules)-1].Order
	}

	_, err = firewallipscontrolpolicies.Update(ctx, svc, id, &apiReq)
	if customErr := failFastOnErrorCodes(err); customErr != nil {
		return infer.UpdateResponse[FirewallIPSRuleState]{}, customErr
	}
	if err != nil {
		return infer.UpdateResponse[FirewallIPSRuleState]{}, err
	}

	intendedOrder := req.Inputs.Order
	intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)

	reorderWithBeforeReorder(
		OrderRule{Order: intendedOrder, Rank: intendedRank},
		id,
		firewallIPSResourceType,
		func() (int, error) {
			allRules, err := firewallipscontrolpolicies.GetAll(ctx, svc)
			if err != nil {
				return 0, err
			}
			return len(allRules), nil
		},
		func(ruleID int, order OrderRule) error {
			rule, err := firewallipscontrolpolicies.Get(ctx, svc, ruleID)
			if err != nil {
				return err
			}
			// to avoid the STALE_CONFIGURATION_ERROR
			rule.LastModifiedTime = 0
			rule.LastModifiedBy = nil
			rule.Order = order.Order
			rule.Rank = order.Rank
			_, err = firewallipscontrolpolicies.Update(ctx, svc, ruleID, rule)
			return err
		},
		nil,
	)
	markOrderRuleAsDone(id, firewallIPSResourceType)
	waitForReorder(firewallIPSResourceType)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[FirewallIPSRuleState]{}, activationErr
		}
	}

	updated, err := firewallipscontrolpolicies.Get(ctx, svc, id)
	if err != nil {
		return infer.UpdateResponse[FirewallIPSRuleState]{Output: FirewallIPSRuleState{FirewallIPSRuleArgs: req.Inputs, RuleID: intPtr(id)}}, nil
	}
	return infer.UpdateResponse[FirewallIPSRuleState]{Output: firewallIPSRuleAPIToState(updated)}, nil
}

func (FirewallIPSRule) Delete(ctx context.Context, req infer.DeleteRequest[FirewallIPSRuleState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid firewall IPS rule ID: %s", req.ID)
	}
	rule, err := firewallipscontrolpolicies.Get(ctx, svc, id)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("error retrieving firewall IPS rule %d: %w", id, err)
	}
	if rule.Predefined {
		return infer.DeleteResponse{}, fmt.Errorf("deletion of predefined rule '%s' is not allowed", rule.Name)
	}
	if _, err := firewallipscontrolpolicies.Delete(ctx, svc, id); err != nil {
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

func (FirewallIPSRule) Annotate(a infer.Annotator) {
	describeResource(a, &FirewallIPSRule{}, "The zia_firewall_ips_rule resource manages firewall IPS (Intrusion Prevention System) rules in the Zscaler Internet Access (ZIA) cloud service. IPS rules allow you to detect and prevent network intrusions by inspecting traffic for known threat signatures and anomalous patterns.\n\nFor more information, see the [ZIA IPS Control Policies documentation](https://help.zscaler.com/zia/ips-control-policies).\n\n{{% examples %}}\n## Example Usage\n\n{{% example %}}\n### Basic Firewall IPS Rule\n\n"+tripleBacktick("typescript")+"\nimport * as zia from \"@bdzscaler/pulumi-zia\";\n\nconst example = new zia.FirewallIPSRule(\"example\", {\n    name: \"Example IPS Rule\",\n    description: \"Block intrusion attempts\",\n    order: 1,\n    state: \"ENABLED\",\n    action: \"BLOCK_DROP\",\n});\n"+tripleBacktick("")+"\n\n"+tripleBacktick("python")+"\nimport zscaler_pulumi_zia as zia\n\nexample = zia.FirewallIPSRule(\"example\",\n    name=\"Example IPS Rule\",\n    description=\"Block intrusion attempts\",\n    order=1,\n    state=\"ENABLED\",\n    action=\"BLOCK_DROP\",\n)\n"+tripleBacktick("")+"\n\n"+tripleBacktick("go")+"\nimport (\n\t\"github.com/pulumi/pulumi/sdk/v3/go/pulumi\"\n\tzia \"github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia\"\n)\n\nfunc main() {\n\tpulumi.Run(func(ctx *pulumi.Context) error {\n\t\t_, err := zia.NewFirewallIPSRule(ctx, \"example\", &zia.FirewallIPSRuleArgs{\n\t\t\tName:        pulumi.String(\"Example IPS Rule\"),\n\t\t\tDescription: pulumi.StringRef(\"Block intrusion attempts\"),\n\t\t\tOrder:       pulumi.Int(1),\n\t\t\tState:       pulumi.StringRef(\"ENABLED\"),\n\t\t\tAction:      pulumi.StringRef(\"BLOCK_DROP\"),\n\t\t})\n\t\treturn err\n\t})\n}\n"+tripleBacktick("")+"\n\n"+tripleBacktick("yaml")+"\nresources:\n  example:\n    type: zia:FirewallIPSRule\n    properties:\n      name: Example IPS Rule\n      description: Block intrusion attempts\n      order: 1\n      state: ENABLED\n      action: BLOCK_DROP\n"+tripleBacktick("")+"\n\n{{% /example %}}\n{{% /examples %}}\n\n## Import\n\nAn existing Firewall IPS Rule can be imported using its resource ID, e.g.\n\n"+tripleBacktick("sh")+"\n$ pulumi import zia:index:FirewallIPSRule example 12345\n"+tripleBacktick("")+"\n")
}

func (a *FirewallIPSRuleArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the firewall IPS rule. Must be unique.")
	ann.Describe(&a.Order, "The order of execution of the rule with respect to other firewall IPS rules.")
	ann.Describe(&a.Description, "Additional information about the firewall IPS rule.")
	ann.Describe(&a.Rank, "Admin rank of the firewall IPS policy rule. Valid values: 0-7. Default: 7.")
	ann.Describe(&a.State, "Rule state. Valid values: `ENABLED`, `DISABLED`.")
	ann.Describe(&a.Action, "The action the rule takes when traffic matches. Valid values: `ALLOW`, `BLOCK_DROP`, `BLOCK_RESET`, `BLOCK_ICMP`.")
	ann.Describe(&a.EnableFullLogging, "If set to true, enables full logging for the rule.")
	ann.Describe(&a.CapturePcap, "If set to true, enables packet capture (PCAP) for the rule.")
	ann.Describe(&a.DefaultRule, "Indicates whether this is the default firewall IPS rule.")
	ann.Describe(&a.Predefined, "Indicates whether this is a predefined rule.")
	ann.Describe(&a.IsEunEnabled, "If set to true, enables end user notification for the rule.")
	ann.Describe(&a.EunTemplateId, "The ID of the end user notification template associated with the rule.")
	ann.Describe(&a.SrcIps, "Source IP addresses or CIDR ranges for the rule.")
	ann.Describe(&a.DestAddresses, "Destination IP addresses, FQDNs, or wildcard FQDNs for the rule.")
	ann.Describe(&a.DestIpCategories, "Destination IP address URL categories for the rule.")
	ann.Describe(&a.ResCategories, "URL categories that apply to the response for the rule.")
	ann.Describe(&a.DestCountries, "Destination countries (ISO 3166-1 alpha-2 codes) for the rule.")
	ann.Describe(&a.SourceCountries, "Source countries (ISO 3166-1 alpha-2 codes) for the rule.")
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
	ann.Describe(&a.ThreatCategories, "IDs of threat categories to which the rule applies.")
	ann.Describe(&a.ZpaAppSegments, "List of ZPA application segments for which this rule is applicable. This field is applicable only for the ZPA gateway forwarding method.")
}

func (s *FirewallIPSRuleState) Annotate(a infer.Annotator) {
	a.Describe(&s.RuleID, "The system-generated ID of the firewall IPS rule.")
}

var _ infer.CustomResource[FirewallIPSRuleArgs, FirewallIPSRuleState] = FirewallIPSRule{}
