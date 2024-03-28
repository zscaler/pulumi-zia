// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zia

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zia/sdk/go/zia/internal"
)

// The **zia_dlp_dictionaries** resource allows the creation and management of ZIA DLP dictionaries in the Zscaler Internet Access cloud or via the API.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/zscaler/pulumi-zia/sdk/go/zia"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := zia.NewDLPDictionaries(ctx, "example", &zia.DLPDictionariesArgs{
//				CustomPhraseMatchType: pulumi.String("MATCH_ALL_CUSTOM_PHRASE_PATTERN_DICTIONARY"),
//				Description:           pulumi.String("Your Description"),
//				DictionaryType:        pulumi.String("PATTERNS_AND_PHRASES"),
//				Patterns: zia.DLPDictionariesPatternArray{
//					&zia.DLPDictionariesPatternArgs{
//						Action:  pulumi.String("PATTERN_COUNT_TYPE_UNIQUE"),
//						Pattern: pulumi.String("YourPattern"),
//					},
//				},
//				Phrases: zia.DLPDictionariesPhraseArray{
//					&zia.DLPDictionariesPhraseArgs{
//						Action: pulumi.String("PHRASE_COUNT_TYPE_ALL"),
//						Phrase: pulumi.String("YourPhrase"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.
//
// # Visit
//
// **zia_dlp_dictionaries** can be imported by using `<DICTIONARY ID>` or `<DICTIONARY_NAME>` as the import ID.
//
// For example:
//
// ```sh
// $ pulumi import zia:index/dLPDictionaries:DLPDictionaries example <dictionary_id>
// ```
//
// or
//
// ```sh
// $ pulumi import zia:index/dLPDictionaries:DLPDictionaries example <dictionary_name>
// ```
type DLPDictionaries struct {
	pulumi.CustomResourceState

	// The list of Bank Identification Number (BIN) values that are included or excluded from the Credit Cards dictionary. BIN values can be specified only for Diners Club, Mastercard, RuPay, and Visa cards. Up to 512 BIN values can be configured in a dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
	BinNumbers pulumi.IntArrayOutput `pulumi:"binNumbers"`
	// The DLP confidence threshold. The following values are supported:
	ConfidenceThreshold pulumi.StringPtrOutput `pulumi:"confidenceThreshold"`
	// The DLP custom phrase match type. Supported values are:
	CustomPhraseMatchType pulumi.StringOutput `pulumi:"customPhraseMatchType"`
	// The desciption of the DLP dictionary
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// ID of the predefined dictionary (original source dictionary) that is used for cloning. This field is applicable only to cloned dictionaries. Only a limited set of identification-based predefined dictionaries (e.g., Credit Cards, Social Security Numbers, National Identification Numbers, etc.) can be cloned. Up to 4 clones can be created from a predefined dictionary.
	DictTemplateId pulumi.IntPtrOutput `pulumi:"dictTemplateId"`
	DictionaryId   pulumi.IntOutput    `pulumi:"dictionaryId"`
	// The DLP dictionary type. The following values are supported:
	DictionaryType pulumi.StringPtrOutput `pulumi:"dictionaryType"`
	// Exact Data Match (EDM) related information for custom DLP dictionaries.
	ExactDataMatchDetails DLPDictionariesExactDataMatchDetailArrayOutput `pulumi:"exactDataMatchDetails"`
	// List of Indexed Document Match (IDM) profiles and their corresponding match accuracy for custom DLP dictionaries.
	IdmProfileMatchAccuracies DLPDictionariesIdmProfileMatchAccuracyArrayOutput `pulumi:"idmProfileMatchAccuracies"`
	// Indicates whether to exclude documents that are a 100%!m(MISSING)atch to already-indexed documents from triggering an Indexed Document Match (IDM) Dictionary.
	IgnoreExactMatchIdmDict pulumi.BoolPtrOutput `pulumi:"ignoreExactMatchIdmDict"`
	// A true value denotes that the specified Bank Identification Number (BIN) values are included in the Credit Cards dictionary. A false value denotes that the specified BIN values are excluded from the Credit Cards dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
	IncludeBinNumbers pulumi.BoolOutput `pulumi:"includeBinNumbers"`
	// The DLP dictionary's name
	Name pulumi.StringOutput `pulumi:"name"`
	// List containing the patterns used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries. Required when `dictionaryType` is `PATTERNS_AND_PHRASES`
	Patterns DLPDictionariesPatternArrayOutput `pulumi:"patterns"`
	// List containing the phrases used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries. Required when `dictionaryType` is `PATTERNS_AND_PHRASES`
	Phrases DLPDictionariesPhraseArrayOutput `pulumi:"phrases"`
	// The DLP dictionary proximity length.
	Proximity pulumi.IntPtrOutput `pulumi:"proximity"`
}

