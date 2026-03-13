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

// Package provider implements the VZEN Cluster resource and invoke.
// Adopted from terraform-provider-zia resource_zia_vzen_cluster.go and data_source_zia_vzen_cluster.go.

package provider

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"time"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/vzen_clusters"
)

// validVzenClusterTypes from terraform-provider-zia resource schema.
var validVzenClusterTypes = map[string]bool{
	"ANY": true, "NONE": true, "SME": true, "SMSM": true, "SMCA": true, "SMUI": true, "SMCDS": true,
	"SMDNSD": true, "SMAA": true, "SMTP": true, "SMQTN": true, "VIP": true, "UIZ": true, "UIAE": true,
	"SITEREVIEW": true, "PAC": true, "S_RELAY": true, "M_RELAY": true, "H_MON": true, "SMIKE": true,
	"NSS": true, "SMEZA": true, "SMLB": true, "SMFCCLT": true, "SMBA": true, "SMBAC": true,
	"SMESXI": true, "SMBAUI": true, "VZEN": true, "ZSCMCLT": true, "SMDLP": true, "ZSQUERY": true,
	"ADP": true, "SMCDSDLP": true, "SMSCIM": true, "ZSAPI": true, "ZSCMCDSSCLT": true,
	"LOCAL_MTS": true, "SVPN": true, "SMCASB": true, "SMFALCONUI": true, "MOBILEAPP_REG": true,
	"SMRESTSVR": true, "FALCONCA": true, "MOBILEAPP_NF": true, "ZIRSVR": true, "SMEDGEUI": true,
	"ALERTEVAL": true, "ALERTNOTIF": true, "SMPARTNERUI": true, "CQM": true, "DATAKEEPER": true,
	"SMBAM": true, "ZWACLT": true,
}

// --- Resource ---

// VzenCluster implements the zia:index:VzenCluster resource.
type VzenCluster struct{}

// VzenClusterArgs are the inputs for VzenCluster.
type VzenClusterArgs struct {
	Name            *string `pulumi:"name,optional"`
	Status          *string `pulumi:"status,optional"`
	IpSecEnabled    *bool   `pulumi:"ipSecEnabled,optional"`
	IpAddress       *string `pulumi:"ipAddress,optional"`
	SubnetMask      *string `pulumi:"subnetMask,optional"`
	DefaultGateway  *string `pulumi:"defaultGateway,optional"`
	Type            *string `pulumi:"type,optional"`
	VirtualZenNodes []int   `pulumi:"virtualZenNodes,optional"`
}

// VzenClusterState is the persisted state.
type VzenClusterState struct {
	VzenClusterArgs
	ClusterId *int `pulumi:"clusterId"`
}

func (VzenCluster) Check(ctx context.Context, req infer.CheckRequest) (infer.CheckResponse[VzenClusterArgs], error) {
	inputs, failures, err := infer.DefaultCheck[VzenClusterArgs](ctx, req.NewInputs)
	if err != nil {
		return infer.CheckResponse[VzenClusterArgs]{}, err
	}
	if len(failures) > 0 {
		return infer.CheckResponse[VzenClusterArgs]{Failures: failures}, nil
	}
	if inputs.Type != nil && *inputs.Type != "" && !validVzenClusterTypes[*inputs.Type] {
		return infer.CheckResponse[VzenClusterArgs]{Failures: []p.CheckFailure{{
			Property: "type",
			Reason:   "type must be one of the valid VZEN cluster types (e.g., ANY, NONE, VZEN, etc.)",
		}}}, nil
	}
	return infer.CheckResponse[VzenClusterArgs]{Inputs: inputs}, nil
}

