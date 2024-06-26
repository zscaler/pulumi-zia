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
    public sealed class GetTrafficForwardingGRETunnelSecondaryDestVipResult
    {
        public readonly string City;
        public readonly string CountryCode;
        public readonly string Datacenter;
        public readonly int Id;
        public readonly int Latitude;
        public readonly int Longitude;
        public readonly bool PrivateServiceEdge;
        public readonly string Region;
        public readonly string VirtualIp;

        [OutputConstructor]
        private GetTrafficForwardingGRETunnelSecondaryDestVipResult(
            string city,

            string countryCode,

            string datacenter,

            int id,

            int latitude,

            int longitude,

            bool privateServiceEdge,

            string region,

            string virtualIp)
        {
            City = city;
            CountryCode = countryCode;
            Datacenter = datacenter;
            Id = id;
            Latitude = latitude;
            Longitude = longitude;
            PrivateServiceEdge = privateServiceEdge;
            Region = region;
            VirtualIp = virtualIp;
        }
    }
}
