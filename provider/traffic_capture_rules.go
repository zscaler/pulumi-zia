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

// Package provider implements the Traffic Capture Rules resource.
// Adopted from terraform-provider-zia resource_zia_traffic_capture_rules.go.

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
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/traffic_capture"
)

const trafficCaptureResourceType = "firewall_filtering_rules"

// TrafficCaptureRule implements the zia:index:TrafficCaptureRule resource.
type TrafficCaptureRule struct{}

// TrafficCaptureRuleArgs are the inputs.
type TrafficCaptureRuleArgs struct {
	Name                string               `pulumi:"name"`
	Order               int                  `pulumi:"order"`
	Description         *string              `pulumi:"description,optional"`
	Rank                *int                 `pulumi:"rank,optional"`
	Action              *string              `pulumi:"action,optional"`
	State               *string              `pulumi:"state,optional"`
	SrcIps              []string             `pulumi:"srcIps,optional"`
	DestAddresses       []string             `pulumi:"destAddresses,optional"`
	DestIpCategories    []string             `pulumi:"destIpCategories,optional"`
	NwApplications      []string             `pulumi:"nwApplications,optional"`
	DefaultRule         *bool                `pulumi:"defaultRule,optional"`
	Predefined          *bool                `pulumi:"predefined,optional"`
	TxnSizeLimit        *string              `pulumi:"txnSizeLimit,optional"`
	TxnSampling         *string              `pulumi:"txnSampling,optional"`
	Locations           []int                `pulumi:"locations,optional"`
	LocationGroups      []int                `pulumi:"locationGroups,optional"`
	Departments         []int                `pulumi:"departments,optional"`
	Groups              []int                `pulumi:"groups,optional"`
	Users               []int                `pulumi:"users,optional"`
	TimeWindows         []int                `pulumi:"timeWindows,optional"`
	SrcIpGroups         []int                `pulumi:"srcIpGroups,optional"`
	DestIpGroups        []int                `pulumi:"destIpGroups,optional"`
	NwServices          []int                `pulumi:"nwServices,optional"`
	NwServiceGroups     []int                `pulumi:"nwServiceGroups,optional"`
	NwApplicationGroups []int                `pulumi:"nwApplicationGroups,optional"`
	AppServiceGroups    []int                `pulumi:"appServiceGroups,optional"`
	Labels              []int                `pulumi:"labels,optional"`
	DeviceGroups        []int                `pulumi:"deviceGroups,optional"`
	Devices             []int                `pulumi:"devices,optional"`
	DestCountries       []string             `pulumi:"destCountries,optional"`
	SourceCountries     []string             `pulumi:"sourceCountries,optional"`
	DeviceTrustLevels   []string             `pulumi:"deviceTrustLevels,optional"`
	WorkloadGroups      []WorkloadGroupInput `pulumi:"workloadGroups,optional"`
}

// TrafficCaptureRuleState is the persisted state.
type TrafficCaptureRuleState struct {
	TrafficCaptureRuleArgs
	RuleID *int `pulumi:"ruleId"`
}

func trafficCaptureRuleArgsToAPI(args *TrafficCaptureRuleArgs, id int) traffic_capture.TrafficCaptureRules {
	order := args.Order
	if order == 0 {
		order = 1
	}
	rank := ptrToIntDefault(args.Rank, 7)
	action := ptrToString(args.Action)
	if action == "" {
		action = "CAPTURE"
	}
	state := ptrToString(args.State)
	if state == "" {
		state = "ENABLED"
	}
	txnSizeLimit := ptrToString(args.TxnSizeLimit)
	if txnSizeLimit == "" {
		txnSizeLimit = "NONE"
	}
	txnSampling := ptrToString(args.TxnSampling)
	if txnSampling == "" {
		txnSampling = "NONE"
	}
	api := traffic_capture.TrafficCaptureRules{
		ID:                  id,
		Name:                args.Name,
		Order:               order,
		Rank:                rank,
		Action:              action,
		State:               state,
		Description:         ptrToString(args.Description),
		TxnSizeLimit:        txnSizeLimit,
		TxnSampling:         txnSampling,
		SrcIps:              args.SrcIps,
		DestAddresses:       args.DestAddresses,
		DestIpCategories:    args.DestIpCategories,
		NwApplications:      args.NwApplications,
		DefaultRule:         ptrToBool(args.DefaultRule),
		Predefined:          ptrToBool(args.Predefined),
		DestCountries:       processCountries(args.DestCountries),
		SourceCountries:     processCountries(args.SourceCountries),
		DeviceTrustLevels:   args.DeviceTrustLevels,
		Locations:           idsToIDNameExtensions(args.Locations),
		LocationsGroups:     idsToIDNameExtensions(args.LocationGroups),
		Departments:         idsToIDNameExtensions(args.Departments),
		Groups:              idsToIDNameExtensions(args.Groups),
		Users:               idsToIDNameExtensions(args.Users),
		TimeWindows:         idsToIDNameExtensions(args.TimeWindows),
		SrcIpGroups:         idsToIDNameExtensions(args.SrcIpGroups),
		DestIpGroups:        idsToIDNameExtensions(args.DestIpGroups),
		NwServices:          idsToIDNameExtensions(args.NwServices),
		NwServiceGroups:     idsToIDNameExtensions(args.NwServiceGroups),
		NwApplicationGroups: idsToIDNameExtensions(args.NwApplicationGroups),
		AppServiceGroups:    idsToIDNameExtensions(args.AppServiceGroups),
		Labels:              idsToIDNameExtensions(args.Labels),
		DeviceGroups:        idsToIDNameExtensions(args.DeviceGroups),
		Devices:             idsToIDNameExtensions(args.Devices),
		WorkloadGroups:      expandWorkloadGroups(args.WorkloadGroups),
	}
	return api
}