func (VzenCluster) Create(ctx context.Context, req infer.CreateRequest[VzenClusterArgs]) (infer.CreateResponse[VzenClusterState], error) {
	if req.DryRun {
		return infer.CreateResponse[VzenClusterState]{
			ID: "preview",
			Output: VzenClusterState{
				VzenClusterArgs: req.Inputs,
				ClusterId:       intPtr(0),
			},
		}, nil
	}

	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[VzenClusterState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := vzenClusterArgsToAPI(req.Inputs, 0)
	resp, _, err := vzen_clusters.Create(ctx, svc, &apiReq)
	if err != nil {
		return infer.CreateResponse[VzenClusterState]{}, fmt.Errorf("creating vzen cluster: %w", err)
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.CreateResponse[VzenClusterState]{}, activationErr
		}
	}

	rule, err := vzen_clusters.Get(ctx, svc, resp.ID)
	if err != nil {
		return infer.CreateResponse[VzenClusterState]{
			ID:     strconv.Itoa(resp.ID),
			Output: VzenClusterState{VzenClusterArgs: req.Inputs, ClusterId: &resp.ID},
		}, nil
	}
	return infer.CreateResponse[VzenClusterState]{
		ID:     strconv.Itoa(resp.ID),
		Output: vzenClusterAPIToState(rule),
	}, nil
}

func (VzenCluster) Read(ctx context.Context, req infer.ReadRequest[VzenClusterArgs, VzenClusterState]) (infer.ReadResponse[VzenClusterArgs, VzenClusterState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[VzenClusterArgs, VzenClusterState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		rule, lookupErr := vzen_clusters.GetClusterByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[VzenClusterArgs, VzenClusterState]{}, nil
			}
			return infer.ReadResponse[VzenClusterArgs, VzenClusterState]{}, lookupErr
		}
		id = rule.ID
	}

	rule, err := vzen_clusters.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[VzenClusterArgs, VzenClusterState]{}, nil
		}
		return infer.ReadResponse[VzenClusterArgs, VzenClusterState]{}, err
	}

	state := vzenClusterAPIToState(rule)
	args := vzenClusterStateToArgs(rule)
	return infer.ReadResponse[VzenClusterArgs, VzenClusterState]{
		ID:     strconv.Itoa(rule.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (VzenCluster) Update(ctx context.Context, req infer.UpdateRequest[VzenClusterArgs, VzenClusterState]) (infer.UpdateResponse[VzenClusterState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[VzenClusterState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[VzenClusterState]{}, fmt.Errorf("invalid cluster id: %w", err)
	}

	apiReq := vzenClusterArgsToAPI(req.Inputs, id)
	if _, _, err := vzen_clusters.Update(ctx, svc, id, &apiReq); err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.UpdateResponse[VzenClusterState]{}, nil
		}
		return infer.UpdateResponse[VzenClusterState]{}, fmt.Errorf("updating vzen cluster: %w", err)
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[VzenClusterState]{}, activationErr
		}
	}

	rule, err := vzen_clusters.Get(ctx, svc, id)
	if err != nil {
		return infer.UpdateResponse[VzenClusterState]{
			Output: VzenClusterState{VzenClusterArgs: req.Inputs, ClusterId: &id},
		}, nil
	}
	return infer.UpdateResponse[VzenClusterState]{Output: vzenClusterAPIToState(rule)}, nil
}

func (VzenCluster) Delete(ctx context.Context, req infer.DeleteRequest[VzenClusterState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid cluster id: %w", err)
	}

	if _, err := vzen_clusters.Delete(ctx, svc, id); err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.DeleteResponse{}, nil
		}
		return infer.DeleteResponse{}, fmt.Errorf("deleting vzen cluster: %w", err)
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}
	return infer.DeleteResponse{}, nil
}

