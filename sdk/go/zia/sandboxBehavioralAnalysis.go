// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zia

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zia/sdk/go/zia/internal"
)

// The **zia_sandbox_behavioral_analysis** resource updates the custom list of MD5 file hashes that are blocked by Sandbox. This overwrites a previously generated blocklist. If you need to completely erase the blocklist, submit an empty list.
//
// **Note**: Only the file types that are supported by Sandbox analysis can be blocked using MD5 hashes.
//
// ## Example Usage
//
// ### Add MD5 Hashes To Sandbox
//
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
//			// Add MD5 Hashes to Sandbox
//			_, err := zia.NewSandboxBehavioralAnalysis(ctx, "this", &zia.SandboxBehavioralAnalysisArgs{
//				FileHashesToBeBlockeds: pulumi.StringArray{
//					pulumi.String("42914d6d213a20a2684064be5c80ffa9"),
//					pulumi.String("c0202cf6aeab8437c638533d14563d35"),
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
//
// ### Remove All MD5 Hashes To Sandbox
//
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
//			// Remove All MD5 Hashes to Sandbox
//			_, err := zia.NewSandboxBehavioralAnalysis(ctx, "this", &zia.SandboxBehavioralAnalysisArgs{
//				FileHashesToBeBlockeds: pulumi.StringArray{},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.
//
// # Visit
//
// **zia_sandbox_behavioral_analysis** can be imported by using `sandbox_settings` as the import ID.
//
// For example:
//
// ```sh
// $ pulumi import zia:index/sandboxBehavioralAnalysis:SandboxBehavioralAnalysis example sandbox_settings
// ```
type SandboxBehavioralAnalysis struct {
	pulumi.CustomResourceState

	// A custom list of unique MD5 file hashes that must be blocked by Sandbox. A maximum of 10000 MD5 file hashes can be
	// blocked
	FileHashesToBeBlockeds pulumi.StringArrayOutput `pulumi:"fileHashesToBeBlockeds"`
}

