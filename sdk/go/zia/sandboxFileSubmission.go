// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zia

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zia/sdk/go/zia/internal"
)

// ## Example Usage
type SandboxFileSubmission struct {
	pulumi.CustomResourceState

	Code pulumi.IntOutput `pulumi:"code"`
	// (Required) The path where the raw or archive files for submission are located.
	FilePath pulumi.StringOutput `pulumi:"filePath"`
	FileType pulumi.StringOutput `pulumi:"fileType"`
	// (Optional) Submit file to sandbox even if found malicious during AV scan and a verdict already exists. Supported values are `true` or `false`
	Force             pulumi.BoolPtrOutput `pulumi:"force"`
	Md5               pulumi.StringOutput  `pulumi:"md5"`
	Message           pulumi.StringOutput  `pulumi:"message"`
	SandboxSubmission pulumi.StringOutput  `pulumi:"sandboxSubmission"`
	// (Required) The submission method to be used. Supportedd values are: `submit` and `discan`
	SubmissionMethod pulumi.StringOutput `pulumi:"submissionMethod"`
	VirusName        pulumi.StringOutput `pulumi:"virusName"`
	VirusType        pulumi.StringOutput `pulumi:"virusType"`
}

// NewSandboxFileSubmission registers a new resource with the given unique name, arguments, and options.
func NewSandboxFileSubmission(ctx *pulumi.Context,
	name string, args *SandboxFileSubmissionArgs, opts ...pulumi.ResourceOption) (*SandboxFileSubmission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FilePath == nil {
		return nil, errors.New("invalid value for required argument 'FilePath'")
	}
	if args.SubmissionMethod == nil {
		return nil, errors.New("invalid value for required argument 'SubmissionMethod'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SandboxFileSubmission
	err := ctx.RegisterResource("zia:index/sandboxFileSubmission:SandboxFileSubmission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSandboxFileSubmission gets an existing SandboxFileSubmission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSandboxFileSubmission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SandboxFileSubmissionState, opts ...pulumi.ResourceOption) (*SandboxFileSubmission, error) {
	var resource SandboxFileSubmission
	err := ctx.ReadResource("zia:index/sandboxFileSubmission:SandboxFileSubmission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SandboxFileSubmission resources.
type sandboxFileSubmissionState struct {
	Code *int `pulumi:"code"`
	// (Required) The path where the raw or archive files for submission are located.
	FilePath *string `pulumi:"filePath"`
	FileType *string `pulumi:"fileType"`
	// (Optional) Submit file to sandbox even if found malicious during AV scan and a verdict already exists. Supported values are `true` or `false`
	Force             *bool   `pulumi:"force"`
	Md5               *string `pulumi:"md5"`
	Message           *string `pulumi:"message"`
	SandboxSubmission *string `pulumi:"sandboxSubmission"`
	// (Required) The submission method to be used. Supportedd values are: `submit` and `discan`
	SubmissionMethod *string `pulumi:"submissionMethod"`
	VirusName        *string `pulumi:"virusName"`
	VirusType        *string `pulumi:"virusType"`
}

type SandboxFileSubmissionState struct {
	Code pulumi.IntPtrInput
	// (Required) The path where the raw or archive files for submission are located.
	FilePath pulumi.StringPtrInput
	FileType pulumi.StringPtrInput
	// (Optional) Submit file to sandbox even if found malicious during AV scan and a verdict already exists. Supported values are `true` or `false`
	Force             pulumi.BoolPtrInput
	Md5               pulumi.StringPtrInput
	Message           pulumi.StringPtrInput
	SandboxSubmission pulumi.StringPtrInput
	// (Required) The submission method to be used. Supportedd values are: `submit` and `discan`
	SubmissionMethod pulumi.StringPtrInput
	VirusName        pulumi.StringPtrInput
	VirusType        pulumi.StringPtrInput
}

func (SandboxFileSubmissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*sandboxFileSubmissionState)(nil)).Elem()
}

type sandboxFileSubmissionArgs struct {
	// (Required) The path where the raw or archive files for submission are located.
	FilePath string `pulumi:"filePath"`
	// (Optional) Submit file to sandbox even if found malicious during AV scan and a verdict already exists. Supported values are `true` or `false`
	Force *bool `pulumi:"force"`
	// (Required) The submission method to be used. Supportedd values are: `submit` and `discan`
	SubmissionMethod string `pulumi:"submissionMethod"`
}

// The set of arguments for constructing a SandboxFileSubmission resource.
type SandboxFileSubmissionArgs struct {
	// (Required) The path where the raw or archive files for submission are located.
	FilePath pulumi.StringInput
	// (Optional) Submit file to sandbox even if found malicious during AV scan and a verdict already exists. Supported values are `true` or `false`
	Force pulumi.BoolPtrInput
	// (Required) The submission method to be used. Supportedd values are: `submit` and `discan`
	SubmissionMethod pulumi.StringInput
}

func (SandboxFileSubmissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sandboxFileSubmissionArgs)(nil)).Elem()
}

type SandboxFileSubmissionInput interface {
	pulumi.Input

	ToSandboxFileSubmissionOutput() SandboxFileSubmissionOutput
	ToSandboxFileSubmissionOutputWithContext(ctx context.Context) SandboxFileSubmissionOutput
}

func (*SandboxFileSubmission) ElementType() reflect.Type {
	return reflect.TypeOf((**SandboxFileSubmission)(nil)).Elem()
}

func (i *SandboxFileSubmission) ToSandboxFileSubmissionOutput() SandboxFileSubmissionOutput {
	return i.ToSandboxFileSubmissionOutputWithContext(context.Background())
}

func (i *SandboxFileSubmission) ToSandboxFileSubmissionOutputWithContext(ctx context.Context) SandboxFileSubmissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SandboxFileSubmissionOutput)
}

