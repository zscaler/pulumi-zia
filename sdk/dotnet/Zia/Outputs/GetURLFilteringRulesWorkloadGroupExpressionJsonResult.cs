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
    public sealed class GetURLFilteringRulesWorkloadGroupExpressionJsonResult
    {
        public readonly ImmutableArray<Outputs.GetURLFilteringRulesWorkloadGroupExpressionJsonExpressionContainerResult> ExpressionContainers;

        [OutputConstructor]
        private GetURLFilteringRulesWorkloadGroupExpressionJsonResult(ImmutableArray<Outputs.GetURLFilteringRulesWorkloadGroupExpressionJsonExpressionContainerResult> expressionContainers)
        {
            ExpressionContainers = expressionContainers;
        }
    }
}
