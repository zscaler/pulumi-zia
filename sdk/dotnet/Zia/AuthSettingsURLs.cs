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
    /// The **zia_auth_settings_urls** resource alows you to add or remove a URL from the cookie authentication exempt list in the Zscaler Internet Access cloud or via the API. To learn more see [URL Format Guidelines](https://help.zscaler.com/zia/url-format-guidelines)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zia = zscaler.PulumiPackage.Zia;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // ZIA User Auth Settings Data Source
    ///     var example = new Zia.AuthSettingsURLs("example", new()
    ///     {
    ///         Urls = new[]
    ///         {
    ///             ".okta.com",
    ///             ".oktacdn.com",
    ///             ".mtls.oktapreview.com",
    ///             ".mtls.okta.com",
    ///             "d3l44rcogcb7iv.cloudfront.net",
    ///             "pac.zdxcloud.net",
    ///             ".windowsazure.com",
    ///             ".fedoraproject.org",
    ///             "login.windows.net",
    ///             "d32a6ru7mhaq0c.cloudfront.net",
    ///             ".kerberos.oktapreview.com",
    ///             ".oktapreview.com",
    ///             "login.zdxcloud.net",
    ///             "login.microsoftonline.com",
    ///             "smres.zdxcloud.net",
    ///             ".kerberos.okta.com",
    ///         },
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
    /// **zia_auth_settings_urls** can be imported by using `all_urls` as the import ID.
    /// 
    /// For example:
    /// 
    /// ```sh
    /// $ pulumi import zia:index/authSettingsURLs:AuthSettingsURLs example all_urls
    /// ```
    /// </summary>
    [ZiaResourceType("zia:index/authSettingsURLs:AuthSettingsURLs")]
    public partial class AuthSettingsURLs : global::Pulumi.CustomResource
    {
        [Output("urls")]
        public Output<ImmutableArray<string>> Urls { get; private set; } = null!;


        /// <summary>
        /// Create a AuthSettingsURLs resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthSettingsURLs(string name, AuthSettingsURLsArgs? args = null, CustomResourceOptions? options = null)
            : base("zia:index/authSettingsURLs:AuthSettingsURLs", name, args ?? new AuthSettingsURLsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthSettingsURLs(string name, Input<string> id, AuthSettingsURLsState? state = null, CustomResourceOptions? options = null)
            : base("zia:index/authSettingsURLs:AuthSettingsURLs", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthSettingsURLs resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthSettingsURLs Get(string name, Input<string> id, AuthSettingsURLsState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthSettingsURLs(name, id, state, options);
        }
    }

    public sealed class AuthSettingsURLsArgs : global::Pulumi.ResourceArgs
    {
        [Input("urls")]
        private InputList<string>? _urls;
        public InputList<string> Urls
        {
            get => _urls ?? (_urls = new InputList<string>());
            set => _urls = value;
        }

        public AuthSettingsURLsArgs()
        {
        }
        public static new AuthSettingsURLsArgs Empty => new AuthSettingsURLsArgs();
    }

    public sealed class AuthSettingsURLsState : global::Pulumi.ResourceArgs
    {
        [Input("urls")]
        private InputList<string>? _urls;
        public InputList<string> Urls
        {
            get => _urls ?? (_urls = new InputList<string>());
            set => _urls = value;
        }

        public AuthSettingsURLsState()
        {
        }
        public static new AuthSettingsURLsState Empty => new AuthSettingsURLsState();
    }
}
