// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zia

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zia/sdk/go/zia/internal"
)

// * [Official documentation](https://help.zscaler.com/zia/about-alert-subscriptions)
// * [API documentation](https://help.zscaler.com/zia/alerts#/alertSubscriptions-get)
//
// Use the **zia_subscription_alert** resource allows the creation and management of Alert Subscriptions in the Zscaler Internet Access.
//
// ## Example Usage
//
// ## Import
//
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.
//
// # Visit
//
// **zia_subscription_alert** can be imported by using `<ALERT_ID>` or `<ALERT_EMAIL>` as the import ID.
//
// For example:
//
// ```sh
// $ pulumi import zia:index/subscriptionAlert:SubscriptionAlert example <alert_id>
// ```
//
// or
//
// ```sh
// $ pulumi import zia:index/subscriptionAlert:SubscriptionAlert example <alert_email>
// ```
type SubscriptionAlert struct {
	pulumi.CustomResourceState

	// System-generated identifier for the alert subscription
	AlertId pulumi.IntOutput `pulumi:"alertId"`
	// (List of String) Lists the severity levels of the Comply Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	ComplySeverities pulumi.StringArrayOutput `pulumi:"complySeverities"`
	// (String) Additional comments or information about the alert subscription
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The email address of the alert recipient
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// (List of String) Lists the severity levels of the Manage Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	ManageSeverities pulumi.StringArrayOutput `pulumi:"manageSeverities"`
	// (List of String) Lists the severity levels of the Patient 0 Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	Pt0Severities pulumi.StringArrayOutput `pulumi:"pt0Severities"`
	// (List of String) Lists the severity levels of the Secure Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	SecureSeverities pulumi.StringArrayOutput `pulumi:"secureSeverities"`
	// (List of String) Lists the severity levels of the System Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	SystemSeverities pulumi.StringArrayOutput `pulumi:"systemSeverities"`
}

// NewSubscriptionAlert registers a new resource with the given unique name, arguments, and options.
func NewSubscriptionAlert(ctx *pulumi.Context,
	name string, args *SubscriptionAlertArgs, opts ...pulumi.ResourceOption) (*SubscriptionAlert, error) {
	if args == nil {
		args = &SubscriptionAlertArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SubscriptionAlert
	err := ctx.RegisterResource("zia:index/subscriptionAlert:SubscriptionAlert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscriptionAlert gets an existing SubscriptionAlert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscriptionAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionAlertState, opts ...pulumi.ResourceOption) (*SubscriptionAlert, error) {
	var resource SubscriptionAlert
	err := ctx.ReadResource("zia:index/subscriptionAlert:SubscriptionAlert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubscriptionAlert resources.
type subscriptionAlertState struct {
	// System-generated identifier for the alert subscription
	AlertId *int `pulumi:"alertId"`
	// (List of String) Lists the severity levels of the Comply Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	ComplySeverities []string `pulumi:"complySeverities"`
	// (String) Additional comments or information about the alert subscription
	Description *string `pulumi:"description"`
	// The email address of the alert recipient
	Email *string `pulumi:"email"`
	// (List of String) Lists the severity levels of the Manage Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	ManageSeverities []string `pulumi:"manageSeverities"`
	// (List of String) Lists the severity levels of the Patient 0 Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	Pt0Severities []string `pulumi:"pt0Severities"`
	// (List of String) Lists the severity levels of the Secure Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	SecureSeverities []string `pulumi:"secureSeverities"`
	// (List of String) Lists the severity levels of the System Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	SystemSeverities []string `pulumi:"systemSeverities"`
}

type SubscriptionAlertState struct {
	// System-generated identifier for the alert subscription
	AlertId pulumi.IntPtrInput
	// (List of String) Lists the severity levels of the Comply Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	ComplySeverities pulumi.StringArrayInput
	// (String) Additional comments or information about the alert subscription
	Description pulumi.StringPtrInput
	// The email address of the alert recipient
	Email pulumi.StringPtrInput
	// (List of String) Lists the severity levels of the Manage Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	ManageSeverities pulumi.StringArrayInput
	// (List of String) Lists the severity levels of the Patient 0 Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	Pt0Severities pulumi.StringArrayInput
	// (List of String) Lists the severity levels of the Secure Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	SecureSeverities pulumi.StringArrayInput
	// (List of String) Lists the severity levels of the System Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	SystemSeverities pulumi.StringArrayInput
}

func (SubscriptionAlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionAlertState)(nil)).Elem()
}

type subscriptionAlertArgs struct {
	// (List of String) Lists the severity levels of the Comply Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	ComplySeverities []string `pulumi:"complySeverities"`
	// (String) Additional comments or information about the alert subscription
	Description *string `pulumi:"description"`
	// The email address of the alert recipient
	Email *string `pulumi:"email"`
	// (List of String) Lists the severity levels of the Manage Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	ManageSeverities []string `pulumi:"manageSeverities"`
	// (List of String) Lists the severity levels of the Patient 0 Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	Pt0Severities []string `pulumi:"pt0Severities"`
	// (List of String) Lists the severity levels of the Secure Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	SecureSeverities []string `pulumi:"secureSeverities"`
	// (List of String) Lists the severity levels of the System Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	SystemSeverities []string `pulumi:"systemSeverities"`
}

// The set of arguments for constructing a SubscriptionAlert resource.
type SubscriptionAlertArgs struct {
	// (List of String) Lists the severity levels of the Comply Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	ComplySeverities pulumi.StringArrayInput
	// (String) Additional comments or information about the alert subscription
	Description pulumi.StringPtrInput
	// The email address of the alert recipient
	Email pulumi.StringPtrInput
	// (List of String) Lists the severity levels of the Manage Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	ManageSeverities pulumi.StringArrayInput
	// (List of String) Lists the severity levels of the Patient 0 Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	Pt0Severities pulumi.StringArrayInput
	// (List of String) Lists the severity levels of the Secure Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	SecureSeverities pulumi.StringArrayInput
	// (List of String) Lists the severity levels of the System Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
	SystemSeverities pulumi.StringArrayInput
}

func (SubscriptionAlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionAlertArgs)(nil)).Elem()
}

