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

// Package provider implements the URL Filtering Rules resource.
// Adopted from terraform-provider-zia resource_zia_url_filtering_rules.go.

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

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/pulumi-zia/provider/internal/zia"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/urlfilteringpolicies"
)

// URLFilteringRule implements the zia:index:URLFilteringRule resource.
type URLFilteringRule struct{}

// URLFilteringRuleArgs are the inputs for URLFilteringRule.
type URLFilteringRuleArgs struct {
	Name                   string               `pulumi:"name"`
	Description            *string              `pulumi:"description,optional"`
	Order                  int                  `pulumi:"order"`
	State                  *string              `pulumi:"state,optional"`
	Rank                   *int                 `pulumi:"rank,optional"`
	EndUserNotificationURL *string              `pulumi:"endUserNotificationUrl,optional"`
	BlockOverride          *bool                `pulumi:"blockOverride,optional"`
	TimeQuota              *int                 `pulumi:"timeQuota,optional"`
	SizeQuota              *int                 `pulumi:"sizeQuota,optional"`
	EnforceTimeValidity    *bool                `pulumi:"enforceTimeValidity,optional"`
	ValidityStartTime      *string              `pulumi:"validityStartTime,optional"`
	ValidityEndTime        *string              `pulumi:"validityEndTime,optional"`
	ValidityTimeZoneID     *string              `pulumi:"validityTimeZoneId,optional"`
	Action                 *string              `pulumi:"action,optional"`
	Ciparule               *bool                `pulumi:"ciparule,optional"`
	BrowserEunTemplateID   *int                 `pulumi:"browserEunTemplateId,optional"`
	CBIProfile             *CBIProfileInput     `pulumi:"cbiProfile,optional"`
	URLCategories          []string             `pulumi:"urlCategories,optional"`
	Protocols              []string             `pulumi:"protocols,optional"`
	Locations              []int                `pulumi:"locations,optional"`
	Groups                 []int                `pulumi:"groups,optional"`
	Departments            []int                `pulumi:"departments,optional"`
	Users                  []int                `pulumi:"users,optional"`
	TimeWindows            []int                `pulumi:"timeWindows,optional"`
	OverrideUsers          []int                `pulumi:"overrideUsers,optional"`
	OverrideGroups         []int                `pulumi:"overrideGroups,optional"`
	DeviceGroups           []int                `pulumi:"deviceGroups,optional"`
	Devices                []int                `pulumi:"devices,optional"`
	LocationGroups         []int                `pulumi:"locationGroups,optional"`
	Labels                 []int                `pulumi:"labels,optional"`
	SourceIPGroups         []int                `pulumi:"sourceIpGroups,optional"`
	SourceCountries        []string             `pulumi:"sourceCountries,optional"`
	DeviceTrustLevels      []string             `pulumi:"deviceTrustLevels,optional"`
	UserRiskScoreLevels    []string             `pulumi:"userRiskScoreLevels,optional"`
	RequestMethods         []string             `pulumi:"requestMethods,optional"`
	UserAgentTypes         []string             `pulumi:"userAgentTypes,optional"`
	WorkloadGroups         []WorkloadGroupInput `pulumi:"workloadGroups,optional"`
}

func (URLFilteringRule) Annotate(a infer.Annotator) {
	describeResource(a, &URLFilteringRule{}, `The zia_url_filtering_rules resource manages URL filtering rules in the Zscaler Internet Access (ZIA) cloud service. URL filtering rules define the actions to take when users access URLs that match specific categories, protocols, locations, departments, groups, or users.

For more information, see the [ZIA URL Filtering documentation](https://help.zscaler.com/zia/url-filtering).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic URL Filtering Rule

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.URLFilteringRule("example", {
    name: "Example URL Filtering Rule",
    description: "Allow access to business URLs",
    order: 1,
    state: "ENABLED",
    action: "ALLOW",
    protocols: ["ANY_RULE"],
    urlCategories: ["ANY"],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.URLFilteringRule("example",
    name="Example URL Filtering Rule",
    description="Allow access to business URLs",
    order=1,
    state="ENABLED",
    action="ALLOW",
    protocols=["ANY_RULE"],
    url_categories=["ANY"],
)
`+tripleBacktick("")+`

`+tripleBacktick("go")+`
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zia "github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := zia.NewURLFilteringRule(ctx, "example", &zia.URLFilteringRuleArgs{
			Name:          pulumi.String("Example URL Filtering Rule"),
			Description:   pulumi.StringRef("Allow access to business URLs"),
			Order:         pulumi.Int(1),
			State:         pulumi.StringRef("ENABLED"),
			Action:        pulumi.StringRef("ALLOW"),
			Protocols:     pulumi.ToStringArray([]string{"ANY_RULE"}),
			UrlCategories: pulumi.ToStringArray([]string{"ANY"}),
		})
		return err
	})
}
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:URLFilteringRule
    properties:
      name: Example URL Filtering Rule
      description: Allow access to business URLs
      order: 1
      state: ENABLED
      action: ALLOW
      protocols:
        - ANY_RULE
      urlCategories:
        - ANY
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing URL Filtering Rule can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:URLFilteringRule example 12345
`+tripleBacktick("")+`
`)
}

