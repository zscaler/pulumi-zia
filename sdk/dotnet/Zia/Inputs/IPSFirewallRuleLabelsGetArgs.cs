// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zia.Inputs
{

    public sealed class IPSFirewallRuleLabelsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Integer) Identifier that uniquely identifies an entity
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public IPSFirewallRuleLabelsGetArgs()
        {
        }
        public static new IPSFirewallRuleLabelsGetArgs Empty => new IPSFirewallRuleLabelsGetArgs();
    }
}
