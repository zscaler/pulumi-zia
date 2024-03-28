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
    public sealed class GetDLPDictionariesIdmProfileMatchAccuracyResult
    {
        /// <summary>
        /// The action applied to a DLP dictionary using patterns
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDLPDictionariesIdmProfileMatchAccuracyAdpIdmProfileResult> AdpIdmProfiles;
        /// <summary>
        /// The IDM template match accuracy.
        /// </summary>
        public readonly string MatchAccuracy;

        [OutputConstructor]
        private GetDLPDictionariesIdmProfileMatchAccuracyResult(
            ImmutableArray<Outputs.GetDLPDictionariesIdmProfileMatchAccuracyAdpIdmProfileResult> adpIdmProfiles,

            string matchAccuracy)
        {
            AdpIdmProfiles = adpIdmProfiles;
            MatchAccuracy = matchAccuracy;
        }
    }
}