func (a *URLFilteringRuleArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the URL filtering rule. Must be unique.")
	ann.Describe(&a.Description, "Additional information about the URL filtering rule. Maximum 10240 characters.")
	ann.Describe(&a.Order, "The order of execution of the rule with respect to other URL filtering rules.")
	ann.Describe(&a.State, "Rule state. Valid values: `ENABLED`, `DISABLED`.")
	ann.Describe(&a.Rank, "Admin rank of the URL filtering policy rule. Valid values: 0-7. Default: 7.")
	ann.Describe(&a.EndUserNotificationURL, "URL of end user notification page to be displayed when the rule is matched. Not applicable if either 'overrideUsers' or 'overrideGroups' is specified.")
	ann.Describe(&a.BlockOverride, "When set to true, a 'BLOCK' action can be overridden. Can only be set when action is 'BLOCK'.")
	ann.Describe(&a.TimeQuota, "Time quota in minutes, after which the URL filtering rule is applied. If not set, no quota is enforced. Valid range: 15-600. Not applicable when action is 'BLOCK'.")
	ann.Describe(&a.SizeQuota, "Size quota in MB beyond which the URL filtering rule is applied. If not set, no quota is enforced. Valid range: 10-100000. Not applicable when action is 'BLOCK'.")
	ann.Describe(&a.EnforceTimeValidity, "Enforce a set validity time period for the URL filtering rule.")
	ann.Describe(&a.ValidityStartTime, "If enforceTimeValidity is set to true, the URL filtering rule is valid starting on this date and time (RFC 1123 format).")
	ann.Describe(&a.ValidityEndTime, "If enforceTimeValidity is set to true, the URL filtering rule ceases to be valid on this end date and time (RFC 1123 format).")
	ann.Describe(&a.ValidityTimeZoneID, "If enforceTimeValidity is set to true, the URL filtering rule date and time is valid based on this time zone ID. Use IANA format (e.g. 'America/Los_Angeles'). See https://nodatime.org/TimeZones for the complete list.")
	ann.Describe(&a.Action, "Action taken when traffic matches rule criteria. Valid values: `BLOCK`, `CAUTION`, `ALLOW`, `ISOLATE`.")
	ann.Describe(&a.Ciparule, "If set to true, the CIPA Compliance rule is enabled.")
	ann.Describe(&a.BrowserEunTemplateID, "Browser End User Notification template ID. Only applicable when action is 'BLOCK' or 'CAUTION'.")
	ann.Describe(&a.CBIProfile, "The Cloud Browser Isolation (CBI) profile. Required when action is 'ISOLATE'.")
	ann.Describe(&a.URLCategories, "List of URL categories to which the rule applies. See the [URL Categories API](https://help.zscaler.com/zia/url-categories#/urlCategories-get) for available categories.")
	ann.Describe(&a.Protocols, "Protocols to which the rule applies. Valid values: `SMRULEF_ZPA_BROKERS_RULE`, `ANY_RULE`, `TCP_RULE`, `UDP_RULE`, `DOHTTPS_RULE`, `TUNNELSSL_RULE`, `HTTP_PROXY`, `FOHTTP_RULE`, `FTP_RULE`, `SSL_RULE`.")
	ann.Describe(&a.Locations, "IDs of locations for which the rule must be applied.")
	ann.Describe(&a.Groups, "IDs of groups for which the rule must be applied.")
	ann.Describe(&a.Departments, "IDs of departments for which the rule must be applied.")
	ann.Describe(&a.Users, "IDs of users for which the rule must be applied.")
	ann.Describe(&a.TimeWindows, "IDs of time intervals during which the rule must be enforced.")
	ann.Describe(&a.OverrideUsers, "IDs of users for which this rule can be overridden. Only applicable when action is 'BLOCK' and blockOverride is true.")
	ann.Describe(&a.OverrideGroups, "IDs of groups for which this rule can be overridden. Only applicable when action is 'BLOCK' and blockOverride is true.")
	ann.Describe(&a.DeviceGroups, "IDs of device groups for which the rule must be applied. Applicable for devices managed using Zscaler Client Connector.")
	ann.Describe(&a.Devices, "IDs of devices for which the rule must be applied.")
	ann.Describe(&a.LocationGroups, "IDs of location groups to which the rule must be applied.")
	ann.Describe(&a.Labels, "IDs of labels associated with the URL filtering rule.")
	ann.Describe(&a.SourceIPGroups, "IDs of source IP address groups.")
	ann.Describe(&a.SourceCountries, "Source countries (ISO 3166-1 alpha-2 codes) for the rule.")
	ann.Describe(&a.DeviceTrustLevels, "Device trust levels for the rule. Valid values: `ANY`, `UNKNOWN_DEVICETRUSTLEVEL`, `LOW_TRUST`, `MEDIUM_TRUST`, `HIGH_TRUST`.")
	ann.Describe(&a.UserRiskScoreLevels, "User risk score levels for the rule. Valid values: `LOW`, `MEDIUM`, `HIGH`, `CRITICAL`.")
	ann.Describe(&a.RequestMethods, "Request methods to which the rule applies. Valid values: `CONNECT`, `DELETE`, `GET`, `HEAD`, `OPTIONS`, `OTHER`, `POST`, `PUT`, `TRACE`.")
	ann.Describe(&a.UserAgentTypes, "User agent types the rule applies to. Valid values: `CHROME`, `FIREFOX`, `MSIE`, `MSEDGE`, `MSCHREDGE`, `OPERA`, `SAFARI`, `OTHER`.")
	ann.Describe(&a.WorkloadGroups, "List of preconfigured workload groups to which the policy must be applied.")
}

