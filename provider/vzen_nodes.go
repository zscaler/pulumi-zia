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

// Package provider implements the VZEN Node resource and invoke.
// Adopted from terraform-provider-zia resource_zia_vzen_nodes.go and data_source_zia_vzen_nodes.go.

package provider

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/vzen_nodes"
)

// --- Resource ---

// VzenNode implements the zia:index:VzenNode resource.
type VzenNode struct{}

// VzenNodeArgs are the inputs for VzenNode.
type VzenNodeArgs struct {
	Name                          *string `pulumi:"name,optional"`
	Status                        *string `pulumi:"status,optional"`
	Type                          *string `pulumi:"type,optional"`
	IpSecEnabled                  *bool   `pulumi:"ipSecEnabled,optional"`
	IpAddress                     *string `pulumi:"ipAddress,optional"`
	SubnetMask                    *string `pulumi:"subnetMask,optional"`
	DefaultGateway                *string `pulumi:"defaultGateway,optional"`
	InProduction                  *bool   `pulumi:"inProduction,optional"`
	OnDemandSupportTunnelEnabled  *bool   `pulumi:"onDemandSupportTunnelEnabled,optional"`
	EstablishSupportTunnelEnabled *bool   `pulumi:"establishSupportTunnelEnabled,optional"`
	LoadBalancerIpAddress         *string `pulumi:"loadBalancerIpAddress,optional"`
	DeploymentMode                *string `pulumi:"deploymentMode,optional"`
	ClusterName                   *string `pulumi:"clusterName,optional"`
	VzenSkuType                   *string `pulumi:"vzenSkuType,optional"`
}

// VzenNodeState is the persisted state.
type VzenNodeState struct {
	VzenNodeArgs
	NodeId *int `pulumi:"nodeId"`
}

func (VzenNode) Check(ctx context.Context, req infer.CheckRequest) (infer.CheckResponse[VzenNodeArgs], error) {
	inputs, failures, err := infer.DefaultCheck[VzenNodeArgs](ctx, req.NewInputs)
	if err != nil {
		return infer.CheckResponse[VzenNodeArgs]{}, err
	}
	if len(failures) > 0 {
		return infer.CheckResponse[VzenNodeArgs]{Failures: failures}, nil
	}
	// Status validation
	if inputs.Status != nil {
		valid := map[string]bool{
			"ENABLED": true, "DISABLED": true, "DISABLED_BY_SERVICE_PROVIDER": true,
			"NOT_PROVISIONED_IN_SERVICE_PROVIDER": true, "IN_TRIAL": true,
		}
		if !valid[*inputs.Status] {
			return infer.CheckResponse[VzenNodeArgs]{Failures: []p.CheckFailure{{
				Property: "status",
				Reason:   "status must be ENABLED, DISABLED, DISABLED_BY_SERVICE_PROVIDER, NOT_PROVISIONED_IN_SERVICE_PROVIDER, or IN_TRIAL",
			}}}, nil
		}
	}
	// IP address validation
	if inputs.IpAddress != nil && *inputs.IpAddress != "" {
		if net.ParseIP(*inputs.IpAddress) == nil {
			return infer.CheckResponse[VzenNodeArgs]{Failures: []p.CheckFailure{{
				Property: "ipAddress",
				Reason:   "ipAddress must be a valid IP address",
			}}}, nil
		}
	}
	if inputs.DefaultGateway != nil && *inputs.DefaultGateway != "" {
		if net.ParseIP(*inputs.DefaultGateway) == nil {
			return infer.CheckResponse[VzenNodeArgs]{Failures: []p.CheckFailure{{
				Property: "defaultGateway",
				Reason:   "defaultGateway must be a valid IP address",
			}}}, nil
		}
	}
	if inputs.LoadBalancerIpAddress != nil && *inputs.LoadBalancerIpAddress != "" {
		if net.ParseIP(*inputs.LoadBalancerIpAddress) == nil {
			return infer.CheckResponse[VzenNodeArgs]{Failures: []p.CheckFailure{{
				Property: "loadBalancerIpAddress",
				Reason:   "loadBalancerIpAddress must be a valid IP address",
			}}}, nil
		}
	}
	// DeploymentMode validation
	if inputs.DeploymentMode != nil {
		if *inputs.DeploymentMode != "STANDALONE" && *inputs.DeploymentMode != "CLUSTER" {
			return infer.CheckResponse[VzenNodeArgs]{Failures: []p.CheckFailure{{
				Property: "deploymentMode",
				Reason:   "deploymentMode must be STANDALONE or CLUSTER",
			}}}, nil
		}
	}
	// VzenSkuType validation
	if inputs.VzenSkuType != nil {
		if *inputs.VzenSkuType != "SMALL" && *inputs.VzenSkuType != "MEDIUM" && *inputs.VzenSkuType != "LARGE" {
			return infer.CheckResponse[VzenNodeArgs]{Failures: []p.CheckFailure{{
				Property: "vzenSkuType",
				Reason:   "vzenSkuType must be SMALL, MEDIUM, or LARGE",
			}}}, nil
		}
	}
	return infer.CheckResponse[VzenNodeArgs]{Inputs: inputs}, nil
}

