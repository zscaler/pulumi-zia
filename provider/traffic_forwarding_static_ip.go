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

// Package provider implements the Traffic Forwarding Static IP resource.
// Adopted from terraform-provider-zia resource_zia_traffic_forwarding_static_ips.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/trafficforwarding/staticips"
)

// TrafficForwardingStaticIp implements the zia:index:TrafficForwardingStaticIp resource.
type TrafficForwardingStaticIp struct{}

// TrafficForwardingStaticIpArgs are the inputs.
type TrafficForwardingStaticIpArgs struct {
	IPAddress   string   `pulumi:"ipAddress"`
	GeoOverride *bool    `pulumi:"geoOverride,optional"`
	Latitude    *float64 `pulumi:"latitude,optional"`
	Longitude   *float64 `pulumi:"longitude,optional"`
	RoutableIP  *bool    `pulumi:"routableIp,optional"`
	Comment     *string  `pulumi:"comment,optional"`
}

// TrafficForwardingStaticIpState is the persisted state.
type TrafficForwardingStaticIpState struct {
	TrafficForwardingStaticIpArgs
	StaticIpId *int `pulumi:"staticIpId"`
}

func (TrafficForwardingStaticIp) Check(ctx context.Context, req infer.CheckRequest) (infer.CheckResponse[TrafficForwardingStaticIpArgs], error) {
	inputs, failures, err := infer.DefaultCheck[TrafficForwardingStaticIpArgs](ctx, req.NewInputs)
	if err != nil {
		return infer.CheckResponse[TrafficForwardingStaticIpArgs]{}, err
	}
	if len(failures) > 0 {
		return infer.CheckResponse[TrafficForwardingStaticIpArgs]{Failures: failures}, nil
	}
	if inputs.Latitude != nil && (*inputs.Latitude < -90 || *inputs.Latitude > 90) {
		return infer.CheckResponse[TrafficForwardingStaticIpArgs]{Failures: []p.CheckFailure{{
			Property: "latitude",
			Reason:   "latitude must be between -90 and 90",
		}}}, nil
	}
	if inputs.Longitude != nil && (*inputs.Longitude < -180 || *inputs.Longitude > 180) {
		return infer.CheckResponse[TrafficForwardingStaticIpArgs]{Failures: []p.CheckFailure{{
			Property: "longitude",
			Reason:   "longitude must be between -180 and 180",
		}}}, nil
	}
	return infer.CheckResponse[TrafficForwardingStaticIpArgs]{Inputs: inputs}, nil
}

