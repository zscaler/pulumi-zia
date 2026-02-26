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

// Package provider implements the SubCloud resource.
// Adopted from terraform-provider-zia resource_zia_sub_cloud.go.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/common"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/trafficforwarding/sub_clouds"

	"github.com/pulumi/pulumi-go-provider/infer"
)

// SubCloud implements the zia:index:SubCloud resource.
// Create and Update both call sub_clouds.Update. Delete is a no-op.
type SubCloud struct{}

// SubCloudExclusionDatacenterInput is the datacenter block within exclusions.
type SubCloudExclusionDatacenterInput struct {
	Id      int     `pulumi:"resourceId"`
	Name    *string `pulumi:"name,optional"`
	Country *string `pulumi:"country,optional"`
}

// SubCloudExclusionInput is a single exclusion.
type SubCloudExclusionInput struct {
	Datacenter SubCloudExclusionDatacenterInput `pulumi:"datacenter"`
	Country    string                           `pulumi:"country"`
	EndTime    *int                             `pulumi:"endTime,optional"`
	EndTimeUtc *string                          `pulumi:"endTimeUtc,optional"`
}

// SubCloudDcOutput is a datacenter in the dcs computed list.
type SubCloudDcOutput struct {
	Id      int    `pulumi:"resourceId"`
	Name    string `pulumi:"name"`
	Country string `pulumi:"country"`
}

// SubCloudArgs are the inputs.
type SubCloudArgs struct {
	CloudId    int                      `pulumi:"cloudId"`
	Exclusions []SubCloudExclusionInput `pulumi:"exclusions,optional"`
}

// SubCloudState is the persisted state.
type SubCloudState struct {
	SubCloudArgs
	Id   string             `pulumi:"resourceId"` // Pulumi reserves "id" for resource identifier
	Name *string            `pulumi:"name,optional"`
	Dcs  []SubCloudDcOutput `pulumi:"dcs,optional"`
}

func expandSubCloudExclusions(in []SubCloudExclusionInput) ([]sub_clouds.Exclusions, error) {
	if len(in) == 0 {
		return nil, nil
	}
	out := make([]sub_clouds.Exclusions, 0, len(in))
	for _, e := range in {
		hasEpoch := e.EndTime != nil && *e.EndTime != 0
		epochVal := 0
		if e.EndTime != nil {
			epochVal = *e.EndTime
		}
		utcStr := ptrToString(e.EndTimeUtc)
		endTime, err := ResolveExclusionTimeEpoch(hasEpoch, epochVal, utcStr)
		if err != nil {
			return nil, fmt.Errorf("exclusion end_time: %w", err)
		}
		exc := sub_clouds.Exclusions{
			Country: e.Country,
			EndTime: endTime,
			Datacenter: &common.IDNameExtensions{
				ID:   e.Datacenter.Id,
				Name: ptrToString(e.Datacenter.Name),
			},
		}
		if e.Datacenter.Country != nil && *e.Datacenter.Country != "" {
			exc.Datacenter.Extensions = map[string]interface{}{"country": *e.Datacenter.Country}
		}
		out = append(out, exc)
	}
	return out, nil
}

func flattenSubCloudDCs(dcs []sub_clouds.DCs) []SubCloudDcOutput {
	if len(dcs) == 0 {
		return nil
	}
	out := make([]SubCloudDcOutput, len(dcs))
	for i, dc := range dcs {
		out[i] = SubCloudDcOutput{Id: dc.ID, Name: dc.Name, Country: dc.Country}
	}
	return out
}

func flattenSubCloudExclusionsForResource(exclusions []sub_clouds.Exclusions) []SubCloudExclusionInput {
	if len(exclusions) == 0 {
		return nil
	}
	out := make([]SubCloudExclusionInput, 0, len(exclusions))
	for _, e := range exclusions {
		dc := SubCloudExclusionDatacenterInput{Id: 0}
		if e.Datacenter != nil {
			dc.Id = e.Datacenter.ID
			dc.Name = stringPtr(e.Datacenter.Name)
			if e.Datacenter.Extensions != nil {
				if c, ok := e.Datacenter.Extensions["country"].(string); ok {
					dc.Country = &c
				}
			}
		}
		exc := SubCloudExclusionInput{
			Datacenter: dc,
			Country:    e.Country,
			EndTime:    intPtr(e.EndTime),
			EndTimeUtc: stringPtr(FormatExclusionTimeUTC(e.EndTime)),
		}
		out = append(out, exc)
	}
	return out
}

