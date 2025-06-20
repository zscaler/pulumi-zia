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
    public static class GetFirewallFilteringApplicationGroups
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/firewall-policies#/networkApplicationGroups/{groupId}-get)
        /// * [API documentation](https://help.zscaler.com/zia/firewall-policies#/networkApplicationGroups/{groupId}-get)
        /// 
        /// Use the **zia_firewall_filtering_network_application_groups** data source to get information about network application groups available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering rule.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// # ZIA IP Source Groups
        /// data "zia_firewall_filtering_network_application_groups" "example" {
        ///     name = "example"
        /// }
        /// ```
        /// </summary>
        public static Task<GetFirewallFilteringApplicationGroupsResult> InvokeAsync(GetFirewallFilteringApplicationGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallFilteringApplicationGroupsResult>("zia:index/getFirewallFilteringApplicationGroups:getFirewallFilteringApplicationGroups", args ?? new GetFirewallFilteringApplicationGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/firewall-policies#/networkApplicationGroups/{groupId}-get)
        /// * [API documentation](https://help.zscaler.com/zia/firewall-policies#/networkApplicationGroups/{groupId}-get)
        /// 
        /// Use the **zia_firewall_filtering_network_application_groups** data source to get information about network application groups available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering rule.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// # ZIA IP Source Groups
        /// data "zia_firewall_filtering_network_application_groups" "example" {
        ///     name = "example"
        /// }
        /// ```
        /// </summary>
        public static Output<GetFirewallFilteringApplicationGroupsResult> Invoke(GetFirewallFilteringApplicationGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallFilteringApplicationGroupsResult>("zia:index/getFirewallFilteringApplicationGroups:getFirewallFilteringApplicationGroups", args ?? new GetFirewallFilteringApplicationGroupsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/firewall-policies#/networkApplicationGroups/{groupId}-get)
        /// * [API documentation](https://help.zscaler.com/zia/firewall-policies#/networkApplicationGroups/{groupId}-get)
        /// 
        /// Use the **zia_firewall_filtering_network_application_groups** data source to get information about network application groups available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering rule.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// # ZIA IP Source Groups
        /// data "zia_firewall_filtering_network_application_groups" "example" {
        ///     name = "example"
        /// }
        /// ```
        /// </summary>
        public static Output<GetFirewallFilteringApplicationGroupsResult> Invoke(GetFirewallFilteringApplicationGroupsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallFilteringApplicationGroupsResult>("zia:index/getFirewallFilteringApplicationGroups:getFirewallFilteringApplicationGroups", args ?? new GetFirewallFilteringApplicationGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallFilteringApplicationGroupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the ip source group resource.
        /// </summary>
        [Input("id")]
        public int? Id { get; set; }

        /// <summary>
        /// The name of the ip source group to be exported.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetFirewallFilteringApplicationGroupsArgs()
        {
        }
        public static new GetFirewallFilteringApplicationGroupsArgs Empty => new GetFirewallFilteringApplicationGroupsArgs();
    }

    public sealed class GetFirewallFilteringApplicationGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the ip source group resource.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// The name of the ip source group to be exported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetFirewallFilteringApplicationGroupsInvokeArgs()
        {
        }
        public static new GetFirewallFilteringApplicationGroupsInvokeArgs Empty => new GetFirewallFilteringApplicationGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallFilteringApplicationGroupsResult
    {
        public readonly string Description;
        public readonly int Id;
        public readonly string Name;
        public readonly ImmutableArray<string> NetworkApplications;

        [OutputConstructor]
        private GetFirewallFilteringApplicationGroupsResult(
            string description,

            int id,

            string name,

            ImmutableArray<string> networkApplications)
        {
            Description = description;
            Id = id;
            Name = name;
            NetworkApplications = networkApplications;
        }
    }
}
