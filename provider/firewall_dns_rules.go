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

// Package provider implements the Firewall DNS Rules resource.
// Uses firewalldnscontrolpolicies package. Predefined rules cannot be deleted.

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
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/firewalldnscontrolpolicies"
)

const firewallDNSResourceType = "firewall_dns_rule"

// FirewallDNSRule implements the zia:index:FirewallDNSRule resource.
type FirewallDNSRule struct{}

// FirewallDNSRuleArgs are the inputs.
type FirewallDNSRuleArgs struct {
	Name                string   `pulumi:"name"`
	Order               int      `pulumi:"order"`
	Description         *string  `pulumi:"description,optional"`
	Rank                *int     `pulumi:"rank,optional"`
	State               *string  `pulumi:"state,optional"`
	Action              *string  `pulumi:"action,optional"`
	RedirectIP          *string  `pulumi:"redirectIp,optional"`
	BlockResponseCode   *string  `pulumi:"blockResponseCode,optional"`
	DefaultRule         *bool    `pulumi:"defaultRule,optional"`
	Predefined          *bool    `pulumi:"predefined,optional"`
	CapturePcap         *bool    `pulumi:"capturePcap,optional"`
	IsWebEunEnabled     *bool    `pulumi:"isWebEunEnabled,optional"`
	SrcIps              []string `pulumi:"srcIps,optional"`
	DestAddresses       []string `pulumi:"destAddresses,optional"`
	DestIpCategories    []string `pulumi:"destIpCategories,optional"`
	DestCountries       []string `pulumi:"destCountries,optional"`
	SourceCountries     []string `pulumi:"sourceCountries,optional"`
	ResCategories       []string `pulumi:"resCategories,optional"`
	Applications        []string `pulumi:"applications,optional"`
	DNSRuleRequestTypes []string `pulumi:"dnsRuleRequestTypes,optional"`
	Protocols           []string `pulumi:"protocols,optional"`
	Locations           []int    `pulumi:"locations,optional"`
	LocationGroups      []int    `pulumi:"locationGroups,optional"`
	Departments         []int    `pulumi:"departments,optional"`
	Groups              []int    `pulumi:"groups,optional"`
	Users               []int    `pulumi:"users,optional"`
	TimeWindows         []int    `pulumi:"timeWindows,optional"`
	Labels              []int    `pulumi:"labels,optional"`
	SrcIpGroups         []int    `pulumi:"srcIpGroups,optional"`
	SrcIpv6Groups       []int    `pulumi:"srcIpv6Groups,optional"`
	DestIpGroups        []int    `pulumi:"destIpGroups,optional"`
	DestIpv6Groups      []int    `pulumi:"destIpv6Groups,optional"`
	DeviceGroups        []int    `pulumi:"deviceGroups,optional"`
	Devices             []int    `pulumi:"devices,optional"`
	DNSGateway          *int     `pulumi:"dnsGateway,optional"`
	ZpaIpGroup          *int     `pulumi:"zpaIpGroup,optional"`
	EdnsEcsObject       *int     `pulumi:"ednsEcsObject,optional"`
}

// FirewallDNSRuleState is the persisted state.
type FirewallDNSRuleState struct {
	FirewallDNSRuleArgs
	RuleID *int `pulumi:"ruleId"`
}

