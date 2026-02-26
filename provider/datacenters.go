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

// Package provider implements the getDatacenters invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_datacenters.go.

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/trafficforwarding/dc_exclusions"
)

type GetDatacentersArgs struct {
	DatacenterId *int    `pulumi:"datacenterId,optional"`
	Name         *string `pulumi:"name,optional"`
	City         *string `pulumi:"city,optional"`
}

type DatacenterInfo struct {
	ID                int     `pulumi:"id"`
	Name              string  `pulumi:"name"`
	Provider          string  `pulumi:"provider"`
	City              string  `pulumi:"city"`
	Timezone          string  `pulumi:"timezone"`
	Lat               int     `pulumi:"lat"`
	Longi             int     `pulumi:"longi"`
	Latitude          float64 `pulumi:"latitude"`
	Longitude         float64 `pulumi:"longitude"`
	GovOnly           bool    `pulumi:"govOnly"`
	ThirdPartyCloud   bool    `pulumi:"thirdPartyCloud"`
	UploadBandwidth   int     `pulumi:"uploadBandwidth"`
	DownloadBandwidth int     `pulumi:"downloadBandwidth"`
	OwnedByCustomer   bool    `pulumi:"ownedByCustomer"`
	ManagedBcp        bool    `pulumi:"managedBcp"`
	DontPublish       bool    `pulumi:"dontPublish"`
	DontProvision     bool    `pulumi:"dontProvision"`
	NotReadyForUse    bool    `pulumi:"notReadyForUse"`
	ForFutureUse      bool    `pulumi:"forFutureUse"`
	RegionalSurcharge bool    `pulumi:"regionalSurcharge"`
	CreateTime        int     `pulumi:"createTime"`
	LastModifiedTime  int     `pulumi:"lastModifiedTime"`
	Virtual           bool    `pulumi:"virtual"`
}

type GetDatacentersResult struct {
	DatacenterId *int             `pulumi:"datacenterId,optional"`
	Name         *string          `pulumi:"name,optional"`
	City         *string          `pulumi:"city,optional"`
	Datacenters  []DatacenterInfo `pulumi:"datacenters"`
}

type GetDatacenters struct{}

func (f *GetDatacenters) Annotate(a infer.Annotator) {
	a.Describe(f, "Use the **zia:index/getDatacenters:getDatacenters** data source to retrieve a list of Zscaler data centers. "+
		"Results can be filtered by datacenter ID, name (case-insensitive partial match), or city (case-insensitive partial match).\n\n"+
		"{{% examples %}}\n"+
		"## Example Usage\n\n"+
		"{{% example %}}\n"+
		"### Retrieve All Datacenters\n\n"+
		tripleBacktick("typescript")+"\n"+
		"import * as pulumi from \"@pulumi/pulumi\";\n"+
		"import * as zia from \"@bdzscaler/pulumi-zia\";\n\n"+
		"const all = zia.getDatacenters({});\n"+
		"export const datacenters = all.then(r => r.datacenters);\n"+
		tripleBacktick("")+"\n\n"+
		tripleBacktick("python")+"\n"+
		"import pulumi\n"+
		"import zscaler_pulumi_zia as zia\n\n"+
		"all = zia.get_datacenters()\n"+
		"pulumi.export(\"datacenters\", all.datacenters)\n"+
		tripleBacktick("")+"\n\n"+
		tripleBacktick("go")+"\n"+
		"package main\n\n"+
		"import (\n"+
		"\t\"github.com/pulumi/pulumi/sdk/v3/go/pulumi\"\n"+
		"\t\"github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia\"\n"+
		")\n\n"+
		"func main() {\n"+
		"\tpulumi.Run(func(ctx *pulumi.Context) error {\n"+
		"\t\tall, err := zia.GetDatacenters(ctx, &zia.GetDatacentersArgs{}, nil)\n"+
		"\t\tif err != nil {\n"+
		"\t\t\treturn err\n"+
		"\t\t}\n"+
		"\t\tctx.Export(\"datacenters\", pulumi.ToStringArray(all.Datacenters))\n"+
		"\t\treturn nil\n"+
		"\t})\n"+
		"}\n"+
		tripleBacktick("")+"\n\n"+
		tripleBacktick("yaml")+"\n"+
		"variables:\n"+
		"  all:\n"+
		"    fn::invoke:\n"+
		"      function: zia:getDatacenters\n"+
		"      arguments: {}\n"+
		"outputs:\n"+
		"  datacenters: ${all.datacenters}\n"+
		tripleBacktick("")+"\n"+
		"{{% /example %}}\n\n"+
		"{{% example %}}\n"+
		"### Filter by City\n\n"+
		tripleBacktick("typescript")+"\n"+
		"import * as pulumi from \"@pulumi/pulumi\";\n"+
		"import * as zia from \"@bdzscaler/pulumi-zia\";\n\n"+
		"const dc = zia.getDatacenters({ city: \"San Jose\" });\n"+
		"export const datacenters = dc.then(r => r.datacenters);\n"+
		tripleBacktick("")+"\n\n"+
		tripleBacktick("python")+"\n"+
		"import pulumi\n"+
		"import zscaler_pulumi_zia as zia\n\n"+
		"dc = zia.get_datacenters(city=\"San Jose\")\n"+
		"pulumi.export(\"datacenters\", dc.datacenters)\n"+
		tripleBacktick("")+"\n"+
		"{{% /example %}}\n"+
		"{{% /examples %}}\n\n"+
		"## Import\n\n"+
		"This data source is read-only and does not support import.")
}

