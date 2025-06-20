// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zia
{
    public static class GetForwardingControlProxies
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/about-third-party-proxies)
        /// * [API documentation](https://help.zscaler.com/zia/forwarding-control-policy#/proxies-get)
        /// 
        /// Use the **zia_forwarding_control_proxies** data source to get information about a third-party proxy service available in the Zscaler Internet Access.
        /// 
        /// ## Example Usage
        /// 
        /// ### Retrieve By Name
        /// 
        /// ```hcl
        /// data "zia_forwarding_control_proxies" "this" {
        ///     name = "Proxy01"
        /// }
        /// ```
        /// 
        /// ### Retrieve By ID
        /// 
        /// ```hcl
        /// data "zia_forwarding_control_proxies" "this" {
        ///     id = "18492370"
        /// }
        /// ```
        /// </summary>
        public static Task<GetForwardingControlProxiesResult> InvokeAsync(GetForwardingControlProxiesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetForwardingControlProxiesResult>("zia:index/getForwardingControlProxies:getForwardingControlProxies", args ?? new GetForwardingControlProxiesArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/about-third-party-proxies)
        /// * [API documentation](https://help.zscaler.com/zia/forwarding-control-policy#/proxies-get)
        /// 
        /// Use the **zia_forwarding_control_proxies** data source to get information about a third-party proxy service available in the Zscaler Internet Access.
        /// 
        /// ## Example Usage
        /// 
        /// ### Retrieve By Name
        /// 
        /// ```hcl
        /// data "zia_forwarding_control_proxies" "this" {
        ///     name = "Proxy01"
        /// }
        /// ```
        /// 
        /// ### Retrieve By ID
        /// 
        /// ```hcl
        /// data "zia_forwarding_control_proxies" "this" {
        ///     id = "18492370"
        /// }
        /// ```
        /// </summary>
        public static Output<GetForwardingControlProxiesResult> Invoke(GetForwardingControlProxiesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetForwardingControlProxiesResult>("zia:index/getForwardingControlProxies:getForwardingControlProxies", args ?? new GetForwardingControlProxiesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/about-third-party-proxies)
        /// * [API documentation](https://help.zscaler.com/zia/forwarding-control-policy#/proxies-get)
        /// 
        /// Use the **zia_forwarding_control_proxies** data source to get information about a third-party proxy service available in the Zscaler Internet Access.
        /// 
        /// ## Example Usage
        /// 
        /// ### Retrieve By Name
        /// 
        /// ```hcl
        /// data "zia_forwarding_control_proxies" "this" {
        ///     name = "Proxy01"
        /// }
        /// ```
        /// 
        /// ### Retrieve By ID
        /// 
        /// ```hcl
        /// data "zia_forwarding_control_proxies" "this" {
        ///     id = "18492370"
        /// }
        /// ```
        /// </summary>
        public static Output<GetForwardingControlProxiesResult> Invoke(GetForwardingControlProxiesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetForwardingControlProxiesResult>("zia:index/getForwardingControlProxies:getForwardingControlProxies", args ?? new GetForwardingControlProxiesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetForwardingControlProxiesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for the third-party proxy services
        /// </summary>
        [Input("id")]
        public int? Id { get; set; }

        /// <summary>
        /// Proxy name for the third-party proxy services
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetForwardingControlProxiesArgs()
        {
        }
        public static new GetForwardingControlProxiesArgs Empty => new GetForwardingControlProxiesArgs();
    }

    public sealed class GetForwardingControlProxiesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for the third-party proxy services
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Proxy name for the third-party proxy services
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetForwardingControlProxiesInvokeArgs()
        {
        }
        public static new GetForwardingControlProxiesInvokeArgs Empty => new GetForwardingControlProxiesInvokeArgs();
    }


    [OutputType]
    public sealed class GetForwardingControlProxiesResult
    {
        /// <summary>
        /// (String) The IP address or the FQDN of the third-party proxy service
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// (Boolean) Flag indicating whether the added X-Authenticated-User header is Base64 encoded. When enabled, the user ID is encoded using the Base64 encoding method.
        /// </summary>
        public readonly bool Base64EncodeXauHeader;
        /// <summary>
        /// (Set of Objects) The root certificate used by the third-party proxy to perform SSL inspection. This root certificate is used by Zscaler to validate the SSL leaf certificates signed by the upstream proxy. The required root certificate appears in this drop-down list only if it is uploaded from the Administration &gt; Root Certificates page.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingControlProxiesCertResult> Certs;
        /// <summary>
        /// (String) Additional notes or information
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (Integer) Identifier that uniquely identifies the certificate
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// (Boolean) Flag indicating whether X-Authenticated-User header is added by the proxy. Enable to automatically insert authenticated user ID to the HTTP header, X-Authenticated-User.
        /// </summary>
        public readonly bool InsertXauHeader;
        public readonly ImmutableArray<Outputs.GetForwardingControlProxiesLastModifiedByResult> LastModifiedBies;
        public readonly int LastModifiedTime;
        public readonly string Name;
        /// <summary>
        /// (integer) The port number on which the third-party proxy service listens to the requests forwarded from Zscaler
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// (String) Gateway type. Returned values: `PROXYCHAIN`, `ZIA`, `ECSELF`
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetForwardingControlProxiesResult(
            string address,

            bool base64EncodeXauHeader,

            ImmutableArray<Outputs.GetForwardingControlProxiesCertResult> certs,

            string description,

            int id,

            bool insertXauHeader,

            ImmutableArray<Outputs.GetForwardingControlProxiesLastModifiedByResult> lastModifiedBies,

            int lastModifiedTime,

            string name,

            int port,

            string type)
        {
            Address = address;
            Base64EncodeXauHeader = base64EncodeXauHeader;
            Certs = certs;
            Description = description;
            Id = id;
            InsertXauHeader = insertXauHeader;
            LastModifiedBies = lastModifiedBies;
            LastModifiedTime = lastModifiedTime;
            Name = name;
            Port = port;
            Type = type;
        }
    }
}