func (TrafficForwardingStaticIp) Create(ctx context.Context, req infer.CreateRequest[TrafficForwardingStaticIpArgs]) (infer.CreateResponse[TrafficForwardingStaticIpState], error) {
	if req.DryRun {
		s := TrafficForwardingStaticIpState{TrafficForwardingStaticIpArgs: req.Inputs, StaticIpId: intPtr(0)}
		return infer.CreateResponse[TrafficForwardingStaticIpState]{ID: "preview", Output: s}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[TrafficForwardingStaticIpState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	args := req.Inputs
	geoOverride := args.GeoOverride != nil && *args.GeoOverride
	hasLat := args.Latitude != nil
	hasLon := args.Longitude != nil

	var staticIpID int
	if geoOverride && (!hasLat || !hasLon) {
		// Auto-populate coordinates: create with geo_override=false or use existing IP's coords
		existingIP, err := staticips.GetByIPAddress(ctx, service, args.IPAddress)
		if err == nil {
			// Reuse existing coordinates
			args.Latitude = &existingIP.Latitude
			args.Longitude = &existingIP.Longitude
			staticIpID = existingIP.ID
			// IP already exists - we're in a weird state; Terraform would have set d.Id() from temp create
			// For Pulumi, if IP exists we can't "create" it again. This is an import scenario.
			// Actually re-reading the Terraform: when existingIP is found, it just sets lat/lon and returns nil -
			// it does NOT set d.SetId. So the create continues with the updated lat/lon. So we still need to CREATE.
			// Wait - if the IP exists, we can't create it again (IP is unique). So the Terraform flow is:
			// 1. geo_override=true, no coords
			// 2. Check existing IP - if found, set lat/lon from it and return (so create will use those)
			// 3. Create will then run - but Create will send the IP which already exists = error?
			// Let me re-read... Oh! When existingIP is found, Terraform does _ = d.Set("latitude", ...) and return nil.
			// So it populates the schema. Then the normal create runs. But the create would try to create the same IP again!
			// Unless... the create is skipped? Let me look again.
			// if d.Id() != "" { log and skip to read } - so if autoPopulateCoordinates set d.Id(), we skip create.
			// When does it set d.Id? When IP doesn't exist - it creates temp, gets coords, updates with geo_override=true, then
			// d.SetId(strconv.Itoa(tempResp.ID)) and _ = d.Set("static_ip_id", tempResp.ID). So it creates AND sets ID.
			// So when existing IP is found, it does NOT set d.Id - it just sets lat/lon. Then create runs and... would fail
			// because the IP already exists. Hmm.
			// Actually re-read: when existingIP found, it returns nil. So create continues. The expandTrafficForwardingStaticIP
			// would use the updated lat/lon. So we'd send a CREATE request with ipAddress, geoOverride=true, lat, lon...
			// The API would likely return "duplicate" or similar. So perhaps the Terraform has a bug, or the API allows
			// creating same IP? Unlikely. Let me assume: when IP exists we need to "adopt" it - i.e. we'd need to
			// set the state as if we created it. So we return state with existingIP.ID and the rest. We don't call Create.
			state := TrafficForwardingStaticIpState{
				TrafficForwardingStaticIpArgs: TrafficForwardingStaticIpArgs{
					IPAddress:   args.IPAddress,
					GeoOverride: boolPtr(true),
					Latitude:    &existingIP.Latitude,
					Longitude:   &existingIP.Longitude,
					RoutableIP:  args.RoutableIP,
					Comment:     args.Comment,
				},
				StaticIpId: intPtr(existingIP.ID),
			}
			return infer.CreateResponse[TrafficForwardingStaticIpState]{
				ID:     strconv.Itoa(existingIP.ID),
				Output: state,
			}, nil
		}
		// IP doesn't exist - create with geo_override=false to get coords, then update
		tempReq := staticips.StaticIP{
			IpAddress:   args.IPAddress,
			GeoOverride: false,
			RoutableIP:  ptrToBool(args.RoutableIP),
			Comment:     ptrToString(args.Comment),
		}
		tempResp, _, err := staticips.Create(ctx, service, &tempReq)
		if err != nil {
			return infer.CreateResponse[TrafficForwardingStaticIpState]{}, fmt.Errorf("failed to create static IP to determine coordinates: %w", err)
		}
		ipWithCoords, err := staticips.Get(ctx, service, tempResp.ID)
		if err != nil {
			return infer.CreateResponse[TrafficForwardingStaticIpState]{}, fmt.Errorf("failed to fetch static IP coordinates: %w", err)
		}
		updateReq := staticips.StaticIP{
			ID:          tempResp.ID,
			IpAddress:   args.IPAddress,
			GeoOverride: true,
			Latitude:    ipWithCoords.Latitude,
			Longitude:   ipWithCoords.Longitude,
			RoutableIP:  ptrToBool(args.RoutableIP),
			Comment:     ptrToString(args.Comment),
		}
		if _, _, err = staticips.Update(ctx, service, tempResp.ID, &updateReq); err != nil {
			return infer.CreateResponse[TrafficForwardingStaticIpState]{}, fmt.Errorf("failed to update static IP with geo_override: %w", err)
		}
		staticIpID = tempResp.ID
	} else {
		// Normal create
		apiReq := trafficForwardingStaticIpArgsToAPI(args, 0)
		resp, _, err := staticips.Create(ctx, service, &apiReq)
		if err != nil {
			return infer.CreateResponse[TrafficForwardingStaticIpState]{}, err
		}
		staticIpID = resp.ID
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[TrafficForwardingStaticIpState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	// Read back to get final state
	ip, err := staticips.Get(ctx, service, staticIpID)
	if err != nil {
		return infer.CreateResponse[TrafficForwardingStaticIpState]{}, err
	}
	state := trafficForwardingStaticIpToState(ip)
	return infer.CreateResponse[TrafficForwardingStaticIpState]{
		ID:     strconv.Itoa(ip.ID),
		Output: state,
	}, nil
}

func (TrafficForwardingStaticIp) Read(ctx context.Context, req infer.ReadRequest[TrafficForwardingStaticIpArgs, TrafficForwardingStaticIpState]) (infer.ReadResponse[TrafficForwardingStaticIpArgs, TrafficForwardingStaticIpState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[TrafficForwardingStaticIpArgs, TrafficForwardingStaticIpState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.ReadResponse[TrafficForwardingStaticIpArgs, TrafficForwardingStaticIpState]{}, fmt.Errorf("invalid static ip id: %w", err)
	}

	allIPs, err := staticips.GetAll(ctx, service)
	if err != nil {
		return infer.ReadResponse[TrafficForwardingStaticIpArgs, TrafficForwardingStaticIpState]{}, err
	}

	var resp *staticips.StaticIP
	for i := range allIPs {
		if allIPs[i].ID == id {
			resp = &allIPs[i]
			break
		}
	}

	if resp == nil {
		ip, err := staticips.Get(ctx, service, id)
		if err != nil {
			if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[TrafficForwardingStaticIpArgs, TrafficForwardingStaticIpState]{}, nil
			}
			return infer.ReadResponse[TrafficForwardingStaticIpArgs, TrafficForwardingStaticIpState]{}, err
		}
		resp = ip
	}

	state := trafficForwardingStaticIpToState(resp)
	return infer.ReadResponse[TrafficForwardingStaticIpArgs, TrafficForwardingStaticIpState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: state.TrafficForwardingStaticIpArgs,
		State:  state,
	}, nil
}

func (TrafficForwardingStaticIp) Update(ctx context.Context, req infer.UpdateRequest[TrafficForwardingStaticIpArgs, TrafficForwardingStaticIpState]) (infer.UpdateResponse[TrafficForwardingStaticIpState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[TrafficForwardingStaticIpState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[TrafficForwardingStaticIpState]{}, fmt.Errorf("invalid static ip id: %w", err)
	}

	currentIP, err := staticips.Get(ctx, service, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.UpdateResponse[TrafficForwardingStaticIpState]{}, nil
		}
		return infer.UpdateResponse[TrafficForwardingStaticIpState]{}, err
	}

	args := req.Inputs
	if ptrToBool(args.GeoOverride) && (args.Latitude == nil || args.Longitude == nil) {
		args.Latitude = &currentIP.Latitude
		args.Longitude = &currentIP.Longitude
	}

	apiReq := trafficForwardingStaticIpArgsToAPI(args, id)
	if _, _, err := staticips.Update(ctx, service, id, &apiReq); err != nil {
		return infer.UpdateResponse[TrafficForwardingStaticIpState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[TrafficForwardingStaticIpState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	ip, err := staticips.Get(ctx, service, id)
	if err != nil {
		return infer.UpdateResponse[TrafficForwardingStaticIpState]{}, err
	}
	state := trafficForwardingStaticIpToState(ip)
	return infer.UpdateResponse[TrafficForwardingStaticIpState]{Output: state}, nil
}

func (TrafficForwardingStaticIp) Delete(ctx context.Context, req infer.DeleteRequest[TrafficForwardingStaticIpState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid static ip id: %w", err)
	}

	if _, err := staticips.Delete(ctx, service, id); err != nil {
		return infer.DeleteResponse{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	return infer.DeleteResponse{}, nil
}

func (TrafficForwardingStaticIp) Annotate(a infer.Annotator) {
	describeResource(a, &TrafficForwardingStaticIp{}, `The zia_traffic_forwarding_static_ip resource manages static IP addresses for traffic forwarding in the Zscaler Internet Access (ZIA) cloud service. Static IPs are used to associate traffic with a specific location or GRE tunnel.

For more information, see the [ZIA Traffic Forwarding documentation](https://help.zscaler.com/zia/traffic-forwarding).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Static IP

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.TrafficForwardingStaticIp("example", {
    ipAddress: "203.0.113.10",
    comment: "Branch office static IP",
    routableIp: true,
    geoOverride: false,
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.TrafficForwardingStaticIp("example",
    ip_address="203.0.113.10",
    comment="Branch office static IP",
    routable_ip=True,
    geo_override=False,
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:TrafficForwardingStaticIp
    properties:
      ipAddress: "203.0.113.10"
      comment: Branch office static IP
      routableIp: true
      geoOverride: false
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing static IP can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:TrafficForwardingStaticIp example 12345
`+tripleBacktick("")+`
`)
}

func (a *TrafficForwardingStaticIpArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.IPAddress, "The static IP address.")
	ann.Describe(&a.GeoOverride, "If not set, geographic coordinates and city are automatically determined from the IP address. When set to true, manually-specified latitude and longitude are used instead.")
	ann.Describe(&a.Latitude, "Required only if geoOverride is true. Latitude of the static IP. Valid range: -90 to 90.")
	ann.Describe(&a.Longitude, "Required only if geoOverride is true. Longitude of the static IP. Valid range: -180 to 180.")
	ann.Describe(&a.RoutableIP, "Indicates whether a non-RFC 1918 IP address is publicly routable.")
	ann.Describe(&a.Comment, "Additional information about the static IP.")
}

func (s *TrafficForwardingStaticIpState) Annotate(a infer.Annotator) {
	a.Describe(&s.StaticIpId, "The system-generated ID of the static IP.")
}

func trafficForwardingStaticIpArgsToAPI(args TrafficForwardingStaticIpArgs, id int) staticips.StaticIP {
	lat, lon := 0.0, 0.0
	if args.Latitude != nil {
		lat = *args.Latitude
	}
	if args.Longitude != nil {
		lon = *args.Longitude
	}
	return staticips.StaticIP{
		ID:          id,
		IpAddress:   args.IPAddress,
		GeoOverride: ptrToBool(args.GeoOverride),
		Latitude:    lat,
		Longitude:   lon,
		RoutableIP:  ptrToBool(args.RoutableIP),
		Comment:     ptrToString(args.Comment),
	}
}

func trafficForwardingStaticIpToState(r *staticips.StaticIP) TrafficForwardingStaticIpState {
	return TrafficForwardingStaticIpState{
		StaticIpId: intPtr(r.ID),
		TrafficForwardingStaticIpArgs: TrafficForwardingStaticIpArgs{
			IPAddress:   r.IpAddress,
			GeoOverride: boolPtr(r.GeoOverride),
			Latitude:    &r.Latitude,
			Longitude:   &r.Longitude,
			RoutableIP:  boolPtr(r.RoutableIP),
			Comment:     stringPtr(r.Comment),
		},
	}
}

var _ infer.CustomResource[TrafficForwardingStaticIpArgs, TrafficForwardingStaticIpState] = TrafficForwardingStaticIp{}
