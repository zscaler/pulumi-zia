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

// Package provider implements the AdminUsers resource.
// Adopted from terraform-provider-zia resource_zia_admin_users.go.

package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/adminuserrolemgmt/admins"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/common"
)

// AdminUsers implements the zia:index:AdminUsers resource.
type AdminUsers struct{}

// AdminUserRoleInput is the role block (single role with id).
type AdminUserRoleInput struct {
	Id *int `pulumi:"id,optional"`
}

// AdminUsersArgs are the inputs for AdminUsers.
type AdminUsersArgs struct {
	LoginName                   string              `pulumi:"loginName"`
	Username                    string              `pulumi:"username"`
	Email                       string              `pulumi:"email"`
	Role                        *AdminUserRoleInput `pulumi:"role,optional"`
	Comments                    *string             `pulumi:"comments,optional"`
	AdminScopeType              *string             `pulumi:"adminScopeType,optional"`
	AdminScopeEntities          []int               `pulumi:"adminScopeEntities,optional"`
	IsNonEditable               *bool               `pulumi:"isNonEditable,optional"`
	Disabled                    *bool               `pulumi:"disabled,optional"`
	IsAuditor                   *bool               `pulumi:"isAuditor,optional"`
	Password                    *string             `pulumi:"password,optional" provider:"secret"`
	IsPasswordLoginAllowed      *bool               `pulumi:"isPasswordLoginAllowed,optional"`
	IsSecurityReportCommEnabled *bool               `pulumi:"isSecurityReportCommEnabled,optional"`
	IsServiceUpdateCommEnabled  *bool               `pulumi:"isServiceUpdateCommEnabled,optional"`
	IsProductUpdateCommEnabled  *bool               `pulumi:"isProductUpdateCommEnabled,optional"`
	IsPasswordExpired           *bool               `pulumi:"isPasswordExpired,optional"`
	IsExecMobileAppEnabled      *bool               `pulumi:"isExecMobileAppEnabled,optional"`
}

// AdminUsersState is the persisted state.
type AdminUsersState struct {
	AdminUsersArgs
	AdminId *int `pulumi:"adminId"`
}

func (AdminUsers) Create(ctx context.Context, req infer.CreateRequest[AdminUsersArgs]) (infer.CreateResponse[AdminUsersState], error) {
	if req.DryRun {
		return infer.CreateResponse[AdminUsersState]{
			ID: "preview",
			Output: AdminUsersState{
				AdminUsersArgs: req.Inputs,
				AdminId:        intPtr(0),
			},
		}, nil
	}

	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[AdminUsersState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := adminUsersArgsToAPI(req.Inputs, 0)
	if err := checkAdminUserPasswordAllowed(apiReq); err != nil {
		return infer.CreateResponse[AdminUsersState]{}, err
	}
	if err := checkAdminUserScopeType(apiReq); err != nil {
		return infer.CreateResponse[AdminUsersState]{}, err
	}

	resp, err := admins.CreateAdminUser(ctx, svc, apiReq)
	if err != nil {
		return infer.CreateResponse[AdminUsersState]{}, fmt.Errorf("creating admin user: %w", err)
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.CreateResponse[AdminUsersState]{}, activationErr
		}
	}

	// Fetch full state (password never returned from API)
	rule, err := admins.GetAdminUsers(ctx, svc, resp.ID)
	if err != nil {
		return infer.CreateResponse[AdminUsersState]{
			ID:     strconv.Itoa(resp.ID),
			Output: adminUsersAPIToState(resp),
		}, nil
	}
	state := adminUsersAPIToState(rule)
	state.Password = req.Inputs.Password // Preserve from inputs for state (never returned from API)
	return infer.CreateResponse[AdminUsersState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (AdminUsers) Read(ctx context.Context, req infer.ReadRequest[AdminUsersArgs, AdminUsersState]) (infer.ReadResponse[AdminUsersArgs, AdminUsersState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[AdminUsersArgs, AdminUsersState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		resp, lookupErr := admins.GetAdminUsersByLoginName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[AdminUsersArgs, AdminUsersState]{}, nil
			}
			return infer.ReadResponse[AdminUsersArgs, AdminUsersState]{}, lookupErr
		}
		id = resp.ID
	}

	resp, err := admins.GetAdminUsers(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[AdminUsersArgs, AdminUsersState]{}, nil
		}
		return infer.ReadResponse[AdminUsersArgs, AdminUsersState]{}, err
	}

	state := adminUsersAPIToState(resp)
	args := adminUsersStateToArgs(state)
	return infer.ReadResponse[AdminUsersArgs, AdminUsersState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (AdminUsers) Update(ctx context.Context, req infer.UpdateRequest[AdminUsersArgs, AdminUsersState]) (infer.UpdateResponse[AdminUsersState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[AdminUsersState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[AdminUsersState]{}, fmt.Errorf("invalid admin user id: %w", err)
	}

	apiReq := adminUsersArgsToAPI(req.Inputs, id)
	if err := checkAdminUserScopeType(apiReq); err != nil {
		return infer.UpdateResponse[AdminUsersState]{}, err
	}

	if _, err := admins.UpdateAdminUser(ctx, svc, id, apiReq); err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.UpdateResponse[AdminUsersState]{}, nil
		}
		return infer.UpdateResponse[AdminUsersState]{}, fmt.Errorf("updating admin user: %w", err)
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[AdminUsersState]{}, activationErr
		}
	}

	rule, err := admins.GetAdminUsers(ctx, svc, id)
	if err != nil {
		return infer.UpdateResponse[AdminUsersState]{Output: adminUsersArgsToState(req.Inputs, &id)}, nil
	}
	state := adminUsersAPIToState(rule)
	state.Password = req.Inputs.Password
	return infer.UpdateResponse[AdminUsersState]{Output: state}, nil
}

