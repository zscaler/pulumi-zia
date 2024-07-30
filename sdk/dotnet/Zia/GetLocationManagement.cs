// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zia
{
    public static class GetLocationManagement
    {
        /// <summary>
        /// Use the **zia_location_management** data source to get information about a location resource available in the Zscaler Internet Access Location Management. This resource can then be referenced in multiple other resources, such as URL Filtering Rules, Firewall rules etc.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zia.GetLocationManagement.Invoke(new()
        ///     {
        ///         Name = "San Jose",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetLocationManagementResult> InvokeAsync(GetLocationManagementArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLocationManagementResult>("zia:index/getLocationManagement:getLocationManagement", args ?? new GetLocationManagementArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zia_location_management** data source to get information about a location resource available in the Zscaler Internet Access Location Management. This resource can then be referenced in multiple other resources, such as URL Filtering Rules, Firewall rules etc.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zia.GetLocationManagement.Invoke(new()
        ///     {
        ///         Name = "San Jose",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetLocationManagementResult> Invoke(GetLocationManagementInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLocationManagementResult>("zia:index/getLocationManagement:getLocationManagement", args ?? new GetLocationManagementInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLocationManagementArgs : global::Pulumi.InvokeArgs
    {
        [Input("basicAuthEnabled")]
        public bool? BasicAuthEnabled { get; set; }

        /// <summary>
        /// The ID of the location to be exported.
        /// </summary>
        [Input("id")]
        public int? Id { get; set; }

        /// <summary>
        /// The name of the location to be exported.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("parentName")]
        public string? ParentName { get; set; }

        public GetLocationManagementArgs()
        {
        }
        public static new GetLocationManagementArgs Empty => new GetLocationManagementArgs();
    }

    public sealed class GetLocationManagementInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("basicAuthEnabled")]
        public Input<bool>? BasicAuthEnabled { get; set; }

        /// <summary>
        /// The ID of the location to be exported.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// The name of the location to be exported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parentName")]
        public Input<string>? ParentName { get; set; }

        public GetLocationManagementInvokeArgs()
        {
        }
        public static new GetLocationManagementInvokeArgs Empty => new GetLocationManagementInvokeArgs();
    }


    [OutputType]
    public sealed class GetLocationManagementResult
    {
        /// <summary>
        /// (Boolean) For First Time AUP Behavior, Block Internet Access. When set, all internet access (including non-HTTP traffic) is disabled until the user accepts the AUP.
        /// </summary>
        public readonly bool AupBlockInternetUntilAccepted;
        /// <summary>
        /// (Boolean) Enable AUP. When set to true, AUP is enabled for the location.
        /// </summary>
        public readonly bool AupEnabled;
        /// <summary>
        /// (Boolean) For First Time AUP Behavior, Force SSL Inspection. When set, Zscaler will force SSL Inspection in order to enforce AUP for HTTPS traffic.
        /// </summary>
        public readonly bool AupForceSslInspection;
        /// <summary>
        /// (Number) Custom AUP Frequency. Refresh time (in days) to re-validate the AUP.
        /// </summary>
        public readonly int AupTimeoutInDays;
        /// <summary>
        /// (Boolean) Enforce Authentication. Required when ports are enabled, IP Surrogate is enabled, or Kerberos Authentication is enabled.
        /// </summary>
        public readonly bool AuthRequired;
        public readonly bool BasicAuthEnabled;
        /// <summary>
        /// (Boolean) Enable Caution. When set to true, a caution notifcation is enabled for the location.
        /// </summary>
        public readonly bool CautionEnabled;
        /// <summary>
        /// (String) Country
        /// </summary>
        public readonly string Country;
        /// <summary>
        /// (String) Additional notes or information regarding the location or sub-location. The description cannot exceed 1024 characters.
        /// </summary>
        public readonly string Description;
        public readonly bool DigestAuthEnabled;
        /// <summary>
        /// (String) Display Time Unit. The time unit to display for IP Surrogate idle time to disassociation.
        /// </summary>
        public readonly string DisplayTimeUnit;
        /// <summary>
        /// (Number) Download bandwidth in bytes. The value `0` implies no Bandwidth Control enforcement.
        /// </summary>
        public readonly int DnBandwidth;
        /// <summary>
        /// (Number) Identifier that uniquely identifies an entity
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// (Number) Idle Time to Disassociation. The user mapping idle time (in minutes) is required if a Surrogate IP is enabled.
        /// </summary>
        public readonly int IdleTimeInMinutes;
        public readonly bool IotDiscoveryEnabled;
        /// <summary>
        /// (List of String) For locations: IP addresses of the egress points that are provisioned in the Zscaler Cloud. Each entry is a single IP address (e.g., `238.10.33.9`). For sub-locations: Egress, internal, or GRE tunnel IP addresses. Each entry is either a single IP address, CIDR (e.g., `10.10.33.0/24`), or range (e.g., `10.10.33.1-10.10.33.10`)).
        /// </summary>
        public readonly ImmutableArray<string> IpAddresses;
        /// <summary>
        /// (Boolean) Enable IPS Control. When set to true, IPS Control is enabled for the location if Firewall is enabled.
        /// </summary>
        public readonly bool IpsControl;
        public readonly bool KerberosAuthEnabled;
        /// <summary>
        /// (String) The configured name of the entity
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// (Boolean) Enable Firewall. When set to true, Firewall is enabled for the location.
        /// </summary>
        public readonly bool OfwEnabled;
        /// <summary>
        /// (Number) - Parent Location ID. If this ID does not exist or is `0`, it is implied that it is a parent location. Otherwise, it is a sub-location whose parent has this ID. x-applicableTo: `SUB`
        /// </summary>
        public readonly int ParentId;
        public readonly string? ParentName;
        /// <summary>
        /// (String) IP ports that are associated with the location.
        /// </summary>
        public readonly string Ports;
        /// <summary>
        /// (String) Profile tag that specifies the location traffic type. If not specified, this tag defaults to `Unassigned`.
        /// </summary>
        public readonly string Profile;
        /// <summary>
        /// (Boolean) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.
        /// </summary>
        public readonly bool SslScanEnabled;
        /// <summary>
        /// (Boolean) Enable Surrogate IP. When set to true, users are mapped to internal device IP addresses.
        /// </summary>
        public readonly bool SurrogateIp;
        /// <summary>
        /// (Boolean) Enforce Surrogate IP for Known Browsers. When set to true, IP Surrogate is enforced for all known browsers.
        /// </summary>
        public readonly bool SurrogateIpEnforcedForKnownBrowsers;
        /// <summary>
        /// (Number) Refresh Time for re-validation of Surrogacy. The surrogate refresh time (in minutes) to re-validate the IP surrogates.
        /// </summary>
        public readonly int SurrogateRefreshTimeInMinutes;
        /// <summary>
        /// (String) Display Refresh Time Unit. The time unit to display for refresh time for re-validation of surrogacy.
        /// </summary>
        public readonly string SurrogateRefreshTimeUnit;
        /// <summary>
        /// (String) Timezone of the location. If not specified, it defaults to GMT.
        /// </summary>
        public readonly string Tz;
        /// <summary>
        /// (Number) Upload bandwidth in bytes. The value `0` implies no Bandwidth Control enforcement.
        /// </summary>
        public readonly int UpBandwidth;
        public readonly ImmutableArray<Outputs.GetLocationManagementVpnCredentialResult> VpnCredentials;
        /// <summary>
        /// (Boolean) Enable XFF Forwarding. When set to true, traffic is passed to Zscaler Cloud via the X-Forwarded-For (XFF) header.
        /// </summary>
        public readonly bool XffForwardEnabled;
        /// <summary>
        /// (Boolean) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.
        /// </summary>
        public readonly bool ZappSslScanEnabled;

        [OutputConstructor]
        private GetLocationManagementResult(
            bool aupBlockInternetUntilAccepted,

            bool aupEnabled,

            bool aupForceSslInspection,

            int aupTimeoutInDays,

            bool authRequired,

            bool basicAuthEnabled,

            bool cautionEnabled,

            string country,

            string description,

            bool digestAuthEnabled,

            string displayTimeUnit,

            int dnBandwidth,

            int? id,

            int idleTimeInMinutes,

            bool iotDiscoveryEnabled,

            ImmutableArray<string> ipAddresses,

            bool ipsControl,

            bool kerberosAuthEnabled,

            string? name,

            bool ofwEnabled,

            int parentId,

            string? parentName,

            string ports,

            string profile,

            bool sslScanEnabled,

            bool surrogateIp,

            bool surrogateIpEnforcedForKnownBrowsers,

            int surrogateRefreshTimeInMinutes,

            string surrogateRefreshTimeUnit,

            string tz,

            int upBandwidth,

            ImmutableArray<Outputs.GetLocationManagementVpnCredentialResult> vpnCredentials,

            bool xffForwardEnabled,

            bool zappSslScanEnabled)
        {
            AupBlockInternetUntilAccepted = aupBlockInternetUntilAccepted;
            AupEnabled = aupEnabled;
            AupForceSslInspection = aupForceSslInspection;
            AupTimeoutInDays = aupTimeoutInDays;
            AuthRequired = authRequired;
            BasicAuthEnabled = basicAuthEnabled;
            CautionEnabled = cautionEnabled;
            Country = country;
            Description = description;
            DigestAuthEnabled = digestAuthEnabled;
            DisplayTimeUnit = displayTimeUnit;
            DnBandwidth = dnBandwidth;
            Id = id;
            IdleTimeInMinutes = idleTimeInMinutes;
            IotDiscoveryEnabled = iotDiscoveryEnabled;
            IpAddresses = ipAddresses;
            IpsControl = ipsControl;
            KerberosAuthEnabled = kerberosAuthEnabled;
            Name = name;
            OfwEnabled = ofwEnabled;
            ParentId = parentId;
            ParentName = parentName;
            Ports = ports;
            Profile = profile;
            SslScanEnabled = sslScanEnabled;
            SurrogateIp = surrogateIp;
            SurrogateIpEnforcedForKnownBrowsers = surrogateIpEnforcedForKnownBrowsers;
            SurrogateRefreshTimeInMinutes = surrogateRefreshTimeInMinutes;
            SurrogateRefreshTimeUnit = surrogateRefreshTimeUnit;
            Tz = tz;
            UpBandwidth = upBandwidth;
            VpnCredentials = vpnCredentials;
            XffForwardEnabled = xffForwardEnabled;
            ZappSslScanEnabled = zappSslScanEnabled;
        }
    }
}
