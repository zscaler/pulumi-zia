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

// Package provider implements the AdminRoles resource.
// Adopted from terraform-provider-zia resource_zia_admin_roles.go.

package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/adminuserrolemgmt/roles"
)

// AdminRoles implements the zia:index:AdminRoles resource.
type AdminRoles struct{}

// AdminRolesArgs are the inputs for AdminRoles.
type AdminRolesArgs struct {
	Name                  *string           `pulumi:"name,optional"`
	Rank                  *int              `pulumi:"rank,optional"`
	PolicyAccess          *string           `pulumi:"policyAccess,optional"`
	AlertingAccess        *string           `pulumi:"alertingAccess,optional"`
	DashboardAccess       *string           `pulumi:"dashboardAccess,optional"`
	ReportAccess          *string           `pulumi:"reportAccess,optional"`
	AnalysisAccess        *string           `pulumi:"analysisAccess,optional"`
	UsernameAccess        *string           `pulumi:"usernameAccess,optional"`
	DeviceInfoAccess      *string           `pulumi:"deviceInfoAccess,optional"`
	AdminAcctAccess       *string           `pulumi:"adminAcctAccess,optional"`
	IsAuditor             *bool             `pulumi:"isAuditor,optional"`
	FeaturePermissions    map[string]string `pulumi:"featurePermissions,optional"`
	ExtFeaturePermissions map[string]string `pulumi:"extFeaturePermissions,optional"`
	IsNonEditable         *bool             `pulumi:"isNonEditable,optional"`
	LogsLimit             *string           `pulumi:"logsLimit,optional"`
	RoleType              *string           `pulumi:"roleType,optional"`
	ReportTimeDuration    *int              `pulumi:"reportTimeDuration,optional"`
	Permissions           []string          `pulumi:"permissions,optional"`
}

// AdminRolesState is the persisted state.
type AdminRolesState struct {
	AdminRolesArgs
	RoleId *int `pulumi:"roleId"`
}

func (AdminRoles) Create(ctx context.Context, req infer.CreateRequest[AdminRolesArgs]) (infer.CreateResponse[AdminRolesState], error) {
	if req.DryRun {
		return infer.CreateResponse[AdminRolesState]{
			ID: "preview",
			Output: AdminRolesState{
				AdminRolesArgs: req.Inputs,
				RoleId:         intPtr(0),
			},
		}, nil
	}

	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[AdminRolesState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := adminRolesArgsToAPI(req.Inputs, 0)
	resp, _, err := roles.Create(ctx, svc, &apiReq)
	if err != nil {
		return infer.CreateResponse[AdminRolesState]{}, fmt.Errorf("creating admin role: %w", err)
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.CreateResponse[AdminRolesState]{}, activationErr
		}
	}

	state := adminRolesAPIToState(resp)
	return infer.CreateResponse[AdminRolesState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (AdminRoles) Read(ctx context.Context, req infer.ReadRequest[AdminRolesArgs, AdminRolesState]) (infer.ReadResponse[AdminRolesArgs, AdminRolesState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[AdminRolesArgs, AdminRolesState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		resp, lookupErr := roles.GetByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[AdminRolesArgs, AdminRolesState]{}, nil
			}
			return infer.ReadResponse[AdminRolesArgs, AdminRolesState]{}, lookupErr
		}
		id = resp.ID
	}

	resp, err := roles.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[AdminRolesArgs, AdminRolesState]{}, nil
		}
		return infer.ReadResponse[AdminRolesArgs, AdminRolesState]{}, err
	}

	state := adminRolesAPIToState(resp)
	args := adminRolesStateToArgs(state)
	return infer.ReadResponse[AdminRolesArgs, AdminRolesState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (AdminRoles) Update(ctx context.Context, req infer.UpdateRequest[AdminRolesArgs, AdminRolesState]) (infer.UpdateResponse[AdminRolesState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[AdminRolesState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[AdminRolesState]{}, fmt.Errorf("invalid admin role id: %w", err)
	}

	apiReq := adminRolesArgsToAPI(req.Inputs, id)
	if _, _, err := roles.Update(ctx, svc, id, &apiReq); err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.UpdateResponse[AdminRolesState]{}, nil
		}
		return infer.UpdateResponse[AdminRolesState]{}, fmt.Errorf("updating admin role: %w", err)
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[AdminRolesState]{}, activationErr
		}
	}

	resp, err := roles.Get(ctx, svc, id)
	if err != nil {
		return infer.UpdateResponse[AdminRolesState]{Output: adminRolesArgsToState(req.Inputs, &id)}, nil
	}
	return infer.UpdateResponse[AdminRolesState]{Output: adminRolesAPIToState(resp)}, nil
}

