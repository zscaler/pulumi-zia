// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zia

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zia/sdk/go/zia/internal"
)

// The **zia_traffic_forwarding_gre_tunnel** resource allows the creation and management of GRE tunnel configuration in the Zscaler Internet Access (ZIA) portal.
//
// > **Note:** The provider automatically query the Zscaler cloud for the primary and secondary destination datacenter and virtual IP address (VIP) of the GRE tunnel. The parameter can be overriden if needed by setting the parameters: `primaryDestVip` and `secondaryDestVip`.
//
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
//			// ZIA Traffic Forwarding - Static IP
//			exampleTrafficForwardingStaticIP, err := zia.NewTrafficForwardingStaticIP(ctx, "exampleTrafficForwardingStaticIP", &zia.TrafficForwardingStaticIPArgs{
//				IpAddress:   pulumi.String("1.1.1.1"),
//				RoutableIp:  pulumi.Bool(true),
//				Comment:     pulumi.String("Example"),
//				GeoOverride: pulumi.Bool(true),
//				Latitude:    pulumi.Float64(37.418171),
//				Longitude:   -121.95314,
//			})
//			if err != nil {
//				return err
//			}
//			// Creates a numbered GRE Tunnel
//			_, err = zia.NewTrafficForwardingGRETunnel(ctx, "exampleTrafficForwardingGRETunnel", &zia.TrafficForwardingGRETunnelArgs{
//				SourceIp:      exampleTrafficForwardingStaticIP.IpAddress,
//				Comment:       pulumi.String("Example"),
//				WithinCountry: pulumi.Bool(true),
//				CountryCode:   pulumi.String("US"),
//				IpUnnumbered:  pulumi.Bool(false),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				exampleTrafficForwardingStaticIP,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
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
//			thisTrafficForwardingStaticIP, err := zia.NewTrafficForwardingStaticIP(ctx, "thisTrafficForwardingStaticIP", &zia.TrafficForwardingStaticIPArgs{
//				IpAddress:   pulumi.String("50.98.112.169"),
//				RoutableIp:  pulumi.Bool(true),
//				Comment:     pulumi.String("Created with Terraform"),
//				GeoOverride: pulumi.Bool(true),
//				Latitude:    pulumi.Float64(49.0526),
//				Longitude:   -122.8291,
//			})
//			if err != nil {
//				return err
//			}
//			thisTrafficForwardingVIPRecommendedList := thisTrafficForwardingStaticIP.IpAddress.ApplyT(func(ipAddress string) (zia.GetTrafficForwardingVIPRecommendedListResult, error) {
//				return zia.GetTrafficForwardingVIPRecommendedListOutput(ctx, zia.GetTrafficForwardingVIPRecommendedListOutputArgs{
//					SourceIp:      ipAddress,
//					RequiredCount: 2,
//				}, nil), nil
//			}).(zia.GetTrafficForwardingVIPRecommendedListResultOutput)
//			_, err = zia.GetTrafficForwardingGREInternalIPRange(ctx, &zia.GetTrafficForwardingGREInternalIPRangeArgs{
//				RequiredCount: pulumi.IntRef(10),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zia.NewTrafficForwardingGRETunnel(ctx, "thisTrafficForwardingGRETunnel", &zia.TrafficForwardingGRETunnelArgs{
//				SourceIp:      thisTrafficForwardingStaticIP.IpAddress,
//				Comment:       pulumi.String("GRE Tunnel Created with Terraform"),
//				WithinCountry: pulumi.Bool(false),
//				CountryCode:   pulumi.String("CA"),
//				IpUnnumbered:  pulumi.Bool(false),
//				PrimaryDestVips: zia.TrafficForwardingGRETunnelPrimaryDestVipArray{
//					&zia.TrafficForwardingGRETunnelPrimaryDestVipArgs{
//						Datacenter: thisTrafficForwardingVIPRecommendedList.ApplyT(func(thisTrafficForwardingVIPRecommendedList zia.GetTrafficForwardingVIPRecommendedListResult) (*string, error) {
//							return &thisTrafficForwardingVIPRecommendedList.Lists[0].Datacenter, nil
//						}).(pulumi.StringPtrOutput),
//						VirtualIp: thisTrafficForwardingVIPRecommendedList.ApplyT(func(thisTrafficForwardingVIPRecommendedList zia.GetTrafficForwardingVIPRecommendedListResult) (*string, error) {
//							return &thisTrafficForwardingVIPRecommendedList.Lists[0].VirtualIp, nil
//						}).(pulumi.StringPtrOutput),
//					},
//				},
//				SecondaryDestVips: zia.TrafficForwardingGRETunnelSecondaryDestVipArray{
//					&zia.TrafficForwardingGRETunnelSecondaryDestVipArgs{
//						Datacenter: thisTrafficForwardingVIPRecommendedList.ApplyT(func(thisTrafficForwardingVIPRecommendedList zia.GetTrafficForwardingVIPRecommendedListResult) (*string, error) {
//							return &thisTrafficForwardingVIPRecommendedList.Lists[1].Datacenter, nil
//						}).(pulumi.StringPtrOutput),
//						VirtualIp: thisTrafficForwardingVIPRecommendedList.ApplyT(func(thisTrafficForwardingVIPRecommendedList zia.GetTrafficForwardingVIPRecommendedListResult) (*string, error) {
//							return &thisTrafficForwardingVIPRecommendedList.Lists[1].VirtualIp, nil
//						}).(pulumi.StringPtrOutput),
//					},
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				thisTrafficForwardingStaticIP,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// > **Note:** Although the example shows 2 valid attributes defined (datacenter, virtual_ip) within the primaryDestVip and secondary_dest_vip, only one attribute is required. If setting the datacenter name as the attribute i.e YVR1. The provider will automatically select the agvaiulable VIP.
//
// > **Note:** To obtain the datacenter codes and/or virtual_ips, refer to the following [Zscaler Portal](https://config.zscaler.com/zscloud.net/cenr) and choose your cloud tenant.
//
// > **Note:** The provider will automatically query and set the Zscaler cloud for the next available `/29` internal IP range to be used in a numbered GRE tunnel.
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
//			// ZIA Traffic Forwarding - Static IP
//			example, err := zia.NewTrafficForwardingStaticIP(ctx, "example", &zia.TrafficForwardingStaticIPArgs{
//				IpAddress:   pulumi.String("1.1.1.1"),
//				RoutableIp:  pulumi.Bool(true),
//				Comment:     pulumi.String("Example"),
//				GeoOverride: pulumi.Bool(true),
//				Latitude:    pulumi.Float64(37.418171),
//				Longitude:   -121.95314,
//			})
//			if err != nil {
//				return err
//			}
//			// Creates an unnumbered GRE Tunnel
//			_, err = zia.NewTrafficForwardingGRETunnel(ctx, "telusHomeInternet01Gre01", &zia.TrafficForwardingGRETunnelArgs{
//				SourceIp:      example.IpAddress,
//				Comment:       pulumi.String("Example"),
//				WithinCountry: pulumi.Bool(true),
//				CountryCode:   pulumi.String("CA"),
//				IpUnnumbered:  pulumi.Bool(true),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				example,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.
//
// # Visit
//
// **zia_traffic_forwarding_gre_tunnel** can be imported by using `<TUNNEL_ID>` as the import ID.
//
// For example:
//
// ```sh
// $ pulumi import zia:index/trafficForwardingGRETunnel:TrafficForwardingGRETunnel example <tunnel_id>
// ```
//
// or
//
// ```sh
// $ pulumi import zia:index/trafficForwardingGRETunnel:TrafficForwardingGRETunnel example <engine_name>
// ```
type TrafficForwardingGRETunnel struct {
	pulumi.CustomResourceState

	// Additional information about this GRE tunnel
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// When withinCountry is enabled, you must set this to the country code.
	CountryCode pulumi.StringOutput `pulumi:"countryCode"`
	// The start of the internal IP address in /29 CIDR range. Automatically set by the provider if `ipUnnumbered` is set to `false`.
	InternalIpRange pulumi.StringOutput `pulumi:"internalIpRange"`
	// This is required to support the automated SD-WAN provisioning of GRE tunnels, when set to true greTunIp and greTunId are set to null
	IpUnnumbered pulumi.BoolOutput `pulumi:"ipUnnumbered"`
	// **` (Optional) The primary destination data center and virtual IP address (VIP) of the GRE tunnel.
	PrimaryDestVips TrafficForwardingGRETunnelPrimaryDestVipArrayOutput `pulumi:"primaryDestVips"`
	// The secondary destination data center and virtual IP address (VIP) of the GRE tunnel.
	SecondaryDestVips TrafficForwardingGRETunnelSecondaryDestVipArrayOutput `pulumi:"secondaryDestVips"`
	// The source IP address of the GRE tunnel. This is typically a static IP address in the organization or SD-WAN. This IP address must be provisioned within the Zscaler service using the /staticIP endpoint.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// The ID of the GRE tunnel.
	TunnelId pulumi.IntOutput `pulumi:"tunnelId"`
	// Restrict the data center virtual IP addresses (VIPs) only to those within the same country as the source IP address
	WithinCountry pulumi.BoolOutput `pulumi:"withinCountry"`
}

