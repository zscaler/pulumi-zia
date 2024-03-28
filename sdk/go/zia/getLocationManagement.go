// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zia

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zia/sdk/go/zia/internal"
)

// Use the **zia_location_management** data source to get information about a location resource available in the Zscaler Internet Access Location Management. This resource can then be referenced in multiple other resources, such as URL Filtering Rules, Firewall rules etc.
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
//			_, err := zia.LookupLocationManagement(ctx, &zia.LookupLocationManagementArgs{
//				Name: pulumi.StringRef("San Jose"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupLocationManagement(ctx *pulumi.Context, args *LookupLocationManagementArgs, opts ...pulumi.InvokeOption) (*LookupLocationManagementResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLocationManagementResult
	err := ctx.Invoke("zia:index/getLocationManagement:getLocationManagement", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocationManagement.
type LookupLocationManagementArgs struct {
	BasicAuthEnabled *bool `pulumi:"basicAuthEnabled"`
	// The ID of the location to be exported.
	Id *int `pulumi:"id"`
	// The name of the location to be exported.
	Name       *string `pulumi:"name"`
	ParentName *string `pulumi:"parentName"`
}

// A collection of values returned by getLocationManagement.
type LookupLocationManagementResult struct {
	// (Boolean) For First Time AUP Behavior, Block Internet Access. When set, all internet access (including non-HTTP traffic) is disabled until the user accepts the AUP.
	AupBlockInternetUntilAccepted bool `pulumi:"aupBlockInternetUntilAccepted"`
	// (Boolean) Enable AUP. When set to true, AUP is enabled for the location.
	AupEnabled bool `pulumi:"aupEnabled"`
	// (Boolean) For First Time AUP Behavior, Force SSL Inspection. When set, Zscaler will force SSL Inspection in order to enforce AUP for HTTPS traffic.
	AupForceSslInspection bool `pulumi:"aupForceSslInspection"`
	// (Number) Custom AUP Frequency. Refresh time (in days) to re-validate the AUP.
	AupTimeoutInDays int `pulumi:"aupTimeoutInDays"`
	// (Boolean) Enforce Authentication. Required when ports are enabled, IP Surrogate is enabled, or Kerberos Authentication is enabled.
	AuthRequired     bool `pulumi:"authRequired"`
	BasicAuthEnabled bool `pulumi:"basicAuthEnabled"`
	// (Boolean) Enable Caution. When set to true, a caution notifcation is enabled for the location.
	CautionEnabled bool `pulumi:"cautionEnabled"`
	// (String) Country
	Country string `pulumi:"country"`
	// (String) Additional notes or information regarding the location or sub-location. The description cannot exceed 1024 characters.
	Description       string `pulumi:"description"`
	DigestAuthEnabled bool   `pulumi:"digestAuthEnabled"`
	// (String) Display Time Unit. The time unit to display for IP Surrogate idle time to disassociation.
	DisplayTimeUnit string `pulumi:"displayTimeUnit"`
	// (Number) Download bandwidth in bytes. The value `0` implies no Bandwidth Control enforcement.
	DnBandwidth int `pulumi:"dnBandwidth"`
	// (Number) Identifier that uniquely identifies an entity
	Id *int `pulumi:"id"`
	// (Number) Idle Time to Disassociation. The user mapping idle time (in minutes) is required if a Surrogate IP is enabled.
	IdleTimeInMinutes   int  `pulumi:"idleTimeInMinutes"`
	IotDiscoveryEnabled bool `pulumi:"iotDiscoveryEnabled"`
	// (List of String) For locations: IP addresses of the egress points that are provisioned in the Zscaler Cloud. Each entry is a single IP address (e.g., `238.10.33.9`). For sub-locations: Egress, internal, or GRE tunnel IP addresses. Each entry is either a single IP address, CIDR (e.g., `10.10.33.0/24`), or range (e.g., `10.10.33.1-10.10.33.10`)).
	IpAddresses []string `pulumi:"ipAddresses"`
	// (Boolean) Enable IPS Control. When set to true, IPS Control is enabled for the location if Firewall is enabled.
	IpsControl          bool `pulumi:"ipsControl"`
	KerberosAuthEnabled bool `pulumi:"kerberosAuthEnabled"`
	// (String) The configured name of the entity
	Name *string `pulumi:"name"`
	// (Boolean) Enable Firewall. When set to true, Firewall is enabled for the location.
	OfwEnabled bool `pulumi:"ofwEnabled"`
	// (Number) - Parent Location ID. If this ID does not exist or is `0`, it is implied that it is a parent location. Otherwise, it is a sub-location whose parent has this ID. x-applicableTo: `SUB`
	ParentId   int     `pulumi:"parentId"`
	ParentName *string `pulumi:"parentName"`
	// (String) IP ports that are associated with the location.
	Ports string `pulumi:"ports"`
	// (String) Profile tag that specifies the location traffic type. If not specified, this tag defaults to `Unassigned`.
	Profile string `pulumi:"profile"`
	// (Boolean) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.
	SslScanEnabled bool `pulumi:"sslScanEnabled"`
	// (Boolean) Enable Surrogate IP. When set to true, users are mapped to internal device IP addresses.
	SurrogateIp bool `pulumi:"surrogateIp"`
	// (Boolean) Enforce Surrogate IP for Known Browsers. When set to true, IP Surrogate is enforced for all known browsers.
	SurrogateIpEnforcedForKnownBrowsers bool `pulumi:"surrogateIpEnforcedForKnownBrowsers"`
	// (Number) Refresh Time for re-validation of Surrogacy. The surrogate refresh time (in minutes) to re-validate the IP surrogates.
	SurrogateRefreshTimeInMinutes int `pulumi:"surrogateRefreshTimeInMinutes"`
	// (String) Display Refresh Time Unit. The time unit to display for refresh time for re-validation of surrogacy.
	SurrogateRefreshTimeUnit string `pulumi:"surrogateRefreshTimeUnit"`
	// (String) Timezone of the location. If not specified, it defaults to GMT.
	Tz string `pulumi:"tz"`
	// (Number) Upload bandwidth in bytes. The value `0` implies no Bandwidth Control enforcement.
	UpBandwidth    int                                  `pulumi:"upBandwidth"`
	VpnCredentials []GetLocationManagementVpnCredential `pulumi:"vpnCredentials"`
	// (Boolean) Enable XFF Forwarding. When set to true, traffic is passed to Zscaler Cloud via the X-Forwarded-For (XFF) header.
	XffForwardEnabled bool `pulumi:"xffForwardEnabled"`
	// (Boolean) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.
	ZappSslScanEnabled bool `pulumi:"zappSslScanEnabled"`
}

func LookupLocationManagementOutput(ctx *pulumi.Context, args LookupLocationManagementOutputArgs, opts ...pulumi.InvokeOption) LookupLocationManagementResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocationManagementResult, error) {
			args := v.(LookupLocationManagementArgs)
			r, err := LookupLocationManagement(ctx, &args, opts...)
			var s LookupLocationManagementResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocationManagementResultOutput)
}

