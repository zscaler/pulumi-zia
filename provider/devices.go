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

// Package provider implements the getDevice invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_devices.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/devicegroups"
)

type GetDeviceArgs struct {
	Id          *int    `pulumi:"resourceId,optional"`
	Name        *string `pulumi:"name,optional"`
	DeviceModel *string `pulumi:"deviceModel,optional"`
	OwnerName   *string `pulumi:"ownerName,optional"`
	OsType      *string `pulumi:"osType,optional"`
	OsVersion   *string `pulumi:"osVersion,optional"`
}

type GetDeviceResult struct {
	Id              int    `pulumi:"resourceId"`
	Name            string `pulumi:"name"`
	DeviceGroupType string `pulumi:"deviceGroupType"`
	DeviceModel     string `pulumi:"deviceModel"`
	OsType          string `pulumi:"osType"`
	OsVersion       string `pulumi:"osVersion"`
	Description     string `pulumi:"description"`
	OwnerUserId     int    `pulumi:"ownerUserId"`
	OwnerName       string `pulumi:"ownerName"`
	Hostname        string `pulumi:"hostname"`
}

type GetDevice struct{}

func (f *GetDevice) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a device by ID, name, model, owner, or OS attributes.")
}

func (a *GetDeviceArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Id, "The ID of the device to look up.")
	ann.Describe(&a.Name, "The name of the device to look up.")
	ann.Describe(&a.DeviceModel, "The device model to filter by.")
	ann.Describe(&a.OwnerName, "The owner name to filter by.")
	ann.Describe(&a.OsType, "The OS type to filter by.")
	ann.Describe(&a.OsVersion, "The OS version to filter by.")
}

func (r *GetDeviceResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the device.")
	a.Describe(&r.Name, "The name of the device.")
	a.Describe(&r.DeviceGroupType, "The device group type.")
	a.Describe(&r.DeviceModel, "The model of the device.")
	a.Describe(&r.OsType, "The OS type of the device.")
	a.Describe(&r.OsVersion, "The OS version of the device.")
	a.Describe(&r.Description, "The description of the device.")
	a.Describe(&r.OwnerUserId, "The user ID of the device owner.")
	a.Describe(&r.OwnerName, "The name of the device owner.")
	a.Describe(&r.Hostname, "The hostname of the device.")
}

func (*GetDevice) Invoke(ctx context.Context, req infer.FunctionRequest[GetDeviceArgs]) (infer.FunctionResponse[GetDeviceResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetDeviceResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	var resp *devicegroups.Devices
	if req.Input.Id != nil && *req.Input.Id != 0 {
		r, err := devicegroups.GetDevicesByID(ctx, svc, *req.Input.Id)
		if err != nil {
			return infer.FunctionResponse[GetDeviceResult]{}, err
		}
		resp = r
	}
	if resp == nil && req.Input.Name != nil && *req.Input.Name != "" {
		r, err := devicegroups.GetDevicesByName(ctx, svc, *req.Input.Name)
		if err != nil {
			return infer.FunctionResponse[GetDeviceResult]{}, err
		}
		resp = r
	}
	if resp == nil && req.Input.DeviceModel != nil && *req.Input.DeviceModel != "" {
		r, err := devicegroups.GetDevicesByModel(ctx, svc, *req.Input.DeviceModel)
		if err != nil {
			return infer.FunctionResponse[GetDeviceResult]{}, err
		}
		resp = r
	}
	if resp == nil && req.Input.OwnerName != nil && *req.Input.OwnerName != "" {
		r, err := devicegroups.GetDevicesByOwner(ctx, svc, *req.Input.OwnerName)
		if err != nil {
			return infer.FunctionResponse[GetDeviceResult]{}, err
		}
		resp = r
	}
	if resp == nil && req.Input.OsType != nil && *req.Input.OsType != "" {
		r, err := devicegroups.GetDevicesByOSType(ctx, svc, *req.Input.OsType)
		if err != nil {
			return infer.FunctionResponse[GetDeviceResult]{}, err
		}
		resp = r
	}
	if resp == nil && req.Input.OsVersion != nil && *req.Input.OsVersion != "" {
		r, err := devicegroups.GetDevicesByOSVersion(ctx, svc, *req.Input.OsVersion)
		if err != nil {
			return infer.FunctionResponse[GetDeviceResult]{}, err
		}
		resp = r
	}
	if resp == nil && ptrToString(req.Input.Name) == "" && ptrToString(req.Input.DeviceModel) == "" &&
		ptrToString(req.Input.OwnerName) == "" && ptrToString(req.Input.OsType) == "" && ptrToString(req.Input.OsVersion) == "" {
		all, err := devicegroups.GetAllDevices(ctx, svc)
		if err != nil {
			return infer.FunctionResponse[GetDeviceResult]{}, err
		}
		if len(all) > 0 {
			resp = &all[0]
		}
	}
	if resp == nil {
		return infer.FunctionResponse[GetDeviceResult]{}, fmt.Errorf("couldn't find any device with the provided attributes")
	}

	return infer.FunctionResponse[GetDeviceResult]{Output: GetDeviceResult{
		Id:              resp.ID,
		Name:            resp.Name,
		DeviceGroupType: resp.DeviceGroupType,
		DeviceModel:     resp.DeviceModel,
		OsType:          resp.OSType,
		OsVersion:       resp.OSVersion,
		Description:     resp.Description,
		OwnerUserId:     resp.OwnerUserId,
		OwnerName:       resp.OwnerName,
		Hostname:        resp.HostName,
	}}, nil
}