func (s *URLFilteringRuleState) Annotate(a infer.Annotator) {
	a.Describe(&s.RuleID, "The system-generated ID of the URL filtering rule.")
}

// URLFilteringRuleState is the persisted state.
type URLFilteringRuleState struct {
	URLFilteringRuleArgs
	RuleID *int `pulumi:"ruleId"`
}

// Check validates and normalizes inputs.
func (URLFilteringRule) Check(ctx context.Context, req infer.CheckRequest) (infer.CheckResponse[URLFilteringRuleArgs], error) {
	inputs, failures, err := infer.DefaultCheck[URLFilteringRuleArgs](ctx, req.NewInputs)
	if err != nil {
		return infer.CheckResponse[URLFilteringRuleArgs]{}, err
	}
	if len(failures) > 0 {
		return infer.CheckResponse[URLFilteringRuleArgs]{Failures: failures}, nil
	}
	// Description max 10240
	if inputs.Description != nil && len(*inputs.Description) > 10240 {
		return infer.CheckResponse[URLFilteringRuleArgs]{Failures: []p.CheckFailure{{
			Property: "description",
			Reason:   "description must be at most 10240 characters",
		}}}, nil
	}
	// Rank 0-7
	if inputs.Rank != nil && (*inputs.Rank < 0 || *inputs.Rank > 7) {
		return infer.CheckResponse[URLFilteringRuleArgs]{Failures: []p.CheckFailure{{
			Property: "rank",
			Reason:   "rank must be between 0 and 7",
		}}}, nil
	}
	// TimeQuota 15-600
	if inputs.TimeQuota != nil && (*inputs.TimeQuota < 15 || *inputs.TimeQuota > 600) {
		return infer.CheckResponse[URLFilteringRuleArgs]{Failures: []p.CheckFailure{{
			Property: "timeQuota",
			Reason:   "timeQuota must be between 15 and 600",
		}}}, nil
	}
	// SizeQuota 10-100000 (MB)
	if inputs.SizeQuota != nil && (*inputs.SizeQuota < 10 || *inputs.SizeQuota > 100000) {
		return infer.CheckResponse[URLFilteringRuleArgs]{Failures: []p.CheckFailure{{
			Property: "sizeQuota",
			Reason:   "sizeQuota must be between 10 and 100000 MB",
		}}}, nil
	}
	// State validation
	if inputs.State != nil {
		s := strings.ToUpper(*inputs.State)
		if s != "ENABLED" && s != "DISABLED" {
			return infer.CheckResponse[URLFilteringRuleArgs]{Failures: []p.CheckFailure{{
				Property: "state",
				Reason:   "state must be ENABLED or DISABLED",
			}}}, nil
		}
	}
	// Action validation
	if inputs.Action != nil {
		act := strings.ToUpper(*inputs.Action)
		if act != "BLOCK" && act != "CAUTION" && act != "ALLOW" && act != "ISOLATE" {
			return infer.CheckResponse[URLFilteringRuleArgs]{Failures: []p.CheckFailure{{
				Property: "action",
				Reason:   "action must be BLOCK, CAUTION, ALLOW, or ISOLATE",
			}}}, nil
		}
	}
	// Order must be >= 1
	if inputs.Order < 1 {
		return infer.CheckResponse[URLFilteringRuleArgs]{Failures: []p.CheckFailure{{
			Property: "order",
			Reason:   "order must be a positive whole number (>= 1)",
		}}}, nil
	}
	// Normalize validity times: parse and re-format so the day-of-week is
	// always correct (prevents drift when the user supplies a wrong weekday).
	if inputs.ValidityStartTime != nil && *inputs.ValidityStartTime != "" {
		if epoch, parseErr := ConvertRFC1123ToEpoch(*inputs.ValidityStartTime); parseErr == nil {
			normalized := time.Unix(int64(epoch), 0).UTC().Format(time.RFC1123)
			inputs.ValidityStartTime = &normalized
		}
	}
	if inputs.ValidityEndTime != nil && *inputs.ValidityEndTime != "" {
		if epoch, parseErr := ConvertRFC1123ToEpoch(*inputs.ValidityEndTime); parseErr == nil {
			normalized := time.Unix(int64(epoch), 0).UTC().Format(time.RFC1123)
			inputs.ValidityEndTime = &normalized
		}
	}

	// Custom validation (CustomizeDiff equivalent)
	err = validateURLFilteringCheck(
		inputs.Action, inputs.BlockOverride, inputs.OverrideUsers, inputs.OverrideGroups,
		inputs.CBIProfile, inputs.UserAgentTypes, inputs.Protocols, inputs.RequestMethods,
		inputs.EnforceTimeValidity, inputs.ValidityStartTime, inputs.ValidityEndTime, inputs.ValidityTimeZoneID,
		inputs.BrowserEunTemplateID,
	)
	if err != nil {
		return infer.CheckResponse[URLFilteringRuleArgs]{Failures: []p.CheckFailure{{
			Property: "",
			Reason:   err.Error(),
		}}}, nil
	}
	return infer.CheckResponse[URLFilteringRuleArgs]{Inputs: inputs}, nil
}

