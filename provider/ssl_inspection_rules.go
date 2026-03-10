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

// Package provider implements the SSL Inspection Rules resource.
// Adopted from terraform-provider-zia resource_zia_ssl_inspection_rules.go.

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
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/sslinspection"
)

const sslInspectionResourceType = "ssl_inspection_rules"

// SslInspectionRule implements the zia:index:SslInspectionRule resource.
type SslInspectionRule struct{}

// SslInterceptionCertInput is the ssl_interception_cert block.
type SslInterceptionCertInput struct {
	Id *int `pulumi:"id,optional"`
}

// DecryptSubActionsInput is decrypt_sub_actions block (when action type is DECRYPT).
type DecryptSubActionsInput struct {
	ServerCertificates              *string `pulumi:"serverCertificates,optional"`
	OcspCheck                       *bool   `pulumi:"ocspCheck,optional"`
	BlockSslTrafficWithNoSniEnabled *bool   `pulumi:"blockSslTrafficWithNoSniEnabled,optional"`
	MinClientTLSVersion             *string `pulumi:"minClientTlsVersion,optional"`
	MinServerTLSVersion             *string `pulumi:"minServerTlsVersion,optional"`
	BlockUndecrypt                  *bool   `pulumi:"blockUndecrypt,optional"`
	HTTP2Enabled                    *bool   `pulumi:"http2Enabled,optional"`
}

// DoNotDecryptSubActionsInput is do_not_decrypt_sub_actions block (when action type is DO_NOT_DECRYPT).
type DoNotDecryptSubActionsInput struct {
	BypassOtherPolicies             *bool   `pulumi:"bypassOtherPolicies,optional"`
	ServerCertificates              *string `pulumi:"serverCertificates,optional"`
	OcspCheck                       *bool   `pulumi:"ocspCheck,optional"`
	BlockSslTrafficWithNoSniEnabled *bool   `pulumi:"blockSslTrafficWithNoSniEnabled,optional"`
	MinTLSVersion                   *string `pulumi:"minTlsVersion,optional"`
}

// SslInspectionActionInput is the action block.
type SslInspectionActionInput struct {
	Type                       *string                      `pulumi:"type,optional"`
	ShowEUN                    *bool                        `pulumi:"showEun,optional"`
	ShowEUNATP                 *bool                        `pulumi:"showEunatp,optional"`
	OverrideDefaultCertificate *bool                        `pulumi:"overrideDefaultCertificate,optional"`
	SSLInterceptionCert        *SslInterceptionCertInput    `pulumi:"sslInterceptionCert,optional"`
	DecryptSubActions          *DecryptSubActionsInput      `pulumi:"decryptSubActions,optional"`
	DoNotDecryptSubActions     *DoNotDecryptSubActionsInput `pulumi:"doNotDecryptSubActions,optional"`
}

// SslInspectionRuleArgs are the inputs.
type SslInspectionRuleArgs struct {
	Name                   string                   `pulumi:"name"`
	Order                  int                      `pulumi:"order"`
	Description            *string                  `pulumi:"description,optional"`
	Rank                   *int                     `pulumi:"rank,optional"`
	State                  *string                  `pulumi:"state,optional"`
	RoadWarriorForKerberos *bool                    `pulumi:"roadWarriorForKerberos,optional"`
	Action                 SslInspectionActionInput `pulumi:"action"`
	CloudApplications      []string                 `pulumi:"cloudApplications,optional"`
	URLCategories          []string                 `pulumi:"urlCategories,optional"`
	DeviceTrustLevels      []string                 `pulumi:"deviceTrustLevels,optional"`
	Platforms              []string                 `pulumi:"platforms,optional"`
	UserAgentTypes         []string                 `pulumi:"userAgentTypes,optional"`
	Locations              []int                    `pulumi:"locations,optional"`
	LocationGroups         []int                    `pulumi:"locationGroups,optional"`
	Groups                 []int                    `pulumi:"groups,optional"`
	Departments            []int                    `pulumi:"departments,optional"`
	Users                  []int                    `pulumi:"users,optional"`
	TimeWindows            []int                    `pulumi:"timeWindows,optional"`
	Labels                 []int                    `pulumi:"labels,optional"`
	DeviceGroups           []int                    `pulumi:"deviceGroups,optional"`
	Devices                []int                    `pulumi:"devices,optional"`
	SourceIPGroups         []int                    `pulumi:"sourceIpGroups,optional"`
	DestIpGroups           []int                    `pulumi:"destIpGroups,optional"`
	ProxyGateways          []int                    `pulumi:"proxyGateways,optional"`
	WorkloadGroups         []WorkloadGroupInput     `pulumi:"workloadGroups,optional"`
}

