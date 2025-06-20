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
    /// <summary>
    /// * [Official documentation](https://help.zscaler.com/zia/about-nat-control)
    /// * [API documentation](https://help.zscaler.com/zia/nat-control-policy#/dnatRules-get)
    /// 
    /// The **zia_nat_control_rules** resource allows the creation and management of NAT Control rules in the Zscaler Internet Access.
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.
    /// 
    /// Visit
    /// 
    /// **zia_nat_control_rules** can be imported by using `&lt;RULE ID&gt;` or `&lt;RULE NAME&gt;` as the import ID.
    /// 
    /// For example:
    /// 
    /// ```sh
    /// $ pulumi import zia:index/natControlRules:NatControlRules example &lt;rule_id&gt;
    /// ```
    /// 
    /// or
    /// 
    /// ```sh
    /// $ pulumi import zia:index/natControlRules:NatControlRules example &lt;rule_name&gt;
    /// ```
    /// </summary>
    [ZiaResourceType("zia:index/natControlRules:NatControlRules")]
    public partial class NatControlRules : global::Pulumi.CustomResource
    {
        /// <summary>
        /// If set to true, the default rule is applied
        /// </summary>
        [Output("defaultRule")]
        public Output<bool?> DefaultRule { get; private set; } = null!;

        /// <summary>
        /// list of departments for which rule must be applied
        /// </summary>
        [Output("departments")]
        public Output<Outputs.NatControlRulesDepartments?> Departments { get; private set; } = null!;

        /// <summary>
        /// Additional information about the rule
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Destination addresses. Supports IPv4, FQDNs, or wildcard FQDNs
        /// </summary>
        [Output("destAddresses")]
        public Output<ImmutableArray<string>> DestAddresses { get; private set; } = null!;

        /// <summary>
        /// Destination countries for which the rule is applicable. If not set, the rule is not restricted to specific destination
        /// countries.
        /// </summary>
        [Output("destCountries")]
        public Output<ImmutableArray<string>> DestCountries { get; private set; } = null!;

        [Output("destIpCategories")]
        public Output<ImmutableArray<string>> DestIpCategories { get; private set; } = null!;

        /// <summary>
        /// list of destination ip groups
        /// </summary>
        [Output("destIpGroups")]
        public Output<Outputs.NatControlRulesDestIpGroups?> DestIpGroups { get; private set; } = null!;

        /// <summary>
        /// list of destination ipv6 groups
        /// </summary>
        [Output("destIpv6Groups")]
        public Output<Outputs.NatControlRulesDestIpv6Groups?> DestIpv6Groups { get; private set; } = null!;

        /// <summary>
        /// This field is applicable for devices that are managed using Zscaler Client Connector.
        /// </summary>
        [Output("deviceGroups")]
        public Output<Outputs.NatControlRulesDeviceGroups?> DeviceGroups { get; private set; } = null!;

        /// <summary>
        /// Name-ID pairs of devices for which rule must be applied.
        /// </summary>
        [Output("devices")]
        public Output<Outputs.NatControlRulesDevices?> Devices { get; private set; } = null!;

        [Output("enableFullLogging")]
        public Output<bool?> EnableFullLogging { get; private set; } = null!;

        /// <summary>
        /// list of groups for which rule must be applied
        /// </summary>
        [Output("groups")]
        public Output<Outputs.NatControlRulesGroups?> Groups { get; private set; } = null!;

        /// <summary>
        /// list of Labels that are applicable to the rule.
        /// </summary>
        [Output("labels")]
        public Output<Outputs.NatControlRulesLabels?> Labels { get; private set; } = null!;

        /// <summary>
        /// list of locations groups
        /// </summary>
        [Output("locationGroups")]
        public Output<Outputs.NatControlRulesLocationGroups?> LocationGroups { get; private set; } = null!;

        /// <summary>
        /// list of locations for which rule must be applied
        /// </summary>
        [Output("locations")]
        public Output<Outputs.NatControlRulesLocations?> Locations { get; private set; } = null!;

        /// <summary>
        /// Name of the nat control policy rule
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// list of nw service groups
        /// </summary>
        [Output("nwServiceGroups")]
        public Output<Outputs.NatControlRulesNwServiceGroups?> NwServiceGroups { get; private set; } = null!;

        /// <summary>
        /// list of nw services
        /// </summary>
        [Output("nwServices")]
        public Output<Outputs.NatControlRulesNwServices?> NwServices { get; private set; } = null!;

        /// <summary>
        /// Rule order number. If omitted, the rule will be added to the end of the rule set.
        /// </summary>
        [Output("order")]
        public Output<int> Order { get; private set; } = null!;

        /// <summary>
        /// If set to true, a predefined rule is applied
        /// </summary>
        [Output("predefined")]
        public Output<bool?> Predefined { get; private set; } = null!;

        /// <summary>
        /// Admin rank of the nat control policy rule
        /// </summary>
        [Output("rank")]
        public Output<int?> Rank { get; private set; } = null!;

        /// <summary>
        /// The action the nat control policy rule takes when packets match the rule
        /// </summary>
        [Output("redirectFqdn")]
        public Output<string?> RedirectFqdn { get; private set; } = null!;

        /// <summary>
        /// The action the nat control policy rule takes when packets match the rule
        /// </summary>
        [Output("redirectIp")]
        public Output<string?> RedirectIp { get; private set; } = null!;

        /// <summary>
        /// The action the nat control policy rule takes when packets match the rule
        /// </summary>
        [Output("redirectPort")]
        public Output<int?> RedirectPort { get; private set; } = null!;

        /// <summary>
        /// List of destination domain categories to which the rule applies
        /// </summary>
        [Output("resCategories")]
        public Output<ImmutableArray<string>> ResCategories { get; private set; } = null!;

        [Output("ruleId")]
        public Output<int> RuleId { get; private set; } = null!;

        /// <summary>
        /// list of source ip groups
        /// </summary>
        [Output("srcIpGroups")]
        public Output<Outputs.NatControlRulesSrcIpGroups?> SrcIpGroups { get; private set; } = null!;

        /// <summary>
        /// User-defined source IP addresses for which the rule is applicable. If not set, the rule is not restricted to a specific
        /// source IP address.
        /// </summary>
        [Output("srcIps")]
        public Output<ImmutableArray<string>> SrcIps { get; private set; } = null!;

        /// <summary>
        /// list of source ipv6 groups
        /// </summary>
        [Output("srcIpv6Groups")]
        public Output<Outputs.NatControlRulesSrcIpv6Groups?> SrcIpv6Groups { get; private set; } = null!;

        /// <summary>
        /// Determines whether the nat control policy rule is enabled or disabled
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        /// <summary>
        /// The time interval in which the nat control policy rule applies
        /// </summary>
        [Output("timeWindows")]
        public Output<Outputs.NatControlRulesTimeWindows?> TimeWindows { get; private set; } = null!;

        /// <summary>
        /// list of users for which rule must be applied
        /// </summary>
        [Output("users")]
        public Output<Outputs.NatControlRulesUsers?> Users { get; private set; } = null!;


        /// <summary>
        /// Create a NatControlRules resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NatControlRules(string name, NatControlRulesArgs? args = null, CustomResourceOptions? options = null)
            : base("zia:index/natControlRules:NatControlRules", name, args ?? new NatControlRulesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NatControlRules(string name, Input<string> id, NatControlRulesState? state = null, CustomResourceOptions? options = null)
            : base("zia:index/natControlRules:NatControlRules", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/zscaler",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NatControlRules resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NatControlRules Get(string name, Input<string> id, NatControlRulesState? state = null, CustomResourceOptions? options = null)
        {
            return new NatControlRules(name, id, state, options);
        }
    }

    public sealed class NatControlRulesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to true, the default rule is applied
        /// </summary>
        [Input("defaultRule")]
        public Input<bool>? DefaultRule { get; set; }

        /// <summary>
        /// list of departments for which rule must be applied
        /// </summary>
        [Input("departments")]
        public Input<Inputs.NatControlRulesDepartmentsArgs>? Departments { get; set; }

        /// <summary>
        /// Additional information about the rule
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destAddresses")]
        private InputList<string>? _destAddresses;

        /// <summary>
        /// Destination addresses. Supports IPv4, FQDNs, or wildcard FQDNs
        /// </summary>
        public InputList<string> DestAddresses
        {
            get => _destAddresses ?? (_destAddresses = new InputList<string>());
            set => _destAddresses = value;
        }

        [Input("destCountries")]
        private InputList<string>? _destCountries;

        /// <summary>
        /// Destination countries for which the rule is applicable. If not set, the rule is not restricted to specific destination
        /// countries.
        /// </summary>
        public InputList<string> DestCountries
        {
            get => _destCountries ?? (_destCountries = new InputList<string>());
            set => _destCountries = value;
        }

        [Input("destIpCategories")]
        private InputList<string>? _destIpCategories;
        public InputList<string> DestIpCategories
        {
            get => _destIpCategories ?? (_destIpCategories = new InputList<string>());
            set => _destIpCategories = value;
        }

        /// <summary>
        /// list of destination ip groups
        /// </summary>
        [Input("destIpGroups")]
        public Input<Inputs.NatControlRulesDestIpGroupsArgs>? DestIpGroups { get; set; }

        /// <summary>
        /// list of destination ipv6 groups
        /// </summary>
        [Input("destIpv6Groups")]
        public Input<Inputs.NatControlRulesDestIpv6GroupsArgs>? DestIpv6Groups { get; set; }

        /// <summary>
        /// This field is applicable for devices that are managed using Zscaler Client Connector.
        /// </summary>
        [Input("deviceGroups")]
        public Input<Inputs.NatControlRulesDeviceGroupsArgs>? DeviceGroups { get; set; }

        /// <summary>
        /// Name-ID pairs of devices for which rule must be applied.
        /// </summary>
        [Input("devices")]
        public Input<Inputs.NatControlRulesDevicesArgs>? Devices { get; set; }

        [Input("enableFullLogging")]
        public Input<bool>? EnableFullLogging { get; set; }

        /// <summary>
        /// list of groups for which rule must be applied
        /// </summary>
        [Input("groups")]
        public Input<Inputs.NatControlRulesGroupsArgs>? Groups { get; set; }

        /// <summary>
        /// list of Labels that are applicable to the rule.
        /// </summary>
        [Input("labels")]
        public Input<Inputs.NatControlRulesLabelsArgs>? Labels { get; set; }

        /// <summary>
        /// list of locations groups
        /// </summary>
        [Input("locationGroups")]
        public Input<Inputs.NatControlRulesLocationGroupsArgs>? LocationGroups { get; set; }

        /// <summary>
        /// list of locations for which rule must be applied
        /// </summary>
        [Input("locations")]
        public Input<Inputs.NatControlRulesLocationsArgs>? Locations { get; set; }

        /// <summary>
        /// Name of the nat control policy rule
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// list of nw service groups
        /// </summary>
        [Input("nwServiceGroups")]
        public Input<Inputs.NatControlRulesNwServiceGroupsArgs>? NwServiceGroups { get; set; }

        /// <summary>
        /// list of nw services
        /// </summary>
        [Input("nwServices")]
        public Input<Inputs.NatControlRulesNwServicesArgs>? NwServices { get; set; }

        /// <summary>
        /// Rule order number. If omitted, the rule will be added to the end of the rule set.
        /// </summary>
        [Input("order")]
        public Input<int>? Order { get; set; }

        /// <summary>
        /// If set to true, a predefined rule is applied
        /// </summary>
        [Input("predefined")]
        public Input<bool>? Predefined { get; set; }

        /// <summary>
        /// Admin rank of the nat control policy rule
        /// </summary>
        [Input("rank")]
        public Input<int>? Rank { get; set; }

        /// <summary>
        /// The action the nat control policy rule takes when packets match the rule
        /// </summary>
        [Input("redirectFqdn")]
        public Input<string>? RedirectFqdn { get; set; }

        /// <summary>
        /// The action the nat control policy rule takes when packets match the rule
        /// </summary>
        [Input("redirectIp")]
        public Input<string>? RedirectIp { get; set; }

        /// <summary>
        /// The action the nat control policy rule takes when packets match the rule
        /// </summary>
        [Input("redirectPort")]
        public Input<int>? RedirectPort { get; set; }

        [Input("resCategories")]
        private InputList<string>? _resCategories;

        /// <summary>
        /// List of destination domain categories to which the rule applies
        /// </summary>
        public InputList<string> ResCategories
        {
            get => _resCategories ?? (_resCategories = new InputList<string>());
            set => _resCategories = value;
        }

        /// <summary>
        /// list of source ip groups
        /// </summary>
        [Input("srcIpGroups")]
        public Input<Inputs.NatControlRulesSrcIpGroupsArgs>? SrcIpGroups { get; set; }

        [Input("srcIps")]
        private InputList<string>? _srcIps;

        /// <summary>
        /// User-defined source IP addresses for which the rule is applicable. If not set, the rule is not restricted to a specific
        /// source IP address.
        /// </summary>
        public InputList<string> SrcIps
        {
            get => _srcIps ?? (_srcIps = new InputList<string>());
            set => _srcIps = value;
        }

        /// <summary>
        /// list of source ipv6 groups
        /// </summary>
        [Input("srcIpv6Groups")]
        public Input<Inputs.NatControlRulesSrcIpv6GroupsArgs>? SrcIpv6Groups { get; set; }

        /// <summary>
        /// Determines whether the nat control policy rule is enabled or disabled
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The time interval in which the nat control policy rule applies
        /// </summary>
        [Input("timeWindows")]
        public Input<Inputs.NatControlRulesTimeWindowsArgs>? TimeWindows { get; set; }

        /// <summary>
        /// list of users for which rule must be applied
        /// </summary>
        [Input("users")]
        public Input<Inputs.NatControlRulesUsersArgs>? Users { get; set; }

        public NatControlRulesArgs()
        {
        }
        public static new NatControlRulesArgs Empty => new NatControlRulesArgs();
    }

    public sealed class NatControlRulesState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to true, the default rule is applied
        /// </summary>
        [Input("defaultRule")]
        public Input<bool>? DefaultRule { get; set; }

        /// <summary>
        /// list of departments for which rule must be applied
        /// </summary>
        [Input("departments")]
        public Input<Inputs.NatControlRulesDepartmentsGetArgs>? Departments { get; set; }

        /// <summary>
        /// Additional information about the rule
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destAddresses")]
        private InputList<string>? _destAddresses;

        /// <summary>
        /// Destination addresses. Supports IPv4, FQDNs, or wildcard FQDNs
        /// </summary>
        public InputList<string> DestAddresses
        {
            get => _destAddresses ?? (_destAddresses = new InputList<string>());
            set => _destAddresses = value;
        }

        [Input("destCountries")]
        private InputList<string>? _destCountries;

        /// <summary>
        /// Destination countries for which the rule is applicable. If not set, the rule is not restricted to specific destination
        /// countries.
        /// </summary>
        public InputList<string> DestCountries
        {
            get => _destCountries ?? (_destCountries = new InputList<string>());
            set => _destCountries = value;
        }

        [Input("destIpCategories")]
        private InputList<string>? _destIpCategories;
        public InputList<string> DestIpCategories
        {
            get => _destIpCategories ?? (_destIpCategories = new InputList<string>());
            set => _destIpCategories = value;
        }

        /// <summary>
        /// list of destination ip groups
        /// </summary>
        [Input("destIpGroups")]
        public Input<Inputs.NatControlRulesDestIpGroupsGetArgs>? DestIpGroups { get; set; }

        /// <summary>
        /// list of destination ipv6 groups
        /// </summary>
        [Input("destIpv6Groups")]
        public Input<Inputs.NatControlRulesDestIpv6GroupsGetArgs>? DestIpv6Groups { get; set; }

        /// <summary>
        /// This field is applicable for devices that are managed using Zscaler Client Connector.
        /// </summary>
        [Input("deviceGroups")]
        public Input<Inputs.NatControlRulesDeviceGroupsGetArgs>? DeviceGroups { get; set; }

        /// <summary>
        /// Name-ID pairs of devices for which rule must be applied.
        /// </summary>
        [Input("devices")]
        public Input<Inputs.NatControlRulesDevicesGetArgs>? Devices { get; set; }

        [Input("enableFullLogging")]
        public Input<bool>? EnableFullLogging { get; set; }

        /// <summary>
        /// list of groups for which rule must be applied
        /// </summary>
        [Input("groups")]
        public Input<Inputs.NatControlRulesGroupsGetArgs>? Groups { get; set; }

        /// <summary>
        /// list of Labels that are applicable to the rule.
        /// </summary>
        [Input("labels")]
        public Input<Inputs.NatControlRulesLabelsGetArgs>? Labels { get; set; }

        /// <summary>
        /// list of locations groups
        /// </summary>
        [Input("locationGroups")]
        public Input<Inputs.NatControlRulesLocationGroupsGetArgs>? LocationGroups { get; set; }

        /// <summary>
        /// list of locations for which rule must be applied
        /// </summary>
        [Input("locations")]
        public Input<Inputs.NatControlRulesLocationsGetArgs>? Locations { get; set; }

        /// <summary>
        /// Name of the nat control policy rule
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// list of nw service groups
        /// </summary>
        [Input("nwServiceGroups")]
        public Input<Inputs.NatControlRulesNwServiceGroupsGetArgs>? NwServiceGroups { get; set; }

        /// <summary>
        /// list of nw services
        /// </summary>
        [Input("nwServices")]
        public Input<Inputs.NatControlRulesNwServicesGetArgs>? NwServices { get; set; }

        /// <summary>
        /// Rule order number. If omitted, the rule will be added to the end of the rule set.
        /// </summary>
        [Input("order")]
        public Input<int>? Order { get; set; }

        /// <summary>
        /// If set to true, a predefined rule is applied
        /// </summary>
        [Input("predefined")]
        public Input<bool>? Predefined { get; set; }

        /// <summary>
        /// Admin rank of the nat control policy rule
        /// </summary>
        [Input("rank")]
        public Input<int>? Rank { get; set; }

        /// <summary>
        /// The action the nat control policy rule takes when packets match the rule
        /// </summary>
        [Input("redirectFqdn")]
        public Input<string>? RedirectFqdn { get; set; }

        /// <summary>
        /// The action the nat control policy rule takes when packets match the rule
        /// </summary>
        [Input("redirectIp")]
        public Input<string>? RedirectIp { get; set; }

        /// <summary>
        /// The action the nat control policy rule takes when packets match the rule
        /// </summary>
        [Input("redirectPort")]
        public Input<int>? RedirectPort { get; set; }

        [Input("resCategories")]
        private InputList<string>? _resCategories;

        /// <summary>
        /// List of destination domain categories to which the rule applies
        /// </summary>
        public InputList<string> ResCategories
        {
            get => _resCategories ?? (_resCategories = new InputList<string>());
            set => _resCategories = value;
        }

        [Input("ruleId")]
        public Input<int>? RuleId { get; set; }

        /// <summary>
        /// list of source ip groups
        /// </summary>
        [Input("srcIpGroups")]
        public Input<Inputs.NatControlRulesSrcIpGroupsGetArgs>? SrcIpGroups { get; set; }

        [Input("srcIps")]
        private InputList<string>? _srcIps;

        /// <summary>
        /// User-defined source IP addresses for which the rule is applicable. If not set, the rule is not restricted to a specific
        /// source IP address.
        /// </summary>
        public InputList<string> SrcIps
        {
            get => _srcIps ?? (_srcIps = new InputList<string>());
            set => _srcIps = value;
        }

        /// <summary>
        /// list of source ipv6 groups
        /// </summary>
        [Input("srcIpv6Groups")]
        public Input<Inputs.NatControlRulesSrcIpv6GroupsGetArgs>? SrcIpv6Groups { get; set; }

        /// <summary>
        /// Determines whether the nat control policy rule is enabled or disabled
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The time interval in which the nat control policy rule applies
        /// </summary>
        [Input("timeWindows")]
        public Input<Inputs.NatControlRulesTimeWindowsGetArgs>? TimeWindows { get; set; }

        /// <summary>
        /// list of users for which rule must be applied
        /// </summary>
        [Input("users")]
        public Input<Inputs.NatControlRulesUsersGetArgs>? Users { get; set; }

        public NatControlRulesState()
        {
        }
        public static new NatControlRulesState Empty => new NatControlRulesState();
    }
}
