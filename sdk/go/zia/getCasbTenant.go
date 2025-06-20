// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zia

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zia/sdk/go/zia/internal"
)

// * [Official documentation](https://help.zscaler.com/zia/about-saas-application-tenants)
// * [API documentation](https://help.zscaler.com/zia/saas-security-api#/casbTenant/lite-get)
//
// Use the **zia_casb_tenant** data source to get information about a ZIA SaaS Application Tenants in the Zscaler Internet Access cloud or via the API.
//
// ## Example Usage
//
// ### By Name
//
// ### By ID
//
// ### Use Optional Parameters
func GetCasbTenant(ctx *pulumi.Context, args *GetCasbTenantArgs, opts ...pulumi.InvokeOption) (*GetCasbTenantResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCasbTenantResult
	err := ctx.Invoke("zia:index/getCasbTenant:getCasbTenant", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCasbTenant.
type GetCasbTenantArgs struct {
	ActiveOnly                  *bool    `pulumi:"activeOnly"`
	App                         *string  `pulumi:"app"`
	AppType                     *string  `pulumi:"appType"`
	FilterByFeatures            []string `pulumi:"filterByFeatures"`
	IncludeBucketReadyS3Tenants *bool    `pulumi:"includeBucketReadyS3Tenants"`
	IncludeDeleted              *bool    `pulumi:"includeDeleted"`
	ScanConfigTenantsOnly       *bool    `pulumi:"scanConfigTenantsOnly"`
	TenantId                    *int     `pulumi:"tenantId"`
	TenantName                  *string  `pulumi:"tenantName"`
}

// A collection of values returned by getCasbTenant.
type GetCasbTenantResult struct {
	ActiveOnly         *bool    `pulumi:"activeOnly"`
	App                *string  `pulumi:"app"`
	AppType            *string  `pulumi:"appType"`
	EnterpriseTenantId string   `pulumi:"enterpriseTenantId"`
	FeaturesSupporteds []string `pulumi:"featuresSupporteds"`
	FilterByFeatures   []string `pulumi:"filterByFeatures"`
	// The provider-assigned unique ID for this managed resource.
	Id                          string                            `pulumi:"id"`
	IncludeBucketReadyS3Tenants *bool                             `pulumi:"includeBucketReadyS3Tenants"`
	IncludeDeleted              *bool                             `pulumi:"includeDeleted"`
	LastTenantValidationTime    int                               `pulumi:"lastTenantValidationTime"`
	ModifiedTime                int                               `pulumi:"modifiedTime"`
	ReAuth                      bool                              `pulumi:"reAuth"`
	SaasApplication             string                            `pulumi:"saasApplication"`
	ScanConfigTenantsOnly       *bool                             `pulumi:"scanConfigTenantsOnly"`
	Statuses                    []string                          `pulumi:"statuses"`
	TenantDeleted               bool                              `pulumi:"tenantDeleted"`
	TenantId                    int                               `pulumi:"tenantId"`
	TenantName                  string                            `pulumi:"tenantName"`
	TenantWebhookEnabled        bool                              `pulumi:"tenantWebhookEnabled"`
	ZscalerAppTenantIds         []GetCasbTenantZscalerAppTenantId `pulumi:"zscalerAppTenantIds"`
}

func GetCasbTenantOutput(ctx *pulumi.Context, args GetCasbTenantOutputArgs, opts ...pulumi.InvokeOption) GetCasbTenantResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetCasbTenantResultOutput, error) {
			args := v.(GetCasbTenantArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("zia:index/getCasbTenant:getCasbTenant", args, GetCasbTenantResultOutput{}, options).(GetCasbTenantResultOutput), nil
		}).(GetCasbTenantResultOutput)
}

// A collection of arguments for invoking getCasbTenant.
type GetCasbTenantOutputArgs struct {
	ActiveOnly                  pulumi.BoolPtrInput     `pulumi:"activeOnly"`
	App                         pulumi.StringPtrInput   `pulumi:"app"`
	AppType                     pulumi.StringPtrInput   `pulumi:"appType"`
	FilterByFeatures            pulumi.StringArrayInput `pulumi:"filterByFeatures"`
	IncludeBucketReadyS3Tenants pulumi.BoolPtrInput     `pulumi:"includeBucketReadyS3Tenants"`
	IncludeDeleted              pulumi.BoolPtrInput     `pulumi:"includeDeleted"`
	ScanConfigTenantsOnly       pulumi.BoolPtrInput     `pulumi:"scanConfigTenantsOnly"`
	TenantId                    pulumi.IntPtrInput      `pulumi:"tenantId"`
	TenantName                  pulumi.StringPtrInput   `pulumi:"tenantName"`
}

