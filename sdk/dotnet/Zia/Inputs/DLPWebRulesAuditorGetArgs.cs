// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zia.Inputs
{

    public sealed class DLPWebRulesAuditorGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique identifier assigned to the workload group
        /// </summary>
        [Input("id", required: true)]
        public Input<int> Id { get; set; } = null!;

        public DLPWebRulesAuditorGetArgs()
        {
        }
        public static new DLPWebRulesAuditorGetArgs Empty => new DLPWebRulesAuditorGetArgs();
    }
}
