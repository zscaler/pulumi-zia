# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetDLPDictionariesResult',
    'AwaitableGetDLPDictionariesResult',
    'get_dlp_dictionaries',
    'get_dlp_dictionaries_output',
]

@pulumi.output_type
class GetDLPDictionariesResult:
    """
    A collection of values returned by getDLPDictionaries.
    """
    def __init__(__self__, bin_numbers=None, confidence_threshold=None, custom=None, custom_phrase_match_type=None, description=None, dict_template_id=None, dictionary_type=None, exact_data_match_details=None, id=None, idm_profile_match_accuracies=None, ignore_exact_match_idm_dict=None, include_bin_numbers=None, name=None, name_l10n_tag=None, patterns=None, phrases=None, predefined_clone=None, proximity=None, proximity_length_enabled=None, threshold_type=None):
        if bin_numbers and not isinstance(bin_numbers, list):
            raise TypeError("Expected argument 'bin_numbers' to be a list")
        pulumi.set(__self__, "bin_numbers", bin_numbers)
        if confidence_threshold and not isinstance(confidence_threshold, str):
            raise TypeError("Expected argument 'confidence_threshold' to be a str")
        pulumi.set(__self__, "confidence_threshold", confidence_threshold)
        if custom and not isinstance(custom, bool):
            raise TypeError("Expected argument 'custom' to be a bool")
        pulumi.set(__self__, "custom", custom)
        if custom_phrase_match_type and not isinstance(custom_phrase_match_type, str):
            raise TypeError("Expected argument 'custom_phrase_match_type' to be a str")
        pulumi.set(__self__, "custom_phrase_match_type", custom_phrase_match_type)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dict_template_id and not isinstance(dict_template_id, int):
            raise TypeError("Expected argument 'dict_template_id' to be a int")
        pulumi.set(__self__, "dict_template_id", dict_template_id)
        if dictionary_type and not isinstance(dictionary_type, str):
            raise TypeError("Expected argument 'dictionary_type' to be a str")
        pulumi.set(__self__, "dictionary_type", dictionary_type)
        if exact_data_match_details and not isinstance(exact_data_match_details, list):
            raise TypeError("Expected argument 'exact_data_match_details' to be a list")
        pulumi.set(__self__, "exact_data_match_details", exact_data_match_details)
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)
        if idm_profile_match_accuracies and not isinstance(idm_profile_match_accuracies, list):
            raise TypeError("Expected argument 'idm_profile_match_accuracies' to be a list")
        pulumi.set(__self__, "idm_profile_match_accuracies", idm_profile_match_accuracies)
        if ignore_exact_match_idm_dict and not isinstance(ignore_exact_match_idm_dict, bool):
            raise TypeError("Expected argument 'ignore_exact_match_idm_dict' to be a bool")
        pulumi.set(__self__, "ignore_exact_match_idm_dict", ignore_exact_match_idm_dict)
        if include_bin_numbers and not isinstance(include_bin_numbers, bool):
            raise TypeError("Expected argument 'include_bin_numbers' to be a bool")
        pulumi.set(__self__, "include_bin_numbers", include_bin_numbers)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_l10n_tag and not isinstance(name_l10n_tag, bool):
            raise TypeError("Expected argument 'name_l10n_tag' to be a bool")
        pulumi.set(__self__, "name_l10n_tag", name_l10n_tag)
        if patterns and not isinstance(patterns, list):
            raise TypeError("Expected argument 'patterns' to be a list")
        pulumi.set(__self__, "patterns", patterns)
        if phrases and not isinstance(phrases, list):
            raise TypeError("Expected argument 'phrases' to be a list")
        pulumi.set(__self__, "phrases", phrases)
        if predefined_clone and not isinstance(predefined_clone, bool):
            raise TypeError("Expected argument 'predefined_clone' to be a bool")
        pulumi.set(__self__, "predefined_clone", predefined_clone)
        if proximity and not isinstance(proximity, int):
            raise TypeError("Expected argument 'proximity' to be a int")
        pulumi.set(__self__, "proximity", proximity)
        if proximity_length_enabled and not isinstance(proximity_length_enabled, bool):
            raise TypeError("Expected argument 'proximity_length_enabled' to be a bool")
        pulumi.set(__self__, "proximity_length_enabled", proximity_length_enabled)
        if threshold_type and not isinstance(threshold_type, str):
            raise TypeError("Expected argument 'threshold_type' to be a str")
        pulumi.set(__self__, "threshold_type", threshold_type)

    @property
    @pulumi.getter(name="binNumbers")
    def bin_numbers(self) -> Sequence[int]:
        """
        (Boolean) The list of Bank Identification Number (BIN) values that are included or excluded from the Credit Cards dictionary. BIN values can be specified only for Diners Club, Mastercard, RuPay, and Visa cards. Up to 512 BIN values can be configured in a dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
        """
        return pulumi.get(self, "bin_numbers")

    @property
    @pulumi.getter(name="confidenceThreshold")
    def confidence_threshold(self) -> str:
        """
        (String) he DLP confidence threshold. [`CONFIDENCE_LEVEL_LOW`, `CONFIDENCE_LEVEL_MEDIUM` `CONFIDENCE_LEVEL_HIGH` ]
        """
        return pulumi.get(self, "confidence_threshold")

    @property
    @pulumi.getter
    def custom(self) -> bool:
        """
        (Boolean) This value is set to true for custom DLP dictionaries.
        """
        return pulumi.get(self, "custom")

    @property
    @pulumi.getter(name="customPhraseMatchType")
    def custom_phrase_match_type(self) -> str:
        """
        (String) The DLP custom phrase match type. [ `MATCH_ALL_CUSTOM_PHRASE_PATTERN_DICTIONARY`, `MATCH_ANY_CUSTOM_PHRASE_PATTERN_DICTIONARY` ]
        """
        return pulumi.get(self, "custom_phrase_match_type")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dictTemplateId")
    def dict_template_id(self) -> int:
        """
        (Number) ID of the predefined dictionary (original source dictionary) that is used for cloning. This field is applicable only to cloned dictionaries. Only a limited set of identification-based predefined dictionaries (e.g., Credit Cards, Social Security Numbers, National Identification Numbers, etc.) can be cloned. Up to 4 clones can be created from a predefined dictionary.
        """
        return pulumi.get(self, "dict_template_id")

    @property
    @pulumi.getter(name="dictionaryType")
    def dictionary_type(self) -> str:
        """
        (String) The DLP dictionary type. The cloud service API only supports custom DLP dictionaries that are using the `PATTERNS_AND_PHRASES` type.
        """
        return pulumi.get(self, "dictionary_type")

    @property
    @pulumi.getter(name="exactDataMatchDetails")
    def exact_data_match_details(self) -> Sequence['outputs.GetDLPDictionariesExactDataMatchDetailResult']:
        return pulumi.get(self, "exact_data_match_details")

    @property
    @pulumi.getter
    def id(self) -> int:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="idmProfileMatchAccuracies")
    def idm_profile_match_accuracies(self) -> Sequence['outputs.GetDLPDictionariesIdmProfileMatchAccuracyResult']:
        return pulumi.get(self, "idm_profile_match_accuracies")

    @property
    @pulumi.getter(name="ignoreExactMatchIdmDict")
    def ignore_exact_match_idm_dict(self) -> bool:
        """
        (Boolean) Indicates whether to exclude documents that are a 100% match to already-indexed documents from triggering an Indexed Document Match (IDM) Dictionary.
        """
        return pulumi.get(self, "ignore_exact_match_idm_dict")

    @property
    @pulumi.getter(name="includeBinNumbers")
    def include_bin_numbers(self) -> bool:
        """
        (Boolean) A true value denotes that the specified Bank Identification Number (BIN) values are included in the Credit Cards dictionary. A false value denotes that the specified BIN values are excluded from the Credit Cards dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
        """
        return pulumi.get(self, "include_bin_numbers")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameL10nTag")
    def name_l10n_tag(self) -> bool:
        """
        (Boolean) Indicates whether the name is localized or not. This is always set to True for predefined DLP dictionaries.
        """
        return pulumi.get(self, "name_l10n_tag")

    @property
    @pulumi.getter
    def patterns(self) -> Sequence['outputs.GetDLPDictionariesPatternResult']:
        return pulumi.get(self, "patterns")

    @property
    @pulumi.getter
    def phrases(self) -> Sequence['outputs.GetDLPDictionariesPhraseResult']:
        return pulumi.get(self, "phrases")

    @property
    @pulumi.getter(name="predefinedClone")
    def predefined_clone(self) -> bool:
        """
        (Boolean) This field is set to true if the dictionary is cloned from a predefined dictionary. Otherwise, it is set to false.
        """
        return pulumi.get(self, "predefined_clone")

    @property
    @pulumi.getter
    def proximity(self) -> int:
        return pulumi.get(self, "proximity")

    @property
    @pulumi.getter(name="proximityLengthEnabled")
    def proximity_length_enabled(self) -> bool:
        return pulumi.get(self, "proximity_length_enabled")

    @property
    @pulumi.getter(name="thresholdType")
    def threshold_type(self) -> str:
        return pulumi.get(self, "threshold_type")