// A collection of arguments for invoking getLocationManagement.
type LookupLocationManagementOutputArgs struct {
	BasicAuthEnabled pulumi.BoolPtrInput `pulumi:"basicAuthEnabled"`
	// The ID of the location to be exported.
	Id pulumi.IntPtrInput `pulumi:"id"`
	// The name of the location to be exported.
	Name       pulumi.StringPtrInput `pulumi:"name"`
	ParentName pulumi.StringPtrInput `pulumi:"parentName"`
}

func (LookupLocationManagementOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationManagementArgs)(nil)).Elem()
}

// A collection of values returned by getLocationManagement.
type LookupLocationManagementResultOutput struct{ *pulumi.OutputState }

func (LookupLocationManagementResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationManagementResult)(nil)).Elem()
}

func (o LookupLocationManagementResultOutput) ToLookupLocationManagementResultOutput() LookupLocationManagementResultOutput {
	return o
}

func (o LookupLocationManagementResultOutput) ToLookupLocationManagementResultOutputWithContext(ctx context.Context) LookupLocationManagementResultOutput {
	return o
}

// (Boolean) For First Time AUP Behavior, Block Internet Access. When set, all internet access (including non-HTTP traffic) is disabled until the user accepts the AUP.
func (o LookupLocationManagementResultOutput) AupBlockInternetUntilAccepted() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) bool { return v.AupBlockInternetUntilAccepted }).(pulumi.BoolOutput)
}

