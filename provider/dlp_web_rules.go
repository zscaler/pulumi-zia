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

// Package provider implements the DLP Web Rules resource.
// Uses dlp/dlp_web_rules package. Skips complex sub-rules/parent/child for initial impl.

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
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/dlp/dlp_web_rules"
)

const dlpWebRulesResourceType = "dlp_web_rules"

// DlpWebRule implements the zia:index:DlpWebRule resource.
type DlpWebRule struct{}

// DlpWebRuleArgs are the inputs. Skips SubRules, ParentRule, Receiver, etc. for initial impl.
type DlpWebRuleArgs struct {
	Name                     string   `pulumi:"name"`
	Order                    int      `pulumi:"order"`
	Description              *string  `pulumi:"description,optional"`
	State                    *string  `pulumi:"state,optional"`
	Rank                     *int     `pulumi:"rank,optional"`
	Action                   *string  `pulumi:"action,optional"`
	FileTypes                []string `pulumi:"fileTypes,optional"`
	Protocols                []string `pulumi:"protocols,optional"`
	CloudApplications        []string `pulumi:"cloudApplications,optional"`
	MinSize                  *int     `pulumi:"minSize,optional"`
	MatchOnly                *bool    `pulumi:"matchOnly,optional"`
	WithoutContentInspection *bool    `pulumi:"withoutContentInspection,optional"`
	OcrEnabled               *bool    `pulumi:"ocrEnabled,optional"`
	DlpDownloadScanEnabled   *bool    `pulumi:"dlpDownloadScanEnabled,optional"`
	ZccNotificationsEnabled  *bool    `pulumi:"zccNotificationsEnabled,optional"`
	ExternalAuditorEmail     *string  `pulumi:"externalAuditorEmail,optional"`
	Locations                []int    `pulumi:"locations,optional"`
	LocationGroups           []int    `pulumi:"locationGroups,optional"`
	Departments              []int    `pulumi:"departments,optional"`
	Groups                   []int    `pulumi:"groups,optional"`
	Users                    []int    `pulumi:"users,optional"`
	TimeWindows              []int    `pulumi:"timeWindows,optional"`
	Labels                   []int    `pulumi:"labels,optional"`
	SourceIpGroups           []int    `pulumi:"sourceIpGroups,optional"`
}

// DlpWebRuleState is the persisted state.
type DlpWebRuleState struct {
	DlpWebRuleArgs
	RuleID *int `pulumi:"ruleId"`
}

func dlpWebRuleArgsToAPI(args *DlpWebRuleArgs, id int) dlp_web_rules.WebDLPRules {
	order := args.Order
	if order == 0 {
		order = 1
	}
	state := ptrToString(args.State)
	if state == "" {
		state = "ENABLED"
	}
	return dlp_web_rules.WebDLPRules{
		ID:                       id,
		Name:                     args.Name,
		Order:                    order,
		Description:              ptrToString(args.Description),
		State:                    state,
		Rank:                     ptrToIntDefault(args.Rank, 7),
		Action:                   ptrToString(args.Action),
		FileTypes:                args.FileTypes,
		Protocols:                args.Protocols,
		CloudApplications:        args.CloudApplications,
		MinSize:                  ptrToIntDefault(args.MinSize, 0),
		MatchOnly:                ptrToBool(args.MatchOnly),
		WithoutContentInspection: ptrToBool(args.WithoutContentInspection),
		OcrEnabled:               ptrToBool(args.OcrEnabled),
		DLPDownloadScanEnabled:   ptrToBool(args.DlpDownloadScanEnabled),
		ZCCNotificationsEnabled:  ptrToBool(args.ZccNotificationsEnabled),
		ExternalAuditorEmail:     ptrToString(args.ExternalAuditorEmail),
		Locations:                idsToIDNameExtensions(args.Locations),
		LocationGroups:           idsToIDNameExtensions(args.LocationGroups),
		Departments:              idsToIDNameExtensions(args.Departments),
		Groups:                   idsToIDNameExtensions(args.Groups),
		Users:                    idsToIDNameExtensions(args.Users),
		TimeWindows:              idsToIDNameExtensions(args.TimeWindows),
		Labels:                   idsToIDNameExtensions(args.Labels),
		SourceIpGroups:           idsToIDNameExtensions(args.SourceIpGroups),
	}
}