class AwaitableGetDLPDictionariesResult(GetDLPDictionariesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDLPDictionariesResult(
            bin_numbers=self.bin_numbers,
            confidence_threshold=self.confidence_threshold,
            custom=self.custom,
            custom_phrase_match_type=self.custom_phrase_match_type,
            description=self.description,
            dict_template_id=self.dict_template_id,
            dictionary_type=self.dictionary_type,
            exact_data_match_details=self.exact_data_match_details,
            id=self.id,
            idm_profile_match_accuracies=self.idm_profile_match_accuracies,
            ignore_exact_match_idm_dict=self.ignore_exact_match_idm_dict,
            include_bin_numbers=self.include_bin_numbers,
            name=self.name,
            name_l10n_tag=self.name_l10n_tag,
            patterns=self.patterns,
            phrases=self.phrases,
            predefined_clone=self.predefined_clone,
            proximity=self.proximity,
            proximity_length_enabled=self.proximity_length_enabled,
            threshold_type=self.threshold_type)


def get_dlp_dictionaries(id: Optional[int] = None,
                         name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDLPDictionariesResult:
    """
    Use the **zia_dlp_dictionaries** data source to get information about a DLP dictionary option available in the Zscaler Internet Access.

    ```python
    import pulumi
    import pulumi_zia as zia

    example = zia.get_dlp_dictionaries(name="SALESFORCE_REPORT_LEAKAGE")
    ```


    :param int id: Unique identifier for the DLP dictionary
    :param str name: DLP dictionary name
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zia:index/getDLPDictionaries:getDLPDictionaries', __args__, opts=opts, typ=GetDLPDictionariesResult).value

    return AwaitableGetDLPDictionariesResult(
        bin_numbers=pulumi.get(__ret__, 'bin_numbers'),
        confidence_threshold=pulumi.get(__ret__, 'confidence_threshold'),
        custom=pulumi.get(__ret__, 'custom'),
        custom_phrase_match_type=pulumi.get(__ret__, 'custom_phrase_match_type'),
        description=pulumi.get(__ret__, 'description'),
        dict_template_id=pulumi.get(__ret__, 'dict_template_id'),
        dictionary_type=pulumi.get(__ret__, 'dictionary_type'),
        exact_data_match_details=pulumi.get(__ret__, 'exact_data_match_details'),
        id=pulumi.get(__ret__, 'id'),
        idm_profile_match_accuracies=pulumi.get(__ret__, 'idm_profile_match_accuracies'),
        ignore_exact_match_idm_dict=pulumi.get(__ret__, 'ignore_exact_match_idm_dict'),
        include_bin_numbers=pulumi.get(__ret__, 'include_bin_numbers'),
        name=pulumi.get(__ret__, 'name'),
        name_l10n_tag=pulumi.get(__ret__, 'name_l10n_tag'),
        patterns=pulumi.get(__ret__, 'patterns'),
        phrases=pulumi.get(__ret__, 'phrases'),
        predefined_clone=pulumi.get(__ret__, 'predefined_clone'),
        proximity=pulumi.get(__ret__, 'proximity'),
        proximity_length_enabled=pulumi.get(__ret__, 'proximity_length_enabled'),
        threshold_type=pulumi.get(__ret__, 'threshold_type'))


@_utilities.lift_output_func(get_dlp_dictionaries)
def get_dlp_dictionaries_output(id: Optional[pulumi.Input[Optional[int]]] = None,
                                name: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDLPDictionariesResult]:
    """
    Use the **zia_dlp_dictionaries** data source to get information about a DLP dictionary option available in the Zscaler Internet Access.

    ```python
    import pulumi
    import pulumi_zia as zia

    example = zia.get_dlp_dictionaries(name="SALESFORCE_REPORT_LEAKAGE")
    ```


    :param int id: Unique identifier for the DLP dictionary
    :param str name: DLP dictionary name
    """
    ...