// NewTrafficForwardingGRETunnel registers a new resource with the given unique name, arguments, and options.
func NewTrafficForwardingGRETunnel(ctx *pulumi.Context,
	name string, args *TrafficForwardingGRETunnelArgs, opts ...pulumi.ResourceOption) (*TrafficForwardingGRETunnel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SourceIp == nil {
		return nil, errors.New("invalid value for required argument 'SourceIp'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TrafficForwardingGRETunnel
	err := ctx.RegisterResource("zia:index/trafficForwardingGRETunnel:TrafficForwardingGRETunnel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficForwardingGRETunnel gets an existing TrafficForwardingGRETunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficForwardingGRETunnel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficForwardingGRETunnelState, opts ...pulumi.ResourceOption) (*TrafficForwardingGRETunnel, error) {
	var resource TrafficForwardingGRETunnel
	err := ctx.ReadResource("zia:index/trafficForwardingGRETunnel:TrafficForwardingGRETunnel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficForwardingGRETunnel resources.
type trafficForwardingGRETunnelState struct {
	// Additional information about this GRE tunnel
	Comment *string `pulumi:"comment"`
	// When withinCountry is enabled, you must set this to the country code.
	CountryCode *string `pulumi:"countryCode"`
	// The start of the internal IP address in /29 CIDR range. Automatically set by the provider if `ipUnnumbered` is set to `false`.
	InternalIpRange *string `pulumi:"internalIpRange"`
	// This is required to support the automated SD-WAN provisioning of GRE tunnels, when set to true greTunIp and greTunId are set to null
	IpUnnumbered *bool `pulumi:"ipUnnumbered"`
	// **` (Optional) The primary destination data center and virtual IP address (VIP) of the GRE tunnel.
	PrimaryDestVips []TrafficForwardingGRETunnelPrimaryDestVip `pulumi:"primaryDestVips"`
	// The secondary destination data center and virtual IP address (VIP) of the GRE tunnel.
	SecondaryDestVips []TrafficForwardingGRETunnelSecondaryDestVip `pulumi:"secondaryDestVips"`
	// The source IP address of the GRE tunnel. This is typically a static IP address in the organization or SD-WAN. This IP address must be provisioned within the Zscaler service using the /staticIP endpoint.
	SourceIp *string `pulumi:"sourceIp"`
	// The ID of the GRE tunnel.
	TunnelId *int `pulumi:"tunnelId"`
	// Restrict the data center virtual IP addresses (VIPs) only to those within the same country as the source IP address
	WithinCountry *bool `pulumi:"withinCountry"`
}

type TrafficForwardingGRETunnelState struct {
	// Additional information about this GRE tunnel
	Comment pulumi.StringPtrInput
	// When withinCountry is enabled, you must set this to the country code.
	CountryCode pulumi.StringPtrInput
	// The start of the internal IP address in /29 CIDR range. Automatically set by the provider if `ipUnnumbered` is set to `false`.
	InternalIpRange pulumi.StringPtrInput
	// This is required to support the automated SD-WAN provisioning of GRE tunnels, when set to true greTunIp and greTunId are set to null
	IpUnnumbered pulumi.BoolPtrInput
	// **` (Optional) The primary destination data center and virtual IP address (VIP) of the GRE tunnel.
	PrimaryDestVips TrafficForwardingGRETunnelPrimaryDestVipArrayInput
	// The secondary destination data center and virtual IP address (VIP) of the GRE tunnel.
	SecondaryDestVips TrafficForwardingGRETunnelSecondaryDestVipArrayInput
	// The source IP address of the GRE tunnel. This is typically a static IP address in the organization or SD-WAN. This IP address must be provisioned within the Zscaler service using the /staticIP endpoint.
	SourceIp pulumi.StringPtrInput
	// The ID of the GRE tunnel.
	TunnelId pulumi.IntPtrInput
	// Restrict the data center virtual IP addresses (VIPs) only to those within the same country as the source IP address
	WithinCountry pulumi.BoolPtrInput
}

func (TrafficForwardingGRETunnelState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficForwardingGRETunnelState)(nil)).Elem()
}

type trafficForwardingGRETunnelArgs struct {
	// Additional information about this GRE tunnel
	Comment *string `pulumi:"comment"`
	// When withinCountry is enabled, you must set this to the country code.
	CountryCode *string `pulumi:"countryCode"`
	// The start of the internal IP address in /29 CIDR range. Automatically set by the provider if `ipUnnumbered` is set to `false`.
	InternalIpRange *string `pulumi:"internalIpRange"`
	// This is required to support the automated SD-WAN provisioning of GRE tunnels, when set to true greTunIp and greTunId are set to null
	IpUnnumbered *bool `pulumi:"ipUnnumbered"`
	// **` (Optional) The primary destination data center and virtual IP address (VIP) of the GRE tunnel.
	PrimaryDestVips []TrafficForwardingGRETunnelPrimaryDestVip `pulumi:"primaryDestVips"`
	// The secondary destination data center and virtual IP address (VIP) of the GRE tunnel.
	SecondaryDestVips []TrafficForwardingGRETunnelSecondaryDestVip `pulumi:"secondaryDestVips"`
	// The source IP address of the GRE tunnel. This is typically a static IP address in the organization or SD-WAN. This IP address must be provisioned within the Zscaler service using the /staticIP endpoint.
	SourceIp string `pulumi:"sourceIp"`
	// Restrict the data center virtual IP addresses (VIPs) only to those within the same country as the source IP address
	WithinCountry *bool `pulumi:"withinCountry"`
}

// The set of arguments for constructing a TrafficForwardingGRETunnel resource.
type TrafficForwardingGRETunnelArgs struct {
	// Additional information about this GRE tunnel
	Comment pulumi.StringPtrInput
	// When withinCountry is enabled, you must set this to the country code.
	CountryCode pulumi.StringPtrInput
	// The start of the internal IP address in /29 CIDR range. Automatically set by the provider if `ipUnnumbered` is set to `false`.
	InternalIpRange pulumi.StringPtrInput
	// This is required to support the automated SD-WAN provisioning of GRE tunnels, when set to true greTunIp and greTunId are set to null
	IpUnnumbered pulumi.BoolPtrInput
	// **` (Optional) The primary destination data center and virtual IP address (VIP) of the GRE tunnel.
	PrimaryDestVips TrafficForwardingGRETunnelPrimaryDestVipArrayInput
	// The secondary destination data center and virtual IP address (VIP) of the GRE tunnel.
	SecondaryDestVips TrafficForwardingGRETunnelSecondaryDestVipArrayInput
	// The source IP address of the GRE tunnel. This is typically a static IP address in the organization or SD-WAN. This IP address must be provisioned within the Zscaler service using the /staticIP endpoint.
	SourceIp pulumi.StringInput
	// Restrict the data center virtual IP addresses (VIPs) only to those within the same country as the source IP address
	WithinCountry pulumi.BoolPtrInput
}

func (TrafficForwardingGRETunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficForwardingGRETunnelArgs)(nil)).Elem()
}

type TrafficForwardingGRETunnelInput interface {
	pulumi.Input

	ToTrafficForwardingGRETunnelOutput() TrafficForwardingGRETunnelOutput
	ToTrafficForwardingGRETunnelOutputWithContext(ctx context.Context) TrafficForwardingGRETunnelOutput
}

func (*TrafficForwardingGRETunnel) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficForwardingGRETunnel)(nil)).Elem()
}