func (AdminRoles) Delete(ctx context.Context, req infer.DeleteRequest[AdminRolesState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid admin role id: %w", err)
	}

	if _, err := roles.Delete(ctx, svc, id); err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.DeleteResponse{}, nil
		}
		return infer.DeleteResponse{}, fmt.Errorf("deleting admin role: %w", err)
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}
	return infer.DeleteResponse{}, nil
}

func (AdminRoles) Annotate(a infer.Annotator) {
	describeResource(a, &AdminRoles{}, `The zia_admin_roles resource manages administrator roles in the Zscaler Internet Access (ZIA) cloud service. Admin roles define the permissions and access levels for administrator users.

For more information, see the [ZIA Admin Role Management documentation](https://help.zscaler.com/zia/admin-role-management).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Admin Role

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.AdminRoles("example", {
    name: "Example Role",
    rank: 7,
    policyAccess: "READ_WRITE",
    dashboardAccess: "READ_ONLY",
    reportAccess: "READ_ONLY",
    alertingAccess: "READ_ONLY",
    usernameAccess: "READ_ONLY",
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.AdminRoles("example",
    name="Example Role",
    rank=7,
    policy_access="READ_WRITE",
    dashboard_access="READ_ONLY",
    report_access="READ_ONLY",
    alerting_access="READ_ONLY",
    username_access="READ_ONLY",
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:AdminRoles
    properties:
      name: Example Role
      rank: 7
      policyAccess: READ_WRITE
      dashboardAccess: READ_ONLY
      reportAccess: READ_ONLY
      alertingAccess: READ_ONLY
      usernameAccess: READ_ONLY
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing admin role can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:AdminRoles example 12345
`+tripleBacktick("")+`
`)
}

func (a *AdminRolesArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the admin role.")
	ann.Describe(&a.Rank, "Admin rank of the role. Default: 7. Valid values: 0-7.")
	ann.Describe(&a.PolicyAccess, "Policy access permission. Valid values: `NONE`, `READ_ONLY`, `READ_WRITE`.")
	ann.Describe(&a.AlertingAccess, "Alerting access permission. Valid values: `NONE`, `READ_ONLY`, `READ_WRITE`.")
	ann.Describe(&a.DashboardAccess, "Dashboard access permission. Valid values: `NONE`, `READ_ONLY`.")
	ann.Describe(&a.ReportAccess, "Report access permission. Valid values: `NONE`, `READ_ONLY`.")
	ann.Describe(&a.AnalysisAccess, "Insights logs access permission. Valid values: `NONE`, `READ_ONLY`, `READ_WRITE`.")
	ann.Describe(&a.UsernameAccess, "Username access permission. Valid values: `NONE`, `READ_ONLY`.")
	ann.Describe(&a.DeviceInfoAccess, "Device info access permission. Valid values: `NONE`, `READ_ONLY`.")
	ann.Describe(&a.AdminAcctAccess, "Admin and role management access permission. Valid values: `NONE`, `READ_ONLY`, `READ_WRITE`.")
	ann.Describe(&a.IsAuditor, "Indicates whether this is an auditor role.")
	ann.Describe(&a.FeaturePermissions, "Map of feature permissions to their access levels.")
	ann.Describe(&a.ExtFeaturePermissions, "Map of extended feature permissions to their access levels.")
	ann.Describe(&a.IsNonEditable, "Indicates whether the role is non-editable (built-in system role).")
	ann.Describe(&a.LogsLimit, "Log range limit. Valid values: `UNRESTRICTED`, `LAST_1_HR`, `LAST_2_HRS`, `LAST_6_HRS`, `LAST_24_HRS`, `LAST_1_MONTH`.")
	ann.Describe(&a.RoleType, "The admin role type. Valid values: `EXEC_INSIGHT_AND_ORG_ADMIN`, `ORG_ADMIN`.")
	ann.Describe(&a.ReportTimeDuration, "Report time duration in days.")
	ann.Describe(&a.Permissions, "List of functional areas to which this role has access (e.g., `POLICY`, `DASHBOARD`).")
}