// SslInspectionRuleState is the persisted state.
type SslInspectionRuleState struct {
	SslInspectionRuleArgs
	RuleID *int `pulumi:"ruleId"`
}

func sslInspectionActionToAPI(in SslInspectionActionInput) sslinspection.Action {
	action := sslinspection.Action{
		Type:                       ptrToString(in.Type),
		ShowEUN:                    ptrToBool(in.ShowEUN),
		ShowEUNATP:                 ptrToBool(in.ShowEUNATP),
		OverrideDefaultCertificate: ptrToBool(in.OverrideDefaultCertificate),
	}
	if in.SSLInterceptionCert != nil && in.SSLInterceptionCert.Id != nil {
		action.SSLInterceptionCert = &sslinspection.SSLInterceptionCert{ID: *in.SSLInterceptionCert.Id}
	}
	if in.DecryptSubActions != nil {
		action.DecryptSubActions = &sslinspection.DecryptSubActions{
			ServerCertificates:              ptrToString(in.DecryptSubActions.ServerCertificates),
			OcspCheck:                       ptrToBool(in.DecryptSubActions.OcspCheck),
			BlockSslTrafficWithNoSniEnabled: ptrToBool(in.DecryptSubActions.BlockSslTrafficWithNoSniEnabled),
			MinClientTLSVersion:             ptrToString(in.DecryptSubActions.MinClientTLSVersion),
			MinServerTLSVersion:             ptrToString(in.DecryptSubActions.MinServerTLSVersion),
			BlockUndecrypt:                  ptrToBool(in.DecryptSubActions.BlockUndecrypt),
			HTTP2Enabled:                    ptrToBool(in.DecryptSubActions.HTTP2Enabled),
		}
	}
	if in.DoNotDecryptSubActions != nil {
		action.DoNotDecryptSubActions = &sslinspection.DoNotDecryptSubActions{
			BypassOtherPolicies:             ptrToBool(in.DoNotDecryptSubActions.BypassOtherPolicies),
			ServerCertificates:              ptrToString(in.DoNotDecryptSubActions.ServerCertificates),
			OcspCheck:                       ptrToBool(in.DoNotDecryptSubActions.OcspCheck),
			BlockSslTrafficWithNoSniEnabled: ptrToBool(in.DoNotDecryptSubActions.BlockSslTrafficWithNoSniEnabled),
			MinTLSVersion:                   ptrToString(in.DoNotDecryptSubActions.MinTLSVersion),
		}
	}
	return action
}

func sslInspectionActionFromAPI(api sslinspection.Action) SslInspectionActionInput {
	in := SslInspectionActionInput{
		Type:                       stringPtr(api.Type),
		ShowEUN:                    boolPtr(api.ShowEUN),
		ShowEUNATP:                 boolPtr(api.ShowEUNATP),
		OverrideDefaultCertificate: boolPtr(api.OverrideDefaultCertificate),
	}
	if api.SSLInterceptionCert != nil {
		in.SSLInterceptionCert = &SslInterceptionCertInput{Id: intPtr(api.SSLInterceptionCert.ID)}
	}
	if api.DecryptSubActions != nil {
		in.DecryptSubActions = &DecryptSubActionsInput{
			ServerCertificates:              stringPtr(api.DecryptSubActions.ServerCertificates),
			OcspCheck:                       boolPtr(api.DecryptSubActions.OcspCheck),
			BlockSslTrafficWithNoSniEnabled: boolPtr(api.DecryptSubActions.BlockSslTrafficWithNoSniEnabled),
			MinClientTLSVersion:             stringPtr(api.DecryptSubActions.MinClientTLSVersion),
			MinServerTLSVersion:             stringPtr(api.DecryptSubActions.MinServerTLSVersion),
			BlockUndecrypt:                  boolPtr(api.DecryptSubActions.BlockUndecrypt),
			HTTP2Enabled:                    boolPtr(api.DecryptSubActions.HTTP2Enabled),
		}
	}
	if api.DoNotDecryptSubActions != nil {
		in.DoNotDecryptSubActions = &DoNotDecryptSubActionsInput{
			BypassOtherPolicies:             boolPtr(api.DoNotDecryptSubActions.BypassOtherPolicies),
			ServerCertificates:              stringPtr(api.DoNotDecryptSubActions.ServerCertificates),
			OcspCheck:                       boolPtr(api.DoNotDecryptSubActions.OcspCheck),
			BlockSslTrafficWithNoSniEnabled: boolPtr(api.DoNotDecryptSubActions.BlockSslTrafficWithNoSniEnabled),
			MinTLSVersion:                   stringPtr(api.DoNotDecryptSubActions.MinTLSVersion),
		}
	}
	return in
}

