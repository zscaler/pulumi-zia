// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zia.Outputs
{

    [OutputType]
    public sealed class GetFirewallDNSRulesSrcIpGroupResult
    {
        public readonly ImmutableDictionary<string, string> Extensions;
        /// <summary>
        /// Unique identifier for the Firewall Filtering policy rule
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Name of the Firewall Filtering policy rule
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetFirewallDNSRulesSrcIpGroupResult(
            ImmutableDictionary<string, string> extensions,

            int id,

            string name)
        {
            Extensions = extensions;
            Id = id;
            Name = name;
        }
    }
}
