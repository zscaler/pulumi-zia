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
    public sealed class DLPWebRulesDlpEngines
    {
        /// <summary>
        /// A unique identifier assigned to the workload group
        /// </summary>
        public readonly ImmutableArray<int> Ids;

        [OutputConstructor]
        private DLPWebRulesDlpEngines(ImmutableArray<int> ids)
        {
            Ids = ids;
        }
    }
}
