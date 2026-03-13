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

// Package provider implements the Extranet resource.
// Adopted from terraform-provider-zia resource_zia_extranet.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/trafficforwarding/extranet"
)

// Extranet implements the zia:index:Extranet resource.
type Extranet struct{}

// ExtranetDnsListInput is a DNS server entry for an extranet.
type ExtranetDnsListInput struct {
	Id                 *int    `pulumi:"id,optional"`
	Name               *string `pulumi:"name,optional"`
	PrimaryDnsServer   *string `pulumi:"primaryDnsServer,optional"`
	SecondaryDnsServer *string `pulumi:"secondaryDnsServer,optional"`
	UseAsDefault       *bool   `pulumi:"useAsDefault,optional"`
}

// ExtranetIpPoolListInput is an IP pool entry for an extranet.
type ExtranetIpPoolListInput struct {
	Id           *int    `pulumi:"id,optional"`
	Name         *string `pulumi:"name,optional"`
	IpStart      *string `pulumi:"ipStart,optional"`
	IpEnd        *string `pulumi:"ipEnd,optional"`
	UseAsDefault *bool   `pulumi:"useAsDefault,optional"`
}

// ExtranetArgs are the inputs.
type ExtranetArgs struct {
	Name               *string                   `pulumi:"name,optional"`
	Description        *string                   `pulumi:"description,optional"`
	ExtranetDnsList    []ExtranetDnsListInput    `pulumi:"extranetDnsList,optional"`
	ExtranetIpPoolList []ExtranetIpPoolListInput `pulumi:"extranetIpPoolList,optional"`
}

// ExtranetState is the persisted state.
type ExtranetState struct {
	ExtranetArgs
	ExtranetId *int `pulumi:"extranetId"`
}

func extranetToAPI(args ExtranetArgs, id int) extranet.Extranet {
	out := extranet.Extranet{
		ID:          id,
		Name:        ptrToString(args.Name),
		Description: ptrToString(args.Description),
	}
	if len(args.ExtranetDnsList) > 0 {
		out.ExtranetDNSList = make([]extranet.ExtranetDNSList, 0, len(args.ExtranetDnsList))
		for _, d := range args.ExtranetDnsList {
			entry := extranet.ExtranetDNSList{
				Name:               ptrToString(d.Name),
				PrimaryDNSServer:   ptrToString(d.PrimaryDnsServer),
				SecondaryDNSServer: ptrToString(d.SecondaryDnsServer),
				UseAsDefault:       ptrToBool(d.UseAsDefault),
			}
			if d.Id != nil && *d.Id > 0 {
				entry.ID = *d.Id
			}
			out.ExtranetDNSList = append(out.ExtranetDNSList, entry)
		}
	}
	if len(args.ExtranetIpPoolList) > 0 {
		out.ExtranetIpPoolList = make([]extranet.ExtranetPoolList, 0, len(args.ExtranetIpPoolList))
		for _, p := range args.ExtranetIpPoolList {
			entry := extranet.ExtranetPoolList{
				Name:         ptrToString(p.Name),
				IPStart:      ptrToString(p.IpStart),
				IPEnd:        ptrToString(p.IpEnd),
				UseAsDefault: ptrToBool(p.UseAsDefault),
			}
			if p.Id != nil && *p.Id > 0 {
				entry.ID = *p.Id
			}
			out.ExtranetIpPoolList = append(out.ExtranetIpPoolList, entry)
		}
	}
	return out
}