// SandboxFileSubmissionArrayInput is an input type that accepts SandboxFileSubmissionArray and SandboxFileSubmissionArrayOutput values.
// You can construct a concrete instance of `SandboxFileSubmissionArrayInput` via:
//
//	SandboxFileSubmissionArray{ SandboxFileSubmissionArgs{...} }
type SandboxFileSubmissionArrayInput interface {
	pulumi.Input

	ToSandboxFileSubmissionArrayOutput() SandboxFileSubmissionArrayOutput
	ToSandboxFileSubmissionArrayOutputWithContext(context.Context) SandboxFileSubmissionArrayOutput
}

type SandboxFileSubmissionArray []SandboxFileSubmissionInput

func (SandboxFileSubmissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SandboxFileSubmission)(nil)).Elem()
}

func (i SandboxFileSubmissionArray) ToSandboxFileSubmissionArrayOutput() SandboxFileSubmissionArrayOutput {
	return i.ToSandboxFileSubmissionArrayOutputWithContext(context.Background())
}

func (i SandboxFileSubmissionArray) ToSandboxFileSubmissionArrayOutputWithContext(ctx context.Context) SandboxFileSubmissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SandboxFileSubmissionArrayOutput)
}

// SandboxFileSubmissionMapInput is an input type that accepts SandboxFileSubmissionMap and SandboxFileSubmissionMapOutput values.
// You can construct a concrete instance of `SandboxFileSubmissionMapInput` via:
//
//	SandboxFileSubmissionMap{ "key": SandboxFileSubmissionArgs{...} }
type SandboxFileSubmissionMapInput interface {
	pulumi.Input

	ToSandboxFileSubmissionMapOutput() SandboxFileSubmissionMapOutput
	ToSandboxFileSubmissionMapOutputWithContext(context.Context) SandboxFileSubmissionMapOutput
}

type SandboxFileSubmissionMap map[string]SandboxFileSubmissionInput

func (SandboxFileSubmissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SandboxFileSubmission)(nil)).Elem()
}

func (i SandboxFileSubmissionMap) ToSandboxFileSubmissionMapOutput() SandboxFileSubmissionMapOutput {
	return i.ToSandboxFileSubmissionMapOutputWithContext(context.Background())
}

func (i SandboxFileSubmissionMap) ToSandboxFileSubmissionMapOutputWithContext(ctx context.Context) SandboxFileSubmissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SandboxFileSubmissionMapOutput)
}

type SandboxFileSubmissionOutput struct{ *pulumi.OutputState }

func (SandboxFileSubmissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SandboxFileSubmission)(nil)).Elem()
}

func (o SandboxFileSubmissionOutput) ToSandboxFileSubmissionOutput() SandboxFileSubmissionOutput {
	return o
}

func (o SandboxFileSubmissionOutput) ToSandboxFileSubmissionOutputWithContext(ctx context.Context) SandboxFileSubmissionOutput {
	return o
}

func (o SandboxFileSubmissionOutput) Code() pulumi.IntOutput {
	return o.ApplyT(func(v *SandboxFileSubmission) pulumi.IntOutput { return v.Code }).(pulumi.IntOutput)
}

