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

// Package provider implements the getCasbTenant invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_casb_tenant.go.

package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/saas_security_api"
)

type GetCasbTenantArgs struct {
	TenantId       *int    `pulumi:"tenantId,optional"`
	TenantName     *string `pulumi:"tenantName,optional"`
	ActiveOnly     *bool   `pulumi:"activeOnly,optional"`
	IncludeDeleted *bool   `pulumi:"includeDeleted,optional"`
	AppType        *string `pulumi:"appType,optional"`
	App            *string `pulumi:"app,optional"`
}

type GetCasbTenantResult struct {
	TenantId                 int      `pulumi:"tenantId"`
	TenantName               string   `pulumi:"tenantName"`
	ModifiedTime             int      `pulumi:"modifiedTime"`
	LastTenantValidationTime int      `pulumi:"lastTenantValidationTime"`
	SaasApplication          string   `pulumi:"saasApplication"`
	EnterpriseTenantId       string   `pulumi:"enterpriseTenantId"`
	TenantWebhookEnabled     bool     `pulumi:"tenantWebhookEnabled"`
	TenantDeleted            bool     `pulumi:"tenantDeleted"`
	ReAuth                   bool     `pulumi:"reAuth"`
	FeaturesSupported        []string `pulumi:"featuresSupported"`
	Status                   []string `pulumi:"status"`
	ZscalerAppTenantId       *int     `pulumi:"zscalerAppTenantId,optional"`
}

type GetCasbTenant struct{}

func (f *GetCasbTenant) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a CASB tenant by ID or name.")
}

func (a *GetCasbTenantArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.TenantId, "The ID of the CASB tenant to look up.")
	ann.Describe(&a.TenantName, "The name of the CASB tenant to look up.")
	ann.Describe(&a.ActiveOnly, "If true, only return active tenants.")
	ann.Describe(&a.IncludeDeleted, "If true, include deleted tenants in the results.")
	ann.Describe(&a.AppType, "The application type to filter by.")
	ann.Describe(&a.App, "The application to filter by.")
}

func (r *GetCasbTenantResult) Annotate(a infer.Annotator) {
	a.Describe(&r.TenantId, "The ID of the CASB tenant.")
	a.Describe(&r.TenantName, "The name of the CASB tenant.")
	a.Describe(&r.ModifiedTime, "The last modification time (epoch).")
	a.Describe(&r.LastTenantValidationTime, "The last tenant validation time (epoch).")
	a.Describe(&r.SaasApplication, "The SaaS application associated with the tenant.")
	a.Describe(&r.EnterpriseTenantId, "The enterprise tenant ID.")
	a.Describe(&r.TenantWebhookEnabled, "Whether tenant webhook is enabled.")
	a.Describe(&r.TenantDeleted, "Whether the tenant has been deleted.")
	a.Describe(&r.ReAuth, "Whether re-authentication is required.")
	a.Describe(&r.FeaturesSupported, "The list of features supported by the tenant.")
	a.Describe(&r.Status, "The status of the tenant.")
	a.Describe(&r.ZscalerAppTenantId, "The Zscaler application tenant ID.")
}

func (*GetCasbTenant) Invoke(ctx context.Context, req infer.FunctionRequest[GetCasbTenantArgs]) (infer.FunctionResponse[GetCasbTenantResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetCasbTenantResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	queryParams := map[string]interface{}{}
	if req.Input.TenantId != nil && *req.Input.TenantId != 0 {
		queryParams["tenantId"] = strconv.Itoa(*req.Input.TenantId)
	}
	if req.Input.TenantName != nil && *req.Input.TenantName != "" {
		queryParams["tenantName"] = *req.Input.TenantName
	}
	if req.Input.ActiveOnly != nil {
		queryParams["activeOnly"] = *req.Input.ActiveOnly
	}
	if req.Input.IncludeDeleted != nil {
		queryParams["includeDeleted"] = *req.Input.IncludeDeleted
	}
	if req.Input.AppType != nil {
		queryParams["appType"] = *req.Input.AppType
	}
	if req.Input.App != nil {
		queryParams["app"] = strings.ToUpper(*req.Input.App)
	}

	tenants, err := saas_security_api.GetCasbTenantLite(ctx, svc, queryParams)
	if err != nil {
		return infer.FunctionResponse[GetCasbTenantResult]{}, err
	}

	id := ptrToIntDefault(req.Input.TenantId, 0)
	name := ptrToString(req.Input.TenantName)
	var matched *saas_security_api.CasbTenants
	for i := range tenants {
		if id != 0 && tenants[i].TenantID == id {
			matched = &tenants[i]
			break
		}
		if name != "" && tenants[i].TenantName == name {
			matched = &tenants[i]
			break
		}
	}
	if matched == nil {
		return infer.FunctionResponse[GetCasbTenantResult]{}, fmt.Errorf("couldn't find any CASB tenant with name '%s' or id '%d'", name, id)
	}

	zscalerId := (*int)(nil)
	if matched.ZscalerAppTenantID != nil {
		zscalerId = &matched.ZscalerAppTenantID.ID
	}
	return infer.FunctionResponse[GetCasbTenantResult]{Output: GetCasbTenantResult{
		TenantId:                 matched.TenantID,
		TenantName:               matched.TenantName,
		ModifiedTime:             matched.ModifiedTime,
		LastTenantValidationTime: matched.LastTenantValidationTime,
		SaasApplication:          matched.SaaSApplication,
		EnterpriseTenantId:       matched.EnterpriseTenantID,
		TenantWebhookEnabled:     matched.TenantWebhookEnabled,
		TenantDeleted:            matched.TenantDeleted,
		ReAuth:                   matched.ReAuth,
		FeaturesSupported:        matched.FeaturesSupported,
		Status:                   matched.Status,
		ZscalerAppTenantId:       zscalerId,
	}}, nil
}