// (Boolean) Enable AUP. When set to true, AUP is enabled for the location.
func (o LookupLocationManagementResultOutput) AupEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) bool { return v.AupEnabled }).(pulumi.BoolOutput)
}

// (Boolean) For First Time AUP Behavior, Force SSL Inspection. When set, Zscaler will force SSL Inspection in order to enforce AUP for HTTPS traffic.
func (o LookupLocationManagementResultOutput) AupForceSslInspection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) bool { return v.AupForceSslInspection }).(pulumi.BoolOutput)
}

// (Number) Custom AUP Frequency. Refresh time (in days) to re-validate the AUP.
func (o LookupLocationManagementResultOutput) AupTimeoutInDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) int { return v.AupTimeoutInDays }).(pulumi.IntOutput)
}

// (Boolean) Enforce Authentication. Required when ports are enabled, IP Surrogate is enabled, or Kerberos Authentication is enabled.
func (o LookupLocationManagementResultOutput) AuthRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) bool { return v.AuthRequired }).(pulumi.BoolOutput)
}

func (o LookupLocationManagementResultOutput) BasicAuthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) bool { return v.BasicAuthEnabled }).(pulumi.BoolOutput)
}

// (Boolean) Enable Caution. When set to true, a caution notifcation is enabled for the location.
func (o LookupLocationManagementResultOutput) CautionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) bool { return v.CautionEnabled }).(pulumi.BoolOutput)
}

// (String) Country
func (o LookupLocationManagementResultOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) string { return v.Country }).(pulumi.StringOutput)
}

// (String) Additional notes or information regarding the location or sub-location. The description cannot exceed 1024 characters.
func (o LookupLocationManagementResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupLocationManagementResultOutput) DigestAuthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) bool { return v.DigestAuthEnabled }).(pulumi.BoolOutput)
}

// (String) Display Time Unit. The time unit to display for IP Surrogate idle time to disassociation.
func (o LookupLocationManagementResultOutput) DisplayTimeUnit() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) string { return v.DisplayTimeUnit }).(pulumi.StringOutput)
}

// (Number) Download bandwidth in bytes. The value `0` implies no Bandwidth Control enforcement.
func (o LookupLocationManagementResultOutput) DnBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) int { return v.DnBandwidth }).(pulumi.IntOutput)
}

// (Number) Identifier that uniquely identifies an entity
func (o LookupLocationManagementResultOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) *int { return v.Id }).(pulumi.IntPtrOutput)
}

// (Number) Idle Time to Disassociation. The user mapping idle time (in minutes) is required if a Surrogate IP is enabled.
func (o LookupLocationManagementResultOutput) IdleTimeInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) int { return v.IdleTimeInMinutes }).(pulumi.IntOutput)
}

func (o LookupLocationManagementResultOutput) IotDiscoveryEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) bool { return v.IotDiscoveryEnabled }).(pulumi.BoolOutput)
}

// (List of String) For locations: IP addresses of the egress points that are provisioned in the Zscaler Cloud. Each entry is a single IP address (e.g., `238.10.33.9`). For sub-locations: Egress, internal, or GRE tunnel IP addresses. Each entry is either a single IP address, CIDR (e.g., `10.10.33.0/24`), or range (e.g., `10.10.33.1-10.10.33.10`)).
func (o LookupLocationManagementResultOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

// (Boolean) Enable IPS Control. When set to true, IPS Control is enabled for the location if Firewall is enabled.
func (o LookupLocationManagementResultOutput) IpsControl() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) bool { return v.IpsControl }).(pulumi.BoolOutput)
}

