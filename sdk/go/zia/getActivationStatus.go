// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zia

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zia/sdk/go/zia/internal"
)

// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/zscaler/pulumi-zia/sdk/go/zia"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := zia.LookupActivationStatus(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupActivationStatus(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*LookupActivationStatusResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupActivationStatusResult
	err := ctx.Invoke("zia:index/getActivationStatus:getActivationStatus", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getActivationStatus.
type LookupActivationStatusResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Status string `pulumi:"status"`
}

func LookupActivationStatusOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) LookupActivationStatusResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (LookupActivationStatusResult, error) {
		r, err := LookupActivationStatus(ctx, opts...)
		var s LookupActivationStatusResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(LookupActivationStatusResultOutput)
}

// A collection of values returned by getActivationStatus.
type LookupActivationStatusResultOutput struct{ *pulumi.OutputState }

func (LookupActivationStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActivationStatusResult)(nil)).Elem()
}

func (o LookupActivationStatusResultOutput) ToLookupActivationStatusResultOutput() LookupActivationStatusResultOutput {
	return o
}

func (o LookupActivationStatusResultOutput) ToLookupActivationStatusResultOutputWithContext(ctx context.Context) LookupActivationStatusResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupActivationStatusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivationStatusResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupActivationStatusResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivationStatusResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupActivationStatusResultOutput{})
}