// (Required) The path where the raw or archive files for submission are located.
func (o SandboxFileSubmissionOutput) FilePath() pulumi.StringOutput {
	return o.ApplyT(func(v *SandboxFileSubmission) pulumi.StringOutput { return v.FilePath }).(pulumi.StringOutput)
}

func (o SandboxFileSubmissionOutput) FileType() pulumi.StringOutput {
	return o.ApplyT(func(v *SandboxFileSubmission) pulumi.StringOutput { return v.FileType }).(pulumi.StringOutput)
}

// (Optional) Submit file to sandbox even if found malicious during AV scan and a verdict already exists. Supported values are `true` or `false`
func (o SandboxFileSubmissionOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SandboxFileSubmission) pulumi.BoolPtrOutput { return v.Force }).(pulumi.BoolPtrOutput)
}

func (o SandboxFileSubmissionOutput) Md5() pulumi.StringOutput {
	return o.ApplyT(func(v *SandboxFileSubmission) pulumi.StringOutput { return v.Md5 }).(pulumi.StringOutput)
}

func (o SandboxFileSubmissionOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v *SandboxFileSubmission) pulumi.StringOutput { return v.Message }).(pulumi.StringOutput)
}

func (o SandboxFileSubmissionOutput) SandboxSubmission() pulumi.StringOutput {
	return o.ApplyT(func(v *SandboxFileSubmission) pulumi.StringOutput { return v.SandboxSubmission }).(pulumi.StringOutput)
}

// (Required) The submission method to be used. Supportedd values are: `submit` and `discan`
func (o SandboxFileSubmissionOutput) SubmissionMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *SandboxFileSubmission) pulumi.StringOutput { return v.SubmissionMethod }).(pulumi.StringOutput)
}

func (o SandboxFileSubmissionOutput) VirusName() pulumi.StringOutput {
	return o.ApplyT(func(v *SandboxFileSubmission) pulumi.StringOutput { return v.VirusName }).(pulumi.StringOutput)
}

func (o SandboxFileSubmissionOutput) VirusType() pulumi.StringOutput {
	return o.ApplyT(func(v *SandboxFileSubmission) pulumi.StringOutput { return v.VirusType }).(pulumi.StringOutput)
}

type SandboxFileSubmissionArrayOutput struct{ *pulumi.OutputState }

func (SandboxFileSubmissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SandboxFileSubmission)(nil)).Elem()
}

func (o SandboxFileSubmissionArrayOutput) ToSandboxFileSubmissionArrayOutput() SandboxFileSubmissionArrayOutput {
	return o
}

func (o SandboxFileSubmissionArrayOutput) ToSandboxFileSubmissionArrayOutputWithContext(ctx context.Context) SandboxFileSubmissionArrayOutput {
	return o
}

func (o SandboxFileSubmissionArrayOutput) Index(i pulumi.IntInput) SandboxFileSubmissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SandboxFileSubmission {
		return vs[0].([]*SandboxFileSubmission)[vs[1].(int)]
	}).(SandboxFileSubmissionOutput)
}

type SandboxFileSubmissionMapOutput struct{ *pulumi.OutputState }

func (SandboxFileSubmissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SandboxFileSubmission)(nil)).Elem()
}

func (o SandboxFileSubmissionMapOutput) ToSandboxFileSubmissionMapOutput() SandboxFileSubmissionMapOutput {
	return o
}

func (o SandboxFileSubmissionMapOutput) ToSandboxFileSubmissionMapOutputWithContext(ctx context.Context) SandboxFileSubmissionMapOutput {
	return o
}

func (o SandboxFileSubmissionMapOutput) MapIndex(k pulumi.StringInput) SandboxFileSubmissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SandboxFileSubmission {
		return vs[0].(map[string]*SandboxFileSubmission)[vs[1].(string)]
	}).(SandboxFileSubmissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SandboxFileSubmissionInput)(nil)).Elem(), &SandboxFileSubmission{})
	pulumi.RegisterInputType(reflect.TypeOf((*SandboxFileSubmissionArrayInput)(nil)).Elem(), SandboxFileSubmissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SandboxFileSubmissionMapInput)(nil)).Elem(), SandboxFileSubmissionMap{})
	pulumi.RegisterOutputType(SandboxFileSubmissionOutput{})
	pulumi.RegisterOutputType(SandboxFileSubmissionArrayOutput{})
	pulumi.RegisterOutputType(SandboxFileSubmissionMapOutput{})
}
