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
    public sealed class GetTrafficForwardingGREInternalIPRangeListResult
    {
        /// <summary>
        /// (String) Starting IP address in the range
        /// </summary>
        public readonly string EndIpAddress;
        /// <summary>
        /// (String) Ending IP address in the range
        /// </summary>
        public readonly string StartIpAddress;

        [OutputConstructor]
        private GetTrafficForwardingGREInternalIPRangeListResult(
            string endIpAddress,

            string startIpAddress)
        {
            EndIpAddress = endIpAddress;
            StartIpAddress = startIpAddress;
        }
    }
}
