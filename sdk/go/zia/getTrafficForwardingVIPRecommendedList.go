// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zia

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zia/sdk/go/zia/internal"
)

// Use the **zia_gre_vip_recommended_list** data source to get information about a list of recommended GRE tunnel virtual IP addresses (VIPs), based on source IP address or latitude/longitude coordinates.
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
//			_, err := zia.GetTrafficForwardingVIPRecommendedList(ctx, &zia.GetTrafficForwardingVIPRecommendedListArgs{
//				RequiredCount: pulumi.IntRef(2),
//				SourceIp:      pulumi.StringRef("1.1.1.1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetTrafficForwardingVIPRecommendedList(ctx *pulumi.Context, args *GetTrafficForwardingVIPRecommendedListArgs, opts ...pulumi.InvokeOption) (*GetTrafficForwardingVIPRecommendedListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTrafficForwardingVIPRecommendedListResult
	err := ctx.Invoke("zia:index/getTrafficForwardingVIPRecommendedList:getTrafficForwardingVIPRecommendedList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTrafficForwardingVIPRecommendedList.
type GetTrafficForwardingVIPRecommendedListArgs struct {
	// Number of IP address to be exported.
	RequiredCount *int `pulumi:"requiredCount"`
	// Filter based on an IP address range.
	SourceIp *string `pulumi:"sourceIp"`
}

// A collection of values returned by getTrafficForwardingVIPRecommendedList.
type GetTrafficForwardingVIPRecommendedListResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id            string                                       `pulumi:"id"`
	Lists         []GetTrafficForwardingVIPRecommendedListList `pulumi:"lists"`
	RequiredCount *int                                         `pulumi:"requiredCount"`
	// (String) The public source IP address.
	SourceIp *string `pulumi:"sourceIp"`
}

func GetTrafficForwardingVIPRecommendedListOutput(ctx *pulumi.Context, args GetTrafficForwardingVIPRecommendedListOutputArgs, opts ...pulumi.InvokeOption) GetTrafficForwardingVIPRecommendedListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTrafficForwardingVIPRecommendedListResult, error) {
			args := v.(GetTrafficForwardingVIPRecommendedListArgs)
			r, err := GetTrafficForwardingVIPRecommendedList(ctx, &args, opts...)
			var s GetTrafficForwardingVIPRecommendedListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTrafficForwardingVIPRecommendedListResultOutput)
}

// A collection of arguments for invoking getTrafficForwardingVIPRecommendedList.
type GetTrafficForwardingVIPRecommendedListOutputArgs struct {
	// Number of IP address to be exported.
	RequiredCount pulumi.IntPtrInput `pulumi:"requiredCount"`
	// Filter based on an IP address range.
	SourceIp pulumi.StringPtrInput `pulumi:"sourceIp"`
}

func (GetTrafficForwardingVIPRecommendedListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTrafficForwardingVIPRecommendedListArgs)(nil)).Elem()
}

// A collection of values returned by getTrafficForwardingVIPRecommendedList.
type GetTrafficForwardingVIPRecommendedListResultOutput struct{ *pulumi.OutputState }

func (GetTrafficForwardingVIPRecommendedListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTrafficForwardingVIPRecommendedListResult)(nil)).Elem()
}

func (o GetTrafficForwardingVIPRecommendedListResultOutput) ToGetTrafficForwardingVIPRecommendedListResultOutput() GetTrafficForwardingVIPRecommendedListResultOutput {
	return o
}

func (o GetTrafficForwardingVIPRecommendedListResultOutput) ToGetTrafficForwardingVIPRecommendedListResultOutputWithContext(ctx context.Context) GetTrafficForwardingVIPRecommendedListResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetTrafficForwardingVIPRecommendedListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrafficForwardingVIPRecommendedListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTrafficForwardingVIPRecommendedListResultOutput) Lists() GetTrafficForwardingVIPRecommendedListListArrayOutput {
	return o.ApplyT(func(v GetTrafficForwardingVIPRecommendedListResult) []GetTrafficForwardingVIPRecommendedListList {
		return v.Lists
	}).(GetTrafficForwardingVIPRecommendedListListArrayOutput)
}

func (o GetTrafficForwardingVIPRecommendedListResultOutput) RequiredCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetTrafficForwardingVIPRecommendedListResult) *int { return v.RequiredCount }).(pulumi.IntPtrOutput)
}

// (String) The public source IP address.
func (o GetTrafficForwardingVIPRecommendedListResultOutput) SourceIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrafficForwardingVIPRecommendedListResult) *string { return v.SourceIp }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTrafficForwardingVIPRecommendedListResultOutput{})
}