func sslInspectionRuleArgsToAPI(args *SslInspectionRuleArgs, id int) sslinspection.SSLInspectionRules {
	order := args.Order
	if order == 0 {
		order = 1
	}
	rank := ptrToIntDefault(args.Rank, 7)
	state := ptrToString(args.State)
	if state == "" {
		state = "ENABLED"
	}
	api := sslinspection.SSLInspectionRules{
		ID:                     id,
		Name:                   args.Name,
		Order:                  order,
		Rank:                   rank,
		State:                  state,
		Description:            ptrToString(args.Description),
		RoadWarriorForKerberos: ptrToBool(args.RoadWarriorForKerberos),
		Action:                 sslInspectionActionToAPI(args.Action),
		CloudApplications:      args.CloudApplications,
		URLCategories:          args.URLCategories,
		DeviceTrustLevels:      args.DeviceTrustLevels,
		Platforms:              args.Platforms,
		UserAgentTypes:         args.UserAgentTypes,
		Locations:              idsToIDNameExtensions(args.Locations),
		LocationGroups:         idsToIDNameExtensions(args.LocationGroups),
		Groups:                 idsToIDNameExtensions(args.Groups),
		Departments:            idsToIDNameExtensions(args.Departments),
		Users:                  idsToIDNameExtensions(args.Users),
		TimeWindows:            idsToIDNameExtensions(args.TimeWindows),
		Labels:                 idsToIDNameExtensions(args.Labels),
		DeviceGroups:           idsToIDNameExtensions(args.DeviceGroups),
		Devices:                idsToIDNameExtensions(args.Devices),
		SourceIPGroups:         idsToIDNameExtensions(args.SourceIPGroups),
		DestIpGroups:           idsToIDNameExtensions(args.DestIpGroups),
		ProxyGateways:          idsToIDNameExtensions(args.ProxyGateways),
		WorkloadGroups:         expandWorkloadGroups(args.WorkloadGroups),
	}
	return api
}

func sslInspectionRuleAPIToState(api *sslinspection.SSLInspectionRules) SslInspectionRuleState {
	state := SslInspectionRuleState{
		SslInspectionRuleArgs: SslInspectionRuleArgs{
			Name:                   api.Name,
			Order:                  api.Order,
			Description:            stringPtr(api.Description),
			Rank:                   intPtr(api.Rank),
			State:                  stringPtr(api.State),
			RoadWarriorForKerberos: boolPtr(api.RoadWarriorForKerberos),
			Action:                 sslInspectionActionFromAPI(api.Action),
			CloudApplications:      api.CloudApplications,
			URLCategories:          api.URLCategories,
			DeviceTrustLevels:      api.DeviceTrustLevels,
			Platforms:              api.Platforms,
			UserAgentTypes:         api.UserAgentTypes,
			Locations:              idNameExtensionsToIDs(api.Locations),
			LocationGroups:         idNameExtensionsToIDs(api.LocationGroups),
			Groups:                 idNameExtensionsToIDs(api.Groups),
			Departments:            idNameExtensionsToIDs(api.Departments),
			Users:                  idNameExtensionsToIDs(api.Users),
			TimeWindows:            idNameExtensionsToIDs(api.TimeWindows),
			Labels:                 idNameExtensionsToIDs(api.Labels),
			DeviceGroups:           idNameExtensionsToIDs(api.DeviceGroups),
			Devices:                idNameExtensionsToIDs(api.Devices),
			SourceIPGroups:         idNameExtensionsToIDs(api.SourceIPGroups),
			DestIpGroups:           idNameExtensionsToIDs(api.DestIpGroups),
			ProxyGateways:          idNameExtensionsToIDs(api.ProxyGateways),
			WorkloadGroups:         workloadGroupOutputsToInputs(flattenWorkloadGroups(api.WorkloadGroups)),
		},
		RuleID: intPtr(api.ID),
	}
	return state
}

