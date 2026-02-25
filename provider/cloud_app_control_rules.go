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
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/common"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/cloudappcontrol"
)

const cloudAppControlResourceType = "cloud_app_control_rules"

type CloudAppControlRule struct{}

type CloudAppControlRuleArgs struct {
	Name                string   `pulumi:"name"`
	Type                string   `pulumi:"type"`
	Order               int      `pulumi:"order"`
	Description         *string  `pulumi:"description,optional"`
	State               *string  `pulumi:"state,optional"`
	Rank                *int     `pulumi:"rank,optional"`
	Actions             []string `pulumi:"actions,optional"`
	Applications        []string `pulumi:"applications,optional"`
	TimeQuota           *int     `pulumi:"timeQuota,optional"`
	SizeQuota           *int     `pulumi:"sizeQuota,optional"`
	EnforceTimeValidity *bool    `pulumi:"enforceTimeValidity,optional"`
	CascadingEnabled    *bool    `pulumi:"cascadingEnabled,optional"`
	EunEnabled          *bool    `pulumi:"eunEnabled,optional"`
	EunTemplateId       *int     `pulumi:"eunTemplateId,optional"`
	BrowserEunTemplateId *int    `pulumi:"browserEunTemplateId,optional"`
	Locations           []int    `pulumi:"locations,optional"`
	LocationGroups      []int    `pulumi:"locationGroups,optional"`
	Groups              []int    `pulumi:"groups,optional"`
	Departments         []int    `pulumi:"departments,optional"`
	Users               []int    `pulumi:"users,optional"`
	TimeWindows         []int    `pulumi:"timeWindows,optional"`
	Labels              []int    `pulumi:"labels,optional"`
	DeviceGroups        []int    `pulumi:"deviceGroups,optional"`
	Devices             []int    `pulumi:"devices,optional"`
	TenancyProfileIds   []int    `pulumi:"tenancyProfileIds,optional"`
	CloudAppRiskProfileId *int   `pulumi:"cloudAppRiskProfileId,optional"`
	CbiProfile          *CBIProfileInput `pulumi:"cbiProfile,optional"`
}

type CloudAppControlRuleState struct {
	CloudAppControlRuleArgs
	RuleID *int `pulumi:"ruleId"`
}

func cloudAppControlRuleArgsToAPI(args *CloudAppControlRuleArgs, id int) cloudappcontrol.WebApplicationRules {
	order := args.Order
	if order == 0 {
		order = 1
	}
	rank := ptrToIntDefault(args.Rank, 7)
	sizeQuota := 0
	if args.SizeQuota != nil {
		sizeQuota = *args.SizeQuota * 1024 // MB to KB
	}
	timeQuota := 0
	if args.TimeQuota != nil {
		timeQuota = *args.TimeQuota
	}
	eunID, browserEunID := 0, 0
	if args.EunTemplateId != nil {
		eunID = *args.EunTemplateId
	}
	if args.BrowserEunTemplateId != nil {
		browserEunID = *args.BrowserEunTemplateId
	}
	cbi := cloudappcontrol.CBIProfile{}
	if args.CbiProfile != nil {
		cbi = cloudappcontrol.CBIProfile{
			ID:   ptrToString(args.CbiProfile.ID),
			Name: ptrToString(args.CbiProfile.Name),
			URL:  ptrToString(args.CbiProfile.URL),
		}
	}
	var cloudAppRisk *common.IDCustom
	if args.CloudAppRiskProfileId != nil && *args.CloudAppRiskProfileId != 0 {
		cloudAppRisk = &common.IDCustom{ID: *args.CloudAppRiskProfileId}
	}
	return cloudappcontrol.WebApplicationRules{
		ID:                   id,
		Name:                 args.Name,
		Type:                 args.Type,
		Description:          ptrToString(args.Description),
		Order:                order,
		Rank:                 rank,
		State:                ptrToString(args.State),
		TimeQuota:            timeQuota,
		SizeQuota:            sizeQuota,
		EnforceTimeValidity:  ptrToBool(args.EnforceTimeValidity),
		CascadingEnabled:     ptrToBool(args.CascadingEnabled),
		EunEnabled:           ptrToBool(args.EunEnabled),
		EunTemplateID:        eunID,
		BrowserEunTemplateID: browserEunID,
		Actions:              args.Actions,
		Applications:         args.Applications,
		Locations:            idsToIDNameExtensions(args.Locations),
		LocationGroups:       idsToIDNameExtensions(args.LocationGroups),
		Groups:               idsToIDNameExtensions(args.Groups),
		Departments:          idsToIDNameExtensions(args.Departments),
		Users:                idsToIDNameExtensions(args.Users),
		TimeWindows:          idsToIDNameExtensions(args.TimeWindows),
		Labels:               idsToIDNameExtensions(args.Labels),
		DeviceGroups:         idsToIDNameExtensions(args.DeviceGroups),
		Devices:              idsToIDNameExtensions(args.Devices),
		TenancyProfileIDs:    idsToIDNameExtensions(args.TenancyProfileIds),
		CloudAppRiskProfile:  cloudAppRisk,
		CBIProfile:           cbi,
	}
}

