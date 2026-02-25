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

// Package provider implements the getLocationGroup invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_location_groups.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/location/locationgroups"
)

type GetLocationGroupArgs struct {
	Id   *int    `pulumi:"resourceId,optional"`
	Name *string `pulumi:"name,optional"`
}

type GetLocationGroupResult struct {
	Id          int    `pulumi:"resourceId"`
	Name        string `pulumi:"name"`
	Deleted     bool   `pulumi:"deleted"`
	GroupType   string `pulumi:"groupType"`
	Comments    string `pulumi:"comments"`
	LastModTime int    `pulumi:"lastModTime"`
	Predefined  bool   `pulumi:"predefined"`
	LocationIds []int  `pulumi:"locationIds"`
}

type GetLocationGroup struct{}

func (f *GetLocationGroup) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a location group by ID or name.")
}

func (a *GetLocationGroupArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Id, "The ID of the location group to look up.")
	ann.Describe(&a.Name, "The name of the location group to look up.")
}

func (r *GetLocationGroupResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the location group.")
	a.Describe(&r.Name, "The name of the location group.")
	a.Describe(&r.Deleted, "Whether the location group has been deleted.")
	a.Describe(&r.GroupType, "The type of the location group.")
	a.Describe(&r.Comments, "Comments or notes about the location group.")
	a.Describe(&r.LastModTime, "The last modification time of the location group (epoch).")
	a.Describe(&r.Predefined, "Whether the location group is predefined by Zscaler.")
	a.Describe(&r.LocationIds, "List of location IDs that belong to this group.")
}

func (*GetLocationGroup) Invoke(ctx context.Context, req infer.FunctionRequest[GetLocationGroupArgs]) (infer.FunctionResponse[GetLocationGroupResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetLocationGroupResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	var resp *locationgroups.LocationGroup
	if req.Input.Id != nil && *req.Input.Id != 0 {
		r, err := locationgroups.GetLocationGroup(ctx, svc, *req.Input.Id)
		if err != nil {
			return infer.FunctionResponse[GetLocationGroupResult]{}, err
		}
		resp = r
	}
	if resp == nil && req.Input.Name != nil && *req.Input.Name != "" {
		r, err := locationgroups.GetLocationGroupByName(ctx, svc, *req.Input.Name)
		if err != nil {
			return infer.FunctionResponse[GetLocationGroupResult]{}, err
		}
		resp = r
	}
	if resp == nil {
		return infer.FunctionResponse[GetLocationGroupResult]{}, fmt.Errorf("couldn't find any location group with id %v or name %s", req.Input.Id, ptrToString(req.Input.Name))
	}

	locIds := idsFromIDNameExtensions(resp.Locations)
	return infer.FunctionResponse[GetLocationGroupResult]{Output: GetLocationGroupResult{
		Id:          resp.ID,
		Name:        resp.Name,
		Deleted:    resp.Deleted,
		GroupType:   resp.GroupType,
		Comments:    resp.Comments,
		LastModTime: resp.LastModTime,
		Predefined:  resp.Predefined,
		LocationIds: locIds,
	}}, nil
}
