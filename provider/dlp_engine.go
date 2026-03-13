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

// Package provider implements the DLP Engine resource.
// Adopted from terraform-provider-zia resource_zia_dlp_engines.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/dlp/dlp_engines"
)

// DlpEngine implements the zia:index:DlpEngine resource.
type DlpEngine struct{}

// DlpEngineArgs are the inputs.
type DlpEngineArgs struct {
	Name             *string `pulumi:"name,optional"`
	Description      *string `pulumi:"description,optional"`
	EngineExpression *string `pulumi:"engineExpression,optional"`
	CustomDlpEngine  *bool   `pulumi:"customDlpEngine,optional"`
}

// DlpEngineState is the persisted state.
type DlpEngineState struct {
	DlpEngineArgs
	EngineId *int `pulumi:"engineId"`
}

func dlpEngineToAPI(args DlpEngineArgs, id int) dlp_engines.DLPEngines {
	return dlp_engines.DLPEngines{
		ID:               id,
		Name:             ptrToString(args.Name),
		Description:      ptrToString(args.Description),
		EngineExpression: ptrToString(args.EngineExpression),
		CustomDlpEngine:  ptrToBool(args.CustomDlpEngine),
	}
}

func (DlpEngine) Create(ctx context.Context, req infer.CreateRequest[DlpEngineArgs]) (infer.CreateResponse[DlpEngineState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[DlpEngineState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := dlpEngineToAPI(req.Inputs, 0)
	resp, _, err := dlp_engines.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[DlpEngineState]{}, err
	}
	log.Printf("[INFO] Created ZIA DLP engine. ID: %v", resp.ID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[DlpEngineState]{}, activationErr
		}
	}

	state := DlpEngineState{
		DlpEngineArgs: req.Inputs,
		EngineId:      &resp.ID,
	}
	return infer.CreateResponse[DlpEngineState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (DlpEngine) Read(ctx context.Context, req infer.ReadRequest[DlpEngineArgs, DlpEngineState]) (infer.ReadResponse[DlpEngineArgs, DlpEngineState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[DlpEngineArgs, DlpEngineState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.EngineId != nil {
		id = *req.State.EngineId
	}
	if id == 0 {
		return infer.ReadResponse[DlpEngineArgs, DlpEngineState]{}, fmt.Errorf("no DLP engine id in state")
	}

	resp, err := dlp_engines.Get(ctx, service, id)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[DlpEngineArgs, DlpEngineState]{ID: ""}, nil
		}
		return infer.ReadResponse[DlpEngineArgs, DlpEngineState]{}, err
	}

	args := DlpEngineArgs{
		Name:             stringPtr(resp.Name),
		Description:      stringPtr(resp.Description),
		EngineExpression: stringPtr(resp.EngineExpression),
		CustomDlpEngine:  boolPtr(resp.CustomDlpEngine),
	}
	state := DlpEngineState{
		DlpEngineArgs: args,
		EngineId:      &resp.ID,
	}
	return infer.ReadResponse[DlpEngineArgs, DlpEngineState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (DlpEngine) Update(ctx context.Context, req infer.UpdateRequest[DlpEngineArgs, DlpEngineState]) (infer.UpdateResponse[DlpEngineState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[DlpEngineState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.EngineId != nil {
		id = *req.State.EngineId
	}
	if id == 0 {
		return infer.UpdateResponse[DlpEngineState]{}, fmt.Errorf("no DLP engine id in state")
	}

	if _, err := dlp_engines.Get(ctx, service, id); err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[DlpEngineState]{}, nil
		}
		return infer.UpdateResponse[DlpEngineState]{}, err
	}

	apiReq := dlpEngineToAPI(req.Inputs, id)
	if _, _, err := dlp_engines.Update(ctx, service, id, &apiReq); err != nil {
		return infer.UpdateResponse[DlpEngineState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[DlpEngineState]{}, activationErr
		}
	}

	state := DlpEngineState{
		DlpEngineArgs: req.Inputs,
		EngineId:      &id,
	}
	return infer.UpdateResponse[DlpEngineState]{Output: state}, nil
}

func (DlpEngine) Delete(ctx context.Context, req infer.DeleteRequest[DlpEngineState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.EngineId != nil {
		id = *req.State.EngineId
	}
	if id != 0 {
		if _, err := dlp_engines.Delete(ctx, service, id); err != nil {
			if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.DeleteResponse{}, nil
			}
			return infer.DeleteResponse{}, err
		}
		log.Printf("[INFO] ZIA DLP engine deleted")
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

func (DlpEngine) Annotate(a infer.Annotator) {
	describeResource(a, &DlpEngine{}, `The zia_dlp_engines resource manages DLP (Data Loss Prevention) engines in the Zscaler Internet Access (ZIA) cloud service. DLP engines combine multiple DLP dictionaries using logical expressions to create sophisticated data detection criteria for DLP policy rules.

For more information, see the [ZIA Data Loss Prevention documentation](https://help.zscaler.com/zia/data-loss-prevention).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic DLP Engine

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.DlpEngine("example", {
    name: "Example DLP Engine",
    description: "Custom DLP engine combining multiple dictionaries",
    engineExpression: "((D63.S > 1))",
    customDlpEngine: true,
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.DlpEngine("example",
    name="Example DLP Engine",
    description="Custom DLP engine combining multiple dictionaries",
    engine_expression="((D63.S > 1))",
    custom_dlp_engine=True,
)
`+tripleBacktick("")+`

`+tripleBacktick("go")+`
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zia "github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := zia.NewDlpEngine(ctx, "example", &zia.DlpEngineArgs{
			Name:             pulumi.StringRef("Example DLP Engine"),
			Description:      pulumi.StringRef("Custom DLP engine combining multiple dictionaries"),
			EngineExpression: pulumi.StringRef("((D63.S > 1))"),
			CustomDlpEngine:  pulumi.BoolRef(true),
		})
		return err
	})
}
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:DlpEngine
    properties:
      name: Example DLP Engine
      description: Custom DLP engine combining multiple dictionaries
      engineExpression: "((D63.S > 1))"
      customDlpEngine: true
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing DLP Engine can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:DlpEngine example 12345
`+tripleBacktick("")+`
`)
}

func (a *DlpEngineArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the DLP engine. Must be unique.")
	ann.Describe(&a.Description, "A description of the DLP engine.")
	ann.Describe(&a.EngineExpression, "The logical expression combining DLP dictionaries. Uses dictionary IDs and operators (e.g., `((D63.S > 1))`).")
	ann.Describe(&a.CustomDlpEngine, "If true, this is a custom DLP engine; false indicates a predefined engine.")
}

func (s *DlpEngineState) Annotate(a infer.Annotator) {
	a.Describe(&s.EngineId, "The system-generated ID of the DLP engine.")
}

var _ infer.CustomResource[DlpEngineArgs, DlpEngineState] = DlpEngine{}
