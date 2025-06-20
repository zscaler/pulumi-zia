// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zia

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zia/sdk/go/zia/internal"
)

// * [Official documentation](https://help.zscaler.com/zia/about-gateways-proxies)
// * [API documentation](https://help.zscaler.com/zia/forwarding-control-policy#/proxyGateways-get)
//
// Use the **zia_forwarding_control_proxy_gateway** data source to retrieve the proxy gateway information. This data source can then be associated with the attribute `proxyGateway` when creating a Forwarding Control Rule via the resource: `ForwardingControlRule`
//
// ## Example Usage
func GetForwardingProxyGateway(ctx *pulumi.Context, args *GetForwardingProxyGatewayArgs, opts ...pulumi.InvokeOption) (*GetForwardingProxyGatewayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetForwardingProxyGatewayResult
	err := ctx.Invoke("zia:index/getForwardingProxyGateway:getForwardingProxyGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getForwardingProxyGateway.
type GetForwardingProxyGatewayArgs struct {
	// The ID of the forwarding control Proxy Gateway resource.
	Id *int `pulumi:"id"`
	// The name of the forwarding control Proxy Gateway to be exported.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getForwardingProxyGateway.
type GetForwardingProxyGatewayResult struct {
	// (string) - Additional details about the Proxy gateway
	Description string `pulumi:"description"`
	// (Boolean) - Indicates whether fail close is enabled to drop the traffic or disabled to allow the traffic when both primary and secondary proxies defined in this gateway are unreachable.
	FailClosed bool `pulumi:"failClosed"`
	// (string) A unique identifier for the secondary proxy gateway
	Id int `pulumi:"id"`
	// (list) -  Information about the admin user that last modified the Proxy gateway
	LastModifiedBies []GetForwardingProxyGatewayLastModifiedBy `pulumi:"lastModifiedBies"`
	// (int) - Timestamp when the ZPA gateway was last modified
	LastModifiedTime int `pulumi:"lastModifiedTime"`
	// (string) The configured name for the secondary proxy gateway
	Name string `pulumi:"name"`
	// (Set of String) - The primary proxy for the gateway. This field is not applicable to the Lite API.
	PrimaryProxies []GetForwardingProxyGatewayPrimaryProxy `pulumi:"primaryProxies"`
	// () - The secondary proxy for the gateway. This field is not applicable to the Lite API.
	SecondaryProxies []GetForwardingProxyGatewaySecondaryProxy `pulumi:"secondaryProxies"`
	// (string) - Indicates whether the type of Proxy gateway. Returned values are: `PROXYCHAIN`, `ZIA`, or `ECSELF`
	Type string `pulumi:"type"`
}

func GetForwardingProxyGatewayOutput(ctx *pulumi.Context, args GetForwardingProxyGatewayOutputArgs, opts ...pulumi.InvokeOption) GetForwardingProxyGatewayResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetForwardingProxyGatewayResultOutput, error) {
			args := v.(GetForwardingProxyGatewayArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("zia:index/getForwardingProxyGateway:getForwardingProxyGateway", args, GetForwardingProxyGatewayResultOutput{}, options).(GetForwardingProxyGatewayResultOutput), nil
		}).(GetForwardingProxyGatewayResultOutput)
}

// A collection of arguments for invoking getForwardingProxyGateway.
type GetForwardingProxyGatewayOutputArgs struct {
	// The ID of the forwarding control Proxy Gateway resource.
	Id pulumi.IntPtrInput `pulumi:"id"`
	// The name of the forwarding control Proxy Gateway to be exported.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetForwardingProxyGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetForwardingProxyGatewayArgs)(nil)).Elem()
}

// A collection of values returned by getForwardingProxyGateway.
type GetForwardingProxyGatewayResultOutput struct{ *pulumi.OutputState }

func (GetForwardingProxyGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetForwardingProxyGatewayResult)(nil)).Elem()
}

func (o GetForwardingProxyGatewayResultOutput) ToGetForwardingProxyGatewayResultOutput() GetForwardingProxyGatewayResultOutput {
	return o
}

func (o GetForwardingProxyGatewayResultOutput) ToGetForwardingProxyGatewayResultOutputWithContext(ctx context.Context) GetForwardingProxyGatewayResultOutput {
	return o
}

// (string) - Additional details about the Proxy gateway
func (o GetForwardingProxyGatewayResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetForwardingProxyGatewayResult) string { return v.Description }).(pulumi.StringOutput)
}

// (Boolean) - Indicates whether fail close is enabled to drop the traffic or disabled to allow the traffic when both primary and secondary proxies defined in this gateway are unreachable.
func (o GetForwardingProxyGatewayResultOutput) FailClosed() pulumi.BoolOutput {
	return o.ApplyT(func(v GetForwardingProxyGatewayResult) bool { return v.FailClosed }).(pulumi.BoolOutput)
}

// (string) A unique identifier for the secondary proxy gateway
func (o GetForwardingProxyGatewayResultOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetForwardingProxyGatewayResult) int { return v.Id }).(pulumi.IntOutput)
}

// (list) -  Information about the admin user that last modified the Proxy gateway
func (o GetForwardingProxyGatewayResultOutput) LastModifiedBies() GetForwardingProxyGatewayLastModifiedByArrayOutput {
	return o.ApplyT(func(v GetForwardingProxyGatewayResult) []GetForwardingProxyGatewayLastModifiedBy {
		return v.LastModifiedBies
	}).(GetForwardingProxyGatewayLastModifiedByArrayOutput)
}

// (int) - Timestamp when the ZPA gateway was last modified
func (o GetForwardingProxyGatewayResultOutput) LastModifiedTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetForwardingProxyGatewayResult) int { return v.LastModifiedTime }).(pulumi.IntOutput)
}

// (string) The configured name for the secondary proxy gateway
func (o GetForwardingProxyGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetForwardingProxyGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

// (Set of String) - The primary proxy for the gateway. This field is not applicable to the Lite API.
func (o GetForwardingProxyGatewayResultOutput) PrimaryProxies() GetForwardingProxyGatewayPrimaryProxyArrayOutput {
	return o.ApplyT(func(v GetForwardingProxyGatewayResult) []GetForwardingProxyGatewayPrimaryProxy {
		return v.PrimaryProxies
	}).(GetForwardingProxyGatewayPrimaryProxyArrayOutput)
}

// () - The secondary proxy for the gateway. This field is not applicable to the Lite API.
func (o GetForwardingProxyGatewayResultOutput) SecondaryProxies() GetForwardingProxyGatewaySecondaryProxyArrayOutput {
	return o.ApplyT(func(v GetForwardingProxyGatewayResult) []GetForwardingProxyGatewaySecondaryProxy {
		return v.SecondaryProxies
	}).(GetForwardingProxyGatewaySecondaryProxyArrayOutput)
}

// (string) - Indicates whether the type of Proxy gateway. Returned values are: `PROXYCHAIN`, `ZIA`, or `ECSELF`
func (o GetForwardingProxyGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetForwardingProxyGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetForwardingProxyGatewayResultOutput{})
}
