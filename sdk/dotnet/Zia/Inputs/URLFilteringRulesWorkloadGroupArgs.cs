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

    public sealed class URLFilteringRulesWorkloadGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique identifier for the resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<int> Id { get; set; } = null!;

        /// <summary>
        /// Name of the Firewall Filtering policy rule
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public URLFilteringRulesWorkloadGroupArgs()
        {
        }
        public static new URLFilteringRulesWorkloadGroupArgs Empty => new URLFilteringRulesWorkloadGroupArgs();
    }
}
