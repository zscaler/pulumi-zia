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

// Package provider implements the FW IP Destination Group resource.
// Adopted from terraform-provider-zia resource_zia_fw_filtering_ip_destination_groups.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/common"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/firewallpolicies/filteringrules"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/firewallpolicies/ipdestinationgroups"
)

// FwIpDestinationGroup implements the zia:index:FwIpDestinationGroup resource.
type FwIpDestinationGroup struct{}

// FwIpDestinationGroupArgs are the inputs.
type FwIpDestinationGroupArgs struct {
	Name         *string  `pulumi:"name,optional"`
	Description  *string  `pulumi:"description,optional"`
	Type         *string  `pulumi:"type,optional"`
	Addresses    []string `pulumi:"addresses,optional"`
	IpCategories []string `pulumi:"ipCategories,optional"`
	Countries    []string `pulumi:"countries,optional"`
}

// FwIpDestinationGroupState is the persisted state.
type FwIpDestinationGroupState struct {
	FwIpDestinationGroupArgs
	GroupId *int `pulumi:"groupId"`
}

func fwIpDestinationGroupToAPI(args FwIpDestinationGroupArgs, id int) ipdestinationgroups.IPDestinationGroups {
	countries := processCountries(args.Countries)
	return ipdestinationgroups.IPDestinationGroups{
		ID:           id,
		Name:         ptrToString(args.Name),
		Description:  ptrToString(args.Description),
		Type:         ptrToString(args.Type),
		Addresses:    args.Addresses,
		IPCategories: args.IpCategories,
		Countries:    countries,
	}
}

func countriesFromAPI(countries []string) []string {
	if len(countries) == 0 {
		return nil
	}
	result := make([]string, len(countries))
	for i, c := range countries {
		result[i] = strings.TrimPrefix(c, "COUNTRY_")
	}
	return result
}