func trafficCaptureRuleAPIToState(api *traffic_capture.TrafficCaptureRules) TrafficCaptureRuleState {
	destCountries := make([]string, len(api.DestCountries))
	for i, c := range api.DestCountries {
		destCountries[i] = strings.TrimPrefix(c, "COUNTRY_")
	}
	sourceCountries := make([]string, len(api.SourceCountries))
	for i, c := range api.SourceCountries {
		sourceCountries[i] = strings.TrimPrefix(c, "COUNTRY_")
	}
	state := TrafficCaptureRuleState{
		TrafficCaptureRuleArgs: TrafficCaptureRuleArgs{
			Name:                api.Name,
			Order:               api.Order,
			Description:         stringPtr(api.Description),
			Rank:                intPtr(api.Rank),
			Action:              stringPtr(api.Action),
			State:               stringPtr(api.State),
			SrcIps:              api.SrcIps,
			DestAddresses:       api.DestAddresses,
			DestIpCategories:    api.DestIpCategories,
			NwApplications:      api.NwApplications,
			DefaultRule:         boolPtr(api.DefaultRule),
			Predefined:          boolPtr(api.Predefined),
			TxnSizeLimit:        stringPtr(api.TxnSizeLimit),
			TxnSampling:         stringPtr(api.TxnSampling),
			DestCountries:       destCountries,
			SourceCountries:     sourceCountries,
			DeviceTrustLevels:   api.DeviceTrustLevels,
			Locations:           idNameExtensionsToIDs(api.Locations),
			LocationGroups:      idNameExtensionsToIDs(api.LocationsGroups),
			Departments:         idNameExtensionsToIDs(api.Departments),
			Groups:              idNameExtensionsToIDs(api.Groups),
			Users:               idNameExtensionsToIDs(api.Users),
			TimeWindows:         idNameExtensionsToIDs(api.TimeWindows),
			SrcIpGroups:         idNameExtensionsToIDs(api.SrcIpGroups),
			DestIpGroups:        idNameExtensionsToIDs(api.DestIpGroups),
			NwServices:          idNameExtensionsToIDs(api.NwServices),
			NwServiceGroups:     idNameExtensionsToIDs(api.NwServiceGroups),
			NwApplicationGroups: idNameExtensionsToIDs(api.NwApplicationGroups),
			AppServiceGroups:    idNameExtensionsToIDs(api.AppServiceGroups),
			Labels:              idNameExtensionsToIDs(api.Labels),
			DeviceGroups:        idNameExtensionsToIDs(api.DeviceGroups),
			Devices:             idNameExtensionsToIDs(api.Devices),
			WorkloadGroups:      workloadGroupOutputsToInputs(flattenWorkloadGroups(api.WorkloadGroups)),
		},
		RuleID: intPtr(api.ID),
	}
	return state
}

