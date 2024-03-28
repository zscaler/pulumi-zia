// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zia

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zia/sdk/go/zia/internal"
)

// The **zia_rule_labels** resource allows the creation and management of rule labels in the Zscaler Internet Access cloud or via the API. This resource can then be associated with resources such as: Firewall Rules and URL filtering rules
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
//			// ZIA Rule Labels Resource
//			_, err := zia.NewRuleLabels(ctx, "example", &zia.RuleLabelsArgs{
//				Description: pulumi.String("Example"),
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
// **zia_rule_labels** can be imported by using `<LABEL_ID>` or `<LABEL_NAME>` as the import ID.
//
// For example:
//
// ```sh
// $ pulumi import zia:index/ruleLabels:RuleLabels example <label_id>
// ```
//
// or
//
// ```sh
// $ pulumi import zia:index/ruleLabels:RuleLabels example <label_name>
// ```
type RuleLabels struct {
	pulumi.CustomResourceState

	// The rule label description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the devices to be created.
	Name        pulumi.StringOutput `pulumi:"name"`
	RuleLabelId pulumi.IntOutput    `pulumi:"ruleLabelId"`
}

// NewRuleLabels registers a new resource with the given unique name, arguments, and options.
func NewRuleLabels(ctx *pulumi.Context,
	name string, args *RuleLabelsArgs, opts ...pulumi.ResourceOption) (*RuleLabels, error) {
	if args == nil {
		args = &RuleLabelsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RuleLabels
	err := ctx.RegisterResource("zia:index/ruleLabels:RuleLabels", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleLabels gets an existing RuleLabels resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleLabels(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleLabelsState, opts ...pulumi.ResourceOption) (*RuleLabels, error) {
	var resource RuleLabels
	err := ctx.ReadResource("zia:index/ruleLabels:RuleLabels", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleLabels resources.
type ruleLabelsState struct {
	// The rule label description.
	Description *string `pulumi:"description"`
	// The name of the devices to be created.
	Name        *string `pulumi:"name"`
	RuleLabelId *int    `pulumi:"ruleLabelId"`
}

type RuleLabelsState struct {
	// The rule label description.
	Description pulumi.StringPtrInput
	// The name of the devices to be created.
	Name        pulumi.StringPtrInput
	RuleLabelId pulumi.IntPtrInput
}

func (RuleLabelsState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleLabelsState)(nil)).Elem()
}

type ruleLabelsArgs struct {
	// The rule label description.
	Description *string `pulumi:"description"`
	// The name of the devices to be created.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a RuleLabels resource.
type RuleLabelsArgs struct {
	// The rule label description.
	Description pulumi.StringPtrInput
	// The name of the devices to be created.
	Name pulumi.StringPtrInput
}

func (RuleLabelsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleLabelsArgs)(nil)).Elem()
}

type RuleLabelsInput interface {
	pulumi.Input

	ToRuleLabelsOutput() RuleLabelsOutput
	ToRuleLabelsOutputWithContext(ctx context.Context) RuleLabelsOutput
}

func (*RuleLabels) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleLabels)(nil)).Elem()
}

func (i *RuleLabels) ToRuleLabelsOutput() RuleLabelsOutput {
	return i.ToRuleLabelsOutputWithContext(context.Background())
}

func (i *RuleLabels) ToRuleLabelsOutputWithContext(ctx context.Context) RuleLabelsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleLabelsOutput)
}

// RuleLabelsArrayInput is an input type that accepts RuleLabelsArray and RuleLabelsArrayOutput values.
// You can construct a concrete instance of `RuleLabelsArrayInput` via:
//
//	RuleLabelsArray{ RuleLabelsArgs{...} }
type RuleLabelsArrayInput interface {
	pulumi.Input

	ToRuleLabelsArrayOutput() RuleLabelsArrayOutput
	ToRuleLabelsArrayOutputWithContext(context.Context) RuleLabelsArrayOutput
}

