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

// Package provider implements the getTimeWindow invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_fw_filtering_time_window.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/firewallpolicies/timewindow"
)

type GetTimeWindowArgs struct {
	Id   *int    `pulumi:"resourceId,optional"`
	Name *string `pulumi:"name,optional"`
}

type GetTimeWindowResult struct {
	Id        int      `pulumi:"resourceId"`
	Name      string   `pulumi:"name"`
	StartTime int      `pulumi:"startTime"`
	EndTime   int      `pulumi:"endTime"`
	DayOfWeek []string `pulumi:"dayOfWeek"`
}

type GetTimeWindow struct{}

func (f *GetTimeWindow) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a firewall time window by ID or name.")
}

func (a *GetTimeWindowArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Id, "The ID of the time window to look up.")
	ann.Describe(&a.Name, "The name of the time window to look up.")
}

func (r *GetTimeWindowResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the time window.")
	a.Describe(&r.Name, "The name of the time window.")
	a.Describe(&r.StartTime, "The start time of the time window (minutes from midnight).")
	a.Describe(&r.EndTime, "The end time of the time window (minutes from midnight).")
	a.Describe(&r.DayOfWeek, "The days of the week the time window applies to.")
}

func (*GetTimeWindow) Invoke(ctx context.Context, req infer.FunctionRequest[GetTimeWindowArgs]) (infer.FunctionResponse[GetTimeWindowResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetTimeWindowResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	var resp *timewindow.TimeWindow
	if req.Input.Id != nil && *req.Input.Id != 0 {
		all, err := timewindow.GetAll(ctx, svc)
		if err != nil {
			return infer.FunctionResponse[GetTimeWindowResult]{}, err
		}
		for i := range all {
			if all[i].ID == *req.Input.Id {
				resp = &all[i]
				break
			}
		}
		if resp == nil {
			return infer.FunctionResponse[GetTimeWindowResult]{}, fmt.Errorf("no time window found with id %d", *req.Input.Id)
		}
	}
	if resp == nil && req.Input.Name != nil && *req.Input.Name != "" {
		r, err := timewindow.GetTimeWindowByName(ctx, svc, *req.Input.Name)
		if err != nil {
			return infer.FunctionResponse[GetTimeWindowResult]{}, err
		}
		resp = r
	}
	if resp == nil {
		return infer.FunctionResponse[GetTimeWindowResult]{}, fmt.Errorf("either 'resourceId' or 'name' must be provided")
	}

	return infer.FunctionResponse[GetTimeWindowResult]{Output: GetTimeWindowResult{
		Id:        resp.ID,
		Name:      resp.Name,
		StartTime: int(resp.StartTime),
		EndTime:   int(resp.EndTime),
		DayOfWeek: resp.DayOfWeek,
	}}, nil
}
