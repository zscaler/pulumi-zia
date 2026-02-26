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

// Package provider implements the FW IP Source Group resource.
// Adopted from terraform-provider-zia resource_zia_fw_filtering_ip_source_groups.go.

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
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/firewallpolicies/ipsourcegroups"
)

// FwIpSourceGroup implements the zia:index:FwIpSourceGroup resource.
type FwIpSourceGroup struct{}

// FwIpSourceGroupArgs are the inputs.
type FwIpSourceGroupArgs struct {
	Name        *string  `pulumi:"name,optional"`
	Description *string  `pulumi:"description,optional"`
	IpAddresses []string `pulumi:"ipAddresses,optional"`
}

// FwIpSourceGroupState is the persisted state.
type FwIpSourceGroupState struct {
	FwIpSourceGroupArgs
	GroupId *int `pulumi:"groupId"`
}

func fwIpSourceGroupToAPI(args FwIpSourceGroupArgs) ipsourcegroups.IPSourceGroups {
	return ipsourcegroups.IPSourceGroups{
		Name:        ptrToString(args.Name),
		Description: ptrToString(args.Description),
		IPAddresses: args.IpAddresses,
	}
}

func (FwIpSourceGroup) Create(ctx context.Context, req infer.CreateRequest[FwIpSourceGroupArgs]) (infer.CreateResponse[FwIpSourceGroupState], error) {
	if req.DryRun {
		s := FwIpSourceGroupState{FwIpSourceGroupArgs: req.Inputs, GroupId: intPtr(0)}
		return infer.CreateResponse[FwIpSourceGroupState]{ID: "preview", Output: s}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[FwIpSourceGroupState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := fwIpSourceGroupToAPI(req.Inputs)
	resp, err := ipsourcegroups.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[FwIpSourceGroupState]{}, err
	}
	log.Printf("[INFO] Created ZIA IP source group. ID: %v", resp.ID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[FwIpSourceGroupState]{}, activationErr
		}
	}

	state := FwIpSourceGroupState{
		FwIpSourceGroupArgs: req.Inputs,
		GroupId:             &resp.ID,
	}
	return infer.CreateResponse[FwIpSourceGroupState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (FwIpSourceGroup) Read(ctx context.Context, req infer.ReadRequest[FwIpSourceGroupArgs, FwIpSourceGroupState]) (infer.ReadResponse[FwIpSourceGroupArgs, FwIpSourceGroupState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[FwIpSourceGroupArgs, FwIpSourceGroupState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.GroupId != nil {
		id = *req.State.GroupId
	}
	if id == 0 {
		return infer.ReadResponse[FwIpSourceGroupArgs, FwIpSourceGroupState]{}, fmt.Errorf("no IP source group id in state")
	}

	resp, err := ipsourcegroups.Get(ctx, service, id)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[FwIpSourceGroupArgs, FwIpSourceGroupState]{ID: ""}, nil
		}
		return infer.ReadResponse[FwIpSourceGroupArgs, FwIpSourceGroupState]{}, err
	}

	args := FwIpSourceGroupArgs{
		Name:        stringPtr(resp.Name),
		Description: stringPtr(resp.Description),
		IpAddresses: resp.IPAddresses,
	}
	state := FwIpSourceGroupState{
		FwIpSourceGroupArgs: args,
		GroupId:             &resp.ID,
	}
	return infer.ReadResponse[FwIpSourceGroupArgs, FwIpSourceGroupState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (FwIpSourceGroup) Update(ctx context.Context, req infer.UpdateRequest[FwIpSourceGroupArgs, FwIpSourceGroupState]) (infer.UpdateResponse[FwIpSourceGroupState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[FwIpSourceGroupState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.GroupId != nil {
		id = *req.State.GroupId
	}
	if id == 0 {
		return infer.UpdateResponse[FwIpSourceGroupState]{}, fmt.Errorf("no IP source group id in state")
	}

	if _, err := ipsourcegroups.Get(ctx, service, id); err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[FwIpSourceGroupState]{}, nil
		}
		return infer.UpdateResponse[FwIpSourceGroupState]{}, err
	}

	apiReq := fwIpSourceGroupToAPI(req.Inputs)
	apiReq.ID = id
	if _, err := ipsourcegroups.Update(ctx, service, id, &apiReq); err != nil {
		return infer.UpdateResponse[FwIpSourceGroupState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[FwIpSourceGroupState]{}, activationErr
		}
	}

	state := FwIpSourceGroupState{
		FwIpSourceGroupArgs: req.Inputs,
		GroupId:             &id,
	}
	return infer.UpdateResponse[FwIpSourceGroupState]{Output: state}, nil
}

func (FwIpSourceGroup) Delete(ctx context.Context, req infer.DeleteRequest[FwIpSourceGroupState]) (infer.DeleteResponse, error) {
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
		if err := detachFromFilteringRules(ctx, client, id, "SrcIpGroups",
			func(r *filteringrules.FirewallFilteringRules) []common.IDNameExtensions { return r.SrcIpGroups },
			func(r *filteringrules.FirewallFilteringRules, ids []common.IDNameExtensions) { r.SrcIpGroups = ids }); err != nil {
			return infer.DeleteResponse{}, err
		}
		if _, err := ipsourcegroups.Delete(ctx, service, id); err != nil {
			return infer.DeleteResponse{}, err
		}
		log.Printf("[INFO] ZIA IP source group deleted")
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

func (FwIpSourceGroup) Annotate(a infer.Annotator) {
	describeResource(a, &FwIpSourceGroup{}, `The zia_fw_ip_source_group resource manages firewall IP source groups in the Zscaler Internet Access (ZIA) cloud service. IP source groups allow you to define groups of source IP addresses that can be referenced in firewall filtering rules.

For more information, see the [ZIA Firewall Policies documentation](https://help.zscaler.com/zia/firewall-policies).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic IP Source Group

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.FwIpSourceGroup("example", {
    name: "Example IP Source Group",
    description: "Group of source IPs",
    ipAddresses: ["192.168.1.0/24", "10.0.0.0/8"],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.FwIpSourceGroup("example",
    name="Example IP Source Group",
    description="Group of source IPs",
    ip_addresses=["192.168.1.0/24", "10.0.0.0/8"],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:FwIpSourceGroup
    properties:
      name: Example IP Source Group
      description: Group of source IPs
      ipAddresses:
        - 192.168.1.0/24
        - 10.0.0.0/8
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing IP source group can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:FwIpSourceGroup example 12345
`+tripleBacktick("")+`
`)
}

func (a *FwIpSourceGroupArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the IP source group.")
	ann.Describe(&a.Description, "Additional information about the IP source group.")
	ann.Describe(&a.IpAddresses, "List of source IP addresses or CIDR ranges included in this group.")
}

func (s *FwIpSourceGroupState) Annotate(a infer.Annotator) {
	a.Describe(&s.GroupId, "The system-generated ID of the IP source group.")
}

var _ infer.CustomResource[FwIpSourceGroupArgs, FwIpSourceGroupState] = FwIpSourceGroup{}
