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
    public sealed class GetTrafficForwardingGRETunnelPrimaryDestVipResult
    {
        public readonly string City;
        /// <summary>
        /// (String) When within_country is enabled, you must set this to the country code.
        /// </summary>
        public readonly string CountryCode;
        public readonly string Datacenter;
        /// <summary>
        /// Unique identifier of the static IP address that is associated to a GRE tunnel
        /// </summary>
        public readonly int Id;
        public readonly int Latitude;
        public readonly int Longitude;
        public readonly bool PrivateServiceEdge;
        public readonly string Region;
        /// <summary>
        /// (String) GRE cluster virtual IP address (VIP)
        /// </summary>
        public readonly string VirtualIp;

        [OutputConstructor]
        private GetTrafficForwardingGRETunnelPrimaryDestVipResult(
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
