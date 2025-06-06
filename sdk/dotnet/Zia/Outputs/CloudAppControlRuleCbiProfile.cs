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
    public sealed class CloudAppControlRuleCbiProfile
    {
        public readonly string? Id;
        public readonly string? Name;
        /// <summary>
        /// The browser isolation profile URL
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private CloudAppControlRuleCbiProfile(
            string? id,

            string? name,

            string? url)
        {
            Id = id;
            Name = name;
            Url = url;
        }
    }
}
