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

    public sealed class FirewallFilteringNetworkServicesSrcTcpPortArgs : global::Pulumi.ResourceArgs
    {
        [Input("end")]
        public Input<int>? End { get; set; }

        [Input("start")]
        public Input<int>? Start { get; set; }

        public FirewallFilteringNetworkServicesSrcTcpPortArgs()
        {
        }
        public static new FirewallFilteringNetworkServicesSrcTcpPortArgs Empty => new FirewallFilteringNetworkServicesSrcTcpPortArgs();
    }
}