func dlpWebRuleAPIToState(api *dlp_web_rules.WebDLPRules) DlpWebRuleState {
	return DlpWebRuleState{
		DlpWebRuleArgs: DlpWebRuleArgs{
			Name:                     api.Name,
			Order:                    api.Order,
			Description:              stringPtr(api.Description),
			State:                    stringPtr(api.State),
			Rank:                     intPtr(api.Rank),
			Action:                   stringPtr(api.Action),
			FileTypes:                api.FileTypes,
			Protocols:                api.Protocols,
			CloudApplications:        api.CloudApplications,
			MinSize:                  intPtr(api.MinSize),
			MatchOnly:                boolPtr(api.MatchOnly),
			WithoutContentInspection: boolPtr(api.WithoutContentInspection),
			OcrEnabled:               boolPtr(api.OcrEnabled),
			DlpDownloadScanEnabled:   boolPtr(api.DLPDownloadScanEnabled),
			ZccNotificationsEnabled:  boolPtr(api.ZCCNotificationsEnabled),
			ExternalAuditorEmail:     stringPtr(api.ExternalAuditorEmail),
			Locations:                idNameExtensionsToIDs(api.Locations),
			LocationGroups:           idNameExtensionsToIDs(api.LocationGroups),
			Departments:              idNameExtensionsToIDs(api.Departments),
			Groups:                   idNameExtensionsToIDs(api.Groups),
			Users:                    idNameExtensionsToIDs(api.Users),
			TimeWindows:              idNameExtensionsToIDs(api.TimeWindows),
			Labels:                   idNameExtensionsToIDs(api.Labels),
			SourceIpGroups:           idNameExtensionsToIDs(api.SourceIpGroups),
		},
		RuleID: intPtr(api.ID),
	}
}