func (VzenNode) Diff(ctx context.Context, req infer.DiffRequest[VzenNodeArgs, VzenNodeState]) (infer.DiffResponse, error) {
	diff, hasChanges := diffSkippingUnsetInputs(req.State.VzenNodeArgs, req.Inputs)
	return infer.DiffResponse{HasChanges: hasChanges, DetailedDiff: diff}, nil
}

func (VzenNode) Create(ctx context.Context, req infer.CreateRequest[VzenNodeArgs]) (infer.CreateResponse[VzenNodeState], error) {
	if req.DryRun {
		return infer.CreateResponse[VzenNodeState]{
			ID: "preview",
			Output: VzenNodeState{
				VzenNodeArgs: req.Inputs,
				NodeId:       intPtr(0),
			},
		}, nil
	}

	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[VzenNodeState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := vzenNodeArgsToAPI(req.Inputs, 0)
	resp, _, err := vzen_nodes.Create(ctx, svc, &apiReq)
	if err != nil {
		return infer.CreateResponse[VzenNodeState]{}, fmt.Errorf("creating vzen node: %w", err)
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.CreateResponse[VzenNodeState]{}, activationErr
		}
	}

	rule, err := vzen_nodes.Get(ctx, svc, resp.ID)
	if err != nil {
		return infer.CreateResponse[VzenNodeState]{
			ID:     strconv.Itoa(resp.ID),
			Output: VzenNodeState{VzenNodeArgs: req.Inputs, NodeId: &resp.ID},
		}, nil
	}
	return infer.CreateResponse[VzenNodeState]{
		ID:     strconv.Itoa(resp.ID),
		Output: vzenNodeAPIToState(rule),
	}, nil
}

