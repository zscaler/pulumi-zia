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

// Package provider implements the getDeviceGroup invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_device_group.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/devicegroups"
)

type GetDeviceGroupArgs struct {
	Id   *int    `pulumi:"resourceId,optional"`
	Name *string `pulumi:"name,optional"`
}

type DeviceGroupItem struct {
	Id          int    `pulumi:"resourceId"`
	Name        string `pulumi:"name"`
	GroupType   string `pulumi:"groupType"`
	Description string `pulumi:"description"`
	OsType      string `pulumi:"osType"`
	Predefined  bool   `pulumi:"predefined"`
	DeviceNames string `pulumi:"deviceNames"`
	DeviceCount int    `pulumi:"deviceCount"`
}

type GetDeviceGroupResult struct {
	Id          *int              `pulumi:"resourceId,optional"`
	Name        *string           `pulumi:"name,optional"`
	GroupType   *string           `pulumi:"groupType,optional"`
	Description *string           `pulumi:"description,optional"`
	OsType      *string           `pulumi:"osType,optional"`
	Predefined  *bool             `pulumi:"predefined,optional"`
	DeviceNames *string           `pulumi:"deviceNames,optional"`
	DeviceCount *int              `pulumi:"deviceCount,optional"`
	List        []DeviceGroupItem `pulumi:"list"`
}

type GetDeviceGroup struct{}

func (f *GetDeviceGroup) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a device group by ID or name.")
}

func (a *GetDeviceGroupArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Id, "The ID of the device group to look up.")
	ann.Describe(&a.Name, "The name of the device group to look up.")
}

func (r *GetDeviceGroupResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the device group.")
	a.Describe(&r.Name, "The name of the device group.")
	a.Describe(&r.GroupType, "The type of the device group.")
	a.Describe(&r.Description, "The description of the device group.")
	a.Describe(&r.OsType, "The OS type of the device group.")
	a.Describe(&r.Predefined, "Whether the device group is predefined.")
	a.Describe(&r.DeviceNames, "The device names in the group.")
	a.Describe(&r.DeviceCount, "The number of devices in the group.")
	a.Describe(&r.List, "The list of all device groups when no specific filter is provided.")
}

func (*GetDeviceGroup) Invoke(ctx context.Context, req infer.FunctionRequest[GetDeviceGroupArgs]) (infer.FunctionResponse[GetDeviceGroupResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetDeviceGroupResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	var resp *devicegroups.DeviceGroups
	if req.Input.Id != nil && *req.Input.Id != 0 {
		all, err := devicegroups.GetAllDevicesGroups(ctx, svc)
		if err != nil {
			return infer.FunctionResponse[GetDeviceGroupResult]{}, err
		}
		for i := range all {
			if all[i].ID == *req.Input.Id {
				resp = &all[i]
				break
			}
		}
	}
	if resp == nil && req.Input.Name != nil && *req.Input.Name != "" {
		r, err := devicegroups.GetDeviceGroupByName(ctx, svc, *req.Input.Name)
		if err != nil {
			return infer.FunctionResponse[GetDeviceGroupResult]{}, err
		}
		resp = r
	}

	if resp != nil {
		return infer.FunctionResponse[GetDeviceGroupResult]{Output: GetDeviceGroupResult{
			Id:          intPtr(resp.ID),
			Name:        stringPtr(resp.Name),
			GroupType:   stringPtr(resp.GroupType),
			Description: stringPtr(resp.Description),
			OsType:      stringPtr(resp.OSType),
			Predefined:  boolPtr(resp.Predefined),
			DeviceNames: stringPtr(resp.DeviceNames),
			DeviceCount: intPtr(resp.DeviceCount),
			List:        []DeviceGroupItem{},
		}}, nil
	}

	// Return all
	all, err := devicegroups.GetAllDevicesGroups(ctx, svc)
	if err != nil {
		return infer.FunctionResponse[GetDeviceGroupResult]{}, err
	}
	if len(all) == 0 {
		return infer.FunctionResponse[GetDeviceGroupResult]{}, fmt.Errorf("no device groups found")
	}

	list := make([]DeviceGroupItem, len(all))
	for i, dg := range all {
		list[i] = DeviceGroupItem{
			Id: dg.ID, Name: dg.Name, GroupType: dg.GroupType, Description: dg.Description,
			OsType: dg.OSType, Predefined: dg.Predefined, DeviceNames: dg.DeviceNames, DeviceCount: dg.DeviceCount,
		}
	}
	first := all[0]
	return infer.FunctionResponse[GetDeviceGroupResult]{Output: GetDeviceGroupResult{
		Id:          intPtr(first.ID),
		Name:        stringPtr(first.Name),
		GroupType:   stringPtr(first.GroupType),
		Description: stringPtr(first.Description),
		OsType:      stringPtr(first.OSType),
		Predefined:  boolPtr(first.Predefined),
		DeviceNames: stringPtr(first.DeviceNames),
		DeviceCount: intPtr(first.DeviceCount),
		List:        list,
	}}, nil
}
