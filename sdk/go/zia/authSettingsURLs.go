// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zia

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zia/sdk/go/zia/internal"
)

// * [Official documentation](https://help.zscaler.com/zia/url-format-guidelines)
// * [API documentation](https://help.zscaler.com/zia/user-authentication-settings#/authSettings/exemptedUrls-get)
//
// The **zia_auth_settings_urls** resource alows you to add or remove a URL from the cookie authentication exempt list in the Zscaler Internet Access cloud or via the API. To learn more see [URL Format Guidelines](https://help.zscaler.com/zia/url-format-guidelines)
//
// ## Example Usage
//
// ## Import
//
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.
//
// # Visit
//
// **zia_auth_settings_urls** can be imported by using `all_urls` as the import ID.
//
// For example:
//
// ```sh
// $ pulumi import zia:index/authSettingsURLs:AuthSettingsURLs example all_urls
// ```
type AuthSettingsURLs struct {
	pulumi.CustomResourceState

	Urls pulumi.StringArrayOutput `pulumi:"urls"`
}

// NewAuthSettingsURLs registers a new resource with the given unique name, arguments, and options.
func NewAuthSettingsURLs(ctx *pulumi.Context,
	name string, args *AuthSettingsURLsArgs, opts ...pulumi.ResourceOption) (*AuthSettingsURLs, error) {
	if args == nil {
		args = &AuthSettingsURLsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuthSettingsURLs
	err := ctx.RegisterResource("zia:index/authSettingsURLs:AuthSettingsURLs", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthSettingsURLs gets an existing AuthSettingsURLs resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthSettingsURLs(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthSettingsURLsState, opts ...pulumi.ResourceOption) (*AuthSettingsURLs, error) {
	var resource AuthSettingsURLs
	err := ctx.ReadResource("zia:index/authSettingsURLs:AuthSettingsURLs", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthSettingsURLs resources.
type authSettingsURLsState struct {
	Urls []string `pulumi:"urls"`
}

type AuthSettingsURLsState struct {
	Urls pulumi.StringArrayInput
}

func (AuthSettingsURLsState) ElementType() reflect.Type {
	return reflect.TypeOf((*authSettingsURLsState)(nil)).Elem()
}

type authSettingsURLsArgs struct {
	Urls []string `pulumi:"urls"`
}

// The set of arguments for constructing a AuthSettingsURLs resource.
type AuthSettingsURLsArgs struct {
	Urls pulumi.StringArrayInput
}

func (AuthSettingsURLsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authSettingsURLsArgs)(nil)).Elem()
}

type AuthSettingsURLsInput interface {
	pulumi.Input

	ToAuthSettingsURLsOutput() AuthSettingsURLsOutput
	ToAuthSettingsURLsOutputWithContext(ctx context.Context) AuthSettingsURLsOutput
}

func (*AuthSettingsURLs) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthSettingsURLs)(nil)).Elem()
}

func (i *AuthSettingsURLs) ToAuthSettingsURLsOutput() AuthSettingsURLsOutput {
	return i.ToAuthSettingsURLsOutputWithContext(context.Background())
}

func (i *AuthSettingsURLs) ToAuthSettingsURLsOutputWithContext(ctx context.Context) AuthSettingsURLsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthSettingsURLsOutput)
}

// AuthSettingsURLsArrayInput is an input type that accepts AuthSettingsURLsArray and AuthSettingsURLsArrayOutput values.
// You can construct a concrete instance of `AuthSettingsURLsArrayInput` via:
//
//	AuthSettingsURLsArray{ AuthSettingsURLsArgs{...} }
type AuthSettingsURLsArrayInput interface {
	pulumi.Input

	ToAuthSettingsURLsArrayOutput() AuthSettingsURLsArrayOutput
	ToAuthSettingsURLsArrayOutputWithContext(context.Context) AuthSettingsURLsArrayOutput
}

type AuthSettingsURLsArray []AuthSettingsURLsInput

func (AuthSettingsURLsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthSettingsURLs)(nil)).Elem()
}

func (i AuthSettingsURLsArray) ToAuthSettingsURLsArrayOutput() AuthSettingsURLsArrayOutput {
	return i.ToAuthSettingsURLsArrayOutputWithContext(context.Background())
}

