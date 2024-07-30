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
    public static class GetDLPDictionaries
    {
        /// <summary>
        /// Use the **zia_dlp_dictionaries** data source to get information about a DLP dictionary option available in the Zscaler Internet Access.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zia.GetDLPDictionaries.Invoke(new()
        ///     {
        ///         Name = "SALESFORCE_REPORT_LEAKAGE",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDLPDictionariesResult> InvokeAsync(GetDLPDictionariesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDLPDictionariesResult>("zia:index/getDLPDictionaries:getDLPDictionaries", args ?? new GetDLPDictionariesArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zia_dlp_dictionaries** data source to get information about a DLP dictionary option available in the Zscaler Internet Access.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zia.GetDLPDictionaries.Invoke(new()
        ///     {
        ///         Name = "SALESFORCE_REPORT_LEAKAGE",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDLPDictionariesResult> Invoke(GetDLPDictionariesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDLPDictionariesResult>("zia:index/getDLPDictionaries:getDLPDictionaries", args ?? new GetDLPDictionariesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDLPDictionariesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for the DLP dictionary
        /// </summary>
        [Input("id")]
        public int? Id { get; set; }

        /// <summary>
        /// DLP dictionary name
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetDLPDictionariesArgs()
        {
        }
        public static new GetDLPDictionariesArgs Empty => new GetDLPDictionariesArgs();
    }

    public sealed class GetDLPDictionariesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for the DLP dictionary
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// DLP dictionary name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetDLPDictionariesInvokeArgs()
        {
        }
        public static new GetDLPDictionariesInvokeArgs Empty => new GetDLPDictionariesInvokeArgs();
    }


    [OutputType]
    public sealed class GetDLPDictionariesResult
    {
        /// <summary>
        /// (Boolean) The list of Bank Identification Number (BIN) values that are included or excluded from the Credit Cards dictionary. BIN values can be specified only for Diners Club, Mastercard, RuPay, and Visa cards. Up to 512 BIN values can be configured in a dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
        /// </summary>
        public readonly ImmutableArray<int> BinNumbers;
        /// <summary>
        /// (String) he DLP confidence threshold. [`CONFIDENCE_LEVEL_LOW`, `CONFIDENCE_LEVEL_MEDIUM` `CONFIDENCE_LEVEL_HIGH` ]
        /// </summary>
        public readonly string ConfidenceThreshold;
        /// <summary>
        /// (Boolean) This value is set to true for custom DLP dictionaries.
        /// </summary>
        public readonly bool Custom;
        /// <summary>
        /// (String) The DLP custom phrase match type. [ `MATCH_ALL_CUSTOM_PHRASE_PATTERN_DICTIONARY`, `MATCH_ANY_CUSTOM_PHRASE_PATTERN_DICTIONARY` ]
        /// </summary>
        public readonly string CustomPhraseMatchType;
        public readonly string Description;
        /// <summary>
        /// (Number) ID of the predefined dictionary (original source dictionary) that is used for cloning. This field is applicable only to cloned dictionaries. Only a limited set of identification-based predefined dictionaries (e.g., Credit Cards, Social Security Numbers, National Identification Numbers, etc.) can be cloned. Up to 4 clones can be created from a predefined dictionary.
        /// </summary>
        public readonly int DictTemplateId;
        /// <summary>
        /// (String) The DLP dictionary type. The cloud service API only supports custom DLP dictionaries that are using the `PATTERNS_AND_PHRASES` type.
        /// </summary>
        public readonly string DictionaryType;
        public readonly ImmutableArray<Outputs.GetDLPDictionariesExactDataMatchDetailResult> ExactDataMatchDetails;
        public readonly int Id;
        public readonly ImmutableArray<Outputs.GetDLPDictionariesIdmProfileMatchAccuracyResult> IdmProfileMatchAccuracies;
        /// <summary>
        /// (Boolean) Indicates whether to exclude documents that are a 100% match to already-indexed documents from triggering an Indexed Document Match (IDM) Dictionary.
        /// </summary>
        public readonly bool IgnoreExactMatchIdmDict;
        /// <summary>
        /// (Boolean) A true value denotes that the specified Bank Identification Number (BIN) values are included in the Credit Cards dictionary. A false value denotes that the specified BIN values are excluded from the Credit Cards dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
        /// </summary>
        public readonly bool IncludeBinNumbers;
        public readonly string Name;
        /// <summary>
        /// (Boolean) Indicates whether the name is localized or not. This is always set to True for predefined DLP dictionaries.
        /// </summary>
        public readonly bool NameL10nTag;
        public readonly ImmutableArray<Outputs.GetDLPDictionariesPatternResult> Patterns;
        public readonly ImmutableArray<Outputs.GetDLPDictionariesPhraseResult> Phrases;
        /// <summary>
        /// (Boolean) This field is set to true if the dictionary is cloned from a predefined dictionary. Otherwise, it is set to false.
        /// </summary>
        public readonly bool PredefinedClone;
        public readonly int Proximity;
        public readonly bool ProximityLengthEnabled;
        public readonly string ThresholdType;

        [OutputConstructor]
        private GetDLPDictionariesResult(
            ImmutableArray<int> binNumbers,

            string confidenceThreshold,

            bool custom,

            string customPhraseMatchType,

            string description,

            int dictTemplateId,

            string dictionaryType,

            ImmutableArray<Outputs.GetDLPDictionariesExactDataMatchDetailResult> exactDataMatchDetails,

            int id,

            ImmutableArray<Outputs.GetDLPDictionariesIdmProfileMatchAccuracyResult> idmProfileMatchAccuracies,

            bool ignoreExactMatchIdmDict,

            bool includeBinNumbers,

            string name,

            bool nameL10nTag,

            ImmutableArray<Outputs.GetDLPDictionariesPatternResult> patterns,

            ImmutableArray<Outputs.GetDLPDictionariesPhraseResult> phrases,

            bool predefinedClone,

            int proximity,

            bool proximityLengthEnabled,

            string thresholdType)
        {
            BinNumbers = binNumbers;
            ConfidenceThreshold = confidenceThreshold;
            Custom = custom;
            CustomPhraseMatchType = customPhraseMatchType;
            Description = description;
            DictTemplateId = dictTemplateId;
            DictionaryType = dictionaryType;
            ExactDataMatchDetails = exactDataMatchDetails;
            Id = id;
            IdmProfileMatchAccuracies = idmProfileMatchAccuracies;
            IgnoreExactMatchIdmDict = ignoreExactMatchIdmDict;
            IncludeBinNumbers = includeBinNumbers;
            Name = name;
            NameL10nTag = nameL10nTag;
            Patterns = patterns;
            Phrases = phrases;
            PredefinedClone = predefinedClone;
            Proximity = proximity;
            ProximityLengthEnabled = proximityLengthEnabled;
            ThresholdType = thresholdType;
        }
    }
}