type RuleLabelsArray []RuleLabelsInput

func (RuleLabelsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuleLabels)(nil)).Elem()
}

func (i RuleLabelsArray) ToRuleLabelsArrayOutput() RuleLabelsArrayOutput {
	return i.ToRuleLabelsArrayOutputWithContext(context.Background())
}

func (i RuleLabelsArray) ToRuleLabelsArrayOutputWithContext(ctx context.Context) RuleLabelsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleLabelsArrayOutput)
}

// RuleLabelsMapInput is an input type that accepts RuleLabelsMap and RuleLabelsMapOutput values.
// You can construct a concrete instance of `RuleLabelsMapInput` via:
//
//	RuleLabelsMap{ "key": RuleLabelsArgs{...} }
type RuleLabelsMapInput interface {
	pulumi.Input

	ToRuleLabelsMapOutput() RuleLabelsMapOutput
	ToRuleLabelsMapOutputWithContext(context.Context) RuleLabelsMapOutput
}

type RuleLabelsMap map[string]RuleLabelsInput

func (RuleLabelsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuleLabels)(nil)).Elem()
}

func (i RuleLabelsMap) ToRuleLabelsMapOutput() RuleLabelsMapOutput {
	return i.ToRuleLabelsMapOutputWithContext(context.Background())
}

func (i RuleLabelsMap) ToRuleLabelsMapOutputWithContext(ctx context.Context) RuleLabelsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleLabelsMapOutput)
}

type RuleLabelsOutput struct{ *pulumi.OutputState }

func (RuleLabelsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleLabels)(nil)).Elem()
}

func (o RuleLabelsOutput) ToRuleLabelsOutput() RuleLabelsOutput {
	return o
}

func (o RuleLabelsOutput) ToRuleLabelsOutputWithContext(ctx context.Context) RuleLabelsOutput {
	return o
}

// The rule label description.
func (o RuleLabelsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleLabels) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the devices to be created.
func (o RuleLabelsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleLabels) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RuleLabelsOutput) RuleLabelId() pulumi.IntOutput {
	return o.ApplyT(func(v *RuleLabels) pulumi.IntOutput { return v.RuleLabelId }).(pulumi.IntOutput)
}

type RuleLabelsArrayOutput struct{ *pulumi.OutputState }

func (RuleLabelsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuleLabels)(nil)).Elem()
}

func (o RuleLabelsArrayOutput) ToRuleLabelsArrayOutput() RuleLabelsArrayOutput {
	return o
}

func (o RuleLabelsArrayOutput) ToRuleLabelsArrayOutputWithContext(ctx context.Context) RuleLabelsArrayOutput {
	return o
}

func (o RuleLabelsArrayOutput) Index(i pulumi.IntInput) RuleLabelsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RuleLabels {
		return vs[0].([]*RuleLabels)[vs[1].(int)]
	}).(RuleLabelsOutput)
}

type RuleLabelsMapOutput struct{ *pulumi.OutputState }

func (RuleLabelsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuleLabels)(nil)).Elem()
}

func (o RuleLabelsMapOutput) ToRuleLabelsMapOutput() RuleLabelsMapOutput {
	return o
}

func (o RuleLabelsMapOutput) ToRuleLabelsMapOutputWithContext(ctx context.Context) RuleLabelsMapOutput {
	return o
}

func (o RuleLabelsMapOutput) MapIndex(k pulumi.StringInput) RuleLabelsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RuleLabels {
		return vs[0].(map[string]*RuleLabels)[vs[1].(string)]
	}).(RuleLabelsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleLabelsInput)(nil)).Elem(), &RuleLabels{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleLabelsArrayInput)(nil)).Elem(), RuleLabelsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleLabelsMapInput)(nil)).Elem(), RuleLabelsMap{})
	pulumi.RegisterOutputType(RuleLabelsOutput{})
	pulumi.RegisterOutputType(RuleLabelsArrayOutput{})
	pulumi.RegisterOutputType(RuleLabelsMapOutput{})
}