// Create creates a new URL Filtering Rule.
func (URLFilteringRule) Diff(ctx context.Context, req infer.DiffRequest[URLFilteringRuleArgs, URLFilteringRuleState]) (infer.DiffResponse, error) {
	diff, hasChanges := diffSkippingUnsetInputs(req.State.URLFilteringRuleArgs, req.Inputs)
	return infer.DiffResponse{HasChanges: hasChanges, DetailedDiff: diff}, nil
}

func (URLFilteringRule) Create(ctx context.Context, req infer.CreateRequest[URLFilteringRuleArgs]) (infer.CreateResponse[URLFilteringRuleState], error) {
	if req.DryRun {
		s := URLFilteringRuleState{URLFilteringRuleArgs: req.Inputs, RuleID: intPtr(0)}
		return infer.CreateResponse[URLFilteringRuleState]{ID: "preview", Output: s}, nil
	}

	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[URLFilteringRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := urlFilteringRuleArgsToAPI(req.Inputs, 0)

	if err := validateURLFilteringActions(apiReq); err != nil {
		return infer.CreateResponse[URLFilteringRuleState]{}, err
	}

	timeout := 60 * time.Minute
	start := time.Now()

	intendedOrder := apiReq.Order
	intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)

	for {
		select {
		case urlFilteringSem <- struct{}{}:
		case <-ctx.Done():
			return infer.CreateResponse[URLFilteringRuleState]{}, fmt.Errorf("creating url filtering rule: %w", ctx.Err())
		}

		urlFilteringOrderMu.Lock()
		if urlFilteringStartingOrder == 0 {
			list, _ := urlfilteringpolicies.GetAll(ctx, svc)
			for _, r := range list {
				if r.Order > urlFilteringStartingOrder {
					urlFilteringStartingOrder = r.Order
				}
			}
			if urlFilteringStartingOrder == 0 {
				urlFilteringStartingOrder = 1
			}
		}

		if intendedRank < 7 {
			apiReq.Rank = 7
		}
		apiReq.Order = urlFilteringStartingOrder
		urlFilteringOrderMu.Unlock()

		resp, err := urlfilteringpolicies.Create(ctx, svc, &apiReq)

		if err == nil {
			urlFilteringOrderMu.Lock()
			urlFilteringStartingOrder++
			urlFilteringOrderMu.Unlock()
		}

		<-urlFilteringSem

		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.CreateResponse[URLFilteringRuleState]{}, customErr
		}

		if err != nil {
			reg := regexp.MustCompile("Rule with rank [0-9]+ is not allowed at order [0-9]+")
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") && reg.MatchString(err.Error()) {
				return infer.CreateResponse[URLFilteringRuleState]{}, fmt.Errorf("error creating resource: %s, please check the order %d vs rank %d, current rules:%s, err:%s",
					apiReq.Name, intendedOrder, intendedRank, currentOrderVsRankWording(ctx, client), err)
			}
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Creating url filtering rule name: %v, got INVALID_INPUT_ARGUMENT\n", apiReq.Name)
				urlFilteringOrderMu.Lock()
				urlFilteringStartingOrder = 0
				urlFilteringOrderMu.Unlock()
				if time.Since(start) < timeout {
					select {
					case <-time.After(10 * time.Second):
					case <-ctx.Done():
						return infer.CreateResponse[URLFilteringRuleState]{}, fmt.Errorf("creating url filtering rule: %w", ctx.Err())
					}
					continue
				}
			}
			return infer.CreateResponse[URLFilteringRuleState]{}, fmt.Errorf("creating url filtering rule: %w", err)
		}

		resourceType := "url_filtering_rules"
		reorderWithBeforeReorder(
			OrderRule{Order: intendedOrder, Rank: intendedRank},
			resp.ID,
			resourceType,
			func() (int, error) {
				allRules, err := urlfilteringpolicies.GetAll(ctx, svc)
				if err != nil {
					return 0, err
				}
				return len(allRules), nil
			},
			func(id int, order OrderRule) error {
				rule, err := urlfilteringpolicies.Get(ctx, svc, id)
				if err != nil {
					return err
				}
				// to avoid the STALE_CONFIGURATION_ERROR
				rule.LastModifiedTime = 0
				rule.LastModifiedBy = nil
				rule.Order = order.Order
				rule.Rank = order.Rank
				_, err = urlfilteringpolicies.Update(ctx, svc, id, rule)
				return err
			},
			nil,
		)

		markOrderRuleAsDone(resp.ID, resourceType)
		waitForReorder(resourceType)

		if shouldActivate() {
			time.Sleep(2 * time.Second)
			if activationErr := triggerActivation(ctx, client); activationErr != nil {
				return infer.CreateResponse[URLFilteringRuleState]{}, activationErr
			}
		}

		// Read to populate state
		rule, err := urlfilteringpolicies.Get(ctx, svc, resp.ID)
		if err != nil {
			state := URLFilteringRuleState{
				URLFilteringRuleArgs: req.Inputs,
				RuleID:               &resp.ID,
			}
			return infer.CreateResponse[URLFilteringRuleState]{
				ID:     strconv.Itoa(resp.ID),
				Output: state,
			}, nil
		}

		state := urlFilteringRuleAPIToState(rule)
		return infer.CreateResponse[URLFilteringRuleState]{
			ID:     strconv.Itoa(resp.ID),
			Output: state,
		}, nil
	}
}

