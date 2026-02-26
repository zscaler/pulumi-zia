// Copyright (c) 2023 Zscaler Technology Alliances, <devrel@zscaler.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package provider implements the native Pulumi provider for Zscaler Internet Access (ZIA).
package provider

import (
	"fmt"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// Name controls how this provider is referenced in package names and elsewhere.
const Name string = "zia"

// SDK package names match the bridged provider (pulumi-zia) for registry compatibility.
// When replacing the bridged provider, PyPI/npm/NuGet names remain unchanged.
const (
	pythonPackageName = "zscaler_pulumi_zia"
	nodeJSPackageName = "@bdzscaler/pulumi-zia"
)

// Provider creates a new instance of the ZIA provider.
func Provider() p.Provider {
	prov, err := infer.NewProviderBuilder().
		WithDisplayName("pulumi-resource-zia").
		WithDescription("A native Pulumi provider for Zscaler Internet Access (ZIA).").
		WithHomepage("https://github.com/zscaler/pulumi-zia").
		WithNamespace("zia").
		WithLanguageMap(map[string]any{
			"go": map[string]any{
				"generateResourceContainerTypes": true,
				"importBasePath":                 "github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia",
				"respectSchemaVersion":           true,
			},
			"nodejs": map[string]any{
				"packageName":          nodeJSPackageName,
				"respectSchemaVersion": true,
			},
			"python": map[string]any{
				"packageName":          pythonPackageName,
				"respectSchemaVersion": true,
				"pyproject": map[string]any{
					"enabled": true,
				},
			},
			"csharp": map[string]any{
				"respectSchemaVersion": true,
				"rootNamespace":        "zscaler.PulumiPackage",
			},
		}).
		WithResources(
			infer.Resource(Activation{}),
			infer.Resource(AdminRoles{}),
			infer.Resource(AdminUsers{}),
			infer.Resource(AdvancedSettings{}),
			infer.Resource(AtpSettings{}),
			infer.Resource(AuthSettingsUrls{}),
			infer.Resource(AtpMaliciousUrls{}),
			infer.Resource(AtpSecurityExceptions{}),
			infer.Resource(AtpMalwareInspection{}),
			infer.Resource(AtpMalwarePolicy{}),
			infer.Resource(AtpMalwareProtocols{}),
			infer.Resource(AtpMalwareSettings{}),
			infer.Resource(SubscriptionAlert{}),
			infer.Resource(RuleLabel{}),
			infer.Resource(NatControlRule{}),
			infer.Resource(ForwardingControlRule{}),
			infer.Resource(ForwardingControlZpaGateway{}),
			infer.Resource(ForwardingControlProxies{}),
			infer.Resource(FirewallIPSRule{}),
			infer.Resource(URLFilteringRule{}),
			infer.Resource(FirewallFilteringRule{}),
			infer.Resource(FirewallDNSRule{}),
			infer.Resource(FileTypeControlRule{}),
			infer.Resource(DlpWebRule{}),
			infer.Resource(CloudAppControlRule{}),
			infer.Resource(CasbDlpRule{}),
			infer.Resource(CasbMalwareRule{}),
			infer.Resource(BandwidthControlRule{}),
			infer.Resource(BandwidthClass{}),
			infer.Resource(BandwidthClassFileSize{}),
			infer.Resource(BandwidthClassWebConferencing{}),
			infer.Resource(BrowserControlPolicy{}),
			infer.Resource(CloudApplicationInstance{}),
			infer.Resource(CloudNssFeed{}),
			infer.Resource(FwNetworkService{}),
			infer.Resource(FwNetworkServiceGroup{}),
			infer.Resource(FwNetworkApplicationGroup{}),
			infer.Resource(FwIpSourceGroup{}),
			infer.Resource(FwIpDestinationGroup{}),
			infer.Resource(FtpControlPolicy{}),
			infer.Resource(LocationManagement{}),
			infer.Resource(CustomFileType{}),
			infer.Resource(DcExclusion{}),
			infer.Resource(DlpDictionary{}),
			infer.Resource(DlpEngine{}),
			infer.Resource(DlpNotificationTemplate{}),
			infer.Resource(EndUserNotification{}),
			infer.Resource(Extranet{}),
			infer.Resource(UrlFilteringCloudAppSettings{}),
			infer.Resource(UrlCategory{}),
			infer.Resource(UrlCategoryPredefined{}),
			infer.Resource(WorkloadGroup{}),
			infer.Resource(VzenNode{}),
			infer.Resource(VzenCluster{}),
			infer.Resource(UserManagementUser{}),
			infer.Resource(TrafficForwardingVpnCredentials{}),
			infer.Resource(TrafficForwardingStaticIp{}),
			infer.Resource(TrafficForwardingGreTunnel{}),
			infer.Resource(SubCloud{}),
			infer.Resource(TrafficCaptureRule{}),
			infer.Resource(SslInspectionRule{}),
			infer.Resource(TenantRestrictionProfile{}),
			infer.Resource(SecurityPolicySettings{}),
			infer.Resource(SandboxSubmission{}),
			infer.Resource(SandboxRule{}),
			infer.Resource(SandboxBehavioralAnalysisAdvancedSettings{}),
			infer.Resource(SandboxBehavioralAnalysisAdvancedSettingsV2{}),
			infer.Resource(MobileMalwareProtectionPolicy{}),
			infer.Resource(RiskProfile{}),
			infer.Resource(NssServer{}),
		).
		WithFunctions(
			infer.Function(&GetWorkloadGroup{}),
			infer.Function(&GetVzenNode{}),
			infer.Function(&GetVzenCluster{}),
			infer.Function(&GetUserManagementUser{}),
			infer.Function(&GetLocationGroup{}),
			infer.Function(&GetDlpIdmProfile{}),
			infer.Function(&GetDlpIcapServer{}),
			infer.Function(&GetFileTypeCategories{}),
			infer.Function(&GetDlpEdmSchema{}),
			infer.Function(&GetDevice{}),
			infer.Function(&GetDeviceGroup{}),
			infer.Function(&GetCloudBrowserIsolationProfile{}),
			infer.Function(&GetCloudApplications{}),
			infer.Function(&GetCasbTombstoneTemplate{}),
			infer.Function(&GetCasbTenant{}),
			infer.Function(&GetCasbEmailLabel{}),
			infer.Function(&GetDlpCloudToCloudIr{}),
			infer.Function(&GetDlpDictionaryPredefinedIdentifiers{}),
			infer.Function(&GetDlpIdmProfileLite{}),
			infer.Function(&GetDlpIncidentReceiverServer{}),
			infer.Function(&GetDomainProfile{}),
			infer.Function(&GetSandboxReport{}),
			infer.Function(&GetUserManagementDepartment{}),
			infer.Function(&GetUserManagementGroup{}),
			infer.Function(&GetTimeWindow{}),
			infer.Function(&GetFwNetworkService{}),
			infer.Function(&GetDatacenters{}),
		).
		WithConfig(infer.Config(&Config{})).
		WithModuleMap(map[tokens.ModuleName]tokens.ModuleName{
			"zia":      "index",
			"provider": "index", // Avoid C# collision: Provider class vs Provider namespace
		}).
		Build()
	if err != nil {
		panic(fmt.Errorf("unable to build ZIA provider: %w", err))
	}
	return prov
}
