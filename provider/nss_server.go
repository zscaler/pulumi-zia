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

// Package provider implements the NSS Server resource.
// Adopted from terraform-provider-zia resource_zia_nss_server.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/cloudnss/nss_servers"
)

// NssServer implements the zia:index:NssServer resource.
type NssServer struct{}

// NssServerArgs are the inputs.
type NssServerArgs struct {
	Name      string  `pulumi:"name"`
	Status    *string `pulumi:"status,optional"`
	Type      *string `pulumi:"type,optional"`
	IcapSvrId *int    `pulumi:"icapSvrId,optional"`
}

// NssServerState is the persisted state.
type NssServerState struct {
	NssServerArgs
	NssId *int `pulumi:"nssId"`
}

func nssServerArgsToAPI(args *NssServerArgs, id int) nss_servers.NSSServers {
	status := ptrToString(args.Status)
	if status == "" {
		status = "ENABLED"
	}
	serverType := ptrToString(args.Type)
	if serverType == "" {
		serverType = "NSS_FOR_FIREWALL"
	}
	return nss_servers.NSSServers{
		ID:        id,
		Name:      args.Name,
		Status:    status,
		Type:      serverType,
		IcapSvrId: ptrToIntDefault(args.IcapSvrId, 0),
	}
}

func nssServerAPIToState(api *nss_servers.NSSServers) NssServerState {
	return NssServerState{
		NssServerArgs: NssServerArgs{
			Name:      api.Name,
			Status:    stringPtr(api.Status),
			Type:      stringPtr(api.Type),
			IcapSvrId: intPtr(api.IcapSvrId),
		},
		NssId: intPtr(api.ID),
	}
}

func (NssServer) Create(ctx context.Context, req infer.CreateRequest[NssServerArgs]) (infer.CreateResponse[NssServerState], error) {
	if req.DryRun {
		s := NssServerState{NssServerArgs: req.Inputs, NssId: intPtr(0)}
		return infer.CreateResponse[NssServerState]{ID: "preview", Output: s}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[NssServerState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	service := client.Service

	apiReq := nssServerArgsToAPI(&req.Inputs, 0)
	resp, err := nss_servers.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[NssServerState]{}, err
	}

	time.Sleep(2 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.CreateResponse[NssServerState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := nssServerAPIToState(resp)
	return infer.CreateResponse[NssServerState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (NssServer) Read(ctx context.Context, req infer.ReadRequest[NssServerArgs, NssServerState]) (infer.ReadResponse[NssServerArgs, NssServerState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[NssServerArgs, NssServerState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		server, lookupErr := nss_servers.GetByName(ctx, service, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[NssServerArgs, NssServerState]{}, fmt.Errorf("nss server not found")
			}
			return infer.ReadResponse[NssServerArgs, NssServerState]{}, lookupErr
		}
		id = server.ID
	}

	resp, err := nss_servers.Get(ctx, service, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[NssServerArgs, NssServerState]{}, fmt.Errorf("nss server not found")
		}
		return infer.ReadResponse[NssServerArgs, NssServerState]{}, err
	}

	state := nssServerAPIToState(resp)
	return infer.ReadResponse[NssServerArgs, NssServerState]{
		ID:     req.ID,
		Inputs: state.NssServerArgs,
		State:  state,
	}, nil
}

func (NssServer) Update(ctx context.Context, req infer.UpdateRequest[NssServerArgs, NssServerState]) (infer.UpdateResponse[NssServerState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[NssServerState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	service := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[NssServerState]{}, fmt.Errorf("invalid nss server ID: %s", req.ID)
	}
	apiReq := nssServerArgsToAPI(&req.Inputs, id)

	if _, err := nss_servers.Update(ctx, service, id, &apiReq); err != nil {
		return infer.UpdateResponse[NssServerState]{}, err
	}

	time.Sleep(2 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[NssServerState]{}, activationErr
		}
	}

	updated, err := nss_servers.Get(ctx, service, id)
	if err != nil {
		return infer.UpdateResponse[NssServerState]{Output: NssServerState{
			NssServerArgs: req.Inputs,
			NssId:         intPtr(id),
		}}, nil
	}
	return infer.UpdateResponse[NssServerState]{Output: nssServerAPIToState(updated)}, nil
}

func (NssServer) Delete(ctx context.Context, req infer.DeleteRequest[NssServerState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	service := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid nss server ID: %s", req.ID)
	}
	if _, err := nss_servers.Delete(ctx, service, id); err != nil {
		return infer.DeleteResponse{}, err
	}

	time.Sleep(2 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}
	return infer.DeleteResponse{}, nil
}

func (NssServer) Annotate(a infer.Annotator) {
	describeResource(a, &NssServer{}, `The zia.NssServer resource manages NSS (Nanolog Streaming Service) server configurations in the
Zscaler Internet Access (ZIA) cloud. NSS servers are used to stream logs from ZIA to external SIEM
or log management systems.

{{% examples %}}
## Example Usage

{{% example %}}
### Basic NSS Server

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.NssServer("example", {
    name: "Example NSS Server",
    status: "ENABLED",
    type: "NSS_FOR_FIREWALL",
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.NssServer("example",
    name="Example NSS Server",
    status="ENABLED",
    type="NSS_FOR_FIREWALL",
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:NssServer
    properties:
      name: Example NSS Server
      status: ENABLED
      type: NSS_FOR_FIREWALL
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing NSS server can be imported using its ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:NssServer example 12345
`+tripleBacktick("")+`
`)
}

func (a *NssServerArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "Name of the NSS server.")
	ann.Describe(&a.Status, "The status of the NSS server. Accepted values: 'ENABLED' or 'DISABLED'. Default: 'ENABLED'.")
	ann.Describe(&a.Type, "The NSS server type. Accepted values: 'NSS_FOR_FIREWALL', 'NSS_FOR_WEB'. Default: 'NSS_FOR_FIREWALL'.")
	ann.Describe(&a.IcapSvrId, "The ICAP server ID associated with the NSS server.")
}

func (s *NssServerState) Annotate(a infer.Annotator) {
	a.Describe(&s.NssId, "The unique identifier for the NSS server assigned by the ZIA cloud.")
}

var _ infer.CustomResource[NssServerArgs, NssServerState] = NssServer{}
