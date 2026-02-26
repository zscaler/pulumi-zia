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

// Package provider implements the getUserManagementDepartment invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_user_management_departments.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/usermanagement/departments"
)

type GetUserManagementDepartmentArgs struct {
	Id   *int    `pulumi:"resourceId,optional"`
	Name *string `pulumi:"name,optional"`
}

type GetUserManagementDepartmentResult struct {
	Id       int    `pulumi:"resourceId"`
	Name     string `pulumi:"name"`
	IdpId    int    `pulumi:"idpId"`
	Comments string `pulumi:"comments"`
	Deleted  bool   `pulumi:"deleted"`
}

type GetUserManagementDepartment struct{}

func (f *GetUserManagementDepartment) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a user management department by ID or name.")
}

func (a *GetUserManagementDepartmentArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Id, "The ID of the department to look up.")
	ann.Describe(&a.Name, "The name of the department to look up.")
}

func (r *GetUserManagementDepartmentResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the department.")
	a.Describe(&r.Name, "The name of the department.")
	a.Describe(&r.IdpId, "The IDP ID associated with the department.")
	a.Describe(&r.Comments, "Comments or notes about the department.")
	a.Describe(&r.Deleted, "Whether the department has been deleted.")
}

func (*GetUserManagementDepartment) Invoke(ctx context.Context, req infer.FunctionRequest[GetUserManagementDepartmentArgs]) (infer.FunctionResponse[GetUserManagementDepartmentResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetUserManagementDepartmentResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	if (req.Input.Id == nil || *req.Input.Id == 0) && (req.Input.Name == nil || *req.Input.Name == "") {
		return infer.FunctionResponse[GetUserManagementDepartmentResult]{}, fmt.Errorf("either 'id' or 'name' must be provided")
	}

	allDepartments, err := departments.GetAll(ctx, svc, nil)
	if err != nil {
		return infer.FunctionResponse[GetUserManagementDepartmentResult]{}, fmt.Errorf("error getting all departments: %w", err)
	}

	id := ptrToIntDefault(req.Input.Id, 0)
	name := ptrToString(req.Input.Name)
	var resp *departments.Department

	if id != 0 {
		for i := range allDepartments {
			if allDepartments[i].ID == id {
				resp = &allDepartments[i]
				break
			}
		}
		if resp == nil {
			return infer.FunctionResponse[GetUserManagementDepartmentResult]{}, fmt.Errorf("error getting department by ID %d: department not found", id)
		}
	}
	if resp == nil && name != "" {
		for i := range allDepartments {
			if allDepartments[i].Name == name {
				resp = &allDepartments[i]
				break
			}
		}
		if resp == nil {
			return infer.FunctionResponse[GetUserManagementDepartmentResult]{}, fmt.Errorf("error getting department by name %s: department not found", name)
		}
	}

	if resp == nil {
		return infer.FunctionResponse[GetUserManagementDepartmentResult]{}, fmt.Errorf("couldn't find any department with name '%s' or id '%d'", name, id)
	}

	return infer.FunctionResponse[GetUserManagementDepartmentResult]{Output: GetUserManagementDepartmentResult{
		Id: resp.ID, Name: resp.Name, IdpId: resp.IdpID, Comments: resp.Comments, Deleted: resp.Deleted,
	}}, nil
}
