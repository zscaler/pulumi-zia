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
    public static class GetTrafficForwardingVIPRecommendedList
    {
        /// <summary>
        /// Use the **zia_gre_vip_recommended_list** data source to get information about a list of recommended GRE tunnel virtual IP addresses (VIPs), based on source IP address or latitude/longitude coordinates.
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
        ///     var @this = Zia.GetTrafficForwardingVIPRecommendedList.Invoke(new()
        ///     {
        ///         RequiredCount = 2,
        ///         SourceIp = "1.1.1.1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetTrafficForwardingVIPRecommendedListResult> InvokeAsync(GetTrafficForwardingVIPRecommendedListArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTrafficForwardingVIPRecommendedListResult>("zia:index/getTrafficForwardingVIPRecommendedList:getTrafficForwardingVIPRecommendedList", args ?? new GetTrafficForwardingVIPRecommendedListArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zia_gre_vip_recommended_list** data source to get information about a list of recommended GRE tunnel virtual IP addresses (VIPs), based on source IP address or latitude/longitude coordinates.
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
        ///     var @this = Zia.GetTrafficForwardingVIPRecommendedList.Invoke(new()
        ///     {
        ///         RequiredCount = 2,
        ///         SourceIp = "1.1.1.1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetTrafficForwardingVIPRecommendedListResult> Invoke(GetTrafficForwardingVIPRecommendedListInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTrafficForwardingVIPRecommendedListResult>("zia:index/getTrafficForwardingVIPRecommendedList:getTrafficForwardingVIPRecommendedList", args ?? new GetTrafficForwardingVIPRecommendedListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTrafficForwardingVIPRecommendedListArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Number of IP address to be exported.
        /// </summary>
        [Input("requiredCount")]
        public int? RequiredCount { get; set; }

        /// <summary>
        /// Filter based on an IP address range.
        /// </summary>
        [Input("sourceIp")]
        public string? SourceIp { get; set; }

        public GetTrafficForwardingVIPRecommendedListArgs()
        {
        }
        public static new GetTrafficForwardingVIPRecommendedListArgs Empty => new GetTrafficForwardingVIPRecommendedListArgs();
    }

    public sealed class GetTrafficForwardingVIPRecommendedListInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Number of IP address to be exported.
        /// </summary>
        [Input("requiredCount")]
        public Input<int>? RequiredCount { get; set; }

        /// <summary>
        /// Filter based on an IP address range.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        public GetTrafficForwardingVIPRecommendedListInvokeArgs()
        {
        }
        public static new GetTrafficForwardingVIPRecommendedListInvokeArgs Empty => new GetTrafficForwardingVIPRecommendedListInvokeArgs();
    }


    [OutputType]
    public sealed class GetTrafficForwardingVIPRecommendedListResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetTrafficForwardingVIPRecommendedListListResult> Lists;
        public readonly int? RequiredCount;
        /// <summary>
        /// (String) The public source IP address.
        /// </summary>
        public readonly string? SourceIp;

        [OutputConstructor]
        private GetTrafficForwardingVIPRecommendedListResult(
            string id,

            ImmutableArray<Outputs.GetTrafficForwardingVIPRecommendedListListResult> lists,

            int? requiredCount,

            string? sourceIp)
        {
            Id = id;
            Lists = lists;
            RequiredCount = requiredCount;
            SourceIp = sourceIp;
        }
    }
}