func (VzenCluster) Annotate(a infer.Annotator) {
	describeResource(a, &VzenCluster{}, `The zia.VzenCluster resource manages Virtual ZEN (VZEN) cluster configurations in the
Zscaler Internet Access (ZIA) cloud. VZEN clusters group multiple VZEN nodes for high availability
and load balancing of traffic processing.

{{% examples %}}
## Example Usage

{{% example %}}
### Basic VZEN Cluster

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.VzenCluster("example", {
    name: "Example VZEN Cluster",
    status: "ENABLED",
    type: "VZEN",
    ipAddress: "10.0.0.20",
    subnetMask: "255.255.255.0",
    defaultGateway: "10.0.0.1",
    ipSecEnabled: true,
    virtualZenNodes: [12345, 67890],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.VzenCluster("example",
    name="Example VZEN Cluster",
    status="ENABLED",
    type="VZEN",
    ip_address="10.0.0.20",
    subnet_mask="255.255.255.0",
    default_gateway="10.0.0.1",
    ip_sec_enabled=True,
    virtual_zen_nodes=[12345, 67890],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:VzenCluster
    properties:
      name: Example VZEN Cluster
      status: ENABLED
      type: VZEN
      ipAddress: "10.0.0.20"
      subnetMask: "255.255.255.0"
      defaultGateway: "10.0.0.1"
      ipSecEnabled: true
      virtualZenNodes:
        - 12345
        - 67890
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing VZEN cluster can be imported using its ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:VzenCluster example 12345
`+tripleBacktick("")+`
`)
}

func (a *VzenClusterArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "Name of the VZEN cluster.")
	ann.Describe(&a.Status, "The status of the cluster (e.g., 'ENABLED', 'DISABLED').")
	ann.Describe(&a.IpSecEnabled, "Whether IPSec is enabled on the cluster.")
	ann.Describe(&a.IpAddress, "The IP address of the VZEN cluster.")
	ann.Describe(&a.SubnetMask, "The subnet mask of the VZEN cluster.")
	ann.Describe(&a.DefaultGateway, "The default gateway of the VZEN cluster.")
	ann.Describe(&a.Type, "The type of the VZEN cluster.")
	ann.Describe(&a.VirtualZenNodes, "List of VZEN node IDs that belong to this cluster.")
}

func (s *VzenClusterState) Annotate(a infer.Annotator) {
	a.Describe(&s.ClusterId, "The unique identifier for the VZEN cluster assigned by the ZIA cloud.")
}

func vzenClusterArgsToAPI(in VzenClusterArgs, existingID int) vzen_clusters.VZENClusters {
	return vzen_clusters.VZENClusters{
		ID:              existingID,
		Name:            ptrToString(in.Name),
		Status:          ptrToString(in.Status),
		Type:            ptrToString(in.Type),
		IpAddress:       ptrToString(in.IpAddress),
		SubnetMask:      ptrToString(in.SubnetMask),
		DefaultGateway:  ptrToString(in.DefaultGateway),
		IpSecEnabled:    ptrToBool(in.IpSecEnabled),
		VirtualZenNodes: idsToIDNameExternalIDs(in.VirtualZenNodes),
	}
}

func vzenClusterAPIToState(rule *vzen_clusters.VZENClusters) VzenClusterState {
	ids := idNameExternalIDsToIDs(rule.VirtualZenNodes)
	sort.Ints(ids)
	return VzenClusterState{
		VzenClusterArgs: VzenClusterArgs{
			Name:            stringPtr(rule.Name),
			Status:          stringPtr(rule.Status),
			Type:            stringPtr(rule.Type),
			IpAddress:       stringPtr(rule.IpAddress),
			SubnetMask:      stringPtr(rule.SubnetMask),
			DefaultGateway:  stringPtr(rule.DefaultGateway),
			IpSecEnabled:    boolPtr(rule.IpSecEnabled),
			VirtualZenNodes: ids,
		},
		ClusterId: &rule.ID,
	}
}

func vzenClusterStateToArgs(rule *vzen_clusters.VZENClusters) VzenClusterArgs {
	return vzenClusterAPIToState(rule).VzenClusterArgs
}

// --- Invoke (data source) ---

// VirtualZenNodeOutput is a node in the cluster.
type VirtualZenNodeOutput struct {
	Id   int    `pulumi:"nodeId"` // Pulumi reserves "id" in function outputs
	Name string `pulumi:"name"`
}

