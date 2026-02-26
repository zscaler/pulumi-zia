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

// Package provider implements the getDlpIdmProfileLite invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_dlp_idm_profiles_lite.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/dlp/dlp_idm_profile_lite"
)

type GetDlpIdmProfileLiteArgs struct {
	ProfileId    *int    `pulumi:"profileId,optional"`
	TemplateName *string `pulumi:"templateName,optional"`
	ActiveOnly   *bool   `pulumi:"activeOnly,optional"`
}

type IdNameExtensionsOutput struct {
	Id   int    `pulumi:"resourceId"`
	Name string `pulumi:"name"`
}

type GetDlpIdmProfileLiteResult struct {
	ProfileId        int                     `pulumi:"profileId"`
	TemplateName     string                  `pulumi:"templateName"`
	NumDocuments     int                     `pulumi:"numDocuments"`
	LastModifiedTime int                     `pulumi:"lastModifiedTime"`
	ClientVm         *IdNameExtensionsOutput `pulumi:"clientVm,optional"`
	LastModifiedBy   *IdNameExtensionsOutput `pulumi:"lastModifiedBy,optional"`
}

type GetDlpIdmProfileLite struct{}

func (f *GetDlpIdmProfileLite) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a DLP IDM profile (lite) by ID or template name.")
}

func (a *GetDlpIdmProfileLiteArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.ProfileId, "The ID of the DLP IDM profile to look up.")
	ann.Describe(&a.TemplateName, "The template name of the DLP IDM profile to look up.")
	ann.Describe(&a.ActiveOnly, "If true, only return active profiles.")
}

func (r *GetDlpIdmProfileLiteResult) Annotate(a infer.Annotator) {
	a.Describe(&r.ProfileId, "The ID of the DLP IDM profile.")
	a.Describe(&r.TemplateName, "The template name of the DLP IDM profile.")
	a.Describe(&r.NumDocuments, "The number of documents in the profile.")
	a.Describe(&r.LastModifiedTime, "The last modification time (epoch).")
	a.Describe(&r.ClientVm, "The client VM associated with the profile.")
	a.Describe(&r.LastModifiedBy, "The user who last modified the profile.")
}

func (*GetDlpIdmProfileLite) Invoke(ctx context.Context, req infer.FunctionRequest[GetDlpIdmProfileLiteArgs]) (infer.FunctionResponse[GetDlpIdmProfileLiteResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetDlpIdmProfileLiteResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	activeOnly := ptrToBool(req.Input.ActiveOnly)

	var resp *dlp_idm_profile_lite.DLPIDMProfileLite
	if req.Input.ProfileId != nil && *req.Input.ProfileId != 0 {
		r, err := dlp_idm_profile_lite.GetDLPProfileLiteID(ctx, svc, *req.Input.ProfileId, activeOnly)
		if err != nil {
			return infer.FunctionResponse[GetDlpIdmProfileLiteResult]{}, err
		}
		resp = r
	}
	if resp == nil && req.Input.TemplateName != nil && *req.Input.TemplateName != "" {
		r, err := dlp_idm_profile_lite.GetDLPProfileLiteByName(ctx, svc, *req.Input.TemplateName, activeOnly)
		if err != nil {
			return infer.FunctionResponse[GetDlpIdmProfileLiteResult]{}, err
		}
		resp = r
	}
	if resp == nil {
		return infer.FunctionResponse[GetDlpIdmProfileLiteResult]{}, fmt.Errorf("couldn't find any DLP IDM profile lite with name '%s' or id '%d'", ptrToString(req.Input.TemplateName), ptrToIntDefault(req.Input.ProfileId, 0))
	}

	result := GetDlpIdmProfileLiteResult{
		ProfileId:        resp.ProfileID,
		TemplateName:     resp.TemplateName,
		NumDocuments:     resp.NumDocuments,
		LastModifiedTime: resp.LastModifiedTime,
	}
	if resp.ClientVM != nil && (resp.ClientVM.ID != 0 || resp.ClientVM.Name != "") {
		result.ClientVm = &IdNameExtensionsOutput{Id: resp.ClientVM.ID, Name: resp.ClientVM.Name}
	}
	if resp.ModifiedBy != nil && (resp.ModifiedBy.ID != 0 || resp.ModifiedBy.Name != "") {
		result.LastModifiedBy = &IdNameExtensionsOutput{Id: resp.ModifiedBy.ID, Name: resp.ModifiedBy.Name}
	}
	return infer.FunctionResponse[GetDlpIdmProfileLiteResult]{Output: result}, nil
}
