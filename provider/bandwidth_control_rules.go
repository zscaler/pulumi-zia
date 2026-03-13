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
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/bandwidth_control/bandwidth_control_rules"
)

const bandwidthControlResourceType = "bandwidth_control_rule"

type BandwidthControlRule struct{}

type BandwidthControlRuleArgs struct {
	Name             string   `pulumi:"name"`
	Order            int      `pulumi:"order"`
	Description      *string  `pulumi:"description,optional"`
	State            *string  `pulumi:"state,optional"`
	Rank             *int     `pulumi:"rank,optional"`
	MinBandwidth     *int     `pulumi:"minBandwidth,optional"`
	MaxBandwidth     *int     `pulumi:"maxBandwidth,optional"`
	BandwidthClasses []int    `pulumi:"bandwidthClasses,optional"`
	Locations        []int    `pulumi:"locations,optional"`
	LocationGroups   []int    `pulumi:"locationGroups,optional"`
	Labels           []int    `pulumi:"labels,optional"`
	TimeWindows      []int    `pulumi:"timeWindows,optional"`
	Protocols        []string `pulumi:"protocols,optional"`
}

type BandwidthControlRuleState struct {
	BandwidthControlRuleArgs
	RuleID *int `pulumi:"ruleId"`
}

func filterOutBandwidthDefaultRule(list []bandwidth_control_rules.BandwidthControlRules) []bandwidth_control_rules.BandwidthControlRules {
	var out []bandwidth_control_rules.BandwidthControlRules
	for _, r := range list {
		if r.Order != 125 && r.Name != "Default Bandwidth Control" {
			out = append(out, r)
		}
	}
	return out
}

func bandwidthControlRuleArgsToAPI(args *BandwidthControlRuleArgs, id int) bandwidth_control_rules.BandwidthControlRules {
	order := args.Order
	if order == 0 {
		order = 1
	}
	minBw, maxBw := 0, 0
	if args.MinBandwidth != nil {
		minBw = *args.MinBandwidth
	}
	if args.MaxBandwidth != nil {
		maxBw = *args.MaxBandwidth
	}
	return bandwidth_control_rules.BandwidthControlRules{
		ID:               id,
		Name:             args.Name,
		Description:      ptrToString(args.Description),
		State:            ptrToString(args.State),
		Order:            order,
		Rank:             ptrToIntDefault(args.Rank, 7),
		MinBandwidth:     minBw,
		MaxBandwidth:     maxBw,
		Protocols:        args.Protocols,
		BandwidthClasses: idsToIDNameExtensions(args.BandwidthClasses),
		Locations:        idsToIDNameExtensions(args.Locations),
		LocationGroups:   idsToIDNameExtensions(args.LocationGroups),
		Labels:           idsToIDNameExtensions(args.Labels),
		TimeWindows:      idsToIDNameExtensions(args.TimeWindows),
	}
}

func bandwidthControlRuleAPIToState(api *bandwidth_control_rules.BandwidthControlRules) BandwidthControlRuleState {
	return BandwidthControlRuleState{
		BandwidthControlRuleArgs: BandwidthControlRuleArgs{
			Name:             api.Name,
			Order:            api.Order,
			Description:      stringPtr(api.Description),
			State:            stringPtr(api.State),
			Rank:             intPtr(api.Rank),
			MinBandwidth:     intPtr(api.MinBandwidth),
			MaxBandwidth:     intPtr(api.MaxBandwidth),
			Protocols:        api.Protocols,
			BandwidthClasses: idNameExtensionsToIDs(api.BandwidthClasses),
			Locations:        idNameExtensionsToIDs(api.Locations),
			LocationGroups:   idNameExtensionsToIDs(api.LocationGroups),
			Labels:           idNameExtensionsToIDs(api.Labels),
			TimeWindows:      idNameExtensionsToIDs(api.TimeWindows),
		},
		RuleID: intPtr(api.ID),
	}
}

