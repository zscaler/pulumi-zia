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

// Package provider implements the getFwNetworkService invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_fw_filtering_network_services.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/firewallpolicies/networkservices"
)

// GetFwNetworkService implements the zia:index:GetFwNetworkService invoke.
type GetFwNetworkService struct{}

func (f *GetFwNetworkService) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a firewall network service by ID or name.")
}

// GetFwNetworkServiceArgs are the inputs for the GetFwNetworkService invoke.
type GetFwNetworkServiceArgs struct {
	Id   *int    `pulumi:"networkServiceId,optional"`
	Name *string `pulumi:"name,optional"`
}

// GetFwNetworkServiceResult is the output of the GetFwNetworkService invoke.
type GetFwNetworkServiceResult struct {
	Id   int    `pulumi:"networkServiceId"`
	Name string `pulumi:"name"`
}

func (a *GetFwNetworkServiceArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Id, "The ID of the network service to look up.")
	ann.Describe(&a.Name, "The name of the network service to look up.")
}

func (r *GetFwNetworkServiceResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the network service.")
	a.Describe(&r.Name, "The name of the network service.")
}

func (*GetFwNetworkService) Invoke(ctx context.Context, req infer.FunctionRequest[GetFwNetworkServiceArgs]) (infer.FunctionResponse[GetFwNetworkServiceResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetFwNetworkServiceResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	var resp *networkservices.NetworkServices
	if req.Input.Id != nil && *req.Input.Id != 0 {
		r, err := networkservices.Get(ctx, svc, *req.Input.Id)
		if err != nil {
			return infer.FunctionResponse[GetFwNetworkServiceResult]{}, err
		}
		resp = r
	}
	if resp == nil && req.Input.Name != nil && *req.Input.Name != "" {
		r, err := networkservices.GetByName(ctx, svc, *req.Input.Name, nil, nil)
		if err != nil {
			return infer.FunctionResponse[GetFwNetworkServiceResult]{}, err
		}
		resp = r
	}
	if resp == nil {
		return infer.FunctionResponse[GetFwNetworkServiceResult]{}, fmt.Errorf("either 'networkServiceId' or 'name' must be provided")
	}

	return infer.FunctionResponse[GetFwNetworkServiceResult]{Output: GetFwNetworkServiceResult{
		Id:   resp.ID,
		Name: resp.Name,
	}}, nil
}