// Read fetches the URL Filtering Rule state.
func (URLFilteringRule) Read(ctx context.Context, req infer.ReadRequest[URLFilteringRuleArgs, URLFilteringRuleState]) (infer.ReadResponse[URLFilteringRuleArgs, URLFilteringRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[URLFilteringRuleArgs, URLFilteringRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		rule, lookupErr := urlfilteringpolicies.GetByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[URLFilteringRuleArgs, URLFilteringRuleState]{}, nil
			}
			return infer.ReadResponse[URLFilteringRuleArgs, URLFilteringRuleState]{}, lookupErr
		}
		id = rule.ID
	}

	rule, err := urlfilteringpolicies.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[URLFilteringRuleArgs, URLFilteringRuleState]{}, nil
		}
		return infer.ReadResponse[URLFilteringRuleArgs, URLFilteringRuleState]{}, err
	}

	state := urlFilteringRuleAPIToState(rule)
	args := urlFilteringRuleStateToArgs(rule)

	return infer.ReadResponse[URLFilteringRuleArgs, URLFilteringRuleState]{
		ID:     strconv.Itoa(rule.ID),
		Inputs: args,
		State:  state,
	}, nil
}

// Update updates an existing URL Filtering Rule.
func (URLFilteringRule) Update(ctx context.Context, req infer.UpdateRequest[URLFilteringRuleArgs, URLFilteringRuleState]) (infer.UpdateResponse[URLFilteringRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[URLFilteringRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[URLFilteringRuleState]{}, fmt.Errorf("invalid rule id: %w", err)
	}

	apiReq := urlFilteringRuleArgsToAPI(req.Inputs, id)
	if err := validateURLFilteringActions(apiReq); err != nil {
		return infer.UpdateResponse[URLFilteringRuleState]{}, err
	}

	existingRules, err := urlfilteringpolicies.GetAll(ctx, svc)
	if err != nil {
		log.Printf("[ERROR] error getting all url filtering rules: %v", err)
	}
	sort.Slice(existingRules, func(i, j int) bool {
		return existingRules[i].Rank < existingRules[j].Rank || (existingRules[i].Rank == existingRules[j].Rank && existingRules[i].Order < existingRules[j].Order)
	})
	intendedOrder := apiReq.Order
	intendedRank := apiReq.Rank
	nextAvailableOrder := existingRules[len(existingRules)-1].Order
	apiReq.Rank = 7
	apiReq.Order = nextAvailableOrder

	_, err = urlfilteringpolicies.Update(ctx, svc, id, &apiReq)
	if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
		return infer.UpdateResponse[URLFilteringRuleState]{}, nil
	}
	if customErr := failFastOnErrorCodes(err); customErr != nil {
		return infer.UpdateResponse[URLFilteringRuleState]{}, customErr
	}
	if err != nil {
		return infer.UpdateResponse[URLFilteringRuleState]{}, fmt.Errorf("updating url filtering rule: %w", err)
	}

	resourceType := "url_filtering_rules"
	reorderWithBeforeReorder(
		OrderRule{Order: intendedOrder, Rank: intendedRank},
		id,
		resourceType,
		func() (int, error) {
			allRules, err := urlfilteringpolicies.GetAll(ctx, svc)
			if err != nil {
				return 0, err
			}
			return len(allRules), nil
		},
		func(ruleID int, order OrderRule) error {
			rule, err := urlfilteringpolicies.Get(ctx, svc, ruleID)
			if err != nil {
				return err
			}
			// to avoid the STALE_CONFIGURATION_ERROR
			rule.LastModifiedTime = 0
			rule.LastModifiedBy = nil
			rule.Order = order.Order
			rule.Rank = order.Rank
			_, err = urlfilteringpolicies.Update(ctx, svc, ruleID, rule)
			return err
		},
		nil,
	)

	markOrderRuleAsDone(id, resourceType)
	waitForReorder(resourceType)

	time.Sleep(2 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[URLFilteringRuleState]{}, activationErr
		}
	}

	rule, err := urlfilteringpolicies.Get(ctx, svc, id)
	if err != nil {
		state := URLFilteringRuleState{URLFilteringRuleArgs: req.Inputs, RuleID: &id}
		return infer.UpdateResponse[URLFilteringRuleState]{Output: state}, nil
	}
	state := urlFilteringRuleAPIToState(rule)
	return infer.UpdateResponse[URLFilteringRuleState]{Output: state}, nil
}