func (SslInspectionRule) Diff(ctx context.Context, req infer.DiffRequest[SslInspectionRuleArgs, SslInspectionRuleState]) (infer.DiffResponse, error) {
	diff, hasChanges := diffSkippingUnsetInputs(req.State.SslInspectionRuleArgs, req.Inputs)
	return infer.DiffResponse{HasChanges: hasChanges, DetailedDiff: diff}, nil
}

func (SslInspectionRule) Create(ctx context.Context, req infer.CreateRequest[SslInspectionRuleArgs]) (infer.CreateResponse[SslInspectionRuleState], error) {
	if req.DryRun {
		s := SslInspectionRuleState{SslInspectionRuleArgs: req.Inputs, RuleID: intPtr(0)}
		return infer.CreateResponse[SslInspectionRuleState]{ID: "preview", Output: s}, nil
	}
	if req.Inputs.Order < 1 {
		return infer.CreateResponse[SslInspectionRuleState]{}, fmt.Errorf("order must be a positive whole number (>= 1), got %d", req.Inputs.Order)
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[SslInspectionRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := sslInspectionRuleArgsToAPI(&req.Inputs, 0)

	sslInspectionLock.Lock()
	if sslInspectionStartingOrder == 0 {
		list, _ := sslinspection.GetAll(ctx, svc)
		for _, r := range list {
			if r.Order > sslInspectionStartingOrder {
				sslInspectionStartingOrder = r.Order
			}
		}
		if sslInspectionStartingOrder == 0 {
			sslInspectionStartingOrder = 1
		}
	}
	sslInspectionLock.Unlock()

	intendedOrder := apiReq.Order
	intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)
	if intendedRank < 7 {
		apiReq.Rank = 7
	}
	apiReq.Order = sslInspectionStartingOrder

	resp, err := sslinspection.Create(ctx, svc, &apiReq)
	if customErr := failFastOnErrorCodes(err); customErr != nil {
		return infer.CreateResponse[SslInspectionRuleState]{}, customErr
	}
	if err != nil {
		reg := regexp.MustCompile("Rule with rank [0-9]+ is not allowed at order [0-9]+")
		if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") && reg.MatchString(err.Error()) {
			return infer.CreateResponse[SslInspectionRuleState]{}, fmt.Errorf("error creating ssl inspection rule: %s, check order %d vs rank %d, err:%s",
				apiReq.Name, intendedOrder, intendedRank, err)
		}
		return infer.CreateResponse[SslInspectionRuleState]{}, fmt.Errorf("creating ssl inspection rule: %w", err)
	}

	reorderWithBeforeReorder(
		OrderRule{Order: intendedOrder, Rank: intendedRank},
		resp.ID,
		sslInspectionResourceType,
		func() (int, error) {
			allRules, err := sslinspection.GetAll(ctx, svc)
			if err != nil {
				return 0, err
			}
			return len(allRules), nil
		},
		func(id int, order OrderRule) error {
			rule, err := sslinspection.Get(ctx, svc, id)
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
			_, err = sslinspection.Update(ctx, svc, id, rule)
			return err
		},
		nil,
	)

	markOrderRuleAsDone(resp.ID, sslInspectionResourceType)
	waitForReorder(sslInspectionResourceType)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.CreateResponse[SslInspectionRuleState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	rule, err := sslinspection.Get(ctx, svc, resp.ID)
	if err != nil {
		state := SslInspectionRuleState{SslInspectionRuleArgs: req.Inputs, RuleID: intPtr(resp.ID)}
		return infer.CreateResponse[SslInspectionRuleState]{ID: strconv.Itoa(resp.ID), Output: state}, nil
	}
	return infer.CreateResponse[SslInspectionRuleState]{
		ID:     strconv.Itoa(resp.ID),
		Output: sslInspectionRuleAPIToState(rule),
	}, nil
}

func (SslInspectionRule) Read(ctx context.Context, req infer.ReadRequest[SslInspectionRuleArgs, SslInspectionRuleState]) (infer.ReadResponse[SslInspectionRuleArgs, SslInspectionRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[SslInspectionRuleArgs, SslInspectionRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		rule, lookupErr := sslinspection.GetByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[SslInspectionRuleArgs, SslInspectionRuleState]{}, fmt.Errorf("ssl inspection rule not found")
			}
			return infer.ReadResponse[SslInspectionRuleArgs, SslInspectionRuleState]{}, lookupErr
		}
		id = rule.ID
	}

	rule, err := sslinspection.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[SslInspectionRuleArgs, SslInspectionRuleState]{}, fmt.Errorf("ssl inspection rule not found")
		}
		return infer.ReadResponse[SslInspectionRuleArgs, SslInspectionRuleState]{}, err
	}
	state := sslInspectionRuleAPIToState(rule)
	return infer.ReadResponse[SslInspectionRuleArgs, SslInspectionRuleState]{
		ID:     req.ID,
		Inputs: state.SslInspectionRuleArgs,
		State:  state,
	}, nil
}