func (GetCasbTenantOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCasbTenantArgs)(nil)).Elem()
}

// A collection of values returned by getCasbTenant.
type GetCasbTenantResultOutput struct{ *pulumi.OutputState }

func (GetCasbTenantResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCasbTenantResult)(nil)).Elem()
}

func (o GetCasbTenantResultOutput) ToGetCasbTenantResultOutput() GetCasbTenantResultOutput {
	return o
}

func (o GetCasbTenantResultOutput) ToGetCasbTenantResultOutputWithContext(ctx context.Context) GetCasbTenantResultOutput {
	return o
}

func (o GetCasbTenantResultOutput) ActiveOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetCasbTenantResult) *bool { return v.ActiveOnly }).(pulumi.BoolPtrOutput)
}

func (o GetCasbTenantResultOutput) App() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCasbTenantResult) *string { return v.App }).(pulumi.StringPtrOutput)
}

func (o GetCasbTenantResultOutput) AppType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCasbTenantResult) *string { return v.AppType }).(pulumi.StringPtrOutput)
}

func (o GetCasbTenantResultOutput) EnterpriseTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCasbTenantResult) string { return v.EnterpriseTenantId }).(pulumi.StringOutput)
}

func (o GetCasbTenantResultOutput) FeaturesSupporteds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCasbTenantResult) []string { return v.FeaturesSupporteds }).(pulumi.StringArrayOutput)
}

func (o GetCasbTenantResultOutput) FilterByFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCasbTenantResult) []string { return v.FilterByFeatures }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCasbTenantResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCasbTenantResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCasbTenantResultOutput) IncludeBucketReadyS3Tenants() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetCasbTenantResult) *bool { return v.IncludeBucketReadyS3Tenants }).(pulumi.BoolPtrOutput)
}

func (o GetCasbTenantResultOutput) IncludeDeleted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetCasbTenantResult) *bool { return v.IncludeDeleted }).(pulumi.BoolPtrOutput)
}

func (o GetCasbTenantResultOutput) LastTenantValidationTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetCasbTenantResult) int { return v.LastTenantValidationTime }).(pulumi.IntOutput)
}

func (o GetCasbTenantResultOutput) ModifiedTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetCasbTenantResult) int { return v.ModifiedTime }).(pulumi.IntOutput)
}

func (o GetCasbTenantResultOutput) ReAuth() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCasbTenantResult) bool { return v.ReAuth }).(pulumi.BoolOutput)
}

func (o GetCasbTenantResultOutput) SaasApplication() pulumi.StringOutput {
	return o.ApplyT(func(v GetCasbTenantResult) string { return v.SaasApplication }).(pulumi.StringOutput)
}

func (o GetCasbTenantResultOutput) ScanConfigTenantsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetCasbTenantResult) *bool { return v.ScanConfigTenantsOnly }).(pulumi.BoolPtrOutput)
}

func (o GetCasbTenantResultOutput) Statuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCasbTenantResult) []string { return v.Statuses }).(pulumi.StringArrayOutput)
}

func (o GetCasbTenantResultOutput) TenantDeleted() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCasbTenantResult) bool { return v.TenantDeleted }).(pulumi.BoolOutput)
}

func (o GetCasbTenantResultOutput) TenantId() pulumi.IntOutput {
	return o.ApplyT(func(v GetCasbTenantResult) int { return v.TenantId }).(pulumi.IntOutput)
}

func (o GetCasbTenantResultOutput) TenantName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCasbTenantResult) string { return v.TenantName }).(pulumi.StringOutput)
}

func (o GetCasbTenantResultOutput) TenantWebhookEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCasbTenantResult) bool { return v.TenantWebhookEnabled }).(pulumi.BoolOutput)
}

func (o GetCasbTenantResultOutput) ZscalerAppTenantIds() GetCasbTenantZscalerAppTenantIdArrayOutput {
	return o.ApplyT(func(v GetCasbTenantResult) []GetCasbTenantZscalerAppTenantId { return v.ZscalerAppTenantIds }).(GetCasbTenantZscalerAppTenantIdArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCasbTenantResultOutput{})
}