func cloudAppControlRuleAPIToState(api *cloudappcontrol.WebApplicationRules) CloudAppControlRuleState {
	sizeQuotaMB := api.SizeQuota / 1024
	if sizeQuotaMB == 0 && api.SizeQuota > 0 {
		sizeQuotaMB = 1
	}
	var cbi *CBIProfileInput
	if api.CBIProfile.ID != "" || api.CBIProfile.Name != "" || api.CBIProfile.URL != "" {
		cbi = &CBIProfileInput{
			ID:   stringPtr(api.CBIProfile.ID),
			Name: stringPtr(api.CBIProfile.Name),
			URL:  stringPtr(api.CBIProfile.URL),
		}
	}
	var cloudAppRiskID *int
	if api.CloudAppRiskProfile != nil && api.CloudAppRiskProfile.ID != 0 {
		cloudAppRiskID = intPtr(api.CloudAppRiskProfile.ID)
	}
	return CloudAppControlRuleState{
		CloudAppControlRuleArgs: CloudAppControlRuleArgs{
			Name:                 api.Name,
			Type:                 api.Type,
			Order:                api.Order,
			Description:          stringPtr(api.Description),
			State:                stringPtr(api.State),
			Rank:                 intPtr(api.Rank),
			Actions:              api.Actions,
			Applications:         api.Applications,
			TimeQuota:            intPtr(api.TimeQuota),
			SizeQuota:            intPtr(sizeQuotaMB),
			EnforceTimeValidity:  boolPtr(api.EnforceTimeValidity),
			CascadingEnabled:     boolPtr(api.CascadingEnabled),
			EunEnabled:           boolPtr(api.EunEnabled),
			EunTemplateId:        intPtr(api.EunTemplateID),
			BrowserEunTemplateId: intPtr(api.BrowserEunTemplateID),
			Locations:            idNameExtensionsToIDs(api.Locations),
			LocationGroups:       idNameExtensionsToIDs(api.LocationGroups),
			Groups:               idNameExtensionsToIDs(api.Groups),
			Departments:          idNameExtensionsToIDs(api.Departments),
			Users:                idNameExtensionsToIDs(api.Users),
			TimeWindows:          idNameExtensionsToIDs(api.TimeWindows),
			Labels:               idNameExtensionsToIDs(api.Labels),
			DeviceGroups:         idNameExtensionsToIDs(api.DeviceGroups),
			Devices:              idNameExtensionsToIDs(api.Devices),
			TenancyProfileIds:    idNameExtensionsToIDs(api.TenancyProfileIDs),
			CloudAppRiskProfileId: cloudAppRiskID,
			CbiProfile:           cbi,
		},
		RuleID: intPtr(api.ID),
	}
}

