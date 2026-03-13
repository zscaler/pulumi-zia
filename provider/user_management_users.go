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

// Package provider implements the User Management User resource and invoke.
// Adopted from terraform-provider-zia resource_zia_user_management_users.go and data_source_zia_user_management_users.go.

package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/common"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/usermanagement/users"
)

// --- Resource ---

// UserDepartmentInput is the department block.
type UserDepartmentInput struct {
	Id       *int    `pulumi:"id,optional"`
	Name     *string `pulumi:"name,optional"`
	IdpId    *int    `pulumi:"idpId,optional"`
	Comments *string `pulumi:"comments,optional"`
	Deleted  *bool   `pulumi:"deleted,optional"`
}

// UserManagementUser implements the zia:index:UserManagementUser resource.
type UserManagementUser struct{}

// UserManagementUserArgs are the inputs for UserManagementUser.
type UserManagementUserArgs struct {
	Name          string               `pulumi:"name"`
	Email         string               `pulumi:"email"`
	Comments      *string              `pulumi:"comments,optional"`
	TempAuthEmail *string              `pulumi:"tempAuthEmail,optional"`
	AuthMethods   []string             `pulumi:"authMethods,optional"`
	Password      string               `pulumi:"password" provider:"secret"`
	Groups        []int                `pulumi:"groups,optional"`
	Department    *UserDepartmentInput `pulumi:"department,optional"`
}

// UserManagementUserState is the persisted state.
type UserManagementUserState struct {
	UserManagementUserArgs
	UserId *int `pulumi:"userId"`
}

func (UserManagementUser) Check(ctx context.Context, req infer.CheckRequest) (infer.CheckResponse[UserManagementUserArgs], error) {
	inputs, failures, err := infer.DefaultCheck[UserManagementUserArgs](ctx, req.NewInputs)
	if err != nil {
		return infer.CheckResponse[UserManagementUserArgs]{}, err
	}
	if len(failures) > 0 {
		return infer.CheckResponse[UserManagementUserArgs]{Failures: failures}, nil
	}
	if len(inputs.Name) > 127 {
		return infer.CheckResponse[UserManagementUserArgs]{Failures: []p.CheckFailure{{
			Property: "name",
			Reason:   "name must be at most 127 characters",
		}}}, nil
	}
	if len(inputs.Email) > 127 {
		return infer.CheckResponse[UserManagementUserArgs]{Failures: []p.CheckFailure{{
			Property: "email",
			Reason:   "email must be at most 127 characters",
		}}}, nil
	}
	if inputs.Comments != nil && len(*inputs.Comments) > 10240 {
		return infer.CheckResponse[UserManagementUserArgs]{Failures: []p.CheckFailure{{
			Property: "comments",
			Reason:   "comments must be at most 10240 characters",
		}}}, nil
	}
	if inputs.Department == nil {
		return infer.CheckResponse[UserManagementUserArgs]{Failures: []p.CheckFailure{{
			Property: "department",
			Reason:   "department is required",
		}}}, nil
	}
	for _, m := range inputs.AuthMethods {
		if m != "BASIC" && m != "DIGEST" {
			return infer.CheckResponse[UserManagementUserArgs]{Failures: []p.CheckFailure{{
				Property: "authMethods",
				Reason:   "authMethods must contain only BASIC and/or DIGEST",
			}}}, nil
		}
	}
	return infer.CheckResponse[UserManagementUserArgs]{Inputs: inputs}, nil
}

func (UserManagementUser) Create(ctx context.Context, req infer.CreateRequest[UserManagementUserArgs]) (infer.CreateResponse[UserManagementUserState], error) {
	if req.DryRun {
		return infer.CreateResponse[UserManagementUserState]{
			ID: "preview",
			Output: UserManagementUserState{
				UserManagementUserArgs: req.Inputs,
				UserId:                 intPtr(0),
			},
		}, nil
	}

	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[UserManagementUserState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := userManagementUserArgsToAPI(req.Inputs, 0)
	resp, err := users.Create(ctx, svc, &apiReq)
	if err != nil {
		return infer.CreateResponse[UserManagementUserState]{}, fmt.Errorf("creating user: %w", err)
	}

	time.Sleep(5 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.CreateResponse[UserManagementUserState]{}, activationErr
		}
	}

	if len(req.Inputs.AuthMethods) > 0 {
		_, err = users.EnrollUser(ctx, svc, resp.ID, users.EnrollUserRequest{
			AuthMethods: req.Inputs.AuthMethods,
			Password:    resp.Password,
		})
		if err != nil {
			return infer.CreateResponse[UserManagementUserState]{}, fmt.Errorf("enrolling user: %w", err)
		}
	}

	state := UserManagementUserState{
		UserManagementUserArgs: req.Inputs,
		UserId:                 &resp.ID,
	}

	time.Sleep(5 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.CreateResponse[UserManagementUserState]{}, activationErr
		}
	}

	rule, err := users.Get(ctx, svc, resp.ID)
	if err != nil {
		return infer.CreateResponse[UserManagementUserState]{
			ID:     strconv.Itoa(resp.ID),
			Output: state,
		}, nil
	}

	fullState := userManagementUserAPIToState(rule)
	return infer.CreateResponse[UserManagementUserState]{
		ID:     strconv.Itoa(resp.ID),
		Output: fullState,
	}, nil
}

