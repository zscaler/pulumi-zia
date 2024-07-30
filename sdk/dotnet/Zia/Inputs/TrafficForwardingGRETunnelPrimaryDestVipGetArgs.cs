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

    public sealed class TrafficForwardingGRETunnelPrimaryDestVipGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Data center information
        /// </summary>
        [Input("datacenter")]
        public Input<string>? Datacenter { get; set; }

        /// <summary>
        /// GRE cluster virtual IP ID
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// GRE cluster virtual IP address (VIP)
        /// </summary>
        [Input("virtualIp")]
        public Input<string>? VirtualIp { get; set; }

        public TrafficForwardingGRETunnelPrimaryDestVipGetArgs()
        {
        }
        public static new TrafficForwardingGRETunnelPrimaryDestVipGetArgs Empty => new TrafficForwardingGRETunnelPrimaryDestVipGetArgs();
    }
}