func (VzenNode) Read(ctx context.Context, req infer.ReadRequest[VzenNodeArgs, VzenNodeState]) (infer.ReadResponse[VzenNodeArgs, VzenNodeState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[VzenNodeArgs, VzenNodeState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		rule, lookupErr := vzen_nodes.GetNodeByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[VzenNodeArgs, VzenNodeState]{}, fmt.Errorf("vzen node not found")
			}
			return infer.ReadResponse[VzenNodeArgs, VzenNodeState]{}, lookupErr
		}
		id = rule.ID
	}

	rule, err := vzen_nodes.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[VzenNodeArgs, VzenNodeState]{}, fmt.Errorf("vzen node not found")
		}
		return infer.ReadResponse[VzenNodeArgs, VzenNodeState]{}, err
	}

	state := vzenNodeAPIToState(rule)
	args := vzenNodeStateToArgs(rule)
	return infer.ReadResponse[VzenNodeArgs, VzenNodeState]{
		ID:     strconv.Itoa(rule.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (VzenNode) Update(ctx context.Context, req infer.UpdateRequest[VzenNodeArgs, VzenNodeState]) (infer.UpdateResponse[VzenNodeState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[VzenNodeState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[VzenNodeState]{}, fmt.Errorf("invalid node id: %w", err)
	}

	apiReq := vzenNodeArgsToAPI(req.Inputs, id)
	if _, _, err := vzen_nodes.Update(ctx, svc, id, &apiReq); err != nil {
		return infer.UpdateResponse[VzenNodeState]{}, fmt.Errorf("updating vzen node: %w", err)
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[VzenNodeState]{}, activationErr
		}
	}

	rule, err := vzen_nodes.Get(ctx, svc, id)
	if err != nil {
		return infer.UpdateResponse[VzenNodeState]{
			Output: VzenNodeState{VzenNodeArgs: req.Inputs, NodeId: &id},
		}, nil
	}
	return infer.UpdateResponse[VzenNodeState]{Output: vzenNodeAPIToState(rule)}, nil
}

func (VzenNode) Delete(ctx context.Context, req infer.DeleteRequest[VzenNodeState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid node id: %w", err)
	}

	if _, err := vzen_nodes.Delete(ctx, svc, id); err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("deleting vzen node: %w", err)
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}
	return infer.DeleteResponse{}, nil
}

func (VzenNode) Annotate(a infer.Annotator) {
	describeResource(a, &VzenNode{}, `The zia.VzenNode resource manages Virtual ZEN (VZEN) node configurations in the Zscaler Internet Access (ZIA) cloud.
VZEN nodes are virtual appliances deployed on-premises to process traffic locally before forwarding to the ZIA cloud.

{{% examples %}}
## Example Usage

{{% example %}}
### Basic VZEN Node

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.VzenNode("example", {
    name: "Example VZEN Node",
    status: "ENABLED",
    type: "VZEN",
    ipAddress: "10.0.0.10",
    subnetMask: "255.255.255.0",
    defaultGateway: "10.0.0.1",
    deploymentMode: "STANDALONE",
    vzenSkuType: "MEDIUM",
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.VzenNode("example",
    name="Example VZEN Node",
    status="ENABLED",
    type="VZEN",
    ip_address="10.0.0.10",
    subnet_mask="255.255.255.0",
    default_gateway="10.0.0.1",
    deployment_mode="STANDALONE",
    vzen_sku_type="MEDIUM",
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:VzenNode
    properties:
      name: Example VZEN Node
      status: ENABLED
      type: VZEN
      ipAddress: "10.0.0.10"
      subnetMask: "255.255.255.0"
      defaultGateway: "10.0.0.1"
      deploymentMode: STANDALONE
      vzenSkuType: MEDIUM
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing VZEN node can be imported using its ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:VzenNode example 12345
`+tripleBacktick("")+`
`)
}

func (a *VzenNodeArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "Name of the VZEN node.")
	ann.Describe(&a.Status, "The status of the node. Accepted values: 'ENABLED', 'DISABLED', 'DISABLED_BY_SERVICE_PROVIDER', 'NOT_PROVISIONED_IN_SERVICE_PROVIDER', 'IN_TRIAL'.")
	ann.Describe(&a.Type, "The type of the VZEN node.")
	ann.Describe(&a.IpSecEnabled, "Whether IPSec is enabled on the node.")
	ann.Describe(&a.IpAddress, "The IP address of the VZEN node.")
	ann.Describe(&a.SubnetMask, "The subnet mask of the VZEN node.")
	ann.Describe(&a.DefaultGateway, "The default gateway of the VZEN node.")
	ann.Describe(&a.InProduction, "Whether the node is in production.")
	ann.Describe(&a.OnDemandSupportTunnelEnabled, "Whether on-demand support tunnel is enabled.")
	ann.Describe(&a.EstablishSupportTunnelEnabled, "Whether establish support tunnel is enabled.")
	ann.Describe(&a.LoadBalancerIpAddress, "The load balancer IP address.")
	ann.Describe(&a.DeploymentMode, "The deployment mode. Accepted values: 'STANDALONE' or 'CLUSTER'.")
	ann.Describe(&a.ClusterName, "The cluster name if deployment mode is CLUSTER.")
	ann.Describe(&a.VzenSkuType, "The VZEN SKU type. Accepted values: 'SMALL', 'MEDIUM', 'LARGE'.")
}

