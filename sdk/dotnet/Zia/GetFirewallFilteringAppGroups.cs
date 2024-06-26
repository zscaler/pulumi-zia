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
    public static class GetFirewallFilteringAppGroups
    {
        public static Task<GetFirewallFilteringAppGroupsResult> InvokeAsync(GetFirewallFilteringAppGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallFilteringAppGroupsResult>("zia:index/getFirewallFilteringAppGroups:getFirewallFilteringAppGroups", args ?? new GetFirewallFilteringAppGroupsArgs(), options.WithDefaults());

        public static Output<GetFirewallFilteringAppGroupsResult> Invoke(GetFirewallFilteringAppGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallFilteringAppGroupsResult>("zia:index/getFirewallFilteringAppGroups:getFirewallFilteringAppGroups", args ?? new GetFirewallFilteringAppGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallFilteringAppGroupsArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public string? Name { get; set; }

        public GetFirewallFilteringAppGroupsArgs()
        {
        }
        public static new GetFirewallFilteringAppGroupsArgs Empty => new GetFirewallFilteringAppGroupsArgs();
    }

    public sealed class GetFirewallFilteringAppGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetFirewallFilteringAppGroupsInvokeArgs()
        {
        }
        public static new GetFirewallFilteringAppGroupsInvokeArgs Empty => new GetFirewallFilteringAppGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallFilteringAppGroupsResult
    {
        public readonly int Id;
        public readonly string Name;
        public readonly bool NameL10nTag;

        [OutputConstructor]
        private GetFirewallFilteringAppGroupsResult(
            int id,

            string name,

            bool nameL10nTag)
        {
            Id = id;
            Name = name;
            NameL10nTag = nameL10nTag;
        }
    }
}
