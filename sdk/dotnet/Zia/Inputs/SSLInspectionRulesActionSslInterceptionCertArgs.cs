// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zia.Inputs
{

    public sealed class SSLInspectionRulesActionSslInterceptionCertArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Integer) - A unique identifier assigned to the workload group
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public SSLInspectionRulesActionSslInterceptionCertArgs()
        {
        }
        public static new SSLInspectionRulesActionSslInterceptionCertArgs Empty => new SSLInspectionRulesActionSslInterceptionCertArgs();
    }
}