func (i *TrafficForwardingGRETunnel) ToTrafficForwardingGRETunnelOutput() TrafficForwardingGRETunnelOutput {
	return i.ToTrafficForwardingGRETunnelOutputWithContext(context.Background())
}

func (i *TrafficForwardingGRETunnel) ToTrafficForwardingGRETunnelOutputWithContext(ctx context.Context) TrafficForwardingGRETunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficForwardingGRETunnelOutput)
}

// TrafficForwardingGRETunnelArrayInput is an input type that accepts TrafficForwardingGRETunnelArray and TrafficForwardingGRETunnelArrayOutput values.
// You can construct a concrete instance of `TrafficForwardingGRETunnelArrayInput` via:
//
//	TrafficForwardingGRETunnelArray{ TrafficForwardingGRETunnelArgs{...} }
type TrafficForwardingGRETunnelArrayInput interface {
	pulumi.Input

	ToTrafficForwardingGRETunnelArrayOutput() TrafficForwardingGRETunnelArrayOutput
	ToTrafficForwardingGRETunnelArrayOutputWithContext(context.Context) TrafficForwardingGRETunnelArrayOutput
}

type TrafficForwardingGRETunnelArray []TrafficForwardingGRETunnelInput

func (TrafficForwardingGRETunnelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficForwardingGRETunnel)(nil)).Elem()
}

