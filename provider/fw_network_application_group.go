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

// Package provider implements the FW Network Application Group resource.
// Adopted from terraform-provider-zia resource_zia_fw_filtering_network_application_groups.go.

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
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/firewallpolicies/networkapplicationgroups"
)

// FwNetworkApplicationGroup implements the zia:index:FwNetworkApplicationGroup resource.
type FwNetworkApplicationGroup struct{}

// FwNetworkApplicationGroupArgs are the inputs.
type FwNetworkApplicationGroupArgs struct {
	Name                *string  `pulumi:"name,optional"`
	Description         *string  `pulumi:"description,optional"`
	NetworkApplications []string `pulumi:"networkApplications,optional"`
}

// FwNetworkApplicationGroupState is the persisted state.
type FwNetworkApplicationGroupState struct {
	FwNetworkApplicationGroupArgs
	AppId *int `pulumi:"appId"`
}

func fwNetworkApplicationGroupToAPI(args FwNetworkApplicationGroupArgs, id int) networkapplicationgroups.NetworkApplicationGroups {
	return networkapplicationgroups.NetworkApplicationGroups{
		ID:                  id,
		Name:                ptrToString(args.Name),
		Description:         ptrToString(args.Description),
		NetworkApplications: args.NetworkApplications,
	}
}