func (s *AdminRolesState) Annotate(a infer.Annotator) {
	a.Describe(&s.RoleId, "The system-generated ID of the admin role.")
}

var _ infer.CustomResource[AdminRolesArgs, AdminRolesState] = AdminRoles{}

func adminRolesArgsToAPI(in AdminRolesArgs, existingID int) roles.AdminRoles {
	featurePerms := make(map[string]interface{})
	for k, v := range in.FeaturePermissions {
		featurePerms[k] = v
	}
	extFeaturePerms := make(map[string]interface{})
	for k, v := range in.ExtFeaturePermissions {
		extFeaturePerms[k] = v
	}
	return roles.AdminRoles{
		ID:                    existingID,
		Name:                  ptrToString(in.Name),
		Rank:                  ptrToIntDefault(in.Rank, 7),
		PolicyAccess:          ptrToString(in.PolicyAccess),
		AlertingAccess:        ptrToString(in.AlertingAccess),
		DashboardAccess:       ptrToString(in.DashboardAccess),
		ReportAccess:          ptrToString(in.ReportAccess),
		AnalysisAccess:        ptrToString(in.AnalysisAccess),
		UsernameAccess:        ptrToString(in.UsernameAccess),
		DeviceInfoAccess:      ptrToString(in.DeviceInfoAccess),
		AdminAcctAccess:       ptrToString(in.AdminAcctAccess),
		IsAuditor:             ptrToBool(in.IsAuditor),
		IsNonEditable:         ptrToBool(in.IsNonEditable),
		LogsLimit:             ptrToString(in.LogsLimit),
		RoleType:              ptrToString(in.RoleType),
		ReportTimeDuration:    ptrToIntDefault(in.ReportTimeDuration, 0),
		Permissions:           in.Permissions,
		FeaturePermissions:    featurePerms,
		ExtFeaturePermissions: extFeaturePerms,
	}
}

func adminRolesAPIToState(r *roles.AdminRoles) AdminRolesState {
	featurePerms := make(map[string]string)
	if r.FeaturePermissions != nil {
		for k, v := range r.FeaturePermissions {
			if s, ok := v.(string); ok {
				featurePerms[k] = s
			}
		}
	}
	extFeaturePerms := make(map[string]string)
	if r.ExtFeaturePermissions != nil {
		for k, v := range r.ExtFeaturePermissions {
			if s, ok := v.(string); ok {
				extFeaturePerms[k] = s
			}
		}
	}
	return AdminRolesState{
		AdminRolesArgs: AdminRolesArgs{
			Name:                  stringPtr(r.Name),
			Rank:                  intPtr(r.Rank),
			PolicyAccess:          stringPtr(r.PolicyAccess),
			AlertingAccess:        stringPtr(r.AlertingAccess),
			DashboardAccess:       stringPtr(r.DashboardAccess),
			ReportAccess:          stringPtr(r.ReportAccess),
			AnalysisAccess:        stringPtr(r.AnalysisAccess),
			UsernameAccess:        stringPtr(r.UsernameAccess),
			DeviceInfoAccess:      stringPtr(r.DeviceInfoAccess),
			AdminAcctAccess:       stringPtr(r.AdminAcctAccess),
			IsAuditor:             boolPtr(r.IsAuditor),
			FeaturePermissions:    featurePerms,
			ExtFeaturePermissions: extFeaturePerms,
			IsNonEditable:         boolPtr(r.IsNonEditable),
			LogsLimit:             stringPtr(r.LogsLimit),
			RoleType:              stringPtr(r.RoleType),
			ReportTimeDuration:    intPtr(r.ReportTimeDuration),
			Permissions:           r.Permissions,
		},
		RoleId: &r.ID,
	}
}

func adminRolesStateToArgs(s AdminRolesState) AdminRolesArgs {
	return s.AdminRolesArgs
}

func adminRolesArgsToState(in AdminRolesArgs, roleId *int) AdminRolesState {
	return AdminRolesState{
		AdminRolesArgs: in,
		RoleId:         roleId,
	}
}
