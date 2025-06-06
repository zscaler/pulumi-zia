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
    public sealed class SSLInspectionRulesWorkloadGroup
    {
        /// <summary>
        /// (Integer) - A unique identifier assigned to the workload group
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private SSLInspectionRulesWorkloadGroup(
            int id,

            string? name)
        {
            Id = id;
            Name = name;
        }
    }
}
