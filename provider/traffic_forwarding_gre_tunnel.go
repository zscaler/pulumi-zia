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

// Package provider implements the Traffic Forwarding GRE Tunnel resource.
// Adopted from terraform-provider-zia resource_zia_traffic_forwarding_gre_tunnels.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/trafficforwarding/gretunnels"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/trafficforwarding/virtualipaddress"
)

// TrafficForwardingGreTunnel implements the zia:index:TrafficForwardingGreTunnel resource.
type TrafficForwardingGreTunnel struct{}

// GreTunnelDestVipInput is primary_dest_vip or secondary_dest_vip block.
type GreTunnelDestVipInput struct {
	Id         *int    `pulumi:"id,optional"`
	VirtualIP  *string `pulumi:"virtualIp,optional"`
	Datacenter *string `pulumi:"datacenter,optional"`
}

// GreTunnelDestVipOutput is output for dest vip blocks.
type GreTunnelDestVipOutput struct {
	Id         *int    `pulumi:"id,optional"`
	VirtualIP  *string `pulumi:"virtualIp,optional"`
	Datacenter *string `pulumi:"datacenter,optional"`
}

// TrafficForwardingGreTunnelArgs are the inputs.
type TrafficForwardingGreTunnelArgs struct {
	SourceIP         string                 `pulumi:"sourceIp"`
	WithinCountry    *bool                  `pulumi:"withinCountry,optional"`
	PrimaryDestVip   *GreTunnelDestVipInput `pulumi:"primaryDestVip,optional"`
	SecondaryDestVip *GreTunnelDestVipInput `pulumi:"secondaryDestVip,optional"`
	InternalIpRange  *string                `pulumi:"internalIpRange,optional"`
	CountryCode      *string                `pulumi:"countryCode,optional"`
	Comment          *string                `pulumi:"comment,optional"`
	IpUnnumbered     *bool                  `pulumi:"ipUnnumbered,optional"`
}

// TrafficForwardingGreTunnelState is the persisted state.
type TrafficForwardingGreTunnelState struct {
	TrafficForwardingGreTunnelArgs
	TunnelId         *int                    `pulumi:"tunnelId"`
	PrimaryDestVip   *GreTunnelDestVipOutput `pulumi:"primaryDestVip,optional"`
	SecondaryDestVip *GreTunnelDestVipOutput `pulumi:"secondaryDestVip,optional"`
}

func assignVipsIfNotSet(ctx context.Context, args *TrafficForwardingGreTunnelArgs, req *gretunnels.GreTunnels, service *zscaler.Service) error {
	primarySet := req.PrimaryDestVip != nil && (req.PrimaryDestVip.VirtualIP != "" || req.PrimaryDestVip.ID != 0)
	secondarySet := req.SecondaryDestVip != nil && (req.SecondaryDestVip.VirtualIP != "" || req.SecondaryDestVip.ID != 0)
	if primarySet && secondarySet {
		return nil
	}
	var pair []virtualipaddress.GREVirtualIPList
	countryCode := ptrToString(args.CountryCode)
	if countryCode != "" {
		vips, err := virtualipaddress.GetPairZSGREVirtualIPsWithinCountry(ctx, service, req.SourceIP, countryCode)
		if err != nil {
			log.Printf("[WARN] GetPairZSGREVirtualIPsWithinCountry failed: %v, falling back to GetZSGREVirtualIPList", err)
			vips, err = virtualipaddress.GetZSGREVirtualIPList(ctx, service, req.SourceIP, 2)
			if err != nil {
				return err
			}
		}
		pair = *vips
	} else {
		vips, err := virtualipaddress.GetZSGREVirtualIPList(ctx, service, req.SourceIP, 2)
		if err != nil {
			return err
		}
		pair = *vips
	}
	req.PrimaryDestVip = &gretunnels.PrimaryDestVip{ID: pair[0].ID, VirtualIP: pair[0].VirtualIp, Datacenter: pair[0].DataCenter}
	req.SecondaryDestVip = &gretunnels.SecondaryDestVip{ID: pair[1].ID, VirtualIP: pair[1].VirtualIp, Datacenter: pair[1].DataCenter}
	return nil
}

func trafficForwardingGreTunnelArgsToAPI(args *TrafficForwardingGreTunnelArgs, id int) gretunnels.GreTunnels {
	withinCountry := ptrToBool(args.WithinCountry)
	ipUnnumbered := ptrToBool(args.IpUnnumbered)
	api := gretunnels.GreTunnels{
		ID:              id,
		SourceIP:        args.SourceIP,
		InternalIpRange: ptrToString(args.InternalIpRange),
		WithinCountry:   &withinCountry,
		Comment:         ptrToString(args.Comment),
		IPUnnumbered:    ipUnnumbered,
	}
	if args.PrimaryDestVip != nil {
		api.PrimaryDestVip = &gretunnels.PrimaryDestVip{
			ID:         ptrToIntDefault(args.PrimaryDestVip.Id, 0),
			VirtualIP:  ptrToString(args.PrimaryDestVip.VirtualIP),
			Datacenter: ptrToString(args.PrimaryDestVip.Datacenter),
		}
	}
	if args.SecondaryDestVip != nil {
		api.SecondaryDestVip = &gretunnels.SecondaryDestVip{
			ID:         ptrToIntDefault(args.SecondaryDestVip.Id, 0),
			VirtualIP:  ptrToString(args.SecondaryDestVip.VirtualIP),
			Datacenter: ptrToString(args.SecondaryDestVip.Datacenter),
		}
	}
	return api
}

