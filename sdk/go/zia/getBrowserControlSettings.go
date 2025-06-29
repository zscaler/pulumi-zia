// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zia

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zia/sdk/go/zia/internal"
)

// * [Official documentation](https://help.zscaler.com/zia/configuring-browser-control-policy)
// * [API documentation](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
//
// Use the **zia_browser_control_policy** data source to retrieves information about the security exceptions configured for the Malware Protection policy. To learn more see [Configuring the Browser Control Policy](https://help.zscaler.com/zia/configuring-browser-control-policy)
//
// ## Example Usage
func LookupBrowserControlSettings(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*LookupBrowserControlSettingsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBrowserControlSettingsResult
	err := ctx.Invoke("zia:index/getBrowserControlSettings:getBrowserControlSettings", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getBrowserControlSettings.
type LookupBrowserControlSettingsResult struct {
	// (Boolean) A Boolean value that specifies whether or not to allow all the browsers and their respective versions access to the internet
	AllowAllBrowsers bool `pulumi:"allowAllBrowsers"`
	// (List) Versions of Google Chrome browser that need to be blocked. If not set, all Google Chrome versions are allowed. See all [Supported values](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
	BlockedChromeVersions []string `pulumi:"blockedChromeVersions"`
	// (List) Versions of Mozilla Firefox browser that need to be blocked. If not set, all Mozilla Firefox versions are allowed. See all [Supported values](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
	BlockedFirefoxVersions []string `pulumi:"blockedFirefoxVersions"`
	// (List) Versions of Microsoft browser that need to be blocked. If not set, all Microsoft browser versions are allowed. See all [Supported values](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
	BlockedInternetExplorerVersions []string `pulumi:"blockedInternetExplorerVersions"`
	// (List) Versions of Opera browser that need to be blocked. If not set, all Opera versions are allowed. See all [Supported values](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
	BlockedOperaVersions []string `pulumi:"blockedOperaVersions"`
	// (List) Versions of Apple Safari browser that need to be blocked. If not set, all Apple Safari versions are allowed. See all [Supported values](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
	BlockedSafariVersions []string `pulumi:"blockedSafariVersions"`
	// (Boolean) If set to true, all the browsers are bypassed for warnings
	BypassAllBrowsers bool `pulumi:"bypassAllBrowsers"`
	// (List) List of applications that need to be bypassed for warnings. This attribute has effect only if the 'enableWarnings' attribute is set to true. If not set, all vulnerable applications are warned. Supported Values:
	// * `ANY`
	// * `NONE`
	// * `OUTLOOKEXP`
	// * `MSOFFICE`
	BypassApplications []string `pulumi:"bypassApplications"`
	// (List) List of plugins that need to be bypassed for warnings. This attribute has effect only if the 'enableWarnings' attribute is set to true. If not set, all vulnerable plugins are warned.Supported Values:
	// * `ANY`
	// * `NONE`
	// * `ACROBAT`
	// * `FLASH`
	// * `SHOCKWAVE`
	// * `QUICKTIME`
	// * `DIVX`
	// * `GOOGLEGEARS`
	// * `DOTNET`
	// * `SILVERLIGHT`
	// * `REALPLAYER`
	// * `JAVA`
	// * `TOTEM`
	// * `WMP`
	BypassPlugins []string `pulumi:"bypassPlugins"`
	// (Boolean) A Boolean value that specifies if Smart Browser Isolation is enabled
	EnableSmartBrowserIsolation bool `pulumi:"enableSmartBrowserIsolation"`
	// (Boolean) A Boolean value that specifies if the warnings are enabled
	EnableWarnings bool `pulumi:"enableWarnings"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (String) Specifies how frequently the service checks browsers and relevant applications to warn users regarding outdated or vulnerable browsers, plugins, and applications. If not set, the warnings are disabled. Supported Values:
	// * `DAILY`
	// * `WEEKLY`
	// * `MONTHLY`,
	// * `EVERY_2_HOURS`
	// * `EVERY_4_HOURS`
	// * `EVERY_6_HOURS`
	// * `EVERY_8_HOURS`
	// * `EVERY_12_HOURS`
	PluginCheckFrequency    string `pulumi:"pluginCheckFrequency"`
	SmartIsolationProfileId int    `pulumi:"smartIsolationProfileId"`
	// (Block, Max: 1) The isolation profile ID used for DLP email alerts sent to the auditor.
	SmartIsolationProfiles []GetBrowserControlSettingsSmartIsolationProfile `pulumi:"smartIsolationProfiles"`
}

func LookupBrowserControlSettingsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) LookupBrowserControlSettingsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (LookupBrowserControlSettingsResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("zia:index/getBrowserControlSettings:getBrowserControlSettings", nil, LookupBrowserControlSettingsResultOutput{}, options).(LookupBrowserControlSettingsResultOutput), nil
	}).(LookupBrowserControlSettingsResultOutput)
}

// A collection of values returned by getBrowserControlSettings.
type LookupBrowserControlSettingsResultOutput struct{ *pulumi.OutputState }

func (LookupBrowserControlSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBrowserControlSettingsResult)(nil)).Elem()
}

func (o LookupBrowserControlSettingsResultOutput) ToLookupBrowserControlSettingsResultOutput() LookupBrowserControlSettingsResultOutput {
	return o
}

func (o LookupBrowserControlSettingsResultOutput) ToLookupBrowserControlSettingsResultOutputWithContext(ctx context.Context) LookupBrowserControlSettingsResultOutput {
	return o
}

// (Boolean) A Boolean value that specifies whether or not to allow all the browsers and their respective versions access to the internet
func (o LookupBrowserControlSettingsResultOutput) AllowAllBrowsers() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBrowserControlSettingsResult) bool { return v.AllowAllBrowsers }).(pulumi.BoolOutput)
}

// (List) Versions of Google Chrome browser that need to be blocked. If not set, all Google Chrome versions are allowed. See all [Supported values](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
func (o LookupBrowserControlSettingsResultOutput) BlockedChromeVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBrowserControlSettingsResult) []string { return v.BlockedChromeVersions }).(pulumi.StringArrayOutput)
}

// (List) Versions of Mozilla Firefox browser that need to be blocked. If not set, all Mozilla Firefox versions are allowed. See all [Supported values](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
func (o LookupBrowserControlSettingsResultOutput) BlockedFirefoxVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBrowserControlSettingsResult) []string { return v.BlockedFirefoxVersions }).(pulumi.StringArrayOutput)
}

// (List) Versions of Microsoft browser that need to be blocked. If not set, all Microsoft browser versions are allowed. See all [Supported values](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
func (o LookupBrowserControlSettingsResultOutput) BlockedInternetExplorerVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBrowserControlSettingsResult) []string { return v.BlockedInternetExplorerVersions }).(pulumi.StringArrayOutput)
}

// (List) Versions of Opera browser that need to be blocked. If not set, all Opera versions are allowed. See all [Supported values](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
func (o LookupBrowserControlSettingsResultOutput) BlockedOperaVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBrowserControlSettingsResult) []string { return v.BlockedOperaVersions }).(pulumi.StringArrayOutput)
}

// (List) Versions of Apple Safari browser that need to be blocked. If not set, all Apple Safari versions are allowed. See all [Supported values](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
func (o LookupBrowserControlSettingsResultOutput) BlockedSafariVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBrowserControlSettingsResult) []string { return v.BlockedSafariVersions }).(pulumi.StringArrayOutput)
}

// (Boolean) If set to true, all the browsers are bypassed for warnings
func (o LookupBrowserControlSettingsResultOutput) BypassAllBrowsers() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBrowserControlSettingsResult) bool { return v.BypassAllBrowsers }).(pulumi.BoolOutput)
}

// (List) List of applications that need to be bypassed for warnings. This attribute has effect only if the 'enableWarnings' attribute is set to true. If not set, all vulnerable applications are warned. Supported Values:
// * `ANY`
// * `NONE`
// * `OUTLOOKEXP`
// * `MSOFFICE`
func (o LookupBrowserControlSettingsResultOutput) BypassApplications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBrowserControlSettingsResult) []string { return v.BypassApplications }).(pulumi.StringArrayOutput)
}

// (List) List of plugins that need to be bypassed for warnings. This attribute has effect only if the 'enableWarnings' attribute is set to true. If not set, all vulnerable plugins are warned.Supported Values:
// * `ANY`
// * `NONE`
// * `ACROBAT`
// * `FLASH`
// * `SHOCKWAVE`
// * `QUICKTIME`
// * `DIVX`
// * `GOOGLEGEARS`
// * `DOTNET`
// * `SILVERLIGHT`
// * `REALPLAYER`
// * `JAVA`
// * `TOTEM`
// * `WMP`
func (o LookupBrowserControlSettingsResultOutput) BypassPlugins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBrowserControlSettingsResult) []string { return v.BypassPlugins }).(pulumi.StringArrayOutput)
}

// (Boolean) A Boolean value that specifies if Smart Browser Isolation is enabled
func (o LookupBrowserControlSettingsResultOutput) EnableSmartBrowserIsolation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBrowserControlSettingsResult) bool { return v.EnableSmartBrowserIsolation }).(pulumi.BoolOutput)
}

// (Boolean) A Boolean value that specifies if the warnings are enabled
func (o LookupBrowserControlSettingsResultOutput) EnableWarnings() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBrowserControlSettingsResult) bool { return v.EnableWarnings }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupBrowserControlSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBrowserControlSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

// (String) Specifies how frequently the service checks browsers and relevant applications to warn users regarding outdated or vulnerable browsers, plugins, and applications. If not set, the warnings are disabled. Supported Values:
// * `DAILY`
// * `WEEKLY`
// * `MONTHLY`,
// * `EVERY_2_HOURS`
// * `EVERY_4_HOURS`
// * `EVERY_6_HOURS`
// * `EVERY_8_HOURS`
// * `EVERY_12_HOURS`
func (o LookupBrowserControlSettingsResultOutput) PluginCheckFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBrowserControlSettingsResult) string { return v.PluginCheckFrequency }).(pulumi.StringOutput)
}

func (o LookupBrowserControlSettingsResultOutput) SmartIsolationProfileId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBrowserControlSettingsResult) int { return v.SmartIsolationProfileId }).(pulumi.IntOutput)
}

// (Block, Max: 1) The isolation profile ID used for DLP email alerts sent to the auditor.
func (o LookupBrowserControlSettingsResultOutput) SmartIsolationProfiles() GetBrowserControlSettingsSmartIsolationProfileArrayOutput {
	return o.ApplyT(func(v LookupBrowserControlSettingsResult) []GetBrowserControlSettingsSmartIsolationProfile {
		return v.SmartIsolationProfiles
	}).(GetBrowserControlSettingsSmartIsolationProfileArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBrowserControlSettingsResultOutput{})
}