func (UserManagementUser) Read(ctx context.Context, req infer.ReadRequest[UserManagementUserArgs, UserManagementUserState]) (infer.ReadResponse[UserManagementUserArgs, UserManagementUserState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[UserManagementUserArgs, UserManagementUserState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		rule, lookupErr := users.GetUserByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[UserManagementUserArgs, UserManagementUserState]{}, nil
			}
			return infer.ReadResponse[UserManagementUserArgs, UserManagementUserState]{}, lookupErr
		}
		id = rule.ID
	}

	rule, err := users.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[UserManagementUserArgs, UserManagementUserState]{}, nil
		}
		return infer.ReadResponse[UserManagementUserArgs, UserManagementUserState]{}, err
	}

	state := userManagementUserAPIToState(rule)
	args := userManagementUserStateToArgs(rule)
	return infer.ReadResponse[UserManagementUserArgs, UserManagementUserState]{
		ID:     strconv.Itoa(rule.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (UserManagementUser) Update(ctx context.Context, req infer.UpdateRequest[UserManagementUserArgs, UserManagementUserState]) (infer.UpdateResponse[UserManagementUserState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[UserManagementUserState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[UserManagementUserState]{}, fmt.Errorf("invalid user id: %w", err)
	}

	apiReq := userManagementUserArgsToAPI(req.Inputs, id)
	if _, _, err := users.Update(ctx, svc, id, &apiReq); err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.UpdateResponse[UserManagementUserState]{}, nil
		}
		return infer.UpdateResponse[UserManagementUserState]{}, fmt.Errorf("updating user: %w", err)
	}

	time.Sleep(5 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[UserManagementUserState]{}, activationErr
		}
	}

	if len(req.Inputs.AuthMethods) > 0 {
		_, err = users.EnrollUser(ctx, svc, id, users.EnrollUserRequest{
			AuthMethods: req.Inputs.AuthMethods,
			Password:    apiReq.Password,
		})
		if err != nil {
			return infer.UpdateResponse[UserManagementUserState]{}, fmt.Errorf("enrolling user: %w", err)
		}
	}

	time.Sleep(5 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[UserManagementUserState]{}, activationErr
		}
	}

	rule, err := users.Get(ctx, svc, id)
	if err != nil {
		return infer.UpdateResponse[UserManagementUserState]{
			Output: UserManagementUserState{UserManagementUserArgs: req.Inputs, UserId: &id},
		}, nil
	}
	return infer.UpdateResponse[UserManagementUserState]{Output: userManagementUserAPIToState(rule)}, nil
}

func (UserManagementUser) Delete(ctx context.Context, req infer.DeleteRequest[UserManagementUserState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid user id: %w", err)
	}

	if err := detachUserFromFilteringRules(ctx, client, id); err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("detaching user from rules: %w", err)
	}

	if _, err := users.Delete(ctx, svc, id); err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.DeleteResponse{}, nil
		}
		return infer.DeleteResponse{}, fmt.Errorf("deleting user: %w", err)
	}

	time.Sleep(5 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}
	return infer.DeleteResponse{}, nil
}