func (SubCloud) Create(ctx context.Context, req infer.CreateRequest[SubCloudArgs]) (infer.CreateResponse[SubCloudState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[SubCloudState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	exclusions, err := expandSubCloudExclusions(req.Inputs.Exclusions)
	if err != nil {
		return infer.CreateResponse[SubCloudState]{}, err
	}
	apiReq := &sub_clouds.SubClouds{
		ID:         req.Inputs.CloudId,
		Exclusions: exclusions,
	}
	resp, _, err := sub_clouds.Update(ctx, service, req.Inputs.CloudId, apiReq)
	if err != nil {
		return infer.CreateResponse[SubCloudState]{}, fmt.Errorf("error creating/updating subcloud: %w", err)
	}
	state := SubCloudState{
		SubCloudArgs: req.Inputs,
		Id:           strconv.Itoa(resp.ID),
		Name:         stringPtr(resp.Name),
		Dcs:          flattenSubCloudDCs(resp.Dcs),
	}
	if len(resp.Exclusions) > 0 {
		state.Exclusions = flattenSubCloudExclusionsForResource(resp.Exclusions)
	}
	return infer.CreateResponse[SubCloudState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (SubCloud) Read(ctx context.Context, req infer.ReadRequest[SubCloudArgs, SubCloudState]) (infer.ReadResponse[SubCloudArgs, SubCloudState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[SubCloudArgs, SubCloudState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.ReadResponse[SubCloudArgs, SubCloudState]{}, fmt.Errorf("invalid subcloud ID: %s", req.ID)
	}
	all, err := sub_clouds.GetAll(ctx, service)
	if err != nil {
		return infer.ReadResponse[SubCloudArgs, SubCloudState]{}, err
	}
	var resp *sub_clouds.SubClouds
	for i := range all {
		if all[i].ID == id {
			resp = &all[i]
			break
		}
	}
	if resp == nil {
		return infer.ReadResponse[SubCloudArgs, SubCloudState]{}, fmt.Errorf("subcloud not found")
	}
	state := SubCloudState{
		SubCloudArgs: SubCloudArgs{
			CloudId:    resp.ID,
			Exclusions: flattenSubCloudExclusionsForResource(resp.Exclusions),
		},
		Id:   strconv.Itoa(resp.ID),
		Name: stringPtr(resp.Name),
		Dcs:  flattenSubCloudDCs(resp.Dcs),
	}
	return infer.ReadResponse[SubCloudArgs, SubCloudState]{
		ID:     req.ID,
		Inputs: state.SubCloudArgs,
		State:  state,
	}, nil
}

func (SubCloud) Update(ctx context.Context, req infer.UpdateRequest[SubCloudArgs, SubCloudState]) (infer.UpdateResponse[SubCloudState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[SubCloudState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	exclusions, err := expandSubCloudExclusions(req.Inputs.Exclusions)
	if err != nil {
		return infer.UpdateResponse[SubCloudState]{}, err
	}
	apiReq := &sub_clouds.SubClouds{
		ID:         req.Inputs.CloudId,
		Exclusions: exclusions,
	}
	resp, _, err := sub_clouds.Update(ctx, service, req.Inputs.CloudId, apiReq)
	if err != nil {
		return infer.UpdateResponse[SubCloudState]{}, fmt.Errorf("error updating subcloud: %w", err)
	}
	state := SubCloudState{
		SubCloudArgs: req.Inputs,
		Id:           strconv.Itoa(resp.ID),
		Name:         stringPtr(resp.Name),
		Dcs:          flattenSubCloudDCs(resp.Dcs),
	}
	if len(resp.Exclusions) > 0 {
		state.Exclusions = flattenSubCloudExclusionsForResource(resp.Exclusions)
	}
	return infer.UpdateResponse[SubCloudState]{Output: state}, nil
}

func (SubCloud) Delete(ctx context.Context, req infer.DeleteRequest[SubCloudState]) (infer.DeleteResponse, error) {
	// No-op: subcloud cannot be deleted via API in the same way; Terraform uses resourceFuncNoOp.
	return infer.DeleteResponse{}, nil
}

func (SubCloud) Annotate(a infer.Annotator) {
	describeResource(a, &SubCloud{}, `The zia.SubCloud resource manages sub-cloud configurations in the Zscaler Internet Access (ZIA) cloud.
Sub-clouds represent regional cloud instances with associated datacenters and exclusion rules.
Create and update both use the same API operation. Deleting the Pulumi resource does not remove
the underlying sub-cloud configuration.

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Sub-Cloud

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.SubCloud("example", {
    cloudId: 1,
    exclusions: [{
        datacenter: {
            resourceId: 100,
            name: "US-East",
            country: "US",
        },
        country: "US",
    }],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.SubCloud("example",
    cloud_id=1,
    exclusions=[{
        "datacenter": {
            "resource_id": 100,
            "name": "US-East",
            "country": "US",
        },
        "country": "US",
    }],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:SubCloud
    properties:
      cloudId: 1
      exclusions:
        - datacenter:
            resourceId: 100
            name: US-East
            country: US
          country: US
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

This resource uses a no-op delete. Import is not typically applicable for sub-cloud resources.
`)
}

func (a *SubCloudArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.CloudId, "The ID of the cloud to which this sub-cloud belongs.")
	ann.Describe(&a.Exclusions, "List of datacenter exclusions for the sub-cloud.")
}

func (s *SubCloudState) Annotate(a infer.Annotator) {
	a.Describe(&s.Id, "The resource ID of the sub-cloud.")
	a.Describe(&s.Name, "The name of the sub-cloud.")
	a.Describe(&s.Dcs, "List of datacenters associated with the sub-cloud.")
}

var _ infer.CustomResource[SubCloudArgs, SubCloudState] = SubCloud{}