func (AdminUsers) Delete(ctx context.Context, req infer.DeleteRequest[AdminUsersState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid admin user id: %w", err)
	}

	if _, err := admins.DeleteAdminUser(ctx, svc, id); err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.DeleteResponse{}, nil
		}
		return infer.DeleteResponse{}, fmt.Errorf("deleting admin user: %w", err)
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}
	return infer.DeleteResponse{}, nil
}

func (AdminUsers) Annotate(a infer.Annotator) {
	describeResource(a, &AdminUsers{}, `The zia_admin_users resource manages administrator users in the Zscaler Internet Access (ZIA) cloud service. Administrator users have access to the ZIA Admin Portal and can manage policies, configurations, and other administrative tasks based on their assigned role.

For more information, see the [ZIA Admin User Management documentation](https://help.zscaler.com/zia/admin-user-management).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Admin User

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

import * as pulumi from "@pulumi/pulumi";

const cfg = new pulumi.Config();
const adminPassword = cfg.requireSecret("adminPassword");

const example = new zia.AdminUsers("example", {
    loginName: "admin@example.com",
    username: "Example Admin",
    email: "admin@example.com",
    password: adminPassword,
    isPasswordLoginAllowed: true,
    role: { id: 12345 },
    adminScopeType: "ORGANIZATION",
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

cfg = pulumi.Config()
admin_password = cfg.require_secret("adminPassword")

example = zia.AdminUsers("example",
    login_name="admin@example.com",
    username="Example Admin",
    email="admin@example.com",
    password=admin_password,
    is_password_login_allowed=True,
    role={"id": 12345},
    admin_scope_type="ORGANIZATION",
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:AdminUsers
    properties:
      loginName: admin@example.com
      username: Example Admin
      email: admin@example.com
      password:
        fn::secret: ${adminPassword}
      isPasswordLoginAllowed: true
      role:
        id: 12345
      adminScopeType: ORGANIZATION
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing admin user can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:AdminUsers example 12345
`+tripleBacktick("")+`
`)
}

func (a *AdminUsersArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.LoginName, "The admin user's login name (email format). Must be unique.")
	ann.Describe(&a.Username, "The admin user's display name.")
	ann.Describe(&a.Email, "The admin user's email address.")
	ann.Describe(&a.Role, "The role assigned to the admin user. Provide the role ID.")
	ann.Describe(&a.Comments, "Additional information about the admin user.")
	ann.Describe(&a.AdminScopeType, "The admin scope type. Valid values: `ORGANIZATION`, `DEPARTMENT`, `LOCATION`, `LOCATION_GROUP`.")
	ann.Describe(&a.AdminScopeEntities, "IDs of the admin scope entities (departments, locations, or location groups) when adminScopeType is not `ORGANIZATION`.")
	ann.Describe(&a.IsNonEditable, "Indicates whether the admin user is non-editable (read-only in the ZIA Admin Portal).")
	ann.Describe(&a.Disabled, "Whether the admin account is disabled.")
	ann.Describe(&a.IsAuditor, "Indicates whether the admin is an auditor.")
	ann.Describe(&a.Password, "The admin user's password. Required when isPasswordLoginAllowed is true. Must be 8 to 100 characters.")
	ann.Describe(&a.IsPasswordLoginAllowed, "Whether password-based login is allowed for the admin user.")
	ann.Describe(&a.IsSecurityReportCommEnabled, "Whether the admin can receive security report communications.")
	ann.Describe(&a.IsServiceUpdateCommEnabled, "Whether the admin can receive service update communications.")
	ann.Describe(&a.IsProductUpdateCommEnabled, "Whether the admin can receive product update communications.")
	ann.Describe(&a.IsPasswordExpired, "Indicates whether the admin user's password has expired.")
	ann.Describe(&a.IsExecMobileAppEnabled, "Whether Executive Insights App access is enabled. Can only be set when adminScopeType is `ORGANIZATION`.")
}

