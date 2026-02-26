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

// Package provider implements the getDlpIcapServer invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_dlp_icap_servers.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/dlp/dlp_icap_servers"
)

type GetDlpIcapServerArgs struct {
	Id   *int    `pulumi:"resourceId,optional"`
	Name *string `pulumi:"name,optional"`
}

type GetDlpIcapServerResult struct {
	Id     int    `pulumi:"resourceId"`
	Name   string `pulumi:"name"`
	Url    string `pulumi:"url"`
	Status string `pulumi:"status"`
}

type GetDlpIcapServer struct{}

func (f *GetDlpIcapServer) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a DLP ICAP server by ID or name.")
}

func (a *GetDlpIcapServerArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Id, "The ID of the DLP ICAP server to look up.")
	ann.Describe(&a.Name, "The name of the DLP ICAP server to look up.")
}

func (r *GetDlpIcapServerResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the DLP ICAP server.")
	a.Describe(&r.Name, "The name of the DLP ICAP server.")
	a.Describe(&r.Url, "The URL of the DLP ICAP server.")
	a.Describe(&r.Status, "The status of the DLP ICAP server.")
}

func (*GetDlpIcapServer) Invoke(ctx context.Context, req infer.FunctionRequest[GetDlpIcapServerArgs]) (infer.FunctionResponse[GetDlpIcapServerResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetDlpIcapServerResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	var resp *dlp_icap_servers.DLPICAPServers
	if req.Input.Id != nil && *req.Input.Id != 0 {
		r, err := dlp_icap_servers.Get(ctx, svc, *req.Input.Id)
		if err != nil {
			return infer.FunctionResponse[GetDlpIcapServerResult]{}, err
		}
		resp = r
	}
	if resp == nil && req.Input.Name != nil && *req.Input.Name != "" {
		r, err := dlp_icap_servers.GetByName(ctx, svc, *req.Input.Name)
		if err != nil {
			return infer.FunctionResponse[GetDlpIcapServerResult]{}, err
		}
		resp = r
	}
	if resp == nil {
		return infer.FunctionResponse[GetDlpIcapServerResult]{}, fmt.Errorf("couldn't find any DLP ICAP server with id %v or name %s", req.Input.Id, ptrToString(req.Input.Name))
	}

	return infer.FunctionResponse[GetDlpIcapServerResult]{Output: GetDlpIcapServerResult{
		Id: resp.ID, Name: resp.Name, Url: resp.URL, Status: resp.Status,
	}}, nil
}
