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
    /// The **zia_sandbox_behavioral_analysis** resource updates the custom list of MD5 file hashes that are blocked by Sandbox. This overwrites a previously generated blocklist. If you need to completely erase the blocklist, submit an empty list.
    /// 
    /// **Note**: Only the file types that are supported by Sandbox analysis can be blocked using MD5 hashes.
    /// 
    /// ## Example Usage
    /// 
    /// ### Add MD5 Hashes To Sandbox
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zia = zscaler.PulumiPackage.Zia;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Add MD5 Hashes to Sandbox
    ///     var @this = new Zia.SandboxBehavioralAnalysis("this", new()
    ///     {
    ///         FileHashesToBeBlockeds = new[]
    ///         {
    ///             "42914d6d213a20a2684064be5c80ffa9",
    ///             "c0202cf6aeab8437c638533d14563d35",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Remove All MD5 Hashes To Sandbox
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zia = zscaler.PulumiPackage.Zia;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Remove All MD5 Hashes to Sandbox
    ///     var @this = new Zia.SandboxBehavioralAnalysis("this", new()
    ///     {
    ///         FileHashesToBeBlockeds = new[] {},
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.
    /// 
    /// Visit
    /// 
    /// **zia_sandbox_behavioral_analysis** can be imported by using `sandbox_settings` as the import ID.
    /// 
    /// For example:
    /// 
    /// ```sh
    /// $ pulumi import zia:index/sandboxBehavioralAnalysis:SandboxBehavioralAnalysis example sandbox_settings
    /// ```
    /// </summary>
    [ZiaResourceType("zia:index/sandboxBehavioralAnalysis:SandboxBehavioralAnalysis")]
    public partial class SandboxBehavioralAnalysis : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A custom list of unique MD5 file hashes that must be blocked by Sandbox. A maximum of 10000 MD5 file hashes can be
        /// blocked
        /// </summary>
        [Output("fileHashesToBeBlockeds")]
        public Output<ImmutableArray<string>> FileHashesToBeBlockeds { get; private set; } = null!;


        /// <summary>
        /// Create a SandboxBehavioralAnalysis resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SandboxBehavioralAnalysis(string name, SandboxBehavioralAnalysisArgs? args = null, CustomResourceOptions? options = null)
            : base("zia:index/sandboxBehavioralAnalysis:SandboxBehavioralAnalysis", name, args ?? new SandboxBehavioralAnalysisArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SandboxBehavioralAnalysis(string name, Input<string> id, SandboxBehavioralAnalysisState? state = null, CustomResourceOptions? options = null)
            : base("zia:index/sandboxBehavioralAnalysis:SandboxBehavioralAnalysis", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SandboxBehavioralAnalysis resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SandboxBehavioralAnalysis Get(string name, Input<string> id, SandboxBehavioralAnalysisState? state = null, CustomResourceOptions? options = null)
        {
            return new SandboxBehavioralAnalysis(name, id, state, options);
        }
    }

    public sealed class SandboxBehavioralAnalysisArgs : global::Pulumi.ResourceArgs
    {
        [Input("fileHashesToBeBlockeds")]
        private InputList<string>? _fileHashesToBeBlockeds;

        /// <summary>
        /// A custom list of unique MD5 file hashes that must be blocked by Sandbox. A maximum of 10000 MD5 file hashes can be
        /// blocked
        /// </summary>
        public InputList<string> FileHashesToBeBlockeds
        {
            get => _fileHashesToBeBlockeds ?? (_fileHashesToBeBlockeds = new InputList<string>());
            set => _fileHashesToBeBlockeds = value;
        }

        public SandboxBehavioralAnalysisArgs()
        {
        }
        public static new SandboxBehavioralAnalysisArgs Empty => new SandboxBehavioralAnalysisArgs();
    }

    public sealed class SandboxBehavioralAnalysisState : global::Pulumi.ResourceArgs
    {
        [Input("fileHashesToBeBlockeds")]
        private InputList<string>? _fileHashesToBeBlockeds;

        /// <summary>
        /// A custom list of unique MD5 file hashes that must be blocked by Sandbox. A maximum of 10000 MD5 file hashes can be
        /// blocked
        /// </summary>
        public InputList<string> FileHashesToBeBlockeds
        {
            get => _fileHashesToBeBlockeds ?? (_fileHashesToBeBlockeds = new InputList<string>());
            set => _fileHashesToBeBlockeds = value;
        }

        public SandboxBehavioralAnalysisState()
        {
        }
        public static new SandboxBehavioralAnalysisState Empty => new SandboxBehavioralAnalysisState();
    }
}