func (SslInspectionRule) Update(ctx context.Context, req infer.UpdateRequest[SslInspectionRuleArgs, SslInspectionRuleState]) (infer.UpdateResponse[SslInspectionRuleState], error) {
	if req.Inputs.Order < 1 {
		return infer.UpdateResponse[SslInspectionRuleState]{}, fmt.Errorf("order must be a positive whole number (>= 1), got %d", req.Inputs.Order)
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[SslInspectionRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[SslInspectionRuleState]{}, fmt.Errorf("invalid ssl inspection rule ID: %s", req.ID)
	}
	apiReq := sslInspectionRuleArgsToAPI(&req.Inputs, id)

	existingRules, err := sslinspection.GetAll(ctx, svc)
	if err != nil {
		return infer.UpdateResponse[SslInspectionRuleState]{}, err
	}
	sort.Slice(existingRules, func(i, j int) bool {
		return existingRules[i].Rank < existingRules[j].Rank || (existingRules[i].Rank == existingRules[j].Rank && existingRules[i].Order < existingRules[j].Order)
	})
	intendedOrder := apiReq.Order
	intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)
	nextAvailableOrder := existingRules[len(existingRules)-1].Order
	apiReq.Rank = 7
	apiReq.Order = nextAvailableOrder

	if _, err = sslinspection.Update(ctx, svc, id, &apiReq); err != nil {
		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.UpdateResponse[SslInspectionRuleState]{}, customErr
		}
		return infer.UpdateResponse[SslInspectionRuleState]{}, err
	}

	reorderWithBeforeReorder(
		OrderRule{Order: intendedOrder, Rank: intendedRank},
		id,
		sslInspectionResourceType,
		func() (int, error) {
			allRules, err := sslinspection.GetAll(ctx, svc)
			if err != nil {
				return 0, err
			}
			return len(allRules), nil
		},
		func(ruleID int, order OrderRule) error {
			rule, err := sslinspection.Get(ctx, svc, ruleID)
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
			_, err = sslinspection.Update(ctx, svc, ruleID, rule)
			return err
		},
		nil,
	)

	markOrderRuleAsDone(id, sslInspectionResourceType)
	waitForReorder(sslInspectionResourceType)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[SslInspectionRuleState]{}, activationErr
		}
	}

	updated, err := sslinspection.Get(ctx, svc, id)
	if err != nil {
		return infer.UpdateResponse[SslInspectionRuleState]{Output: SslInspectionRuleState{
			SslInspectionRuleArgs: req.Inputs,
			RuleID:                intPtr(id),
		}}, nil
	}
	return infer.UpdateResponse[SslInspectionRuleState]{Output: sslInspectionRuleAPIToState(updated)}, nil
}

