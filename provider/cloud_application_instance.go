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

// Package provider implements the Cloud Application Instance resource.
// Adopted from terraform-provider-zia resource_zia_cloud_application_instance.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/cloud_app_instances"
)

// CloudApplicationInstance implements the zia:index:CloudApplicationInstance resource.
type CloudApplicationInstance struct{}

// InstanceIdentifierInput is a nested input for instance identifiers.
type InstanceIdentifierInput struct {
	InstanceId             *int    `pulumi:"instanceId,optional"`
	InstanceIdentifier     *string `pulumi:"instanceIdentifier,optional"`
	InstanceIdentifierName *string `pulumi:"instanceIdentifierName,optional"`
	IdentifierType         *string `pulumi:"identifierType,optional"`
}

// CloudApplicationInstanceArgs are the inputs.
type CloudApplicationInstanceArgs struct {
	Name                *string                   `pulumi:"name,optional"`
	InstanceType        *string                   `pulumi:"instanceType,optional"`
	InstanceIdentifiers []InstanceIdentifierInput `pulumi:"instanceIdentifiers,optional"`
}

// CloudApplicationInstanceState is the persisted state.
type CloudApplicationInstanceState struct {
	CloudApplicationInstanceArgs
	InstanceId *int `pulumi:"instanceId"`
}

func instanceIdentifiersToAPI(list []InstanceIdentifierInput) []cloud_app_instances.InstanceIdentifiers {
	if len(list) == 0 {
		return nil
	}
	result := make([]cloud_app_instances.InstanceIdentifiers, len(list))
	for i, in := range list {
		instanceID := 0
		if in.InstanceId != nil {
			instanceID = *in.InstanceId
		}
		result[i] = cloud_app_instances.InstanceIdentifiers{
			InstanceID:             instanceID,
			InstanceIdentifier:     ptrToString(in.InstanceIdentifier),
			InstanceIdentifierName: ptrToString(in.InstanceIdentifierName),
			IdentifierType:         ptrToString(in.IdentifierType),
		}
	}
	return result
}

func instanceIdentifiersFromAPI(list []cloud_app_instances.InstanceIdentifiers) []InstanceIdentifierInput {
	if len(list) == 0 {
		return nil
	}
	result := make([]InstanceIdentifierInput, len(list))
	for i, api := range list {
		result[i] = InstanceIdentifierInput{
			InstanceId:             intPtr(api.InstanceID),
			InstanceIdentifier:     stringPtr(api.InstanceIdentifier),
			InstanceIdentifierName: stringPtr(api.InstanceIdentifierName),
			IdentifierType:         stringPtr(api.IdentifierType),
		}
	}
	return result
}

func cloudApplicationInstanceToAPI(args CloudApplicationInstanceArgs, instanceID int) cloud_app_instances.CloudApplicationInstances {
	return cloud_app_instances.CloudApplicationInstances{
		InstanceID:          instanceID,
		InstanceName:        ptrToString(args.Name),
		InstanceType:        ptrToString(args.InstanceType),
		InstanceIdentifiers: instanceIdentifiersToAPI(args.InstanceIdentifiers),
	}
}

