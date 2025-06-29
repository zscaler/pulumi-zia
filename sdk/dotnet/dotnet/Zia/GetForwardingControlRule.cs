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
    public static class GetForwardingControlRule
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/configuring-forwarding-policy)
        /// * [API documentation](https://help.zscaler.com/zia/forwarding-control-policy#/forwardingRules-get)
        /// 
        /// Use the **zia_forwarding_control_rule** data source to get information about a forwarding control rule which is used to forward selective Zscaler traffic to specific destinations based on your needs.For example, if you want to forward specific web traffic to a third-party proxy service or if you want to forward source IP anchored application traffic to a specific Zscaler Private Access (ZPA) App Connector or internal application traffic through ZIA threat and data protection engines, use forwarding control by configuring appropriate rules.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// # ZIA Forwarding Control - ZPA Gateway
        /// data "zia_forwarding_control_rule" "this" {
        ///   name = "FWD_RULE01"
        /// }
        /// ```
        /// </summary>
        public static Task<GetForwardingControlRuleResult> InvokeAsync(GetForwardingControlRuleArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetForwardingControlRuleResult>("zia:index/getForwardingControlRule:getForwardingControlRule", args ?? new GetForwardingControlRuleArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/configuring-forwarding-policy)
        /// * [API documentation](https://help.zscaler.com/zia/forwarding-control-policy#/forwardingRules-get)
        /// 
        /// Use the **zia_forwarding_control_rule** data source to get information about a forwarding control rule which is used to forward selective Zscaler traffic to specific destinations based on your needs.For example, if you want to forward specific web traffic to a third-party proxy service or if you want to forward source IP anchored application traffic to a specific Zscaler Private Access (ZPA) App Connector or internal application traffic through ZIA threat and data protection engines, use forwarding control by configuring appropriate rules.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// # ZIA Forwarding Control - ZPA Gateway
        /// data "zia_forwarding_control_rule" "this" {
        ///   name = "FWD_RULE01"
        /// }
        /// ```
        /// </summary>
        public static Output<GetForwardingControlRuleResult> Invoke(GetForwardingControlRuleInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetForwardingControlRuleResult>("zia:index/getForwardingControlRule:getForwardingControlRule", args ?? new GetForwardingControlRuleInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/configuring-forwarding-policy)
        /// * [API documentation](https://help.zscaler.com/zia/forwarding-control-policy#/forwardingRules-get)
        /// 
        /// Use the **zia_forwarding_control_rule** data source to get information about a forwarding control rule which is used to forward selective Zscaler traffic to specific destinations based on your needs.For example, if you want to forward specific web traffic to a third-party proxy service or if you want to forward source IP anchored application traffic to a specific Zscaler Private Access (ZPA) App Connector or internal application traffic through ZIA threat and data protection engines, use forwarding control by configuring appropriate rules.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// # ZIA Forwarding Control - ZPA Gateway
        /// data "zia_forwarding_control_rule" "this" {
        ///   name = "FWD_RULE01"
        /// }
        /// ```
        /// </summary>
        public static Output<GetForwardingControlRuleResult> Invoke(GetForwardingControlRuleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetForwardingControlRuleResult>("zia:index/getForwardingControlRule:getForwardingControlRule", args ?? new GetForwardingControlRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetForwardingControlRuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier assigned to the forwarding rule.
        /// </summary>
        [Input("id")]
        public int? Id { get; set; }

        /// <summary>
        /// The name of the forwarding rule.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// (string) -  The rule type selected from the available options
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetForwardingControlRuleArgs()
        {
        }
        public static new GetForwardingControlRuleArgs Empty => new GetForwardingControlRuleArgs();
    }

    public sealed class GetForwardingControlRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier assigned to the forwarding rule.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// The name of the forwarding rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (string) -  The rule type selected from the available options
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetForwardingControlRuleInvokeArgs()
        {
        }
        public static new GetForwardingControlRuleInvokeArgs Empty => new GetForwardingControlRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetForwardingControlRuleResult
    {
        /// <summary>
        /// (list) Apply to any number of departments When not used it implies `Any` to apply the rule to all departments.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleDepartmentResult> Departments;
        /// <summary>
        /// (string) - Additional information about the forwarding rule
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// ** - (list) -  IP addresses and fully qualified domain names (FQDNs), if the domain has multiple destination IP addresses or if its IP addresses may change. For IP addresses, you can enter individual IP addresses, subnets, or address ranges. If adding multiple items, hit Enter after each entry.
        /// </summary>
        public readonly ImmutableArray<string> DestAddresses;
        /// <summary>
        /// ** - (list) estination countries for which the rule is applicable. If not set, the rule is not restricted to specific destination countries. Provide a 2 letter [ISO3166 Alpha2 Country code](https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes).
        /// </summary>
        public readonly ImmutableArray<string> DestCountries;
        /// <summary>
        /// ** - (list) identify destinations based on the URL category of the domain, select Any to apply the rule to all categories or select the specific categories you want to control.
        /// </summary>
        public readonly ImmutableArray<string> DestIpCategories;
        /// <summary>
        /// ** - (list) Any number of destination IP address groups that you want to control with this rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleDestIpGroupResult> DestIpGroups;
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleDestIpv6GroupResult> DestIpv6Groups;
        /// <summary>
        /// (list) Name-ID pairs of device groups for which the rule must be applied. This field is applicable for devices that are managed using Zscaler Client Connector. If no value is set, this field is ignored during the policy evaluation.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleDeviceGroupResult> DeviceGroups;
        /// <summary>
        /// (list) Name-ID pairs of devices for which the rule must be applied. Specifies devices that are managed using Zscaler Client Connector. If no value is set, this field is ignored during the policy evaluation.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleDeviceResult> Devices;
        /// <summary>
        /// (list) - Name-ID pairs of the Zscaler Cloud Connector groups to which the forwarding rule applies
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleEcGroupResult> EcGroups;
        /// <summary>
        /// (string) - The type of traffic forwarding method selected from the available options.
        /// </summary>
        public readonly string ForwardMethod;
        /// <summary>
        /// (list) You can manually select up to `8` groups. When not used it implies `Any` to apply the rule to all groups.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleGroupResult> Groups;
        /// <summary>
        /// (int) Identifier that uniquely identifies an entity
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// (list) Labels that are applicable to the rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleLabelResult> Labels;
        /// <summary>
        /// (Optional) You can manually select up to `32` location groups. When not used it implies `Any` to apply the rule to all location groups.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleLocationGroupResult> LocationGroups;
        /// <summary>
        /// (Optional) You can manually select up to `8` locations. When not used it implies `Any` to apply the rule to all groups.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleLocationResult> Locations;
        /// <summary>
        /// (string) The configured name of the entity
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// (list) Any number of application groups that you want to control with this rule. The service provides predefined applications that you can group, but not modify
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleNwApplicationGroupResult> NwApplicationGroups;
        /// <summary>
        /// (Optional) When not used it applies the rule to all applications. The service provides predefined applications, which you can group, but not modify.
        /// </summary>
        public readonly ImmutableArray<string> NwApplications;
        /// <summary>
        /// (list) Any number of predefined or custom network service groups to which the rule applies.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleNwServiceGroupResult> NwServiceGroups;
        /// <summary>
        /// (list) When not used it applies the rule to all network services or you can select specific network services. The Zscaler firewall has predefined services and you can configure up to `1,024` additional custom services.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleNwServiceResult> NwServices;
        /// <summary>
        /// (string) - The order of execution for the forwarding rule order.
        /// </summary>
        public readonly int Order;
        /// <summary>
        /// (set) The proxy gateway for which the rule is applicable. This field is applicable only for the `PROXYCHAIN` forwarding method.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleProxyGatewayResult> ProxyGateways;
        public readonly int Rank;
        /// <summary>
        /// ** - (list) List of destination domain categories to which the rule applies.
        /// </summary>
        public readonly ImmutableArray<string> ResCategories;
        /// <summary>
        /// (list) Any number of source IP address groups that you want to control with this rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleSrcIpGroupResult> SrcIpGroups;
        /// <summary>
        /// (Optional) You can enter individual IP addresses, subnets, or address ranges.
        /// </summary>
        public readonly ImmutableArray<string> SrcIps;
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleSrcIpv6GroupResult> SrcIpv6Groups;
        /// <summary>
        /// (string) - Indicates whether the forwarding rule is enabled or disabled.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// (string) -  The rule type selected from the available options
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// (list) You can manually select up to `4` general and/or special users. When not used it implies `Any` to apply the rule to all users.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleUserResult> Users;
        /// <summary>
        /// (set) The list of ZPA Application Segments for which this rule is applicable. This field is applicable only for the `ZPA` Gateway forwarding method.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleZpaAppSegmentResult> ZpaAppSegments;
        /// <summary>
        /// (set) List of ZPA Application Segment Groups for which this rule is applicable. This field is applicable only for the `ECZPA` forwarding method (used for Zscaler Cloud Connector).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleZpaApplicationSegmentGroupResult> ZpaApplicationSegmentGroups;
        /// <summary>
        /// (set) List of ZPA Application Segments for which this rule is applicable. This field is applicable only for the `ECZPA` forwarding method (used for Zscaler Cloud Connector).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleZpaApplicationSegmentResult> ZpaApplicationSegments;
        public readonly bool ZpaBrokerRule;
        /// <summary>
        /// (set) The ZPA Gateway for which this rule is applicable. This field is applicable only for the `ZPA` forwarding method.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlRuleZpaGatewayResult> ZpaGateways;

        [OutputConstructor]
        private GetForwardingControlRuleResult(
            ImmutableArray<Outputs.GetForwardingControlRuleDepartmentResult> departments,

            string description,

            ImmutableArray<string> destAddresses,

            ImmutableArray<string> destCountries,

            ImmutableArray<string> destIpCategories,

            ImmutableArray<Outputs.GetForwardingControlRuleDestIpGroupResult> destIpGroups,

            ImmutableArray<Outputs.GetForwardingControlRuleDestIpv6GroupResult> destIpv6Groups,

            ImmutableArray<Outputs.GetForwardingControlRuleDeviceGroupResult> deviceGroups,

            ImmutableArray<Outputs.GetForwardingControlRuleDeviceResult> devices,

            ImmutableArray<Outputs.GetForwardingControlRuleEcGroupResult> ecGroups,

            string forwardMethod,

            ImmutableArray<Outputs.GetForwardingControlRuleGroupResult> groups,

            int? id,

            ImmutableArray<Outputs.GetForwardingControlRuleLabelResult> labels,

            ImmutableArray<Outputs.GetForwardingControlRuleLocationGroupResult> locationGroups,

            ImmutableArray<Outputs.GetForwardingControlRuleLocationResult> locations,

            string? name,

            ImmutableArray<Outputs.GetForwardingControlRuleNwApplicationGroupResult> nwApplicationGroups,

            ImmutableArray<string> nwApplications,

            ImmutableArray<Outputs.GetForwardingControlRuleNwServiceGroupResult> nwServiceGroups,

            ImmutableArray<Outputs.GetForwardingControlRuleNwServiceResult> nwServices,

            int order,

            ImmutableArray<Outputs.GetForwardingControlRuleProxyGatewayResult> proxyGateways,

            int rank,

            ImmutableArray<string> resCategories,

            ImmutableArray<Outputs.GetForwardingControlRuleSrcIpGroupResult> srcIpGroups,

            ImmutableArray<string> srcIps,

            ImmutableArray<Outputs.GetForwardingControlRuleSrcIpv6GroupResult> srcIpv6Groups,

            string state,

            string? type,

            ImmutableArray<Outputs.GetForwardingControlRuleUserResult> users,

            ImmutableArray<Outputs.GetForwardingControlRuleZpaAppSegmentResult> zpaAppSegments,

            ImmutableArray<Outputs.GetForwardingControlRuleZpaApplicationSegmentGroupResult> zpaApplicationSegmentGroups,

            ImmutableArray<Outputs.GetForwardingControlRuleZpaApplicationSegmentResult> zpaApplicationSegments,

            bool zpaBrokerRule,

            ImmutableArray<Outputs.GetForwardingControlRuleZpaGatewayResult> zpaGateways)
        {
            Departments = departments;
            Description = description;
            DestAddresses = destAddresses;
            DestCountries = destCountries;
            DestIpCategories = destIpCategories;
            DestIpGroups = destIpGroups;
            DestIpv6Groups = destIpv6Groups;
            DeviceGroups = deviceGroups;
            Devices = devices;
            EcGroups = ecGroups;
            ForwardMethod = forwardMethod;
            Groups = groups;
            Id = id;
            Labels = labels;
            LocationGroups = locationGroups;
            Locations = locations;
            Name = name;
            NwApplicationGroups = nwApplicationGroups;
            NwApplications = nwApplications;
            NwServiceGroups = nwServiceGroups;
            NwServices = nwServices;
            Order = order;
            ProxyGateways = proxyGateways;
            Rank = rank;
            ResCategories = resCategories;
            SrcIpGroups = srcIpGroups;
            SrcIps = srcIps;
            SrcIpv6Groups = srcIpv6Groups;
            State = state;
            Type = type;
            Users = users;
            ZpaAppSegments = zpaAppSegments;
            ZpaApplicationSegmentGroups = zpaApplicationSegmentGroups;
            ZpaApplicationSegments = zpaApplicationSegments;
            ZpaBrokerRule = zpaBrokerRule;
            ZpaGateways = zpaGateways;
        }
    }
}
