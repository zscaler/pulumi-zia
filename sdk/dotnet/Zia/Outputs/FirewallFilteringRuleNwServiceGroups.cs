// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
    public sealed class FirewallFilteringRuleNwServiceGroups
    {
        public readonly ImmutableArray<int> Ids;

        [OutputConstructor]
        private FirewallFilteringRuleNwServiceGroups(ImmutableArray<int> ids)
        {
            Ids = ids;
        }
    }
}
