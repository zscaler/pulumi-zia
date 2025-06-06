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
    public sealed class GetSandboxReportSummaryResult
    {
        public readonly string Category;
        public readonly int Duration;
        public readonly string FileType;
        public readonly int StartTime;
        public readonly string Status;

        [OutputConstructor]
        private GetSandboxReportSummaryResult(
            string category,

            int duration,

            string fileType,

            int startTime,

            string status)
        {
            Category = category;
            Duration = duration;
            FileType = fileType;
            StartTime = startTime;
            Status = status;
        }
    }
}