// NewSandboxBehavioralAnalysis registers a new resource with the given unique name, arguments, and options.
func NewSandboxBehavioralAnalysis(ctx *pulumi.Context,
	name string, args *SandboxBehavioralAnalysisArgs, opts ...pulumi.ResourceOption) (*SandboxBehavioralAnalysis, error) {
	if args == nil {
		args = &SandboxBehavioralAnalysisArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SandboxBehavioralAnalysis
	err := ctx.RegisterResource("zia:index/sandboxBehavioralAnalysis:SandboxBehavioralAnalysis", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSandboxBehavioralAnalysis gets an existing SandboxBehavioralAnalysis resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSandboxBehavioralAnalysis(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SandboxBehavioralAnalysisState, opts ...pulumi.ResourceOption) (*SandboxBehavioralAnalysis, error) {
	var resource SandboxBehavioralAnalysis
	err := ctx.ReadResource("zia:index/sandboxBehavioralAnalysis:SandboxBehavioralAnalysis", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SandboxBehavioralAnalysis resources.
type sandboxBehavioralAnalysisState struct {
	// A custom list of unique MD5 file hashes that must be blocked by Sandbox. A maximum of 10000 MD5 file hashes can be
	// blocked
	FileHashesToBeBlockeds []string `pulumi:"fileHashesToBeBlockeds"`
}

type SandboxBehavioralAnalysisState struct {
	// A custom list of unique MD5 file hashes that must be blocked by Sandbox. A maximum of 10000 MD5 file hashes can be
	// blocked
	FileHashesToBeBlockeds pulumi.StringArrayInput
}

func (SandboxBehavioralAnalysisState) ElementType() reflect.Type {
	return reflect.TypeOf((*sandboxBehavioralAnalysisState)(nil)).Elem()
}

type sandboxBehavioralAnalysisArgs struct {
	// A custom list of unique MD5 file hashes that must be blocked by Sandbox. A maximum of 10000 MD5 file hashes can be
	// blocked
	FileHashesToBeBlockeds []string `pulumi:"fileHashesToBeBlockeds"`
}

// The set of arguments for constructing a SandboxBehavioralAnalysis resource.
type SandboxBehavioralAnalysisArgs struct {
	// A custom list of unique MD5 file hashes that must be blocked by Sandbox. A maximum of 10000 MD5 file hashes can be
	// blocked
	FileHashesToBeBlockeds pulumi.StringArrayInput
}

func (SandboxBehavioralAnalysisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sandboxBehavioralAnalysisArgs)(nil)).Elem()
}

type SandboxBehavioralAnalysisInput interface {
	pulumi.Input

	ToSandboxBehavioralAnalysisOutput() SandboxBehavioralAnalysisOutput
	ToSandboxBehavioralAnalysisOutputWithContext(ctx context.Context) SandboxBehavioralAnalysisOutput
}

func (*SandboxBehavioralAnalysis) ElementType() reflect.Type {
	return reflect.TypeOf((**SandboxBehavioralAnalysis)(nil)).Elem()
}

func (i *SandboxBehavioralAnalysis) ToSandboxBehavioralAnalysisOutput() SandboxBehavioralAnalysisOutput {
	return i.ToSandboxBehavioralAnalysisOutputWithContext(context.Background())
}

func (i *SandboxBehavioralAnalysis) ToSandboxBehavioralAnalysisOutputWithContext(ctx context.Context) SandboxBehavioralAnalysisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SandboxBehavioralAnalysisOutput)
}

// SandboxBehavioralAnalysisArrayInput is an input type that accepts SandboxBehavioralAnalysisArray and SandboxBehavioralAnalysisArrayOutput values.
// You can construct a concrete instance of `SandboxBehavioralAnalysisArrayInput` via:
//
//	SandboxBehavioralAnalysisArray{ SandboxBehavioralAnalysisArgs{...} }
type SandboxBehavioralAnalysisArrayInput interface {
	pulumi.Input

	ToSandboxBehavioralAnalysisArrayOutput() SandboxBehavioralAnalysisArrayOutput
	ToSandboxBehavioralAnalysisArrayOutputWithContext(context.Context) SandboxBehavioralAnalysisArrayOutput
}

type SandboxBehavioralAnalysisArray []SandboxBehavioralAnalysisInput

func (SandboxBehavioralAnalysisArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SandboxBehavioralAnalysis)(nil)).Elem()
}

func (i SandboxBehavioralAnalysisArray) ToSandboxBehavioralAnalysisArrayOutput() SandboxBehavioralAnalysisArrayOutput {
	return i.ToSandboxBehavioralAnalysisArrayOutputWithContext(context.Background())
}

func (i SandboxBehavioralAnalysisArray) ToSandboxBehavioralAnalysisArrayOutputWithContext(ctx context.Context) SandboxBehavioralAnalysisArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SandboxBehavioralAnalysisArrayOutput)
}

// SandboxBehavioralAnalysisMapInput is an input type that accepts SandboxBehavioralAnalysisMap and SandboxBehavioralAnalysisMapOutput values.
// You can construct a concrete instance of `SandboxBehavioralAnalysisMapInput` via:
//
//	SandboxBehavioralAnalysisMap{ "key": SandboxBehavioralAnalysisArgs{...} }
type SandboxBehavioralAnalysisMapInput interface {
	pulumi.Input

	ToSandboxBehavioralAnalysisMapOutput() SandboxBehavioralAnalysisMapOutput
	ToSandboxBehavioralAnalysisMapOutputWithContext(context.Context) SandboxBehavioralAnalysisMapOutput
}

type SandboxBehavioralAnalysisMap map[string]SandboxBehavioralAnalysisInput

func (SandboxBehavioralAnalysisMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SandboxBehavioralAnalysis)(nil)).Elem()
}