func (s *VzenNodeState) Annotate(a infer.Annotator) {
	a.Describe(&s.NodeId, "The unique identifier for the VZEN node assigned by the ZIA cloud.")
}

func vzenNodeArgsToAPI(in VzenNodeArgs, existingID int) vzen_nodes.VZENNodes {
	return vzen_nodes.VZENNodes{
		ID:                            existingID,
		Name:                          ptrToString(in.Name),
		Status:                        ptrToString(in.Status),
		Type:                          ptrToString(in.Type),
		IPAddress:                     ptrToString(in.IpAddress),
		SubnetMask:                    ptrToString(in.SubnetMask),
		DefaultGateway:                ptrToString(in.DefaultGateway),
		IPSecEnabled:                  ptrToBool(in.IpSecEnabled),
		InProduction:                  ptrToBool(in.InProduction),
		OnDemandSupportTunnelEnabled:  ptrToBool(in.OnDemandSupportTunnelEnabled),
		EstablishSupportTunnelEnabled: ptrToBool(in.EstablishSupportTunnelEnabled),
		LoadBalancerIPAddress:         ptrToString(in.LoadBalancerIpAddress),
		DeploymentMode:                ptrToString(in.DeploymentMode),
		ClusterName:                   ptrToString(in.ClusterName),
		VzenSkuType:                   ptrToString(in.VzenSkuType),
	}
}

func vzenNodeAPIToState(rule *vzen_nodes.VZENNodes) VzenNodeState {
	return VzenNodeState{
		VzenNodeArgs: VzenNodeArgs{
			Name:                          stringPtr(rule.Name),
			Status:                        stringPtr(rule.Status),
			Type:                          stringPtr(rule.Type),
			IpSecEnabled:                  boolPtr(rule.IPSecEnabled),
			IpAddress:                     stringPtr(rule.IPAddress),
			SubnetMask:                    stringPtr(rule.SubnetMask),
			DefaultGateway:                stringPtr(rule.DefaultGateway),
			InProduction:                  boolPtr(rule.InProduction),
			OnDemandSupportTunnelEnabled:  boolPtr(rule.OnDemandSupportTunnelEnabled),
			EstablishSupportTunnelEnabled: boolPtr(rule.EstablishSupportTunnelEnabled),
			LoadBalancerIpAddress:         stringPtr(rule.LoadBalancerIPAddress),
			DeploymentMode:                stringPtr(rule.DeploymentMode),
			ClusterName:                   stringPtr(rule.ClusterName),
			VzenSkuType:                   stringPtr(rule.VzenSkuType),
		},
		NodeId: &rule.ID,
	}
}

func vzenNodeStateToArgs(rule *vzen_nodes.VZENNodes) VzenNodeArgs {
	return vzenNodeAPIToState(rule).VzenNodeArgs
}

// --- Invoke (data source) ---

// GetVzenNodeArgs are the inputs for the GetVzenNode invoke.
type GetVzenNodeArgs struct {
	Id   *int    `pulumi:"nodeId,optional"` // Pulumi reserves "id" in function I/O
	Name *string `pulumi:"name,optional"`
}

// GetVzenNodeResult is the output of the GetVzenNode invoke.
type GetVzenNodeResult struct {
	Id                            int    `pulumi:"nodeId"` // Pulumi reserves "id" in function outputs
	Name                          string `pulumi:"name"`
	Status                        string `pulumi:"status"`
	Type                          string `pulumi:"type"`
	IpSecEnabled                  bool   `pulumi:"ipSecEnabled"`
	IpAddress                     string `pulumi:"ipAddress"`
	SubnetMask                    string `pulumi:"subnetMask"`
	DefaultGateway                string `pulumi:"defaultGateway"`
	ZGatewayId                    int    `pulumi:"zGatewayId"`
	InProduction                  bool   `pulumi:"inProduction"`
	OnDemandSupportTunnelEnabled  bool   `pulumi:"onDemandSupportTunnelEnabled"`
	EstablishSupportTunnelEnabled bool   `pulumi:"establishSupportTunnelEnabled"`
	LoadBalancerIpAddress         string `pulumi:"loadBalancerIpAddress"`
	DeploymentMode                string `pulumi:"deploymentMode"`
	ClusterName                   string `pulumi:"clusterName"`
	VzenSkuType                   string `pulumi:"vzenSkuType"`
}