func (a *GetDatacentersArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.DatacenterId, "Filter datacenters by ID. When exactly one result is returned, this is set to that datacenter's ID.")
	ann.Describe(&a.Name, "Filter datacenters by name (case-insensitive partial match). When exactly one result is returned, this is set to that datacenter's name.")
	ann.Describe(&a.City, "Filter datacenters by city (case-insensitive partial match). When exactly one result is returned, this is set to that datacenter's city.")
}

func (r *GetDatacentersResult) Annotate(a infer.Annotator) {
	a.Describe(&r.DatacenterId, "The datacenter ID when exactly one result is returned.")
	a.Describe(&r.Name, "The datacenter name when exactly one result is returned.")
	a.Describe(&r.City, "The datacenter city when exactly one result is returned.")
	a.Describe(&r.Datacenters, "List of datacenters matching the filter criteria.")
}

func (*GetDatacenters) Invoke(ctx context.Context, req infer.FunctionRequest[GetDatacentersArgs]) (infer.FunctionResponse[GetDatacentersResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetDatacentersResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	allDatacenters, err := dc_exclusions.GetDatacenters(ctx, svc)
	if err != nil {
		return infer.FunctionResponse[GetDatacentersResult]{}, fmt.Errorf("error getting datacenters: %w", err)
	}

	filtered := allDatacenters
	hasID := req.Input.DatacenterId != nil && *req.Input.DatacenterId != 0
	hasName := req.Input.Name != nil && *req.Input.Name != ""
	hasCity := req.Input.City != nil && *req.Input.City != ""

	if hasID || hasName || hasCity {
		var result []dc_exclusions.Datacenter
		for _, dc := range allDatacenters {
			matched := true

			if hasID && dc.ID != *req.Input.DatacenterId {
				matched = false
			}

			if matched && hasName {
				if !strings.Contains(strings.ToLower(dc.Name), strings.ToLower(*req.Input.Name)) {
					matched = false
				}
			}

			if matched && hasCity {
				if !strings.Contains(strings.ToLower(dc.City), strings.ToLower(*req.Input.City)) {
					matched = false
				}
			}

			if matched {
				result = append(result, dc)
			}
		}
		filtered = result
	}

	out := make([]DatacenterInfo, 0, len(filtered))
	for _, dc := range filtered {
		out = append(out, DatacenterInfo{
			ID:                dc.ID,
			Name:              dc.Name,
			Provider:          dc.Provider,
			City:              dc.City,
			Timezone:          dc.Timezone,
			Lat:               dc.Lat,
			Longi:             dc.Longi,
			Latitude:          dc.Latitude,
			Longitude:         dc.Longitude,
			GovOnly:           dc.GovOnly,
			ThirdPartyCloud:   dc.ThirdPartyCloud,
			UploadBandwidth:   dc.UploadBandwidth,
			DownloadBandwidth: dc.DownloadBandwidth,
			OwnedByCustomer:   dc.OwnedByCustomer,
			ManagedBcp:        dc.ManagedBcp,
			DontPublish:       dc.DontPublish,
			DontProvision:     dc.DontProvision,
			NotReadyForUse:    dc.NotReadyForUse,
			ForFutureUse:      dc.ForFutureUse,
			RegionalSurcharge: dc.RegionalSurcharge,
			CreateTime:        dc.CreateTime,
			LastModifiedTime:  dc.LastModifiedTime,
			Virtual:           dc.Virtual,
		})
	}

	res := GetDatacentersResult{
		Datacenters: out,
	}

	if len(filtered) == 1 {
		res.DatacenterId = &filtered[0].ID
		res.Name = &filtered[0].Name
		res.City = &filtered[0].City
	} else {
		res.DatacenterId = req.Input.DatacenterId
		res.Name = req.Input.Name
		res.City = req.Input.City
	}

	return infer.FunctionResponse[GetDatacentersResult]{Output: res}, nil
}