func (CloudAppControlRule) Create(ctx context.Context, req infer.CreateRequest[CloudAppControlRuleArgs]) (infer.CreateResponse[CloudAppControlRuleState], error) {
	if req.DryRun {
		return infer.CreateResponse[CloudAppControlRuleState]{ID: "preview", Output: CloudAppControlRuleState{CloudAppControlRuleArgs: req.Inputs, RuleID: intPtr(0)}}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[CloudAppControlRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := cloudAppControlRuleArgsToAPI(&req.Inputs, 0)
	ruleType := apiReq.Type
	timeout := 60 * time.Minute
	start := time.Now()

	for {
		cloudAppRuleLock.Lock()
		if cloudAppRuleStartingOrder == 0 {
			rules, _ := cloudappcontrol.GetByRuleType(ctx, svc, ruleType)
			for _, r := range rules {
				if r.Order > cloudAppRuleStartingOrder {
					cloudAppRuleStartingOrder = r.Order
				}
			}
			if cloudAppRuleStartingOrder == 0 {
				cloudAppRuleStartingOrder = 1
			}
		}
		cloudAppRuleLock.Unlock()

		intendedOrder := apiReq.Order
		intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)
		if intendedRank < 7 {
			apiReq.Rank = 7
		}
		apiReq.Order = cloudAppRuleStartingOrder

		resp, err := cloudappcontrol.Create(ctx, svc, ruleType, &apiReq)
		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.CreateResponse[CloudAppControlRuleState]{}, customErr
		}
		if err != nil {
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Creating cloud app control rule name: %v, got INVALID_INPUT_ARGUMENT\n", apiReq.Name)
				if time.Since(start) < timeout {
					time.Sleep(10 * time.Second)
					continue
				}
			}
			return infer.CreateResponse[CloudAppControlRuleState]{}, fmt.Errorf("creating cloud app control rule: %w", err)
		}

		reorderWithBeforeReorder(
			OrderRule{Order: intendedOrder, Rank: intendedRank},
			resp.ID,
			cloudAppControlResourceType,
			func() (int, error) {
				rules, err := cloudappcontrol.GetByRuleType(ctx, svc, ruleType)
				if err != nil {
					return 0, err
				}
				return len(rules), nil
			},
			func(id int, order OrderRule) error {
				rule, err := cloudappcontrol.GetByRuleID(ctx, svc, ruleType, id)
				if err != nil {
					return err
				}
				// to avoid the STALE_CONFIGURATION_ERROR
				rule.LastModifiedTime = 0
				rule.Order = order.Order
				rule.Rank = order.Rank
				_, err = cloudappcontrol.Update(ctx, svc, ruleType, id, rule)
				return err
			},
			nil,
		)
		markOrderRuleAsDone(resp.ID, cloudAppControlResourceType)
		waitForReorder(cloudAppControlResourceType)

		if shouldActivate() {
			time.Sleep(2 * time.Second)
			if activationErr := triggerActivation(ctx, client); activationErr != nil {
				return infer.CreateResponse[CloudAppControlRuleState]{}, activationErr
			}
		} else {
			log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
		}

		rule, err := cloudappcontrol.GetByRuleID(ctx, svc, ruleType, resp.ID)
		if err != nil {
			return infer.CreateResponse[CloudAppControlRuleState]{ID: strconv.Itoa(resp.ID), Output: CloudAppControlRuleState{CloudAppControlRuleArgs: req.Inputs, RuleID: intPtr(resp.ID)}}, nil
		}
		return infer.CreateResponse[CloudAppControlRuleState]{ID: strconv.Itoa(resp.ID), Output: cloudAppControlRuleAPIToState(rule)}, nil
	}
}

func (CloudAppControlRule) Read(ctx context.Context, req infer.ReadRequest[CloudAppControlRuleArgs, CloudAppControlRuleState]) (infer.ReadResponse[CloudAppControlRuleArgs, CloudAppControlRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[CloudAppControlRuleArgs, CloudAppControlRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.ReadResponse[CloudAppControlRuleArgs, CloudAppControlRuleState]{}, fmt.Errorf("cloud app control rule ID must be numeric")
	}
	ruleType := req.State.Type
	if ruleType == "" {
		ruleType = req.Inputs.Type
	}
	if ruleType == "" {
		return infer.ReadResponse[CloudAppControlRuleArgs, CloudAppControlRuleState]{}, fmt.Errorf("rule type is required")
	}

	rule, err := cloudappcontrol.GetByRuleID(ctx, svc, ruleType, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[CloudAppControlRuleArgs, CloudAppControlRuleState]{}, fmt.Errorf("cloud app control rule not found")
		}
		return infer.ReadResponse[CloudAppControlRuleArgs, CloudAppControlRuleState]{}, err
	}

	state := cloudAppControlRuleAPIToState(rule)
	state.Type = ruleType
	return infer.ReadResponse[CloudAppControlRuleArgs, CloudAppControlRuleState]{ID: req.ID, Inputs: state.CloudAppControlRuleArgs, State: state}, nil
}