func (i TrafficForwardingGRETunnelArray) ToTrafficForwardingGRETunnelArrayOutput() TrafficForwardingGRETunnelArrayOutput {
	return i.ToTrafficForwardingGRETunnelArrayOutputWithContext(context.Background())
}

func (i TrafficForwardingGRETunnelArray) ToTrafficForwardingGRETunnelArrayOutputWithContext(ctx context.Context) TrafficForwardingGRETunnelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficForwardingGRETunnelArrayOutput)
}

// TrafficForwardingGRETunnelMapInput is an input type that accepts TrafficForwardingGRETunnelMap and TrafficForwardingGRETunnelMapOutput values.
// You can construct a concrete instance of `TrafficForwardingGRETunnelMapInput` via:
//
//	TrafficForwardingGRETunnelMap{ "key": TrafficForwardingGRETunnelArgs{...} }
type TrafficForwardingGRETunnelMapInput interface {
	pulumi.Input

	ToTrafficForwardingGRETunnelMapOutput() TrafficForwardingGRETunnelMapOutput
	ToTrafficForwardingGRETunnelMapOutputWithContext(context.Context) TrafficForwardingGRETunnelMapOutput
}

type TrafficForwardingGRETunnelMap map[string]TrafficForwardingGRETunnelInput

func (TrafficForwardingGRETunnelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficForwardingGRETunnel)(nil)).Elem()
}

