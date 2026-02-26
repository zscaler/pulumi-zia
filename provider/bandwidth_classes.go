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

// Package provider implements the Bandwidth Classes resource.
// Adopted from terraform-provider-zia resource_zia_bandwidth_classes.go.

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

// BandwidthClass implements the zia:index:BandwidthClass resource.
type BandwidthClass struct{}

// BandwidthClassArgs are the inputs.
type BandwidthClassArgs struct {
	Name            *string  `pulumi:"name,optional"`
	Urls            []string `pulumi:"urls,optional"`
	UrlCategories   []string `pulumi:"urlCategories,optional"`
	WebApplications []string `pulumi:"webApplications,optional"`
}

// BandwidthClassState is the persisted state.
type BandwidthClassState struct {
	BandwidthClassArgs
	ClassId *int `pulumi:"classId"`
}

func bandwidthClassToAPI(args BandwidthClassArgs, id int) bandwidth_classes.BandwidthClasses {
	return bandwidth_classes.BandwidthClasses{
		ID:              id,
		Name:            ptrToString(args.Name),
		Urls:            args.Urls,
		UrlCategories:   args.UrlCategories,
		WebApplications: args.WebApplications,
	}
}

func (BandwidthClass) Create(ctx context.Context, req infer.CreateRequest[BandwidthClassArgs]) (infer.CreateResponse[BandwidthClassState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[BandwidthClassState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	name := ptrToString(req.Inputs.Name)
	if name == "" {
		return infer.CreateResponse[BandwidthClassState]{}, fmt.Errorf("name must be provided for bandwidth class")
	}

	// Idempotent create: adopt existing class if one with this name already exists.
	// This handles (1) Preview-then-Up where Preview creates but state isn't persisted,
	// and (2) leftover resources from failed runs. Update existing to match desired config.
	existing, err := bandwidth_classes.GetByName(ctx, service, name)
	if err == nil {
		log.Printf("[INFO] Adopting existing ZIA bandwidth class %q with ID: %v", name, existing.ID)
		apiReq := bandwidthClassToAPI(req.Inputs, existing.ID)
		if _, _, updateErr := bandwidth_classes.Update(ctx, service, existing.ID, &apiReq); updateErr != nil {
			return infer.CreateResponse[BandwidthClassState]{}, updateErr
		}
		if shouldActivate() {
			time.Sleep(2 * time.Second)
			if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
				return infer.CreateResponse[BandwidthClassState]{}, activationErr
			}
		}
		sortedInputs := BandwidthClassArgs{
			Name:            req.Inputs.Name,
			Urls:            sortStringSlice(req.Inputs.Urls),
			UrlCategories:   sortStringSlice(req.Inputs.UrlCategories),
			WebApplications: sortStringSlice(req.Inputs.WebApplications),
		}
		state := BandwidthClassState{
			BandwidthClassArgs: sortedInputs,
			ClassId:            &existing.ID,
		}
		return infer.CreateResponse[BandwidthClassState]{
			ID:     strconv.Itoa(existing.ID),
			Output: state,
		}, nil
	}

	apiReq := bandwidthClassToAPI(req.Inputs, 0)
	resp, _, err := bandwidth_classes.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[BandwidthClassState]{}, err
	}
	log.Printf("[INFO] Created ZIA bandwidth class. ID: %v", resp.ID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[BandwidthClassState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	sortedInputs := BandwidthClassArgs{
		Name:            req.Inputs.Name,
		Urls:            sortStringSlice(req.Inputs.Urls),
		UrlCategories:   sortStringSlice(req.Inputs.UrlCategories),
		WebApplications: sortStringSlice(req.Inputs.WebApplications),
	}
	state := BandwidthClassState{
		BandwidthClassArgs: sortedInputs,
		ClassId:            &resp.ID,
	}
	return infer.CreateResponse[BandwidthClassState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (BandwidthClass) Read(ctx context.Context, req infer.ReadRequest[BandwidthClassArgs, BandwidthClassState]) (infer.ReadResponse[BandwidthClassArgs, BandwidthClassState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[BandwidthClassArgs, BandwidthClassState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	classID := 0
	if req.State.ClassId != nil {
		classID = *req.State.ClassId
	}
	if classID == 0 {
		return infer.ReadResponse[BandwidthClassArgs, BandwidthClassState]{}, fmt.Errorf("no bandwidth class ID in state")
	}

	resp, err := bandwidth_classes.Get(ctx, service, classID)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[BandwidthClassArgs, BandwidthClassState]{ID: ""}, nil
		}
		return infer.ReadResponse[BandwidthClassArgs, BandwidthClassState]{}, err
	}

	args := BandwidthClassArgs{
		Name:            stringPtr(resp.Name),
		Urls:            sortStringSlice(resp.Urls),
		UrlCategories:   sortStringSlice(resp.UrlCategories),
		WebApplications: sortStringSlice(resp.WebApplications),
	}
	state := BandwidthClassState{
		BandwidthClassArgs: args,
		ClassId:            &resp.ID,
	}
	return infer.ReadResponse[BandwidthClassArgs, BandwidthClassState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (BandwidthClass) Update(ctx context.Context, req infer.UpdateRequest[BandwidthClassArgs, BandwidthClassState]) (infer.UpdateResponse[BandwidthClassState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[BandwidthClassState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	classID := 0
	if req.State.ClassId != nil {
		classID = *req.State.ClassId
	}
	if classID == 0 {
		return infer.UpdateResponse[BandwidthClassState]{}, fmt.Errorf("no bandwidth class ID in state")
	}

	if _, err := bandwidth_classes.Get(ctx, service, classID); err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[BandwidthClassState]{}, nil
		}
		return infer.UpdateResponse[BandwidthClassState]{}, err
	}

	apiReq := bandwidthClassToAPI(req.Inputs, classID)
	if _, _, err := bandwidth_classes.Update(ctx, service, classID, &apiReq); err != nil {
		return infer.UpdateResponse[BandwidthClassState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[BandwidthClassState]{}, activationErr
		}
	}

	sortedInputs := BandwidthClassArgs{
		Name:            req.Inputs.Name,
		Urls:            sortStringSlice(req.Inputs.Urls),
		UrlCategories:   sortStringSlice(req.Inputs.UrlCategories),
		WebApplications: sortStringSlice(req.Inputs.WebApplications),
	}
	state := BandwidthClassState{
		BandwidthClassArgs: sortedInputs,
		ClassId:            &classID,
	}
	return infer.UpdateResponse[BandwidthClassState]{Output: state}, nil
}

func (BandwidthClass) Delete(ctx context.Context, req infer.DeleteRequest[BandwidthClassState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	classID := 0
	if req.State.ClassId != nil {
		classID = *req.State.ClassId
	}
	if classID != 0 {
		if _, err := bandwidth_classes.Delete(ctx, service, classID); err != nil {
			return infer.DeleteResponse{}, err
		}
		log.Printf("[INFO] ZIA bandwidth class deleted")
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

func (BandwidthClass) Annotate(a infer.Annotator) {
	describeResource(a, &BandwidthClass{}, `The zia_bandwidth_class resource manages bandwidth classes in the Zscaler Internet Access (ZIA) cloud service. Bandwidth classes define traffic categories based on URLs, URL categories, and web applications that can be referenced in bandwidth control rules to apply specific bandwidth limits.

For more information, see the [ZIA Bandwidth Control documentation](https://help.zscaler.com/zia/bandwidth-control).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Bandwidth Class

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.BandwidthClass("example", {
    name: "Example Bandwidth Class",
    webApplications: ["STREAMING_MEDIA"],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.BandwidthClass("example",
    name="Example Bandwidth Class",
    web_applications=["STREAMING_MEDIA"],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:BandwidthClass
    properties:
      name: Example Bandwidth Class
      webApplications:
        - STREAMING_MEDIA
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing Bandwidth Class can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:BandwidthClass example 12345
`+tripleBacktick("")+`
`)
}

func (a *BandwidthClassArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the bandwidth class. Must be unique.")
	ann.Describe(&a.Urls, "List of URLs associated with the bandwidth class.")
	ann.Describe(&a.UrlCategories, "List of URL categories associated with the bandwidth class.")
	ann.Describe(&a.WebApplications, "List of web applications associated with the bandwidth class.")
}

func (s *BandwidthClassState) Annotate(a infer.Annotator) {
	a.Describe(&s.ClassId, "The system-generated ID of the bandwidth class.")
}

var _ infer.CustomResource[BandwidthClassArgs, BandwidthClassState] = BandwidthClass{}