func (FwIpDestinationGroup) Create(ctx context.Context, req infer.CreateRequest[FwIpDestinationGroupArgs]) (infer.CreateResponse[FwIpDestinationGroupState], error) {
	if req.DryRun {
		s := FwIpDestinationGroupState{FwIpDestinationGroupArgs: req.Inputs, GroupId: intPtr(0)}
		return infer.CreateResponse[FwIpDestinationGroupState]{ID: "preview", Output: s}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[FwIpDestinationGroupState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := fwIpDestinationGroupToAPI(req.Inputs, 0)
	resp, err := ipdestinationgroups.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[FwIpDestinationGroupState]{}, err
	}
	log.Printf("[INFO] Created ZIA IP destination group. ID: %v", resp.ID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[FwIpDestinationGroupState]{}, activationErr
		}
	}

	state := FwIpDestinationGroupState{
		FwIpDestinationGroupArgs: req.Inputs,
		GroupId:                  &resp.ID,
	}
	return infer.CreateResponse[FwIpDestinationGroupState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (FwIpDestinationGroup) Read(ctx context.Context, req infer.ReadRequest[FwIpDestinationGroupArgs, FwIpDestinationGroupState]) (infer.ReadResponse[FwIpDestinationGroupArgs, FwIpDestinationGroupState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[FwIpDestinationGroupArgs, FwIpDestinationGroupState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.GroupId != nil {
		id = *req.State.GroupId
	}
	if id == 0 {
		return infer.ReadResponse[FwIpDestinationGroupArgs, FwIpDestinationGroupState]{}, fmt.Errorf("no IP destination group id in state")
	}

	resp, err := ipdestinationgroups.Get(ctx, service, id)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[FwIpDestinationGroupArgs, FwIpDestinationGroupState]{ID: ""}, nil
		}
		return infer.ReadResponse[FwIpDestinationGroupArgs, FwIpDestinationGroupState]{}, err
	}

	args := FwIpDestinationGroupArgs{
		Name:         stringPtr(resp.Name),
		Description:  stringPtr(resp.Description),
		Type:         stringPtr(resp.Type),
		Addresses:    resp.Addresses,
		IpCategories: resp.IPCategories,
		Countries:    countriesFromAPI(resp.Countries),
	}
	state := FwIpDestinationGroupState{
		FwIpDestinationGroupArgs: args,
		GroupId:                  &resp.ID,
	}
	return infer.ReadResponse[FwIpDestinationGroupArgs, FwIpDestinationGroupState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (FwIpDestinationGroup) Update(ctx context.Context, req infer.UpdateRequest[FwIpDestinationGroupArgs, FwIpDestinationGroupState]) (infer.UpdateResponse[FwIpDestinationGroupState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[FwIpDestinationGroupState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.GroupId != nil {
		id = *req.State.GroupId
	}
	if id == 0 {
		return infer.UpdateResponse[FwIpDestinationGroupState]{}, fmt.Errorf("no IP destination group id in state")
	}

	if _, err := ipdestinationgroups.Get(ctx, service, id); err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[FwIpDestinationGroupState]{}, nil
		}
		return infer.UpdateResponse[FwIpDestinationGroupState]{}, err
	}

	apiReq := fwIpDestinationGroupToAPI(req.Inputs, id)
	if _, _, err := ipdestinationgroups.Update(ctx, service, id, &apiReq, nil); err != nil {
		return infer.UpdateResponse[FwIpDestinationGroupState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[FwIpDestinationGroupState]{}, activationErr
		}
	}

	state := FwIpDestinationGroupState{
		FwIpDestinationGroupArgs: req.Inputs,
		GroupId:                  &id,
	}
	return infer.UpdateResponse[FwIpDestinationGroupState]{Output: state}, nil
}

func (FwIpDestinationGroup) Delete(ctx context.Context, req infer.DeleteRequest[FwIpDestinationGroupState]) (infer.DeleteResponse, error) {
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
		if err := detachFromFilteringRules(ctx, client, id, "DestIpGroups",
			func(r *filteringrules.FirewallFilteringRules) []common.IDNameExtensions { return r.DestIpGroups },
			func(r *filteringrules.FirewallFilteringRules, ids []common.IDNameExtensions) { r.DestIpGroups = ids }); err != nil {
			return infer.DeleteResponse{}, err
		}
		if _, err := ipdestinationgroups.Delete(ctx, service, id); err != nil {
			if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.DeleteResponse{}, nil
			}
			return infer.DeleteResponse{}, err
		}
		log.Printf("[INFO] ZIA IP destination group deleted")
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

func (FwIpDestinationGroup) Annotate(a infer.Annotator) {
	describeResource(a, &FwIpDestinationGroup{}, `The zia_fw_ip_destination_group resource manages firewall IP destination groups in the Zscaler Internet Access (ZIA) cloud service. IP destination groups allow you to define groups of destination IP addresses, FQDNs, or countries that can be referenced in firewall filtering rules.

For more information, see the [ZIA Firewall Policies documentation](https://help.zscaler.com/zia/firewall-policies).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic IP Destination Group

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.FwIpDestinationGroup("example", {
    name: "Example IP Destination Group",
    description: "Group of destination IPs",
    type: "DSTN_IP",
    addresses: ["203.0.113.0/24", "198.51.100.0/24"],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.FwIpDestinationGroup("example",
    name="Example IP Destination Group",
    description="Group of destination IPs",
    type="DSTN_IP",
    addresses=["203.0.113.0/24", "198.51.100.0/24"],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:FwIpDestinationGroup
    properties:
      name: Example IP Destination Group
      description: Group of destination IPs
      type: DSTN_IP
      addresses:
        - 203.0.113.0/24
        - 198.51.100.0/24
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing IP destination group can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:FwIpDestinationGroup example 12345
`+tripleBacktick("")+`
`)
}

func (a *FwIpDestinationGroupArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the IP destination group.")
	ann.Describe(&a.Description, "Additional information about the IP destination group.")
	ann.Describe(&a.Type, "Destination group type. Valid values: `DSTN_IP`, `DSTN_FQDN`, `DSTN_DOMAIN`, `DSTN_OTHER`.")
	ann.Describe(&a.Addresses, "List of destination IP addresses, FQDNs, or wildcard FQDNs in this group.")
	ann.Describe(&a.IpCategories, "List of URL/IP categories allowed for this group.")
	ann.Describe(&a.Countries, "List of destination countries (ISO 3166-1 alpha-2 codes). The COUNTRY_ prefix is added automatically.")
}

func (s *FwIpDestinationGroupState) Annotate(a infer.Annotator) {
	a.Describe(&s.GroupId, "The system-generated ID of the IP destination group.")
}

var _ infer.CustomResource[FwIpDestinationGroupArgs, FwIpDestinationGroupState] = FwIpDestinationGroup{}
