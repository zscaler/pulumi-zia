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

// Package provider implements the FW Network Service resource.
// Adopted from terraform-provider-zia resource_zia_fw_filtering_network_services.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/common"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/firewallpolicies/filteringrules"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/firewallpolicies/networkservices"
)

// FwNetworkService implements the zia:index:FwNetworkService resource.
type FwNetworkService struct{}

// NetworkPortInput is a port range (start, end).
type NetworkPortInput struct {
	Start *int `pulumi:"start,optional"`
	End   *int `pulumi:"end,optional"`
}

// FwNetworkServiceArgs are the inputs.
type FwNetworkServiceArgs struct {
	Name          *string             `pulumi:"name,optional"`
	Tag           *string             `pulumi:"tag,optional"`
	Description   *string             `pulumi:"description,optional"`
	SrcTcpPorts   []NetworkPortInput  `pulumi:"srcTcpPorts,optional"`
	DestTcpPorts  []NetworkPortInput  `pulumi:"destTcpPorts,optional"`
	SrcUdpPorts   []NetworkPortInput  `pulumi:"srcUdpPorts,optional"`
	DestUdpPorts  []NetworkPortInput  `pulumi:"destUdpPorts,optional"`
	Type          *string             `pulumi:"type,optional"`
	IsNameL10nTag *bool                `pulumi:"isNameL10nTag,optional"`
}

// FwNetworkServiceState is the persisted state.
type FwNetworkServiceState struct {
	FwNetworkServiceArgs
	NetworkServiceId *int `pulumi:"networkServiceId"`
}

func networkPortsToAPI(list []NetworkPortInput) []networkservices.NetworkPorts {
	if len(list) == 0 {
		return nil
	}
	result := make([]networkservices.NetworkPorts, len(list))
	for i, p := range list {
		start, end := 0, 0
		if p.Start != nil {
			start = *p.Start
		}
		if p.End != nil {
			end = *p.End
		}
		result[i] = networkservices.NetworkPorts{Start: start, End: end}
	}
	return result
}

func networkPortsFromAPI(list []networkservices.NetworkPorts) []NetworkPortInput {
	if len(list) == 0 {
		return nil
	}
	result := make([]NetworkPortInput, len(list))
	for i, p := range list {
		result[i] = NetworkPortInput{Start: intPtr(p.Start), End: intPtr(p.End)}
	}
	return result
}

func fwNetworkServiceToAPI(args FwNetworkServiceArgs, id int) networkservices.NetworkServices {
	return networkservices.NetworkServices{
		ID:            id,
		Name:          ptrToString(args.Name),
		Tag:           ptrToString(args.Tag),
		Description:   ptrToString(args.Description),
		Type:          ptrToString(args.Type),
		IsNameL10nTag: ptrToBool(args.IsNameL10nTag),
		SrcTCPPorts:   networkPortsToAPI(args.SrcTcpPorts),
		DestTCPPorts:  networkPortsToAPI(args.DestTcpPorts),
		SrcUDPPorts:   networkPortsToAPI(args.SrcUdpPorts),
		DestUDPPorts:  networkPortsToAPI(args.DestUdpPorts),
	}
}