func (UserManagementUser) Annotate(a infer.Annotator) {
	describeResource(a, &UserManagementUser{}, `The zia.UserManagementUser resource manages user accounts in the Zscaler Internet Access (ZIA) cloud.
Users can be assigned to departments and groups, and enrolled with authentication methods such as
BASIC or DIGEST.

{{% examples %}}
## Example Usage

{{% example %}}
### Basic User Management

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

import * as pulumi from "@pulumi/pulumi";

const cfg = new pulumi.Config();
const userPassword = cfg.requireSecret("userPassword");

const example = new zia.UserManagementUser("example", {
    name: "John Doe",
    email: "john.doe@example.com",
    password: userPassword,
    authMethods: ["BASIC"],
    groups: [12345],
    department: {
        id: 67890,
    },
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

cfg = pulumi.Config()
user_password = cfg.require_secret("userPassword")

example = zia.UserManagementUser("example",
    name="John Doe",
    email="john.doe@example.com",
    password=user_password,
    auth_methods=["BASIC"],
    groups=[12345],
    department={
        "id": 67890,
    },
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:UserManagementUser
    properties:
      name: John Doe
      email: john.doe@example.com
      password:
        fn::secret: ${userPassword}
      authMethods:
        - BASIC
      groups:
        - 12345
      department:
        id: 67890
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing user can be imported using its ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:UserManagementUser example 12345
`+tripleBacktick("")+`
`)
}

func (a *UserManagementUserArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The user's full name. Maximum 127 characters.")
	ann.Describe(&a.Email, "The user's email address. Maximum 127 characters.")
	ann.Describe(&a.Comments, "Comments or notes about the user. Maximum 10240 characters.")
	ann.Describe(&a.TempAuthEmail, "Temporary authentication email for the user.")
	ann.Describe(&a.AuthMethods, "Authentication methods for the user. Accepted values: 'BASIC', 'DIGEST'.")
	ann.Describe(&a.Password, "The user's password. This is a secret and will not be stored in plaintext in the state.")
	ann.Describe(&a.Groups, "List of group IDs the user belongs to.")
	ann.Describe(&a.Department, "The department the user belongs to.")
}

func (s *UserManagementUserState) Annotate(a infer.Annotator) {
	a.Describe(&s.UserId, "The unique identifier for the user assigned by the ZIA cloud.")
}

func userManagementUserArgsToAPI(in UserManagementUserArgs, existingID int) users.Users {
	result := users.Users{
		ID:            existingID,
		Name:          in.Name,
		Email:         in.Email,
		Comments:      ptrToString(in.Comments),
		TempAuthEmail: ptrToString(in.TempAuthEmail),
		Password:      in.Password,
		Groups:        idsToUserGroups(in.Groups),
	}
	if in.Department != nil {
		result.Department = &common.UserDepartment{
			ID:       ptrToIntDefault(in.Department.Id, 0),
			Name:     ptrToString(in.Department.Name),
			IdpID:    ptrToIntDefault(in.Department.IdpId, 0),
			Comments: ptrToString(in.Department.Comments),
			Deleted:  ptrToBool(in.Department.Deleted),
		}
	}
	return result
}

func idsToUserGroups(ids []int) []common.UserGroups {
	if len(ids) == 0 {
		return nil
	}
	result := make([]common.UserGroups, len(ids))
	for i, id := range ids {
		result[i] = common.UserGroups{ID: id}
	}
	return result
}

func userGroupsToIDs(list []common.UserGroups) []int {
	if len(list) == 0 {
		return nil
	}
	ids := make([]int, len(list))
	for i, g := range list {
		ids[i] = g.ID
	}
	return ids
}

func userManagementUserAPIToState(rule *users.Users) UserManagementUserState {
	state := UserManagementUserState{
		UserManagementUserArgs: UserManagementUserArgs{
			Name:          rule.Name,
			Email:         rule.Email,
			Comments:      stringPtr(rule.Comments),
			TempAuthEmail: stringPtr(rule.TempAuthEmail),
			AuthMethods:   rule.AuthMethods,
			Groups:        userGroupsToIDs(rule.Groups),
			Password:      "", // never return password in state
		},
		UserId: &rule.ID,
	}
	if rule.Department != nil {
		state.Department = &UserDepartmentInput{
			Id:       intPtr(rule.Department.ID),
			Name:     stringPtr(rule.Department.Name),
			IdpId:    intPtr(rule.Department.IdpID),
			Comments: stringPtr(rule.Department.Comments),
			Deleted:  boolPtr(rule.Department.Deleted),
		}
	}
	return state
}

func userManagementUserStateToArgs(rule *users.Users) UserManagementUserArgs {
	s := userManagementUserAPIToState(rule)
	// Preserve password from inputs for update - but Read doesn't have it, so we need empty
	s.Password = ""
	return s.UserManagementUserArgs
}

// --- Invoke (data source) ---

// UserGroupOutput is a group in the invoke result.
type UserGroupOutput struct {
	Id       int    `pulumi:"groupId"` // Pulumi reserves "id" in function outputs
	Name     string `pulumi:"name"`
	IdpId    int    `pulumi:"idpId"`
	Comments string `pulumi:"comments"`
}