func (Extranet) Create(ctx context.Context, req infer.CreateRequest[ExtranetArgs]) (infer.CreateResponse[ExtranetState], error) {
	if req.DryRun {
		s := ExtranetState{ExtranetArgs: req.Inputs, ExtranetId: intPtr(0)}
		return infer.CreateResponse[ExtranetState]{ID: "preview", Output: s}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[ExtranetState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := extranetToAPI(req.Inputs, 0)
	resp, _, err := extranet.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[ExtranetState]{}, err
	}
	log.Printf("[INFO] Created ZIA extranet. ID: %v", resp.ID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[ExtranetState]{}, activationErr
		}
	}

	state := buildExtranetStateFromAPI(resp)
	return infer.CreateResponse[ExtranetState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (Extranet) Read(ctx context.Context, req infer.ReadRequest[ExtranetArgs, ExtranetState]) (infer.ReadResponse[ExtranetArgs, ExtranetState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[ExtranetArgs, ExtranetState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.ExtranetId != nil {
		id = *req.State.ExtranetId
	}
	if id == 0 {
		return infer.ReadResponse[ExtranetArgs, ExtranetState]{}, fmt.Errorf("no extranet id in state")
	}

	resp, err := extranet.Get(ctx, service, id)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[ExtranetArgs, ExtranetState]{ID: ""}, nil
		}
		return infer.ReadResponse[ExtranetArgs, ExtranetState]{}, err
	}

	state := buildExtranetStateFromAPI(resp)
	return infer.ReadResponse[ExtranetArgs, ExtranetState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: state.ExtranetArgs,
		State:  state,
	}, nil
}

func (Extranet) Update(ctx context.Context, req infer.UpdateRequest[ExtranetArgs, ExtranetState]) (infer.UpdateResponse[ExtranetState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[ExtranetState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.ExtranetId != nil {
		id = *req.State.ExtranetId
	}
	if id == 0 {
		return infer.UpdateResponse[ExtranetState]{}, fmt.Errorf("no extranet id in state")
	}

	if _, err := extranet.Get(ctx, service, id); err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[ExtranetState]{}, nil
		}
		return infer.UpdateResponse[ExtranetState]{}, err
	}

	merged := mergeExtranetNestedIDs(req.Inputs, req.State.ExtranetArgs)
	apiReq := extranetToAPI(merged, id)
	resp, err := extranet.Update(ctx, service, id, &apiReq)
	if err != nil {
		return infer.UpdateResponse[ExtranetState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[ExtranetState]{}, activationErr
		}
	}

	state := buildExtranetStateFromAPI(resp)
	return infer.UpdateResponse[ExtranetState]{Output: state}, nil
}

func (Extranet) Delete(ctx context.Context, req infer.DeleteRequest[ExtranetState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.ExtranetId != nil {
		id = *req.State.ExtranetId
	}
	if id != 0 {
		if _, err := extranet.Delete(ctx, service, id); err != nil {
			if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.DeleteResponse{}, nil
			}
			return infer.DeleteResponse{}, err
		}
		log.Printf("[INFO] ZIA extranet deleted")
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

// buildExtranetStateFromAPI constructs a full ExtranetState from an API response,
// including nested item IDs so that subsequent Updates can reference them.
func buildExtranetStateFromAPI(resp *extranet.Extranet) ExtranetState {
	dnsList := make([]ExtranetDnsListInput, 0, len(resp.ExtranetDNSList))
	for _, d := range resp.ExtranetDNSList {
		entry := ExtranetDnsListInput{
			Name:               stringPtr(d.Name),
			PrimaryDnsServer:   stringPtr(d.PrimaryDNSServer),
			SecondaryDnsServer: stringPtr(d.SecondaryDNSServer),
		}
		if d.ID > 0 {
			entry.Id = intPtr(d.ID)
		}
		if d.UseAsDefault {
			entry.UseAsDefault = boolPtr(true)
		}
		dnsList = append(dnsList, entry)
	}
	ipPoolList := make([]ExtranetIpPoolListInput, 0, len(resp.ExtranetIpPoolList))
	for _, p := range resp.ExtranetIpPoolList {
		entry := ExtranetIpPoolListInput{
			Name:    stringPtr(p.Name),
			IpStart: stringPtr(p.IPStart),
			IpEnd:   stringPtr(p.IPEnd),
		}
		if p.ID > 0 {
			entry.Id = intPtr(p.ID)
		}
		if p.UseAsDefault {
			entry.UseAsDefault = boolPtr(true)
		}
		ipPoolList = append(ipPoolList, entry)
	}
	args := ExtranetArgs{
		Name:               stringPtr(resp.Name),
		Description:        stringPtr(resp.Description),
		ExtranetDnsList:    dnsList,
		ExtranetIpPoolList: ipPoolList,
	}
	return ExtranetState{ExtranetArgs: args, ExtranetId: &resp.ID}
}

// mergeExtranetNestedIDs copies nested item IDs from the old state into the new
// inputs (matched by name) so the API can identify existing items on update.
func mergeExtranetNestedIDs(inputs, state ExtranetArgs) ExtranetArgs {
	dnsIDByName := map[string]int{}
	for _, d := range state.ExtranetDnsList {
		if d.Id != nil && d.Name != nil {
			dnsIDByName[*d.Name] = *d.Id
		}
	}
	poolIDByName := map[string]int{}
	for _, p := range state.ExtranetIpPoolList {
		if p.Id != nil && p.Name != nil {
			poolIDByName[*p.Name] = *p.Id
		}
	}
	for i := range inputs.ExtranetDnsList {
		if inputs.ExtranetDnsList[i].Name != nil {
			if nid, ok := dnsIDByName[*inputs.ExtranetDnsList[i].Name]; ok {
				inputs.ExtranetDnsList[i].Id = intPtr(nid)
			}
		}
	}
	for i := range inputs.ExtranetIpPoolList {
		if inputs.ExtranetIpPoolList[i].Name != nil {
			if nid, ok := poolIDByName[*inputs.ExtranetIpPoolList[i].Name]; ok {
				inputs.ExtranetIpPoolList[i].Id = intPtr(nid)
			}
		}
	}
	return inputs
}

func (Extranet) Diff(ctx context.Context, req infer.DiffRequest[ExtranetArgs, ExtranetState]) (infer.DiffResponse, error) {
	diff, hasChanges := diffSkippingUnsetInputs(req.State.ExtranetArgs, req.Inputs)
	return infer.DiffResponse{HasChanges: hasChanges, DetailedDiff: diff}, nil
}

func (Extranet) Annotate(a infer.Annotator) {
	describeResource(a, &Extranet{}, `The zia.Extranet resource manages extranet configurations in the Zscaler Internet Access (ZIA) cloud.
Extranets define DNS and IP pool settings for traffic forwarding.

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Extranet

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.Extranet("example", {
    name: "Example Extranet",
    description: "Managed by Pulumi",
    extranetDnsList: [{
        name: "Primary DNS",
        primaryDnsServer: "8.8.8.8",
        secondaryDnsServer: "8.8.4.4",
    }],
    extranetIpPoolList: [{
        name: "IP Pool 1",
        ipStart: "10.0.0.1",
        ipEnd: "10.0.0.254",
    }],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.Extranet("example",
    name="Example Extranet",
    description="Managed by Pulumi",
    extranet_dns_list=[{
        "name": "Primary DNS",
        "primary_dns_server": "8.8.8.8",
        "secondary_dns_server": "8.8.4.4",
    }],
    extranet_ip_pool_list=[{
        "name": "IP Pool 1",
        "ip_start": "10.0.0.1",
        "ip_end": "10.0.0.254",
    }],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:Extranet
    properties:
      name: Example Extranet
      description: Managed by Pulumi
      extranetDnsList:
        - name: Primary DNS
          primaryDnsServer: "8.8.8.8"
          secondaryDnsServer: "8.8.4.4"
      extranetIpPoolList:
        - name: IP Pool 1
          ipStart: "10.0.0.1"
          ipEnd: "10.0.0.254"
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing extranet can be imported using its ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:Extranet example 12345
`+tripleBacktick("")+`
`)
}

func (a *ExtranetArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "Name of the extranet.")
	ann.Describe(&a.Description, "Description of the extranet.")
	ann.Describe(&a.ExtranetDnsList, "List of DNS server entries for the extranet.")
	ann.Describe(&a.ExtranetIpPoolList, "List of IP pool entries for the extranet.")
}

func (s *ExtranetState) Annotate(a infer.Annotator) {
	a.Describe(&s.ExtranetId, "The unique identifier for the extranet assigned by the ZIA cloud.")
}

var _ infer.CustomResource[ExtranetArgs, ExtranetState] = Extranet{}
