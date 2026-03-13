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

// Package provider implements the Traffic Forwarding VPN Credentials resource.
// Adopted from terraform-provider-zia resource_zia_traffic_forwarding_vpn_credentials.go.

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
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/trafficforwarding/vpncredentials"
)

// TrafficForwardingVpnCredentials implements the zia:index:TrafficForwardingVpnCredentials resource.
type TrafficForwardingVpnCredentials struct{}

// TrafficForwardingVpnCredentialsArgs are the inputs.
type TrafficForwardingVpnCredentialsArgs struct {
	Type         *string `pulumi:"type,optional"`
	Fqdn         *string `pulumi:"fqdn,optional"`
	IPAddress    *string `pulumi:"ipAddress,optional"`
	PreSharedKey *string `pulumi:"preSharedKey,optional" provider:"secret"`
	Comments     *string `pulumi:"comments,optional"`
}

// TrafficForwardingVpnCredentialsState is the persisted state.
type TrafficForwardingVpnCredentialsState struct {
	TrafficForwardingVpnCredentialsArgs
	VpnId *int `pulumi:"vpnId"`
}

func (TrafficForwardingVpnCredentials) Check(ctx context.Context, req infer.CheckRequest) (infer.CheckResponse[TrafficForwardingVpnCredentialsArgs], error) {
	inputs, failures, err := infer.DefaultCheck[TrafficForwardingVpnCredentialsArgs](ctx, req.NewInputs)
	if err != nil {
		return infer.CheckResponse[TrafficForwardingVpnCredentialsArgs]{}, err
	}
	if len(failures) > 0 {
		return infer.CheckResponse[TrafficForwardingVpnCredentialsArgs]{Failures: failures}, nil
	}
	t := ptrToString(inputs.Type)
	if t == "IP" && (inputs.IPAddress == nil || *inputs.IPAddress == "") {
		return infer.CheckResponse[TrafficForwardingVpnCredentialsArgs]{Failures: []p.CheckFailure{{
			Property: "ipAddress",
			Reason:   "ipAddress is required when type is IP",
		}}}, nil
	}
	if t == "UFQDN" && (inputs.Fqdn == nil || *inputs.Fqdn == "") {
		return infer.CheckResponse[TrafficForwardingVpnCredentialsArgs]{Failures: []p.CheckFailure{{
			Property: "fqdn",
			Reason:   "fqdn is required when type is UFQDN",
		}}}, nil
	}
	if inputs.Comments != nil && len(*inputs.Comments) > 10240 {
		return infer.CheckResponse[TrafficForwardingVpnCredentialsArgs]{Failures: []p.CheckFailure{{
			Property: "comments",
			Reason:   "comments must be at most 10240 characters",
		}}}, nil
	}
	return infer.CheckResponse[TrafficForwardingVpnCredentialsArgs]{Inputs: inputs}, nil
}

