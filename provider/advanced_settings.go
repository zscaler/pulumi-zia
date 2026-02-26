// Copyright (c) 2023 Zscaler Technology Alliances, <zscaler-partner-labs@z-bd.com>
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

// Package provider implements the AdvancedSettings resource.
// Adopted from terraform-provider-zia resource_zia_advanced_settings.go.
// Singleton: UpdateAdvancedSettings for create/update, GetAdvancedSettings for read. Delete no-op.

package provider

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/advanced_settings"
)

const advancedSettingsID = "advanced_settings"

// AdvancedSettings implements the zia:index:AdvancedSettings resource.
type AdvancedSettings struct{}

// AdvancedSettingsArgs are the inputs.
type AdvancedSettingsArgs struct {
	AuthBypassUrls                                      []string `pulumi:"authBypassUrls,optional"`
	AuthBypassApps                                      []string `pulumi:"authBypassApps,optional"`
	KerberosBypassApps                                  []string `pulumi:"kerberosBypassApps,optional"`
	BasicBypassApps                                     []string `pulumi:"basicBypassApps,optional"`
	DigestAuthBypassApps                                []string `pulumi:"digestAuthBypassApps,optional"`
	DnsResolutionOnTransparentProxyExemptApps           []string `pulumi:"dnsResolutionOnTransparentProxyExemptApps,optional"`
	DnsResolutionOnTransparentProxyIpv6ExemptApps       []string `pulumi:"dnsResolutionOnTransparentProxyIpv6ExemptApps,optional"`
	DnsResolutionOnTransparentProxyApps                 []string `pulumi:"dnsResolutionOnTransparentProxyApps,optional"`
	DnsResolutionOnTransparentProxyIpv6Apps             []string `pulumi:"dnsResolutionOnTransparentProxyIpv6Apps,optional"`
	BlockDomainFrontingApps                             []string `pulumi:"blockDomainFrontingApps,optional"`
	PreferSniOverConnHostApps                           []string `pulumi:"preferSniOverConnHostApps,optional"`
	DnsResolutionOnTransparentProxyExemptUrlCategories   []string `pulumi:"dnsResolutionOnTransparentProxyExemptUrlCategories,optional"`
	DnsResolutionOnTransparentProxyIpv6ExemptUrlCategories []string `pulumi:"dnsResolutionOnTransparentProxyIpv6ExemptUrlCategories,optional"`
	DnsResolutionOnTransparentProxyUrlCategories        []string `pulumi:"dnsResolutionOnTransparentProxyUrlCategories,optional"`
	DnsResolutionOnTransparentProxyIpv6UrlCategories    []string `pulumi:"dnsResolutionOnTransparentProxyIpv6UrlCategories,optional"`
	AuthBypassUrlCategories                             []string `pulumi:"authBypassUrlCategories,optional"`
	DomainFrontingBypassUrlCategories                   []string `pulumi:"domainFrontingBypassUrlCategories,optional"`
	KerberosBypassUrlCategories                         []string `pulumi:"kerberosBypassUrlCategories,optional"`
	BasicBypassUrlCategories                            []string `pulumi:"basicBypassUrlCategories,optional"`
	HttpRangeHeaderRemoveUrlCategories                  []string `pulumi:"httpRangeHeaderRemoveUrlCategories,optional"`
	DigestAuthBypassUrlCategories                       []string `pulumi:"digestAuthBypassUrlCategories,optional"`
	SniDnsOptimizationBypassUrlCategories               []string `pulumi:"sniDnsOptimizationBypassUrlCategories,optional"`
	KerberosBypassUrls                                  []string `pulumi:"kerberosBypassUrls,optional"`
	DigestAuthBypassUrls                                []string `pulumi:"digestAuthBypassUrls,optional"`
	DnsResolutionOnTransparentProxyExemptUrls           []string `pulumi:"dnsResolutionOnTransparentProxyExemptUrls,optional"`
	DnsResolutionOnTransparentProxyUrls                 []string `pulumi:"dnsResolutionOnTransparentProxyUrls,optional"`
	EnableDnsResolutionOnTransparentProxy               *bool    `pulumi:"enableDnsResolutionOnTransparentProxy,optional"`
	EnableIpv6DnsResolutionOnTransparentProxy           *bool    `pulumi:"enableIpv6DnsResolutionOnTransparentProxy,optional"`
	EnableIpv6DnsOptimizationOnAllTransparentProxy      *bool    `pulumi:"enableIpv6DnsOptimizationOnAllTransparentProxy,optional"`
	EnableEvaluatePolicyOnGlobalSslBypass               *bool    `pulumi:"enableEvaluatePolicyOnGlobalSslBypass,optional"`
	EnableOffice365                                     *bool    `pulumi:"enableOffice365,optional"`
	LogInternalIp                                       *bool    `pulumi:"logInternalIp,optional"`
	EnforceSurrogateIpForWindowsApp                    *bool    `pulumi:"enforceSurrogateIpForWindowsApp,optional"`
	TrackHttpTunnelOnHttpPorts                          *bool    `pulumi:"trackHttpTunnelOnHttpPorts,optional"`
	BlockHttpTunnelOnNonHttpPorts                       *bool    `pulumi:"blockHttpTunnelOnNonHttpPorts,optional"`
	BlockDomainFrontingOnHostHeader                     *bool    `pulumi:"blockDomainFrontingOnHostHeader,optional"`
	ZscalerClientConnector1AndPacRoadWarriorInFirewall  *bool    `pulumi:"zscalerClientConnector1AndPacRoadWarriorInFirewall,optional"`
	CascadeUrlFiltering                                 *bool    `pulumi:"cascadeUrlFiltering,optional"`
	EnablePolicyForUnauthenticatedTraffic               *bool    `pulumi:"enablePolicyForUnauthenticatedTraffic,optional"`
	BlockNonCompliantHttpRequestOnHttpPorts            *bool    `pulumi:"blockNonCompliantHttpRequestOnHttpPorts,optional"`
	EnableAdminRankAccess                               *bool    `pulumi:"enableAdminRankAccess,optional"`
	Http2NonbrowserTrafficEnabled                       *bool    `pulumi:"http2NonbrowserTrafficEnabled,optional"`
	EcsForAllEnabled                                    *bool    `pulumi:"ecsForAllEnabled,optional"`
	DynamicUserRiskEnabled                              *bool    `pulumi:"dynamicUserRiskEnabled,optional"`
	BlockConnectHostSniMismatch                         *bool    `pulumi:"blockConnectHostSniMismatch,optional"`
	PreferSniOverConnHost                               *bool    `pulumi:"preferSniOverConnHost,optional"`
	SipaXffHeaderEnabled                                *bool    `pulumi:"sipaXffHeaderEnabled,optional"`
	BlockNonHttpOnHttpPortEnabled                       *bool    `pulumi:"blockNonHttpOnHttpPortEnabled,optional"`
	UiSessionTimeout                                    *int     `pulumi:"uiSessionTimeout,optional"`
}

