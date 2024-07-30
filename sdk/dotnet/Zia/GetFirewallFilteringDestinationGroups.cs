// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zia
{
    public static class GetFirewallFilteringDestinationGroups
    {
        /// <summary>
        /// Use the **zia_firewall_filtering_destination_groups** data source to get information about IP destination groups option available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering rule.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zia.GetFirewallFilteringDestinationGroups.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetFirewallFilteringDestinationGroupsResult> InvokeAsync(GetFirewallFilteringDestinationGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallFilteringDestinationGroupsResult>("zia:index/getFirewallFilteringDestinationGroups:getFirewallFilteringDestinationGroups", args ?? new GetFirewallFilteringDestinationGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zia_firewall_filtering_destination_groups** data source to get information about IP destination groups option available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering rule.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zia.GetFirewallFilteringDestinationGroups.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetFirewallFilteringDestinationGroupsResult> Invoke(GetFirewallFilteringDestinationGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallFilteringDestinationGroupsResult>("zia:index/getFirewallFilteringDestinationGroups:getFirewallFilteringDestinationGroups", args ?? new GetFirewallFilteringDestinationGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallFilteringDestinationGroupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the destination group resource.
        /// </summary>
        [Input("id")]
        public int? Id { get; set; }

        /// <summary>
        /// The name of the destination group to be exported.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetFirewallFilteringDestinationGroupsArgs()
        {
        }
        public static new GetFirewallFilteringDestinationGroupsArgs Empty => new GetFirewallFilteringDestinationGroupsArgs();
    }

    public sealed class GetFirewallFilteringDestinationGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the destination group resource.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// The name of the destination group to be exported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetFirewallFilteringDestinationGroupsInvokeArgs()
        {
        }
        public static new GetFirewallFilteringDestinationGroupsInvokeArgs Empty => new GetFirewallFilteringDestinationGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallFilteringDestinationGroupsResult
    {
        /// <summary>
        /// (List of String) Destination IP addresses within the group
        /// </summary>
        public readonly ImmutableArray<string> Addresses;
        /// <summary>
        /// (List of String) Destination IP address counties. You can identify destinations based on the location of a server.
        /// </summary>
        public readonly ImmutableArray<string> Countries;
        /// <summary>
        /// (String) Additional information about the destination IP group
        /// </summary>
        public readonly string Description;
        public readonly int Id;
        /// <summary>
        /// (List of String) Destination IP address URL categories. You can identify destinations based on the URL category of the domain. See list of all IP Categories [Here](https://help.zscaler.com/zia/firewall-policies#/ipDestinationGroups-get)
        /// * !&gt; **WARNING:** The `ip_categories` attribute only accepts custom URL categories.
        /// </summary>
        public readonly ImmutableArray<string> IpCategories;
        public readonly string Name;
        /// <summary>
        /// (String) Destination IP group type (i.e., the group can contain destination IP addresses or FQDNs)
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetFirewallFilteringDestinationGroupsResult(
            ImmutableArray<string> addresses,

            ImmutableArray<string> countries,

            string description,

            int id,

            ImmutableArray<string> ipCategories,

            string name,

            string type)
        {
            Addresses = addresses;
            Countries = countries;
            Description = description;
            Id = id;
            IpCategories = ipCategories;
            Name = name;
            Type = type;
        }
    }
}