type SubscriptionAlertInput interface {
	pulumi.Input

	ToSubscriptionAlertOutput() SubscriptionAlertOutput
	ToSubscriptionAlertOutputWithContext(ctx context.Context) SubscriptionAlertOutput
}

func (*SubscriptionAlert) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionAlert)(nil)).Elem()
}

func (i *SubscriptionAlert) ToSubscriptionAlertOutput() SubscriptionAlertOutput {
	return i.ToSubscriptionAlertOutputWithContext(context.Background())
}

func (i *SubscriptionAlert) ToSubscriptionAlertOutputWithContext(ctx context.Context) SubscriptionAlertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionAlertOutput)
}

// SubscriptionAlertArrayInput is an input type that accepts SubscriptionAlertArray and SubscriptionAlertArrayOutput values.
// You can construct a concrete instance of `SubscriptionAlertArrayInput` via:
//
//	SubscriptionAlertArray{ SubscriptionAlertArgs{...} }
type SubscriptionAlertArrayInput interface {
	pulumi.Input

	ToSubscriptionAlertArrayOutput() SubscriptionAlertArrayOutput
	ToSubscriptionAlertArrayOutputWithContext(context.Context) SubscriptionAlertArrayOutput
}

type SubscriptionAlertArray []SubscriptionAlertInput

func (SubscriptionAlertArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubscriptionAlert)(nil)).Elem()
}

func (i SubscriptionAlertArray) ToSubscriptionAlertArrayOutput() SubscriptionAlertArrayOutput {
	return i.ToSubscriptionAlertArrayOutputWithContext(context.Background())
}

func (i SubscriptionAlertArray) ToSubscriptionAlertArrayOutputWithContext(ctx context.Context) SubscriptionAlertArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionAlertArrayOutput)
}

// SubscriptionAlertMapInput is an input type that accepts SubscriptionAlertMap and SubscriptionAlertMapOutput values.
// You can construct a concrete instance of `SubscriptionAlertMapInput` via:
//
//	SubscriptionAlertMap{ "key": SubscriptionAlertArgs{...} }
type SubscriptionAlertMapInput interface {
	pulumi.Input

	ToSubscriptionAlertMapOutput() SubscriptionAlertMapOutput
	ToSubscriptionAlertMapOutputWithContext(context.Context) SubscriptionAlertMapOutput
}

type SubscriptionAlertMap map[string]SubscriptionAlertInput

func (SubscriptionAlertMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubscriptionAlert)(nil)).Elem()
}

func (i SubscriptionAlertMap) ToSubscriptionAlertMapOutput() SubscriptionAlertMapOutput {
	return i.ToSubscriptionAlertMapOutputWithContext(context.Background())
}