// Delete removes the URL Filtering Rule.
func (URLFilteringRule) Delete(ctx context.Context, req infer.DeleteRequest[URLFilteringRuleState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid rule id: %w", err)
	}

	if _, err := urlfilteringpolicies.Delete(ctx, svc, id); err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.DeleteResponse{}, nil
		}
		return infer.DeleteResponse{}, fmt.Errorf("deleting url filtering rule: %w", err)
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}
	return infer.DeleteResponse{}, nil
}

func urlFilteringRuleArgsToAPI(in URLFilteringRuleArgs, existingID int) urlfilteringpolicies.URLFilteringRule {
	order := in.Order
	if order == 0 {
		order = 1
	}

	var validityStartTime, validityEndTime int
	if in.ValidityStartTime != nil && *in.ValidityStartTime != "" {
		t, err := ConvertRFC1123ToEpoch(*in.ValidityStartTime)
		if err == nil {
			validityStartTime = t
		}
	}
	if in.ValidityEndTime != nil && *in.ValidityEndTime != "" {
		t, err := ConvertRFC1123ToEpoch(*in.ValidityEndTime)
		if err == nil {
			validityEndTime = t
		}
	}

	sizeQuotaKB := 0
	if in.SizeQuota != nil && *in.SizeQuota > 0 {
		if k, err := convertAndValidateSizeQuota(*in.SizeQuota); err == nil {
			sizeQuotaKB = k
		}
	}

	urlCategories := in.URLCategories
	if len(urlCategories) == 0 {
		urlCategories = []string{"ANY"}
	}

	return urlfilteringpolicies.URLFilteringRule{
		ID:                     existingID,
		Name:                   in.Name,
		Description:            ptrToString(in.Description),
		Order:                  order,
		Protocols:              in.Protocols,
		URLCategories:          urlCategories,
		UserRiskScoreLevels:    in.UserRiskScoreLevels,
		DeviceTrustLevels:      in.DeviceTrustLevels,
		RequestMethods:         in.RequestMethods,
		UserAgentTypes:         in.UserAgentTypes,
		State:                  ptrToString(in.State),
		Rank:                   ptrToIntDefault(in.Rank, 7),
		EndUserNotificationURL: ptrToString(in.EndUserNotificationURL),
		BlockOverride:          ptrToBool(in.BlockOverride),
		TimeQuota:              ptrToIntDefault(in.TimeQuota, 0),
		BrowserEunTemplateID:   ptrToIntDefault(in.BrowserEunTemplateID, 0),
		SizeQuota:              sizeQuotaKB,
		ValidityStartTime:      validityStartTime,
		ValidityEndTime:        validityEndTime,
		SourceCountries:        processCountries(in.SourceCountries),
		ValidityTimeZoneID:     ptrToString(in.ValidityTimeZoneID),
		EnforceTimeValidity:    ptrToBool(in.EnforceTimeValidity),
		Action:                 ptrToString(in.Action),
		Ciparule:               ptrToBool(in.Ciparule),
		Locations:              idsToIDNameExtensions(in.Locations),
		Groups:                 idsToIDNameExtensions(in.Groups),
		Departments:            idsToIDNameExtensions(in.Departments),
		Users:                  idsToIDNameExtensions(in.Users),
		TimeWindows:            idsToIDNameExtensions(in.TimeWindows),
		OverrideUsers:          idsToIDNameExtensions(in.OverrideUsers),
		OverrideGroups:         idsToIDNameExtensions(in.OverrideGroups),
		LocationGroups:         idsToIDNameExtensions(in.LocationGroups),
		Labels:                 idsToIDNameExtensions(in.Labels),
		DeviceGroups:           idsToIDNameExtensions(in.DeviceGroups),
		Devices:                idsToIDNameExtensions(in.Devices),
		SourceIPGroups:         idsToIDNameExtensions(in.SourceIPGroups),
		WorkloadGroups:         expandWorkloadGroups(in.WorkloadGroups),
		CBIProfile:             expandCBIProfile(in.CBIProfile),
	}
}

