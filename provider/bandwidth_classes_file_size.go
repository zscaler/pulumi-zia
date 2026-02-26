// Copyright (c) 2023 Zscaler Technology Alliances, <zscaler-partner-labs@z-bd.com>
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

// Package provider implements the Bandwidth Class File Size resource.
// Adopted from terraform-provider-zia resource_zia_bandwidth_classes_file_size.go.
// Updates an existing bandwidth class by name (e.g., LARGE_FILE). Delete is a no-op.

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

// BandwidthClassFileSize implements the zia:index:BandwidthClassFileSize resource.
type BandwidthClassFileSize struct{}

// BandwidthClassFileSizeArgs are the inputs.
type BandwidthClassFileSizeArgs struct {
	Name     *string `pulumi:"name,optional"`
	Type     *string `pulumi:"type,optional"`
	FileSize *string `pulumi:"fileSize,optional"`
}

// BandwidthClassFileSizeState is the persisted state.
type BandwidthClassFileSizeState struct {
	BandwidthClassFileSizeArgs
	ClassId *int `pulumi:"classId"`
}

// bandwidthClassFileSizeToAPI builds the API request, merging in existing Urls, UrlCategories,
// WebApplications, Applications so the update does not clear them (API requires at least one).
func bandwidthClassFileSizeToAPI(args BandwidthClassFileSizeArgs, id int, existing *bandwidth_classes.BandwidthClasses) bandwidth_classes.BandwidthClasses {
	name := ptrToStringDefault(args.Name, "BANDWIDTH_CAT_LARGE_FILE")
	typ := ptrToStringDefault(args.Type, "BANDWIDTH_CAT_LARGE_FILE")
	req := bandwidth_classes.BandwidthClasses{
		ID:       id,
		Name:     name,
		Type:     typ,
		FileSize: ptrToString(args.FileSize),
	}
	if existing != nil {
		req.Urls = existing.Urls
		req.UrlCategories = existing.UrlCategories
		req.WebApplications = existing.WebApplications
		req.Applications = existing.Applications
	}
	return req
}