func (FwNetworkApplicationGroup) Create(ctx context.Context, req infer.CreateRequest[FwNetworkApplicationGroupArgs]) (infer.CreateResponse[FwNetworkApplicationGroupState], error) {
	if req.DryRun {
		s := FwNetworkApplicationGroupState{FwNetworkApplicationGroupArgs: req.Inputs, AppId: intPtr(0)}
		return infer.CreateResponse[FwNetworkApplicationGroupState]{ID: "preview", Output: s}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[FwNetworkApplicationGroupState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := fwNetworkApplicationGroupToAPI(req.Inputs, 0)
	resp, err := networkapplicationgroups.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[FwNetworkApplicationGroupState]{}, err
	}
	log.Printf("[INFO] Created ZIA network application group. ID: %v", resp.ID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[FwNetworkApplicationGroupState]{}, activationErr
		}
	}

	state := FwNetworkApplicationGroupState{
		FwNetworkApplicationGroupArgs: req.Inputs,
		AppId:                         &resp.ID,
	}
	return infer.CreateResponse[FwNetworkApplicationGroupState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (FwNetworkApplicationGroup) Read(ctx context.Context, req infer.ReadRequest[FwNetworkApplicationGroupArgs, FwNetworkApplicationGroupState]) (infer.ReadResponse[FwNetworkApplicationGroupArgs, FwNetworkApplicationGroupState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[FwNetworkApplicationGroupArgs, FwNetworkApplicationGroupState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.AppId != nil {
		id = *req.State.AppId
	}
	if id == 0 {
		return infer.ReadResponse[FwNetworkApplicationGroupArgs, FwNetworkApplicationGroupState]{}, fmt.Errorf("no network application group id in state")
	}

	resp, err := networkapplicationgroups.GetNetworkApplicationGroups(ctx, service, id)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[FwNetworkApplicationGroupArgs, FwNetworkApplicationGroupState]{ID: ""}, nil
		}
		return infer.ReadResponse[FwNetworkApplicationGroupArgs, FwNetworkApplicationGroupState]{}, err
	}

	args := FwNetworkApplicationGroupArgs{
		Name:                stringPtr(resp.Name),
		Description:         stringPtr(resp.Description),
		NetworkApplications: resp.NetworkApplications,
	}
	state := FwNetworkApplicationGroupState{
		FwNetworkApplicationGroupArgs: args,
		AppId:                         &resp.ID,
	}
	return infer.ReadResponse[FwNetworkApplicationGroupArgs, FwNetworkApplicationGroupState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (FwNetworkApplicationGroup) Update(ctx context.Context, req infer.UpdateRequest[FwNetworkApplicationGroupArgs, FwNetworkApplicationGroupState]) (infer.UpdateResponse[FwNetworkApplicationGroupState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[FwNetworkApplicationGroupState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.AppId != nil {
		id = *req.State.AppId
	}
	if id == 0 {
		return infer.UpdateResponse[FwNetworkApplicationGroupState]{}, fmt.Errorf("no network application group id in state")
	}

	if _, err := networkapplicationgroups.GetNetworkApplicationGroups(ctx, service, id); err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[FwNetworkApplicationGroupState]{}, nil
		}
		return infer.UpdateResponse[FwNetworkApplicationGroupState]{}, err
	}

	apiReq := fwNetworkApplicationGroupToAPI(req.Inputs, id)
	if _, _, err := networkapplicationgroups.Update(ctx, service, id, &apiReq); err != nil {
		return infer.UpdateResponse[FwNetworkApplicationGroupState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[FwNetworkApplicationGroupState]{}, activationErr
		}
	}

	state := FwNetworkApplicationGroupState{
		FwNetworkApplicationGroupArgs: req.Inputs,
		AppId:                         &id,
	}
	return infer.UpdateResponse[FwNetworkApplicationGroupState]{Output: state}, nil
}

func (FwNetworkApplicationGroup) Delete(ctx context.Context, req infer.DeleteRequest[FwNetworkApplicationGroupState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	service := client.Service

	id := 0
	if req.State.AppId != nil {
		id = *req.State.AppId
	}
	if id != 0 {
		if err := detachFromFilteringRules(ctx, client, id, "NwApplicationGroups",
			func(r *filteringrules.FirewallFilteringRules) []common.IDNameExtensions { return r.NwApplicationGroups },
			func(r *filteringrules.FirewallFilteringRules, ids []common.IDNameExtensions) {
				r.NwApplicationGroups = ids
			}); err != nil {
			return infer.DeleteResponse{}, err
		}
		if _, err := networkapplicationgroups.Delete(ctx, service, id); err != nil {
			if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.DeleteResponse{}, nil
			}
			return infer.DeleteResponse{}, err
		}
		log.Printf("[INFO] ZIA network application group deleted")
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

func (FwNetworkApplicationGroup) Annotate(a infer.Annotator) {
	describeResource(a, &FwNetworkApplicationGroup{}, `The zia_fw_network_application_group resource manages firewall network application groups in the Zscaler Internet Access (ZIA) cloud service. Network application groups allow you to bundle multiple network applications together for use in firewall filtering rules.

For more information, see the [ZIA Firewall Policies documentation](https://help.zscaler.com/zia/firewall-policies).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Network Application Group

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.FwNetworkApplicationGroup("example", {
    name: "Example App Group",
    description: "Group of network applications",
    networkApplications: ["APNS", "APPSTORE"],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.FwNetworkApplicationGroup("example",
    name="Example App Group",
    description="Group of network applications",
    network_applications=["APNS", "APPSTORE"],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:FwNetworkApplicationGroup
    properties:
      name: Example App Group
      description: Group of network applications
      networkApplications:
        - APNS
        - APPSTORE
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing network application group can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:FwNetworkApplicationGroup example 12345
`+tripleBacktick("")+`
`)
}

func (a *FwNetworkApplicationGroupArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the network application group.")
	ann.Describe(&a.Description, "Additional information about the network application group.")
	ann.Describe(&a.NetworkApplications, "List of network application identifiers that belong to this group (e.g., `APNS`, `APPSTORE`).")
}

func (s *FwNetworkApplicationGroupState) Annotate(a infer.Annotator) {
	a.Describe(&s.AppId, "The system-generated ID of the network application group.")
}

var _ infer.CustomResource[FwNetworkApplicationGroupArgs, FwNetworkApplicationGroupState] = FwNetworkApplicationGroup{}
