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
    public sealed class GetDLPWebRulesWorkloadGroupExpressionJsonResult
    {
        public readonly ImmutableArray<Outputs.GetDLPWebRulesWorkloadGroupExpressionJsonExpressionContainerResult> ExpressionContainers;

        [OutputConstructor]
        private GetDLPWebRulesWorkloadGroupExpressionJsonResult(ImmutableArray<Outputs.GetDLPWebRulesWorkloadGroupExpressionJsonExpressionContainerResult> expressionContainers)
        {
            ExpressionContainers = expressionContainers;
        }
    }
}