func (BandwidthControlRule) Create(ctx context.Context, req infer.CreateRequest[BandwidthControlRuleArgs]) (infer.CreateResponse[BandwidthControlRuleState], error) {
	if req.DryRun {
		return infer.CreateResponse[BandwidthControlRuleState]{ID: "preview", Output: BandwidthControlRuleState{BandwidthControlRuleArgs: req.Inputs, RuleID: intPtr(0)}}, nil
	}
	if req.Inputs.Order < 1 {
		return infer.CreateResponse[BandwidthControlRuleState]{}, fmt.Errorf("order must be a positive whole number (>= 1), got %d", req.Inputs.Order)
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[BandwidthControlRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := bandwidthControlRuleArgsToAPI(&req.Inputs, 0)
	timeout := 60 * time.Minute
	start := time.Now()

	intendedOrder := apiReq.Order

	for {
		select {
		case bandwidthControlSem <- struct{}{}:
		case <-ctx.Done():
			return infer.CreateResponse[BandwidthControlRuleState]{}, fmt.Errorf("creating bandwidth control rule: %w", ctx.Err())
		}

		bandwidthControlOrderMu.Lock()
		if bandwidthControlStartingOrder == 0 {
			list, _ := bandwidth_control_rules.GetAll(ctx, svc)
			for _, r := range list {
				if r.Order == 125 || r.Name == "Default Bandwidth Control" {
					continue
				}
				if r.Order > bandwidthControlStartingOrder {
					bandwidthControlStartingOrder = r.Order
				}
			}
			if bandwidthControlStartingOrder == 0 {
				bandwidthControlStartingOrder = 1
			}
		}
		apiReq.Order = bandwidthControlStartingOrder
		bandwidthControlOrderMu.Unlock()

		resp, err := bandwidth_control_rules.Create(ctx, svc, &apiReq)

		if err == nil {
			bandwidthControlOrderMu.Lock()
			bandwidthControlStartingOrder++
			bandwidthControlOrderMu.Unlock()
		}

		<-bandwidthControlSem

		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.CreateResponse[BandwidthControlRuleState]{}, customErr
		}
		if err != nil {
			reg := regexp.MustCompile("Rule with rank [0-9]+ is not allowed at order [0-9]+")
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") && reg.MatchString(err.Error()) {
				return infer.CreateResponse[BandwidthControlRuleState]{}, fmt.Errorf("error creating bandwidth control rule: %s, check order %d vs rank %d, err:%s", apiReq.Name, intendedOrder, apiReq.Rank, err)
			}
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Creating bandwidth control rule name: %v, got INVALID_INPUT_ARGUMENT\n", apiReq.Name)
				bandwidthControlOrderMu.Lock()
				bandwidthControlStartingOrder = 0
				bandwidthControlOrderMu.Unlock()
				if time.Since(start) < timeout {
					select {
					case <-time.After(10 * time.Second):
					case <-ctx.Done():
						return infer.CreateResponse[BandwidthControlRuleState]{}, fmt.Errorf("creating bandwidth control rule: %w", ctx.Err())
					}
					continue
				}
			}
			return infer.CreateResponse[BandwidthControlRuleState]{}, fmt.Errorf("creating bandwidth control rule: %w", err)
		}

		reorderWithBeforeReorder(
			OrderRule{Order: intendedOrder, Rank: apiReq.Rank},
			resp.ID,
			bandwidthControlResourceType,
			func() (int, error) {
				list, err := bandwidth_control_rules.GetAll(ctx, svc)
				if err != nil {
					return 0, err
				}
				return len(filterOutBandwidthDefaultRule(list)), nil
			},
			func(id int, order OrderRule) error {
				rule, err := bandwidth_control_rules.Get(ctx, svc, id)
				if err != nil {
					return err
				}
				if rule.Order == order.Order {
					return nil
				}
				// Strip read-only fields that cause "Request body is invalid" for predefined rules
				rule.DefaultRule = false
				rule.AccessControl = ""
				rule.Order = order.Order
				_, err = bandwidth_control_rules.Update(ctx, svc, id, rule)
				return err
			},
			nil,
		)
		markOrderRuleAsDone(resp.ID, bandwidthControlResourceType)
		waitForReorder(bandwidthControlResourceType)

		if shouldActivate() {
			time.Sleep(2 * time.Second)
			if activationErr := triggerActivation(ctx, client); activationErr != nil {
				return infer.CreateResponse[BandwidthControlRuleState]{}, activationErr
			}
		} else {
			log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
		}

		rule, err := bandwidth_control_rules.Get(ctx, svc, resp.ID)
		if err != nil {
			return infer.CreateResponse[BandwidthControlRuleState]{ID: strconv.Itoa(resp.ID), Output: BandwidthControlRuleState{BandwidthControlRuleArgs: req.Inputs, RuleID: intPtr(resp.ID)}}, nil
		}
		return infer.CreateResponse[BandwidthControlRuleState]{ID: strconv.Itoa(resp.ID), Output: bandwidthControlRuleAPIToState(rule)}, nil
	}
}