func (DlpWebRule) Create(ctx context.Context, req infer.CreateRequest[DlpWebRuleArgs]) (infer.CreateResponse[DlpWebRuleState], error) {
	if req.DryRun {
		s := DlpWebRuleState{DlpWebRuleArgs: req.Inputs, RuleID: intPtr(0)}
		return infer.CreateResponse[DlpWebRuleState]{ID: "preview", Output: s}, nil
	}
	if req.Inputs.Order < 1 {
		return infer.CreateResponse[DlpWebRuleState]{}, fmt.Errorf("order must be a positive whole number (>= 1), got %d", req.Inputs.Order)
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[DlpWebRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := dlpWebRuleArgsToAPI(&req.Inputs, 0)
	timeout := 60 * time.Minute
	start := time.Now()

	intendedOrder := apiReq.Order
	intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)

	for {
		select {
		case dlpWebRulesSem <- struct{}{}:
		case <-ctx.Done():
			return infer.CreateResponse[DlpWebRuleState]{}, fmt.Errorf("creating DLP web rule: %w", ctx.Err())
		}

		dlpWebOrderMu.Lock()
		if dlpWebStartingOrder == 0 {
			list, _ := dlp_web_rules.GetAll(ctx, svc)
			for _, r := range list {
				if r.Order > dlpWebStartingOrder {
					dlpWebStartingOrder = r.Order
				}
			}
			if dlpWebStartingOrder == 0 {
				dlpWebStartingOrder = 1
			}
		}

		if intendedRank < 7 {
			apiReq.Rank = 7
		}
		apiReq.Order = dlpWebStartingOrder
		dlpWebOrderMu.Unlock()

		resp, err := dlp_web_rules.Create(ctx, svc, &apiReq)

		if err == nil {
			dlpWebOrderMu.Lock()
			dlpWebStartingOrder++
			dlpWebOrderMu.Unlock()
		}

		<-dlpWebRulesSem

		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.CreateResponse[DlpWebRuleState]{}, customErr
		}
		if err != nil {
			reg := regexp.MustCompile("Rule with rank [0-9]+ is not allowed at order [0-9]+")
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") && reg.MatchString(err.Error()) {
				return infer.CreateResponse[DlpWebRuleState]{}, fmt.Errorf("error creating DLP web rule: %s, check order %d vs rank %d, err:%s",
					apiReq.Name, intendedOrder, intendedRank, err)
			}
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Creating DLP web rule name: %v, got INVALID_INPUT_ARGUMENT\n", apiReq.Name)
				dlpWebOrderMu.Lock()
				dlpWebStartingOrder = 0
				dlpWebOrderMu.Unlock()
				if time.Since(start) < timeout {
					select {
					case <-time.After(10 * time.Second):
					case <-ctx.Done():
						return infer.CreateResponse[DlpWebRuleState]{}, fmt.Errorf("creating DLP web rule: %w", ctx.Err())
					}
					continue
				}
			}
			return infer.CreateResponse[DlpWebRuleState]{}, fmt.Errorf("creating DLP web rule: %w", err)
		}

		reorderWithBeforeReorder(
			OrderRule{Order: intendedOrder, Rank: intendedRank},
			resp.ID,
			dlpWebRulesResourceType,
			func() (int, error) {
				allRules, err := dlp_web_rules.GetAll(ctx, svc)
				if err != nil {
					return 0, err
				}
				return len(allRules), nil
			},
			func(id int, order OrderRule) error {
				rule, err := dlp_web_rules.Get(ctx, svc, id)
				if err != nil {
					return err
				}
				// Strip read-only fields that cause "Request body is invalid" for predefined rules
				rule.AccessControl = ""
				rule.Order = order.Order
				rule.Rank = order.Rank
				_, err = dlp_web_rules.Update(ctx, svc, id, rule)
				return err
			},
			nil,
		)

		markOrderRuleAsDone(resp.ID, dlpWebRulesResourceType)
		waitForReorder(dlpWebRulesResourceType)

		if shouldActivate() {
			time.Sleep(2 * time.Second)
			if activationErr := triggerActivation(ctx, client); activationErr != nil {
				return infer.CreateResponse[DlpWebRuleState]{}, activationErr
			}
		} else {
			log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
		}

		rule, err := dlp_web_rules.Get(ctx, svc, resp.ID)
		if err != nil {
			state := DlpWebRuleState{DlpWebRuleArgs: req.Inputs, RuleID: intPtr(resp.ID)}
			return infer.CreateResponse[DlpWebRuleState]{ID: strconv.Itoa(resp.ID), Output: state}, nil
		}
		return infer.CreateResponse[DlpWebRuleState]{
			ID:     strconv.Itoa(resp.ID),
			Output: dlpWebRuleAPIToState(rule),
		}, nil
	}
}

func (DlpWebRule) Read(ctx context.Context, req infer.ReadRequest[DlpWebRuleArgs, DlpWebRuleState]) (infer.ReadResponse[DlpWebRuleArgs, DlpWebRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[DlpWebRuleArgs, DlpWebRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		rule, lookupErr := dlp_web_rules.GetByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[DlpWebRuleArgs, DlpWebRuleState]{}, fmt.Errorf("DLP web rule not found")
			}
			return infer.ReadResponse[DlpWebRuleArgs, DlpWebRuleState]{}, lookupErr
		}
		id = rule.ID
	}

	rule, err := dlp_web_rules.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[DlpWebRuleArgs, DlpWebRuleState]{}, fmt.Errorf("DLP web rule not found")
		}
		return infer.ReadResponse[DlpWebRuleArgs, DlpWebRuleState]{}, err
	}

	state := dlpWebRuleAPIToState(rule)
	return infer.ReadResponse[DlpWebRuleArgs, DlpWebRuleState]{
		ID:     req.ID,
		Inputs: state.DlpWebRuleArgs,
		State:  state,
	}, nil
}