func (i SubscriptionAlertMap) ToSubscriptionAlertMapOutputWithContext(ctx context.Context) SubscriptionAlertMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionAlertMapOutput)
}

type SubscriptionAlertOutput struct{ *pulumi.OutputState }

func (SubscriptionAlertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionAlert)(nil)).Elem()
}

func (o SubscriptionAlertOutput) ToSubscriptionAlertOutput() SubscriptionAlertOutput {
	return o
}

func (o SubscriptionAlertOutput) ToSubscriptionAlertOutputWithContext(ctx context.Context) SubscriptionAlertOutput {
	return o
}

// System-generated identifier for the alert subscription
func (o SubscriptionAlertOutput) AlertId() pulumi.IntOutput {
	return o.ApplyT(func(v *SubscriptionAlert) pulumi.IntOutput { return v.AlertId }).(pulumi.IntOutput)
}

// (List of String) Lists the severity levels of the Comply Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
func (o SubscriptionAlertOutput) ComplySeverities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubscriptionAlert) pulumi.StringArrayOutput { return v.ComplySeverities }).(pulumi.StringArrayOutput)
}

// (String) Additional comments or information about the alert subscription
func (o SubscriptionAlertOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionAlert) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The email address of the alert recipient
func (o SubscriptionAlertOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionAlert) pulumi.StringPtrOutput { return v.Email }).(pulumi.StringPtrOutput)
}

// (List of String) Lists the severity levels of the Manage Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
func (o SubscriptionAlertOutput) ManageSeverities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubscriptionAlert) pulumi.StringArrayOutput { return v.ManageSeverities }).(pulumi.StringArrayOutput)
}

// (List of String) Lists the severity levels of the Patient 0 Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
func (o SubscriptionAlertOutput) Pt0Severities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubscriptionAlert) pulumi.StringArrayOutput { return v.Pt0Severities }).(pulumi.StringArrayOutput)
}

// (List of String) Lists the severity levels of the Secure Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
func (o SubscriptionAlertOutput) SecureSeverities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubscriptionAlert) pulumi.StringArrayOutput { return v.SecureSeverities }).(pulumi.StringArrayOutput)
}

// (List of String) Lists the severity levels of the System Severity Alert class information that the recipient receives. Supported Values: `CRITICAL`, `MAJOR`, `MINOR`, `INFO`, `DEBUG`
func (o SubscriptionAlertOutput) SystemSeverities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubscriptionAlert) pulumi.StringArrayOutput { return v.SystemSeverities }).(pulumi.StringArrayOutput)
}

type SubscriptionAlertArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionAlertArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubscriptionAlert)(nil)).Elem()
}

func (o SubscriptionAlertArrayOutput) ToSubscriptionAlertArrayOutput() SubscriptionAlertArrayOutput {
	return o
}

func (o SubscriptionAlertArrayOutput) ToSubscriptionAlertArrayOutputWithContext(ctx context.Context) SubscriptionAlertArrayOutput {
	return o
}

func (o SubscriptionAlertArrayOutput) Index(i pulumi.IntInput) SubscriptionAlertOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SubscriptionAlert {
		return vs[0].([]*SubscriptionAlert)[vs[1].(int)]
	}).(SubscriptionAlertOutput)
}

type SubscriptionAlertMapOutput struct{ *pulumi.OutputState }

func (SubscriptionAlertMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubscriptionAlert)(nil)).Elem()
}

func (o SubscriptionAlertMapOutput) ToSubscriptionAlertMapOutput() SubscriptionAlertMapOutput {
	return o
}

func (o SubscriptionAlertMapOutput) ToSubscriptionAlertMapOutputWithContext(ctx context.Context) SubscriptionAlertMapOutput {
	return o
}

func (o SubscriptionAlertMapOutput) MapIndex(k pulumi.StringInput) SubscriptionAlertOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SubscriptionAlert {
		return vs[0].(map[string]*SubscriptionAlert)[vs[1].(string)]
	}).(SubscriptionAlertOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionAlertInput)(nil)).Elem(), &SubscriptionAlert{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionAlertArrayInput)(nil)).Elem(), SubscriptionAlertArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionAlertMapInput)(nil)).Elem(), SubscriptionAlertMap{})
	pulumi.RegisterOutputType(SubscriptionAlertOutput{})
	pulumi.RegisterOutputType(SubscriptionAlertArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionAlertMapOutput{})
}