// AdvancedSettingsState is the persisted state.
type AdvancedSettingsState struct {
	AdvancedSettingsArgs
	ResourceId string `pulumi:"resourceId"`
}

func (AdvancedSettings) Create(ctx context.Context, req infer.CreateRequest[AdvancedSettingsArgs]) (infer.CreateResponse[AdvancedSettingsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[AdvancedSettingsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := advancedSettingsArgsToAPI(req.Inputs, nil)
	if _, _, err := advanced_settings.UpdateAdvancedSettings(ctx, service, &apiReq); err != nil {
		return infer.CreateResponse[AdvancedSettingsState]{}, err
	}

	if shouldActivate() {
		time.Sleep(1 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[AdvancedSettingsState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	// Read back to populate full state
	resp, err := advanced_settings.GetAdvancedSettings(ctx, service)
	if err != nil {
		state := AdvancedSettingsState{
			AdvancedSettingsArgs: req.Inputs,
			ResourceId:           advancedSettingsID,
		}
		return infer.CreateResponse[AdvancedSettingsState]{
			ID:     advancedSettingsID,
			Output: state,
		}, nil
	}
	state := advancedSettingsAPIToState(resp)
	return infer.CreateResponse[AdvancedSettingsState]{
		ID:     advancedSettingsID,
		Output: state,
	}, nil
}

func (AdvancedSettings) Read(ctx context.Context, req infer.ReadRequest[AdvancedSettingsArgs, AdvancedSettingsState]) (infer.ReadResponse[AdvancedSettingsArgs, AdvancedSettingsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[AdvancedSettingsArgs, AdvancedSettingsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	resp, err := advanced_settings.GetAdvancedSettings(ctx, service)
	if err != nil {
		return infer.ReadResponse[AdvancedSettingsArgs, AdvancedSettingsState]{}, err
	}
	if resp == nil {
		return infer.ReadResponse[AdvancedSettingsArgs, AdvancedSettingsState]{}, fmt.Errorf("couldn't read advanced settings")
	}

	state := advancedSettingsAPIToState(resp)
	return infer.ReadResponse[AdvancedSettingsArgs, AdvancedSettingsState]{
		ID:     advancedSettingsID,
		Inputs: state.AdvancedSettingsArgs,
		State:  state,
	}, nil
}

func (AdvancedSettings) Update(ctx context.Context, req infer.UpdateRequest[AdvancedSettingsArgs, AdvancedSettingsState]) (infer.UpdateResponse[AdvancedSettingsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[AdvancedSettingsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := advancedSettingsArgsToAPI(req.Inputs, nil)
	if _, _, err := advanced_settings.UpdateAdvancedSettings(ctx, service, &apiReq); err != nil {
		return infer.UpdateResponse[AdvancedSettingsState]{}, err
	}

	if shouldActivate() {
		time.Sleep(1 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[AdvancedSettingsState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	resp, err := advanced_settings.GetAdvancedSettings(ctx, service)
	if err != nil {
		state := AdvancedSettingsState{
			AdvancedSettingsArgs: req.Inputs,
			ResourceId:           advancedSettingsID,
		}
		return infer.UpdateResponse[AdvancedSettingsState]{Output: state}, nil
	}
	return infer.UpdateResponse[AdvancedSettingsState]{Output: advancedSettingsAPIToState(resp)}, nil
}

func (AdvancedSettings) Delete(ctx context.Context, req infer.DeleteRequest[AdvancedSettingsState]) (infer.DeleteResponse, error) {
	return infer.DeleteResponse{}, nil
}

func (AdvancedSettings) Annotate(a infer.Annotator) {
	describeResource(a, &AdvancedSettings{}, `The zia_advanced_settings resource manages advanced settings in the Zscaler Internet Access (ZIA) cloud service. This singleton resource controls a wide range of advanced proxy, authentication, DNS resolution, and security settings including domain fronting protection, HTTP tunnel tracking, surrogate IP enforcement, and session timeout configuration.

For more information, see the [ZIA Advanced Settings documentation](https://help.zscaler.com/zia/advanced-settings).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Advanced Settings

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.AdvancedSettings("example", {
    enableOffice365: true,
    logInternalIp: true,
    blockHttpTunnelOnNonHttpPorts: true,
    blockDomainFrontingOnHostHeader: true,
    authBypassUrls: [".example.com"],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.AdvancedSettings("example",
    enable_office365=True,
    log_internal_ip=True,
    block_http_tunnel_on_non_http_ports=True,
    block_domain_fronting_on_host_header=True,
    auth_bypass_urls=[".example.com"],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:AdvancedSettings
    properties:
      enableOffice365: true
      logInternalIp: true
      blockHttpTunnelOnNonHttpPorts: true
      blockDomainFrontingOnHostHeader: true
      authBypassUrls:
        - .example.com
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

This is a singleton resource and does not support traditional import. It is automatically managed by the provider.
`)
}

func (a *AdvancedSettingsArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.AuthBypassUrls, "URLs that bypass authentication.")
	ann.Describe(&a.AuthBypassApps, "Cloud applications that bypass authentication.")
	ann.Describe(&a.KerberosBypassApps, "Cloud applications that bypass Kerberos authentication.")
	ann.Describe(&a.BasicBypassApps, "Cloud applications that bypass basic authentication.")
	ann.Describe(&a.DigestAuthBypassApps, "Cloud applications that bypass digest authentication.")
	ann.Describe(&a.DnsResolutionOnTransparentProxyExemptApps, "Cloud applications exempt from DNS resolution on transparent proxy.")
	ann.Describe(&a.DnsResolutionOnTransparentProxyIpv6ExemptApps, "Cloud applications exempt from IPv6 DNS resolution on transparent proxy.")
	ann.Describe(&a.DnsResolutionOnTransparentProxyApps, "Cloud applications with DNS resolution on transparent proxy enabled.")
	ann.Describe(&a.DnsResolutionOnTransparentProxyIpv6Apps, "Cloud applications with IPv6 DNS resolution on transparent proxy enabled.")
	ann.Describe(&a.BlockDomainFrontingApps, "Cloud applications for which domain fronting is blocked.")
	ann.Describe(&a.PreferSniOverConnHostApps, "Cloud applications that prefer SNI over CONNECT host header.")
	ann.Describe(&a.DnsResolutionOnTransparentProxyExemptUrlCategories, "URL categories exempt from DNS resolution on transparent proxy.")
	ann.Describe(&a.DnsResolutionOnTransparentProxyIpv6ExemptUrlCategories, "URL categories exempt from IPv6 DNS resolution on transparent proxy.")
	ann.Describe(&a.DnsResolutionOnTransparentProxyUrlCategories, "URL categories with DNS resolution on transparent proxy enabled.")
	ann.Describe(&a.DnsResolutionOnTransparentProxyIpv6UrlCategories, "URL categories with IPv6 DNS resolution on transparent proxy enabled.")
	ann.Describe(&a.AuthBypassUrlCategories, "URL categories that bypass authentication.")
	ann.Describe(&a.DomainFrontingBypassUrlCategories, "URL categories that bypass domain fronting detection.")
	ann.Describe(&a.KerberosBypassUrlCategories, "URL categories that bypass Kerberos authentication.")
	ann.Describe(&a.BasicBypassUrlCategories, "URL categories that bypass basic authentication.")
	ann.Describe(&a.HttpRangeHeaderRemoveUrlCategories, "URL categories for which HTTP range headers are removed.")
	ann.Describe(&a.DigestAuthBypassUrlCategories, "URL categories that bypass digest authentication.")
	ann.Describe(&a.SniDnsOptimizationBypassUrlCategories, "URL categories that bypass SNI/DNS optimization.")
	ann.Describe(&a.KerberosBypassUrls, "URLs that bypass Kerberos authentication.")
	ann.Describe(&a.DigestAuthBypassUrls, "URLs that bypass digest authentication.")
	ann.Describe(&a.DnsResolutionOnTransparentProxyExemptUrls, "URLs exempt from DNS resolution on transparent proxy.")
	ann.Describe(&a.DnsResolutionOnTransparentProxyUrls, "URLs with DNS resolution on transparent proxy enabled.")
	ann.Describe(&a.EnableDnsResolutionOnTransparentProxy, "Enable DNS resolution on transparent proxy.")
	ann.Describe(&a.EnableIpv6DnsResolutionOnTransparentProxy, "Enable IPv6 DNS resolution on transparent proxy.")
	ann.Describe(&a.EnableIpv6DnsOptimizationOnAllTransparentProxy, "Enable IPv6 DNS optimization on all transparent proxy connections.")
	ann.Describe(&a.EnableEvaluatePolicyOnGlobalSslBypass, "Enable policy evaluation on global SSL bypass.")
	ann.Describe(&a.EnableOffice365, "Enable Office 365 one-click configuration.")
	ann.Describe(&a.LogInternalIp, "Enable logging of internal IP addresses.")
	ann.Describe(&a.EnforceSurrogateIpForWindowsApp, "Enforce surrogate IP for Windows applications.")
	ann.Describe(&a.TrackHttpTunnelOnHttpPorts, "Track HTTP tunnels on HTTP ports.")
	ann.Describe(&a.BlockHttpTunnelOnNonHttpPorts, "Block HTTP tunnels on non-HTTP ports.")
	ann.Describe(&a.BlockDomainFrontingOnHostHeader, "Block domain fronting when the host header mismatches the SNI.")
	ann.Describe(&a.ZscalerClientConnector1AndPacRoadWarriorInFirewall, "Include Zscaler Client Connector and PAC road warrior traffic in firewall policy.")
	ann.Describe(&a.CascadeUrlFiltering, "Enable cascading URL filtering.")
	ann.Describe(&a.EnablePolicyForUnauthenticatedTraffic, "Enable policy evaluation for unauthenticated traffic.")
	ann.Describe(&a.BlockNonCompliantHttpRequestOnHttpPorts, "Block non-compliant HTTP requests on HTTP ports.")
	ann.Describe(&a.EnableAdminRankAccess, "Enable admin rank-based access control.")
	ann.Describe(&a.Http2NonbrowserTrafficEnabled, "Enable HTTP/2 for non-browser traffic.")
	ann.Describe(&a.EcsForAllEnabled, "Enable EDNS Client Subnet (ECS) for all DNS queries.")
	ann.Describe(&a.DynamicUserRiskEnabled, "Enable dynamic user risk scoring.")
	ann.Describe(&a.BlockConnectHostSniMismatch, "Block connections where CONNECT host and SNI mismatch.")
	ann.Describe(&a.PreferSniOverConnHost, "Prefer SNI over CONNECT host header for policy evaluation.")
	ann.Describe(&a.SipaXffHeaderEnabled, "Enable X-Forwarded-For header for SIPA traffic.")
	ann.Describe(&a.BlockNonHttpOnHttpPortEnabled, "Block non-HTTP traffic on HTTP ports.")
	ann.Describe(&a.UiSessionTimeout, "UI session timeout in minutes.")
}

func (s *AdvancedSettingsState) Annotate(a infer.Annotator) {
	a.Describe(&s.ResourceId, "The internal resource identifier for the advanced settings.")
}

func (AdvancedSettings) Diff(ctx context.Context, req infer.DiffRequest[AdvancedSettingsArgs, AdvancedSettingsState]) (infer.DiffResponse, error) {
	diff, hasChanges := diffSkippingUnsetInputs(req.State.AdvancedSettingsArgs, req.Inputs)
	return infer.DiffResponse{HasChanges: hasChanges, DetailedDiff: diff}, nil
}

var _ infer.CustomResource[AdvancedSettingsArgs, AdvancedSettingsState] = AdvancedSettings{}

func advancedSettingsArgsToAPI(in AdvancedSettingsArgs, existing *advanced_settings.AdvancedSettings) advanced_settings.AdvancedSettings {
	use := func(b *bool, def bool) bool {
		if b != nil {
			return *b
		}
		if existing != nil {
			return def
		}
		return def
	}
	useInt := func(i *int, def int) int {
		if i != nil {
			return *i
		}
		if existing != nil {
			return def
		}
		return def
	}
	result := advanced_settings.AdvancedSettings{
		AuthBypassUrls:                                emptyToNil(in.AuthBypassUrls),
		KerberosBypassUrls:                            emptyToNil(in.KerberosBypassUrls),
		DigestAuthBypassUrls:                          emptyToNil(in.DigestAuthBypassUrls),
		DnsResolutionOnTransparentProxyExemptUrls:     emptyToNil(in.DnsResolutionOnTransparentProxyExemptUrls),
		DnsResolutionOnTransparentProxyUrls:           emptyToNil(in.DnsResolutionOnTransparentProxyUrls),
		AuthBypassApps:                                emptyToNil(in.AuthBypassApps),
		KerberosBypassApps:                            emptyToNil(in.KerberosBypassApps),
		BasicBypassApps:                               emptyToNil(in.BasicBypassApps),
		DigestAuthBypassApps:                          emptyToNil(in.DigestAuthBypassApps),
		DnsResolutionOnTransparentProxyExemptApps:     emptyToNil(in.DnsResolutionOnTransparentProxyExemptApps),
		DnsResolutionOnTransparentProxyIPv6ExemptApps: emptyToNil(in.DnsResolutionOnTransparentProxyIpv6ExemptApps),
		DnsResolutionOnTransparentProxyApps:           emptyToNil(in.DnsResolutionOnTransparentProxyApps),
		DnsResolutionOnTransparentProxyIPv6Apps:      emptyToNil(in.DnsResolutionOnTransparentProxyIpv6Apps),
		BlockDomainFrontingApps:                       emptyToNil(in.BlockDomainFrontingApps),
		PreferSniOverConnHostApps:                     emptyToNil(in.PreferSniOverConnHostApps),
		DnsResolutionOnTransparentProxyExemptUrlCategories:     emptyToNil(in.DnsResolutionOnTransparentProxyExemptUrlCategories),
		DnsResolutionOnTransparentProxyIPv6ExemptUrlCategories: emptyToNil(in.DnsResolutionOnTransparentProxyIpv6ExemptUrlCategories),
		DnsResolutionOnTransparentProxyUrlCategories:           emptyToNil(in.DnsResolutionOnTransparentProxyUrlCategories),
		DnsResolutionOnTransparentProxyIPv6UrlCategories:       emptyToNil(in.DnsResolutionOnTransparentProxyIpv6UrlCategories),
		AuthBypassUrlCategories:                       emptyToNil(in.AuthBypassUrlCategories),
		DomainFrontingBypassUrlCategories:              emptyToNil(in.DomainFrontingBypassUrlCategories),
		KerberosBypassUrlCategories:                  emptyToNil(in.KerberosBypassUrlCategories),
		BasicBypassUrlCategories:                      emptyToNil(in.BasicBypassUrlCategories),
		HttpRangeHeaderRemoveUrlCategories:           emptyToNil(in.HttpRangeHeaderRemoveUrlCategories),
		DigestAuthBypassUrlCategories:                 emptyToNil(in.DigestAuthBypassUrlCategories),
		SniDnsOptimizationBypassUrlCategories:        emptyToNil(in.SniDnsOptimizationBypassUrlCategories),
		EnableDnsResolutionOnTransparentProxy:         use(in.EnableDnsResolutionOnTransparentProxy, false),
		EnableIPv6DnsResolutionOnTransparentProxy:      use(in.EnableIpv6DnsResolutionOnTransparentProxy, false),
		EnableIPv6DnsOptimizationOnAllTransparentProxy: use(in.EnableIpv6DnsOptimizationOnAllTransparentProxy, false),
		EnableEvaluatePolicyOnGlobalSSLBypass:         use(in.EnableEvaluatePolicyOnGlobalSslBypass, false),
		EnableOffice365:                               use(in.EnableOffice365, false),
		LogInternalIp:                                 use(in.LogInternalIp, false),
		EnforceSurrogateIpForWindowsApp:                use(in.EnforceSurrogateIpForWindowsApp, false),
		TrackHttpTunnelOnHttpPorts:                     use(in.TrackHttpTunnelOnHttpPorts, false),
		BlockHttpTunnelOnNonHttpPorts:                 use(in.BlockHttpTunnelOnNonHttpPorts, false),
		BlockDomainFrontingOnHostHeader:               use(in.BlockDomainFrontingOnHostHeader, false),
		ZscalerClientConnector1AndPacRoadWarriorInFirewall: use(in.ZscalerClientConnector1AndPacRoadWarriorInFirewall, false),
		CascadeUrlFiltering:                           use(in.CascadeUrlFiltering, false),
		EnablePolicyForUnauthenticatedTraffic:          use(in.EnablePolicyForUnauthenticatedTraffic, false),
		BlockNonCompliantHttpRequestOnHttpPorts:       use(in.BlockNonCompliantHttpRequestOnHttpPorts, false),
		EnableAdminRankAccess:                         use(in.EnableAdminRankAccess, false),
		Http2NonbrowserTrafficEnabled:                 use(in.Http2NonbrowserTrafficEnabled, false),
		EcsForAllEnabled:                              use(in.EcsForAllEnabled, false),
		DynamicUserRiskEnabled:                        use(in.DynamicUserRiskEnabled, false),
		BlockConnectHostSniMismatch:                    use(in.BlockConnectHostSniMismatch, false),
		PreferSniOverConnHost:                         use(in.PreferSniOverConnHost, false),
		SipaXffHeaderEnabled:                          use(in.SipaXffHeaderEnabled, false),
		BlockNonHttpOnHttpPortEnabled:                 use(in.BlockNonHttpOnHttpPortEnabled, false),
		UISessionTimeout:                              useInt(in.UiSessionTimeout, 0),
	}
	return result
}

func emptyToNil(s []string) []string {
	if len(s) == 0 {
		return nil
	}
	return s
}

func advancedSettingsAPIToState(r *advanced_settings.AdvancedSettings) AdvancedSettingsState {
	return AdvancedSettingsState{
		AdvancedSettingsArgs: AdvancedSettingsArgs{
			AuthBypassUrls:                                r.AuthBypassUrls,
			AuthBypassApps:                                r.AuthBypassApps,
			KerberosBypassApps:                            r.KerberosBypassApps,
			BasicBypassApps:                               r.BasicBypassApps,
			DigestAuthBypassApps:                          r.DigestAuthBypassApps,
			DnsResolutionOnTransparentProxyExemptApps:     r.DnsResolutionOnTransparentProxyExemptApps,
			DnsResolutionOnTransparentProxyIpv6ExemptApps: r.DnsResolutionOnTransparentProxyIPv6ExemptApps,
			DnsResolutionOnTransparentProxyApps:           r.DnsResolutionOnTransparentProxyApps,
			DnsResolutionOnTransparentProxyIpv6Apps:       r.DnsResolutionOnTransparentProxyIPv6Apps,
			BlockDomainFrontingApps:                       r.BlockDomainFrontingApps,
			PreferSniOverConnHostApps:                     r.PreferSniOverConnHostApps,
			DnsResolutionOnTransparentProxyExemptUrlCategories:   r.DnsResolutionOnTransparentProxyExemptUrlCategories,
			DnsResolutionOnTransparentProxyIpv6ExemptUrlCategories: r.DnsResolutionOnTransparentProxyIPv6ExemptUrlCategories,
			DnsResolutionOnTransparentProxyUrlCategories:   r.DnsResolutionOnTransparentProxyUrlCategories,
			DnsResolutionOnTransparentProxyIpv6UrlCategories: r.DnsResolutionOnTransparentProxyIPv6UrlCategories,
			AuthBypassUrlCategories:                   r.AuthBypassUrlCategories,
			DomainFrontingBypassUrlCategories:         r.DomainFrontingBypassUrlCategories,
			KerberosBypassUrlCategories:               r.KerberosBypassUrlCategories,
			BasicBypassUrlCategories:                   r.BasicBypassUrlCategories,
			HttpRangeHeaderRemoveUrlCategories:         r.HttpRangeHeaderRemoveUrlCategories,
			DigestAuthBypassUrlCategories:              r.DigestAuthBypassUrlCategories,
			SniDnsOptimizationBypassUrlCategories:      r.SniDnsOptimizationBypassUrlCategories,
			KerberosBypassUrls:                          r.KerberosBypassUrls,
			DigestAuthBypassUrls:                        r.DigestAuthBypassUrls,
			DnsResolutionOnTransparentProxyExemptUrls:   r.DnsResolutionOnTransparentProxyExemptUrls,
			DnsResolutionOnTransparentProxyUrls:         r.DnsResolutionOnTransparentProxyUrls,
			EnableDnsResolutionOnTransparentProxy:       boolPtr(r.EnableDnsResolutionOnTransparentProxy),
			EnableIpv6DnsResolutionOnTransparentProxy:   boolPtr(r.EnableIPv6DnsResolutionOnTransparentProxy),
			EnableIpv6DnsOptimizationOnAllTransparentProxy: boolPtr(r.EnableIPv6DnsOptimizationOnAllTransparentProxy),
			EnableEvaluatePolicyOnGlobalSslBypass:       boolPtr(r.EnableEvaluatePolicyOnGlobalSSLBypass),
			EnableOffice365:                             boolPtr(r.EnableOffice365),
			LogInternalIp:                               boolPtr(r.LogInternalIp),
			EnforceSurrogateIpForWindowsApp:            boolPtr(r.EnforceSurrogateIpForWindowsApp),
			TrackHttpTunnelOnHttpPorts:                  boolPtr(r.TrackHttpTunnelOnHttpPorts),
			BlockHttpTunnelOnNonHttpPorts:               boolPtr(r.BlockHttpTunnelOnNonHttpPorts),
			BlockDomainFrontingOnHostHeader:             boolPtr(r.BlockDomainFrontingOnHostHeader),
			ZscalerClientConnector1AndPacRoadWarriorInFirewall: boolPtr(r.ZscalerClientConnector1AndPacRoadWarriorInFirewall),
			CascadeUrlFiltering:                         boolPtr(r.CascadeUrlFiltering),
			EnablePolicyForUnauthenticatedTraffic:        boolPtr(r.EnablePolicyForUnauthenticatedTraffic),
			BlockNonCompliantHttpRequestOnHttpPorts:     boolPtr(r.BlockNonCompliantHttpRequestOnHttpPorts),
			EnableAdminRankAccess:                       boolPtr(r.EnableAdminRankAccess),
			Http2NonbrowserTrafficEnabled:               boolPtr(r.Http2NonbrowserTrafficEnabled),
			EcsForAllEnabled:                            boolPtr(r.EcsForAllEnabled),
			DynamicUserRiskEnabled:                     boolPtr(r.DynamicUserRiskEnabled),
			BlockConnectHostSniMismatch:                  boolPtr(r.BlockConnectHostSniMismatch),
			PreferSniOverConnHost:                       boolPtr(r.PreferSniOverConnHost),
			SipaXffHeaderEnabled:                       boolPtr(r.SipaXffHeaderEnabled),
			BlockNonHttpOnHttpPortEnabled:              boolPtr(r.BlockNonHttpOnHttpPortEnabled),
			UiSessionTimeout:                            intPtr(r.UISessionTimeout),
		},
		ResourceId: advancedSettingsID,
	}
}