func (TrafficCaptureRule) Create(ctx context.Context, req infer.CreateRequest[TrafficCaptureRuleArgs]) (infer.CreateResponse[TrafficCaptureRuleState], error) {
	if req.DryRun {
		s := TrafficCaptureRuleState{TrafficCaptureRuleArgs: req.Inputs, RuleID: intPtr(0)}
		return infer.CreateResponse[TrafficCaptureRuleState]{ID: "preview", Output: s}, nil
	}
	if req.Inputs.Order < 1 {
		return infer.CreateResponse[TrafficCaptureRuleState]{}, fmt.Errorf("order must be a positive whole number (>= 1), got %d", req.Inputs.Order)
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[TrafficCaptureRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := trafficCaptureRuleArgsToAPI(&req.Inputs, 0)
	timeout := 60 * time.Minute
	start := time.Now()

	for {
		trafficCaptureLock.Lock()
		if trafficCaptureStartingOrder == 0 {
			list, _ := traffic_capture.GetAll(ctx, svc, nil)
			for _, r := range list {
				if r.Order > trafficCaptureStartingOrder {
					trafficCaptureStartingOrder = r.Order
				}
			}
			if trafficCaptureStartingOrder == 0 {
				trafficCaptureStartingOrder = 1
			} else {
				trafficCaptureStartingOrder++
			}
		}
		trafficCaptureLock.Unlock()

		intendedOrder := apiReq.Order
		intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)
		if intendedRank < 7 {
			apiReq.Rank = 7
		}
		apiReq.Order = trafficCaptureStartingOrder

		resp, err := traffic_capture.Create(ctx, svc, &apiReq)
		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.CreateResponse[TrafficCaptureRuleState]{}, customErr
		}
		if err != nil {
			reg := regexp.MustCompile("Rule with rank [0-9]+ is not allowed at order [0-9]+")
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") && reg.MatchString(err.Error()) {
				return infer.CreateResponse[TrafficCaptureRuleState]{}, fmt.Errorf("error creating traffic capture rule: %s, check order %d vs rank %d, err:%s",
					apiReq.Name, intendedOrder, intendedRank, err)
			}
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Creating traffic capture rule name: %v, got INVALID_INPUT_ARGUMENT\n", apiReq.Name)
				if time.Since(start) < timeout {
					time.Sleep(10 * time.Second)
					continue
				}
			}
			return infer.CreateResponse[TrafficCaptureRuleState]{}, fmt.Errorf("creating traffic capture rule: %w", err)
		}

		reorderWithBeforeReorder(
			OrderRule{Order: intendedOrder, Rank: intendedRank},
			resp.ID,
			trafficCaptureResourceType,
			func() (int, error) {
				allRules, err := traffic_capture.GetAll(ctx, svc, nil)
				if err != nil {
					return 0, err
				}
				return len(allRules), nil
			},
			func(id int, order OrderRule) error {
				rule, err := traffic_capture.Get(ctx, svc, id)
				if err != nil {
					return err
				}
				rule.LastModifiedTime = 0
				rule.LastModifiedBy = nil
				// Strip read-only fields that cause "Request body is invalid" for predefined rules
				rule.Predefined = false
				rule.DefaultRule = false
				rule.AccessControl = ""
				rule.Order = order.Order
				rule.Rank = order.Rank
				_, err = traffic_capture.Update(ctx, svc, id, rule)
				return err
			},
			nil,
		)

		markOrderRuleAsDone(resp.ID, trafficCaptureResourceType)
		waitForReorder(trafficCaptureResourceType)

		if shouldActivate() {
			time.Sleep(2 * time.Second)
			if activationErr := triggerActivation(ctx, client); activationErr != nil {
				return infer.CreateResponse[TrafficCaptureRuleState]{}, activationErr
			}
		} else {
			log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
		}

		rule, err := traffic_capture.Get(ctx, svc, resp.ID)
		if err != nil {
			state := TrafficCaptureRuleState{
				TrafficCaptureRuleArgs: req.Inputs,
				RuleID:                 intPtr(resp.ID),
			}
			return infer.CreateResponse[TrafficCaptureRuleState]{ID: strconv.Itoa(resp.ID), Output: state}, nil
		}
		return infer.CreateResponse[TrafficCaptureRuleState]{
			ID:     strconv.Itoa(resp.ID),
			Output: trafficCaptureRuleAPIToState(rule),
		}, nil
	}
}

