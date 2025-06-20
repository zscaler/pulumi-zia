# Changelog

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