func (TrafficForwardingVpnCredentials) Create(ctx context.Context, req infer.CreateRequest[TrafficForwardingVpnCredentialsArgs]) (infer.CreateResponse[TrafficForwardingVpnCredentialsState], error) {
	if req.DryRun {
		s := TrafficForwardingVpnCredentialsState{TrafficForwardingVpnCredentialsArgs: req.Inputs, VpnId: intPtr(0)}
		return infer.CreateResponse[TrafficForwardingVpnCredentialsState]{ID: "preview", Output: s}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[TrafficForwardingVpnCredentialsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := trafficForwardingVpnCredentialsArgsToAPI(req.Inputs, 0)
	resp, _, err := vpncredentials.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[TrafficForwardingVpnCredentialsState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[TrafficForwardingVpnCredentialsState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := TrafficForwardingVpnCredentialsState{
		TrafficForwardingVpnCredentialsArgs: req.Inputs,
		VpnId:                               intPtr(resp.ID),
	}
	return infer.CreateResponse[TrafficForwardingVpnCredentialsState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (TrafficForwardingVpnCredentials) Read(ctx context.Context, req infer.ReadRequest[TrafficForwardingVpnCredentialsArgs, TrafficForwardingVpnCredentialsState]) (infer.ReadResponse[TrafficForwardingVpnCredentialsArgs, TrafficForwardingVpnCredentialsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[TrafficForwardingVpnCredentialsArgs, TrafficForwardingVpnCredentialsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.ReadResponse[TrafficForwardingVpnCredentialsArgs, TrafficForwardingVpnCredentialsState]{}, fmt.Errorf("invalid vpn credential id: %w", err)
	}

	allCreds, err := vpncredentials.GetAll(ctx, service)
	if err != nil {
		return infer.ReadResponse[TrafficForwardingVpnCredentialsArgs, TrafficForwardingVpnCredentialsState]{}, err
	}

	var resp *vpncredentials.VPNCredentials
	for i := range allCreds {
		if allCreds[i].ID == id {
			resp = &allCreds[i]
			break
		}
	}

	if resp == nil {
		cred, err := vpncredentials.Get(ctx, service, id)
		if err != nil {
			if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[TrafficForwardingVpnCredentialsArgs, TrafficForwardingVpnCredentialsState]{}, nil
			}
			return infer.ReadResponse[TrafficForwardingVpnCredentialsArgs, TrafficForwardingVpnCredentialsState]{}, err
		}
		resp = cred
	}

	state := trafficForwardingVpnCredentialsToState(resp)
	// The API never returns preSharedKey; preserve it from the prior state
	// so that a refresh does not show a spurious diff.
	if req.State.PreSharedKey != nil {
		state.PreSharedKey = req.State.PreSharedKey
	}
	args := state.TrafficForwardingVpnCredentialsArgs
	return infer.ReadResponse[TrafficForwardingVpnCredentialsArgs, TrafficForwardingVpnCredentialsState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (TrafficForwardingVpnCredentials) Update(ctx context.Context, req infer.UpdateRequest[TrafficForwardingVpnCredentialsArgs, TrafficForwardingVpnCredentialsState]) (infer.UpdateResponse[TrafficForwardingVpnCredentialsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[TrafficForwardingVpnCredentialsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[TrafficForwardingVpnCredentialsState]{}, fmt.Errorf("invalid vpn credential id: %w", err)
	}

	_, getErr := vpncredentials.Get(ctx, service, id)
	if getErr != nil {
		if respErr, ok := getErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.UpdateResponse[TrafficForwardingVpnCredentialsState]{}, nil
		}
		return infer.UpdateResponse[TrafficForwardingVpnCredentialsState]{}, getErr
	}

	apiReq := trafficForwardingVpnCredentialsArgsToAPI(req.Inputs, id)
	if _, _, err := vpncredentials.Update(ctx, service, id, &apiReq); err != nil {
		return infer.UpdateResponse[TrafficForwardingVpnCredentialsState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[TrafficForwardingVpnCredentialsState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := TrafficForwardingVpnCredentialsState{
		TrafficForwardingVpnCredentialsArgs: req.Inputs,
		VpnId:                               intPtr(id),
	}
	return infer.UpdateResponse[TrafficForwardingVpnCredentialsState]{Output: state}, nil
}

func (TrafficForwardingVpnCredentials) Delete(ctx context.Context, req infer.DeleteRequest[TrafficForwardingVpnCredentialsState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid vpn credential id: %w", err)
	}

	if err := vpncredentials.Delete(ctx, service, id); err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.DeleteResponse{}, nil
		}
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

func (TrafficForwardingVpnCredentials) Diff(ctx context.Context, req infer.DiffRequest[TrafficForwardingVpnCredentialsArgs, TrafficForwardingVpnCredentialsState]) (infer.DiffResponse, error) {
	diff, hasChanges := diffSkippingUnsetInputs(req.State.TrafficForwardingVpnCredentialsArgs, req.Inputs)
	return infer.DiffResponse{HasChanges: hasChanges, DetailedDiff: diff}, nil
}

func (TrafficForwardingVpnCredentials) Annotate(a infer.Annotator) {
	describeResource(a, &TrafficForwardingVpnCredentials{}, `The zia_traffic_forwarding_vpn_credentials resource manages VPN credentials for traffic forwarding in the Zscaler Internet Access (ZIA) cloud service. VPN credentials are used to authenticate IPSec VPN tunnels between on-premises equipment and the Zscaler cloud.

For more information, see the [ZIA Traffic Forwarding documentation](https://help.zscaler.com/zia/traffic-forwarding).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic VPN Credentials (UFQDN)

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

import * as pulumi from "@pulumi/pulumi";

const cfg = new pulumi.Config();
const vpnPreSharedKey = cfg.requireSecret("vpnPreSharedKey");

const example = new zia.TrafficForwardingVpnCredentials("example", {
    type: "UFQDN",
    fqdn: "user@example.com",
    preSharedKey: vpnPreSharedKey,
    comments: "Branch office VPN credentials",
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import pulumi
import zscaler_pulumi_zia as zia

cfg = pulumi.Config()
vpn_pre_shared_key = cfg.require_secret("vpnPreSharedKey")

example = zia.TrafficForwardingVpnCredentials("example",
    type="UFQDN",
    fqdn="user@example.com",
    pre_shared_key=vpn_pre_shared_key,
    comments="Branch office VPN credentials",
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:TrafficForwardingVpnCredentials
    properties:
      type: UFQDN
      fqdn: user@example.com
      preSharedKey:
        fn::secret: ${vpnPreSharedKey}
      comments: Branch office VPN credentials
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing VPN credential can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:TrafficForwardingVpnCredentials example 12345
`+tripleBacktick("")+`
`)
}

func (a *TrafficForwardingVpnCredentialsArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Type, "VPN credential type. Valid values: `IP`, `UFQDN`.")
	ann.Describe(&a.Fqdn, "Fully Qualified Domain Name (FQDN). Required when type is `UFQDN`.")
	ann.Describe(&a.IPAddress, "The static IP address associated with the VPN credential. Required when type is `IP`.")
	ann.Describe(&a.PreSharedKey, "Pre-shared key (PSK) for the VPN credential. This is a secret value.")
	ann.Describe(&a.Comments, "Additional information about the VPN credential. Maximum 10240 characters.")
}

func (s *TrafficForwardingVpnCredentialsState) Annotate(a infer.Annotator) {
	a.Describe(&s.VpnId, "The system-generated ID of the VPN credential.")
}

func trafficForwardingVpnCredentialsArgsToAPI(args TrafficForwardingVpnCredentialsArgs, id int) vpncredentials.VPNCredentials {
	return vpncredentials.VPNCredentials{
		ID:           id,
		Type:         ptrToString(args.Type),
		FQDN:         ptrToString(args.Fqdn),
		IPAddress:    ptrToString(args.IPAddress),
		PreSharedKey: ptrToString(args.PreSharedKey),
		Comments:     ptrToString(args.Comments),
	}
}

func trafficForwardingVpnCredentialsToState(r *vpncredentials.VPNCredentials) TrafficForwardingVpnCredentialsState {
	return TrafficForwardingVpnCredentialsState{
		VpnId: intPtr(r.ID),
		TrafficForwardingVpnCredentialsArgs: TrafficForwardingVpnCredentialsArgs{
			Type:      stringPtr(r.Type),
			Fqdn:      stringPtr(r.FQDN),
			IPAddress: stringPtr(r.IPAddress),
			Comments:  stringPtr(r.Comments),
		},
	}
}

var _ infer.CustomResource[TrafficForwardingVpnCredentialsArgs, TrafficForwardingVpnCredentialsState] = TrafficForwardingVpnCredentials{}