func (TrafficCaptureRule) Read(ctx context.Context, req infer.ReadRequest[TrafficCaptureRuleArgs, TrafficCaptureRuleState]) (infer.ReadResponse[TrafficCaptureRuleArgs, TrafficCaptureRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[TrafficCaptureRuleArgs, TrafficCaptureRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		rule, lookupErr := traffic_capture.GetByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[TrafficCaptureRuleArgs, TrafficCaptureRuleState]{}, fmt.Errorf("traffic capture rule not found")
			}
			return infer.ReadResponse[TrafficCaptureRuleArgs, TrafficCaptureRuleState]{}, lookupErr
		}
		id = rule.ID
	}

	rule, err := traffic_capture.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[TrafficCaptureRuleArgs, TrafficCaptureRuleState]{}, fmt.Errorf("traffic capture rule not found")
		}
		return infer.ReadResponse[TrafficCaptureRuleArgs, TrafficCaptureRuleState]{}, err
	}
	state := trafficCaptureRuleAPIToState(rule)
	return infer.ReadResponse[TrafficCaptureRuleArgs, TrafficCaptureRuleState]{
		ID:     req.ID,
		Inputs: state.TrafficCaptureRuleArgs,
		State:  state,
	}, nil
}

func (TrafficCaptureRule) Update(ctx context.Context, req infer.UpdateRequest[TrafficCaptureRuleArgs, TrafficCaptureRuleState]) (infer.UpdateResponse[TrafficCaptureRuleState], error) {
	if req.Inputs.Order < 1 {
		return infer.UpdateResponse[TrafficCaptureRuleState]{}, fmt.Errorf("order must be a positive whole number (>= 1), got %d", req.Inputs.Order)
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[TrafficCaptureRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[TrafficCaptureRuleState]{}, fmt.Errorf("invalid traffic capture rule ID: %s", req.ID)
	}
	apiReq := trafficCaptureRuleArgsToAPI(&req.Inputs, id)

	existingRules, err := traffic_capture.GetAll(ctx, svc, nil)
	if err != nil {
		return infer.UpdateResponse[TrafficCaptureRuleState]{}, err
	}
	sort.Slice(existingRules, func(i, j int) bool {
		return existingRules[i].Rank < existingRules[j].Rank || (existingRules[i].Rank == existingRules[j].Rank && existingRules[i].Order < existingRules[j].Order)
	})
	intendedOrder := apiReq.Order
	intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)
	nextAvailableOrder := existingRules[len(existingRules)-1].Order
	apiReq.Rank = 7
	apiReq.Order = nextAvailableOrder

	if _, err = traffic_capture.Update(ctx, svc, id, &apiReq); err != nil {
		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.UpdateResponse[TrafficCaptureRuleState]{}, customErr
		}
		return infer.UpdateResponse[TrafficCaptureRuleState]{}, err
	}

	reorderWithBeforeReorder(
		OrderRule{Order: intendedOrder, Rank: intendedRank},
		id,
		trafficCaptureResourceType,
		func() (int, error) {
			allRules, err := traffic_capture.GetAll(ctx, svc, nil)
			if err != nil {
				return 0, err
			}
			return len(allRules), nil
		},
		func(ruleID int, order OrderRule) error {
			rule, err := traffic_capture.Get(ctx, svc, ruleID)
			if err != nil {
				return err
			}
			if rule.Order == order.Order && rule.Rank == order.Rank {
				return nil
			}
			rule.LastModifiedTime = 0
			rule.LastModifiedBy = nil
			// Strip read-only fields that cause "Request body is invalid" for predefined rules
			rule.Predefined = false
			rule.DefaultRule = false
			rule.AccessControl = ""
			rule.Order = order.Order
			rule.Rank = order.Rank
			_, err = traffic_capture.Update(ctx, svc, ruleID, rule)
			return err
		},
		nil,
	)

	markOrderRuleAsDone(id, trafficCaptureResourceType)
	waitForReorder(trafficCaptureResourceType)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[TrafficCaptureRuleState]{}, activationErr
		}
	}

	updated, err := traffic_capture.Get(ctx, svc, id)
	if err != nil {
		return infer.UpdateResponse[TrafficCaptureRuleState]{Output: TrafficCaptureRuleState{
			TrafficCaptureRuleArgs: req.Inputs,
			RuleID:                 intPtr(id),
		}}, nil
	}
	return infer.UpdateResponse[TrafficCaptureRuleState]{Output: trafficCaptureRuleAPIToState(updated)}, nil
}

