// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zia

import (
	"fmt"
	"unicode"

	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"
	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tks "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/zscaler/pulumi-zia/provider/pkg/version"
	"github.com/zscaler/terraform-provider-zia/v4/zia"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	ziaPkg = "zia"
	// modules:
	ziaMod = "index"
)

// ziaMember manufactures a type token for the zia package and the given module and type.
func ziaMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(ziaPkg + ":" + mod + ":" + mem)
}

// ziaType manufactures a type token for the zia package and the given module and type.
func ziaType(mod string, typ string) tokens.Type {
	return tokens.Type(ziaMember(mod, typ))
}

// ziaDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the zia package and names the file by simply lower casing the data
// source's first character.
func ziaDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return ziaMember(mod+"/"+fn, res)
}

// ziaResource manufactures a standard resource token given a module and resource name.
// It automatically uses the zia package and names the file by simply lower casing the resource's
// first character.
func ziaResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return ziaType(mod+"/"+fn, res)
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

//go:embed cmd/pulumi-resource-zia/bridge-metadata.json
var metadata []byte

func refProviderLicense(license tfbridge.TFProviderLicense) *tfbridge.TFProviderLicense {
	return &license
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(zia.ZIAProvider())
	prov := tfbridge.ProviderInfo{
		P:                       p,
		Name:                    "zia",
		Description:             "A Pulumi package for creating and managing zia cloud resources.",
		Keywords:                []string{"pulumi", "zia", "zscaler", "category/cloud"},
		TFProviderLicense:       refProviderLicense(tfbridge.MITLicenseType),
		License:                 "MIT",
		LogoURL:                 "https://raw.githubusercontent.com/zscaler/pulumi-zia/master/assets/zscaler.png", // nolint[:lll]
		Homepage:                "https://www.zscaler.com",
		Repository:              "https://github.com/zscaler/pulumi-zia",
		PluginDownloadURL:       "github://api.github.com/zscaler",
		GitHubOrg:               "zscaler",
		Publisher:               "Zscaler",
		DisplayName:             "Zscaler Internet Access",
		TFProviderModuleVersion: "v4",
		Version:                 version.Version,
		Config: map[string]*tfbridge.SchemaInfo{
			// Authentication Parameters for OneAPI Mode
			"client_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZSCALER_CLIENT_ID"},
				},
			},
			"client_secret": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZSCALER_CLIENT_SECRET"},
				},
			},
			"private_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZSCALER_PRIVATE_KEY"},
				},
			},
			"vanity_domain": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZSCALER_VANITY_DOMAIN"},
				},
			},
			"zscaler_cloud": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZSCALER_CLOUD"},
				},
			},
			"username": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZIA_USERNAME"},
				},
			},
			"password": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZIA_PASSWORD"},
				},
			},
			"api_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZIA_API_KEY"},
				},
			},
			"zia_cloud": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZIA_CLOUD"},
				},
			},
			"sandbox_token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZSCALER_SANDBOX_TOKEN"},
				},
			},
			"sandbox_cloud": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZSCALER_SANDBOX_CLOUD"},
				},
			},
			"use_legacy_client": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZSCALER_USE_LEGACY_CLIENT"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"zia_activation_status":          {Tok: ziaResource(ziaMod, "ActivationStatus")},
			"zia_advanced_settings":          {Tok: ziaResource(ziaMod, "AdvancedSettings")},
			"zia_advanced_threat_settings":   {Tok: ziaResource(ziaMod, "AdvancedThreatSettings")},
			"zia_atp_malicious_urls":         {Tok: ziaResource(ziaMod, "ATPMaliciousURLs")},
			"zia_atp_malware_inspection":     {Tok: ziaResource(ziaMod, "ATPMalwareInspection")},
			"zia_atp_malware_policy":         {Tok: ziaResource(ziaMod, "ATPMalwarePolicy")},
			"zia_atp_malware_protocols":      {Tok: ziaResource(ziaMod, "ATPMalwareProtocols")},
			"zia_atp_malware_settings":       {Tok: ziaResource(ziaMod, "ATPMalwareSettings")},
			"zia_atp_security_exceptions":    {Tok: ziaResource(ziaMod, "ATPSecurityExceptions")},
			"zia_cloud_app_control_rule":     {Tok: ziaResource(ziaMod, "CloudAppControlRule")},
			"zia_admin_users":                {Tok: ziaResource(ziaMod, "AdminUsers")},
			"zia_dlp_dictionaries":           {Tok: ziaResource(ziaMod, "DLPDictionaries")},
			"zia_dlp_engines":                {Tok: ziaResource(ziaMod, "DLPEngines")},
			"zia_dlp_notification_templates": {Tok: ziaResource(ziaMod, "DLPNotificationTemplates")},
			"zia_dlp_web_rules":              {Tok: ziaResource(ziaMod, "DLPWebRules")},
			"zia_end_user_notification":      {Tok: ziaResource(ziaMod, "EndUserNotification")},
			"zia_file_type_control_rules": {
				Tok: ziaResource(ziaMod, "FileTypeControlRules"),
				// Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zia_firewall_dns_rule": {
				Tok: ziaResource(ziaMod, "FirewallDNSRule"),
				// Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zia_firewall_ips_rule": {
				Tok: ziaResource(ziaMod, "IPSFirewallRule"),
				// Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zia_firewall_filtering_rule":                       {Tok: ziaResource(ziaMod, "FirewallFilteringRule")},
			"zia_firewall_filtering_destination_groups":         {Tok: ziaResource(ziaMod, "FirewallFilteringDestinationGroups")},
			"zia_firewall_filtering_ip_source_groups":           {Tok: ziaResource(ziaMod, "FirewallFilteringSourceGroups")},
			"zia_firewall_filtering_network_service":            {Tok: ziaResource(ziaMod, "FirewallFilteringNetworkServices")},
			"zia_firewall_filtering_network_service_groups":     {Tok: ziaResource(ziaMod, "FirewallFilteringServiceGroups")},
			"zia_firewall_filtering_network_application_groups": {Tok: ziaResource(ziaMod, "FirewallFilteringApplicationGroups")},
			"zia_forwarding_control_rule":                       {Tok: ziaResource(ziaMod, "ForwardingControlRule")},
			"zia_forwarding_control_zpa_gateway":                {Tok: ziaResource(ziaMod, "ForwardingControlZPAGateway")},
			"zia_ssl_inspection_rules":                          {Tok: ziaResource(ziaMod, "SSLInspectionRules")},
			"zia_sandbox_rules":                                 {Tok: ziaResource(ziaMod, "SandboxRules")},
			"zia_sandbox_behavioral_analysis":                   {Tok: ziaResource(ziaMod, "SandboxBehavioralAnalysis")},
			"zia_sandbox_file_submission":                       {Tok: ziaResource(ziaMod, "SandboxFileSubmission")},
			"zia_traffic_forwarding_gre_tunnel":                 {Tok: ziaResource(ziaMod, "TrafficForwardingGRETunnel")},
			"zia_traffic_forwarding_static_ip":                  {Tok: ziaResource(ziaMod, "TrafficForwardingStaticIP")},
			"zia_traffic_forwarding_vpn_credentials":            {Tok: ziaResource(ziaMod, "TrafficForwardingVPNCredentials")},
			"zia_location_management":                           {Tok: ziaResource(ziaMod, "LocationManagement")},
			"zia_url_categories":                                {Tok: ziaResource(ziaMod, "URLCategories")},
			"zia_url_filtering_rules":                           {Tok: ziaResource(ziaMod, "URLFilteringRules")},
			"zia_url_filtering_and_cloud_app_settings":          {Tok: ziaResource(ziaMod, "URLFilteringCloudAppSettings")},
			"zia_user_management":                               {Tok: ziaResource(ziaMod, "UserManagement")},
			"zia_rule_labels":                                   {Tok: ziaResource(ziaMod, "RuleLabels")},
			"zia_auth_settings_urls":                            {Tok: ziaResource(ziaMod, "AuthSettingsURLs")},
			"zia_security_settings": {
				Tok: ziaResource(ziaMod, "SecuritySettings"),
				// Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"zia_activation_status": {
				Tok: ziaDataSource(ziaMod, "getActivationStatus"),
			},
			"zia_cloud_applications": {
				Tok: ziaDataSource(ziaMod, "getCloudApplications"),
			},
			"zia_advanced_settings": {
				Tok: ziaDataSource(ziaMod, "getAdvancedSettings"),
			},
			"zia_advanced_threat_settings": {
				Tok: ziaDataSource(ziaMod, "getAdvancedThreatSettings"),
			},
			"zia_atp_malicious_urls": {
				Tok: ziaDataSource(ziaMod, "getATPMaliciousURLs"),
			},
			"zia_atp_malware_inspection": {
				Tok: ziaDataSource(ziaMod, "getMalwareInspection"),
			},
			"zia_atp_malware_policy": {
				Tok: ziaDataSource(ziaMod, "getMalwarePolicy"),
			},
			"zia_atp_malware_protocols": {
				Tok: ziaDataSource(ziaMod, "getMalwareProtocols"),
			},
			"zia_atp_malware_settings": {
				Tok: ziaDataSource(ziaMod, "getMalwareSettings"),
			},
			"zia_atp_security_exceptions": {
				Tok: ziaDataSource(ziaMod, "getATPSecurityExceptions"),
			},
			"zia_admin_users": {
				Tok: ziaDataSource(ziaMod, "getAdminUsers"),
			},
			"zia_admin_roles": {
				Tok: ziaDataSource(ziaMod, "getAdminRoles"),
			},
			"zia_cloud_app_control_rule": {
				Tok: ziaDataSource(ziaMod, "getCloudAppControlRule"),
			},
			"zia_cloud_browser_isolation_profile": {
				Tok: ziaDataSource(ziaMod, "getCbiProfile"),
			},
			"zia_end_user_notification": {
				Tok: ziaDataSource(ziaMod, "getEndUserNotification"),
			},
			"zia_user_management": {
				Tok: ziaDataSource(ziaMod, "getUserManagement"),
			},
			"zia_group_management": {
				Tok: ziaDataSource(ziaMod, "getGroupManagement"),
				// Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zia_department_management": {
				Tok: ziaDataSource(ziaMod, "getDepartmentManagement"),
				// Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zia_file_type_control_rules": {
				Tok: ziaDataSource(ziaMod, "getFileTypeControlRules"),
				// Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zia_firewall_dns_rule": {
				Tok: ziaDataSource(ziaMod, "getFirewallDNSRules"),
				// Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zia_firewall_ips_rule": {
				Tok:  ziaDataSource(ziaMod, "getIPSFirewallRule"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zia_firewall_filtering_rule": {
				Tok: ziaDataSource(ziaMod, "getFirewallFilteringRule"),
			},
			"zia_firewall_filtering_destination_groups": {
				Tok: ziaDataSource(ziaMod, "getFirewallFilteringDestinationGroups"),
			},
			"zia_firewall_filtering_ip_source_groups": {
				Tok: ziaDataSource(ziaMod, "getFirewallFilteringSourceIPGroups"),
			},
			"zia_firewall_filtering_application_services": {
				Tok: ziaDataSource(ziaMod, "getFirewallFilteringAppServices"),
				// Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zia_firewall_filtering_application_services_group": {
				Tok: ziaDataSource(ziaMod, "getFirewallFilteringAppGroup"),
				// Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zia_firewall_filtering_network_service": {
				Tok: ziaDataSource(ziaMod, "getFirewallFilteringNetworkServices"),
			},
			"zia_firewall_filtering_network_service_groups": {
				Tok: ziaDataSource(ziaMod, "getFirewallFilteringNetworkServiceGroups"),
			},
			"zia_firewall_filtering_network_application_groups": {
				Tok: ziaDataSource(ziaMod, "getFirewallFilteringApplicationGroups"),
			},
			"zia_firewall_filtering_network_application": {
				Tok: ziaDataSource(ziaMod, "getFirewallFilteringApplication"),
			},
			"zia_forwarding_control_rule": {
				Tok: ziaDataSource(ziaMod, "getForwardingControlRule"),
			},
			"zia_forwarding_control_zpa_gateway": {
				Tok: ziaDataSource(ziaMod, "getForwardingControlZPAGateway"),
			},
			"zia_sandbox_behavioral_analysis": {
				Tok: ziaDataSource(ziaMod, "getSandboxBehavioralAnalysis"),
			},
			"zia_sandbox_report": {
				Tok: ziaDataSource(ziaMod, "getSandboxReport"),
			},
			"zia_traffic_forwarding_gre_tunnel": {
				Tok: ziaDataSource(ziaMod, "getTrafficForwardingGRETunnel"),
				// Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zia_traffic_forwarding_public_node_vips": {
				Tok: ziaDataSource(ziaMod, "getTrafficForwardingNodeVIPs"),
			},
			"zia_traffic_forwarding_gre_vip_recommended_list": {
				Tok: ziaDataSource(ziaMod, "getTrafficForwardingVIPRecommendedList"),
			},
			"zia_traffic_forwarding_gre_tunnel_info": {
				Tok: ziaDataSource(ziaMod, "getTrafficForwardingGRETunnelInfo"),
			},
			"zia_gre_internal_ip_range_list": {
				Tok: ziaDataSource(ziaMod, "getTrafficForwardingGREInternalIPRange"),
				// Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zia_traffic_forwarding_static_ip": {
				Tok: ziaDataSource(ziaMod, "getTrafficForwardingStaticIP"),
			},
			"zia_traffic_forwarding_vpn_credentials": {
				Tok: ziaDataSource(ziaMod, "getTrafficForwardingVPNCredentials"),
			},
			"zia_location_management": {
				Tok: ziaDataSource(ziaMod, "getLocationManagement"),
			},
			"zia_location_groups": {
				Tok: ziaDataSource(ziaMod, "getLocationGroups"),
			},
			"zia_location_lite": {
				Tok: ziaDataSource(ziaMod, "getLocationLite"),
			},
			"zia_url_categories": {
				Tok: ziaDataSource(ziaMod, "getURLCategories"),
			},
			"zia_url_filtering_rules": {
				Tok: ziaDataSource(ziaMod, "getURLFilteringRules"),
			},
			"zia_url_filtering_and_cloud_app_settings": {
				Tok: ziaDataSource(ziaMod, "getURLFilteringCloudAppSettings"),
			},
			"zia_dlp_engines": {
				Tok: ziaDataSource(ziaMod, "getDLPEngines"),
			},
			"zia_dlp_dictionaries": {
				Tok: ziaDataSource(ziaMod, "getDLPDictionaries"),
			},
			"zia_dlp_notification_templates": {
				Tok: ziaDataSource(ziaMod, "getDLPNotificationTemplates"),
			},
			"zia_dlp_web_rules": {
				Tok: ziaDataSource(ziaMod, "getDLPWebRules"),
			},
			"zia_dlp_edm_schema": {
				Tok: ziaDataSource(ziaMod, "getDLPEDMSchema"),
			},
			"zia_dlp_icap_servers": {
				Tok: ziaDataSource(ziaMod, "getIcapServers"),
			},
			"zia_dlp_idm_profile_lite": {
				Tok: ziaDataSource(ziaMod, "getDLPIDMProfileLite"),
				// Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zia_dlp_idm_profiles": {
				Tok: ziaDataSource(ziaMod, "getDLPIDMProfiles"),
			},
			"zia_dlp_incident_receiver_servers": {
				Tok: ziaDataSource(ziaMod, "getDLPIncidentReceiverServers"),
			},
			"zia_rule_labels": {
				Tok: ziaDataSource(ziaMod, "getRuleLabels"),
			},
			"zia_device_groups": {
				Tok: ziaDataSource(ziaMod, "getDeviceGroups"),
			},
			"zia_devices": {
				Tok: ziaDataSource(ziaMod, "getDevices"),
			},
			"zia_auth_settings_urls": {
				Tok: ziaDataSource(ziaMod, "getAuthSettingsURLs"),
			},
			"zia_sandbox_rules": {
				Tok: ziaDataSource(ziaMod, "getSandboxRules"),
			},
			"zia_ssl_inspection_rules": {
				Tok: ziaDataSource(ziaMod, "getSSLInspectionRules"),
			},
			"zia_security_settings": {
				Tok: ziaDataSource(ziaMod, "getSecuritySettings"),
				// Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"zia_firewall_filtering_time_window": {
				Tok: ziaDataSource(ziaMod, "getTimeWindow"),
			},
			"zia_workload_groups": {
				Tok: ziaDataSource(ziaMod, "getWorkloadGroups"),
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@bdzscaler/pulumi-zia",

			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "zscaler_pulumi_zia",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/zscaler/pulumi-%[1]s/sdk/", ziaPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				ziaPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "zscaler.PulumiPackage",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
	}

	prov.MustComputeTokens(tks.SingleModule("zia_", ziaMod, tks.MakeStandard(ziaPkg)))

	prov.SetAutonaming(255, "-")

	prov.MustApplyAutoAliases()

	return prov
}