func (DlpWebRule) Update(ctx context.Context, req infer.UpdateRequest[DlpWebRuleArgs, DlpWebRuleState]) (infer.UpdateResponse[DlpWebRuleState], error) {
	if req.Inputs.Order < 1 {
		return infer.UpdateResponse[DlpWebRuleState]{}, fmt.Errorf("order must be a positive whole number (>= 1), got %d", req.Inputs.Order)
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[DlpWebRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[DlpWebRuleState]{}, fmt.Errorf("invalid DLP web rule ID: %s", req.ID)
	}
	apiReq := dlpWebRuleArgsToAPI(&req.Inputs, id)

	intendedOrder := apiReq.Order
	intendedRank := apiReq.Rank

	timeout := 60 * time.Minute
	start := time.Now()
	for {
		if _, err := dlp_web_rules.Update(ctx, svc, id, &apiReq); err != nil {
			if customErr := failFastOnErrorCodes(err); customErr != nil {
				return infer.UpdateResponse[DlpWebRuleState]{}, customErr
			}
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Updating DLP web rule ID: %v, got INVALID_INPUT_ARGUMENT\n", id)
				if time.Since(start) < timeout {
					time.Sleep(10 * time.Second)
					continue
				}
			}
			return infer.UpdateResponse[DlpWebRuleState]{}, err
		}

		reorderWithBeforeReorder(
			OrderRule{Order: intendedOrder, Rank: intendedRank},
			id,
			dlpWebRulesResourceType,
			func() (int, error) {
				allRules, err := dlp_web_rules.GetAll(ctx, svc)
				if err != nil {
					return 0, err
				}
				return len(allRules), nil
			},
			func(ruleID int, order OrderRule) error {
				rule, err := dlp_web_rules.Get(ctx, svc, ruleID)
				if err != nil {
					return err
				}
				// Optional: avoid unnecessary updates if the current order is already correct
				if rule.Order == order.Order && rule.Rank == order.Rank {
					return nil
				}
				// Strip read-only fields that cause "Request body is invalid" for predefined rules
				rule.AccessControl = ""
				rule.Order = order.Order
				rule.Rank = order.Rank
				_, err = dlp_web_rules.Update(ctx, svc, ruleID, rule)
				return err
			},
			nil,
		)

		markOrderRuleAsDone(id, dlpWebRulesResourceType)
		waitForReorder(dlpWebRulesResourceType)
		break
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[DlpWebRuleState]{}, activationErr
		}
	}

	updated, err := dlp_web_rules.Get(ctx, svc, id)
	if err != nil {
		return infer.UpdateResponse[DlpWebRuleState]{Output: DlpWebRuleState{
			DlpWebRuleArgs: req.Inputs,
			RuleID:         intPtr(id),
		}}, nil
	}
	return infer.UpdateResponse[DlpWebRuleState]{Output: dlpWebRuleAPIToState(updated)}, nil
}