func firewallDNSRuleArgsToAPI(args *FirewallDNSRuleArgs, id int) firewalldnscontrolpolicies.FirewallDNSRules {
	order := args.Order
	if order == 0 {
		order = 1
	}
	rank := ptrToIntDefault(args.Rank, 7)
	state := ptrToString(args.State)
	if state == "" {
		state = "ENABLED"
	}
	return firewalldnscontrolpolicies.FirewallDNSRules{
		ID:                  id,
		Name:                args.Name,
		Order:               order,
		Rank:                rank,
		State:               state,
		Action:              ptrToString(args.Action),
		Description:         ptrToString(args.Description),
		RedirectIP:          ptrToString(args.RedirectIP),
		BlockResponseCode:   ptrToString(args.BlockResponseCode),
		DefaultRule:         ptrToBool(args.DefaultRule),
		Predefined:          ptrToBool(args.Predefined),
		CapturePCAP:         ptrToBool(args.CapturePcap),
		IsWebEUNEnabled:     ptrToBool(args.IsWebEunEnabled),
		SrcIps:              args.SrcIps,
		DestAddresses:       args.DestAddresses,
		DestIpCategories:    args.DestIpCategories,
		DestCountries:       processCountries(args.DestCountries),
		SourceCountries:     processCountries(args.SourceCountries),
		ResCategories:       args.ResCategories,
		Applications:        args.Applications,
		DNSRuleRequestTypes: args.DNSRuleRequestTypes,
		Protocols:           args.Protocols,
		Locations:           idsToIDNameExtensions(args.Locations),
		LocationsGroups:     idsToIDNameExtensions(args.LocationGroups),
		Departments:         idsToIDNameExtensions(args.Departments),
		Groups:              idsToIDNameExtensions(args.Groups),
		Users:               idsToIDNameExtensions(args.Users),
		TimeWindows:         idsToIDNameExtensions(args.TimeWindows),
		Labels:              idsToIDNameExtensions(args.Labels),
		SrcIpGroups:         idsToIDNameExtensions(args.SrcIpGroups),
		SrcIpv6Groups:       idsToIDNameExtensions(args.SrcIpv6Groups),
		DestIpGroups:        idsToIDNameExtensions(args.DestIpGroups),
		DestIpv6Groups:      idsToIDNameExtensions(args.DestIpv6Groups),
		DeviceGroups:        idsToIDNameExtensions(args.DeviceGroups),
		Devices:             idsToIDNameExtensions(args.Devices),
		DNSGateway:          idToOptionalIDName(ptrToIntDefault(args.DNSGateway, 0)), // idToOptionalIDName returns nil when id is 0
		ZPAIPGroup:          idToOptionalIDName(ptrToIntDefault(args.ZpaIpGroup, 0)),
		EDNSEcsObject:       idToOptionalIDName(ptrToIntDefault(args.EdnsEcsObject, 0)),
	}
}

func firewallDNSRuleAPIToState(api *firewalldnscontrolpolicies.FirewallDNSRules) FirewallDNSRuleState {
	return FirewallDNSRuleState{
		FirewallDNSRuleArgs: FirewallDNSRuleArgs{
			Name:                api.Name,
			Order:               api.Order,
			Description:         stringPtr(api.Description),
			Rank:                intPtr(api.Rank),
			State:               stringPtr(api.State),
			Action:              stringPtr(api.Action),
			RedirectIP:          stringPtr(api.RedirectIP),
			BlockResponseCode:   stringPtr(api.BlockResponseCode),
			DefaultRule:         boolPtr(api.DefaultRule),
			Predefined:          boolPtr(api.Predefined),
			CapturePcap:         boolPtr(api.CapturePCAP),
			IsWebEunEnabled:     boolPtr(api.IsWebEUNEnabled),
			SrcIps:              api.SrcIps,
			DestAddresses:       api.DestAddresses,
			DestIpCategories:    api.DestIpCategories,
			DestCountries:       processCountriesFromAPI(api.DestCountries),
			SourceCountries:     processCountriesFromAPI(api.SourceCountries),
			ResCategories:       api.ResCategories,
			Applications:        api.Applications,
			DNSRuleRequestTypes: api.DNSRuleRequestTypes,
			Protocols:           api.Protocols,
			Locations:           idNameExtensionsToIDs(api.Locations),
			LocationGroups:      idNameExtensionsToIDs(api.LocationsGroups),
			Departments:         idNameExtensionsToIDs(api.Departments),
			Groups:              idNameExtensionsToIDs(api.Groups),
			Users:               idNameExtensionsToIDs(api.Users),
			TimeWindows:         idNameExtensionsToIDs(api.TimeWindows),
			Labels:              idNameExtensionsToIDs(api.Labels),
			SrcIpGroups:         idNameExtensionsToIDs(api.SrcIpGroups),
			SrcIpv6Groups:       idNameExtensionsToIDs(api.SrcIpv6Groups),
			DestIpGroups:        idNameExtensionsToIDs(api.DestIpGroups),
			DestIpv6Groups:      idNameExtensionsToIDs(api.DestIpv6Groups),
			DeviceGroups:        idNameExtensionsToIDs(api.DeviceGroups),
			Devices:             idNameExtensionsToIDs(api.Devices),
			DNSGateway:          idNameToOptionalID(api.DNSGateway),
			ZpaIpGroup:          idNameToOptionalID(api.ZPAIPGroup),
			EdnsEcsObject:       idNameToOptionalID(api.EDNSEcsObject),
		},
		RuleID: intPtr(api.ID),
	}
}

