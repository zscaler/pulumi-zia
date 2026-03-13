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

// Package provider implements the FW Network Service Group resource.
// Adopted from terraform-provider-zia resource_zia_fw_filtering_network_services_groups.go.

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
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/firewallpolicies/networkservicegroups"
)

// FwNetworkServiceGroup implements the zia:index:FwNetworkServiceGroup resource.
type FwNetworkServiceGroup struct{}

// FwNetworkServiceGroupArgs are the inputs.
type FwNetworkServiceGroupArgs struct {
	Name        *string `pulumi:"name,optional"`
	Description *string `pulumi:"description,optional"`
	ServiceIds  []int   `pulumi:"serviceIds,optional"`
}

// FwNetworkServiceGroupState is the persisted state.
type FwNetworkServiceGroupState struct {
	FwNetworkServiceGroupArgs
	GroupId *int `pulumi:"groupId"`
}

func serviceIdsToAPI(ids []int) []networkservicegroups.Services {
	if len(ids) == 0 {
		return nil
	}
	result := make([]networkservicegroups.Services, len(ids))
	for i, id := range ids {
		result[i] = networkservicegroups.Services{ID: id}
	}
	return result
}

func serviceIdsFromAPI(list []networkservicegroups.Services) []int {
	if len(list) == 0 {
		return nil
	}
	result := make([]int, len(list))
	for i, s := range list {
		result[i] = s.ID
	}
	return result
}

func fwNetworkServiceGroupToAPI(args FwNetworkServiceGroupArgs, id int) networkservicegroups.NetworkServiceGroups {
	return networkservicegroups.NetworkServiceGroups{
		ID:          id,
		Name:        ptrToString(args.Name),
		Description: ptrToString(args.Description),
		Services:    serviceIdsToAPI(args.ServiceIds),
	}
}