func (BandwidthControlRule) Read(ctx context.Context, req infer.ReadRequest[BandwidthControlRuleArgs, BandwidthControlRuleState]) (infer.ReadResponse[BandwidthControlRuleArgs, BandwidthControlRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[BandwidthControlRuleArgs, BandwidthControlRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		rule, lookupErr := bandwidth_control_rules.GetByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[BandwidthControlRuleArgs, BandwidthControlRuleState]{}, nil
			}
			return infer.ReadResponse[BandwidthControlRuleArgs, BandwidthControlRuleState]{}, lookupErr
		}
		id = rule.ID
	}

	rule, err := bandwidth_control_rules.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[BandwidthControlRuleArgs, BandwidthControlRuleState]{}, nil
		}
		return infer.ReadResponse[BandwidthControlRuleArgs, BandwidthControlRuleState]{}, err
	}

	state := bandwidthControlRuleAPIToState(rule)
	return infer.ReadResponse[BandwidthControlRuleArgs, BandwidthControlRuleState]{ID: req.ID, Inputs: state.BandwidthControlRuleArgs, State: state}, nil
}

func (BandwidthControlRule) Update(ctx context.Context, req infer.UpdateRequest[BandwidthControlRuleArgs, BandwidthControlRuleState]) (infer.UpdateResponse[BandwidthControlRuleState], error) {
	if req.Inputs.Order < 1 {
		return infer.UpdateResponse[BandwidthControlRuleState]{}, fmt.Errorf("order must be a positive whole number (>= 1), got %d", req.Inputs.Order)
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[BandwidthControlRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[BandwidthControlRuleState]{}, fmt.Errorf("invalid bandwidth control rule ID: %s", req.ID)
	}
	apiReq := bandwidthControlRuleArgsToAPI(&req.Inputs, id)
	timeout := 60 * time.Minute
	start := time.Now()

	for {
		_, err = bandwidth_control_rules.Update(ctx, svc, id, &apiReq)
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.UpdateResponse[BandwidthControlRuleState]{}, nil
		}
		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.UpdateResponse[BandwidthControlRuleState]{}, customErr
		}
		if err != nil {
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Updating bandwidth control rule ID: %v, got INVALID_INPUT_ARGUMENT\n", id)
				if time.Since(start) < timeout {
					time.Sleep(10 * time.Second)
					continue
				}
			}
			return infer.UpdateResponse[BandwidthControlRuleState]{}, err
		}

		reorderWithBeforeReorder(
			OrderRule{Order: apiReq.Order, Rank: apiReq.Rank},
			id,
			bandwidthControlResourceType,
			func() (int, error) {
				list, err := bandwidth_control_rules.GetAll(ctx, svc)
				if err != nil {
					return 0, err
				}
				return len(filterOutBandwidthDefaultRule(list)), nil
			},
			func(ruleID int, order OrderRule) error {
				rule, err := bandwidth_control_rules.Get(ctx, svc, ruleID)
				if err != nil {
					return err
				}
				if rule.Order == order.Order {
					return nil
				}
				// Strip read-only fields that cause "Request body is invalid" for predefined rules
				rule.DefaultRule = false
				rule.AccessControl = ""
				rule.Order = order.Order
				_, err = bandwidth_control_rules.Update(ctx, svc, ruleID, rule)
				return err
			},
			nil,
		)
		markOrderRuleAsDone(id, bandwidthControlResourceType)
		waitForReorder(bandwidthControlResourceType)
		break
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[BandwidthControlRuleState]{}, activationErr
		}
	}

	updated, err := bandwidth_control_rules.Get(ctx, svc, id)
	if err != nil {
		return infer.UpdateResponse[BandwidthControlRuleState]{Output: BandwidthControlRuleState{BandwidthControlRuleArgs: req.Inputs, RuleID: intPtr(id)}}, nil
	}
	return infer.UpdateResponse[BandwidthControlRuleState]{Output: bandwidthControlRuleAPIToState(updated)}, nil
}

