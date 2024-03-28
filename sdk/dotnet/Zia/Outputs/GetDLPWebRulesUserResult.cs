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
    public sealed class GetDLPWebRulesUserResult
    {
        public readonly ImmutableDictionary<string, string> Extensions;
        /// <summary>
        /// A unique identifier assigned to the workload group
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// The name of the workload group
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetDLPWebRulesUserResult(
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