// NewDLPDictionaries registers a new resource with the given unique name, arguments, and options.
func NewDLPDictionaries(ctx *pulumi.Context,
	name string, args *DLPDictionariesArgs, opts ...pulumi.ResourceOption) (*DLPDictionaries, error) {
	if args == nil {
		args = &DLPDictionariesArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DLPDictionaries
	err := ctx.RegisterResource("zia:index/dLPDictionaries:DLPDictionaries", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDLPDictionaries gets an existing DLPDictionaries resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDLPDictionaries(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DLPDictionariesState, opts ...pulumi.ResourceOption) (*DLPDictionaries, error) {
	var resource DLPDictionaries
	err := ctx.ReadResource("zia:index/dLPDictionaries:DLPDictionaries", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DLPDictionaries resources.
type dlpdictionariesState struct {
	// The list of Bank Identification Number (BIN) values that are included or excluded from the Credit Cards dictionary. BIN values can be specified only for Diners Club, Mastercard, RuPay, and Visa cards. Up to 512 BIN values can be configured in a dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
	BinNumbers []int `pulumi:"binNumbers"`
	// The DLP confidence threshold. The following values are supported:
	ConfidenceThreshold *string `pulumi:"confidenceThreshold"`
	// The DLP custom phrase match type. Supported values are:
	CustomPhraseMatchType *string `pulumi:"customPhraseMatchType"`
	// The desciption of the DLP dictionary
	Description *string `pulumi:"description"`
	// ID of the predefined dictionary (original source dictionary) that is used for cloning. This field is applicable only to cloned dictionaries. Only a limited set of identification-based predefined dictionaries (e.g., Credit Cards, Social Security Numbers, National Identification Numbers, etc.) can be cloned. Up to 4 clones can be created from a predefined dictionary.
	DictTemplateId *int `pulumi:"dictTemplateId"`
	DictionaryId   *int `pulumi:"dictionaryId"`
	// The DLP dictionary type. The following values are supported:
	DictionaryType *string `pulumi:"dictionaryType"`
	// Exact Data Match (EDM) related information for custom DLP dictionaries.
	ExactDataMatchDetails []DLPDictionariesExactDataMatchDetail `pulumi:"exactDataMatchDetails"`
	// List of Indexed Document Match (IDM) profiles and their corresponding match accuracy for custom DLP dictionaries.
	IdmProfileMatchAccuracies []DLPDictionariesIdmProfileMatchAccuracy `pulumi:"idmProfileMatchAccuracies"`
	// Indicates whether to exclude documents that are a 100%!m(MISSING)atch to already-indexed documents from triggering an Indexed Document Match (IDM) Dictionary.
	IgnoreExactMatchIdmDict *bool `pulumi:"ignoreExactMatchIdmDict"`
	// A true value denotes that the specified Bank Identification Number (BIN) values are included in the Credit Cards dictionary. A false value denotes that the specified BIN values are excluded from the Credit Cards dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
	IncludeBinNumbers *bool `pulumi:"includeBinNumbers"`
	// The DLP dictionary's name
	Name *string `pulumi:"name"`
	// List containing the patterns used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries. Required when `dictionaryType` is `PATTERNS_AND_PHRASES`
	Patterns []DLPDictionariesPattern `pulumi:"patterns"`
	// List containing the phrases used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries. Required when `dictionaryType` is `PATTERNS_AND_PHRASES`
	Phrases []DLPDictionariesPhrase `pulumi:"phrases"`
	// The DLP dictionary proximity length.
	Proximity *int `pulumi:"proximity"`
}

type DLPDictionariesState struct {
	// The list of Bank Identification Number (BIN) values that are included or excluded from the Credit Cards dictionary. BIN values can be specified only for Diners Club, Mastercard, RuPay, and Visa cards. Up to 512 BIN values can be configured in a dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
	BinNumbers pulumi.IntArrayInput
	// The DLP confidence threshold. The following values are supported:
	ConfidenceThreshold pulumi.StringPtrInput
	// The DLP custom phrase match type. Supported values are:
	CustomPhraseMatchType pulumi.StringPtrInput
	// The desciption of the DLP dictionary
	Description pulumi.StringPtrInput
	// ID of the predefined dictionary (original source dictionary) that is used for cloning. This field is applicable only to cloned dictionaries. Only a limited set of identification-based predefined dictionaries (e.g., Credit Cards, Social Security Numbers, National Identification Numbers, etc.) can be cloned. Up to 4 clones can be created from a predefined dictionary.
	DictTemplateId pulumi.IntPtrInput
	DictionaryId   pulumi.IntPtrInput
	// The DLP dictionary type. The following values are supported:
	DictionaryType pulumi.StringPtrInput
	// Exact Data Match (EDM) related information for custom DLP dictionaries.
	ExactDataMatchDetails DLPDictionariesExactDataMatchDetailArrayInput
	// List of Indexed Document Match (IDM) profiles and their corresponding match accuracy for custom DLP dictionaries.
	IdmProfileMatchAccuracies DLPDictionariesIdmProfileMatchAccuracyArrayInput
	// Indicates whether to exclude documents that are a 100%!m(MISSING)atch to already-indexed documents from triggering an Indexed Document Match (IDM) Dictionary.
	IgnoreExactMatchIdmDict pulumi.BoolPtrInput
	// A true value denotes that the specified Bank Identification Number (BIN) values are included in the Credit Cards dictionary. A false value denotes that the specified BIN values are excluded from the Credit Cards dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
	IncludeBinNumbers pulumi.BoolPtrInput
	// The DLP dictionary's name
	Name pulumi.StringPtrInput
	// List containing the patterns used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries. Required when `dictionaryType` is `PATTERNS_AND_PHRASES`
	Patterns DLPDictionariesPatternArrayInput
	// List containing the phrases used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries. Required when `dictionaryType` is `PATTERNS_AND_PHRASES`
	Phrases DLPDictionariesPhraseArrayInput
	// The DLP dictionary proximity length.
	Proximity pulumi.IntPtrInput
}

func (DLPDictionariesState) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpdictionariesState)(nil)).Elem()
}

type dlpdictionariesArgs struct {
	// The list of Bank Identification Number (BIN) values that are included or excluded from the Credit Cards dictionary. BIN values can be specified only for Diners Club, Mastercard, RuPay, and Visa cards. Up to 512 BIN values can be configured in a dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
	BinNumbers []int `pulumi:"binNumbers"`
	// The DLP confidence threshold. The following values are supported:
	ConfidenceThreshold *string `pulumi:"confidenceThreshold"`
	// The DLP custom phrase match type. Supported values are:
	CustomPhraseMatchType *string `pulumi:"customPhraseMatchType"`
	// The desciption of the DLP dictionary
	Description *string `pulumi:"description"`
	// ID of the predefined dictionary (original source dictionary) that is used for cloning. This field is applicable only to cloned dictionaries. Only a limited set of identification-based predefined dictionaries (e.g., Credit Cards, Social Security Numbers, National Identification Numbers, etc.) can be cloned. Up to 4 clones can be created from a predefined dictionary.
	DictTemplateId *int `pulumi:"dictTemplateId"`
	// The DLP dictionary type. The following values are supported:
	DictionaryType *string `pulumi:"dictionaryType"`
	// Exact Data Match (EDM) related information for custom DLP dictionaries.
	ExactDataMatchDetails []DLPDictionariesExactDataMatchDetail `pulumi:"exactDataMatchDetails"`
	// List of Indexed Document Match (IDM) profiles and their corresponding match accuracy for custom DLP dictionaries.
	IdmProfileMatchAccuracies []DLPDictionariesIdmProfileMatchAccuracy `pulumi:"idmProfileMatchAccuracies"`
	// Indicates whether to exclude documents that are a 100%!m(MISSING)atch to already-indexed documents from triggering an Indexed Document Match (IDM) Dictionary.
	IgnoreExactMatchIdmDict *bool `pulumi:"ignoreExactMatchIdmDict"`
	// A true value denotes that the specified Bank Identification Number (BIN) values are included in the Credit Cards dictionary. A false value denotes that the specified BIN values are excluded from the Credit Cards dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
	IncludeBinNumbers *bool `pulumi:"includeBinNumbers"`
	// The DLP dictionary's name
	Name *string `pulumi:"name"`
	// List containing the patterns used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries. Required when `dictionaryType` is `PATTERNS_AND_PHRASES`
	Patterns []DLPDictionariesPattern `pulumi:"patterns"`
	// List containing the phrases used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries. Required when `dictionaryType` is `PATTERNS_AND_PHRASES`
	Phrases []DLPDictionariesPhrase `pulumi:"phrases"`
	// The DLP dictionary proximity length.
	Proximity *int `pulumi:"proximity"`
}

// The set of arguments for constructing a DLPDictionaries resource.
type DLPDictionariesArgs struct {
	// The list of Bank Identification Number (BIN) values that are included or excluded from the Credit Cards dictionary. BIN values can be specified only for Diners Club, Mastercard, RuPay, and Visa cards. Up to 512 BIN values can be configured in a dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
	BinNumbers pulumi.IntArrayInput
	// The DLP confidence threshold. The following values are supported:
	ConfidenceThreshold pulumi.StringPtrInput
	// The DLP custom phrase match type. Supported values are:
	CustomPhraseMatchType pulumi.StringPtrInput
	// The desciption of the DLP dictionary
	Description pulumi.StringPtrInput
	// ID of the predefined dictionary (original source dictionary) that is used for cloning. This field is applicable only to cloned dictionaries. Only a limited set of identification-based predefined dictionaries (e.g., Credit Cards, Social Security Numbers, National Identification Numbers, etc.) can be cloned. Up to 4 clones can be created from a predefined dictionary.
	DictTemplateId pulumi.IntPtrInput
	// The DLP dictionary type. The following values are supported:
	DictionaryType pulumi.StringPtrInput
	// Exact Data Match (EDM) related information for custom DLP dictionaries.
	ExactDataMatchDetails DLPDictionariesExactDataMatchDetailArrayInput
	// List of Indexed Document Match (IDM) profiles and their corresponding match accuracy for custom DLP dictionaries.
	IdmProfileMatchAccuracies DLPDictionariesIdmProfileMatchAccuracyArrayInput
	// Indicates whether to exclude documents that are a 100%!m(MISSING)atch to already-indexed documents from triggering an Indexed Document Match (IDM) Dictionary.
	IgnoreExactMatchIdmDict pulumi.BoolPtrInput
	// A true value denotes that the specified Bank Identification Number (BIN) values are included in the Credit Cards dictionary. A false value denotes that the specified BIN values are excluded from the Credit Cards dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
	IncludeBinNumbers pulumi.BoolPtrInput
	// The DLP dictionary's name
	Name pulumi.StringPtrInput
	// List containing the patterns used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries. Required when `dictionaryType` is `PATTERNS_AND_PHRASES`
	Patterns DLPDictionariesPatternArrayInput
	// List containing the phrases used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries. Required when `dictionaryType` is `PATTERNS_AND_PHRASES`
	Phrases DLPDictionariesPhraseArrayInput
	// The DLP dictionary proximity length.
	Proximity pulumi.IntPtrInput
}

func (DLPDictionariesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpdictionariesArgs)(nil)).Elem()
}

