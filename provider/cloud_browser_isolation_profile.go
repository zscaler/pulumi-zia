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

// Package provider implements the getCloudBrowserIsolationProfile invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_cloud_browser_isolation_profile.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/browser_isolation"
)

type GetCloudBrowserIsolationProfileArgs struct {
	Name *string `pulumi:"name,optional"`
}

type GetCloudBrowserIsolationProfileResult struct {
	Id             string `pulumi:"resourceId"`
	Name           string `pulumi:"name"`
	Url            string `pulumi:"url"`
	DefaultProfile bool   `pulumi:"defaultProfile"`
}

type GetCloudBrowserIsolationProfile struct{}

func (f *GetCloudBrowserIsolationProfile) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a cloud browser isolation profile by name.")
}

func (a *GetCloudBrowserIsolationProfileArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the cloud browser isolation profile to look up.")
}

func (r *GetCloudBrowserIsolationProfileResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the cloud browser isolation profile.")
	a.Describe(&r.Name, "The name of the cloud browser isolation profile.")
	a.Describe(&r.Url, "The URL of the cloud browser isolation profile.")
	a.Describe(&r.DefaultProfile, "Whether this is the default profile.")
}

func (*GetCloudBrowserIsolationProfile) Invoke(ctx context.Context, req infer.FunctionRequest[GetCloudBrowserIsolationProfileArgs]) (infer.FunctionResponse[GetCloudBrowserIsolationProfileResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetCloudBrowserIsolationProfileResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	if req.Input.Name == nil || *req.Input.Name == "" {
		return infer.FunctionResponse[GetCloudBrowserIsolationProfileResult]{}, fmt.Errorf("name is required")
	}
	resp, err := browser_isolation.GetByName(ctx, svc, *req.Input.Name)
	if err != nil {
		return infer.FunctionResponse[GetCloudBrowserIsolationProfileResult]{}, err
	}

	return infer.FunctionResponse[GetCloudBrowserIsolationProfileResult]{Output: GetCloudBrowserIsolationProfileResult{
		Id: resp.ID, Name: resp.Name, Url: resp.URL, DefaultProfile: resp.DefaultProfile,
	}}, nil
}