func (i TrafficForwardingGRETunnelMap) ToTrafficForwardingGRETunnelMapOutput() TrafficForwardingGRETunnelMapOutput {
	return i.ToTrafficForwardingGRETunnelMapOutputWithContext(context.Background())
}

func (i TrafficForwardingGRETunnelMap) ToTrafficForwardingGRETunnelMapOutputWithContext(ctx context.Context) TrafficForwardingGRETunnelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficForwardingGRETunnelMapOutput)
}

type TrafficForwardingGRETunnelOutput struct{ *pulumi.OutputState }

func (TrafficForwardingGRETunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficForwardingGRETunnel)(nil)).Elem()
}

func (o TrafficForwardingGRETunnelOutput) ToTrafficForwardingGRETunnelOutput() TrafficForwardingGRETunnelOutput {
	return o
}

func (o TrafficForwardingGRETunnelOutput) ToTrafficForwardingGRETunnelOutputWithContext(ctx context.Context) TrafficForwardingGRETunnelOutput {
	return o
}

// Additional information about this GRE tunnel
func (o TrafficForwardingGRETunnelOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficForwardingGRETunnel) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// When withinCountry is enabled, you must set this to the country code.
func (o TrafficForwardingGRETunnelOutput) CountryCode() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficForwardingGRETunnel) pulumi.StringOutput { return v.CountryCode }).(pulumi.StringOutput)
}

// The start of the internal IP address in /29 CIDR range. Automatically set by the provider if `ipUnnumbered` is set to `false`.
func (o TrafficForwardingGRETunnelOutput) InternalIpRange() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficForwardingGRETunnel) pulumi.StringOutput { return v.InternalIpRange }).(pulumi.StringOutput)
}

// This is required to support the automated SD-WAN provisioning of GRE tunnels, when set to true greTunIp and greTunId are set to null
func (o TrafficForwardingGRETunnelOutput) IpUnnumbered() pulumi.BoolOutput {
	return o.ApplyT(func(v *TrafficForwardingGRETunnel) pulumi.BoolOutput { return v.IpUnnumbered }).(pulumi.BoolOutput)
}