func (TrafficCaptureRule) Delete(ctx context.Context, req infer.DeleteRequest[TrafficCaptureRuleState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid traffic capture rule ID: %s", req.ID)
	}
	rule, err := traffic_capture.Get(ctx, svc, id)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("error retrieving traffic capture rule %d: %w", id, err)
	}
	if rule.Predefined {
		return infer.DeleteResponse{}, fmt.Errorf("deletion of predefined rule '%s' is not allowed", rule.Name)
	}
	if _, err := traffic_capture.Delete(ctx, svc, id); err != nil {
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

func (TrafficCaptureRule) Annotate(a infer.Annotator) {
	describeResource(a, &TrafficCaptureRule{}, `The zia.TrafficCaptureRule resource manages traffic capture rules in the Zscaler Internet Access (ZIA) cloud.
Traffic capture rules define criteria for capturing network traffic for analysis, specifying which traffic
to capture based on source/destination IPs, locations, departments, users, applications, and other criteria.

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Traffic Capture Rule

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.TrafficCaptureRule("example", {
    name: "Example Capture Rule",
    order: 1,
    description: "Managed by Pulumi",
    state: "ENABLED",
    action: "CAPTURE",
    srcIps: ["192.168.1.0/24"],
    destAddresses: ["10.0.0.0/8"],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.TrafficCaptureRule("example",
    name="Example Capture Rule",
    order=1,
    description="Managed by Pulumi",
    state="ENABLED",
    action="CAPTURE",
    src_ips=["192.168.1.0/24"],
    dest_addresses=["10.0.0.0/8"],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:TrafficCaptureRule
    properties:
      name: Example Capture Rule
      order: 1
      description: Managed by Pulumi
      state: ENABLED
      action: CAPTURE
      srcIps:
        - 192.168.1.0/24
      destAddresses:
        - 10.0.0.0/8
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing traffic capture rule can be imported using its ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:TrafficCaptureRule example 12345
`+tripleBacktick("")+`
`)
}

func (a *TrafficCaptureRuleArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "Name of the traffic capture rule.")
	ann.Describe(&a.Order, "The rule order of execution for the traffic capture rule.")
	ann.Describe(&a.Description, "Description of the traffic capture rule.")
	ann.Describe(&a.Rank, "The admin rank of the rule. Default is 7.")
	ann.Describe(&a.Action, "The action taken when traffic matches the rule (e.g., 'CAPTURE'). Default: 'CAPTURE'.")
	ann.Describe(&a.State, "The rule state. Accepted values: 'ENABLED' or 'DISABLED'. Default: 'ENABLED'.")
	ann.Describe(&a.SrcIps, "List of source IP addresses or CIDR ranges.")
	ann.Describe(&a.DestAddresses, "List of destination addresses.")
	ann.Describe(&a.DestIpCategories, "List of destination IP categories.")
	ann.Describe(&a.NwApplications, "List of network applications.")
	ann.Describe(&a.DefaultRule, "Whether this is a default rule.")
	ann.Describe(&a.Predefined, "Whether this is a predefined rule.")
	ann.Describe(&a.TxnSizeLimit, "Transaction size limit. Default: 'NONE'.")
	ann.Describe(&a.TxnSampling, "Transaction sampling mode. Default: 'NONE'.")
	ann.Describe(&a.Locations, "List of location IDs.")
	ann.Describe(&a.LocationGroups, "List of location group IDs.")
	ann.Describe(&a.Departments, "List of department IDs.")
	ann.Describe(&a.Groups, "List of group IDs.")
	ann.Describe(&a.Users, "List of user IDs.")
	ann.Describe(&a.TimeWindows, "List of time window IDs.")
	ann.Describe(&a.SrcIpGroups, "List of source IP group IDs.")
	ann.Describe(&a.DestIpGroups, "List of destination IP group IDs.")
	ann.Describe(&a.NwServices, "List of network service IDs.")
	ann.Describe(&a.NwServiceGroups, "List of network service group IDs.")
	ann.Describe(&a.NwApplicationGroups, "List of network application group IDs.")
	ann.Describe(&a.AppServiceGroups, "List of application service group IDs.")
	ann.Describe(&a.Labels, "List of label IDs.")
	ann.Describe(&a.DeviceGroups, "List of device group IDs.")
	ann.Describe(&a.Devices, "List of device IDs.")
	ann.Describe(&a.DestCountries, "List of destination country codes.")
	ann.Describe(&a.SourceCountries, "List of source country codes.")
	ann.Describe(&a.DeviceTrustLevels, "List of device trust levels.")
	ann.Describe(&a.WorkloadGroups, "List of workload groups.")
}

func (s *TrafficCaptureRuleState) Annotate(a infer.Annotator) {
	a.Describe(&s.RuleID, "The unique identifier for the traffic capture rule assigned by the ZIA cloud.")
}

var _ infer.CustomResource[TrafficCaptureRuleArgs, TrafficCaptureRuleState] = TrafficCaptureRule{}
