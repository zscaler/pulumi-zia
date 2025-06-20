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
    public static class GetForwardingProxyGateway
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/about-gateways-proxies)
        /// * [API documentation](https://help.zscaler.com/zia/forwarding-control-policy#/proxyGateways-get)
        /// 
        /// Use the **zia_forwarding_control_proxy_gateway** data source to retrieve the proxy gateway information. This data source can then be associated with the attribute `proxy_gateway` when creating a Forwarding Control Rule via the resource: `zia.ForwardingControlRule`
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// # ZIA Forwarding Control - Proxy Gateway
        /// data "zia_forwarding_control_proxy_gateway" "this" {
        ///   name = "Proxy_GW01"
        /// }
        /// ```
        /// </summary>
        public static Task<GetForwardingProxyGatewayResult> InvokeAsync(GetForwardingProxyGatewayArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetForwardingProxyGatewayResult>("zia:index/getForwardingProxyGateway:getForwardingProxyGateway", args ?? new GetForwardingProxyGatewayArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/about-gateways-proxies)
        /// * [API documentation](https://help.zscaler.com/zia/forwarding-control-policy#/proxyGateways-get)
        /// 
        /// Use the **zia_forwarding_control_proxy_gateway** data source to retrieve the proxy gateway information. This data source can then be associated with the attribute `proxy_gateway` when creating a Forwarding Control Rule via the resource: `zia.ForwardingControlRule`
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// # ZIA Forwarding Control - Proxy Gateway
        /// data "zia_forwarding_control_proxy_gateway" "this" {
        ///   name = "Proxy_GW01"
        /// }
        /// ```
        /// </summary>
        public static Output<GetForwardingProxyGatewayResult> Invoke(GetForwardingProxyGatewayInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetForwardingProxyGatewayResult>("zia:index/getForwardingProxyGateway:getForwardingProxyGateway", args ?? new GetForwardingProxyGatewayInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/about-gateways-proxies)
        /// * [API documentation](https://help.zscaler.com/zia/forwarding-control-policy#/proxyGateways-get)
        /// 
        /// Use the **zia_forwarding_control_proxy_gateway** data source to retrieve the proxy gateway information. This data source can then be associated with the attribute `proxy_gateway` when creating a Forwarding Control Rule via the resource: `zia.ForwardingControlRule`
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// # ZIA Forwarding Control - Proxy Gateway
        /// data "zia_forwarding_control_proxy_gateway" "this" {
        ///   name = "Proxy_GW01"
        /// }
        /// ```
        /// </summary>
        public static Output<GetForwardingProxyGatewayResult> Invoke(GetForwardingProxyGatewayInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetForwardingProxyGatewayResult>("zia:index/getForwardingProxyGateway:getForwardingProxyGateway", args ?? new GetForwardingProxyGatewayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetForwardingProxyGatewayArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the forwarding control Proxy Gateway resource.
        /// </summary>
        [Input("id")]
        public int? Id { get; set; }

        /// <summary>
        /// The name of the forwarding control Proxy Gateway to be exported.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetForwardingProxyGatewayArgs()
        {
        }
        public static new GetForwardingProxyGatewayArgs Empty => new GetForwardingProxyGatewayArgs();
    }

    public sealed class GetForwardingProxyGatewayInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the forwarding control Proxy Gateway resource.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// The name of the forwarding control Proxy Gateway to be exported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetForwardingProxyGatewayInvokeArgs()
        {
        }
        public static new GetForwardingProxyGatewayInvokeArgs Empty => new GetForwardingProxyGatewayInvokeArgs();
    }


    [OutputType]
    public sealed class GetForwardingProxyGatewayResult
    {
        /// <summary>
        /// (string) - Additional details about the Proxy gateway
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (Boolean) - Indicates whether fail close is enabled to drop the traffic or disabled to allow the traffic when both primary and secondary proxies defined in this gateway are unreachable.
        /// </summary>
        public readonly bool FailClosed;
        /// <summary>
        /// (string) A unique identifier for the secondary proxy gateway
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// (list) -  Information about the admin user that last modified the Proxy gateway
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingProxyGatewayLastModifiedByResult> LastModifiedBies;
        /// <summary>
        /// (int) - Timestamp when the ZPA gateway was last modified
        /// </summary>
        public readonly int LastModifiedTime;
        /// <summary>
        /// (string) The configured name for the secondary proxy gateway
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// (Set of String) - The primary proxy for the gateway. This field is not applicable to the Lite API.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingProxyGatewayPrimaryProxyResult> PrimaryProxies;
        /// <summary>
        /// () - The secondary proxy for the gateway. This field is not applicable to the Lite API.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetForwardingProxyGatewaySecondaryProxyResult> SecondaryProxies;
        /// <summary>
        /// (string) - Indicates whether the type of Proxy gateway. Returned values are: `PROXYCHAIN`, `ZIA`, or `ECSELF`
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetForwardingProxyGatewayResult(
            string description,

            bool failClosed,

            int id,

            ImmutableArray<Outputs.GetForwardingProxyGatewayLastModifiedByResult> lastModifiedBies,

            int lastModifiedTime,

            string name,

            ImmutableArray<Outputs.GetForwardingProxyGatewayPrimaryProxyResult> primaryProxies,

            ImmutableArray<Outputs.GetForwardingProxyGatewaySecondaryProxyResult> secondaryProxies,

            string type)
        {
            Description = description;
            FailClosed = failClosed;
            Id = id;
            LastModifiedBies = lastModifiedBies;
            LastModifiedTime = lastModifiedTime;
            Name = name;
            PrimaryProxies = primaryProxies;
            SecondaryProxies = secondaryProxies;
            Type = type;
        }
    }
}