// **` (Optional) The primary destination data center and virtual IP address (VIP) of the GRE tunnel.
func (o TrafficForwardingGRETunnelOutput) PrimaryDestVips() TrafficForwardingGRETunnelPrimaryDestVipArrayOutput {
	return o.ApplyT(func(v *TrafficForwardingGRETunnel) TrafficForwardingGRETunnelPrimaryDestVipArrayOutput {
		return v.PrimaryDestVips
	}).(TrafficForwardingGRETunnelPrimaryDestVipArrayOutput)
}

// The secondary destination data center and virtual IP address (VIP) of the GRE tunnel.
func (o TrafficForwardingGRETunnelOutput) SecondaryDestVips() TrafficForwardingGRETunnelSecondaryDestVipArrayOutput {
	return o.ApplyT(func(v *TrafficForwardingGRETunnel) TrafficForwardingGRETunnelSecondaryDestVipArrayOutput {
		return v.SecondaryDestVips
	}).(TrafficForwardingGRETunnelSecondaryDestVipArrayOutput)
}

// The source IP address of the GRE tunnel. This is typically a static IP address in the organization or SD-WAN. This IP address must be provisioned within the Zscaler service using the /staticIP endpoint.
func (o TrafficForwardingGRETunnelOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficForwardingGRETunnel) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// The ID of the GRE tunnel.
func (o TrafficForwardingGRETunnelOutput) TunnelId() pulumi.IntOutput {
	return o.ApplyT(func(v *TrafficForwardingGRETunnel) pulumi.IntOutput { return v.TunnelId }).(pulumi.IntOutput)
}

// Restrict the data center virtual IP addresses (VIPs) only to those within the same country as the source IP address
func (o TrafficForwardingGRETunnelOutput) WithinCountry() pulumi.BoolOutput {
	return o.ApplyT(func(v *TrafficForwardingGRETunnel) pulumi.BoolOutput { return v.WithinCountry }).(pulumi.BoolOutput)
}

type TrafficForwardingGRETunnelArrayOutput struct{ *pulumi.OutputState }

func (TrafficForwardingGRETunnelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficForwardingGRETunnel)(nil)).Elem()
}

func (o TrafficForwardingGRETunnelArrayOutput) ToTrafficForwardingGRETunnelArrayOutput() TrafficForwardingGRETunnelArrayOutput {
	return o
}

func (o TrafficForwardingGRETunnelArrayOutput) ToTrafficForwardingGRETunnelArrayOutputWithContext(ctx context.Context) TrafficForwardingGRETunnelArrayOutput {
	return o
}

func (o TrafficForwardingGRETunnelArrayOutput) Index(i pulumi.IntInput) TrafficForwardingGRETunnelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TrafficForwardingGRETunnel {
		return vs[0].([]*TrafficForwardingGRETunnel)[vs[1].(int)]
	}).(TrafficForwardingGRETunnelOutput)
}

type TrafficForwardingGRETunnelMapOutput struct{ *pulumi.OutputState }

func (TrafficForwardingGRETunnelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficForwardingGRETunnel)(nil)).Elem()
}

func (o TrafficForwardingGRETunnelMapOutput) ToTrafficForwardingGRETunnelMapOutput() TrafficForwardingGRETunnelMapOutput {
	return o
}

func (o TrafficForwardingGRETunnelMapOutput) ToTrafficForwardingGRETunnelMapOutputWithContext(ctx context.Context) TrafficForwardingGRETunnelMapOutput {
	return o
}

func (o TrafficForwardingGRETunnelMapOutput) MapIndex(k pulumi.StringInput) TrafficForwardingGRETunnelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TrafficForwardingGRETunnel {
		return vs[0].(map[string]*TrafficForwardingGRETunnel)[vs[1].(string)]
	}).(TrafficForwardingGRETunnelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficForwardingGRETunnelInput)(nil)).Elem(), &TrafficForwardingGRETunnel{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficForwardingGRETunnelArrayInput)(nil)).Elem(), TrafficForwardingGRETunnelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficForwardingGRETunnelMapInput)(nil)).Elem(), TrafficForwardingGRETunnelMap{})
	pulumi.RegisterOutputType(TrafficForwardingGRETunnelOutput{})
	pulumi.RegisterOutputType(TrafficForwardingGRETunnelArrayOutput{})
	pulumi.RegisterOutputType(TrafficForwardingGRETunnelMapOutput{})
}
