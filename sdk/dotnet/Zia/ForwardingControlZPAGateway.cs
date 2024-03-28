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
    /// <summary>
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zia = zscaler.PulumiPackage.Zia;
    /// using Zpa = Pulumi.Zpa;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var thisServerGroup = Zpa.GetServerGroup.Invoke(new()
    ///     {
    ///         Name = "Server_Group_IP_Source_Anchoring",
    ///     });
    /// 
    ///     var this1 = Zpa.GetApplicationSegment.Invoke(new()
    ///     {
    ///         Name = "App_Segment_IP_Source_Anchoring",
    ///     });
    /// 
    ///     var this2 = Zpa.GetApplicationSegment.Invoke(new()
    ///     {
    ///         Name = "App_Segment_IP_Source_Anchoring2",
    ///     });
    /// 
    ///     var thisForwardingControlZPAGateway = new Zia.ForwardingControlZPAGateway("thisForwardingControlZPAGateway", new()
    ///     {
    ///         Description = "ZPA_GW01",
    ///         Type = "ZPA",
    ///         ZpaServerGroup = new Zia.Inputs.ForwardingControlZPAGatewayZpaServerGroupArgs
    ///         {
    ///             ExternalId = thisServerGroup.Apply(getServerGroupResult =&gt; getServerGroupResult.Id),
    ///             Name = thisServerGroup.Apply(getServerGroupResult =&gt; getServerGroupResult.Id),
    ///         },
    ///         ZpaAppSegments = new[]
    ///         {
    ///             new Zia.Inputs.ForwardingControlZPAGatewayZpaAppSegmentArgs
    ///             {
    ///                 ExternalId = this1.Apply(getApplicationSegmentResult =&gt; getApplicationSegmentResult.Id),
    ///                 Name = this1.Apply(getApplicationSegmentResult =&gt; getApplicationSegmentResult.Name),
    ///             },
    ///             new Zia.Inputs.ForwardingControlZPAGatewayZpaAppSegmentArgs
    ///             {
    ///                 ExternalId = this2.Apply(getApplicationSegmentResult =&gt; getApplicationSegmentResult.Id),
    ///                 Name = this2.Apply(getApplicationSegmentResult =&gt; getApplicationSegmentResult.Name),
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.
    /// 
    /// Visit
    /// 
    /// **forwarding_control_zpa_gateway** can be imported by using `&lt;GATEWAY_ID&gt;` or `&lt;GATEWAY_NAME&gt;` as the import ID.
    /// 
    /// For example:
    /// 
    /// ```sh
    /// $ pulumi import zia:index/forwardingControlZPAGateway:ForwardingControlZPAGateway example &lt;gateway_id&gt;
    /// ```
    /// 
    /// or
    /// 
    /// ```sh
    /// $ pulumi import zia:index/forwardingControlZPAGateway:ForwardingControlZPAGateway example &lt;gateway_name&gt;
    /// ```
    /// </summary>
    [ZiaResourceType("zia:index/forwardingControlZPAGateway:ForwardingControlZPAGateway")]
    public partial class ForwardingControlZPAGateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (string) - Additional details about the ZPA gateway
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("gatewayId")]
        public Output<int> GatewayId { get; private set; } = null!;

        /// <summary>
        /// The configured name of the entity
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// (string) - Indicates whether the ZPA gateway is configured for Zscaler Internet Access (using option ZPA) or Zscaler Cloud Connector (using option ECZPA). Supported values: ``ZPA`` and ``ECZPA``
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// The ZPA Server Group that is configured for Source IP Anchoring
        /// </summary>
        [Output("zpaAppSegments")]
        public Output<ImmutableArray<Outputs.ForwardingControlZPAGatewayZpaAppSegment>> ZpaAppSegments { get; private set; } = null!;

        /// <summary>
        /// The ZPA Server Group that is configured for Source IP Anchoring
        /// </summary>
        [Output("zpaServerGroup")]
        public Output<Outputs.ForwardingControlZPAGatewayZpaServerGroup> ZpaServerGroup { get; private set; } = null!;


        /// <summary>
        /// Create a ForwardingControlZPAGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ForwardingControlZPAGateway(string name, ForwardingControlZPAGatewayArgs args, CustomResourceOptions? options = null)
            : base("zia:index/forwardingControlZPAGateway:ForwardingControlZPAGateway", name, args ?? new ForwardingControlZPAGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ForwardingControlZPAGateway(string name, Input<string> id, ForwardingControlZPAGatewayState? state = null, CustomResourceOptions? options = null)
            : base("zia:index/forwardingControlZPAGateway:ForwardingControlZPAGateway", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/zscaler",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ForwardingControlZPAGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ForwardingControlZPAGateway Get(string name, Input<string> id, ForwardingControlZPAGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new ForwardingControlZPAGateway(name, id, state, options);
        }
    }

    public sealed class ForwardingControlZPAGatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (string) - Additional details about the ZPA gateway
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The configured name of the entity
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (string) - Indicates whether the ZPA gateway is configured for Zscaler Internet Access (using option ZPA) or Zscaler Cloud Connector (using option ECZPA). Supported values: ``ZPA`` and ``ECZPA``
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("zpaAppSegments", required: true)]
        private InputList<Inputs.ForwardingControlZPAGatewayZpaAppSegmentArgs>? _zpaAppSegments;

        /// <summary>
        /// The ZPA Server Group that is configured for Source IP Anchoring
        /// </summary>
        public InputList<Inputs.ForwardingControlZPAGatewayZpaAppSegmentArgs> ZpaAppSegments
        {
            get => _zpaAppSegments ?? (_zpaAppSegments = new InputList<Inputs.ForwardingControlZPAGatewayZpaAppSegmentArgs>());
            set => _zpaAppSegments = value;
        }

        /// <summary>
        /// The ZPA Server Group that is configured for Source IP Anchoring
        /// </summary>
        [Input("zpaServerGroup", required: true)]
        public Input<Inputs.ForwardingControlZPAGatewayZpaServerGroupArgs> ZpaServerGroup { get; set; } = null!;

        public ForwardingControlZPAGatewayArgs()
        {
        }
        public static new ForwardingControlZPAGatewayArgs Empty => new ForwardingControlZPAGatewayArgs();
    }

    public sealed class ForwardingControlZPAGatewayState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (string) - Additional details about the ZPA gateway
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("gatewayId")]
        public Input<int>? GatewayId { get; set; }

        /// <summary>
        /// The configured name of the entity
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (string) - Indicates whether the ZPA gateway is configured for Zscaler Internet Access (using option ZPA) or Zscaler Cloud Connector (using option ECZPA). Supported values: ``ZPA`` and ``ECZPA``
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("zpaAppSegments")]
        private InputList<Inputs.ForwardingControlZPAGatewayZpaAppSegmentGetArgs>? _zpaAppSegments;

        /// <summary>
        /// The ZPA Server Group that is configured for Source IP Anchoring
        /// </summary>
        public InputList<Inputs.ForwardingControlZPAGatewayZpaAppSegmentGetArgs> ZpaAppSegments
        {
            get => _zpaAppSegments ?? (_zpaAppSegments = new InputList<Inputs.ForwardingControlZPAGatewayZpaAppSegmentGetArgs>());
            set => _zpaAppSegments = value;
        }

        /// <summary>
        /// The ZPA Server Group that is configured for Source IP Anchoring
        /// </summary>
        [Input("zpaServerGroup")]
        public Input<Inputs.ForwardingControlZPAGatewayZpaServerGroupGetArgs>? ZpaServerGroup { get; set; }

        public ForwardingControlZPAGatewayState()
        {
        }
        public static new ForwardingControlZPAGatewayState Empty => new ForwardingControlZPAGatewayState();
    }
}