func (s *AdminUsersState) Annotate(a infer.Annotator) {
	a.Describe(&s.AdminId, "The system-generated ID of the admin user.")
}

var _ infer.CustomResource[AdminUsersArgs, AdminUsersState] = AdminUsers{}

func adminUsersArgsToAPI(in AdminUsersArgs, existingID int) admins.AdminUsers {
	result := admins.AdminUsers{
		ID:                          existingID,
		LoginName:                   in.LoginName,
		UserName:                    in.Username,
		Email:                       in.Email,
		Comments:                    ptrToString(in.Comments),
		IsNonEditable:               ptrToBool(in.IsNonEditable),
		Disabled:                    ptrToBool(in.Disabled),
		IsAuditor:                   ptrToBool(in.IsAuditor),
		Password:                    ptrToString(in.Password),
		AdminScopeType:              ptrToString(in.AdminScopeType),
		IsPasswordLoginAllowed:      ptrToBool(in.IsPasswordLoginAllowed),
		IsSecurityReportCommEnabled: ptrToBool(in.IsSecurityReportCommEnabled),
		IsServiceUpdateCommEnabled:  ptrToBool(in.IsServiceUpdateCommEnabled),
		IsProductUpdateCommEnabled:  ptrToBool(in.IsProductUpdateCommEnabled),
		IsPasswordExpired:           ptrToBool(in.IsPasswordExpired),
		IsExecMobileAppEnabled:      ptrToBool(in.IsExecMobileAppEnabled),
		AdminScopeEntities:          idsToIDNameExtensions(in.AdminScopeEntities),
	}
	if in.Role != nil && in.Role.Id != nil && *in.Role.Id != 0 {
		result.Role = &admins.Role{ID: *in.Role.Id}
	}
	return result
}

func adminUsersAPIToState(r *admins.AdminUsers) AdminUsersState {
	args := AdminUsersArgs{
		LoginName:                   r.LoginName,
		Username:                    r.UserName,
		Email:                       r.Email,
		Comments:                    stringPtr(r.Comments),
		AdminScopeType:              stringPtr(r.AdminScopeType),
		AdminScopeEntities:          adminScopeEntitiesToIDs(r.AdminScopeEntities),
		IsNonEditable:               boolPtr(r.IsNonEditable),
		Disabled:                    boolPtr(r.Disabled),
		IsAuditor:                   boolPtr(r.IsAuditor),
		Password:                    nil, // never returned from API
		IsPasswordLoginAllowed:      boolPtr(r.IsPasswordLoginAllowed),
		IsSecurityReportCommEnabled: boolPtr(r.IsSecurityReportCommEnabled),
		IsServiceUpdateCommEnabled:  boolPtr(r.IsServiceUpdateCommEnabled),
		IsProductUpdateCommEnabled:  boolPtr(r.IsProductUpdateCommEnabled),
		IsPasswordExpired:           boolPtr(r.IsPasswordExpired),
		IsExecMobileAppEnabled:      boolPtr(r.IsExecMobileAppEnabled),
	}
	if r.Role != nil {
		args.Role = &AdminUserRoleInput{Id: intPtr(r.Role.ID)}
	}
	return AdminUsersState{
		AdminUsersArgs: args,
		AdminId:        &r.ID,
	}
}

func adminScopeEntitiesToIDs(list []common.IDNameExtensions) []int {
	return idNameExtensionsToIDs(list)
}

func adminUsersStateToArgs(s AdminUsersState) AdminUsersArgs {
	return s.AdminUsersArgs
}

func adminUsersArgsToState(in AdminUsersArgs, adminId *int) AdminUsersState {
	return AdminUsersState{
		AdminUsersArgs: in,
		AdminId:        adminId,
	}
}

func checkAdminUserPasswordAllowed(p admins.AdminUsers) error {
	if p.IsPasswordLoginAllowed && p.Password == "" {
		return fmt.Errorf("enter a password for the admin when isPasswordLoginAllowed is true; password must be 8 to 100 characters")
	}
	return nil
}

func checkAdminUserScopeType(p admins.AdminUsers) error {
	if p.IsExecMobileAppEnabled && p.AdminScopeType != "ORGANIZATION" {
		return fmt.Errorf("mobile app access can only be enabled for an admin with organization scope")
	}
	return nil
}