func (FwNetworkServiceGroup) Create(ctx context.Context, req infer.CreateRequest[FwNetworkServiceGroupArgs]) (infer.CreateResponse[FwNetworkServiceGroupState], error) {
	if req.DryRun {
		s := FwNetworkServiceGroupState{FwNetworkServiceGroupArgs: req.Inputs, GroupId: intPtr(0)}
		return infer.CreateResponse[FwNetworkServiceGroupState]{ID: "preview", Output: s}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[FwNetworkServiceGroupState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := fwNetworkServiceGroupToAPI(req.Inputs, 0)
	resp, err := networkservicegroups.CreateNetworkServiceGroups(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[FwNetworkServiceGroupState]{}, err
	}
	log.Printf("[INFO] Created ZIA network service group. ID: %v", resp.ID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[FwNetworkServiceGroupState]{}, activationErr
		}
	}

	state := FwNetworkServiceGroupState{
		FwNetworkServiceGroupArgs: req.Inputs,
		GroupId:                   &resp.ID,
	}
	return infer.CreateResponse[FwNetworkServiceGroupState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (FwNetworkServiceGroup) Read(ctx context.Context, req infer.ReadRequest[FwNetworkServiceGroupArgs, FwNetworkServiceGroupState]) (infer.ReadResponse[FwNetworkServiceGroupArgs, FwNetworkServiceGroupState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[FwNetworkServiceGroupArgs, FwNetworkServiceGroupState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.GroupId != nil {
		id = *req.State.GroupId
	}
	if id == 0 {
		return infer.ReadResponse[FwNetworkServiceGroupArgs, FwNetworkServiceGroupState]{}, fmt.Errorf("no network service group id in state")
	}

	resp, err := networkservicegroups.GetNetworkServiceGroups(ctx, service, id)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[FwNetworkServiceGroupArgs, FwNetworkServiceGroupState]{ID: ""}, nil
		}
		return infer.ReadResponse[FwNetworkServiceGroupArgs, FwNetworkServiceGroupState]{}, err
	}

	args := FwNetworkServiceGroupArgs{
		Name:        stringPtr(resp.Name),
		Description: stringPtr(resp.Description),
		ServiceIds:  serviceIdsFromAPI(resp.Services),
	}
	state := FwNetworkServiceGroupState{
		FwNetworkServiceGroupArgs: args,
		GroupId:                   &resp.ID,
	}
	return infer.ReadResponse[FwNetworkServiceGroupArgs, FwNetworkServiceGroupState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (FwNetworkServiceGroup) Update(ctx context.Context, req infer.UpdateRequest[FwNetworkServiceGroupArgs, FwNetworkServiceGroupState]) (infer.UpdateResponse[FwNetworkServiceGroupState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[FwNetworkServiceGroupState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.GroupId != nil {
		id = *req.State.GroupId
	}
	if id == 0 {
		return infer.UpdateResponse[FwNetworkServiceGroupState]{}, fmt.Errorf("no network service group id in state")
	}

	if _, err := networkservicegroups.GetNetworkServiceGroups(ctx, service, id); err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[FwNetworkServiceGroupState]{}, nil
		}
		return infer.UpdateResponse[FwNetworkServiceGroupState]{}, err
	}

	apiReq := fwNetworkServiceGroupToAPI(req.Inputs, id)
	if _, _, err := networkservicegroups.UpdateNetworkServiceGroups(ctx, service, id, &apiReq); err != nil {
		return infer.UpdateResponse[FwNetworkServiceGroupState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[FwNetworkServiceGroupState]{}, activationErr
		}
	}

	state := FwNetworkServiceGroupState{
		FwNetworkServiceGroupArgs: req.Inputs,
		GroupId:                   &id,
	}
	return infer.UpdateResponse[FwNetworkServiceGroupState]{Output: state}, nil
}

func (FwNetworkServiceGroup) Delete(ctx context.Context, req infer.DeleteRequest[FwNetworkServiceGroupState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	service := client.Service

	id := 0
	if req.State.GroupId != nil {
		id = *req.State.GroupId
	}
	if id != 0 {
		if err := detachFromFilteringRules(ctx, client, id, "NwServiceGroups",
			func(r *filteringrules.FirewallFilteringRules) []common.IDNameExtensions { return r.NwServiceGroups },
			func(r *filteringrules.FirewallFilteringRules, ids []common.IDNameExtensions) { r.NwServiceGroups = ids }); err != nil {
			return infer.DeleteResponse{}, err
		}
		if _, err := networkservicegroups.DeleteNetworkServiceGroups(ctx, service, id); err != nil {
			if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.DeleteResponse{}, nil
			}
			return infer.DeleteResponse{}, err
		}
		log.Printf("[INFO] ZIA network service group deleted")
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

func (FwNetworkServiceGroup) Annotate(a infer.Annotator) {
	describeResource(a, &FwNetworkServiceGroup{}, `The zia_fw_network_service_group resource manages firewall network service groups in the Zscaler Internet Access (ZIA) cloud service. Network service groups allow you to bundle multiple network services together for use in firewall filtering rules.

For more information, see the [ZIA Firewall Policies documentation](https://help.zscaler.com/zia/firewall-policies).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Network Service Group

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.FwNetworkServiceGroup("example", {
    name: "Example Service Group",
    description: "Group of network services",
    serviceIds: [12345, 67890],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.FwNetworkServiceGroup("example",
    name="Example Service Group",
    description="Group of network services",
    service_ids=[12345, 67890],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:FwNetworkServiceGroup
    properties:
      name: Example Service Group
      description: Group of network services
      serviceIds:
        - 12345
        - 67890
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing network service group can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:FwNetworkServiceGroup example 12345
`+tripleBacktick("")+`
`)
}

func (a *FwNetworkServiceGroupArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the network service group.")
	ann.Describe(&a.Description, "Additional information about the network service group.")
	ann.Describe(&a.ServiceIds, "IDs of network services that belong to this group.")
}

func (s *FwNetworkServiceGroupState) Annotate(a infer.Annotator) {
	a.Describe(&s.GroupId, "The system-generated ID of the network service group.")
}

var _ infer.CustomResource[FwNetworkServiceGroupArgs, FwNetworkServiceGroupState] = FwNetworkServiceGroup{}