func (FirewallDNSRule) Diff(ctx context.Context, req infer.DiffRequest[FirewallDNSRuleArgs, FirewallDNSRuleState]) (infer.DiffResponse, error) {
	diff, hasChanges := diffSkippingUnsetInputs(req.State.FirewallDNSRuleArgs, req.Inputs)
	return infer.DiffResponse{HasChanges: hasChanges, DetailedDiff: diff}, nil
}

func (FirewallDNSRule) Create(ctx context.Context, req infer.CreateRequest[FirewallDNSRuleArgs]) (infer.CreateResponse[FirewallDNSRuleState], error) {
	if req.DryRun {
		s := FirewallDNSRuleState{FirewallDNSRuleArgs: req.Inputs, RuleID: intPtr(0)}
		return infer.CreateResponse[FirewallDNSRuleState]{ID: "preview", Output: s}, nil
	}
	if req.Inputs.Order < 1 {
		return infer.CreateResponse[FirewallDNSRuleState]{}, fmt.Errorf("order must be a positive whole number (>= 1), got %d", req.Inputs.Order)
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[FirewallDNSRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := firewallDNSRuleArgsToAPI(&req.Inputs, 0)
	timeout := 60 * time.Minute
	start := time.Now()

	intendedOrder := apiReq.Order
	intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)

	for {
		select {
		case firewallDNSSem <- struct{}{}:
		case <-ctx.Done():
			return infer.CreateResponse[FirewallDNSRuleState]{}, fmt.Errorf("creating firewall DNS rule: %w", ctx.Err())
		}

		firewallDNSOrderMu.Lock()
		if firewallDNSStartingOrder == 0 {
			list, _ := firewalldnscontrolpolicies.GetAll(ctx, svc)
			for _, r := range list {
				if r.Order > firewallDNSStartingOrder {
					firewallDNSStartingOrder = r.Order
				}
			}
			if firewallDNSStartingOrder == 0 {
				firewallDNSStartingOrder = 1
			}
		}

		if intendedRank < 7 {
			apiReq.Rank = 7
		}
		apiReq.Order = firewallDNSStartingOrder
		firewallDNSOrderMu.Unlock()

		resp, err := firewalldnscontrolpolicies.Create(ctx, svc, &apiReq)

		if err == nil {
			firewallDNSOrderMu.Lock()
			firewallDNSStartingOrder++
			firewallDNSOrderMu.Unlock()
		}

		<-firewallDNSSem

		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.CreateResponse[FirewallDNSRuleState]{}, customErr
		}
		if err != nil {
			reg := regexp.MustCompile("Rule with rank [0-9]+ is not allowed at order [0-9]+")
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") && reg.MatchString(err.Error()) {
				return infer.CreateResponse[FirewallDNSRuleState]{}, fmt.Errorf("error creating firewall DNS rule: %s, check order %d vs rank %d, err:%s",
					apiReq.Name, intendedOrder, intendedRank, err)
			}
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Creating firewall DNS rule name: %v, got INVALID_INPUT_ARGUMENT\n", apiReq.Name)
				firewallDNSOrderMu.Lock()
				firewallDNSStartingOrder = 0
				firewallDNSOrderMu.Unlock()
				if time.Since(start) < timeout {
					select {
					case <-time.After(10 * time.Second):
					case <-ctx.Done():
						return infer.CreateResponse[FirewallDNSRuleState]{}, fmt.Errorf("creating firewall DNS rule: %w", ctx.Err())
					}
					continue
				}
			}
			return infer.CreateResponse[FirewallDNSRuleState]{}, fmt.Errorf("creating firewall DNS rule: %w", err)
		}

		reorderWithBeforeReorder(
			OrderRule{Order: intendedOrder, Rank: intendedRank},
			resp.ID,
			firewallDNSResourceType,
			func() (int, error) {
				allRules, err := firewalldnscontrolpolicies.GetAll(ctx, svc)
				if err != nil {
					return 0, err
				}
				return len(allRules), nil
			},
			func(id int, order OrderRule) error {
				rule, err := firewalldnscontrolpolicies.Get(ctx, svc, id)
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
				_, err = firewalldnscontrolpolicies.Update(ctx, svc, id, rule)
				return err
			},
			nil,
		)

		markOrderRuleAsDone(resp.ID, firewallDNSResourceType)
		waitForReorder(firewallDNSResourceType)

		if shouldActivate() {
			time.Sleep(2 * time.Second)
			if activationErr := triggerActivation(ctx, client); activationErr != nil {
				return infer.CreateResponse[FirewallDNSRuleState]{}, activationErr
			}
		} else {
			log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
		}

		rule, err := firewalldnscontrolpolicies.Get(ctx, svc, resp.ID)
		if err != nil {
			state := FirewallDNSRuleState{FirewallDNSRuleArgs: req.Inputs, RuleID: intPtr(resp.ID)}
			return infer.CreateResponse[FirewallDNSRuleState]{ID: strconv.Itoa(resp.ID), Output: state}, nil
		}
		return infer.CreateResponse[FirewallDNSRuleState]{
			ID:     strconv.Itoa(resp.ID),
			Output: firewallDNSRuleAPIToState(rule),
		}, nil
	}
}