func urlFilteringRuleAPIToState(rule *urlfilteringpolicies.URLFilteringRule) URLFilteringRuleState {
	urlCategories := rule.URLCategories
	if len(urlCategories) == 0 {
		urlCategories = []string{"ANY"}
	}

	sizeQuotaMB := 0
	if rule.SizeQuota > 0 {
		sizeQuotaMB = rule.SizeQuota / 1024
	}

	var validityStartTime, validityEndTime *string
	if rule.ValidityStartTime != 0 {
		s := time.Unix(int64(rule.ValidityStartTime), 0).UTC().Format(time.RFC1123)
		validityStartTime = &s
	}
	if rule.ValidityEndTime != 0 {
		s := time.Unix(int64(rule.ValidityEndTime), 0).UTC().Format(time.RFC1123)
		validityEndTime = &s
	}

	state := URLFilteringRuleState{
		URLFilteringRuleArgs: URLFilteringRuleArgs{
			Name:                   rule.Name,
			Description:            stringPtr(rule.Description),
			Order:                  rule.Order,
			Protocols:              rule.Protocols,
			URLCategories:          urlCategories,
			SourceCountries:        processCountriesFromAPI(rule.SourceCountries),
			State:                  stringPtr(rule.State),
			UserAgentTypes:         rule.UserAgentTypes,
			Rank:                   intPtr(rule.Rank),
			EndUserNotificationURL: stringPtr(rule.EndUserNotificationURL),
			BlockOverride:          boolPtr(rule.BlockOverride),
			BrowserEunTemplateID:   intPtr(rule.BrowserEunTemplateID),
			TimeQuota:              intPtr(rule.TimeQuota),
			SizeQuota:              intPtr(sizeQuotaMB),
			RequestMethods:         rule.RequestMethods,
			ValidityStartTime:      validityStartTime,
			ValidityEndTime:        validityEndTime,
			ValidityTimeZoneID:     stringPtr(rule.ValidityTimeZoneID),
			EnforceTimeValidity:    boolPtr(rule.EnforceTimeValidity),
			Action:                 stringPtr(rule.Action),
			Ciparule:               boolPtr(rule.Ciparule),
			Locations:              idNameExtensionsToIDs(rule.Locations),
			Groups:                 idNameExtensionsToIDs(rule.Groups),
			Departments:            idNameExtensionsToIDs(rule.Departments),
			Users:                  idNameExtensionsToIDs(rule.Users),
			TimeWindows:            idNameExtensionsToIDs(rule.TimeWindows),
			LocationGroups:         idNameExtensionsToIDs(rule.LocationGroups),
			Labels:                 idNameExtensionsToIDs(rule.Labels),
			DeviceGroups:           idNameExtensionsToIDs(rule.DeviceGroups),
			Devices:                idNameExtensionsToIDs(rule.Devices),
			SourceIPGroups:         idNameExtensionsToIDs(rule.SourceIPGroups),
			DeviceTrustLevels:      rule.DeviceTrustLevels,
			UserRiskScoreLevels:    rule.UserRiskScoreLevels,
		},
		RuleID: &rule.ID,
	}

	if rule.Action == "BLOCK" && rule.BlockOverride {
		state.OverrideUsers = idNameExtensionsToIDs(rule.OverrideUsers)
		state.OverrideGroups = idNameExtensionsToIDs(rule.OverrideGroups)
	}

	// Workload groups
	if len(rule.WorkloadGroups) > 0 {
		state.WorkloadGroups = make([]WorkloadGroupInput, len(rule.WorkloadGroups))
		for i, wg := range rule.WorkloadGroups {
			state.WorkloadGroups[i] = WorkloadGroupInput{ID: wg.ID, Name: stringPtr(wg.Name)}
		}
	}

	// CBI profile (preserve for ISOLATE when API doesn't return it)
	if rule.Action == "ISOLATE" && (rule.CBIProfile == nil || rule.CBIProfile.ID == "") {
		log.Printf("[DEBUG] API did not return cbi_profile for ISOLATE rule %d", rule.ID)
	} else if fp := flattenCBIProfileSimple(rule.CBIProfile); fp != nil {
		state.CBIProfile = &CBIProfileInput{
			ID:   stringPtr(fp.ID),
			Name: stringPtr(fp.Name),
			URL:  stringPtr(fp.URL),
		}
	}

	return state
}

func urlFilteringRuleStateToArgs(rule *urlfilteringpolicies.URLFilteringRule) URLFilteringRuleArgs {
	state := urlFilteringRuleAPIToState(rule)
	return state.URLFilteringRuleArgs
}

func currentOrderVsRankWording(ctx context.Context, client *zia.Client) string {
	list, err := urlfilteringpolicies.GetAll(ctx, client.Service)
	if err != nil {
		return ""
	}
	var parts []string
	for _, r := range list {
		parts = append(parts, fmt.Sprintf("Rank %d VS Order %d", r.Rank, r.Order))
	}
	return strings.Join(parts, ", ")
}
