// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package activation

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/zscaler/pulumi-zia/sdk/go/zia/Activation"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Activation.GetActivationStatus(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupActivationStatus(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*LookupActivationStatusResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupActivationStatusResult
	err := ctx.Invoke("zia:Activation/getActivationStatus:getActivationStatus", nil, &rv, opts...)
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