func (FirewallDNSRule) Read(ctx context.Context, req infer.ReadRequest[FirewallDNSRuleArgs, FirewallDNSRuleState]) (infer.ReadResponse[FirewallDNSRuleArgs, FirewallDNSRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[FirewallDNSRuleArgs, FirewallDNSRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		rule, lookupErr := firewalldnscontrolpolicies.GetByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[FirewallDNSRuleArgs, FirewallDNSRuleState]{}, fmt.Errorf("firewall DNS rule not found")
			}
			return infer.ReadResponse[FirewallDNSRuleArgs, FirewallDNSRuleState]{}, lookupErr
		}
		id = rule.ID
	}

	rule, err := firewalldnscontrolpolicies.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[FirewallDNSRuleArgs, FirewallDNSRuleState]{}, fmt.Errorf("firewall DNS rule not found")
		}
		return infer.ReadResponse[FirewallDNSRuleArgs, FirewallDNSRuleState]{}, err
	}

	state := firewallDNSRuleAPIToState(rule)
	return infer.ReadResponse[FirewallDNSRuleArgs, FirewallDNSRuleState]{
		ID:     req.ID,
		Inputs: state.FirewallDNSRuleArgs,
		State:  state,
	}, nil
}

func (FirewallDNSRule) Update(ctx context.Context, req infer.UpdateRequest[FirewallDNSRuleArgs, FirewallDNSRuleState]) (infer.UpdateResponse[FirewallDNSRuleState], error) {
	if req.Inputs.Order < 1 {
		return infer.UpdateResponse[FirewallDNSRuleState]{}, fmt.Errorf("order must be a positive whole number (>= 1), got %d", req.Inputs.Order)
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[FirewallDNSRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[FirewallDNSRuleState]{}, fmt.Errorf("invalid firewall DNS rule ID: %s", req.ID)
	}
	apiReq := firewallDNSRuleArgsToAPI(&req.Inputs, id)

	// Store intended order/rank from inputs before overwriting (matches Terraform)
	intendedOrder := apiReq.Order
	intendedRank := apiReq.Rank

	existingRules, err := firewalldnscontrolpolicies.GetAll(ctx, svc)
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
		if _, err := firewalldnscontrolpolicies.Update(ctx, svc, id, &apiReq); err != nil {
			if customErr := failFastOnErrorCodes(err); customErr != nil {
				return infer.UpdateResponse[FirewallDNSRuleState]{}, customErr
			}
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Updating firewall DNS rule ID: %v, got INVALID_INPUT_ARGUMENT\n", id)
				if time.Since(start) < timeout {
					time.Sleep(10 * time.Second)
					continue
				}
			}
			return infer.UpdateResponse[FirewallDNSRuleState]{}, err
		}

		reorderWithBeforeReorder(
			OrderRule{Order: intendedOrder, Rank: intendedRank},
			id,
			firewallDNSResourceType,
			func() (int, error) {
				allRules, err := firewalldnscontrolpolicies.GetAll(ctx, svc)
				if err != nil {
					return 0, err
				}
				return len(allRules), nil
			},
			func(ruleID int, order OrderRule) error {
				rule, err := firewalldnscontrolpolicies.Get(ctx, svc, ruleID)
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
				_, err = firewalldnscontrolpolicies.Update(ctx, svc, ruleID, rule)
				return err
			},
			nil,
		)

		markOrderRuleAsDone(id, firewallDNSResourceType)
		waitForReorder(firewallDNSResourceType)
		break
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[FirewallDNSRuleState]{}, activationErr
		}
	}

	updated, err := firewalldnscontrolpolicies.Get(ctx, svc, id)
	if err != nil {
		return infer.UpdateResponse[FirewallDNSRuleState]{Output: FirewallDNSRuleState{
			FirewallDNSRuleArgs: req.Inputs,
			RuleID:              intPtr(id),
		}}, nil
	}
	return infer.UpdateResponse[FirewallDNSRuleState]{Output: firewallDNSRuleAPIToState(updated)}, nil
}

