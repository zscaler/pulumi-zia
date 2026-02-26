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

// Package provider implements the ZIA Activation resource.
// Adopted from terraform-provider-zia resource_zia_activation.go.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/activation"
)

const activationID = "activation"

// Activation implements the zia:index:Activation resource.
// Triggers organization policy activation. Delete is a no-op.
type Activation struct{}

// ActivationArgs are the inputs.
type ActivationArgs struct {
	Status string `pulumi:"status"` // Must be "ACTIVE"
}

// ActivationState is the persisted state.
type ActivationState struct {
	ActivationArgs
	ResourceId string `pulumi:"resourceId"`
}

func (Activation) Create(ctx context.Context, req infer.CreateRequest[ActivationArgs]) (infer.CreateResponse[ActivationState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[ActivationState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	if req.Inputs.Status != "ACTIVE" {
		return infer.CreateResponse[ActivationState]{}, fmt.Errorf("status must be ACTIVE")
	}

	apiReq := activation.Activation{Status: req.Inputs.Status}
	log.Printf("[INFO] Performing configuration activation")
	resp, err := activation.CreateActivation(ctx, service, apiReq)
	if err != nil {
		return infer.CreateResponse[ActivationState]{}, err
	}
	log.Printf("[INFO] Configuration activation successful. %v", resp.Status)

	state := ActivationState{
		ActivationArgs: req.Inputs,
		ResourceId:     activationID,
	}
	return infer.CreateResponse[ActivationState]{
		ID:     activationID,
		Output: state,
	}, nil
}

func (Activation) Read(ctx context.Context, req infer.ReadRequest[ActivationArgs, ActivationState]) (infer.ReadResponse[ActivationArgs, ActivationState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[ActivationArgs, ActivationState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	resp, err := activation.GetActivationStatus(ctx, service)
	if err != nil {
		return infer.ReadResponse[ActivationArgs, ActivationState]{}, err
	}

	state := ActivationState{
		ActivationArgs: ActivationArgs{Status: resp.Status},
		ResourceId:     activationID,
	}
	return infer.ReadResponse[ActivationArgs, ActivationState]{
		ID:     activationID,
		Inputs: state.ActivationArgs,
		State:  state,
	}, nil
}

func (Activation) Delete(ctx context.Context, req infer.DeleteRequest[ActivationState]) (infer.DeleteResponse, error) {
	return infer.DeleteResponse{}, nil
}

func (Activation) Annotate(a infer.Annotator) {
	describeResource(a, &Activation{}, `The zia_activation resource triggers the activation of ZIA configuration changes in the Zscaler Internet Access (ZIA) cloud service. After making configuration changes to ZIA resources, this resource can be used to activate and push those changes to the ZIA cloud. Delete is a no-op.

For more information, see the [ZIA Configuration Activation documentation](https://help.zscaler.com/zia/activating-configuration-changes).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Configuration Activation

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.Activation("example", {
    status: "ACTIVE",
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.Activation("example",
    status="ACTIVE",
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:Activation
    properties:
      status: ACTIVE
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

This is a singleton resource and does not support traditional import. It is automatically managed by the provider.
`)
}

func (a *ActivationArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Status, "The activation status. Must be `ACTIVE` to trigger configuration activation.")
}

func (s *ActivationState) Annotate(a infer.Annotator) {
	a.Describe(&s.ResourceId, "The internal resource identifier for the activation.")
}

var _ infer.CustomResource[ActivationArgs, ActivationState] = Activation{}