func (BandwidthControlRule) Delete(ctx context.Context, req infer.DeleteRequest[BandwidthControlRuleState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid bandwidth control rule ID: %s", req.ID)
	}
	if _, err := bandwidth_control_rules.Delete(ctx, svc, id); err != nil {
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

func (BandwidthControlRule) Annotate(a infer.Annotator) {
	describeResource(a, &BandwidthControlRule{}, `The zia_bandwidth_control_rule resource manages bandwidth control rules in the Zscaler Internet Access (ZIA) cloud service. Bandwidth control rules allow administrators to define minimum and maximum bandwidth limits for specific traffic, locations, and time windows to ensure quality of service across the network.

For more information, see the [ZIA Bandwidth Control documentation](https://help.zscaler.com/zia/bandwidth-control).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Bandwidth Control Rule

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.BandwidthControlRule("example", {
    name: "Example Bandwidth Control Rule",
    description: "Limit streaming bandwidth",
    order: 1,
    state: "ENABLED",
    maxBandwidth: 50,
    protocols: ["ANY_RULE"],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.BandwidthControlRule("example",
    name="Example Bandwidth Control Rule",
    description="Limit streaming bandwidth",
    order=1,
    state="ENABLED",
    max_bandwidth=50,
    protocols=["ANY_RULE"],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:BandwidthControlRule
    properties:
      name: Example Bandwidth Control Rule
      description: Limit streaming bandwidth
      order: 1
      state: ENABLED
      maxBandwidth: 50
      protocols:
        - ANY_RULE
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing Bandwidth Control Rule can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:BandwidthControlRule example 12345
`+tripleBacktick("")+`
`)
}

func (a *BandwidthControlRuleArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the bandwidth control rule. Must be unique.")
	ann.Describe(&a.Order, "The order of execution of the rule with respect to other bandwidth control rules.")
	ann.Describe(&a.Description, "Additional information about the bandwidth control rule.")
	ann.Describe(&a.State, "Rule state. Valid values: `ENABLED`, `DISABLED`.")
	ann.Describe(&a.Rank, "Admin rank of the bandwidth control rule. Valid values: 0-7. Default: 7.")
	ann.Describe(&a.MinBandwidth, "The minimum bandwidth percentage allocated. Valid range: 0-100.")
	ann.Describe(&a.MaxBandwidth, "The maximum bandwidth percentage allowed. Valid range: 0-100.")
	ann.Describe(&a.BandwidthClasses, "IDs of bandwidth classes associated with this rule.")
	ann.Describe(&a.Locations, "IDs of locations for which the rule must be applied.")
	ann.Describe(&a.LocationGroups, "IDs of location groups for which the rule must be applied.")
	ann.Describe(&a.Labels, "IDs of labels associated with the bandwidth control rule.")
	ann.Describe(&a.TimeWindows, "IDs of time intervals during which the rule must be enforced.")
	ann.Describe(&a.Protocols, "Protocols to which the rule applies. Valid values: `ANY_RULE`, `TCP_RULE`, `UDP_RULE`, `SSL_RULE`.")
}

func (s *BandwidthControlRuleState) Annotate(a infer.Annotator) {
	a.Describe(&s.RuleID, "The system-generated ID of the bandwidth control rule.")
}

var _ infer.CustomResource[BandwidthControlRuleArgs, BandwidthControlRuleState] = BandwidthControlRule{}