func (CloudAppControlRule) Update(ctx context.Context, req infer.UpdateRequest[CloudAppControlRuleArgs, CloudAppControlRuleState]) (infer.UpdateResponse[CloudAppControlRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[CloudAppControlRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[CloudAppControlRuleState]{}, fmt.Errorf("invalid cloud app control rule ID: %s", req.ID)
	}
	apiReq := cloudAppControlRuleArgsToAPI(&req.Inputs, id)
	ruleType := apiReq.Type

	existingRules, err := cloudappcontrol.GetByRuleType(ctx, svc, ruleType)
	if err == nil && len(existingRules) > 0 {
		sort.Slice(existingRules, func(i, j int) bool {
			return existingRules[i].Rank < existingRules[j].Rank || (existingRules[i].Rank == existingRules[j].Rank && existingRules[i].Order < existingRules[j].Order)
		})
		apiReq.Rank = 7
		apiReq.Order = existingRules[len(existingRules)-1].Order
	}

	_, err = cloudappcontrol.Update(ctx, svc, ruleType, id, &apiReq)
	if customErr := failFastOnErrorCodes(err); customErr != nil {
		return infer.UpdateResponse[CloudAppControlRuleState]{}, customErr
	}
	if err != nil {
		return infer.UpdateResponse[CloudAppControlRuleState]{}, err
	}

	intendedOrder := req.Inputs.Order
	intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)

	reorderWithBeforeReorder(
		OrderRule{Order: intendedOrder, Rank: intendedRank},
		id,
		cloudAppControlResourceType,
		func() (int, error) {
			rules, err := cloudappcontrol.GetByRuleType(ctx, svc, ruleType)
			if err != nil {
				return 0, err
			}
			return len(rules), nil
		},
		func(ruleID int, order OrderRule) error {
			rule, err := cloudappcontrol.GetByRuleID(ctx, svc, ruleType, ruleID)
			if err != nil {
				return err
			}
			// to avoid the STALE_CONFIGURATION_ERROR
			rule.LastModifiedTime = 0
			rule.Order = order.Order
			rule.Rank = order.Rank
			_, err = cloudappcontrol.Update(ctx, svc, ruleType, ruleID, rule)
			return err
		},
		nil,
	)
	markOrderRuleAsDone(id, cloudAppControlResourceType)
	waitForReorder(cloudAppControlResourceType)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[CloudAppControlRuleState]{}, activationErr
		}
	}

	updated, err := cloudappcontrol.GetByRuleID(ctx, svc, ruleType, id)
	if err != nil {
		return infer.UpdateResponse[CloudAppControlRuleState]{Output: CloudAppControlRuleState{CloudAppControlRuleArgs: req.Inputs, RuleID: intPtr(id)}}, nil
	}
	return infer.UpdateResponse[CloudAppControlRuleState]{Output: cloudAppControlRuleAPIToState(updated)}, nil
}