func (o LookupLocationManagementResultOutput) KerberosAuthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) bool { return v.KerberosAuthEnabled }).(pulumi.BoolOutput)
}

// (String) The configured name of the entity
func (o LookupLocationManagementResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// (Boolean) Enable Firewall. When set to true, Firewall is enabled for the location.
func (o LookupLocationManagementResultOutput) OfwEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) bool { return v.OfwEnabled }).(pulumi.BoolOutput)
}

// (Number) - Parent Location ID. If this ID does not exist or is `0`, it is implied that it is a parent location. Otherwise, it is a sub-location whose parent has this ID. x-applicableTo: `SUB`
func (o LookupLocationManagementResultOutput) ParentId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) int { return v.ParentId }).(pulumi.IntOutput)
}

func (o LookupLocationManagementResultOutput) ParentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) *string { return v.ParentName }).(pulumi.StringPtrOutput)
}

// (String) IP ports that are associated with the location.
func (o LookupLocationManagementResultOutput) Ports() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) string { return v.Ports }).(pulumi.StringOutput)
}

// (String) Profile tag that specifies the location traffic type. If not specified, this tag defaults to `Unassigned`.
func (o LookupLocationManagementResultOutput) Profile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) string { return v.Profile }).(pulumi.StringOutput)
}

// (Boolean) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.
func (o LookupLocationManagementResultOutput) SslScanEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) bool { return v.SslScanEnabled }).(pulumi.BoolOutput)
}

// (Boolean) Enable Surrogate IP. When set to true, users are mapped to internal device IP addresses.
func (o LookupLocationManagementResultOutput) SurrogateIp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) bool { return v.SurrogateIp }).(pulumi.BoolOutput)
}

// (Boolean) Enforce Surrogate IP for Known Browsers. When set to true, IP Surrogate is enforced for all known browsers.
func (o LookupLocationManagementResultOutput) SurrogateIpEnforcedForKnownBrowsers() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) bool { return v.SurrogateIpEnforcedForKnownBrowsers }).(pulumi.BoolOutput)
}

// (Number) Refresh Time for re-validation of Surrogacy. The surrogate refresh time (in minutes) to re-validate the IP surrogates.
func (o LookupLocationManagementResultOutput) SurrogateRefreshTimeInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) int { return v.SurrogateRefreshTimeInMinutes }).(pulumi.IntOutput)
}

// (String) Display Refresh Time Unit. The time unit to display for refresh time for re-validation of surrogacy.
func (o LookupLocationManagementResultOutput) SurrogateRefreshTimeUnit() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) string { return v.SurrogateRefreshTimeUnit }).(pulumi.StringOutput)
}

// (String) Timezone of the location. If not specified, it defaults to GMT.
func (o LookupLocationManagementResultOutput) Tz() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) string { return v.Tz }).(pulumi.StringOutput)
}

// (Number) Upload bandwidth in bytes. The value `0` implies no Bandwidth Control enforcement.
func (o LookupLocationManagementResultOutput) UpBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) int { return v.UpBandwidth }).(pulumi.IntOutput)
}

func (o LookupLocationManagementResultOutput) VpnCredentials() GetLocationManagementVpnCredentialArrayOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) []GetLocationManagementVpnCredential { return v.VpnCredentials }).(GetLocationManagementVpnCredentialArrayOutput)
}

// (Boolean) Enable XFF Forwarding. When set to true, traffic is passed to Zscaler Cloud via the X-Forwarded-For (XFF) header.
func (o LookupLocationManagementResultOutput) XffForwardEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) bool { return v.XffForwardEnabled }).(pulumi.BoolOutput)
}

// (Boolean) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.
func (o LookupLocationManagementResultOutput) ZappSslScanEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLocationManagementResult) bool { return v.ZappSslScanEnabled }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocationManagementResultOutput{})
}