func (i AuthSettingsURLsArray) ToAuthSettingsURLsArrayOutputWithContext(ctx context.Context) AuthSettingsURLsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthSettingsURLsArrayOutput)
}

// AuthSettingsURLsMapInput is an input type that accepts AuthSettingsURLsMap and AuthSettingsURLsMapOutput values.
// You can construct a concrete instance of `AuthSettingsURLsMapInput` via:
//
//	AuthSettingsURLsMap{ "key": AuthSettingsURLsArgs{...} }
type AuthSettingsURLsMapInput interface {
	pulumi.Input

	ToAuthSettingsURLsMapOutput() AuthSettingsURLsMapOutput
	ToAuthSettingsURLsMapOutputWithContext(context.Context) AuthSettingsURLsMapOutput
}

type AuthSettingsURLsMap map[string]AuthSettingsURLsInput

func (AuthSettingsURLsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthSettingsURLs)(nil)).Elem()
}

func (i AuthSettingsURLsMap) ToAuthSettingsURLsMapOutput() AuthSettingsURLsMapOutput {
	return i.ToAuthSettingsURLsMapOutputWithContext(context.Background())
}

func (i AuthSettingsURLsMap) ToAuthSettingsURLsMapOutputWithContext(ctx context.Context) AuthSettingsURLsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthSettingsURLsMapOutput)
}

type AuthSettingsURLsOutput struct{ *pulumi.OutputState }

func (AuthSettingsURLsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthSettingsURLs)(nil)).Elem()
}

func (o AuthSettingsURLsOutput) ToAuthSettingsURLsOutput() AuthSettingsURLsOutput {
	return o
}

func (o AuthSettingsURLsOutput) ToAuthSettingsURLsOutputWithContext(ctx context.Context) AuthSettingsURLsOutput {
	return o
}

func (o AuthSettingsURLsOutput) Urls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthSettingsURLs) pulumi.StringArrayOutput { return v.Urls }).(pulumi.StringArrayOutput)
}

type AuthSettingsURLsArrayOutput struct{ *pulumi.OutputState }

func (AuthSettingsURLsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthSettingsURLs)(nil)).Elem()
}

func (o AuthSettingsURLsArrayOutput) ToAuthSettingsURLsArrayOutput() AuthSettingsURLsArrayOutput {
	return o
}

func (o AuthSettingsURLsArrayOutput) ToAuthSettingsURLsArrayOutputWithContext(ctx context.Context) AuthSettingsURLsArrayOutput {
	return o
}

func (o AuthSettingsURLsArrayOutput) Index(i pulumi.IntInput) AuthSettingsURLsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthSettingsURLs {
		return vs[0].([]*AuthSettingsURLs)[vs[1].(int)]
	}).(AuthSettingsURLsOutput)
}

type AuthSettingsURLsMapOutput struct{ *pulumi.OutputState }

func (AuthSettingsURLsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthSettingsURLs)(nil)).Elem()
}

func (o AuthSettingsURLsMapOutput) ToAuthSettingsURLsMapOutput() AuthSettingsURLsMapOutput {
	return o
}

func (o AuthSettingsURLsMapOutput) ToAuthSettingsURLsMapOutputWithContext(ctx context.Context) AuthSettingsURLsMapOutput {
	return o
}

func (o AuthSettingsURLsMapOutput) MapIndex(k pulumi.StringInput) AuthSettingsURLsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthSettingsURLs {
		return vs[0].(map[string]*AuthSettingsURLs)[vs[1].(string)]
	}).(AuthSettingsURLsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthSettingsURLsInput)(nil)).Elem(), &AuthSettingsURLs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthSettingsURLsArrayInput)(nil)).Elem(), AuthSettingsURLsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthSettingsURLsMapInput)(nil)).Elem(), AuthSettingsURLsMap{})
	pulumi.RegisterOutputType(AuthSettingsURLsOutput{})
	pulumi.RegisterOutputType(AuthSettingsURLsArrayOutput{})
	pulumi.RegisterOutputType(AuthSettingsURLsMapOutput{})
}