func (CloudApplicationInstance) Create(ctx context.Context, req infer.CreateRequest[CloudApplicationInstanceArgs]) (infer.CreateResponse[CloudApplicationInstanceState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[CloudApplicationInstanceState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := cloudApplicationInstanceToAPI(req.Inputs, 0)
	resp, _, err := cloud_app_instances.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[CloudApplicationInstanceState]{}, err
	}
	log.Printf("[INFO] Created ZIA cloud application instance. ID: %v", resp.InstanceID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[CloudApplicationInstanceState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := CloudApplicationInstanceState{
		CloudApplicationInstanceArgs: req.Inputs,
		InstanceId:                   &resp.InstanceID,
	}
	return infer.CreateResponse[CloudApplicationInstanceState]{
		ID:     strconv.Itoa(resp.InstanceID),
		Output: state,
	}, nil
}

func (CloudApplicationInstance) Read(ctx context.Context, req infer.ReadRequest[CloudApplicationInstanceArgs, CloudApplicationInstanceState]) (infer.ReadResponse[CloudApplicationInstanceArgs, CloudApplicationInstanceState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[CloudApplicationInstanceArgs, CloudApplicationInstanceState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	instanceID := 0
	if req.State.InstanceId != nil {
		instanceID = *req.State.InstanceId
	}
	if instanceID == 0 {
		return infer.ReadResponse[CloudApplicationInstanceArgs, CloudApplicationInstanceState]{}, fmt.Errorf("no cloud application instance id in state")
	}

	resp, err := cloud_app_instances.Get(ctx, service, instanceID)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[CloudApplicationInstanceArgs, CloudApplicationInstanceState]{ID: ""}, nil
		}
		return infer.ReadResponse[CloudApplicationInstanceArgs, CloudApplicationInstanceState]{}, err
	}

	args := CloudApplicationInstanceArgs{
		Name:                stringPtr(resp.InstanceName),
		InstanceType:        stringPtr(resp.InstanceType),
		InstanceIdentifiers: instanceIdentifiersFromAPI(resp.InstanceIdentifiers),
	}
	state := CloudApplicationInstanceState{
		CloudApplicationInstanceArgs: args,
		InstanceId:                   &resp.InstanceID,
	}
	return infer.ReadResponse[CloudApplicationInstanceArgs, CloudApplicationInstanceState]{
		ID:     strconv.Itoa(resp.InstanceID),
		Inputs: args,
		State:  state,
	}, nil
}

func (CloudApplicationInstance) Update(ctx context.Context, req infer.UpdateRequest[CloudApplicationInstanceArgs, CloudApplicationInstanceState]) (infer.UpdateResponse[CloudApplicationInstanceState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[CloudApplicationInstanceState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	instanceID := 0
	if req.State.InstanceId != nil {
		instanceID = *req.State.InstanceId
	}
	if instanceID == 0 {
		return infer.UpdateResponse[CloudApplicationInstanceState]{}, fmt.Errorf("no cloud application instance id in state")
	}

	if _, err := cloud_app_instances.Get(ctx, service, instanceID); err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[CloudApplicationInstanceState]{}, nil
		}
		return infer.UpdateResponse[CloudApplicationInstanceState]{}, err
	}

	apiReq := cloudApplicationInstanceToAPI(req.Inputs, instanceID)
	if _, _, err := cloud_app_instances.Update(ctx, service, instanceID, &apiReq); err != nil {
		return infer.UpdateResponse[CloudApplicationInstanceState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[CloudApplicationInstanceState]{}, activationErr
		}
	}

	state := CloudApplicationInstanceState{
		CloudApplicationInstanceArgs: req.Inputs,
		InstanceId:                   &instanceID,
	}
	return infer.UpdateResponse[CloudApplicationInstanceState]{Output: state}, nil
}

func (CloudApplicationInstance) Delete(ctx context.Context, req infer.DeleteRequest[CloudApplicationInstanceState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	instanceID := 0
	if req.State.InstanceId != nil {
		instanceID = *req.State.InstanceId
	}
	if instanceID != 0 {
		if _, err := cloud_app_instances.Delete(ctx, service, instanceID); err != nil {
			return infer.DeleteResponse{}, err
		}
		log.Printf("[INFO] ZIA cloud application instance deleted")
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

func (CloudApplicationInstance) Annotate(a infer.Annotator) {
	describeResource(a, &CloudApplicationInstance{}, `The zia_cloud_application_instance resource manages cloud application instances in the Zscaler Internet Access (ZIA) cloud service. Cloud application instances allow you to define specific tenants or instances of cloud applications for granular policy control.

{{% examples %}}
## Example Usage

{{% example %}}
### Cloud Application Instance

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.CloudApplicationInstance("example", {
    name: "Example Instance",
    instanceType: "SALESFORCE",
    instanceIdentifiers: [{
        instanceIdentifier: "example.my.salesforce.com",
        identifierType: "URL",
    }],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.CloudApplicationInstance("example",
    name="Example Instance",
    instance_type="SALESFORCE",
    instance_identifiers=[{
        "instance_identifier": "example.my.salesforce.com",
        "identifier_type": "URL",
    }],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:CloudApplicationInstance
    properties:
      name: Example Instance
      instanceType: SALESFORCE
      instanceIdentifiers:
        - instanceIdentifier: example.my.salesforce.com
          identifierType: URL
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing Cloud Application Instance can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:CloudApplicationInstance example 12345
`+tripleBacktick("")+`
`)
}

func (a *CloudApplicationInstanceArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the cloud application instance.")
	ann.Describe(&a.InstanceType, "The type of cloud application (e.g. `SALESFORCE`, `SLACK`, `OFFICE365`).")
	ann.Describe(&a.InstanceIdentifiers, "List of instance identifiers for the cloud application.")
}

func (s *CloudApplicationInstanceState) Annotate(a infer.Annotator) {
	a.Describe(&s.InstanceId, "The system-generated ID of the cloud application instance.")
}

var _ infer.CustomResource[CloudApplicationInstanceArgs, CloudApplicationInstanceState] = CloudApplicationInstance{}
