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

// Package provider implements the Custom File Type resource.
// Adopted from terraform-provider-zia resource_zia_custom_file_types.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/filetypecontrol/custom_file_types"
)

// CustomFileType implements the zia:index:CustomFileType resource.
type CustomFileType struct{}

// CustomFileTypeArgs are the inputs.
type CustomFileTypeArgs struct {
	Name        *string `pulumi:"name,optional"`
	Description *string `pulumi:"description,optional"`
	Extension   *string `pulumi:"extension,optional"`
	FileTypeId  *int    `pulumi:"fileTypeId,optional"`
}

// CustomFileTypeState is the persisted state.
type CustomFileTypeState struct {
	CustomFileTypeArgs
	FileId *int `pulumi:"fileId"`
}

func customFileTypeToAPI(args CustomFileTypeArgs, id int) custom_file_types.CustomFileTypes {
	return custom_file_types.CustomFileTypes{
		ID:          id,
		Name:        ptrToString(args.Name),
		Description: ptrToString(args.Description),
		Extension:   ptrToString(args.Extension),
		FileTypeID:  ptrToIntDefault(args.FileTypeId, 0),
	}
}

func (CustomFileType) Create(ctx context.Context, req infer.CreateRequest[CustomFileTypeArgs]) (infer.CreateResponse[CustomFileTypeState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[CustomFileTypeState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := customFileTypeToAPI(req.Inputs, 0)
	resp, err := custom_file_types.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[CustomFileTypeState]{}, err
	}
	log.Printf("[INFO] Created ZIA custom file type. ID: %v", resp.ID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[CustomFileTypeState]{}, activationErr
		}
	}

	state := CustomFileTypeState{
		CustomFileTypeArgs: req.Inputs,
		FileId:             &resp.ID,
	}
	return infer.CreateResponse[CustomFileTypeState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (CustomFileType) Read(ctx context.Context, req infer.ReadRequest[CustomFileTypeArgs, CustomFileTypeState]) (infer.ReadResponse[CustomFileTypeArgs, CustomFileTypeState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[CustomFileTypeArgs, CustomFileTypeState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.FileId != nil {
		id = *req.State.FileId
	}
	if id == 0 {
		return infer.ReadResponse[CustomFileTypeArgs, CustomFileTypeState]{}, fmt.Errorf("no custom file type id in state")
	}

	resp, err := custom_file_types.Get(ctx, service, id)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[CustomFileTypeArgs, CustomFileTypeState]{ID: ""}, nil
		}
		return infer.ReadResponse[CustomFileTypeArgs, CustomFileTypeState]{}, err
	}

	args := CustomFileTypeArgs{
		Name:        stringPtr(resp.Name),
		Description: stringPtr(resp.Description),
		Extension:   stringPtr(resp.Extension),
		FileTypeId:  intPtr(resp.FileTypeID),
	}
	state := CustomFileTypeState{
		CustomFileTypeArgs: args,
		FileId:             &resp.ID,
	}
	return infer.ReadResponse[CustomFileTypeArgs, CustomFileTypeState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (CustomFileType) Update(ctx context.Context, req infer.UpdateRequest[CustomFileTypeArgs, CustomFileTypeState]) (infer.UpdateResponse[CustomFileTypeState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[CustomFileTypeState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.FileId != nil {
		id = *req.State.FileId
	}
	if id == 0 {
		return infer.UpdateResponse[CustomFileTypeState]{}, fmt.Errorf("no custom file type id in state")
	}

	if _, err := custom_file_types.Get(ctx, service, id); err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[CustomFileTypeState]{}, nil
		}
		return infer.UpdateResponse[CustomFileTypeState]{}, err
	}

	apiReq := customFileTypeToAPI(req.Inputs, id)
	if _, err := custom_file_types.Update(ctx, service, id, &apiReq); err != nil {
		return infer.UpdateResponse[CustomFileTypeState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[CustomFileTypeState]{}, activationErr
		}
	}

	state := CustomFileTypeState{
		CustomFileTypeArgs: req.Inputs,
		FileId:             &id,
	}
	return infer.UpdateResponse[CustomFileTypeState]{Output: state}, nil
}

func (CustomFileType) Delete(ctx context.Context, req infer.DeleteRequest[CustomFileTypeState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.FileId != nil {
		id = *req.State.FileId
	}
	if id != 0 {
		if _, err := custom_file_types.Delete(ctx, service, id); err != nil {
			return infer.DeleteResponse{}, err
		}
		log.Printf("[INFO] ZIA custom file type deleted")
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

func (CustomFileType) Annotate(a infer.Annotator) {
	describeResource(a, &CustomFileType{}, `The zia_custom_file_type resource manages custom file type controls in the Zscaler Internet Access (ZIA) cloud service. Custom file types allow you to define file extensions and types for use in file type control policies.

{{% examples %}}
## Example Usage

{{% example %}}
### Custom File Type

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.CustomFileType("example", {
    name: "Custom Archive",
    description: "Custom archive file type",
    extension: "myarch",
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.CustomFileType("example",
    name="Custom Archive",
    description="Custom archive file type",
    extension="myarch",
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:CustomFileType
    properties:
      name: Custom Archive
      description: Custom archive file type
      extension: myarch
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing Custom File Type can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:CustomFileType example 12345
`+tripleBacktick("")+`
`)
}

func (a *CustomFileTypeArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the custom file type.")
	ann.Describe(&a.Description, "A description of the custom file type.")
	ann.Describe(&a.Extension, "The file extension for this custom file type (e.g. `myarch`).")
	ann.Describe(&a.FileTypeId, "The file type category ID to associate this custom file type with.")
}

func (s *CustomFileTypeState) Annotate(a infer.Annotator) {
	a.Describe(&s.FileId, "The system-generated ID of the custom file type.")
}

var _ infer.CustomResource[CustomFileTypeArgs, CustomFileTypeState] = CustomFileType{}
