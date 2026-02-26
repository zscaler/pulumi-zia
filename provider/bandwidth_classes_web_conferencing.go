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

// Package provider implements the Bandwidth Class Web Conferencing resource.
// Adopted from terraform-provider-zia resource_zia_bandwidth_classes_web_conferencing.go.
// Updates an existing bandwidth class by name (BANDWIDTH_CAT_WEBCONF or BANDWIDTH_CAT_VOIP). Delete is a no-op.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/bandwidth_control/bandwidth_classes"
)

// BandwidthClassWebConferencing implements the zia:index:BandwidthClassWebConferencing resource.
type BandwidthClassWebConferencing struct{}

// BandwidthClassWebConferencingArgs are the inputs.
type BandwidthClassWebConferencingArgs struct {
	Name         *string  `pulumi:"name,optional"`
	Type         *string  `pulumi:"type,optional"`
	Applications []string `pulumi:"applications,optional"`
}

// BandwidthClassWebConferencingState is the persisted state.
type BandwidthClassWebConferencingState struct {
	BandwidthClassWebConferencingArgs
	ClassId *int `pulumi:"classId"`
}

func bandwidthClassWebConferencingToAPI(args BandwidthClassWebConferencingArgs, id int) bandwidth_classes.BandwidthClasses {
	name := ptrToString(args.Name)
	typ := ptrToString(args.Type)
	return bandwidth_classes.BandwidthClasses{
		ID:           id,
		Name:         name,
		Type:         typ,
		Applications: args.Applications,
	}
}

