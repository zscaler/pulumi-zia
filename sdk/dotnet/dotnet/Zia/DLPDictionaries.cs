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
    /// <summary>
    /// * [Official documentation](https://help.zscaler.com/zia/adding-custom-dlp-dictionary)
    /// * [API documentation](https://help.zscaler.com/zia/data-loss-prevention#/dlpDictionaries-post)
    /// 
    /// The **zia_dlp_dictionaries** resource allows the creation and management of ZIA DLP dictionaries in the Zscaler Internet Access cloud or via the API.
    /// 
    /// ## Example Usage
    /// 
    /// ### With Hierarchical Identifiers
    /// 
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.
    /// 
    /// Visit
    /// 
    /// **zia_dlp_dictionaries** can be imported by using `&lt;DICTIONARY ID&gt;` or `&lt;DICTIONARY_NAME&gt;` as the import ID.
    /// 
    /// For example:
    /// 
    /// ```sh
    /// $ pulumi import zia:index/dLPDictionaries:DLPDictionaries example &lt;dictionary_id&gt;
    /// ```
    /// 
    /// or
    /// 
    /// ```sh
    /// $ pulumi import zia:index/dLPDictionaries:DLPDictionaries example &lt;dictionary_name&gt;
    /// ```
    /// </summary>
    [ZiaResourceType("zia:index/dLPDictionaries:DLPDictionaries")]
    public partial class DLPDictionaries : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The list of Bank Identification Number (BIN) values that are included or excluded from the Credit Cards dictionary. BIN
        /// values can be specified only for Diners Club, Mastercard, RuPay, and Visa cards. Up to 512 BIN values can be configured
        /// in a dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
        /// </summary>
        [Output("binNumbers")]
        public Output<ImmutableArray<int>> BinNumbers { get; private set; } = null!;

        /// <summary>
        /// The DLP confidence threshold for predefined dictionaries
        /// </summary>
        [Output("confidenceLevelForPredefinedDict")]
        public Output<string?> ConfidenceLevelForPredefinedDict { get; private set; } = null!;

        /// <summary>
        /// The DLP confidence threshold
        /// </summary>
        [Output("confidenceThreshold")]
        public Output<string?> ConfidenceThreshold { get; private set; } = null!;

        /// <summary>
        /// The DLP dictionary proximity length.
        /// </summary>
        [Output("custom")]
        public Output<bool> Custom { get; private set; } = null!;

        [Output("customPhraseMatchType")]
        public Output<string> CustomPhraseMatchType { get; private set; } = null!;

        /// <summary>
        /// The desciption of the DLP dictionary
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// ID of the predefined dictionary (original source dictionary) that is used for cloning. This field is applicable only to
        /// cloned dictionaries. Only a limited set of identification-based predefined dictionaries (e.g., Credit Cards, Social
        /// Security Numbers, National Identification Numbers, etc.) can be cloned. Up to 4 clones can be created from a predefined
        /// dictionary.
        /// </summary>
        [Output("dictTemplateId")]
        public Output<int?> DictTemplateId { get; private set; } = null!;

        [Output("dictionaryId")]
        public Output<int> DictionaryId { get; private set; } = null!;

        /// <summary>
        /// The DLP dictionary type.
        /// </summary>
        [Output("dictionaryType")]
        public Output<string?> DictionaryType { get; private set; } = null!;

        /// <summary>
        /// Exact Data Match (EDM) related information for custom DLP dictionaries.
        /// </summary>
        [Output("exactDataMatchDetails")]
        public Output<ImmutableArray<Outputs.DLPDictionariesExactDataMatchDetail>> ExactDataMatchDetails { get; private set; } = null!;

        /// <summary>
        /// List of hierarchical identifiers for the DLP dictionary.
        /// </summary>
        [Output("hierarchicalIdentifiers")]
        public Output<ImmutableArray<string>> HierarchicalIdentifiers { get; private set; } = null!;

        /// <summary>
        /// List of Indexed Document Match (IDM) profiles and their corresponding match accuracy for custom DLP dictionaries.
        /// </summary>
        [Output("idmProfileMatchAccuracies")]
        public Output<ImmutableArray<Outputs.DLPDictionariesIdmProfileMatchAccuracy>> IdmProfileMatchAccuracies { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to exclude documents that are a 100% match to already-indexed documents from triggering an Indexed
        /// Document Match (IDM) Dictionary.
        /// </summary>
        [Output("ignoreExactMatchIdmDict")]
        public Output<bool?> IgnoreExactMatchIdmDict { get; private set; } = null!;

        /// <summary>
        /// A true value denotes that the specified Bank Identification Number (BIN) values are included in the Credit Cards
        /// dictionary. A false value denotes that the specified BIN values are excluded from the Credit Cards dictionary.Note: This
        /// field is applicable only to the predefined Credit Cards dictionary and its clones.
        /// </summary>
        [Output("includeBinNumbers")]
        public Output<bool> IncludeBinNumbers { get; private set; } = null!;

        /// <summary>
        /// The DLP dictionary's name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List containing the patterns used within a custom DLP dictionary. This attribute is not applicable to predefined DLP
        /// dictionaries
        /// </summary>
        [Output("patterns")]
        public Output<ImmutableArray<Outputs.DLPDictionariesPattern>> Patterns { get; private set; } = null!;

        [Output("phrases")]
        public Output<ImmutableArray<Outputs.DLPDictionariesPhrase>> Phrases { get; private set; } = null!;

        /// <summary>
        /// The DLP dictionary proximity length.
        /// </summary>
        [Output("proximity")]
        public Output<int?> Proximity { get; private set; } = null!;


        /// <summary>
        /// Create a DLPDictionaries resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DLPDictionaries(string name, DLPDictionariesArgs? args = null, CustomResourceOptions? options = null)
            : base("zia:index/dLPDictionaries:DLPDictionaries", name, args ?? new DLPDictionariesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DLPDictionaries(string name, Input<string> id, DLPDictionariesState? state = null, CustomResourceOptions? options = null)
            : base("zia:index/dLPDictionaries:DLPDictionaries", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DLPDictionaries resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DLPDictionaries Get(string name, Input<string> id, DLPDictionariesState? state = null, CustomResourceOptions? options = null)
        {
            return new DLPDictionaries(name, id, state, options);
        }
    }

    public sealed class DLPDictionariesArgs : global::Pulumi.ResourceArgs
    {
        [Input("binNumbers")]
        private InputList<int>? _binNumbers;

        /// <summary>
        /// The list of Bank Identification Number (BIN) values that are included or excluded from the Credit Cards dictionary. BIN
        /// values can be specified only for Diners Club, Mastercard, RuPay, and Visa cards. Up to 512 BIN values can be configured
        /// in a dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
        /// </summary>
        public InputList<int> BinNumbers
        {
            get => _binNumbers ?? (_binNumbers = new InputList<int>());
            set => _binNumbers = value;
        }

        /// <summary>
        /// The DLP confidence threshold for predefined dictionaries
        /// </summary>
        [Input("confidenceLevelForPredefinedDict")]
        public Input<string>? ConfidenceLevelForPredefinedDict { get; set; }

        /// <summary>
        /// The DLP confidence threshold
        /// </summary>
        [Input("confidenceThreshold")]
        public Input<string>? ConfidenceThreshold { get; set; }

        /// <summary>
        /// The DLP dictionary proximity length.
        /// </summary>
        [Input("custom")]
        public Input<bool>? Custom { get; set; }

        [Input("customPhraseMatchType")]
        public Input<string>? CustomPhraseMatchType { get; set; }

        /// <summary>
        /// The desciption of the DLP dictionary
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID of the predefined dictionary (original source dictionary) that is used for cloning. This field is applicable only to
        /// cloned dictionaries. Only a limited set of identification-based predefined dictionaries (e.g., Credit Cards, Social
        /// Security Numbers, National Identification Numbers, etc.) can be cloned. Up to 4 clones can be created from a predefined
        /// dictionary.
        /// </summary>
        [Input("dictTemplateId")]
        public Input<int>? DictTemplateId { get; set; }

        /// <summary>
        /// The DLP dictionary type.
        /// </summary>
        [Input("dictionaryType")]
        public Input<string>? DictionaryType { get; set; }

        [Input("exactDataMatchDetails")]
        private InputList<Inputs.DLPDictionariesExactDataMatchDetailArgs>? _exactDataMatchDetails;

        /// <summary>
        /// Exact Data Match (EDM) related information for custom DLP dictionaries.
        /// </summary>
        public InputList<Inputs.DLPDictionariesExactDataMatchDetailArgs> ExactDataMatchDetails
        {
            get => _exactDataMatchDetails ?? (_exactDataMatchDetails = new InputList<Inputs.DLPDictionariesExactDataMatchDetailArgs>());
            set => _exactDataMatchDetails = value;
        }

        [Input("hierarchicalIdentifiers")]
        private InputList<string>? _hierarchicalIdentifiers;

        /// <summary>
        /// List of hierarchical identifiers for the DLP dictionary.
        /// </summary>
        public InputList<string> HierarchicalIdentifiers
        {
            get => _hierarchicalIdentifiers ?? (_hierarchicalIdentifiers = new InputList<string>());
            set => _hierarchicalIdentifiers = value;
        }

        [Input("idmProfileMatchAccuracies")]
        private InputList<Inputs.DLPDictionariesIdmProfileMatchAccuracyArgs>? _idmProfileMatchAccuracies;

        /// <summary>
        /// List of Indexed Document Match (IDM) profiles and their corresponding match accuracy for custom DLP dictionaries.
        /// </summary>
        public InputList<Inputs.DLPDictionariesIdmProfileMatchAccuracyArgs> IdmProfileMatchAccuracies
        {
            get => _idmProfileMatchAccuracies ?? (_idmProfileMatchAccuracies = new InputList<Inputs.DLPDictionariesIdmProfileMatchAccuracyArgs>());
            set => _idmProfileMatchAccuracies = value;
        }

        /// <summary>
        /// Indicates whether to exclude documents that are a 100% match to already-indexed documents from triggering an Indexed
        /// Document Match (IDM) Dictionary.
        /// </summary>
        [Input("ignoreExactMatchIdmDict")]
        public Input<bool>? IgnoreExactMatchIdmDict { get; set; }

        /// <summary>
        /// A true value denotes that the specified Bank Identification Number (BIN) values are included in the Credit Cards
        /// dictionary. A false value denotes that the specified BIN values are excluded from the Credit Cards dictionary.Note: This
        /// field is applicable only to the predefined Credit Cards dictionary and its clones.
        /// </summary>
        [Input("includeBinNumbers")]
        public Input<bool>? IncludeBinNumbers { get; set; }

        /// <summary>
        /// The DLP dictionary's name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("patterns")]
        private InputList<Inputs.DLPDictionariesPatternArgs>? _patterns;

        /// <summary>
        /// List containing the patterns used within a custom DLP dictionary. This attribute is not applicable to predefined DLP
        /// dictionaries
        /// </summary>
        public InputList<Inputs.DLPDictionariesPatternArgs> Patterns
        {
            get => _patterns ?? (_patterns = new InputList<Inputs.DLPDictionariesPatternArgs>());
            set => _patterns = value;
        }

        [Input("phrases")]
        private InputList<Inputs.DLPDictionariesPhraseArgs>? _phrases;
        public InputList<Inputs.DLPDictionariesPhraseArgs> Phrases
        {
            get => _phrases ?? (_phrases = new InputList<Inputs.DLPDictionariesPhraseArgs>());
            set => _phrases = value;
        }

        /// <summary>
        /// The DLP dictionary proximity length.
        /// </summary>
        [Input("proximity")]
        public Input<int>? Proximity { get; set; }

        public DLPDictionariesArgs()
        {
        }
        public static new DLPDictionariesArgs Empty => new DLPDictionariesArgs();
    }

    public sealed class DLPDictionariesState : global::Pulumi.ResourceArgs
    {
        [Input("binNumbers")]
        private InputList<int>? _binNumbers;

        /// <summary>
        /// The list of Bank Identification Number (BIN) values that are included or excluded from the Credit Cards dictionary. BIN
        /// values can be specified only for Diners Club, Mastercard, RuPay, and Visa cards. Up to 512 BIN values can be configured
        /// in a dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
        /// </summary>
        public InputList<int> BinNumbers
        {
            get => _binNumbers ?? (_binNumbers = new InputList<int>());
            set => _binNumbers = value;
        }

        /// <summary>
        /// The DLP confidence threshold for predefined dictionaries
        /// </summary>
        [Input("confidenceLevelForPredefinedDict")]
        public Input<string>? ConfidenceLevelForPredefinedDict { get; set; }

        /// <summary>
        /// The DLP confidence threshold
        /// </summary>
        [Input("confidenceThreshold")]
        public Input<string>? ConfidenceThreshold { get; set; }

        /// <summary>
        /// The DLP dictionary proximity length.
        /// </summary>
        [Input("custom")]
        public Input<bool>? Custom { get; set; }

        [Input("customPhraseMatchType")]
        public Input<string>? CustomPhraseMatchType { get; set; }

        /// <summary>
        /// The desciption of the DLP dictionary
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID of the predefined dictionary (original source dictionary) that is used for cloning. This field is applicable only to
        /// cloned dictionaries. Only a limited set of identification-based predefined dictionaries (e.g., Credit Cards, Social
        /// Security Numbers, National Identification Numbers, etc.) can be cloned. Up to 4 clones can be created from a predefined
        /// dictionary.
        /// </summary>
        [Input("dictTemplateId")]
        public Input<int>? DictTemplateId { get; set; }

        [Input("dictionaryId")]
        public Input<int>? DictionaryId { get; set; }

        /// <summary>
        /// The DLP dictionary type.
        /// </summary>
        [Input("dictionaryType")]
        public Input<string>? DictionaryType { get; set; }

        [Input("exactDataMatchDetails")]
        private InputList<Inputs.DLPDictionariesExactDataMatchDetailGetArgs>? _exactDataMatchDetails;

        /// <summary>
        /// Exact Data Match (EDM) related information for custom DLP dictionaries.
        /// </summary>
        public InputList<Inputs.DLPDictionariesExactDataMatchDetailGetArgs> ExactDataMatchDetails
        {
            get => _exactDataMatchDetails ?? (_exactDataMatchDetails = new InputList<Inputs.DLPDictionariesExactDataMatchDetailGetArgs>());
            set => _exactDataMatchDetails = value;
        }

        [Input("hierarchicalIdentifiers")]
        private InputList<string>? _hierarchicalIdentifiers;

        /// <summary>
        /// List of hierarchical identifiers for the DLP dictionary.
        /// </summary>
        public InputList<string> HierarchicalIdentifiers
        {
            get => _hierarchicalIdentifiers ?? (_hierarchicalIdentifiers = new InputList<string>());
            set => _hierarchicalIdentifiers = value;
        }

        [Input("idmProfileMatchAccuracies")]
        private InputList<Inputs.DLPDictionariesIdmProfileMatchAccuracyGetArgs>? _idmProfileMatchAccuracies;

        /// <summary>
        /// List of Indexed Document Match (IDM) profiles and their corresponding match accuracy for custom DLP dictionaries.
        /// </summary>
        public InputList<Inputs.DLPDictionariesIdmProfileMatchAccuracyGetArgs> IdmProfileMatchAccuracies
        {
            get => _idmProfileMatchAccuracies ?? (_idmProfileMatchAccuracies = new InputList<Inputs.DLPDictionariesIdmProfileMatchAccuracyGetArgs>());
            set => _idmProfileMatchAccuracies = value;
        }

        /// <summary>
        /// Indicates whether to exclude documents that are a 100% match to already-indexed documents from triggering an Indexed
        /// Document Match (IDM) Dictionary.
        /// </summary>
        [Input("ignoreExactMatchIdmDict")]
        public Input<bool>? IgnoreExactMatchIdmDict { get; set; }

        /// <summary>
        /// A true value denotes that the specified Bank Identification Number (BIN) values are included in the Credit Cards
        /// dictionary. A false value denotes that the specified BIN values are excluded from the Credit Cards dictionary.Note: This
        /// field is applicable only to the predefined Credit Cards dictionary and its clones.
        /// </summary>
        [Input("includeBinNumbers")]
        public Input<bool>? IncludeBinNumbers { get; set; }

        /// <summary>
        /// The DLP dictionary's name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("patterns")]
        private InputList<Inputs.DLPDictionariesPatternGetArgs>? _patterns;

        /// <summary>
        /// List containing the patterns used within a custom DLP dictionary. This attribute is not applicable to predefined DLP
        /// dictionaries
        /// </summary>
        public InputList<Inputs.DLPDictionariesPatternGetArgs> Patterns
        {
            get => _patterns ?? (_patterns = new InputList<Inputs.DLPDictionariesPatternGetArgs>());
            set => _patterns = value;
        }

        [Input("phrases")]
        private InputList<Inputs.DLPDictionariesPhraseGetArgs>? _phrases;
        public InputList<Inputs.DLPDictionariesPhraseGetArgs> Phrases
        {
            get => _phrases ?? (_phrases = new InputList<Inputs.DLPDictionariesPhraseGetArgs>());
            set => _phrases = value;
        }

        /// <summary>
        /// The DLP dictionary proximity length.
        /// </summary>
        [Input("proximity")]
        public Input<int>? Proximity { get; set; }

        public DLPDictionariesState()
        {
        }
        public static new DLPDictionariesState Empty => new DLPDictionariesState();
    }
}
