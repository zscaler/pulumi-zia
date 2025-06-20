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
    public sealed class GetRiskProfilesCustomTagResult
    {
        /// <summary>
        /// Unique identifier for the risk profile.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Cloud application risk profile name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetRiskProfilesCustomTagResult(
            int id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
