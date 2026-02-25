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

// Package provider implements the Forwarding Control ZPA Gateway resource.
// Adopted from terraform-provider-zia resource_zia_forwarding_control_zpa_gateway.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/forwarding_control_policy/zpa_gateways"
)

// ForwardingControlZpaGateway implements the zia:index:ForwardingControlZpaGateway resource.
type ForwardingControlZpaGateway struct{}

// ZpaServerGroupInput is the ZPA Server Group block.
type ZpaServerGroupInput struct {
	ExternalId *string `pulumi:"externalId,optional"`
	Name       *string `pulumi:"name,optional"`
}

// ZpaAppSegmentInput is an Application Segment for the ZPA Gateway.
type ZpaAppSegmentInput struct {
	Name       *string `pulumi:"name,optional"`
	ExternalId *string `pulumi:"externalId,optional"`
}

// ForwardingControlZpaGatewayArgs are the inputs.
type ForwardingControlZpaGatewayArgs struct {
	Name            *string             `pulumi:"name,optional"`
	Description     *string             `pulumi:"description,optional"`
	Type            *string             `pulumi:"type,optional"` // ZPA, ECZPA
	ZpaServerGroup  *ZpaServerGroupInput `pulumi:"zpaServerGroup,optional"`
	ZpaAppSegments  []ZpaAppSegmentInput `pulumi:"zpaAppSegments,optional"`
}

// ForwardingControlZpaGatewayState is the persisted state.
type ForwardingControlZpaGatewayState struct {
	ForwardingControlZpaGatewayArgs
	GatewayId *int `pulumi:"gatewayId"`
}

func validateZpaGatewayPredefined(gw zpa_gateways.ZPAGateways) error {
	if gw.Name == "Auto ZPA Gateway" {
		return fmt.Errorf("predefined zpa gateway '%s' cannot be deleted", gw.Name)
	}
	return nil
}

func forwardingControlZpaGatewayToAPI(args ForwardingControlZpaGatewayArgs, id int) zpa_gateways.ZPAGateways {
	gatewayType := ptrToString(args.Type)
	if gatewayType == "" {
		gatewayType = "ZPA"
	}
	out := zpa_gateways.ZPAGateways{
		ID:          id,
		Name:        ptrToString(args.Name),
		Description: ptrToString(args.Description),
		Type:        gatewayType,
	}
	if args.ZpaServerGroup != nil {
		out.ZPAServerGroup = zpa_gateways.ZPAServerGroup{
			ExternalID: ptrToString(args.ZpaServerGroup.ExternalId),
			Name:       ptrToString(args.ZpaServerGroup.Name),
		}
	}
	if len(args.ZpaAppSegments) > 0 {
		out.ZPAAppSegments = make([]zpa_gateways.ZPAAppSegments, 0, len(args.ZpaAppSegments))
		for _, s := range args.ZpaAppSegments {
			out.ZPAAppSegments = append(out.ZPAAppSegments, zpa_gateways.ZPAAppSegments{
				Name:       ptrToString(s.Name),
				ExternalID: ptrToString(s.ExternalId),
			})
		}
	}
	return out
}