// GetVzenClusterArgs are the inputs for the GetVzenCluster invoke.
type GetVzenClusterArgs struct {
	Id   *int    `pulumi:"clusterId,optional"` // Pulumi reserves "id" in function I/O
	Name *string `pulumi:"name,optional"`
}

// GetVzenClusterResult is the output of the GetVzenCluster invoke.
type GetVzenClusterResult struct {
	Id              int                    `pulumi:"clusterId"` // Pulumi reserves "id" in function outputs
	Name            string                 `pulumi:"name"`
	Status          string                 `pulumi:"status"`
	Type            string                 `pulumi:"type"`
	IpAddress       string                 `pulumi:"ipAddress"`
	SubnetMask      string                 `pulumi:"subnetMask"`
	DefaultGateway  string                 `pulumi:"defaultGateway"`
	IpSecEnabled    bool                   `pulumi:"ipSecEnabled"`
	VirtualZenNodes []VirtualZenNodeOutput `pulumi:"virtualZenNodes"`
}

// GetVzenCluster implements the zia:index:GetVzenCluster invoke.
type GetVzenCluster struct{}

func (f *GetVzenCluster) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a VZEN cluster by ID or name.")
}

func (a *GetVzenClusterArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Id, "The ID of the VZEN cluster to look up.")
	ann.Describe(&a.Name, "The name of the VZEN cluster to look up.")
}

func (r *GetVzenClusterResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the VZEN cluster.")
	a.Describe(&r.Name, "The name of the VZEN cluster.")
	a.Describe(&r.Status, "The status of the VZEN cluster.")
	a.Describe(&r.Type, "The type of the VZEN cluster.")
	a.Describe(&r.IpAddress, "The IP address of the VZEN cluster.")
	a.Describe(&r.SubnetMask, "The subnet mask of the VZEN cluster.")
	a.Describe(&r.DefaultGateway, "The default gateway of the VZEN cluster.")
	a.Describe(&r.IpSecEnabled, "Whether IPSec is enabled on the cluster.")
	a.Describe(&r.VirtualZenNodes, "The list of virtual ZEN nodes in this cluster.")
}

func (*GetVzenCluster) Invoke(ctx context.Context, req infer.FunctionRequest[GetVzenClusterArgs]) (infer.FunctionResponse[GetVzenClusterResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetVzenClusterResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	var resp *vzen_clusters.VZENClusters
	if req.Input.Id != nil {
		r, err := vzen_clusters.Get(ctx, svc, *req.Input.Id)
		if err != nil {
			return infer.FunctionResponse[GetVzenClusterResult]{}, err
		}
		resp = r
	}
	if resp == nil && req.Input.Name != nil && *req.Input.Name != "" {
		r, err := vzen_clusters.GetClusterByName(ctx, svc, *req.Input.Name)
		if err != nil {
			return infer.FunctionResponse[GetVzenClusterResult]{}, err
		}
		resp = r
	}

	if resp == nil {
		return infer.FunctionResponse[GetVzenClusterResult]{}, fmt.Errorf("couldn't find any vzen cluster with id %v or name %v", req.Input.Id, ptrToString(req.Input.Name))
	}

	nodes := make([]VirtualZenNodeOutput, len(resp.VirtualZenNodes))
	for i, n := range resp.VirtualZenNodes {
		nodes[i] = VirtualZenNodeOutput{Id: n.ID, Name: n.Name}
	}

	return infer.FunctionResponse[GetVzenClusterResult]{
		Output: GetVzenClusterResult{
			Id:              resp.ID,
			Name:            resp.Name,
			Status:          resp.Status,
			Type:            resp.Type,
			IpAddress:       resp.IpAddress,
			SubnetMask:      resp.SubnetMask,
			DefaultGateway:  resp.DefaultGateway,
			IpSecEnabled:    resp.IpSecEnabled,
			VirtualZenNodes: nodes,
		},
	}, nil
}