func (FirewallDNSRule) Delete(ctx context.Context, req infer.DeleteRequest[FirewallDNSRuleState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid firewall DNS rule ID: %s", req.ID)
	}
	rule, err := firewalldnscontrolpolicies.Get(ctx, svc, id)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("error retrieving firewall DNS rule %d: %w", id, err)
	}
	if rule.Predefined {
		return infer.DeleteResponse{}, fmt.Errorf("deletion of predefined rule '%s' is not allowed", rule.Name)
	}
	if _, err := firewalldnscontrolpolicies.Delete(ctx, svc, id); err != nil {
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

func (FirewallDNSRule) Annotate(a infer.Annotator) {
	describeResource(a, &FirewallDNSRule{}, `The zia_firewall_dns_rule resource manages firewall DNS control rules in the Zscaler Internet Access (ZIA) cloud service. DNS control rules allow you to control DNS traffic by allowing, blocking, or redirecting DNS requests based on various criteria such as source, destination, applications, and DNS request types.

For more information, see the [ZIA DNS Control Policies documentation](https://help.zscaler.com/zia/dns-control-policies).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Firewall DNS Rule

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.FirewallDNSRule("example", {
    name: "Example DNS Rule",
    description: "Block malicious DNS requests",
    order: 1,
    state: "ENABLED",
    action: "BLOCK_DROP",
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.FirewallDNSRule("example",
    name="Example DNS Rule",
    description="Block malicious DNS requests",
    order=1,
    state="ENABLED",
    action="BLOCK_DROP",
)
`+tripleBacktick("")+`

`+tripleBacktick("go")+`
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zia "github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := zia.NewFirewallDNSRule(ctx, "example", &zia.FirewallDNSRuleArgs{
			Name:        pulumi.String("Example DNS Rule"),
			Description: pulumi.StringRef("Block malicious DNS requests"),
			Order:       pulumi.Int(1),
			State:       pulumi.StringRef("ENABLED"),
			Action:      pulumi.StringRef("BLOCK_DROP"),
		})
		return err
	})
}
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:FirewallDNSRule
    properties:
      name: Example DNS Rule
      description: Block malicious DNS requests
      order: 1
      state: ENABLED
      action: BLOCK_DROP
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing Firewall DNS Rule can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:FirewallDNSRule example 12345
`+tripleBacktick("")+`
`)
}

func (a *FirewallDNSRuleArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the firewall DNS rule. Must be unique.")
	ann.Describe(&a.Order, "The order of execution of the rule with respect to other firewall DNS rules.")
	ann.Describe(&a.Description, "Additional information about the firewall DNS rule.")
	ann.Describe(&a.Rank, "Admin rank of the firewall DNS policy rule. Valid values: 0-7. Default: 7.")
	ann.Describe(&a.State, "Rule state. Valid values: `ENABLED`, `DISABLED`.")
	ann.Describe(&a.Action, "The action the rule takes when traffic matches. Valid values: `ALLOW`, `BLOCK_DROP`, `BLOCK_RESET`, `BLOCK_ICMP`, `REDIR_REQ`.")
	ann.Describe(&a.RedirectIP, "The IP address to redirect DNS requests to. Required when action is `REDIR_REQ`.")
	ann.Describe(&a.BlockResponseCode, "The DNS response code to return when blocking. Valid values: `ANY`, `NONE`, `FORMERR`, `SERVFAIL`, `NXDOMAIN`, `NOTIMP`, `REFUSED`, `NOTAUTH`, `NXRRSET`.")
	ann.Describe(&a.DefaultRule, "Indicates whether this is the default firewall DNS rule.")
	ann.Describe(&a.Predefined, "Indicates whether this is a predefined rule.")
	ann.Describe(&a.CapturePcap, "If set to true, enables packet capture (PCAP) for the rule.")
	ann.Describe(&a.IsWebEunEnabled, "If set to true, enables web end user notification for the rule.")
	ann.Describe(&a.SrcIps, "Source IP addresses or CIDR ranges for the rule.")
	ann.Describe(&a.DestAddresses, "Destination IP addresses, FQDNs, or wildcard FQDNs for the rule.")
	ann.Describe(&a.DestIpCategories, "Destination IP address URL categories for the rule.")
	ann.Describe(&a.DestCountries, "Destination countries (ISO 3166-1 alpha-2 codes) for the rule.")
	ann.Describe(&a.SourceCountries, "Source countries (ISO 3166-1 alpha-2 codes) for the rule.")
	ann.Describe(&a.ResCategories, "URL categories that apply to the response for the rule.")
	ann.Describe(&a.Applications, "DNS application values to which the rule applies.")
	ann.Describe(&a.DNSRuleRequestTypes, "DNS request types to which the rule applies. Valid values: `A`, `AAAA`, `CNAME`, `MX`, `NS`, `SOA`, `TXT`, `SRV`, `PTR`, `ANY`.")
	ann.Describe(&a.Protocols, "Protocols to which the rule applies. Valid values: `ANY_RULE`, `TCP_RULE`, `UDP_RULE`.")
	ann.Describe(&a.Locations, "IDs of locations to which the rule must be applied.")
	ann.Describe(&a.LocationGroups, "IDs of location groups to which the rule must be applied.")
	ann.Describe(&a.Departments, "IDs of departments to which the rule must be applied.")
	ann.Describe(&a.Groups, "IDs of groups to which the rule must be applied.")
	ann.Describe(&a.Users, "IDs of users to which the rule must be applied.")
	ann.Describe(&a.TimeWindows, "IDs of time intervals during which the rule must be enforced.")
	ann.Describe(&a.Labels, "IDs of labels associated with the rule.")
	ann.Describe(&a.SrcIpGroups, "IDs of source IP address groups for the rule.")
	ann.Describe(&a.SrcIpv6Groups, "IDs of source IPv6 address groups for the rule.")
	ann.Describe(&a.DestIpGroups, "IDs of destination IP address groups for the rule.")
	ann.Describe(&a.DestIpv6Groups, "IDs of destination IPv6 address groups for the rule.")
	ann.Describe(&a.DeviceGroups, "IDs of device groups for which the rule must be applied. Applicable for devices managed using Zscaler Client Connector.")
	ann.Describe(&a.Devices, "IDs of devices for which the rule must be applied.")
	ann.Describe(&a.DNSGateway, "The ID of the DNS gateway associated with the rule.")
	ann.Describe(&a.ZpaIpGroup, "The ID of the ZPA IP group associated with the rule.")
	ann.Describe(&a.EdnsEcsObject, "The ID of the EDNS ECS object associated with the rule.")
}

func (s *FirewallDNSRuleState) Annotate(a infer.Annotator) {
	a.Describe(&s.RuleID, "The system-generated ID of the firewall DNS rule.")
}

var _ infer.CustomResource[FirewallDNSRuleArgs, FirewallDNSRuleState] = FirewallDNSRule{}
