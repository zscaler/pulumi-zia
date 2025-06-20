// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zia
{
    public static class GetIPSFirewallRule
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/ips-control-policy#/firewallIpsRules-get)
        /// * [API documentation](https://help.zscaler.com/zia/configuring-ips-control-policy)
        /// 
        /// Use the **zia_firewall_ips_rule** data source to get information about a cloud firewall IPS rule available in the Zscaler Internet Access.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// # ZIA Firewall IPS Rule by name
        /// data "zia_firewall_ips_rule" "this" {
        ///     name = "Default Cloud IPS Rule"
        /// }
        /// ```
        /// 
        /// ```hcl
        /// # ZIA Firewall IPS Rule by ID
        /// data "zia_firewall_ips_rule" "this" {
        ///     id = "12365478"
        /// }
        /// ```
        /// </summary>
        public static Task<GetIPSFirewallRuleResult> InvokeAsync(GetIPSFirewallRuleArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIPSFirewallRuleResult>("zia:index/getIPSFirewallRule:getIPSFirewallRule", args ?? new GetIPSFirewallRuleArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/ips-control-policy#/firewallIpsRules-get)
        /// * [API documentation](https://help.zscaler.com/zia/configuring-ips-control-policy)
        /// 
        /// Use the **zia_firewall_ips_rule** data source to get information about a cloud firewall IPS rule available in the Zscaler Internet Access.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// # ZIA Firewall IPS Rule by name
        /// data "zia_firewall_ips_rule" "this" {
        ///     name = "Default Cloud IPS Rule"
        /// }
        /// ```
        /// 
        /// ```hcl
        /// # ZIA Firewall IPS Rule by ID
        /// data "zia_firewall_ips_rule" "this" {
        ///     id = "12365478"
        /// }
        /// ```
        /// </summary>
        public static Output<GetIPSFirewallRuleResult> Invoke(GetIPSFirewallRuleInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIPSFirewallRuleResult>("zia:index/getIPSFirewallRule:getIPSFirewallRule", args ?? new GetIPSFirewallRuleInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/ips-control-policy#/firewallIpsRules-get)
        /// * [API documentation](https://help.zscaler.com/zia/configuring-ips-control-policy)
        /// 
        /// Use the **zia_firewall_ips_rule** data source to get information about a cloud firewall IPS rule available in the Zscaler Internet Access.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// # ZIA Firewall IPS Rule by name
        /// data "zia_firewall_ips_rule" "this" {
        ///     name = "Default Cloud IPS Rule"
        /// }
        /// ```
        /// 
        /// ```hcl
        /// # ZIA Firewall IPS Rule by ID
        /// data "zia_firewall_ips_rule" "this" {
        ///     id = "12365478"
        /// }
        /// ```
        /// </summary>
        public static Output<GetIPSFirewallRuleResult> Invoke(GetIPSFirewallRuleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIPSFirewallRuleResult>("zia:index/getIPSFirewallRule:getIPSFirewallRule", args ?? new GetIPSFirewallRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIPSFirewallRuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for the Firewall Filtering policy rule
        /// </summary>
        [Input("id")]
        public int? Id { get; set; }

        /// <summary>
        /// Name of the Firewall Filtering policy rule
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetIPSFirewallRuleArgs()
        {
        }
        public static new GetIPSFirewallRuleArgs Empty => new GetIPSFirewallRuleArgs();
    }

    public sealed class GetIPSFirewallRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for the Firewall Filtering policy rule
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Name of the Firewall Filtering policy rule
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetIPSFirewallRuleInvokeArgs()
        {
        }
        public static new GetIPSFirewallRuleInvokeArgs Empty => new GetIPSFirewallRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetIPSFirewallRuleResult
    {
        /// <summary>
        /// (String) The action configured for the rule that must take place if the traffic matches the rule criteria, such as allowing or blocking the traffic or bypassing the rule. The following actions are accepted: `ALLOW`, `BLOCK_DROP`, `BLOCK_RESET`, `BYPASS_IPS`
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// (Boolean) Value that indicates whether packet capture (PCAP) is enabled or not
        /// </summary>
        public readonly bool CapturePcap;
        /// <summary>
        /// (Boolean) Value that indicates whether the rule is the Default Cloud IPS Rule or not
        /// </summary>
        public readonly bool DefaultRule;
        /// <summary>
        /// (List of Objects) Apply to any number of departments When not used it implies `Any` to apply the rule to all departments.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleDepartmentResult> Departments;
        /// <summary>
        /// (String) Enter additional notes or information. The description cannot exceed 10,240 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (Set of String) Destination IP addresses or FQDNs to which the rule applies. If not set, the rule is not restricted to a specific destination IP address. Each IP entry can be a single IP address, CIDR (e.g., 10.10.33.0/24), or an IP range (e.g., 10.10.33.1-10.10.33.10).
        /// </summary>
        public readonly ImmutableArray<string> DestAddresses;
        /// <summary>
        /// (Set of String) Identify destinations based on the location of a server, select Any to apply the rule to all countries or select the countries to which you want to control traffic.
        /// **NOTE**: Provide a 2 letter [ISO3166 Alpha2 Country code](https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes). i.e ``"US"``, ``"CA"``
        /// </summary>
        public readonly ImmutableArray<string> DestCountries;
        /// <summary>
        /// (Set of String)  identify destinations based on the URL category of the domain, select Any to apply the rule to all categories or select the specific categories you want to control.
        /// </summary>
        public readonly ImmutableArray<string> DestIpCategories;
        /// <summary>
        /// ** - (List of Objects) Any number of destination IP address groups that you want to control with this rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleDestIpGroupResult> DestIpGroups;
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleDestIpv6GroupResult> DestIpv6Groups;
        /// <summary>
        /// (List of Objects) Device groups to which the rule applies. This field is applicable for devices that are managed using Zscaler Client Connector. If no value is set, this field is ignored during the policy evaluation.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleDeviceGroupResult> DeviceGroups;
        /// <summary>
        /// (List of Objects) Devices to which the rule applies. This field is applicable for devices that are managed using Zscaler Client Connector. If no value is set, this field is ignored during the policy evaluation.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleDeviceResult> Devices;
        /// <summary>
        /// (Integer) A Boolean value that indicates whether full logging is enabled. A true value indicates that full logging is enabled, whereas a false value indicates that aggregate logging is enabled.
        /// </summary>
        public readonly bool EnableFullLogging;
        /// <summary>
        /// (List of Objects) You can manually select up to `8` groups. When not used it implies `Any` to apply the rule to all groups.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleGroupResult> Groups;
        /// <summary>
        /// (Integer) Identifier that uniquely identifies an entity
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// (List of Objects) Labels that are applicable to the rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleLabelResult> Labels;
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleLastModifiedByResult> LastModifiedBies;
        public readonly int LastModifiedTime;
        /// <summary>
        /// (List of Objects)You can manually select up to `32` location groups. When not used it implies `Any` to apply the rule to all location groups.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleLocationGroupResult> LocationGroups;
        /// <summary>
        /// (List of Objects) You can manually select up to `8` locations. When not used it implies `Any` to apply the rule to all groups.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleLocationResult> Locations;
        /// <summary>
        /// (String) The configured name of the entity
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// (List of Objects) Any number of predefined or custom network service groups to which the rule applies.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleNwServiceGroupResult> NwServiceGroups;
        /// <summary>
        /// (List of Objects) When not used it applies the rule to all network services or you can select specific network services. The Zscaler firewall has predefined services and you can configure up to `1,024` additional custom services.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleNwServiceResult> NwServices;
        /// <summary>
        /// (Integer) Policy rules are evaluated in ascending numerical order (Rule 1 before Rule 2, and so on), and the Rule Order reflects this rule's place in the order.
        /// </summary>
        public readonly int Order;
        /// <summary>
        /// (Boolean) A Boolean field that indicates that the rule is predefined by using a true value
        /// </summary>
        public readonly bool Predefined;
        /// <summary>
        /// (Integer) By default, the admin ranking is disabled. To use this feature, you must enable admin rank. The default value is `7`.
        /// </summary>
        public readonly int Rank;
        /// <summary>
        /// (Set of String) URL categories associated with resolved IP addresses to which the rule applies. If not set, the rule is not restricted to a specific URL category.
        /// </summary>
        public readonly ImmutableArray<string> ResCategories;
        /// <summary>
        /// (Set of String) The countries of origin of traffic for which the rule is applicable. If not set, the rule is not restricted to specific source countries.
        /// **NOTE**: Provide a 2 letter [ISO3166 Alpha2 Country code](https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes). i.e ``"US"``, ``"CA"``
        /// </summary>
        public readonly ImmutableArray<string> SourceCountries;
        /// <summary>
        /// (List of Objects)Source IP address groups for which the rule is applicable. If not set, the rule is not restricted to a specific source IP address group.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleSrcIpGroupResult> SrcIpGroups;
        /// <summary>
        /// (Set of String) Source IP addresses or FQDNs to which the rule applies. If not set, the rule is not restricted to a specific source IP address. Each IP entry can be a single IP address, CIDR (e.g., 10.10.33.0/24), or an IP range (e.g., 10.10.33.1-10.10.33.10).
        /// </summary>
        public readonly ImmutableArray<string> SrcIps;
        /// <summary>
        /// (List of Objects) Source IPv6 address groups for which the rule is applicable. If not set, the rule is not restricted to a specific source IPv6 address group.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleSrcIpv6GroupResult> SrcIpv6Groups;
        /// <summary>
        /// (String) An enabled rule is actively enforced. A disabled rule is not actively enforced but does not lose its place in the Rule Order. The service skips it and moves to the next rule.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// (List of Objects) Advanced threat categories to which the rule applies
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleThreatCategoryResult> ThreatCategories;
        /// <summary>
        /// (List of Objects) You can manually select up to `1` time intervals. When not used it implies `always` to apply the rule to all time intervals.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleTimeWindowResult> TimeWindows;
        /// <summary>
        /// (List of Objects) You can manually select up to `4` general and/or special users. When not used it implies `Any` to apply the rule to all users.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleUserResult> Users;
        /// <summary>
        /// (List of Objects) The ZPA application segments to which the rule applies
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIPSFirewallRuleZpaAppSegmentResult> ZpaAppSegments;

        [OutputConstructor]
        private GetIPSFirewallRuleResult(
            string action,

            bool capturePcap,

            bool defaultRule,

            ImmutableArray<Outputs.GetIPSFirewallRuleDepartmentResult> departments,

            string description,

            ImmutableArray<string> destAddresses,

            ImmutableArray<string> destCountries,

            ImmutableArray<string> destIpCategories,

            ImmutableArray<Outputs.GetIPSFirewallRuleDestIpGroupResult> destIpGroups,

            ImmutableArray<Outputs.GetIPSFirewallRuleDestIpv6GroupResult> destIpv6Groups,

            ImmutableArray<Outputs.GetIPSFirewallRuleDeviceGroupResult> deviceGroups,

            ImmutableArray<Outputs.GetIPSFirewallRuleDeviceResult> devices,

            bool enableFullLogging,

            ImmutableArray<Outputs.GetIPSFirewallRuleGroupResult> groups,

            int id,

            ImmutableArray<Outputs.GetIPSFirewallRuleLabelResult> labels,

            ImmutableArray<Outputs.GetIPSFirewallRuleLastModifiedByResult> lastModifiedBies,

            int lastModifiedTime,

            ImmutableArray<Outputs.GetIPSFirewallRuleLocationGroupResult> locationGroups,

            ImmutableArray<Outputs.GetIPSFirewallRuleLocationResult> locations,

            string name,

            ImmutableArray<Outputs.GetIPSFirewallRuleNwServiceGroupResult> nwServiceGroups,

            ImmutableArray<Outputs.GetIPSFirewallRuleNwServiceResult> nwServices,

            int order,

            bool predefined,

            int rank,

            ImmutableArray<string> resCategories,

            ImmutableArray<string> sourceCountries,

            ImmutableArray<Outputs.GetIPSFirewallRuleSrcIpGroupResult> srcIpGroups,

            ImmutableArray<string> srcIps,

            ImmutableArray<Outputs.GetIPSFirewallRuleSrcIpv6GroupResult> srcIpv6Groups,

            string state,

            ImmutableArray<Outputs.GetIPSFirewallRuleThreatCategoryResult> threatCategories,

            ImmutableArray<Outputs.GetIPSFirewallRuleTimeWindowResult> timeWindows,

            ImmutableArray<Outputs.GetIPSFirewallRuleUserResult> users,

            ImmutableArray<Outputs.GetIPSFirewallRuleZpaAppSegmentResult> zpaAppSegments)
        {
            Action = action;
            CapturePcap = capturePcap;
            DefaultRule = defaultRule;
            Departments = departments;
            Description = description;
            DestAddresses = destAddresses;
            DestCountries = destCountries;
            DestIpCategories = destIpCategories;
            DestIpGroups = destIpGroups;
            DestIpv6Groups = destIpv6Groups;
            DeviceGroups = deviceGroups;
            Devices = devices;
            EnableFullLogging = enableFullLogging;
            Groups = groups;
            Id = id;
            Labels = labels;
            LastModifiedBies = lastModifiedBies;
            LastModifiedTime = lastModifiedTime;
            LocationGroups = locationGroups;
            Locations = locations;
            Name = name;
            NwServiceGroups = nwServiceGroups;
            NwServices = nwServices;
            Order = order;
            Predefined = predefined;
            Rank = rank;
            ResCategories = resCategories;
            SourceCountries = sourceCountries;
            SrcIpGroups = srcIpGroups;
            SrcIps = srcIps;
            SrcIpv6Groups = srcIpv6Groups;
            State = state;
            ThreatCategories = threatCategories;
            TimeWindows = timeWindows;
            Users = users;
            ZpaAppSegments = zpaAppSegments;
        }
    }
}
