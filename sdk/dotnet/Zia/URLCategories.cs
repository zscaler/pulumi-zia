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
    /// The **zia_url_categories** resource creates and manages a new custom URL category. If keywords are included within the request, they will be added to the new category.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zia = zscaler.PulumiPackage.Zia;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Zia.URLCategories("example", new()
    ///     {
    ///         SuperCategory = "USER_DEFINED",
    ///         ConfiguredName = "MCAS Unsanctioned Apps",
    ///         Description = "MCAS Unsanctioned Apps",
    ///         Keywords = new[]
    ///         {
    ///             "microsoft",
    ///         },
    ///         CustomCategory = true,
    ///         Type = "URL_CATEGORY",
    ///         Scopes = new[]
    ///         {
    ///             new Zia.Inputs.URLCategoriesScopeArgs
    ///             {
    ///                 Type = "LOCATION",
    ///                 ScopeEntities = new Zia.Inputs.URLCategoriesScopeScopeEntitiesArgs
    ///                 {
    ///                     Ids = new[]
    ///                     {
    ///                         data.Zia_location_management.Nyc_site.Id,
    ///                     },
    ///                 },
    ///                 ScopeGroupMemberEntities = new Zia.Inputs.URLCategoriesScopeScopeGroupMemberEntitiesArgs
    ///                 {
    ///                     Ids = new[]
    ///                     {
    ///                         data.Zia_group_management.Engineering.Id,
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         Urls = new[]
    ///         {
    ///             ".coupons.com",
    ///             ".resource.alaskaair.net",
    ///             ".techrepublic.com",
    ///             ".dailymotion.com",
    ///             ".osiriscomm.com",
    ///             ".uefa.com",
    ///             ".Logz.io",
    ///             ".alexa.com",
    ///             ".baidu.com",
    ///             ".cnn.com",
    ///             ".level3.com",
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
    /// **zia_url_categories** can be imported by using `&lt;CATEGORY_ID&gt;` or `&lt;CATEGORY_NAME&gt;` as the import ID.
    /// 
    /// For example:
    /// 
    /// ```sh
    /// $ pulumi import zia:index/uRLCategories:URLCategories example &lt;category_id&gt;
    /// ```
    /// 
    /// or
    /// 
    /// ```sh
    /// $ pulumi import zia:index/uRLCategories:URLCategories example &lt;category_name&gt;
    /// ```
    /// 
    /// ⚠️ **NOTE :**:  This provider only supports the importing of custom URL categories. The importing of built-in categories is not supported.
    /// </summary>
    [ZiaResourceType("zia:index/uRLCategories:URLCategories")]
    public partial class URLCategories : global::Pulumi.CustomResource
    {
        [Output("categoryId")]
        public Output<string> CategoryId { get; private set; } = null!;

        /// <summary>
        /// Name of the URL category. This is only required for custom URL categories.
        /// </summary>
        [Output("configuredName")]
        public Output<string?> ConfiguredName { get; private set; } = null!;

        /// <summary>
        /// Set to true for custom URL category. Up to 48 custom URL categories can be added per organization.
        /// </summary>
        [Output("customCategory")]
        public Output<bool?> CustomCategory { get; private set; } = null!;

        /// <summary>
        /// The number of custom IP address ranges associated to the URL category.
        /// </summary>
        [Output("customIpRangesCount")]
        public Output<int> CustomIpRangesCount { get; private set; } = null!;

        /// <summary>
        /// The number of custom URLs associated to the URL category.
        /// </summary>
        [Output("customUrlsCount")]
        public Output<int> CustomUrlsCount { get; private set; } = null!;

        /// <summary>
        /// URLs added to a custom URL category are also retained under the original parent URL category (i.e., the predefined category the URL previously belonged to).
        /// </summary>
        [Output("dbCategorizedUrls")]
        public Output<ImmutableArray<string>> DbCategorizedUrls { get; private set; } = null!;

        /// <summary>
        /// Description of the category.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Value is set to false for custom URL category when due to scope user does not have edit permission
        /// </summary>
        [Output("editable")]
        public Output<bool> Editable { get; private set; } = null!;

        /// <summary>
        /// Custom IP address ranges associated to a URL category. Up to 2000 custom IP address ranges and retaining parent custom IP address ranges can be added, per organization, across all categories.
        /// ⚠️ **NOTE :**: This field is available only if the option to configure custom IP ranges is enabled for your organization. To enable this option, contact Zscaler Support.
        /// </summary>
        [Output("ipRanges")]
        public Output<ImmutableArray<string>> IpRanges { get; private set; } = null!;

        /// <summary>
        /// The retaining parent custom IP address ranges associated to a URL category. Up to 2000 custom IP ranges and retaining parent custom IP address ranges can be added, per organization, across all categories.
        /// </summary>
        [Output("ipRangesRetainingParentCategories")]
        public Output<ImmutableArray<string>> IpRangesRetainingParentCategories { get; private set; } = null!;

        /// <summary>
        /// The number of custom IP address ranges associated to the URL category, that also need to be retained under the original parent category.
        /// </summary>
        [Output("ipRangesRetainingParentCategoryCount")]
        public Output<int> IpRangesRetainingParentCategoryCount { get; private set; } = null!;

        /// <summary>
        /// Custom keywords associated to a URL category. Up to 2048 custom keywords can be added per organization across all categories (including bandwidth classes).
        /// </summary>
        [Output("keywords")]
        public Output<ImmutableArray<string>> Keywords { get; private set; } = null!;

        [Output("keywordsRetainingParentCategories")]
        public Output<ImmutableArray<string>> KeywordsRetainingParentCategories { get; private set; } = null!;

        /// <summary>
        /// Scope of the custom categories.
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<Outputs.URLCategoriesScope>> Scopes { get; private set; } = null!;

        /// <summary>
        /// Super Category of the URL category. This field is required when creating custom URL categories.
        /// </summary>
        [Output("superCategory")]
        public Output<string?> SuperCategory { get; private set; } = null!;

        /// <summary>
        /// The admin scope type. The attribute name is subject to change. `ORGANIZATION`, `DEPARTMENT`, `LOCATION`, `LOCATION_GROUP`
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// URL and keyword counts for the category.
        /// </summary>
        [Output("urlKeywordCounts")]
        public Output<Outputs.URLCategoriesUrlKeywordCounts> UrlKeywordCounts { get; private set; } = null!;

        /// <summary>
        /// Custom URLs to add to a URL category. Up to 25,000 custom URLs can be added per organization across all categories (including bandwidth classes).
        /// </summary>
        [Output("urls")]
        public Output<ImmutableArray<string>> Urls { get; private set; } = null!;

        /// <summary>
        /// The number of custom IP address ranges associated to the URL category, that also need to be retained under the original parent category.
        /// </summary>
        [Output("urlsRetainingParentCategoryCount")]
        public Output<int> UrlsRetainingParentCategoryCount { get; private set; } = null!;


        /// <summary>
        /// Create a URLCategories resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public URLCategories(string name, URLCategoriesArgs? args = null, CustomResourceOptions? options = null)
            : base("zia:index/uRLCategories:URLCategories", name, args ?? new URLCategoriesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private URLCategories(string name, Input<string> id, URLCategoriesState? state = null, CustomResourceOptions? options = null)
            : base("zia:index/uRLCategories:URLCategories", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing URLCategories resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static URLCategories Get(string name, Input<string> id, URLCategoriesState? state = null, CustomResourceOptions? options = null)
        {
            return new URLCategories(name, id, state, options);
        }
    }

    public sealed class URLCategoriesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the URL category. This is only required for custom URL categories.
        /// </summary>
        [Input("configuredName")]
        public Input<string>? ConfiguredName { get; set; }

        /// <summary>
        /// Set to true for custom URL category. Up to 48 custom URL categories can be added per organization.
        /// </summary>
        [Input("customCategory")]
        public Input<bool>? CustomCategory { get; set; }

        /// <summary>
        /// The number of custom IP address ranges associated to the URL category.
        /// </summary>
        [Input("customIpRangesCount")]
        public Input<int>? CustomIpRangesCount { get; set; }

        /// <summary>
        /// The number of custom URLs associated to the URL category.
        /// </summary>
        [Input("customUrlsCount")]
        public Input<int>? CustomUrlsCount { get; set; }

        [Input("dbCategorizedUrls")]
        private InputList<string>? _dbCategorizedUrls;

        /// <summary>
        /// URLs added to a custom URL category are also retained under the original parent URL category (i.e., the predefined category the URL previously belonged to).
        /// </summary>
        public InputList<string> DbCategorizedUrls
        {
            get => _dbCategorizedUrls ?? (_dbCategorizedUrls = new InputList<string>());
            set => _dbCategorizedUrls = value;
        }

        /// <summary>
        /// Description of the category.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Value is set to false for custom URL category when due to scope user does not have edit permission
        /// </summary>
        [Input("editable")]
        public Input<bool>? Editable { get; set; }

        [Input("ipRanges")]
        private InputList<string>? _ipRanges;

        /// <summary>
        /// Custom IP address ranges associated to a URL category. Up to 2000 custom IP address ranges and retaining parent custom IP address ranges can be added, per organization, across all categories.
        /// ⚠️ **NOTE :**: This field is available only if the option to configure custom IP ranges is enabled for your organization. To enable this option, contact Zscaler Support.
        /// </summary>
        public InputList<string> IpRanges
        {
            get => _ipRanges ?? (_ipRanges = new InputList<string>());
            set => _ipRanges = value;
        }

        [Input("ipRangesRetainingParentCategories")]
        private InputList<string>? _ipRangesRetainingParentCategories;

        /// <summary>
        /// The retaining parent custom IP address ranges associated to a URL category. Up to 2000 custom IP ranges and retaining parent custom IP address ranges can be added, per organization, across all categories.
        /// </summary>
        public InputList<string> IpRangesRetainingParentCategories
        {
            get => _ipRangesRetainingParentCategories ?? (_ipRangesRetainingParentCategories = new InputList<string>());
            set => _ipRangesRetainingParentCategories = value;
        }

        /// <summary>
        /// The number of custom IP address ranges associated to the URL category, that also need to be retained under the original parent category.
        /// </summary>
        [Input("ipRangesRetainingParentCategoryCount")]
        public Input<int>? IpRangesRetainingParentCategoryCount { get; set; }

        [Input("keywords")]
        private InputList<string>? _keywords;

        /// <summary>
        /// Custom keywords associated to a URL category. Up to 2048 custom keywords can be added per organization across all categories (including bandwidth classes).
        /// </summary>
        public InputList<string> Keywords
        {
            get => _keywords ?? (_keywords = new InputList<string>());
            set => _keywords = value;
        }

        [Input("keywordsRetainingParentCategories")]
        private InputList<string>? _keywordsRetainingParentCategories;
        public InputList<string> KeywordsRetainingParentCategories
        {
            get => _keywordsRetainingParentCategories ?? (_keywordsRetainingParentCategories = new InputList<string>());
            set => _keywordsRetainingParentCategories = value;
        }

        [Input("scopes")]
        private InputList<Inputs.URLCategoriesScopeArgs>? _scopes;

        /// <summary>
        /// Scope of the custom categories.
        /// </summary>
        public InputList<Inputs.URLCategoriesScopeArgs> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<Inputs.URLCategoriesScopeArgs>());
            set => _scopes = value;
        }

        /// <summary>
        /// Super Category of the URL category. This field is required when creating custom URL categories.
        /// </summary>
        [Input("superCategory")]
        public Input<string>? SuperCategory { get; set; }

        /// <summary>
        /// The admin scope type. The attribute name is subject to change. `ORGANIZATION`, `DEPARTMENT`, `LOCATION`, `LOCATION_GROUP`
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// URL and keyword counts for the category.
        /// </summary>
        [Input("urlKeywordCounts")]
        public Input<Inputs.URLCategoriesUrlKeywordCountsArgs>? UrlKeywordCounts { get; set; }

        [Input("urls")]
        private InputList<string>? _urls;

        /// <summary>
        /// Custom URLs to add to a URL category. Up to 25,000 custom URLs can be added per organization across all categories (including bandwidth classes).
        /// </summary>
        public InputList<string> Urls
        {
            get => _urls ?? (_urls = new InputList<string>());
            set => _urls = value;
        }

        /// <summary>
        /// The number of custom IP address ranges associated to the URL category, that also need to be retained under the original parent category.
        /// </summary>
        [Input("urlsRetainingParentCategoryCount")]
        public Input<int>? UrlsRetainingParentCategoryCount { get; set; }

        public URLCategoriesArgs()
        {
        }
        public static new URLCategoriesArgs Empty => new URLCategoriesArgs();
    }

    public sealed class URLCategoriesState : global::Pulumi.ResourceArgs
    {
        [Input("categoryId")]
        public Input<string>? CategoryId { get; set; }

        /// <summary>
        /// Name of the URL category. This is only required for custom URL categories.
        /// </summary>
        [Input("configuredName")]
        public Input<string>? ConfiguredName { get; set; }

        /// <summary>
        /// Set to true for custom URL category. Up to 48 custom URL categories can be added per organization.
        /// </summary>
        [Input("customCategory")]
        public Input<bool>? CustomCategory { get; set; }

        /// <summary>
        /// The number of custom IP address ranges associated to the URL category.
        /// </summary>
        [Input("customIpRangesCount")]
        public Input<int>? CustomIpRangesCount { get; set; }

        /// <summary>
        /// The number of custom URLs associated to the URL category.
        /// </summary>
        [Input("customUrlsCount")]
        public Input<int>? CustomUrlsCount { get; set; }

        [Input("dbCategorizedUrls")]
        private InputList<string>? _dbCategorizedUrls;

        /// <summary>
        /// URLs added to a custom URL category are also retained under the original parent URL category (i.e., the predefined category the URL previously belonged to).
        /// </summary>
        public InputList<string> DbCategorizedUrls
        {
            get => _dbCategorizedUrls ?? (_dbCategorizedUrls = new InputList<string>());
            set => _dbCategorizedUrls = value;
        }

        /// <summary>
        /// Description of the category.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Value is set to false for custom URL category when due to scope user does not have edit permission
        /// </summary>
        [Input("editable")]
        public Input<bool>? Editable { get; set; }

        [Input("ipRanges")]
        private InputList<string>? _ipRanges;

        /// <summary>
        /// Custom IP address ranges associated to a URL category. Up to 2000 custom IP address ranges and retaining parent custom IP address ranges can be added, per organization, across all categories.
        /// ⚠️ **NOTE :**: This field is available only if the option to configure custom IP ranges is enabled for your organization. To enable this option, contact Zscaler Support.
        /// </summary>
        public InputList<string> IpRanges
        {
            get => _ipRanges ?? (_ipRanges = new InputList<string>());
            set => _ipRanges = value;
        }

        [Input("ipRangesRetainingParentCategories")]
        private InputList<string>? _ipRangesRetainingParentCategories;

        /// <summary>
        /// The retaining parent custom IP address ranges associated to a URL category. Up to 2000 custom IP ranges and retaining parent custom IP address ranges can be added, per organization, across all categories.
        /// </summary>
        public InputList<string> IpRangesRetainingParentCategories
        {
            get => _ipRangesRetainingParentCategories ?? (_ipRangesRetainingParentCategories = new InputList<string>());
            set => _ipRangesRetainingParentCategories = value;
        }

        /// <summary>
        /// The number of custom IP address ranges associated to the URL category, that also need to be retained under the original parent category.
        /// </summary>
        [Input("ipRangesRetainingParentCategoryCount")]
        public Input<int>? IpRangesRetainingParentCategoryCount { get; set; }

        [Input("keywords")]
        private InputList<string>? _keywords;

        /// <summary>
        /// Custom keywords associated to a URL category. Up to 2048 custom keywords can be added per organization across all categories (including bandwidth classes).
        /// </summary>
        public InputList<string> Keywords
        {
            get => _keywords ?? (_keywords = new InputList<string>());
            set => _keywords = value;
        }

        [Input("keywordsRetainingParentCategories")]
        private InputList<string>? _keywordsRetainingParentCategories;
        public InputList<string> KeywordsRetainingParentCategories
        {
            get => _keywordsRetainingParentCategories ?? (_keywordsRetainingParentCategories = new InputList<string>());
            set => _keywordsRetainingParentCategories = value;
        }

        [Input("scopes")]
        private InputList<Inputs.URLCategoriesScopeGetArgs>? _scopes;

        /// <summary>
        /// Scope of the custom categories.
        /// </summary>
        public InputList<Inputs.URLCategoriesScopeGetArgs> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<Inputs.URLCategoriesScopeGetArgs>());
            set => _scopes = value;
        }

        /// <summary>
        /// Super Category of the URL category. This field is required when creating custom URL categories.
        /// </summary>
        [Input("superCategory")]
        public Input<string>? SuperCategory { get; set; }

        /// <summary>
        /// The admin scope type. The attribute name is subject to change. `ORGANIZATION`, `DEPARTMENT`, `LOCATION`, `LOCATION_GROUP`
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// URL and keyword counts for the category.
        /// </summary>
        [Input("urlKeywordCounts")]
        public Input<Inputs.URLCategoriesUrlKeywordCountsGetArgs>? UrlKeywordCounts { get; set; }

        [Input("urls")]
        private InputList<string>? _urls;

        /// <summary>
        /// Custom URLs to add to a URL category. Up to 25,000 custom URLs can be added per organization across all categories (including bandwidth classes).
        /// </summary>
        public InputList<string> Urls
        {
            get => _urls ?? (_urls = new InputList<string>());
            set => _urls = value;
        }

        /// <summary>
        /// The number of custom IP address ranges associated to the URL category, that also need to be retained under the original parent category.
        /// </summary>
        [Input("urlsRetainingParentCategoryCount")]
        public Input<int>? UrlsRetainingParentCategoryCount { get; set; }

        public URLCategoriesState()
        {
        }
        public static new URLCategoriesState Empty => new URLCategoriesState();
    }
}
