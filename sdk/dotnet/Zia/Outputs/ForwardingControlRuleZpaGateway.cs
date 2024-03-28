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
    public sealed class ForwardingControlRuleZpaGateway
    {
        /// <summary>
        /// (int) Identifier that uniquely identifies an entity
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Name of the Firewall Filtering policy rule
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private ForwardingControlRuleZpaGateway(
            int id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
