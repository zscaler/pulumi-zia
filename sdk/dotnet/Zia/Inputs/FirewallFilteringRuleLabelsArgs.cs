// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zia.Inputs
{

    public sealed class FirewallFilteringRuleLabelsArgs : global::Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<int>? Id { get; set; }

        public FirewallFilteringRuleLabelsArgs()
        {
        }
        public static new FirewallFilteringRuleLabelsArgs Empty => new FirewallFilteringRuleLabelsArgs();
    }
}