func (FwNetworkService) Create(ctx context.Context, req infer.CreateRequest[FwNetworkServiceArgs]) (infer.CreateResponse[FwNetworkServiceState], error) {
	if req.DryRun {
		s := FwNetworkServiceState{FwNetworkServiceArgs: req.Inputs, NetworkServiceId: intPtr(0)}
		return infer.CreateResponse[FwNetworkServiceState]{ID: "preview", Output: s}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[FwNetworkServiceState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := fwNetworkServiceToAPI(req.Inputs, 0)
	resp, err := networkservices.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[FwNetworkServiceState]{}, err
	}
	log.Printf("[INFO] Created ZIA network service. ID: %v", resp.ID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[FwNetworkServiceState]{}, activationErr
		}
	}

	state := FwNetworkServiceState{
		FwNetworkServiceArgs: req.Inputs,
		NetworkServiceId:    &resp.ID,
	}
	return infer.CreateResponse[FwNetworkServiceState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (FwNetworkService) Read(ctx context.Context, req infer.ReadRequest[FwNetworkServiceArgs, FwNetworkServiceState]) (infer.ReadResponse[FwNetworkServiceArgs, FwNetworkServiceState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[FwNetworkServiceArgs, FwNetworkServiceState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.NetworkServiceId != nil {
		id = *req.State.NetworkServiceId
	}
	if id == 0 {
		return infer.ReadResponse[FwNetworkServiceArgs, FwNetworkServiceState]{}, fmt.Errorf("no network service id in state")
	}

	resp, err := networkservices.Get(ctx, service, id)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[FwNetworkServiceArgs, FwNetworkServiceState]{ID: ""}, nil
		}
		return infer.ReadResponse[FwNetworkServiceArgs, FwNetworkServiceState]{}, err
	}

	args := FwNetworkServiceArgs{
		Name:          stringPtr(resp.Name),
		Tag:           stringPtr(resp.Tag),
		Description:   stringPtr(resp.Description),
		Type:          stringPtr(resp.Type),
		IsNameL10nTag: boolPtr(resp.IsNameL10nTag),
		SrcTcpPorts:   networkPortsFromAPI(resp.SrcTCPPorts),
		DestTcpPorts:  networkPortsFromAPI(resp.DestTCPPorts),
		SrcUdpPorts:   networkPortsFromAPI(resp.SrcUDPPorts),
		DestUdpPorts:  networkPortsFromAPI(resp.DestUDPPorts),
	}
	state := FwNetworkServiceState{
		FwNetworkServiceArgs: args,
		NetworkServiceId:     &resp.ID,
	}
	return infer.ReadResponse[FwNetworkServiceArgs, FwNetworkServiceState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (FwNetworkService) Update(ctx context.Context, req infer.UpdateRequest[FwNetworkServiceArgs, FwNetworkServiceState]) (infer.UpdateResponse[FwNetworkServiceState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[FwNetworkServiceState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.NetworkServiceId != nil {
		id = *req.State.NetworkServiceId
	}
	if id == 0 {
		return infer.UpdateResponse[FwNetworkServiceState]{}, fmt.Errorf("no network service id in state")
	}

	if _, err := networkservices.Get(ctx, service, id); err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[FwNetworkServiceState]{}, nil
		}
		return infer.UpdateResponse[FwNetworkServiceState]{}, err
	}

	apiReq := fwNetworkServiceToAPI(req.Inputs, id)
	if _, _, err := networkservices.Update(ctx, service, id, &apiReq); err != nil {
		return infer.UpdateResponse[FwNetworkServiceState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[FwNetworkServiceState]{}, activationErr
		}
	}

	state := FwNetworkServiceState{
		FwNetworkServiceArgs: req.Inputs,
		NetworkServiceId:     &id,
	}
	return infer.UpdateResponse[FwNetworkServiceState]{Output: state}, nil
}

func (FwNetworkService) Delete(ctx context.Context, req infer.DeleteRequest[FwNetworkServiceState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	service := client.Service

	id := 0
	if req.State.NetworkServiceId != nil {
		id = *req.State.NetworkServiceId
	}
	if id != 0 {
		if err := detachFromFilteringRules(ctx, client, id, "NwServices",
			func(r *filteringrules.FirewallFilteringRules) []common.IDNameExtensions { return r.NwServices },
			func(r *filteringrules.FirewallFilteringRules, ids []common.IDNameExtensions) { r.NwServices = ids }); err != nil {
			return infer.DeleteResponse{}, err
		}
		if _, err := networkservices.Delete(ctx, service, id); err != nil {
			return infer.DeleteResponse{}, err
		}
		log.Printf("[INFO] ZIA network service deleted")
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

func (FwNetworkService) Annotate(a infer.Annotator) {
	describeResource(a, &FwNetworkService{}, `The zia_fw_network_service resource manages firewall network services in the Zscaler Internet Access (ZIA) cloud service. Network services define the TCP/UDP port ranges used in firewall filtering rules.

For more information, see the [ZIA Firewall Policies documentation](https://help.zscaler.com/zia/firewall-policies).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Network Service

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.FwNetworkService("example", {
    name: "Example Network Service",
    description: "Custom network service",
    destTcpPorts: [
        { start: 443, end: 443 },
        { start: 8080, end: 8090 },
    ],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.FwNetworkService("example",
    name="Example Network Service",
    description="Custom network service",
    dest_tcp_ports=[
        {"start": 443, "end": 443},
        {"start": 8080, "end": 8090},
    ],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:FwNetworkService
    properties:
      name: Example Network Service
      description: Custom network service
      destTcpPorts:
        - start: 443
          end: 443
        - start: 8080
          end: 8090
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing network service can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:FwNetworkService example 12345
`+tripleBacktick("")+`
`)
}

func (a *FwNetworkServiceArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the network service.")
	ann.Describe(&a.Tag, "The tag associated with the network service.")
	ann.Describe(&a.Description, "Additional information about the network service.")
	ann.Describe(&a.SrcTcpPorts, "Source TCP port ranges. Each entry specifies a start and end port.")
	ann.Describe(&a.DestTcpPorts, "Destination TCP port ranges. Each entry specifies a start and end port.")
	ann.Describe(&a.SrcUdpPorts, "Source UDP port ranges. Each entry specifies a start and end port.")
	ann.Describe(&a.DestUdpPorts, "Destination UDP port ranges. Each entry specifies a start and end port.")
	ann.Describe(&a.Type, "The network service type. Valid values: `STANDARD`, `PREDEFINED`, `CUSTOM`.")
	ann.Describe(&a.IsNameL10nTag, "Indicates whether the name is a localization tag.")
}

func (s *FwNetworkServiceState) Annotate(a infer.Annotator) {
	a.Describe(&s.NetworkServiceId, "The system-generated ID of the network service.")
}

var _ infer.CustomResource[FwNetworkServiceArgs, FwNetworkServiceState] = FwNetworkService{}
