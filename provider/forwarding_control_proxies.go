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

// Package provider implements the Forwarding Control Proxies resource.
// Adopted from terraform-provider-zia resource_zia_forwarding_control_proxies.go.

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
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/forwarding_control_policy/proxies"
)

// ForwardingControlProxies implements the zia:index:ForwardingControlProxies resource.
type ForwardingControlProxies struct{}

// ForwardingControlProxiesArgs are the inputs.
type ForwardingControlProxiesArgs struct {
	Name                  *string `pulumi:"name,optional"`
	Description           *string `pulumi:"description,optional"`
	Type                  *string `pulumi:"type,optional"` // PROXYCHAIN, ZIA, ECSELF
	Address               *string `pulumi:"address,optional"`
	Port                  *int    `pulumi:"port,optional"`
	InsertXauHeader       *bool   `pulumi:"insertXauHeader,optional"`
	Base64EncodeXauHeader *bool   `pulumi:"base64EncodeXauHeader,optional"`
	CertId                *int    `pulumi:"certId,optional"`
}

// ForwardingControlProxiesState is the persisted state.
type ForwardingControlProxiesState struct {
	ForwardingControlProxiesArgs
	ProxyId *int `pulumi:"proxyId"`
}

func forwardingControlProxiesToAPI(args ForwardingControlProxiesArgs, id int) proxies.Proxies {
	out := proxies.Proxies{
		ID:                    id,
		Name:                  ptrToString(args.Name),
		Description:           ptrToString(args.Description),
		Type:                  ptrToString(args.Type),
		Address:               ptrToString(args.Address),
		Port:                  ptrToIntDefault(args.Port, 0),
		InsertXauHeader:       ptrToBool(args.InsertXauHeader),
		Base64EncodeXauHeader: ptrToBool(args.Base64EncodeXauHeader),
	}
	if args.CertId != nil && *args.CertId != 0 {
		out.Cert = &common.IDNameExternalID{ID: *args.CertId}
	}
	return out
}