type DLPDictionariesInput interface {
	pulumi.Input

	ToDLPDictionariesOutput() DLPDictionariesOutput
	ToDLPDictionariesOutputWithContext(ctx context.Context) DLPDictionariesOutput
}

func (*DLPDictionaries) ElementType() reflect.Type {
	return reflect.TypeOf((**DLPDictionaries)(nil)).Elem()
}

func (i *DLPDictionaries) ToDLPDictionariesOutput() DLPDictionariesOutput {
	return i.ToDLPDictionariesOutputWithContext(context.Background())
}

func (i *DLPDictionaries) ToDLPDictionariesOutputWithContext(ctx context.Context) DLPDictionariesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DLPDictionariesOutput)
}

// DLPDictionariesArrayInput is an input type that accepts DLPDictionariesArray and DLPDictionariesArrayOutput values.
// You can construct a concrete instance of `DLPDictionariesArrayInput` via:
//
//	DLPDictionariesArray{ DLPDictionariesArgs{...} }
type DLPDictionariesArrayInput interface {
	pulumi.Input

	ToDLPDictionariesArrayOutput() DLPDictionariesArrayOutput
	ToDLPDictionariesArrayOutputWithContext(context.Context) DLPDictionariesArrayOutput
}

type DLPDictionariesArray []DLPDictionariesInput

