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
    public sealed class GetSandboxReportNetworkingResult
    {
        public readonly string Risk;
        public readonly string Signature;
        public readonly ImmutableArray<string> SignatureSources;

        [OutputConstructor]
        private GetSandboxReportNetworkingResult(
            string risk,

            string signature,

            ImmutableArray<string> signatureSources)
        {
            Risk = risk;
            Signature = signature;
            SignatureSources = signatureSources;
        }
    }
}