func (ForwardingControlProxies) Create(ctx context.Context, req infer.CreateRequest[ForwardingControlProxiesArgs]) (infer.CreateResponse[ForwardingControlProxiesState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[ForwardingControlProxiesState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := forwardingControlProxiesToAPI(req.Inputs, 0)
	resp, _, err := proxies.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[ForwardingControlProxiesState]{}, err
	}
	log.Printf("[INFO] Created ZIA Forwarding Control Proxy. ID: %v", resp.ID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[ForwardingControlProxiesState]{}, activationErr
		}
	}

	state := ForwardingControlProxiesState{
		ForwardingControlProxiesArgs: req.Inputs,
		ProxyId:                      &resp.ID,
	}
	return infer.CreateResponse[ForwardingControlProxiesState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (ForwardingControlProxies) Read(ctx context.Context, req infer.ReadRequest[ForwardingControlProxiesArgs, ForwardingControlProxiesState]) (infer.ReadResponse[ForwardingControlProxiesArgs, ForwardingControlProxiesState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[ForwardingControlProxiesArgs, ForwardingControlProxiesState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.ProxyId != nil {
		id = *req.State.ProxyId
	}
	if id == 0 {
		return infer.ReadResponse[ForwardingControlProxiesArgs, ForwardingControlProxiesState]{}, fmt.Errorf("no proxy id in state")
	}

	resp, err := proxies.Get(ctx, service, id)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[ForwardingControlProxiesArgs, ForwardingControlProxiesState]{ID: ""}, nil
		}
		return infer.ReadResponse[ForwardingControlProxiesArgs, ForwardingControlProxiesState]{}, err
	}

	certId := (*int)(nil)
	if resp.Cert != nil && resp.Cert.ID != 0 {
		certId = &resp.Cert.ID
	}

	args := ForwardingControlProxiesArgs{
		Name:                  stringPtr(resp.Name),
		Description:           stringPtr(resp.Description),
		Type:                  stringPtr(resp.Type),
		Address:               stringPtr(resp.Address),
		Port:                  intPtr(resp.Port),
		InsertXauHeader:       boolPtr(resp.InsertXauHeader),
		Base64EncodeXauHeader: boolPtr(resp.Base64EncodeXauHeader),
		CertId:                certId,
	}
	state := ForwardingControlProxiesState{
		ForwardingControlProxiesArgs: args,
		ProxyId:                      &resp.ID,
	}
	return infer.ReadResponse[ForwardingControlProxiesArgs, ForwardingControlProxiesState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (ForwardingControlProxies) Update(ctx context.Context, req infer.UpdateRequest[ForwardingControlProxiesArgs, ForwardingControlProxiesState]) (infer.UpdateResponse[ForwardingControlProxiesState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[ForwardingControlProxiesState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.ProxyId != nil {
		id = *req.State.ProxyId
	}
	if id == 0 {
		return infer.UpdateResponse[ForwardingControlProxiesState]{}, fmt.Errorf("no proxy id in state")
	}

	if _, err := proxies.Get(ctx, service, id); err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[ForwardingControlProxiesState]{}, nil
		}
		return infer.UpdateResponse[ForwardingControlProxiesState]{}, err
	}

	apiReq := forwardingControlProxiesToAPI(req.Inputs, id)
	if _, _, err := proxies.Update(ctx, service, id, &apiReq); err != nil {
		return infer.UpdateResponse[ForwardingControlProxiesState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[ForwardingControlProxiesState]{}, activationErr
		}
	}

	state := ForwardingControlProxiesState{
		ForwardingControlProxiesArgs: req.Inputs,
		ProxyId:                      &id,
	}
	return infer.UpdateResponse[ForwardingControlProxiesState]{Output: state}, nil
}

func (ForwardingControlProxies) Delete(ctx context.Context, req infer.DeleteRequest[ForwardingControlProxiesState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.ProxyId != nil {
		id = *req.State.ProxyId
	}
	if id != 0 {
		if _, err := proxies.Delete(ctx, service, id); err != nil {
			return infer.DeleteResponse{}, err
		}
		log.Printf("[INFO] ZIA Forwarding Control Proxy deleted")
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

func (ForwardingControlProxies) Annotate(a infer.Annotator) {
	describeResource(a, &ForwardingControlProxies{}, `The zia.ForwardingControlProxies resource manages forwarding control proxy configurations in the
Zscaler Internet Access (ZIA) cloud. Proxies are used in forwarding control rules to direct traffic
through proxy chains, ZIA, or EC-self proxy types.

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Forwarding Control Proxy

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.ForwardingControlProxies("example", {
    name: "Example Proxy",
    description: "Managed by Pulumi",
    type: "PROXYCHAIN",
    address: "proxy.example.com",
    port: 8080,
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.ForwardingControlProxies("example",
    name="Example Proxy",
    description="Managed by Pulumi",
    type="PROXYCHAIN",
    address="proxy.example.com",
    port=8080,
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:ForwardingControlProxies
    properties:
      name: Example Proxy
      description: Managed by Pulumi
      type: PROXYCHAIN
      address: proxy.example.com
      port: 8080
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing forwarding control proxy can be imported using its ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:ForwardingControlProxies example 12345
`+tripleBacktick("")+`
`)
}

func (a *ForwardingControlProxiesArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "Name of the forwarding control proxy.")
	ann.Describe(&a.Description, "Description of the forwarding control proxy.")
	ann.Describe(&a.Type, "The proxy type. Accepted values: 'PROXYCHAIN', 'ZIA', 'ECSELF'.")
	ann.Describe(&a.Address, "The address of the proxy server.")
	ann.Describe(&a.Port, "The port number of the proxy server.")
	ann.Describe(&a.InsertXauHeader, "Whether to insert the X-Authenticated-User header.")
	ann.Describe(&a.Base64EncodeXauHeader, "Whether to base64-encode the X-Authenticated-User header.")
	ann.Describe(&a.CertId, "The certificate ID used for the proxy.")
}

func (s *ForwardingControlProxiesState) Annotate(a infer.Annotator) {
	a.Describe(&s.ProxyId, "The unique identifier for the proxy assigned by the ZIA cloud.")
}

var _ infer.CustomResource[ForwardingControlProxiesArgs, ForwardingControlProxiesState] = ForwardingControlProxies{}
