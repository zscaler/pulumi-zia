// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zia

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zia/sdk/go/zia/internal"
)

// Use the **zia_dlp_incident_receiver_servers** data source to get information about a ZIA DLP Incident Receiver Server in the Zscaler Internet Access cloud or via the API.
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
//			_, err := zia.GetDLPIncidentReceiverServers(ctx, &zia.GetDLPIncidentReceiverServersArgs{
//				Name: pulumi.StringRef("ZS_Incident_Receiver"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetDLPIncidentReceiverServers(ctx *pulumi.Context, args *GetDLPIncidentReceiverServersArgs, opts ...pulumi.InvokeOption) (*GetDLPIncidentReceiverServersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDLPIncidentReceiverServersResult
	err := ctx.Invoke("zia:index/getDLPIncidentReceiverServers:getDLPIncidentReceiverServers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDLPIncidentReceiverServers.
type GetDLPIncidentReceiverServersArgs struct {
	Name *string `pulumi:"name"`
}

// A collection of values returned by getDLPIncidentReceiverServers.
type GetDLPIncidentReceiverServersResult struct {
	Flags  int     `pulumi:"flags"`
	Id     int     `pulumi:"id"`
	Name   *string `pulumi:"name"`
	Status string  `pulumi:"status"`
	Url    string  `pulumi:"url"`
}

func GetDLPIncidentReceiverServersOutput(ctx *pulumi.Context, args GetDLPIncidentReceiverServersOutputArgs, opts ...pulumi.InvokeOption) GetDLPIncidentReceiverServersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDLPIncidentReceiverServersResult, error) {
			args := v.(GetDLPIncidentReceiverServersArgs)
			r, err := GetDLPIncidentReceiverServers(ctx, &args, opts...)
			var s GetDLPIncidentReceiverServersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDLPIncidentReceiverServersResultOutput)
}

// A collection of arguments for invoking getDLPIncidentReceiverServers.
type GetDLPIncidentReceiverServersOutputArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetDLPIncidentReceiverServersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDLPIncidentReceiverServersArgs)(nil)).Elem()
}

// A collection of values returned by getDLPIncidentReceiverServers.
type GetDLPIncidentReceiverServersResultOutput struct{ *pulumi.OutputState }

func (GetDLPIncidentReceiverServersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDLPIncidentReceiverServersResult)(nil)).Elem()
}

func (o GetDLPIncidentReceiverServersResultOutput) ToGetDLPIncidentReceiverServersResultOutput() GetDLPIncidentReceiverServersResultOutput {
	return o
}

func (o GetDLPIncidentReceiverServersResultOutput) ToGetDLPIncidentReceiverServersResultOutputWithContext(ctx context.Context) GetDLPIncidentReceiverServersResultOutput {
	return o
}

func (o GetDLPIncidentReceiverServersResultOutput) Flags() pulumi.IntOutput {
	return o.ApplyT(func(v GetDLPIncidentReceiverServersResult) int { return v.Flags }).(pulumi.IntOutput)
}

func (o GetDLPIncidentReceiverServersResultOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetDLPIncidentReceiverServersResult) int { return v.Id }).(pulumi.IntOutput)
}

func (o GetDLPIncidentReceiverServersResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDLPIncidentReceiverServersResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetDLPIncidentReceiverServersResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetDLPIncidentReceiverServersResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o GetDLPIncidentReceiverServersResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v GetDLPIncidentReceiverServersResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDLPIncidentReceiverServersResultOutput{})
}
