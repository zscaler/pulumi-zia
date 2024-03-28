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
    public sealed class GetURLCategoriesUrlKeywordCountResult
    {
        /// <summary>
        /// (Number) Count of total keywords with retain parent category.
        /// </summary>
        public readonly int RetainParentKeywordCount;
        /// <summary>
        /// (Number) Count of URLs with retain parent category.
        /// </summary>
        public readonly int RetainParentUrlCount;
        /// <summary>
        /// (Number) Total keyword count for the category.
        /// </summary>
        public readonly int TotalKeywordCount;
        /// <summary>
        /// (Number) Custom URL count for the category.
        /// </summary>
        public readonly int TotalUrlCount;

        [OutputConstructor]
        private GetURLCategoriesUrlKeywordCountResult(
            int retainParentKeywordCount,

            int retainParentUrlCount,

            int totalKeywordCount,

            int totalUrlCount)
        {
            RetainParentKeywordCount = retainParentKeywordCount;
            RetainParentUrlCount = retainParentUrlCount;
            TotalKeywordCount = totalKeywordCount;
            TotalUrlCount = totalUrlCount;
        }
    }
}
