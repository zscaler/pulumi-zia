# Changelog

## Unreleased

### Maintenance

- Temporarily removed the .NET SDK. Access to the NuGet account that publishes `zscaler.PulumiPackage.Zia` was lost and is being migrated, so the release workflow could no longer publish it (1.4.1 was not published to NuGet). The `sdk/dotnet` SDK, the `examples/dotnet` example and test, the `csharp` language settings in the provider schema, the `dotnet` CI and release jobs, and the .NET Makefile targets were removed. Versions already on NuGet (up to 1.4.0) remain available. .NET support will return once publishing access is restored.

## 1.4.1 (November, 18 2026)

### Notes

- Release date: **(November, 18 2026)**
- Supported Terraform version: **v1.x**
- Golang: **v1.26**

### Maintenance

- [PR #85](https://github.com/zscaler/pulumi-zia/pull/85) - Updated to [Zscaler-SDK-GO v3.8.60](https://github.com/zscaler/zscaler-sdk-go/releases/tag/v3.8.60).
- [PR #85](https://github.com/zscaler/pulumi-zia/pull/85) - Go modules, the workspace, and the CI/release workflows now use Go 1.26.
- [PR #85](https://github.com/zscaler/pulumi-zia/pull/85) - CI/release workflows now build the Python SDK with Python 3.12. GitHub is moving `ubuntu-latest` to Ubuntu 26.04, which has no Python 3.9 build, so the Python SDK job failed at `Setup Python`. The published package still supports Python 3.9 and later (`requires-python = ">=3.9"`).
- [PR #85](https://github.com/zscaler/pulumi-zia/pull/85) - CI/release workflows now build the Node.js SDK with Node.js 24 (was 20). `@pulumi/pulumi` 3.249.0 and later require Node.js 22 or newer, and the SDK's `^3.142.0` range resolves to the latest release, so `yarn install` failed with `The engine "node" is incompatible with this module`. The published package's dependencies are unchanged.
- [PR #85](https://github.com/zscaler/pulumi-zia/pull/85) - CI/release workflows now install the .NET 10 SDK instead of the end-of-life .NET 6.0 and 3.1 SDKs. The NuGet package still targets `net6.0`.
- [PR #85](https://github.com/zscaler/pulumi-zia/pull/85) - Reverted the `github.com/pulumi/pulumi/sdk/v3` v3.249.0 bump ([PR #81](https://github.com/zscaler/pulumi-zia/pull/81), [PR #82](https://github.com/zscaler/pulumi-zia/pull/82)) to v3.217.0. It does not build against `pulumi/pkg/v3` v3.217.0, the version `pulumi-go-provider` v1.3.0 requires.

### Bug Fix

- [PR #85](https://github.com/zscaler/pulumi-zia/pull/85) [Issue #83](https://github.com/zscaler/pulumi-zia/issues/83) - `pulumi preview` no longer sends write requests (`POST`/`PUT`) to the ZIA API. Previously, a preview could change the live tenant while Pulumi state stayed unchanged:
  - **In-place updates:** none of the 69 resource `Update` methods checked `req.DryRun`. A planned update was applied during preview, for example a `UrlCategory` URL change sent `PUT /urlCategories/{id}?action=ADD_TO_LIST`, and rule updates ran policy updates and rule reorder `PUT`s.
  - **Creates:** 33 resource `Create` methods did not check `req.DryRun`, including `Activation`, `DlpDictionary`, `DlpEngine`, `TrafficForwardingGreTunnel`, `SandboxSubmission` (which uploaded the file to Sandbox), and the ATP, advanced, and security settings resources.
  - **Activation:** previewing a new `Activation` resource activated the tenant's staged configuration, and with `ZIA_ACTIVATION=true` any previewed update or create could trigger activation.

  Every resource `Create` and `Update` now returns the predicted state during preview before making any API call. Updates keep computed values such as IDs from the current state. Two input checks now run at `pulumi up` instead of preview: `Activation` `status` must be `ACTIVE`, and `DcExclusion` requires `datacenterId`.
- Added regression tests that fail if any resource `Create` or `Update` does not start with a `req.DryRun` guard, or if a preview reaches the ZIA client.

## 1.4.0 (June, 30 2026)

### Notes

- Release date: **(June, 30 2026)**
- Supported Terraform version: **v1.x**

### Bug Fix

- [PR #80](https://github.com/zscaler/pulumi-zia/pull/80) - Fixed non-deterministic rule ordering (scrambled `order` values, drift on re-apply, and predefined-rule movement) when creating multiple order-based rules concurrently. Replaced the single-pass, count-based reorder with a diff-based convergent engine across all 15 rule-based resources (SSL inspection, URL filtering, cloud firewall filtering/DNS/IPS, NAT, forwarding control, file type control, web DLP, cloud app control, CASB DLP/malware, bandwidth control, sandbox, and traffic capture), porting the reconciliation logic to match the Terraform provider. The engine now only updates rules whose position actually drifted, defers out-of-range orders, and converges over multiple passes, eliminating the need for `--parallel 1`.


## 1.3.9 (April, 2 2026)

### Notes

- Release date: **(April, 2 2026)**
- Supported Terraform version: **v1.x**

### Maintanance

- [PR #73](https://github.com/zscaler/pulumi-zia/pull/73) - Updated to [Zscaler-SDK-GO v3.8.29](https://github.com/zscaler/zscaler-sdk-go/releases/tag/v3.8.29)

## 1.3.8 (March, 13 2026)

### Notes

- Release date: **(March, 13 2026)**
- Supported Terraform version: **v1.x**

### Bug Fix

- [PR #68](https://github.com/zscaler/pulumi-zia/pull/68) - Marked `authenticationToken` as secret in `zia_cloud_nss_feed` to prevent plaintext storage in state. Improved graceful handling of externally deleted resources across all Read, Update, and Delete functions by returning empty state on not-found instead of erroring, matching Terraform provider behavior.

- [PR #68](https://github.com/zscaler/pulumi-zia/pull/68) - Implemented `DiffConfig` on the provider configuration to prevent resource replacement (`replace [diff: ~provider]`) when upgrading the provider version.

## 1.3.7 (March, 13 2026)

### Notes

- Release date: **(March, 13 2026)**
- Supported Terraform version: **v1.x**

### Bug Fix

- [PR #67](https://github.com/zscaler/pulumi-zia/pull/67) - Added context-aware concurrency safeguards to all rule-based resources to improve parallel creation reliability and graceful cancellation. Optimized global reorder mutex to avoid cross-resource-type blocking during bulk deployments.

## 1.3.6 (March, 12 2026)

### Notes

- Release date: **(March, 12 2026)**
- Supported Terraform version: **v1.x**

### Bug Fix

- [PR #66](https://github.com/zscaler/pulumi-zia/pull/66) - Fixed tautological nil check in `getFileTypeCategories` data source.
- Removed hardcoded credential placeholder from `zia_traffic_forwarding_vpn_credentials` example documentation; examples now use `pulumi.Config.requireSecret` for pre-shared keys.

### Enhancements

- [PR #66](https://github.com/zscaler/pulumi-zia/pull/66) - Added `dedicatedIpGatewayId` attribute to resource `zia_forwarding_control_rule` for Dedicated IP Gateway support.
- [PR #66](https://github.com/zscaler/pulumi-zia/pull/66) - Added new attributes to resource `zia_url_filtering_and_cloud_app_settings`: `enableDeepSeekPrompt`, `enableWriterPrompt`, `enableGrokPrompt`, `enableMistralAiPrompt`, `enableClaudePrompt`, `enableGrammarlyPrompt`, `zveloDbLookupDisabled`, `enableCreativeCommonsSearchResults`.
- [PR #66](https://github.com/zscaler/pulumi-zia/pull/66) - Fixed Go SDK version stamping: release workflow now commits regenerated Go SDK with the correct version and updates the `sdk/` tag, resolving `1.0.0-alpha.0+dev` resolution errors for Go users.

## 1.3.5 (March, 11 2026)

### Notes

- Release date: **(March, 11 2026)**
- Supported Terraform version: **v1.x**

### Bug Fix

- [PR #65](https://github.com/zscaler/pulumi-zia/pull/65) - Fixed `validityStartTime` / `validityEndTime` drift in `zia_url_filtering_rules` by normalizing RFC1123 day-of-week during input validation, preventing perpetual diffs caused by incorrect weekday names.
- [PR #65](https://github.com/zscaler/pulumi-zia/pull/65) - Fixed `cbiProfile` drift in `zia_url_filtering_rules` by excluding server-computed `profileSeq` from state and updating struct pointer comparison in diff logic to skip nil sub-fields in user inputs.

## 1.3.4 (March, 10 2026)

### Notes

- Release date: **(March, 10 2026)**
- Supported Terraform version: **v1.x**

### Bug Fix

- [PR #60](https://github.com/zscaler/pulumi-zia/pull/60) - Fixed rule reordering across all 15 rule-base resources by stripping read-only fields (`Predefined`, `DefaultRule`, `AccessControl`) during `updateOrder`, preventing `"Request body is invalid"` API errors on predefined and default rules.
- [PR #60](https://github.com/zscaler/pulumi-zia/pull/60) - Added order validation (`order >= 1`) to Create and Update callbacks across all rule-base resources: `ssl_inspection`, `firewall_filtering`, `firewall_dns`, `firewall_ips`, `nat_control`, `traffic_capture`, `cloud_app_control`, `sandbox`, `bandwidth_control`, `dlp_web`, `file_type_control`, `casb_dlp`, `casb_malware`, `url_filtering`, and `forwarding_control`.

### Enhancements

- [PR #61](https://github.com/zscaler/pulumi-zia/pull/61) - Added `dedicatedIpGatewayId` attribute to resource `zia_forwarding_control_rule` for Dedicated IP Gateway support.
- [PR #61](https://github.com/zscaler/pulumi-zia/pull/61) - Added new AI prompt attributes to resource `zia_url_filtering_and_cloud_app_settings`:
  - `enableDeepSeekPrompt`
  - `enableWriterPrompt`
  - `enableGrokPrompt`
  - `enableMistralAiPrompt`
  - `enableClaudePrompt`
  - `enableGrammarlyPrompt`
  - `zveloDbLookupDisabled`
  - `enableCreativeCommonsSearchResults`

## 1.3.3 (February, 27 2026)

### Notes

- Release date: **(February, 26 2026)**
- Supported Terraform version: **v1.x**

### Bug Fix

- [PR #59](https://github.com/zscaler/pulumi-zia/pull/59) - Added `pluginDownloadURL` (`github://api.github.com/zscaler`) and `publisher` (`Zscaler`) to provider schema and Go SDK, enabling automatic plugin download from GitHub Releases, fixing `1.0.0-alpha.0+dev` version resolution errors, and resolving Pulumi Registry metadata generation failures.

## 1.3.2 (February, 26 2026)

### Notes

- Release date: **(February, 26 2026)**
- Supported Terraform version: **v1.x**

### Bug Fix

- [PR #58](https://github.com/zscaler/pulumi-zia/pull/58) - Fixed Go SDK module structure: moved `go.mod` from `sdk/go/pulumi-zia/` to `sdk/` and corrected module path to `github.com/zscaler/pulumi-zia/sdk`, restoring standard Go module resolution and `sdk/vX.Y.Z` tag compatibility.

## 1.3.1 (February, 25 2026)

### Notes

- Release date: **(February, 25 2026)**
- Supported Terraform version: **v1.x**

### Internal

- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Migrated provider from the Pulumi Terraform bridged framework to a **native Go provider** built with `pulumi-go-provider`. The provider now communicates directly with the Zscaler API via `zscaler-sdk-go/v3` without a Terraform runtime dependency. All resources, data sources, attributes, and authentication methods remain unchanged.

### Enhancements

- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Added new data source `zia_datacenters` (`getDatacenters`) - Retrieves the list of Zscaler data centers with optional filtering by ID, name, or city.
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Replaced Apache 2.0 license headers with MIT license across all provider and test Go files (143 files).
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Fixed `.goreleaser.yml` - replaced boilerplate references (`pulumi-provider-boilerplate`) with correct provider names (`pulumi-resource-zia`), fixed ldflags version path, removed inapplicable S3 blob upload.
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Fixed `release.yaml` - added `PROVIDER_VERSION` derivation from git tags (SDKs were being published as `1.0.0-alpha.0+dev`), fixed NuGet package path, updated Go version to 1.24, upgraded GitHub Actions to latest versions.
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Improved `ci.yaml` - added provider build, schema generation, unit tests, and SDK build steps (was previously a no-op).
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Disabled `upgrade.yml` workflow (`pulumi-upgrade-provider-action` is not applicable to native providers).
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Fixed Makefile `generate_schema` target referencing undefined `SCHEMA_PATH` variable (corrected to `SCHEMA_FILE`).
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Added `CLAUDE.md` for AI-assisted development guidance (architecture, conventions, build commands, SDK source of truth).
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Added `.cursor/rules/` with Cursor IDE rules for provider conventions, add-resource, add-datasource, and testing workflows.
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Added `.claude/commands/` with Claude Code slash commands for add-resource, add-datasource, and release workflows.

## 1.3.0 (February, 25 2026)

### Notes

- Release date: **(February, 25 2026)**
- Supported Terraform version: **v1.x**

### Internal

- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Migrated provider from the Pulumi Terraform bridged framework to a **native Go provider** built with `pulumi-go-provider`. The provider now communicates directly with the Zscaler API via `zscaler-sdk-go/v3` without a Terraform runtime dependency. All resources, data sources, attributes, and authentication methods remain unchanged.

### Enhancements

- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Added new data source `zia_datacenters` (`getDatacenters`) - Retrieves the list of Zscaler data centers with optional filtering by ID, name, or city.
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Replaced Apache 2.0 license headers with MIT license across all provider and test Go files (143 files).
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Fixed `.goreleaser.yml` - replaced boilerplate references (`pulumi-provider-boilerplate`) with correct provider names (`pulumi-resource-zia`), fixed ldflags version path, removed inapplicable S3 blob upload.
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Fixed `release.yaml` - added `PROVIDER_VERSION` derivation from git tags (SDKs were being published as `1.0.0-alpha.0+dev`), fixed NuGet package path, updated Go version to 1.24, upgraded GitHub Actions to latest versions.
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Improved `ci.yaml` - added provider build, schema generation, unit tests, and SDK build steps (was previously a no-op).
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Disabled `upgrade.yml` workflow (`pulumi-upgrade-provider-action` is not applicable to native providers).
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Fixed Makefile `generate_schema` target referencing undefined `SCHEMA_PATH` variable (corrected to `SCHEMA_FILE`).
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Added `CLAUDE.md` for AI-assisted development guidance (architecture, conventions, build commands, SDK source of truth).
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Added `.cursor/rules/` with Cursor IDE rules for provider conventions, add-resource, add-datasource, and testing workflows.
- [PR #57](https://github.com/zscaler/pulumi-zia/pull/57) - Added `.claude/commands/` with Claude Code slash commands for add-resource, add-datasource, and release workflows.


## 1.2.0 (February, 19 2026)

### Notes

- Release date: **(February, 19 2026)**
- Supported Terraform version: **v1.x**

### Enhancements

- [PR #56](https://github.com/zscaler/pulumi-zia/pull/56) - Added new data source and resource `zia_custom_file_types`
- [PR #56](https://github.com/zscaler/pulumi-zia/pull/56)  - Improved `zia_url_categories` resource READ function for better state refresh and rate limiting conservation.
- [PR #56](https://github.com/zscaler/pulumi-zia/pull/56)  - Added new resource `zia_workload_groups`.
- [PR #56](https://github.com/zscaler/pulumi-zia/pull/56)- Added the following new resources
  - `zia_url_categories_predefined` - Manages predefined URL categories. See [Documentation](https://registry.terraform.io/providers/zscaler/zia/latest/docs/resources/zia_url_categories_predefined) for details.
- [PR #56](https://github.com/zscaler/pulumi-zia/pull/56) - Added new datasource resource `zia_virtual_service_edge_node` - Retrieves the Virtual Service Edge Nodes (VZEN) configured in the ZIA Admin Portal. This data source can be used to set the corresponding node when configuring the resource `zia_virtual_service_edge_cluster`.
- [PR #56](https://github.com/zscaler/pulumi-zia/pull/56) - Added the following new datasources
  - `zia_datacenters` - Retrieves the list of Zscaler data centers (DCs) that can be excluded from service
- [PR #56](https://github.com/zscaler/pulumi-zia/pull/56) - Added the following new datasources and resources:
  - `zia_sub_cloud` - Manage Zscaler Sub-Clouds in ZIA
  - `zia_extranet` - Manage Extranet configurations in ZIA
  - `zia_dc_exclusions` - Manage Extranet configurations in ZIA

## 1.1.1 (June, 24 2025)

### Notes

- Release date: **(June, 24 2025)**
- Supported Terraform version: **v1.x**

### Bug Fixes

- [PR #35](https://github.com/zscaler/pulumi-zia/pull/35) - Upgraded to ZIA Terraform Provider v4.3.2 to address bug related to OneAPI message parsing.

## 1.1.0 (June, 20 2025)

### Notes

- Release date: **(June, 20 2025)**
- Supported Terraform version: **v1.x**

### NEW - RESOURCES, DATA SOURCES

- [PR #32](https://github.com/zscaler/pulumi-zia/pull/32) - The following new resources and data sources have been introduced:

- [PR #32](https://github.com/zscaler/pulumi-zia/pull/32) - Added the datasource and resource ``zia_browser_control_policy`` - Browser Control Policy
- [PR #32](https://github.com/zscaler/pulumi-zia/pull/32) - Added the datasource and resource ``zia_casb_dlp_rules`` - SaaS Security API (Casb DLP Rules)
- [PR #32](https://github.com/zscaler/pulumi-zia/pull/32) - Added the datasource and resource ``zia_casb_malware_rules`` - SaaS Security API (Casb Malware Rules)
- [PR #32](https://github.com/zscaler/pulumi-zia/pull/32) - Added the datasource and resource ``zia_cloud_application_instance`` - Cloud Application Instance
- [PR #32](https://github.com/zscaler/pulumi-zia/pull/32) - Added the datasource and resource ``zia_risk_profiles`` - Risk Profiles

### NEW DATA SOURCES

- [PR #32](https://github.com/zscaler/pulumi-zia/pull/32) - Added the datasource and resource ``zia_casb_tenant`` - SaaS Security API (Casb Tenant)
- [PR #32](https://github.com/zscaler/pulumi-zia/pull/32) - Added the datasource and resource ``zia_casb_email_label`` - SaaS Security API (Casb Email Label)
- [PR #32](https://github.com/zscaler/pulumi-zia/pull/32) - Added the datasource and resource ``zia_casb_tombstone_template`` - SaaS Security API (Casb Quarantine Tombstone Template)
- [PR #32](https://github.com/zscaler/pulumi-zia/pull/32) - Added the datasource and resource ``zia_casb_tombstone_template`` - SaaS Security API (Casb Quarantine Tombstone Template)
- [PR #32](https://github.com/zscaler/pulumi-zia/pull/32) - Added the datasource and resource ``zia_domain_profiles`` - SaaS Security API (Casb Domain Profiles)
- [PR #32](https://github.com/zscaler/pulumi-zia/pull/32) - Added the datasource and resource ``zia_tenant_restriction_profile`` - Tenant Restriction Profile

### Enhancements

- [PR #32](https://github.com/zscaler/pulumi-zia/pull/32) - The resource `zia_location_management` now supports the following attributes.
- `extranet` - The ID of the extranet resource that must be assigned to the location
- `extranetip_pool` - The ID of the traffic selector specified in the extranet
- `extranet_dns` - The ID of the DNS server configuration used in the extranet
- `default_extranet_ts_pool` - A Boolean value indicating that the traffic selector specified in the extranet is the designated default traffic selector
- `default_extranet_dns` - A Boolean value indicating that the DNS server configuration used in the extranet is the designated default DNS server

### Bug Fixes

- [PR #32](https://github.com/zscaler/pulumi-zia/pull/32) - Added validation to ``zia_dlp_web_rules`` to prevent conflict between attributes: `auditor`, `external_auditor_email` and `notification_template`

## 1.0.0 (June, 5 2025) - BREAKING CHANGES

### Notes

- Release date: **(June, 5 2025)**

#### Enhancements - Zscaler OneAPI Support

[PR #31](https://github.com/zscaler/pulumi-zia/pull/31): The ZIA Terraform Provider now offers support for [OneAPI](https://help.zscaler.com/oneapi/understanding-oneapi) Oauth2 authentication through [Zidentity](https://help.zscaler.com/zidentity/what-zidentity).

**NOTE** As of version v1.0.0, this Terraform provider offers backwards compatibility to the Zscaler legacy API framework. This is the recommended authentication method for organizations whose tenants are still not migrated to [Zidentity](https://help.zscaler.com/zidentity/what-zidentity).

⚠️ **WARNING**: Please refer to the [Index Page](https://github.com/zscaler/terraform-provider-zia/blob/master/docs/index.md) page for details on authentication requirements prior to upgrading your provider configuration.

⚠️ **WARNING**: Attention Government customers. OneAPI and Zidentity is not currently supported for the following clouds: `zscalergov` and `zscalerten`. Refer to the [Legacy API Framework](https://github.com/zscaler/terraform-provider-zpa/blob/master/docs/index) section for more information on how authenticate to these environments using the legacy method.

### NEW - RESOURCES, DATA SOURCES, PROPERTIES, ATTRIBUTES, ENV VARS

#### ENV VARS: ZIA Sandbox Submission - BREAKING CHANGES

[PR #31](https://github.com/zscaler/pulumi-zia/pull/31): Authentication to Zscaler Sandbox service now use the following attributes.

- `sandboxToken` - Can also be sourced from the `ZSCALER_SANDBOX_TOKEN` environment variable.
- `sandboxCloud` - Can also be sourced from the `ZSCALER_SANDBOX_CLOUD` environment variable.

The use of the previous envioronment variables combination `ZIA_SANDBOX_TOKEN` and `ZIA_CLOUD` is now deprecated.

### NEW - RESOURCES, DATA SOURCES

[PR #31](https://github.com/zscaler/pulumi-zia/pull/31): The following new resources and data sources have been introduced:

- Added the datasource and resource ``zia_sandbox_rules`` [PR #31](https://github.com/zscaler/pulumi-zia/pull/31) :rocket: - Manage Sandbox Rules
- Added the datasource and resource ``zia_firewall_dns_rule``[PR #31](https://github.com/zscaler/pulumi-zia/pull/31) :rocket: - Manage Cloud Firewall DNS Rules
- Added the datasource and resource ``zia_firewall_ips_rule`` [PR #31](https://github.com/zscaler/pulumi-zia/pull/31) :rocket: - Manage Cloud Firewall IPS Rules
- Added the datasource and resource ``zia_file_type_control_rules`` [PR #31](https://github.com/zscaler/pulumi-zia/pull/31) :rocket: - Manage File Type Control Rules
- Added the datasource and resource ``zia_advanced_threat_settings`` [PR #31](https://github.com/zscaler/pulumi-zia/pull/31) :rocket: - Manages advanced threat configuration settings
- Added the datasource and resource ``zia_atp_malicious_urls`` [PR #31](https://github.com/zscaler/pulumi-zia/pull/31) :rocket: - Manages malicious URLs added to the denylist in ATP policy
- Added the datasource and resource ``zia_atp_security_exceptions`` [PR #31](https://github.com/zscaler/pulumi-zia/pull/31) :rocket: - Manages Security Exceptions (URL Bypass List) for the ATP policy
- Added the datasource and resource ``zia_advanced_settings`` [PR #31](https://github.com/zscaler/pulumi-zia/pull/31) :rocket: - Manages Advanced Settings configuration. [Configuring Advanced Settings](https://help.zscaler.com/zia/configuring-advanced-settings)
- Added the datasource and resource ``zia_atp_malware_inspection`` [PR #31](https://github.com/zscaler/pulumi-zia/pull/31) :rocket: - Manages Advanced Threat Protection Malware Inspection configuration. [Malware Protection](https://help.zscaler.com/zia/policies/malware-protection)
- Added the datasource and resource ``zia_atp_malware_protocols`` [PR #31](https://github.com/zscaler/pulumi-zia/pull/31) :rocket: - Manages Advanced Threat Protection Malware Protocols configuration. [Malware Protection](https://help.zscaler.com/zia/policies/malware-protection)
- Added the datasource and resource ``zia_atp_malware_settings`` [PR #31](https://github.com/zscaler/pulumi-zia/pull/31) :rocket: - Manages Advanced Threat Protection Malware Settings. [Malware Protection](https://help.zscaler.com/zia/policies/malware-protection)
- Added the datasource and resource ``zia_atp_malware_policy`` [PR #31](https://github.com/zscaler/pulumi-zia/pull/31) :rocket: - Manages Advanced Threat Protection Malware Policy. [Malware Protection](https://help.zscaler.com/zia/policies/malware-protection)
- Added the datasource and resource ``zia_end_user_notification`` [PR #31](https://github.com/zscaler/pulumi-zia/pull/31) :rocket: - Retrieves information of browser-based end user notification (EUN) configuration details.[Understanding Browser-Based End User Notifications](https://help.zscaler.com/zia/understanding-browser-based-end-user-notifications)
- Added the datasource and resource ``zia_url_filtering_and_cloud_app_settings`` [PR #31](https://github.com/zscaler/pulumi-zia/pull/31) :rocket: - Manages the URL and Cloud App Control advanced policy settings.[Configuring Advanced Policy Settings](https://help.zscaler.com/zia/configuring-advanced-policy-settings)
- Added the datasource ``zia_cloud_applications`` [PR #31](https://github.com/zscaler/pulumi-zia/pull/31) :rocket: - Retrieves Predefined and User Defined Cloud Applications associated with the DLP rules, Cloud App Control rules, Advanced Settings, Bandwidth Classes, File Type Control rules, and SSL Inspection rules.
- Added the datasource ``zia_forwarding_control_proxy_gateway`` [PR #31](https://github.com/zscaler/pulumi-zia/pull/31) :rocket: - Retrieves information of existing Proxy Gateway configuration.
- Added the datasource and resource ``zia_ssl_inspection_rules`` [PR #31](https://github.com/zscaler/pulumi-zia/pull/31) :rocket: - Manages SSL Inspection Rules.

## 0.0.1 (March 27, 2024)

### Notes

- Release date: **(March 27, 2024)**

🎉 **Initial Release** 🎉
