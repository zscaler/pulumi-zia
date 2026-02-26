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

// Package provider implements the getCloudApplications invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_cloud_applications.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/cloudapplications/cloudapplications"
)

type GetCloudApplicationsArgs struct {
	PolicyType *string  `pulumi:"policyType"`
	AppClass   []string `pulumi:"appClass,optional"`
	AppName    *string  `pulumi:"appName,optional"`
}

type CloudApplicationItem struct {
	App        string `pulumi:"app"`
	AppName    string `pulumi:"appName"`
	Parent     string `pulumi:"parent"`
	ParentName string `pulumi:"parentName"`
}

type GetCloudApplicationsResult struct {
	Applications []CloudApplicationItem `pulumi:"applications"`
}

type GetCloudApplications struct{}

func (f *GetCloudApplications) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up cloud applications by policy type, application class, or application name.")
}

func (a *GetCloudApplicationsArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.PolicyType, "The policy type to filter by. Accepted values: 'cloud_application_policy', 'cloud_application_ssl_policy'.")
	ann.Describe(&a.AppClass, "The application class(es) to filter by.")
	ann.Describe(&a.AppName, "The application name to filter by.")
}

func (r *GetCloudApplicationsResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Applications, "The list of cloud applications matching the filter criteria.")
}

func (*GetCloudApplications) Invoke(ctx context.Context, req infer.FunctionRequest[GetCloudApplicationsArgs]) (infer.FunctionResponse[GetCloudApplicationsResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetCloudApplicationsResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	if req.Input.PolicyType == nil || *req.Input.PolicyType == "" {
		return infer.FunctionResponse[GetCloudApplicationsResult]{}, fmt.Errorf("policyType is required")
	}
	policyType := *req.Input.PolicyType

	params := map[string]interface{}{}
	if len(req.Input.AppClass) > 0 {
		var iface []interface{}
		for _, s := range req.Input.AppClass {
			iface = append(iface, s)
		}
		params["appClass"] = iface
	}

	var apps []cloudapplications.CloudApplications
	var err error
	switch policyType {
	case "cloud_application_policy":
		apps, err = cloudapplications.GetCloudApplicationPolicy(ctx, svc, params)
	case "cloud_application_ssl_policy":
		apps, err = cloudapplications.GetCloudApplicationSSLPolicy(ctx, svc, params)
	default:
		return infer.FunctionResponse[GetCloudApplicationsResult]{}, fmt.Errorf("invalid policyType: %s", policyType)
	}
	if err != nil {
		return infer.FunctionResponse[GetCloudApplicationsResult]{}, err
	}
	if len(apps) == 0 {
		return infer.FunctionResponse[GetCloudApplicationsResult]{}, fmt.Errorf("no cloud applications found")
	}

	list := make([]CloudApplicationItem, 0)
	for _, app := range apps {
		if req.Input.AppName != nil && *req.Input.AppName != "" && app.AppName != *req.Input.AppName {
			continue
		}
		list = append(list, CloudApplicationItem{
			App: app.App, AppName: app.AppName, Parent: app.Parent, ParentName: app.ParentName,
		})
	}
	return infer.FunctionResponse[GetCloudApplicationsResult]{Output: GetCloudApplicationsResult{Applications: list}}, nil
}