// GetVzenNode implements the zia:index:GetVzenNode invoke.
type GetVzenNode struct{}

func (f *GetVzenNode) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a VZEN node by ID or name.")
}

func (a *GetVzenNodeArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Id, "The ID of the VZEN node to look up.")
	ann.Describe(&a.Name, "The name of the VZEN node to look up.")
}

func (r *GetVzenNodeResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the VZEN node.")
	a.Describe(&r.Name, "The name of the VZEN node.")
	a.Describe(&r.Status, "The status of the VZEN node.")
	a.Describe(&r.Type, "The type of the VZEN node.")
	a.Describe(&r.IpSecEnabled, "Whether IPSec is enabled on the node.")
	a.Describe(&r.IpAddress, "The IP address of the VZEN node.")
	a.Describe(&r.SubnetMask, "The subnet mask of the VZEN node.")
	a.Describe(&r.DefaultGateway, "The default gateway of the VZEN node.")
	a.Describe(&r.ZGatewayId, "The ZGateway ID associated with the VZEN node.")
	a.Describe(&r.InProduction, "Whether the node is in production.")
	a.Describe(&r.OnDemandSupportTunnelEnabled, "Whether on-demand support tunnel is enabled.")
	a.Describe(&r.EstablishSupportTunnelEnabled, "Whether establish support tunnel is enabled.")
	a.Describe(&r.LoadBalancerIpAddress, "The load balancer IP address.")
	a.Describe(&r.DeploymentMode, "The deployment mode (STANDALONE or CLUSTER).")
	a.Describe(&r.ClusterName, "The cluster name if deployment mode is CLUSTER.")
	a.Describe(&r.VzenSkuType, "The VZEN SKU type (SMALL, MEDIUM, or LARGE).")
}

func (*GetVzenNode) Invoke(ctx context.Context, req infer.FunctionRequest[GetVzenNodeArgs]) (infer.FunctionResponse[GetVzenNodeResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetVzenNodeResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	var resp *vzen_nodes.VZENNodes
	if req.Input.Id != nil {
		r, err := vzen_nodes.Get(ctx, svc, *req.Input.Id)
		if err != nil {
			return infer.FunctionResponse[GetVzenNodeResult]{}, err
		}
		resp = r
	}
	if resp == nil && req.Input.Name != nil && *req.Input.Name != "" {
		r, err := vzen_nodes.GetNodeByName(ctx, svc, *req.Input.Name)
		if err != nil {
			return infer.FunctionResponse[GetVzenNodeResult]{}, err
		}
		resp = r
	}

	if resp == nil {
		return infer.FunctionResponse[GetVzenNodeResult]{}, fmt.Errorf("couldn't find any vzen node with id %v or name %v", req.Input.Id, ptrToString(req.Input.Name))
	}

	return infer.FunctionResponse[GetVzenNodeResult]{
		Output: GetVzenNodeResult{
			Id:                            resp.ID,
			Name:                          resp.Name,
			Status:                        resp.Status,
			Type:                          resp.Type,
			IpSecEnabled:                  resp.IPSecEnabled,
			IpAddress:                     resp.IPAddress,
			SubnetMask:                    resp.SubnetMask,
			DefaultGateway:                resp.DefaultGateway,
			ZGatewayId:                    resp.ZGatewayID,
			InProduction:                  resp.InProduction,
			OnDemandSupportTunnelEnabled:  resp.OnDemandSupportTunnelEnabled,
			EstablishSupportTunnelEnabled: resp.EstablishSupportTunnelEnabled,
			LoadBalancerIpAddress:         resp.LoadBalancerIPAddress,
			DeploymentMode:                resp.DeploymentMode,
			ClusterName:                   resp.ClusterName,
			VzenSkuType:                   resp.VzenSkuType,
		},
	}, nil
}