func (DLPDictionariesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DLPDictionaries)(nil)).Elem()
}

func (i DLPDictionariesArray) ToDLPDictionariesArrayOutput() DLPDictionariesArrayOutput {
	return i.ToDLPDictionariesArrayOutputWithContext(context.Background())
}

func (i DLPDictionariesArray) ToDLPDictionariesArrayOutputWithContext(ctx context.Context) DLPDictionariesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DLPDictionariesArrayOutput)
}

// DLPDictionariesMapInput is an input type that accepts DLPDictionariesMap and DLPDictionariesMapOutput values.
// You can construct a concrete instance of `DLPDictionariesMapInput` via:
//
//	DLPDictionariesMap{ "key": DLPDictionariesArgs{...} }
type DLPDictionariesMapInput interface {
	pulumi.Input

	ToDLPDictionariesMapOutput() DLPDictionariesMapOutput
	ToDLPDictionariesMapOutputWithContext(context.Context) DLPDictionariesMapOutput
}

type DLPDictionariesMap map[string]DLPDictionariesInput

func (DLPDictionariesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DLPDictionaries)(nil)).Elem()
}

func (i DLPDictionariesMap) ToDLPDictionariesMapOutput() DLPDictionariesMapOutput {
	return i.ToDLPDictionariesMapOutputWithContext(context.Background())
}