func trafficForwardingGreTunnelToState(api *gretunnels.GreTunnels) TrafficForwardingGreTunnelState {
	state := TrafficForwardingGreTunnelState{
		TrafficForwardingGreTunnelArgs: TrafficForwardingGreTunnelArgs{
			SourceIP:        api.SourceIP,
			InternalIpRange: stringPtr(api.InternalIpRange),
			Comment:         stringPtr(api.Comment),
			IpUnnumbered:    boolPtr(api.IPUnnumbered),
		},
		TunnelId: intPtr(api.ID),
	}
	if api.WithinCountry != nil {
		state.WithinCountry = boolPtr(*api.WithinCountry)
	}
	if api.PrimaryDestVip != nil {
		state.PrimaryDestVip = &GreTunnelDestVipOutput{
			Id:         intPtr(api.PrimaryDestVip.ID),
			VirtualIP:  stringPtr(api.PrimaryDestVip.VirtualIP),
			Datacenter: stringPtr(api.PrimaryDestVip.Datacenter),
		}
		state.TrafficForwardingGreTunnelArgs.PrimaryDestVip = &GreTunnelDestVipInput{
			Id:         intPtr(api.PrimaryDestVip.ID),
			VirtualIP:  stringPtr(api.PrimaryDestVip.VirtualIP),
			Datacenter: stringPtr(api.PrimaryDestVip.Datacenter),
		}
	}
	if api.SecondaryDestVip != nil {
		state.SecondaryDestVip = &GreTunnelDestVipOutput{
			Id:         intPtr(api.SecondaryDestVip.ID),
			VirtualIP:  stringPtr(api.SecondaryDestVip.VirtualIP),
			Datacenter: stringPtr(api.SecondaryDestVip.Datacenter),
		}
		state.TrafficForwardingGreTunnelArgs.SecondaryDestVip = &GreTunnelDestVipInput{
			Id:         intPtr(api.SecondaryDestVip.ID),
			VirtualIP:  stringPtr(api.SecondaryDestVip.VirtualIP),
			Datacenter: stringPtr(api.SecondaryDestVip.Datacenter),
		}
	}
	return state
}

