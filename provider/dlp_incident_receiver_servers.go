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

// Package provider implements the getDlpIncidentReceiverServer invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_dlp_incident_receiver_servers.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/dlp/dlp_incident_receiver_servers"
)

type GetDlpIncidentReceiverServerArgs struct {
	Id   *int    `pulumi:"resourceId,optional"`
	Name *string `pulumi:"name,optional"`
}

type GetDlpIncidentReceiverServerResult struct {
	Id     int    `pulumi:"resourceId"`
	Name   string `pulumi:"name"`
	Url    string `pulumi:"url"`
	Status string `pulumi:"status"`
	Flags  int    `pulumi:"flags"`
}

type GetDlpIncidentReceiverServer struct{}

func (f *GetDlpIncidentReceiverServer) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a DLP incident receiver server by ID or name.")
}

func (a *GetDlpIncidentReceiverServerArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Id, "The ID of the DLP incident receiver server to look up.")
	ann.Describe(&a.Name, "The name of the DLP incident receiver server to look up.")
}

func (r *GetDlpIncidentReceiverServerResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the DLP incident receiver server.")
	a.Describe(&r.Name, "The name of the DLP incident receiver server.")
	a.Describe(&r.Url, "The URL of the DLP incident receiver server.")
	a.Describe(&r.Status, "The status of the DLP incident receiver server.")
	a.Describe(&r.Flags, "The flags associated with the DLP incident receiver server.")
}

func (*GetDlpIncidentReceiverServer) Invoke(ctx context.Context, req infer.FunctionRequest[GetDlpIncidentReceiverServerArgs]) (infer.FunctionResponse[GetDlpIncidentReceiverServerResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetDlpIncidentReceiverServerResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	var resp *dlp_incident_receiver_servers.IncidentReceiverServers
	if req.Input.Id != nil && *req.Input.Id != 0 {
		r, err := dlp_incident_receiver_servers.Get(ctx, svc, *req.Input.Id)
		if err != nil {
			return infer.FunctionResponse[GetDlpIncidentReceiverServerResult]{}, err
		}
		resp = r
	}
	if resp == nil && req.Input.Name != nil && *req.Input.Name != "" {
		r, err := dlp_incident_receiver_servers.GetByName(ctx, svc, *req.Input.Name)
		if err != nil {
			return infer.FunctionResponse[GetDlpIncidentReceiverServerResult]{}, err
		}
		resp = r
	}
	if resp == nil {
		return infer.FunctionResponse[GetDlpIncidentReceiverServerResult]{}, fmt.Errorf("couldn't find any DLP incident receiver server with id %v or name %s", req.Input.Id, ptrToString(req.Input.Name))
	}

	return infer.FunctionResponse[GetDlpIncidentReceiverServerResult]{Output: GetDlpIncidentReceiverServerResult{
		Id: resp.ID, Name: resp.Name, Url: resp.URL, Status: resp.Status, Flags: resp.Flags,
	}}, nil
}