func (BandwidthClassWebConferencing) Create(ctx context.Context, req infer.CreateRequest[BandwidthClassWebConferencingArgs]) (infer.CreateResponse[BandwidthClassWebConferencingState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[BandwidthClassWebConferencingState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	name := ptrToString(req.Inputs.Name)
	if name == "" {
		return infer.CreateResponse[BandwidthClassWebConferencingState]{}, fmt.Errorf("name must be provided to locate the existing bandwidth class")
	}

	existing, err := bandwidth_classes.GetByName(ctx, service, name)
	if err != nil {
		return infer.CreateResponse[BandwidthClassWebConferencingState]{}, fmt.Errorf("failed to find existing bandwidth class by name %q: %w", name, err)
	}
	log.Printf("[INFO] Found existing bandwidth class %q with ID %d", existing.Name, existing.ID)

	apiReq := bandwidthClassWebConferencingToAPI(req.Inputs, existing.ID)
	if _, _, err := bandwidth_classes.Update(ctx, service, existing.ID, &apiReq); err != nil {
		return infer.CreateResponse[BandwidthClassWebConferencingState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[BandwidthClassWebConferencingState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := BandwidthClassWebConferencingState{
		BandwidthClassWebConferencingArgs: req.Inputs,
		ClassId:                           &existing.ID,
	}
	return infer.CreateResponse[BandwidthClassWebConferencingState]{
		ID:     strconv.Itoa(existing.ID),
		Output: state,
	}, nil
}

func (BandwidthClassWebConferencing) Read(ctx context.Context, req infer.ReadRequest[BandwidthClassWebConferencingArgs, BandwidthClassWebConferencingState]) (infer.ReadResponse[BandwidthClassWebConferencingArgs, BandwidthClassWebConferencingState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[BandwidthClassWebConferencingArgs, BandwidthClassWebConferencingState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	classID := 0
	if req.State.ClassId != nil {
		classID = *req.State.ClassId
	}
	if classID == 0 {
		return infer.ReadResponse[BandwidthClassWebConferencingArgs, BandwidthClassWebConferencingState]{}, fmt.Errorf("no bandwidth class id in state")
	}

	resp, err := bandwidth_classes.Get(ctx, service, classID)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[BandwidthClassWebConferencingArgs, BandwidthClassWebConferencingState]{ID: ""}, nil
		}
		return infer.ReadResponse[BandwidthClassWebConferencingArgs, BandwidthClassWebConferencingState]{}, err
	}

	args := BandwidthClassWebConferencingArgs{
		Name:         stringPtr(resp.Name),
		Type:         stringPtr(resp.Type),
		Applications: resp.Applications,
	}
	state := BandwidthClassWebConferencingState{
		BandwidthClassWebConferencingArgs: args,
		ClassId:                           &resp.ID,
	}
	return infer.ReadResponse[BandwidthClassWebConferencingArgs, BandwidthClassWebConferencingState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (BandwidthClassWebConferencing) Update(ctx context.Context, req infer.UpdateRequest[BandwidthClassWebConferencingArgs, BandwidthClassWebConferencingState]) (infer.UpdateResponse[BandwidthClassWebConferencingState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[BandwidthClassWebConferencingState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	classID := 0
	if req.State.ClassId != nil {
		classID = *req.State.ClassId
	}
	if classID == 0 {
		name := ptrToString(req.Inputs.Name)
		if name == "" {
			return infer.UpdateResponse[BandwidthClassWebConferencingState]{}, fmt.Errorf("either classId or name must be set for update")
		}
		existing, err := bandwidth_classes.GetByName(ctx, service, name)
		if err != nil {
			return infer.UpdateResponse[BandwidthClassWebConferencingState]{}, fmt.Errorf("failed to find bandwidth class with name %q: %w", name, err)
		}
		classID = existing.ID
	}

	if _, err := bandwidth_classes.Get(ctx, service, classID); err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[BandwidthClassWebConferencingState]{}, nil
		}
		return infer.UpdateResponse[BandwidthClassWebConferencingState]{}, err
	}

	apiReq := bandwidthClassWebConferencingToAPI(req.Inputs, classID)
	if _, _, err := bandwidth_classes.Update(ctx, service, classID, &apiReq); err != nil {
		return infer.UpdateResponse[BandwidthClassWebConferencingState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[BandwidthClassWebConferencingState]{}, activationErr
		}
	}

	state := BandwidthClassWebConferencingState{
		BandwidthClassWebConferencingArgs: req.Inputs,
		ClassId:                           &classID,
	}
	return infer.UpdateResponse[BandwidthClassWebConferencingState]{Output: state}, nil
}

func (BandwidthClassWebConferencing) Delete(ctx context.Context, req infer.DeleteRequest[BandwidthClassWebConferencingState]) (infer.DeleteResponse, error) {
	// No-op: this resource updates an existing class; deleting the Pulumi resource does not remove the underlying class
	return infer.DeleteResponse{}, nil
}

func (BandwidthClassWebConferencing) Annotate(a infer.Annotator) {
	describeResource(a, &BandwidthClassWebConferencing{}, `The zia_bandwidth_classes_web_conferencing resource manages bandwidth class settings for web conferencing and VoIP in the Zscaler Internet Access (ZIA) cloud service. This resource updates an existing built-in bandwidth class by name (e.g. BANDWIDTH_CAT_WEBCONF or BANDWIDTH_CAT_VOIP). Deleting the Pulumi resource does not remove the underlying class.

{{% examples %}}
## Example Usage

{{% example %}}
### Bandwidth Class Web Conferencing

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.BandwidthClassWebConferencing("example", {
    name: "BANDWIDTH_CAT_WEBCONF",
    type: "WEB_CONF",
    applications: ["ZOOM", "WEBEX"],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.BandwidthClassWebConferencing("example",
    name="BANDWIDTH_CAT_WEBCONF",
    type="WEB_CONF",
    applications=["ZOOM", "WEBEX"],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:BandwidthClassWebConferencing
    properties:
      name: BANDWIDTH_CAT_WEBCONF
      type: WEB_CONF
      applications:
        - ZOOM
        - WEBEX
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing Bandwidth Class Web Conferencing can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:BandwidthClassWebConferencing example 12345
`+tripleBacktick("")+`
`)
}

func (a *BandwidthClassWebConferencingArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the bandwidth class (e.g. `BANDWIDTH_CAT_WEBCONF` or `BANDWIDTH_CAT_VOIP`).")
	ann.Describe(&a.Type, "The type of the bandwidth class (e.g. `WEB_CONF`, `VOIP`).")
	ann.Describe(&a.Applications, "List of applications associated with this bandwidth class (e.g. `ZOOM`, `WEBEX`, `TEAMS`).")
}

func (s *BandwidthClassWebConferencingState) Annotate(a infer.Annotator) {
	a.Describe(&s.ClassId, "The system-generated ID of the bandwidth class.")
}

var _ infer.CustomResource[BandwidthClassWebConferencingArgs, BandwidthClassWebConferencingState] = BandwidthClassWebConferencing{}
