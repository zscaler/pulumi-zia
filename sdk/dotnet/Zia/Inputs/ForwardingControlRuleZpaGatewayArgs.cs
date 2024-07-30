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

    public sealed class ForwardingControlRuleZpaGatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (int) Identifier that uniquely identifies an entity
        /// </summary>
        [Input("id", required: true)]
        public Input<int> Id { get; set; } = null!;

        /// <summary>
        /// (string) The configured name of the entity
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ForwardingControlRuleZpaGatewayArgs()
        {
        }
        public static new ForwardingControlRuleZpaGatewayArgs Empty => new ForwardingControlRuleZpaGatewayArgs();
    }
}
