// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zia

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zia/sdk/go/zia/internal"
)

// * [Official documentation](https://help.zscaler.com/zia/about-cloud-application-instances)
// * [API documentation](https://help.zscaler.com/zia/cloud-app-control-policy#/cloudApplicationInstances-get)
//
// Use the **zia_cloud_application_instance** data source to get information about cloud application instances in the Zscaler Internet Access cloud or via the API.
//
// ## Example Usage
//
// ### By Name
//
// ### By ID
func LookupCloudApplicationInstance(ctx *pulumi.Context, args *LookupCloudApplicationInstanceArgs, opts ...pulumi.InvokeOption) (*LookupCloudApplicationInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCloudApplicationInstanceResult
	err := ctx.Invoke("zia:index/getCloudApplicationInstance:getCloudApplicationInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudApplicationInstance.
type LookupCloudApplicationInstanceArgs struct {
	Id   *int    `pulumi:"id"`
	Name *string `pulumi:"name"`
}

// A collection of values returned by getCloudApplicationInstance.
type LookupCloudApplicationInstanceResult struct {
	Id                  int                                             `pulumi:"id"`
	InstanceIdentifiers []GetCloudApplicationInstanceInstanceIdentifier `pulumi:"instanceIdentifiers"`
	InstanceType        string                                          `pulumi:"instanceType"`
	LastModifiedBies    []GetCloudApplicationInstanceLastModifiedBy     `pulumi:"lastModifiedBies"`
	ModifiedAt          int                                             `pulumi:"modifiedAt"`
	Name                string                                          `pulumi:"name"`
}

func LookupCloudApplicationInstanceOutput(ctx *pulumi.Context, args LookupCloudApplicationInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupCloudApplicationInstanceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCloudApplicationInstanceResultOutput, error) {
			args := v.(LookupCloudApplicationInstanceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("zia:index/getCloudApplicationInstance:getCloudApplicationInstance", args, LookupCloudApplicationInstanceResultOutput{}, options).(LookupCloudApplicationInstanceResultOutput), nil
		}).(LookupCloudApplicationInstanceResultOutput)
}

// A collection of arguments for invoking getCloudApplicationInstance.
type LookupCloudApplicationInstanceOutputArgs struct {
	Id   pulumi.IntPtrInput    `pulumi:"id"`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupCloudApplicationInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudApplicationInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getCloudApplicationInstance.
type LookupCloudApplicationInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupCloudApplicationInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudApplicationInstanceResult)(nil)).Elem()
}

func (o LookupCloudApplicationInstanceResultOutput) ToLookupCloudApplicationInstanceResultOutput() LookupCloudApplicationInstanceResultOutput {
	return o
}

func (o LookupCloudApplicationInstanceResultOutput) ToLookupCloudApplicationInstanceResultOutputWithContext(ctx context.Context) LookupCloudApplicationInstanceResultOutput {
	return o
}

func (o LookupCloudApplicationInstanceResultOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudApplicationInstanceResult) int { return v.Id }).(pulumi.IntOutput)
}

func (o LookupCloudApplicationInstanceResultOutput) InstanceIdentifiers() GetCloudApplicationInstanceInstanceIdentifierArrayOutput {
	return o.ApplyT(func(v LookupCloudApplicationInstanceResult) []GetCloudApplicationInstanceInstanceIdentifier {
		return v.InstanceIdentifiers
	}).(GetCloudApplicationInstanceInstanceIdentifierArrayOutput)
}

func (o LookupCloudApplicationInstanceResultOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudApplicationInstanceResult) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o LookupCloudApplicationInstanceResultOutput) LastModifiedBies() GetCloudApplicationInstanceLastModifiedByArrayOutput {
	return o.ApplyT(func(v LookupCloudApplicationInstanceResult) []GetCloudApplicationInstanceLastModifiedBy {
		return v.LastModifiedBies
	}).(GetCloudApplicationInstanceLastModifiedByArrayOutput)
}

func (o LookupCloudApplicationInstanceResultOutput) ModifiedAt() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudApplicationInstanceResult) int { return v.ModifiedAt }).(pulumi.IntOutput)
}

func (o LookupCloudApplicationInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudApplicationInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudApplicationInstanceResultOutput{})
}