func (i SandboxBehavioralAnalysisMap) ToSandboxBehavioralAnalysisMapOutput() SandboxBehavioralAnalysisMapOutput {
	return i.ToSandboxBehavioralAnalysisMapOutputWithContext(context.Background())
}

func (i SandboxBehavioralAnalysisMap) ToSandboxBehavioralAnalysisMapOutputWithContext(ctx context.Context) SandboxBehavioralAnalysisMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SandboxBehavioralAnalysisMapOutput)
}

type SandboxBehavioralAnalysisOutput struct{ *pulumi.OutputState }

func (SandboxBehavioralAnalysisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SandboxBehavioralAnalysis)(nil)).Elem()
}

func (o SandboxBehavioralAnalysisOutput) ToSandboxBehavioralAnalysisOutput() SandboxBehavioralAnalysisOutput {
	return o
}

func (o SandboxBehavioralAnalysisOutput) ToSandboxBehavioralAnalysisOutputWithContext(ctx context.Context) SandboxBehavioralAnalysisOutput {
	return o
}

// A custom list of unique MD5 file hashes that must be blocked by Sandbox. A maximum of 10000 MD5 file hashes can be
// blocked
func (o SandboxBehavioralAnalysisOutput) FileHashesToBeBlockeds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SandboxBehavioralAnalysis) pulumi.StringArrayOutput { return v.FileHashesToBeBlockeds }).(pulumi.StringArrayOutput)
}

type SandboxBehavioralAnalysisArrayOutput struct{ *pulumi.OutputState }

func (SandboxBehavioralAnalysisArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SandboxBehavioralAnalysis)(nil)).Elem()
}

func (o SandboxBehavioralAnalysisArrayOutput) ToSandboxBehavioralAnalysisArrayOutput() SandboxBehavioralAnalysisArrayOutput {
	return o
}

func (o SandboxBehavioralAnalysisArrayOutput) ToSandboxBehavioralAnalysisArrayOutputWithContext(ctx context.Context) SandboxBehavioralAnalysisArrayOutput {
	return o
}

func (o SandboxBehavioralAnalysisArrayOutput) Index(i pulumi.IntInput) SandboxBehavioralAnalysisOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SandboxBehavioralAnalysis {
		return vs[0].([]*SandboxBehavioralAnalysis)[vs[1].(int)]
	}).(SandboxBehavioralAnalysisOutput)
}

type SandboxBehavioralAnalysisMapOutput struct{ *pulumi.OutputState }

func (SandboxBehavioralAnalysisMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SandboxBehavioralAnalysis)(nil)).Elem()
}

func (o SandboxBehavioralAnalysisMapOutput) ToSandboxBehavioralAnalysisMapOutput() SandboxBehavioralAnalysisMapOutput {
	return o
}

func (o SandboxBehavioralAnalysisMapOutput) ToSandboxBehavioralAnalysisMapOutputWithContext(ctx context.Context) SandboxBehavioralAnalysisMapOutput {
	return o
}

func (o SandboxBehavioralAnalysisMapOutput) MapIndex(k pulumi.StringInput) SandboxBehavioralAnalysisOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SandboxBehavioralAnalysis {
		return vs[0].(map[string]*SandboxBehavioralAnalysis)[vs[1].(string)]
	}).(SandboxBehavioralAnalysisOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SandboxBehavioralAnalysisInput)(nil)).Elem(), &SandboxBehavioralAnalysis{})
	pulumi.RegisterInputType(reflect.TypeOf((*SandboxBehavioralAnalysisArrayInput)(nil)).Elem(), SandboxBehavioralAnalysisArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SandboxBehavioralAnalysisMapInput)(nil)).Elem(), SandboxBehavioralAnalysisMap{})
	pulumi.RegisterOutputType(SandboxBehavioralAnalysisOutput{})
	pulumi.RegisterOutputType(SandboxBehavioralAnalysisArrayOutput{})
	pulumi.RegisterOutputType(SandboxBehavioralAnalysisMapOutput{})
}
