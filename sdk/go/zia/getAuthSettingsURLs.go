// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zia

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zia/sdk/go/zia/internal"
)

// Use the **zia_auth_settings_urls** data source to get a list of URLs that were exempted from cookie authentiation and SSL Inspection in the Zscaler Internet Access cloud or via the API. To learn more see [URL Format Guidelines](https://help.zscaler.com/zia/url-format-guidelines)
//
// ## Example Usage
//
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
//			_, err := zia.LookupAuthSettingsURLs(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupAuthSettingsURLs(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*LookupAuthSettingsURLsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAuthSettingsURLsResult
	err := ctx.Invoke("zia:index/getAuthSettingsURLs:getAuthSettingsURLs", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getAuthSettingsURLs.
type LookupAuthSettingsURLsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id   string   `pulumi:"id"`
	Urls []string `pulumi:"urls"`
}

func LookupAuthSettingsURLsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) LookupAuthSettingsURLsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (LookupAuthSettingsURLsResult, error) {
		r, err := LookupAuthSettingsURLs(ctx, opts...)
		var s LookupAuthSettingsURLsResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(LookupAuthSettingsURLsResultOutput)
}

// A collection of values returned by getAuthSettingsURLs.
type LookupAuthSettingsURLsResultOutput struct{ *pulumi.OutputState }

func (LookupAuthSettingsURLsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthSettingsURLsResult)(nil)).Elem()
}

func (o LookupAuthSettingsURLsResultOutput) ToLookupAuthSettingsURLsResultOutput() LookupAuthSettingsURLsResultOutput {
	return o
}

func (o LookupAuthSettingsURLsResultOutput) ToLookupAuthSettingsURLsResultOutputWithContext(ctx context.Context) LookupAuthSettingsURLsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAuthSettingsURLsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthSettingsURLsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAuthSettingsURLsResultOutput) Urls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthSettingsURLsResult) []string { return v.Urls }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthSettingsURLsResultOutput{})
}