func (i DLPDictionariesMap) ToDLPDictionariesMapOutputWithContext(ctx context.Context) DLPDictionariesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DLPDictionariesMapOutput)
}

type DLPDictionariesOutput struct{ *pulumi.OutputState }

func (DLPDictionariesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DLPDictionaries)(nil)).Elem()
}

func (o DLPDictionariesOutput) ToDLPDictionariesOutput() DLPDictionariesOutput {
	return o
}

func (o DLPDictionariesOutput) ToDLPDictionariesOutputWithContext(ctx context.Context) DLPDictionariesOutput {
	return o
}

// The list of Bank Identification Number (BIN) values that are included or excluded from the Credit Cards dictionary. BIN values can be specified only for Diners Club, Mastercard, RuPay, and Visa cards. Up to 512 BIN values can be configured in a dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
func (o DLPDictionariesOutput) BinNumbers() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *DLPDictionaries) pulumi.IntArrayOutput { return v.BinNumbers }).(pulumi.IntArrayOutput)
}

// The DLP confidence threshold. The following values are supported:
func (o DLPDictionariesOutput) ConfidenceThreshold() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DLPDictionaries) pulumi.StringPtrOutput { return v.ConfidenceThreshold }).(pulumi.StringPtrOutput)
}

// The DLP custom phrase match type. Supported values are:
func (o DLPDictionariesOutput) CustomPhraseMatchType() pulumi.StringOutput {
	return o.ApplyT(func(v *DLPDictionaries) pulumi.StringOutput { return v.CustomPhraseMatchType }).(pulumi.StringOutput)
}

// The desciption of the DLP dictionary
func (o DLPDictionariesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DLPDictionaries) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// ID of the predefined dictionary (original source dictionary) that is used for cloning. This field is applicable only to cloned dictionaries. Only a limited set of identification-based predefined dictionaries (e.g., Credit Cards, Social Security Numbers, National Identification Numbers, etc.) can be cloned. Up to 4 clones can be created from a predefined dictionary.
func (o DLPDictionariesOutput) DictTemplateId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DLPDictionaries) pulumi.IntPtrOutput { return v.DictTemplateId }).(pulumi.IntPtrOutput)
}

func (o DLPDictionariesOutput) DictionaryId() pulumi.IntOutput {
	return o.ApplyT(func(v *DLPDictionaries) pulumi.IntOutput { return v.DictionaryId }).(pulumi.IntOutput)
}

// The DLP dictionary type. The following values are supported:
func (o DLPDictionariesOutput) DictionaryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DLPDictionaries) pulumi.StringPtrOutput { return v.DictionaryType }).(pulumi.StringPtrOutput)
}

// Exact Data Match (EDM) related information for custom DLP dictionaries.
func (o DLPDictionariesOutput) ExactDataMatchDetails() DLPDictionariesExactDataMatchDetailArrayOutput {
	return o.ApplyT(func(v *DLPDictionaries) DLPDictionariesExactDataMatchDetailArrayOutput {
		return v.ExactDataMatchDetails
	}).(DLPDictionariesExactDataMatchDetailArrayOutput)
}