func (ForwardingControlZpaGateway) Create(ctx context.Context, req infer.CreateRequest[ForwardingControlZpaGatewayArgs]) (infer.CreateResponse[ForwardingControlZpaGatewayState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[ForwardingControlZpaGatewayState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := forwardingControlZpaGatewayToAPI(req.Inputs, 0)
	if err := validateZpaGatewayPredefined(apiReq); err != nil {
		return infer.CreateResponse[ForwardingControlZpaGatewayState]{}, err
	}

	resp, err := zpa_gateways.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[ForwardingControlZpaGatewayState]{}, err
	}
	log.Printf("[INFO] Created ZIA Forwarding Control ZPA Gateway. ID: %v", resp.ID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[ForwardingControlZpaGatewayState]{}, activationErr
		}
	}

	state := ForwardingControlZpaGatewayState{
		ForwardingControlZpaGatewayArgs: req.Inputs,
		GatewayId:                       &resp.ID,
	}
	return infer.CreateResponse[ForwardingControlZpaGatewayState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (ForwardingControlZpaGateway) Read(ctx context.Context, req infer.ReadRequest[ForwardingControlZpaGatewayArgs, ForwardingControlZpaGatewayState]) (infer.ReadResponse[ForwardingControlZpaGatewayArgs, ForwardingControlZpaGatewayState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[ForwardingControlZpaGatewayArgs, ForwardingControlZpaGatewayState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.GatewayId != nil {
		id = *req.State.GatewayId
	}
	if id == 0 {
		return infer.ReadResponse[ForwardingControlZpaGatewayArgs, ForwardingControlZpaGatewayState]{}, fmt.Errorf("no ZPA gateway id in state")
	}

	// Terraform uses GetAll to avoid API bug where Get by ID returns incorrect app segments
	all, err := zpa_gateways.GetAll(ctx, service)
	if err != nil {
		return infer.ReadResponse[ForwardingControlZpaGatewayArgs, ForwardingControlZpaGatewayState]{}, err
	}
	var resp *zpa_gateways.ZPAGateways
	for i := range all {
		if all[i].ID == id {
			resp = &all[i]
			break
		}
	}
	if resp == nil {
		return infer.ReadResponse[ForwardingControlZpaGatewayArgs, ForwardingControlZpaGatewayState]{ID: ""}, nil
	}

	gatewayType := resp.Type
	if gatewayType == "" {
		gatewayType = "ZPA"
	}

	var zpaServerGroup *ZpaServerGroupInput
	if resp.ZPAServerGroup.Name != "" || resp.ZPAServerGroup.ExternalID != "" {
		zpaServerGroup = &ZpaServerGroupInput{
			ExternalId: stringPtr(resp.ZPAServerGroup.ExternalID),
			Name:       stringPtr(resp.ZPAServerGroup.Name),
		}
	}

	zpaAppSegments := make([]ZpaAppSegmentInput, 0, len(resp.ZPAAppSegments))
	for _, s := range resp.ZPAAppSegments {
		zpaAppSegments = append(zpaAppSegments, ZpaAppSegmentInput{
			Name:       stringPtr(s.Name),
			ExternalId: stringPtr(s.ExternalID),
		})
	}

	args := ForwardingControlZpaGatewayArgs{
		Name:           stringPtr(resp.Name),
		Description:    stringPtr(resp.Description),
		Type:           stringPtr(gatewayType),
		ZpaServerGroup: zpaServerGroup,
		ZpaAppSegments:  zpaAppSegments,
	}
	state := ForwardingControlZpaGatewayState{
		ForwardingControlZpaGatewayArgs: args,
		GatewayId:                       &resp.ID,
	}
	return infer.ReadResponse[ForwardingControlZpaGatewayArgs, ForwardingControlZpaGatewayState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (ForwardingControlZpaGateway) Update(ctx context.Context, req infer.UpdateRequest[ForwardingControlZpaGatewayArgs, ForwardingControlZpaGatewayState]) (infer.UpdateResponse[ForwardingControlZpaGatewayState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[ForwardingControlZpaGatewayState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.GatewayId != nil {
		id = *req.State.GatewayId
	}
	if id == 0 {
		return infer.UpdateResponse[ForwardingControlZpaGatewayState]{}, fmt.Errorf("no ZPA gateway id in state")
	}

	apiReq := forwardingControlZpaGatewayToAPI(req.Inputs, id)
	if err := validateZpaGatewayPredefined(apiReq); err != nil {
		return infer.UpdateResponse[ForwardingControlZpaGatewayState]{}, err
	}

	all, err := zpa_gateways.GetAll(ctx, service)
	if err != nil {
		return infer.UpdateResponse[ForwardingControlZpaGatewayState]{}, err
	}
	var exists bool
	for i := range all {
		if all[i].ID == id {
			exists = true
			break
		}
	}
	if !exists {
		return infer.UpdateResponse[ForwardingControlZpaGatewayState]{}, nil
	}

	if _, err := zpa_gateways.Update(ctx, service, id, &apiReq); err != nil {
		return infer.UpdateResponse[ForwardingControlZpaGatewayState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[ForwardingControlZpaGatewayState]{}, activationErr
		}
	}

	state := ForwardingControlZpaGatewayState{
		ForwardingControlZpaGatewayArgs: req.Inputs,
		GatewayId:                       &id,
	}
	return infer.UpdateResponse[ForwardingControlZpaGatewayState]{Output: state}, nil
}

func (ForwardingControlZpaGateway) Delete(ctx context.Context, req infer.DeleteRequest[ForwardingControlZpaGatewayState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.GatewayId != nil {
		id = *req.State.GatewayId
	}
	if id == 0 {
		return infer.DeleteResponse{}, nil
	}

	gw, err := zpa_gateways.Get(ctx, service, id)
	if err != nil {
		return infer.DeleteResponse{}, err
	}
	if err := validateZpaGatewayPredefined(*gw); err != nil {
		return infer.DeleteResponse{}, err
	}

	if _, err := zpa_gateways.Delete(ctx, service, id); err != nil {
		return infer.DeleteResponse{}, err
	}
	log.Printf("[INFO] ZIA Forwarding Control ZPA Gateway deleted")

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

func (ForwardingControlZpaGateway) Annotate(a infer.Annotator) {
	describeResource(a, &ForwardingControlZpaGateway{}, `The zia.ForwardingControlZpaGateway resource manages forwarding control ZPA gateway configurations in the
Zscaler Internet Access (ZIA) cloud. ZPA gateways are used in forwarding control rules to direct traffic
to Zscaler Private Access (ZPA) application segments through ZPA server groups.

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Forwarding Control ZPA Gateway

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.ForwardingControlZpaGateway("example", {
    name: "Example ZPA Gateway",
    description: "Managed by Pulumi",
    type: "ZPA",
    zpaServerGroup: {
        externalId: "server-group-external-id",
        name: "Example Server Group",
    },
    zpaAppSegments: [{
        externalId: "app-segment-external-id",
        name: "Example App Segment",
    }],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.ForwardingControlZpaGateway("example",
    name="Example ZPA Gateway",
    description="Managed by Pulumi",
    type="ZPA",
    zpa_server_group={
        "external_id": "server-group-external-id",
        "name": "Example Server Group",
    },
    zpa_app_segments=[{
        "external_id": "app-segment-external-id",
        "name": "Example App Segment",
    }],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:ForwardingControlZpaGateway
    properties:
      name: Example ZPA Gateway
      description: Managed by Pulumi
      type: ZPA
      zpaServerGroup:
        externalId: server-group-external-id
        name: Example Server Group
      zpaAppSegments:
        - externalId: app-segment-external-id
          name: Example App Segment
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing forwarding control ZPA gateway can be imported using its ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:ForwardingControlZpaGateway example 12345
`+tripleBacktick("")+`
`)
}

func (a *ForwardingControlZpaGatewayArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "Name of the ZPA gateway.")
	ann.Describe(&a.Description, "Description of the ZPA gateway.")
	ann.Describe(&a.Type, "The gateway type. Accepted values: 'ZPA' or 'ECZPA'.")
	ann.Describe(&a.ZpaServerGroup, "The ZPA server group associated with the gateway.")
	ann.Describe(&a.ZpaAppSegments, "List of ZPA application segments associated with the gateway.")
}

func (s *ForwardingControlZpaGatewayState) Annotate(a infer.Annotator) {
	a.Describe(&s.GatewayId, "The unique identifier for the ZPA gateway assigned by the ZIA cloud.")
}

var _ infer.CustomResource[ForwardingControlZpaGatewayArgs, ForwardingControlZpaGatewayState] = ForwardingControlZpaGateway{}