func (SslInspectionRule) Delete(ctx context.Context, req infer.DeleteRequest[SslInspectionRuleState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid ssl inspection rule ID: %s", req.ID)
	}
	if _, err := sslinspection.Delete(ctx, svc, id); err != nil {
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

func (SslInspectionRule) Annotate(a infer.Annotator) {
	describeResource(a, &SslInspectionRule{}, `The zia_ssl_inspection_rules resource manages SSL inspection rules in the Zscaler Internet Access (ZIA) cloud service. SSL inspection rules determine whether to decrypt, not decrypt, or block SSL/TLS traffic based on criteria such as locations, departments, groups, users, URL categories, cloud applications, and platforms.

For more information, see the [ZIA SSL Inspection documentation](https://help.zscaler.com/zia/about-ssl-inspection-policies).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic SSL Inspection Rule

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.SslInspectionRule("example", {
    name: "Example SSL Inspection Rule",
    description: "Decrypt corporate traffic",
    order: 1,
    state: "ENABLED",
    action: {
        type: "DECRYPT",
        showEun: false,
        decryptSubActions: {
            serverCertificates: "ALLOW",
            ocspCheck: true,
            http2Enabled: true,
        },
    },
    urlCategories: ["ANY"],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.SslInspectionRule("example",
    name="Example SSL Inspection Rule",
    description="Decrypt corporate traffic",
    order=1,
    state="ENABLED",
    action={
        "type": "DECRYPT",
        "show_eun": False,
        "decrypt_sub_actions": {
            "server_certificates": "ALLOW",
            "ocsp_check": True,
            "http2_enabled": True,
        },
    },
    url_categories=["ANY"],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:SslInspectionRule
    properties:
      name: Example SSL Inspection Rule
      description: Decrypt corporate traffic
      order: 1
      state: ENABLED
      action:
        type: DECRYPT
        showEun: false
        decryptSubActions:
          serverCertificates: ALLOW
          ocspCheck: true
          http2Enabled: true
      urlCategories:
        - ANY
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing SSL Inspection Rule can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:SslInspectionRule example 12345
`+tripleBacktick("")+`
`)
}

func (a *SslInspectionRuleArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the SSL inspection rule. Must be unique.")
	ann.Describe(&a.Order, "The order of execution of the rule with respect to other SSL inspection rules.")
	ann.Describe(&a.Description, "Additional information about the SSL inspection rule.")
	ann.Describe(&a.Rank, "Admin rank of the SSL inspection policy rule. Valid values: 0-7. Default: 7.")
	ann.Describe(&a.State, "Rule state. Valid values: `ENABLED`, `DISABLED`.")
	ann.Describe(&a.RoadWarriorForKerberos, "Indicates whether the rule applies to road warrior (remote) users using Kerberos authentication.")
	ann.Describe(&a.Action, "The action configuration for the SSL inspection rule, including decrypt/do-not-decrypt sub-actions.")
	ann.Describe(&a.CloudApplications, "List of cloud application names to which the rule applies.")
	ann.Describe(&a.URLCategories, "List of URL categories to which the rule applies.")
	ann.Describe(&a.DeviceTrustLevels, "Device trust levels for the rule. Valid values: `ANY`, `UNKNOWN_DEVICETRUSTLEVEL`, `LOW_TRUST`, `MEDIUM_TRUST`, `HIGH_TRUST`.")
	ann.Describe(&a.Platforms, "Platforms to which the rule applies (e.g., `SCAN_IOS`, `SCAN_ANDROID`, `SCAN_MACOS`, `SCAN_WINDOWS`, `SCAN_LINUX`).")
	ann.Describe(&a.UserAgentTypes, "User agent types the rule applies to.")
	ann.Describe(&a.Locations, "IDs of locations to which the rule applies.")
	ann.Describe(&a.LocationGroups, "IDs of location groups to which the rule applies.")
	ann.Describe(&a.Groups, "IDs of groups to which the rule applies.")
	ann.Describe(&a.Departments, "IDs of departments to which the rule applies.")
	ann.Describe(&a.Users, "IDs of users to which the rule applies.")
	ann.Describe(&a.TimeWindows, "IDs of time intervals during which the rule must be enforced.")
	ann.Describe(&a.Labels, "IDs of labels associated with the rule.")
	ann.Describe(&a.DeviceGroups, "IDs of device groups to which the rule applies.")
	ann.Describe(&a.Devices, "IDs of devices to which the rule applies.")
	ann.Describe(&a.SourceIPGroups, "IDs of source IP address groups for the rule.")
	ann.Describe(&a.DestIpGroups, "IDs of destination IP address groups for the rule.")
	ann.Describe(&a.ProxyGateways, "IDs of proxy gateway configurations for the rule.")
	ann.Describe(&a.WorkloadGroups, "List of preconfigured workload groups to which the policy must be applied.")
}

func (s *SslInspectionRuleState) Annotate(a infer.Annotator) {
	a.Describe(&s.RuleID, "The system-generated ID of the SSL inspection rule.")
}

var _ infer.CustomResource[SslInspectionRuleArgs, SslInspectionRuleState] = SslInspectionRule{}