// List of Indexed Document Match (IDM) profiles and their corresponding match accuracy for custom DLP dictionaries.
func (o DLPDictionariesOutput) IdmProfileMatchAccuracies() DLPDictionariesIdmProfileMatchAccuracyArrayOutput {
	return o.ApplyT(func(v *DLPDictionaries) DLPDictionariesIdmProfileMatchAccuracyArrayOutput {
		return v.IdmProfileMatchAccuracies
	}).(DLPDictionariesIdmProfileMatchAccuracyArrayOutput)
}

// Indicates whether to exclude documents that are a 100%!m(MISSING)atch to already-indexed documents from triggering an Indexed Document Match (IDM) Dictionary.
func (o DLPDictionariesOutput) IgnoreExactMatchIdmDict() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DLPDictionaries) pulumi.BoolPtrOutput { return v.IgnoreExactMatchIdmDict }).(pulumi.BoolPtrOutput)
}

// A true value denotes that the specified Bank Identification Number (BIN) values are included in the Credit Cards dictionary. A false value denotes that the specified BIN values are excluded from the Credit Cards dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.
func (o DLPDictionariesOutput) IncludeBinNumbers() pulumi.BoolOutput {
	return o.ApplyT(func(v *DLPDictionaries) pulumi.BoolOutput { return v.IncludeBinNumbers }).(pulumi.BoolOutput)
}

// The DLP dictionary's name
func (o DLPDictionariesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DLPDictionaries) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List containing the patterns used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries. Required when `dictionaryType` is `PATTERNS_AND_PHRASES`
func (o DLPDictionariesOutput) Patterns() DLPDictionariesPatternArrayOutput {
	return o.ApplyT(func(v *DLPDictionaries) DLPDictionariesPatternArrayOutput { return v.Patterns }).(DLPDictionariesPatternArrayOutput)
}

// List containing the phrases used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries. Required when `dictionaryType` is `PATTERNS_AND_PHRASES`
func (o DLPDictionariesOutput) Phrases() DLPDictionariesPhraseArrayOutput {
	return o.ApplyT(func(v *DLPDictionaries) DLPDictionariesPhraseArrayOutput { return v.Phrases }).(DLPDictionariesPhraseArrayOutput)
}

// The DLP dictionary proximity length.
func (o DLPDictionariesOutput) Proximity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DLPDictionaries) pulumi.IntPtrOutput { return v.Proximity }).(pulumi.IntPtrOutput)
}

type DLPDictionariesArrayOutput struct{ *pulumi.OutputState }

func (DLPDictionariesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DLPDictionaries)(nil)).Elem()
}

func (o DLPDictionariesArrayOutput) ToDLPDictionariesArrayOutput() DLPDictionariesArrayOutput {
	return o
}

func (o DLPDictionariesArrayOutput) ToDLPDictionariesArrayOutputWithContext(ctx context.Context) DLPDictionariesArrayOutput {
	return o
}

func (o DLPDictionariesArrayOutput) Index(i pulumi.IntInput) DLPDictionariesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DLPDictionaries {
		return vs[0].([]*DLPDictionaries)[vs[1].(int)]
	}).(DLPDictionariesOutput)
}

type DLPDictionariesMapOutput struct{ *pulumi.OutputState }

func (DLPDictionariesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DLPDictionaries)(nil)).Elem()
}

func (o DLPDictionariesMapOutput) ToDLPDictionariesMapOutput() DLPDictionariesMapOutput {
	return o
}

func (o DLPDictionariesMapOutput) ToDLPDictionariesMapOutputWithContext(ctx context.Context) DLPDictionariesMapOutput {
	return o
}

func (o DLPDictionariesMapOutput) MapIndex(k pulumi.StringInput) DLPDictionariesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DLPDictionaries {
		return vs[0].(map[string]*DLPDictionaries)[vs[1].(string)]
	}).(DLPDictionariesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DLPDictionariesInput)(nil)).Elem(), &DLPDictionaries{})
	pulumi.RegisterInputType(reflect.TypeOf((*DLPDictionariesArrayInput)(nil)).Elem(), DLPDictionariesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DLPDictionariesMapInput)(nil)).Elem(), DLPDictionariesMap{})
	pulumi.RegisterOutputType(DLPDictionariesOutput{})
	pulumi.RegisterOutputType(DLPDictionariesArrayOutput{})
	pulumi.RegisterOutputType(DLPDictionariesMapOutput{})
}