func (TrafficForwardingGreTunnel) Create(ctx context.Context, req infer.CreateRequest[TrafficForwardingGreTunnelArgs]) (infer.CreateResponse[TrafficForwardingGreTunnelState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[TrafficForwardingGreTunnelState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	service := client.Service

	apiReq := trafficForwardingGreTunnelArgsToAPI(&req.Inputs, 0)
	if err := assignVipsIfNotSet(ctx, &req.Inputs, &apiReq, service); err != nil {
		return infer.CreateResponse[TrafficForwardingGreTunnelState]{}, fmt.Errorf("error assigning VIPs: %w", err)
	}

	resp, _, createErr := gretunnels.CreateGreTunnels(ctx, service, &apiReq)
	if createErr != nil {
		return infer.CreateResponse[TrafficForwardingGreTunnelState]{}, fmt.Errorf("error creating GRE tunnel: %w", createErr)
	}

	time.Sleep(2 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.CreateResponse[TrafficForwardingGreTunnelState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	updated, err := gretunnels.GetGreTunnels(ctx, service, resp.ID)
	if err != nil {
		state := TrafficForwardingGreTunnelState{
			TrafficForwardingGreTunnelArgs: req.Inputs,
			TunnelId:                       intPtr(resp.ID),
		}
		return infer.CreateResponse[TrafficForwardingGreTunnelState]{
			ID:     strconv.Itoa(resp.ID),
			Output: state,
		}, nil
	}
	state := trafficForwardingGreTunnelToState(updated)
	return infer.CreateResponse[TrafficForwardingGreTunnelState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (TrafficForwardingGreTunnel) Read(ctx context.Context, req infer.ReadRequest[TrafficForwardingGreTunnelArgs, TrafficForwardingGreTunnelState]) (infer.ReadResponse[TrafficForwardingGreTunnelArgs, TrafficForwardingGreTunnelState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[TrafficForwardingGreTunnelArgs, TrafficForwardingGreTunnelState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.ReadResponse[TrafficForwardingGreTunnelArgs, TrafficForwardingGreTunnelState]{}, fmt.Errorf("invalid GRE tunnel ID: %s", req.ID)
	}
	resp, err := gretunnels.GetGreTunnels(ctx, service, id)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[TrafficForwardingGreTunnelArgs, TrafficForwardingGreTunnelState]{}, nil
		}
		return infer.ReadResponse[TrafficForwardingGreTunnelArgs, TrafficForwardingGreTunnelState]{}, err
	}
	state := trafficForwardingGreTunnelToState(resp)
	return infer.ReadResponse[TrafficForwardingGreTunnelArgs, TrafficForwardingGreTunnelState]{
		ID:     req.ID,
		Inputs: state.TrafficForwardingGreTunnelArgs,
		State:  state,
	}, nil
}

func (TrafficForwardingGreTunnel) Update(ctx context.Context, req infer.UpdateRequest[TrafficForwardingGreTunnelArgs, TrafficForwardingGreTunnelState]) (infer.UpdateResponse[TrafficForwardingGreTunnelState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[TrafficForwardingGreTunnelState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	service := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[TrafficForwardingGreTunnelState]{}, fmt.Errorf("invalid GRE tunnel ID: %s", req.ID)
	}
	apiReq := trafficForwardingGreTunnelArgsToAPI(&req.Inputs, id)
	if err := assignVipsIfNotSet(ctx, &req.Inputs, &apiReq, service); err != nil {
		return infer.UpdateResponse[TrafficForwardingGreTunnelState]{}, fmt.Errorf("error assigning VIPs: %w", err)
	}
	if _, _, err := gretunnels.UpdateGreTunnels(ctx, service, id, &apiReq); err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.UpdateResponse[TrafficForwardingGreTunnelState]{}, nil
		}
		return infer.UpdateResponse[TrafficForwardingGreTunnelState]{}, err
	}
	time.Sleep(2 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[TrafficForwardingGreTunnelState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}
	updated, err := gretunnels.GetGreTunnels(ctx, service, id)
	if err != nil {
		return infer.UpdateResponse[TrafficForwardingGreTunnelState]{
			Output: TrafficForwardingGreTunnelState{
				TrafficForwardingGreTunnelArgs: req.Inputs,
				TunnelId:                       intPtr(id),
			},
		}, nil
	}
	return infer.UpdateResponse[TrafficForwardingGreTunnelState]{Output: trafficForwardingGreTunnelToState(updated)}, nil
}

func (TrafficForwardingGreTunnel) Delete(ctx context.Context, req infer.DeleteRequest[TrafficForwardingGreTunnelState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	service := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid GRE tunnel ID: %s", req.ID)
	}
	if _, err := gretunnels.DeleteGreTunnels(ctx, service, id); err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.DeleteResponse{}, nil
		}
		return infer.DeleteResponse{}, err
	}
	time.Sleep(2 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}
	return infer.DeleteResponse{}, nil
}

func (TrafficForwardingGreTunnel) Annotate(a infer.Annotator) {
	describeResource(a, &TrafficForwardingGreTunnel{}, `The zia_traffic_forwarding_gre_tunnel resource manages GRE (Generic Routing Encapsulation) tunnels for traffic forwarding in the Zscaler Internet Access (ZIA) cloud service. GRE tunnels are used to forward traffic from on-premises networks to the Zscaler cloud.

For more information, see the [ZIA Traffic Forwarding documentation](https://help.zscaler.com/zia/traffic-forwarding).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic GRE Tunnel

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.TrafficForwardingGreTunnel("example", {
    sourceIp: "203.0.113.10",
    comment: "Branch office GRE tunnel",
    withinCountry: true,
    ipUnnumbered: true,
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.TrafficForwardingGreTunnel("example",
    source_ip="203.0.113.10",
    comment="Branch office GRE tunnel",
    within_country=True,
    ip_unnumbered=True,
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:TrafficForwardingGreTunnel
    properties:
      sourceIp: "203.0.113.10"
      comment: Branch office GRE tunnel
      withinCountry: true
      ipUnnumbered: true
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing GRE tunnel can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:TrafficForwardingGreTunnel example 12345
`+tripleBacktick("")+`
`)
}

func (a *TrafficForwardingGreTunnelArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.SourceIP, "The source IP address of the GRE tunnel. This is typically a static IP associated with the location.")
	ann.Describe(&a.WithinCountry, "Restrict the data center virtual IP addresses (VIPs) only to those within the same country as the source IP.")
	ann.Describe(&a.PrimaryDestVip, "The primary destination data center and virtual IP address (VIP) of the GRE tunnel.")
	ann.Describe(&a.SecondaryDestVip, "The secondary destination data center and virtual IP address (VIP) of the GRE tunnel.")
	ann.Describe(&a.InternalIpRange, "The start of the internal IP address in /29 CIDR range. Automatically assigned if not provided.")
	ann.Describe(&a.CountryCode, "Country code (ISO 3166-1 alpha-2) used when withinCountry is true to restrict VIP selection.")
	ann.Describe(&a.Comment, "Additional information about the GRE tunnel.")
	ann.Describe(&a.IpUnnumbered, "When set to true, indicates that the GRE tunnel interface is unnumbered (no internal IP range is assigned).")
}

func (s *TrafficForwardingGreTunnelState) Annotate(a infer.Annotator) {
	a.Describe(&s.TunnelId, "The system-generated ID of the GRE tunnel.")
}

var _ infer.CustomResource[TrafficForwardingGreTunnelArgs, TrafficForwardingGreTunnelState] = TrafficForwardingGreTunnel{}
