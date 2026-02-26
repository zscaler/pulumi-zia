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

// Package provider implements the getDlpCloudToCloudIr invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_dlp_cloud_to_cloud_ir.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/c2c_incident_receiver"
)

type GetDlpCloudToCloudIrArgs struct {
	Id   *int    `pulumi:"resourceId,optional"`
	Name *string `pulumi:"name,optional"`
}

type LastModifiedByOutput struct {
	Id   int    `pulumi:"resourceId"`
	Name string `pulumi:"name"`
}

type LastValidationMsgOutput struct {
	ErrorMsg  string `pulumi:"errorMsg"`
	ErrorCode string `pulumi:"errorCode"`
}

type OnboardableEntityOutput struct {
	Id                 int                      `pulumi:"resourceId"`
	Name               string                   `pulumi:"name"`
	Type               string                   `pulumi:"type"`
	EnterpriseTenantId string                   `pulumi:"enterpriseTenantId"`
	Application        string                   `pulumi:"application"`
	LastValidationMsg  *LastValidationMsgOutput `pulumi:"lastValidationMsg,optional"`
	ZscalerAppTenantId *LastModifiedByOutput    `pulumi:"zscalerAppTenantId,optional"`
}

type GetDlpCloudToCloudIrResult struct {
	Id                       int                      `pulumi:"resourceId"`
	Name                     string                   `pulumi:"name"`
	Status                   []string                 `pulumi:"status"`
	ModifiedTime             int                      `pulumi:"modifiedTime"`
	LastTenantValidationTime int                      `pulumi:"lastTenantValidationTime"`
	LastModifiedBy           *LastModifiedByOutput    `pulumi:"lastModifiedBy,optional"`
	LastValidationMsg        *LastValidationMsgOutput `pulumi:"lastValidationMsg,optional"`
	OnboardableEntity        *OnboardableEntityOutput `pulumi:"onboardableEntity,optional"`
}

type GetDlpCloudToCloudIr struct{}

func (f *GetDlpCloudToCloudIr) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a DLP cloud-to-cloud incident receiver by ID or name.")
}

func (a *GetDlpCloudToCloudIrArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Id, "The ID of the cloud-to-cloud incident receiver to look up.")
	ann.Describe(&a.Name, "The name of the cloud-to-cloud incident receiver to look up.")
}

func (r *GetDlpCloudToCloudIrResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the cloud-to-cloud incident receiver.")
	a.Describe(&r.Name, "The name of the cloud-to-cloud incident receiver.")
	a.Describe(&r.Status, "The status of the incident receiver.")
	a.Describe(&r.ModifiedTime, "The last modification time (epoch).")
	a.Describe(&r.LastTenantValidationTime, "The last tenant validation time (epoch).")
	a.Describe(&r.LastModifiedBy, "The user who last modified the incident receiver.")
	a.Describe(&r.LastValidationMsg, "The last validation message.")
	a.Describe(&r.OnboardableEntity, "The onboardable entity associated with the incident receiver.")
}

func (*GetDlpCloudToCloudIr) Invoke(ctx context.Context, req infer.FunctionRequest[GetDlpCloudToCloudIrArgs]) (infer.FunctionResponse[GetDlpCloudToCloudIrResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetDlpCloudToCloudIrResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	if (req.Input.Id == nil || *req.Input.Id == 0) && (req.Input.Name == nil || *req.Input.Name == "") {
		return infer.FunctionResponse[GetDlpCloudToCloudIrResult]{}, fmt.Errorf("either 'id' or 'name' must be provided")
	}

	var resp *c2c_incident_receiver.C2CIncidentReceiver
	if req.Input.Id != nil && *req.Input.Id != 0 {
		r, err := c2c_incident_receiver.Get(ctx, svc, *req.Input.Id)
		if err != nil {
			return infer.FunctionResponse[GetDlpCloudToCloudIrResult]{}, fmt.Errorf("failed to get cloud-to-cloud IR with ID %d: %w", *req.Input.Id, err)
		}
		resp = r
	}
	if resp == nil && req.Input.Name != nil && *req.Input.Name != "" {
		r, err := c2c_incident_receiver.GetC2CIRName(ctx, svc, *req.Input.Name)
		if err != nil {
			return infer.FunctionResponse[GetDlpCloudToCloudIrResult]{}, fmt.Errorf("failed to get cloud-to-cloud IR with name %s: %w", *req.Input.Name, err)
		}
		resp = r
	}
	if resp == nil {
		return infer.FunctionResponse[GetDlpCloudToCloudIrResult]{}, fmt.Errorf("received nil response from cloud-to-cloud IR API")
	}

	result := GetDlpCloudToCloudIrResult{
		Id:                       resp.ID,
		Name:                     resp.Name,
		Status:                   resp.Status,
		ModifiedTime:             resp.ModifiedTime,
		LastTenantValidationTime: resp.LastTenantValidationTime,
	}
	if resp.LastModifiedBy != nil {
		result.LastModifiedBy = &LastModifiedByOutput{Id: resp.LastModifiedBy.ID, Name: resp.LastModifiedBy.Name}
	}
	if resp.LastValidationMsg != nil && (resp.LastValidationMsg.ErrorMsg != "" || resp.LastValidationMsg.ErrorCode != "") {
		result.LastValidationMsg = &LastValidationMsgOutput{
			ErrorMsg: resp.LastValidationMsg.ErrorMsg, ErrorCode: resp.LastValidationMsg.ErrorCode,
		}
	}
	if resp.OnboardableEntity != nil {
		oe := &OnboardableEntityOutput{
			Id:                 resp.OnboardableEntity.ID,
			Name:               resp.OnboardableEntity.Name,
			Type:               resp.OnboardableEntity.Type,
			EnterpriseTenantId: resp.OnboardableEntity.EnterpriseTenantID,
			Application:        resp.OnboardableEntity.Application,
		}
		if resp.OnboardableEntity.LastValidationMsg.ErrorMsg != "" || resp.OnboardableEntity.LastValidationMsg.ErrorCode != "" {
			oe.LastValidationMsg = &LastValidationMsgOutput{
				ErrorMsg:  resp.OnboardableEntity.LastValidationMsg.ErrorMsg,
				ErrorCode: resp.OnboardableEntity.LastValidationMsg.ErrorCode,
			}
		}
		if resp.OnboardableEntity.ZscalerAppTenantID != nil {
			oe.ZscalerAppTenantId = &LastModifiedByOutput{Id: resp.OnboardableEntity.ZscalerAppTenantID.ID, Name: resp.OnboardableEntity.ZscalerAppTenantID.Name}
		}
		result.OnboardableEntity = oe
	}
	return infer.FunctionResponse[GetDlpCloudToCloudIrResult]{Output: result}, nil
}
