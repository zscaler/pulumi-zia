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

// Package provider implements the Tenant Restriction Profile resource.
// Adopted from terraform-provider-zia resource_zia_tenant_restriction_profile.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/tenancy_restriction"
)

// TenantRestrictionProfile implements the zia:index:TenantRestrictionProfile resource.
type TenantRestrictionProfile struct{}

// TenantRestrictionProfileArgs are the inputs.
type TenantRestrictionProfileArgs struct {
	Name                        string   `pulumi:"name"`
	Description                 *string  `pulumi:"description,optional"`
	AppType                     *string  `pulumi:"appType,optional"`
	ItemTypePrimary             *string  `pulumi:"itemTypePrimary,optional"`
	ItemTypeSecondary           *string  `pulumi:"itemTypeSecondary,optional"`
	RestrictPersonalO365Domains *bool    `pulumi:"restrictPersonalO365Domains,optional"`
	AllowGoogleConsumers        *bool    `pulumi:"allowGoogleConsumers,optional"`
	MsLoginServicesTrV2         *bool    `pulumi:"msLoginServicesTrV2,optional"`
	AllowGoogleVisitors         *bool    `pulumi:"allowGoogleVisitors,optional"`
	AllowGcpCloudStorageRead    *bool    `pulumi:"allowGcpCloudStorageRead,optional"`
	ItemDataPrimary             []string `pulumi:"itemDataPrimary,optional"`
	ItemDataSecondary           []string `pulumi:"itemDataSecondary,optional"`
	ItemValue                   []string `pulumi:"itemValue,optional"`
}

// TenantRestrictionProfileState is the persisted state.
type TenantRestrictionProfileState struct {
	TenantRestrictionProfileArgs
	ProfileId *int `pulumi:"profileId"`
}