func (DlpWebRule) Delete(ctx context.Context, req infer.DeleteRequest[DlpWebRuleState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid DLP web rule ID: %s", req.ID)
	}
	if _, err := dlp_web_rules.Delete(ctx, svc, id); err != nil {
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

func (DlpWebRule) Annotate(a infer.Annotator) {
	describeResource(a, &DlpWebRule{}, `The zia_dlp_web_rules resource manages DLP (Data Loss Prevention) web rules in the Zscaler Internet Access (ZIA) cloud service. DLP web rules define how sensitive data is handled in web traffic, allowing organizations to control and monitor the transfer of confidential information.

For more information, see the [ZIA Data Loss Prevention documentation](https://help.zscaler.com/zia/data-loss-prevention).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic DLP Web Rule

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.DlpWebRule("example", {
    name: "Example DLP Web Rule",
    description: "Block sensitive data uploads",
    order: 1,
    state: "ENABLED",
    action: "BLOCK",
    protocols: ["FTP_RULE", "HTTPS_RULE", "HTTP_RULE"],
    fileTypes: ["ALL_OUTBOUND"],
    zccNotificationsEnabled: true,
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.DlpWebRule("example",
    name="Example DLP Web Rule",
    description="Block sensitive data uploads",
    order=1,
    state="ENABLED",
    action="BLOCK",
    protocols=["FTP_RULE", "HTTPS_RULE", "HTTP_RULE"],
    file_types=["ALL_OUTBOUND"],
    zcc_notifications_enabled=True,
)
`+tripleBacktick("")+`

`+tripleBacktick("go")+`
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zia "github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := zia.NewDlpWebRule(ctx, "example", &zia.DlpWebRuleArgs{
			Name:                    pulumi.String("Example DLP Web Rule"),
			Description:             pulumi.StringRef("Block sensitive data uploads"),
			Order:                   pulumi.Int(1),
			State:                   pulumi.StringRef("ENABLED"),
			Action:                  pulumi.StringRef("BLOCK"),
			Protocols:               pulumi.ToStringArray([]string{"FTP_RULE", "HTTPS_RULE", "HTTP_RULE"}),
			FileTypes:               pulumi.ToStringArray([]string{"ALL_OUTBOUND"}),
			ZccNotificationsEnabled: pulumi.BoolRef(true),
		})
		return err
	})
}
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:DlpWebRule
    properties:
      name: Example DLP Web Rule
      description: Block sensitive data uploads
      order: 1
      state: ENABLED
      action: BLOCK
      protocols:
        - FTP_RULE
        - HTTPS_RULE
        - HTTP_RULE
      fileTypes:
        - ALL_OUTBOUND
      zccNotificationsEnabled: true
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing DLP Web Rule can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:DlpWebRule example 12345
`+tripleBacktick("")+`
`)
}

func (a *DlpWebRuleArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the DLP web rule. Must be unique.")
	ann.Describe(&a.Order, "The order of execution of the rule with respect to other DLP web rules.")
	ann.Describe(&a.Description, "Additional information about the DLP web rule.")
	ann.Describe(&a.State, "Rule state. Valid values: `ENABLED`, `DISABLED`.")
	ann.Describe(&a.Rank, "Admin rank of the DLP web rule. Valid values: 0-7. Default: 7.")
	ann.Describe(&a.Action, "Action taken when the rule is matched. Valid values: `ALLOW`, `BLOCK`, `ICAP_RESPONSE`.")
	ann.Describe(&a.FileTypes, "List of file types to which the DLP policy rule must be applied.")
	ann.Describe(&a.Protocols, "Protocols to which the rule applies. Valid values: `FTP_RULE`, `HTTPS_RULE`, `HTTP_RULE`.")
	ann.Describe(&a.CloudApplications, "List of cloud application names for which the rule is applied.")
	ann.Describe(&a.MinSize, "Minimum file size (in KB) used for evaluating the DLP policy rule.")
	ann.Describe(&a.MatchOnly, "If true, the rule matches but does not enforce the action.")
	ann.Describe(&a.WithoutContentInspection, "If true, the DLP rule is applied without inspecting content.")
	ann.Describe(&a.OcrEnabled, "If true, Optical Character Recognition (OCR) is enabled for the DLP rule.")
	ann.Describe(&a.DlpDownloadScanEnabled, "If true, DLP scanning is enabled for file downloads.")
	ann.Describe(&a.ZccNotificationsEnabled, "If true, Zscaler Client Connector notifications are enabled for this rule.")
	ann.Describe(&a.ExternalAuditorEmail, "The email address of an external auditor to whom DLP email notifications are sent.")
	ann.Describe(&a.Locations, "IDs of locations for which the rule must be applied.")
	ann.Describe(&a.LocationGroups, "IDs of location groups for which the rule must be applied.")
	ann.Describe(&a.Departments, "IDs of departments for which the rule must be applied.")
	ann.Describe(&a.Groups, "IDs of groups for which the rule must be applied.")
	ann.Describe(&a.Users, "IDs of users for which the rule must be applied.")
	ann.Describe(&a.TimeWindows, "IDs of time intervals during which the rule must be enforced.")
	ann.Describe(&a.Labels, "IDs of labels associated with the DLP web rule.")
	ann.Describe(&a.SourceIpGroups, "IDs of source IP address groups for which the rule must be applied.")
}

func (s *DlpWebRuleState) Annotate(a infer.Annotator) {
	a.Describe(&s.RuleID, "The system-generated ID of the DLP web rule.")
}

var _ infer.CustomResource[DlpWebRuleArgs, DlpWebRuleState] = DlpWebRule{}