func (CloudAppControlRule) Delete(ctx context.Context, req infer.DeleteRequest[CloudAppControlRuleState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid cloud app control rule ID: %s", req.ID)
	}
	ruleType := req.State.Type
	if ruleType == "" {
		return infer.DeleteResponse{}, fmt.Errorf("rule type is required for delete")
	}
	if _, err := cloudappcontrol.Delete(ctx, svc, ruleType, id); err != nil {
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

func (CloudAppControlRule) Annotate(a infer.Annotator) {
	describeResource(a, &CloudAppControlRule{}, `The zia_cloud_app_control_rules resource manages cloud application control rules in the Zscaler Internet Access (ZIA) cloud service. Cloud app control rules define policies that govern user access to cloud applications, allowing administrators to allow, block, or isolate specific application activities.

For more information, see the [ZIA Cloud App Control documentation](https://help.zscaler.com/zia/cloud-app-control).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Cloud App Control Rule

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.CloudAppControlRule("example", {
    name: "Example Cloud App Control Rule",
    description: "Block file sharing uploads",
    type: "STREAMING_MEDIA",
    order: 1,
    state: "ENABLED",
    actions: ["BLOCK"],
    applications: ["YOUTUBE"],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.CloudAppControlRule("example",
    name="Example Cloud App Control Rule",
    description="Block file sharing uploads",
    type="STREAMING_MEDIA",
    order=1,
    state="ENABLED",
    actions=["BLOCK"],
    applications=["YOUTUBE"],
)
`+tripleBacktick("")+`

`+tripleBacktick("go")+`
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zia "github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := zia.NewCloudAppControlRule(ctx, "example", &zia.CloudAppControlRuleArgs{
			Name:         pulumi.String("Example Cloud App Control Rule"),
			Description:  pulumi.StringRef("Block file sharing uploads"),
			Type:         pulumi.String("STREAMING_MEDIA"),
			Order:        pulumi.Int(1),
			State:        pulumi.StringRef("ENABLED"),
			Actions:      pulumi.ToStringArray([]string{"BLOCK"}),
			Applications: pulumi.ToStringArray([]string{"YOUTUBE"}),
		})
		return err
	})
}
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:CloudAppControlRule
    properties:
      name: Example Cloud App Control Rule
      description: Block file sharing uploads
      type: STREAMING_MEDIA
      order: 1
      state: ENABLED
      actions:
        - BLOCK
      applications:
        - YOUTUBE
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing Cloud App Control Rule can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:CloudAppControlRule example 12345
`+tripleBacktick("")+`
`)
}

func (a *CloudAppControlRuleArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the cloud app control rule. Must be unique.")
	ann.Describe(&a.Type, "The rule type, corresponding to the cloud application category. Valid values: `STREAMING_MEDIA`, `SOCIAL_NETWORKING`, `WEBMAIL`, `INSTANT_MESSAGING`, `FILE_SHARE`, `BUSINESS_PRODUCTIVITY`, `SYSTEM_AND_DEVELOPMENT`, `CONSUMER`, `HOSTING_PROVIDER`, `DNS_OVER_HTTPS`, `ENTERPRISE_COLLABORATION`, `GENERATIVE_AI`, `SALES_AND_MARKETING`, `HEALTH_CARE`, `LEGAL`, `HUMAN_RESOURCES`, `FINANCE`.")
	ann.Describe(&a.Order, "The order of execution of the rule with respect to other cloud app control rules.")
	ann.Describe(&a.Description, "Additional information about the cloud app control rule.")
	ann.Describe(&a.State, "Rule state. Valid values: `ENABLED`, `DISABLED`.")
	ann.Describe(&a.Rank, "Admin rank of the cloud app control rule. Valid values: 0-7. Default: 7.")
	ann.Describe(&a.Actions, "Actions taken when traffic matches rule criteria. Valid values: `ALLOW`, `BLOCK`, `CAUTION`, `ISOLATE`.")
	ann.Describe(&a.Applications, "List of cloud application names to which the rule applies.")
	ann.Describe(&a.TimeQuota, "Time quota in minutes, after which the rule is applied. Not applicable when action is 'BLOCK'.")
	ann.Describe(&a.SizeQuota, "Size quota in MB beyond which the rule is applied. Not applicable when action is 'BLOCK'.")
	ann.Describe(&a.EnforceTimeValidity, "Enforce a set validity time period for the rule.")
	ann.Describe(&a.CascadingEnabled, "If true, cascading to other rules is enabled when this rule matches.")
	ann.Describe(&a.EunEnabled, "If true, End User Notification is enabled for this rule.")
	ann.Describe(&a.EunTemplateId, "The ID of the End User Notification template.")
	ann.Describe(&a.BrowserEunTemplateId, "The ID of the Browser End User Notification template.")
	ann.Describe(&a.Locations, "IDs of locations for which the rule must be applied.")
	ann.Describe(&a.LocationGroups, "IDs of location groups for which the rule must be applied.")
	ann.Describe(&a.Groups, "IDs of groups for which the rule must be applied.")
	ann.Describe(&a.Departments, "IDs of departments for which the rule must be applied.")
	ann.Describe(&a.Users, "IDs of users for which the rule must be applied.")
	ann.Describe(&a.TimeWindows, "IDs of time intervals during which the rule must be enforced.")
	ann.Describe(&a.Labels, "IDs of labels associated with the cloud app control rule.")
	ann.Describe(&a.DeviceGroups, "IDs of device groups for which the rule must be applied.")
	ann.Describe(&a.Devices, "IDs of devices for which the rule must be applied.")
	ann.Describe(&a.TenancyProfileIds, "IDs of tenancy profiles for which the rule must be applied.")
	ann.Describe(&a.CloudAppRiskProfileId, "The ID of the cloud application risk profile associated with this rule.")
	ann.Describe(&a.CbiProfile, "The Cloud Browser Isolation (CBI) profile. Required when action is 'ISOLATE'.")
}

func (s *CloudAppControlRuleState) Annotate(a infer.Annotator) {
	a.Describe(&s.RuleID, "The system-generated ID of the cloud app control rule.")
}

var _ infer.CustomResource[CloudAppControlRuleArgs, CloudAppControlRuleState] = CloudAppControlRule{}