// UserDepartmentOutput is the department in the invoke result.
type UserDepartmentOutput struct {
	Id       int    `pulumi:"departmentId"` // Pulumi reserves "id" in function outputs
	Name     string `pulumi:"name"`
	IdpId    int    `pulumi:"idpId"`
	Comments string `pulumi:"comments"`
	Deleted  bool   `pulumi:"deleted"`
}

// GetUserManagementUserArgs are the inputs for the GetUserManagementUser invoke.
type GetUserManagementUserArgs struct {
	Id   *int    `pulumi:"userId,optional"` // Pulumi reserves "id" in function I/O
	Name *string `pulumi:"name,optional"`
}

// GetUserManagementUserResult is the output of the GetUserManagementUser invoke.
type GetUserManagementUserResult struct {
	Id            int                   `pulumi:"userId"` // Pulumi reserves "id" in function outputs
	Name          string                `pulumi:"name"`
	Email         string                `pulumi:"email"`
	Comments      string                `pulumi:"comments"`
	TempAuthEmail string                `pulumi:"tempAuthEmail"`
	AuthMethods   []string              `pulumi:"authMethods"`
	AdminUser     bool                  `pulumi:"adminUser"`
	Type          string                `pulumi:"type"`
	Department    *UserDepartmentOutput `pulumi:"department,optional"`
	Groups        []UserGroupOutput     `pulumi:"groups"`
}

// GetUserManagementUser implements the zia:index:GetUserManagementUser invoke.
type GetUserManagementUser struct{}

func (f *GetUserManagementUser) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a user management user by ID or name.")
}

func (a *GetUserManagementUserArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Id, "The ID of the user to look up.")
	ann.Describe(&a.Name, "The name of the user to look up.")
}

func (r *GetUserManagementUserResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the user.")
	a.Describe(&r.Name, "The full name of the user.")
	a.Describe(&r.Email, "The email address of the user.")
	a.Describe(&r.Comments, "Comments or notes about the user.")
	a.Describe(&r.TempAuthEmail, "The temporary authentication email for the user.")
	a.Describe(&r.AuthMethods, "The authentication methods configured for the user.")
	a.Describe(&r.AdminUser, "Whether the user is an admin user.")
	a.Describe(&r.Type, "The type of the user.")
	a.Describe(&r.Department, "The department the user belongs to.")
	a.Describe(&r.Groups, "The list of groups the user belongs to.")
}

func (*GetUserManagementUser) Invoke(ctx context.Context, req infer.FunctionRequest[GetUserManagementUserArgs]) (infer.FunctionResponse[GetUserManagementUserResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetUserManagementUserResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	var resp *users.Users
	if req.Input.Id != nil {
		r, err := users.Get(ctx, svc, *req.Input.Id)
		if err != nil {
			return infer.FunctionResponse[GetUserManagementUserResult]{}, err
		}
		resp = r
	}
	if resp == nil && req.Input.Name != nil && *req.Input.Name != "" {
		all, err := users.GetAllUsers(ctx, svc, nil)
		if err != nil {
			return infer.FunctionResponse[GetUserManagementUserResult]{}, err
		}
		for i := range all {
			if all[i].Name == *req.Input.Name {
				resp = &all[i]
				break
			}
		}
	}

	if resp == nil {
		return infer.FunctionResponse[GetUserManagementUserResult]{}, fmt.Errorf("couldn't find any user with id %v or name %v", req.Input.Id, ptrToString(req.Input.Name))
	}

	grps := make([]UserGroupOutput, len(resp.Groups))
	for i, g := range resp.Groups {
		grps[i] = UserGroupOutput{Id: g.ID, Name: g.Name, IdpId: g.IdpID, Comments: g.Comments}
	}

	result := GetUserManagementUserResult{
		Id:            resp.ID,
		Name:          resp.Name,
		Email:         resp.Email,
		Comments:      resp.Comments,
		TempAuthEmail: resp.TempAuthEmail,
		AuthMethods:   resp.AuthMethods,
		AdminUser:     resp.AdminUser,
		Type:          resp.Type,
		Groups:        grps,
	}
	if resp.Department != nil {
		result.Department = &UserDepartmentOutput{
			Id:       resp.Department.ID,
			Name:     resp.Department.Name,
			IdpId:    resp.Department.IdpID,
			Comments: resp.Department.Comments,
			Deleted:  resp.Department.Deleted,
		}
	}
	return infer.FunctionResponse[GetUserManagementUserResult]{Output: result}, nil
}