func (BandwidthClassFileSize) Create(ctx context.Context, req infer.CreateRequest[BandwidthClassFileSizeArgs]) (infer.CreateResponse[BandwidthClassFileSizeState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[BandwidthClassFileSizeState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	name := ptrToStringDefault(req.Inputs.Name, "BANDWIDTH_CAT_LARGE_FILE")
	if name == "" {
		return infer.CreateResponse[BandwidthClassFileSizeState]{}, fmt.Errorf("name must be provided to locate the existing bandwidth class")
	}

	existing, err := bandwidth_classes.GetByName(ctx, service, name)
	if err != nil {
		return infer.CreateResponse[BandwidthClassFileSizeState]{}, fmt.Errorf("failed to find existing bandwidth class by name %q: %w", name, err)
	}
	log.Printf("[INFO] Found existing bandwidth class %q with ID %d", existing.Name, existing.ID)

	apiReq := bandwidthClassFileSizeToAPI(req.Inputs, existing.ID, existing)
	updated, _, err := bandwidth_classes.Update(ctx, service, existing.ID, &apiReq)
	if err != nil {
		return infer.CreateResponse[BandwidthClassFileSizeState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[BandwidthClassFileSizeState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	// Populate state with actual values (from API when input was nil) so outputs are never nil
	state := BandwidthClassFileSizeState{
		BandwidthClassFileSizeArgs: BandwidthClassFileSizeArgs{
			Name:     ptrOrDefault(req.Inputs.Name, updated.Name),
			Type:     ptrOrDefault(req.Inputs.Type, updated.Type),
			FileSize: ptrOrDefault(req.Inputs.FileSize, updated.FileSize),
		},
		ClassId: &existing.ID,
	}
	return infer.CreateResponse[BandwidthClassFileSizeState]{
		ID:     strconv.Itoa(existing.ID),
		Output: state,
	}, nil
}

func (BandwidthClassFileSize) Read(ctx context.Context, req infer.ReadRequest[BandwidthClassFileSizeArgs, BandwidthClassFileSizeState]) (infer.ReadResponse[BandwidthClassFileSizeArgs, BandwidthClassFileSizeState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[BandwidthClassFileSizeArgs, BandwidthClassFileSizeState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	classID := 0
	if req.State.ClassId != nil {
		classID = *req.State.ClassId
	}
	if classID == 0 {
		return infer.ReadResponse[BandwidthClassFileSizeArgs, BandwidthClassFileSizeState]{}, fmt.Errorf("no bandwidth class file size id in state")
	}

	resp, err := bandwidth_classes.Get(ctx, service, classID)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[BandwidthClassFileSizeArgs, BandwidthClassFileSizeState]{ID: ""}, nil
		}
		return infer.ReadResponse[BandwidthClassFileSizeArgs, BandwidthClassFileSizeState]{}, err
	}

	args := BandwidthClassFileSizeArgs{
		Name:     stringPtr(resp.Name),
		Type:     stringPtr(resp.Type),
		FileSize: stringPtr(resp.FileSize),
	}
	state := BandwidthClassFileSizeState{
		BandwidthClassFileSizeArgs: args,
		ClassId:                    &resp.ID,
	}
	return infer.ReadResponse[BandwidthClassFileSizeArgs, BandwidthClassFileSizeState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (BandwidthClassFileSize) Update(ctx context.Context, req infer.UpdateRequest[BandwidthClassFileSizeArgs, BandwidthClassFileSizeState]) (infer.UpdateResponse[BandwidthClassFileSizeState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[BandwidthClassFileSizeState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	classID := 0
	if req.State.ClassId != nil {
		classID = *req.State.ClassId
	}
	if classID == 0 {
		name := ptrToString(req.Inputs.Name)
		if name == "" {
			return infer.UpdateResponse[BandwidthClassFileSizeState]{}, fmt.Errorf("either classId or name must be set for update")
		}
		existing, err := bandwidth_classes.GetByName(ctx, service, name)
		if err != nil {
			return infer.UpdateResponse[BandwidthClassFileSizeState]{}, fmt.Errorf("failed to find bandwidth class with name %q: %w", name, err)
		}
		classID = existing.ID
	}

	existing, err := bandwidth_classes.Get(ctx, service, classID)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[BandwidthClassFileSizeState]{}, nil
		}
		return infer.UpdateResponse[BandwidthClassFileSizeState]{}, err
	}

	apiReq := bandwidthClassFileSizeToAPI(req.Inputs, classID, existing)
	updated, _, err := bandwidth_classes.Update(ctx, service, classID, &apiReq)
	if err != nil {
		return infer.UpdateResponse[BandwidthClassFileSizeState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[BandwidthClassFileSizeState]{}, activationErr
		}
	}

	// Populate state with actual values so outputs are never nil
	state := BandwidthClassFileSizeState{
		BandwidthClassFileSizeArgs: BandwidthClassFileSizeArgs{
			Name:     ptrOrDefault(req.Inputs.Name, updated.Name),
			Type:     ptrOrDefault(req.Inputs.Type, updated.Type),
			FileSize: ptrOrDefault(req.Inputs.FileSize, updated.FileSize),
		},
		ClassId: &classID,
	}
	return infer.UpdateResponse[BandwidthClassFileSizeState]{Output: state}, nil
}

func (BandwidthClassFileSize) Delete(ctx context.Context, req infer.DeleteRequest[BandwidthClassFileSizeState]) (infer.DeleteResponse, error) {
	// No-op: this resource updates an existing class; deleting the Pulumi resource does not remove the underlying class
	return infer.DeleteResponse{}, nil
}

func (BandwidthClassFileSize) Annotate(a infer.Annotator) {
	describeResource(a, &BandwidthClassFileSize{}, `The zia_bandwidth_class_file_size resource manages the file size configuration for an existing bandwidth class in the Zscaler Internet Access (ZIA) cloud service. This resource updates a pre-existing bandwidth class (e.g., LARGE_FILE) with the specified file size threshold. Delete is a no-op; the underlying bandwidth class is not removed when the Pulumi resource is destroyed.

For more information, see the [ZIA Bandwidth Control documentation](https://help.zscaler.com/zia/bandwidth-control).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Bandwidth Class File Size

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.BandwidthClassFileSize("example", {
    name: "BANDWIDTH_CAT_LARGE_FILE",
    type: "BANDWIDTH_CAT_LARGE_FILE",
    fileSize: "100MB",
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.BandwidthClassFileSize("example",
    name="BANDWIDTH_CAT_LARGE_FILE",
    type="BANDWIDTH_CAT_LARGE_FILE",
    file_size="100MB",
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:BandwidthClassFileSize
    properties:
      name: BANDWIDTH_CAT_LARGE_FILE
      type: BANDWIDTH_CAT_LARGE_FILE
      fileSize: 100MB
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

This resource updates an existing bandwidth class by name and does not support traditional import. It is automatically managed by the provider.
`)
}

func (a *BandwidthClassFileSizeArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the existing bandwidth class to update (e.g., `BANDWIDTH_CAT_LARGE_FILE`).")
	ann.Describe(&a.Type, "The type of the bandwidth class (e.g., `BANDWIDTH_CAT_LARGE_FILE`).")
	ann.Describe(&a.FileSize, "The file size threshold for the bandwidth class (e.g., `100MB`).")
}

func (s *BandwidthClassFileSizeState) Annotate(a infer.Annotator) {
	a.Describe(&s.ClassId, "The system-generated ID of the bandwidth class.")
}

var _ infer.CustomResource[BandwidthClassFileSizeArgs, BandwidthClassFileSizeState] = BandwidthClassFileSize{}