func (TenantRestrictionProfile) Create(ctx context.Context, req infer.CreateRequest[TenantRestrictionProfileArgs]) (infer.CreateResponse[TenantRestrictionProfileState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[TenantRestrictionProfileState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := tenantRestrictionProfileArgsToAPI(req.Inputs, 0)
	resp, _, err := tenancy_restriction.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[TenantRestrictionProfileState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[TenantRestrictionProfileState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := tenantRestrictionProfileToState(resp)
	return infer.CreateResponse[TenantRestrictionProfileState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (TenantRestrictionProfile) Read(ctx context.Context, req infer.ReadRequest[TenantRestrictionProfileArgs, TenantRestrictionProfileState]) (infer.ReadResponse[TenantRestrictionProfileArgs, TenantRestrictionProfileState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[TenantRestrictionProfileArgs, TenantRestrictionProfileState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.ReadResponse[TenantRestrictionProfileArgs, TenantRestrictionProfileState]{}, fmt.Errorf("invalid profile id: %w", err)
	}

	resp, err := tenancy_restriction.Get(ctx, service, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[TenantRestrictionProfileArgs, TenantRestrictionProfileState]{}, nil
		}
		return infer.ReadResponse[TenantRestrictionProfileArgs, TenantRestrictionProfileState]{}, err
	}

	state := tenantRestrictionProfileToState(resp)
	return infer.ReadResponse[TenantRestrictionProfileArgs, TenantRestrictionProfileState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: state.TenantRestrictionProfileArgs,
		State:  state,
	}, nil
}

func (TenantRestrictionProfile) Update(ctx context.Context, req infer.UpdateRequest[TenantRestrictionProfileArgs, TenantRestrictionProfileState]) (infer.UpdateResponse[TenantRestrictionProfileState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[TenantRestrictionProfileState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[TenantRestrictionProfileState]{}, fmt.Errorf("invalid profile id: %w", err)
	}

	_, err = tenancy_restriction.Get(ctx, service, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.UpdateResponse[TenantRestrictionProfileState]{}, nil
		}
		return infer.UpdateResponse[TenantRestrictionProfileState]{}, err
	}

	apiReq := tenantRestrictionProfileArgsToAPI(req.Inputs, id)
	if _, _, err := tenancy_restriction.Update(ctx, service, id, &apiReq); err != nil {
		return infer.UpdateResponse[TenantRestrictionProfileState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[TenantRestrictionProfileState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	updated, err := tenancy_restriction.Get(ctx, service, id)
	if err != nil {
		return infer.UpdateResponse[TenantRestrictionProfileState]{}, err
	}
	state := tenantRestrictionProfileToState(updated)
	return infer.UpdateResponse[TenantRestrictionProfileState]{Output: state}, nil
}

func (TenantRestrictionProfile) Delete(ctx context.Context, req infer.DeleteRequest[TenantRestrictionProfileState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid profile id: %w", err)
	}

	if _, err := tenancy_restriction.Delete(ctx, service, id); err != nil {
		return infer.DeleteResponse{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	return infer.DeleteResponse{}, nil
}

func (TenantRestrictionProfile) Annotate(a infer.Annotator) {
	describeResource(a, &TenantRestrictionProfile{}, `The zia.TenantRestrictionProfile resource manages tenant restriction profiles in the
Zscaler Internet Access (ZIA) cloud. Tenant restriction profiles control access to cloud
application tenants (e.g., Microsoft 365, Google Workspace) by restricting users to
authorized tenant domains.

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Tenant Restriction Profile

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.TenantRestrictionProfile("example", {
    name: "Example Tenant Profile",
    description: "Managed by Pulumi",
    appType: "MICROSOFT",
    itemTypePrimary: "TENANT_ID",
    itemDataPrimary: ["tenant-id-12345"],
    restrictPersonalO365Domains: true,
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.TenantRestrictionProfile("example",
    name="Example Tenant Profile",
    description="Managed by Pulumi",
    app_type="MICROSOFT",
    item_type_primary="TENANT_ID",
    item_data_primary=["tenant-id-12345"],
    restrict_personal_o365_domains=True,
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:TenantRestrictionProfile
    properties:
      name: Example Tenant Profile
      description: Managed by Pulumi
      appType: MICROSOFT
      itemTypePrimary: TENANT_ID
      itemDataPrimary:
        - tenant-id-12345
      restrictPersonalO365Domains: true
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing tenant restriction profile can be imported using its ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:TenantRestrictionProfile example 12345
`+tripleBacktick("")+`
`)
}

func (a *TenantRestrictionProfileArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "Name of the tenant restriction profile.")
	ann.Describe(&a.Description, "Description of the tenant restriction profile.")
	ann.Describe(&a.AppType, "The cloud application type (e.g., 'MICROSOFT', 'GOOGLE').")
	ann.Describe(&a.ItemTypePrimary, "The primary item type (e.g., 'TENANT_ID').")
	ann.Describe(&a.ItemTypeSecondary, "The secondary item type.")
	ann.Describe(&a.RestrictPersonalO365Domains, "Whether to restrict personal Office 365 domains.")
	ann.Describe(&a.AllowGoogleConsumers, "Whether to allow Google consumer accounts.")
	ann.Describe(&a.MsLoginServicesTrV2, "Whether to enable Microsoft login services tenant restriction v2.")
	ann.Describe(&a.AllowGoogleVisitors, "Whether to allow Google visitor accounts.")
	ann.Describe(&a.AllowGcpCloudStorageRead, "Whether to allow GCP Cloud Storage read access.")
	ann.Describe(&a.ItemDataPrimary, "List of primary item data values (e.g., tenant IDs).")
	ann.Describe(&a.ItemDataSecondary, "List of secondary item data values.")
	ann.Describe(&a.ItemValue, "List of item values.")
}

func (s *TenantRestrictionProfileState) Annotate(a infer.Annotator) {
	a.Describe(&s.ProfileId, "The unique identifier for the tenant restriction profile assigned by the ZIA cloud.")
}

func tenantRestrictionProfileArgsToAPI(args TenantRestrictionProfileArgs, id int) tenancy_restriction.TenancyRestrictionProfile {
	return tenancy_restriction.TenancyRestrictionProfile{
		ID:                          id,
		Name:                        args.Name,
		Description:                 ptrToString(args.Description),
		AppType:                     ptrToString(args.AppType),
		ItemTypePrimary:             ptrToString(args.ItemTypePrimary),
		ItemTypeSecondary:           ptrToString(args.ItemTypeSecondary),
		RestrictPersonalO365Domains: ptrToBool(args.RestrictPersonalO365Domains),
		AllowGoogleConsumers:        ptrToBool(args.AllowGoogleConsumers),
		MsLoginServicesTrV2:         ptrToBool(args.MsLoginServicesTrV2),
		AllowGoogleVisitors:         ptrToBool(args.AllowGoogleVisitors),
		AllowGcpCloudStorageRead:    ptrToBool(args.AllowGcpCloudStorageRead),
		ItemDataPrimary:             args.ItemDataPrimary,
		ItemDataSecondary:           args.ItemDataSecondary,
		ItemValue:                   args.ItemValue,
	}
}

func tenantRestrictionProfileToState(r *tenancy_restriction.TenancyRestrictionProfile) TenantRestrictionProfileState {
	return TenantRestrictionProfileState{
		TenantRestrictionProfileArgs: TenantRestrictionProfileArgs{
			Name:                        r.Name,
			Description:                 stringPtr(r.Description),
			AppType:                     stringPtr(r.AppType),
			ItemTypePrimary:             stringPtr(r.ItemTypePrimary),
			ItemTypeSecondary:           stringPtr(r.ItemTypeSecondary),
			RestrictPersonalO365Domains: boolPtr(r.RestrictPersonalO365Domains),
			AllowGoogleConsumers:        boolPtr(r.AllowGoogleConsumers),
			MsLoginServicesTrV2:         boolPtr(r.MsLoginServicesTrV2),
			AllowGoogleVisitors:         boolPtr(r.AllowGoogleVisitors),
			AllowGcpCloudStorageRead:    boolPtr(r.AllowGcpCloudStorageRead),
			ItemDataPrimary:             r.ItemDataPrimary,
			ItemDataSecondary:           r.ItemDataSecondary,
			ItemValue:                   r.ItemValue,
		},
		ProfileId: intPtr(r.ID),
	}
}

var _ infer.CustomResource[TenantRestrictionProfileArgs, TenantRestrictionProfileState] = TenantRestrictionProfile{}
