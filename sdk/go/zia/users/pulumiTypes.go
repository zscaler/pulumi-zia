// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package users

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type UserManagementDepartment struct {
	Comments *string `pulumi:"comments"`
	Deleted  *bool   `pulumi:"deleted"`
	// Department ID
	Id    *int `pulumi:"id"`
	IdpId *int `pulumi:"idpId"`
	// User name. This appears when choosing users for policies.
	Name *string `pulumi:"name"`
}

// UserManagementDepartmentInput is an input type that accepts UserManagementDepartmentArgs and UserManagementDepartmentOutput values.
// You can construct a concrete instance of `UserManagementDepartmentInput` via:
//
//	UserManagementDepartmentArgs{...}
type UserManagementDepartmentInput interface {
	pulumi.Input

	ToUserManagementDepartmentOutput() UserManagementDepartmentOutput
	ToUserManagementDepartmentOutputWithContext(context.Context) UserManagementDepartmentOutput
}

type UserManagementDepartmentArgs struct {
	Comments pulumi.StringPtrInput `pulumi:"comments"`
	Deleted  pulumi.BoolPtrInput   `pulumi:"deleted"`
	// Department ID
	Id    pulumi.IntPtrInput `pulumi:"id"`
	IdpId pulumi.IntPtrInput `pulumi:"idpId"`
	// User name. This appears when choosing users for policies.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (UserManagementDepartmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserManagementDepartment)(nil)).Elem()
}

func (i UserManagementDepartmentArgs) ToUserManagementDepartmentOutput() UserManagementDepartmentOutput {
	return i.ToUserManagementDepartmentOutputWithContext(context.Background())
}

func (i UserManagementDepartmentArgs) ToUserManagementDepartmentOutputWithContext(ctx context.Context) UserManagementDepartmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserManagementDepartmentOutput)
}

func (i UserManagementDepartmentArgs) ToUserManagementDepartmentPtrOutput() UserManagementDepartmentPtrOutput {
	return i.ToUserManagementDepartmentPtrOutputWithContext(context.Background())
}

func (i UserManagementDepartmentArgs) ToUserManagementDepartmentPtrOutputWithContext(ctx context.Context) UserManagementDepartmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserManagementDepartmentOutput).ToUserManagementDepartmentPtrOutputWithContext(ctx)
}

// UserManagementDepartmentPtrInput is an input type that accepts UserManagementDepartmentArgs, UserManagementDepartmentPtr and UserManagementDepartmentPtrOutput values.
// You can construct a concrete instance of `UserManagementDepartmentPtrInput` via:
//
//	        UserManagementDepartmentArgs{...}
//
//	or:
//
//	        nil
type UserManagementDepartmentPtrInput interface {
	pulumi.Input

	ToUserManagementDepartmentPtrOutput() UserManagementDepartmentPtrOutput
	ToUserManagementDepartmentPtrOutputWithContext(context.Context) UserManagementDepartmentPtrOutput
}

type userManagementDepartmentPtrType UserManagementDepartmentArgs

func UserManagementDepartmentPtr(v *UserManagementDepartmentArgs) UserManagementDepartmentPtrInput {
	return (*userManagementDepartmentPtrType)(v)
}

func (*userManagementDepartmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserManagementDepartment)(nil)).Elem()
}

func (i *userManagementDepartmentPtrType) ToUserManagementDepartmentPtrOutput() UserManagementDepartmentPtrOutput {
	return i.ToUserManagementDepartmentPtrOutputWithContext(context.Background())
}

func (i *userManagementDepartmentPtrType) ToUserManagementDepartmentPtrOutputWithContext(ctx context.Context) UserManagementDepartmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserManagementDepartmentPtrOutput)
}

type UserManagementDepartmentOutput struct{ *pulumi.OutputState }

func (UserManagementDepartmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserManagementDepartment)(nil)).Elem()
}

func (o UserManagementDepartmentOutput) ToUserManagementDepartmentOutput() UserManagementDepartmentOutput {
	return o
}

func (o UserManagementDepartmentOutput) ToUserManagementDepartmentOutputWithContext(ctx context.Context) UserManagementDepartmentOutput {
	return o
}

func (o UserManagementDepartmentOutput) ToUserManagementDepartmentPtrOutput() UserManagementDepartmentPtrOutput {
	return o.ToUserManagementDepartmentPtrOutputWithContext(context.Background())
}

func (o UserManagementDepartmentOutput) ToUserManagementDepartmentPtrOutputWithContext(ctx context.Context) UserManagementDepartmentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserManagementDepartment) *UserManagementDepartment {
		return &v
	}).(UserManagementDepartmentPtrOutput)
}

func (o UserManagementDepartmentOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserManagementDepartment) *string { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o UserManagementDepartmentOutput) Deleted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UserManagementDepartment) *bool { return v.Deleted }).(pulumi.BoolPtrOutput)
}

// Department ID
func (o UserManagementDepartmentOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v UserManagementDepartment) *int { return v.Id }).(pulumi.IntPtrOutput)
}

func (o UserManagementDepartmentOutput) IdpId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v UserManagementDepartment) *int { return v.IdpId }).(pulumi.IntPtrOutput)
}

// User name. This appears when choosing users for policies.
func (o UserManagementDepartmentOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserManagementDepartment) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type UserManagementDepartmentPtrOutput struct{ *pulumi.OutputState }

func (UserManagementDepartmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserManagementDepartment)(nil)).Elem()
}

func (o UserManagementDepartmentPtrOutput) ToUserManagementDepartmentPtrOutput() UserManagementDepartmentPtrOutput {
	return o
}

func (o UserManagementDepartmentPtrOutput) ToUserManagementDepartmentPtrOutputWithContext(ctx context.Context) UserManagementDepartmentPtrOutput {
	return o
}

func (o UserManagementDepartmentPtrOutput) Elem() UserManagementDepartmentOutput {
	return o.ApplyT(func(v *UserManagementDepartment) UserManagementDepartment {
		if v != nil {
			return *v
		}
		var ret UserManagementDepartment
		return ret
	}).(UserManagementDepartmentOutput)
}

func (o UserManagementDepartmentPtrOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserManagementDepartment) *string {
		if v == nil {
			return nil
		}
		return v.Comments
	}).(pulumi.StringPtrOutput)
}

func (o UserManagementDepartmentPtrOutput) Deleted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserManagementDepartment) *bool {
		if v == nil {
			return nil
		}
		return v.Deleted
	}).(pulumi.BoolPtrOutput)
}

// Department ID
func (o UserManagementDepartmentPtrOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *UserManagementDepartment) *int {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.IntPtrOutput)
}

func (o UserManagementDepartmentPtrOutput) IdpId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *UserManagementDepartment) *int {
		if v == nil {
			return nil
		}
		return v.IdpId
	}).(pulumi.IntPtrOutput)
}

// User name. This appears when choosing users for policies.
func (o UserManagementDepartmentPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserManagementDepartment) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type UserManagementGroups struct {
	// Department ID
	Ids []int `pulumi:"ids"`
}

// UserManagementGroupsInput is an input type that accepts UserManagementGroupsArgs and UserManagementGroupsOutput values.
// You can construct a concrete instance of `UserManagementGroupsInput` via:
//
//	UserManagementGroupsArgs{...}
type UserManagementGroupsInput interface {
	pulumi.Input

	ToUserManagementGroupsOutput() UserManagementGroupsOutput
	ToUserManagementGroupsOutputWithContext(context.Context) UserManagementGroupsOutput
}

type UserManagementGroupsArgs struct {
	// Department ID
	Ids pulumi.IntArrayInput `pulumi:"ids"`
}

func (UserManagementGroupsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserManagementGroups)(nil)).Elem()
}

func (i UserManagementGroupsArgs) ToUserManagementGroupsOutput() UserManagementGroupsOutput {
	return i.ToUserManagementGroupsOutputWithContext(context.Background())
}

func (i UserManagementGroupsArgs) ToUserManagementGroupsOutputWithContext(ctx context.Context) UserManagementGroupsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserManagementGroupsOutput)
}

func (i UserManagementGroupsArgs) ToUserManagementGroupsPtrOutput() UserManagementGroupsPtrOutput {
	return i.ToUserManagementGroupsPtrOutputWithContext(context.Background())
}

func (i UserManagementGroupsArgs) ToUserManagementGroupsPtrOutputWithContext(ctx context.Context) UserManagementGroupsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserManagementGroupsOutput).ToUserManagementGroupsPtrOutputWithContext(ctx)
}

// UserManagementGroupsPtrInput is an input type that accepts UserManagementGroupsArgs, UserManagementGroupsPtr and UserManagementGroupsPtrOutput values.
// You can construct a concrete instance of `UserManagementGroupsPtrInput` via:
//
//	        UserManagementGroupsArgs{...}
//
//	or:
//
//	        nil
type UserManagementGroupsPtrInput interface {
	pulumi.Input

	ToUserManagementGroupsPtrOutput() UserManagementGroupsPtrOutput
	ToUserManagementGroupsPtrOutputWithContext(context.Context) UserManagementGroupsPtrOutput
}

type userManagementGroupsPtrType UserManagementGroupsArgs

func UserManagementGroupsPtr(v *UserManagementGroupsArgs) UserManagementGroupsPtrInput {
	return (*userManagementGroupsPtrType)(v)
}

func (*userManagementGroupsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserManagementGroups)(nil)).Elem()
}

func (i *userManagementGroupsPtrType) ToUserManagementGroupsPtrOutput() UserManagementGroupsPtrOutput {
	return i.ToUserManagementGroupsPtrOutputWithContext(context.Background())
}

func (i *userManagementGroupsPtrType) ToUserManagementGroupsPtrOutputWithContext(ctx context.Context) UserManagementGroupsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserManagementGroupsPtrOutput)
}

type UserManagementGroupsOutput struct{ *pulumi.OutputState }

func (UserManagementGroupsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserManagementGroups)(nil)).Elem()
}

func (o UserManagementGroupsOutput) ToUserManagementGroupsOutput() UserManagementGroupsOutput {
	return o
}

func (o UserManagementGroupsOutput) ToUserManagementGroupsOutputWithContext(ctx context.Context) UserManagementGroupsOutput {
	return o
}

func (o UserManagementGroupsOutput) ToUserManagementGroupsPtrOutput() UserManagementGroupsPtrOutput {
	return o.ToUserManagementGroupsPtrOutputWithContext(context.Background())
}

func (o UserManagementGroupsOutput) ToUserManagementGroupsPtrOutputWithContext(ctx context.Context) UserManagementGroupsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserManagementGroups) *UserManagementGroups {
		return &v
	}).(UserManagementGroupsPtrOutput)
}

// Department ID
func (o UserManagementGroupsOutput) Ids() pulumi.IntArrayOutput {
	return o.ApplyT(func(v UserManagementGroups) []int { return v.Ids }).(pulumi.IntArrayOutput)
}

type UserManagementGroupsPtrOutput struct{ *pulumi.OutputState }

func (UserManagementGroupsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserManagementGroups)(nil)).Elem()
}

func (o UserManagementGroupsPtrOutput) ToUserManagementGroupsPtrOutput() UserManagementGroupsPtrOutput {
	return o
}

func (o UserManagementGroupsPtrOutput) ToUserManagementGroupsPtrOutputWithContext(ctx context.Context) UserManagementGroupsPtrOutput {
	return o
}

func (o UserManagementGroupsPtrOutput) Elem() UserManagementGroupsOutput {
	return o.ApplyT(func(v *UserManagementGroups) UserManagementGroups {
		if v != nil {
			return *v
		}
		var ret UserManagementGroups
		return ret
	}).(UserManagementGroupsOutput)
}

// Department ID
func (o UserManagementGroupsPtrOutput) Ids() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *UserManagementGroups) []int {
		if v == nil {
			return nil
		}
		return v.Ids
	}).(pulumi.IntArrayOutput)
}

type GetUserManagementDepartment struct {
	// (String) Additional information about the group
	Comments string `pulumi:"comments"`
	// (Boolean) default: `false`
	Deleted bool `pulumi:"deleted"`
	// The ID of the time window resource.
	Id int `pulumi:"id"`
	// (Number) Unique identfier for the identity provider (IdP)
	IdpId int `pulumi:"idpId"`
	// User name. This appears when choosing users for policies.
	Name string `pulumi:"name"`
}

// GetUserManagementDepartmentInput is an input type that accepts GetUserManagementDepartmentArgs and GetUserManagementDepartmentOutput values.
// You can construct a concrete instance of `GetUserManagementDepartmentInput` via:
//
//	GetUserManagementDepartmentArgs{...}
type GetUserManagementDepartmentInput interface {
	pulumi.Input

	ToGetUserManagementDepartmentOutput() GetUserManagementDepartmentOutput
	ToGetUserManagementDepartmentOutputWithContext(context.Context) GetUserManagementDepartmentOutput
}

type GetUserManagementDepartmentArgs struct {
	// (String) Additional information about the group
	Comments pulumi.StringInput `pulumi:"comments"`
	// (Boolean) default: `false`
	Deleted pulumi.BoolInput `pulumi:"deleted"`
	// The ID of the time window resource.
	Id pulumi.IntInput `pulumi:"id"`
	// (Number) Unique identfier for the identity provider (IdP)
	IdpId pulumi.IntInput `pulumi:"idpId"`
	// User name. This appears when choosing users for policies.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetUserManagementDepartmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserManagementDepartment)(nil)).Elem()
}

func (i GetUserManagementDepartmentArgs) ToGetUserManagementDepartmentOutput() GetUserManagementDepartmentOutput {
	return i.ToGetUserManagementDepartmentOutputWithContext(context.Background())
}

func (i GetUserManagementDepartmentArgs) ToGetUserManagementDepartmentOutputWithContext(ctx context.Context) GetUserManagementDepartmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetUserManagementDepartmentOutput)
}

// GetUserManagementDepartmentArrayInput is an input type that accepts GetUserManagementDepartmentArray and GetUserManagementDepartmentArrayOutput values.
// You can construct a concrete instance of `GetUserManagementDepartmentArrayInput` via:
//
//	GetUserManagementDepartmentArray{ GetUserManagementDepartmentArgs{...} }
type GetUserManagementDepartmentArrayInput interface {
	pulumi.Input

	ToGetUserManagementDepartmentArrayOutput() GetUserManagementDepartmentArrayOutput
	ToGetUserManagementDepartmentArrayOutputWithContext(context.Context) GetUserManagementDepartmentArrayOutput
}

type GetUserManagementDepartmentArray []GetUserManagementDepartmentInput

func (GetUserManagementDepartmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetUserManagementDepartment)(nil)).Elem()
}

func (i GetUserManagementDepartmentArray) ToGetUserManagementDepartmentArrayOutput() GetUserManagementDepartmentArrayOutput {
	return i.ToGetUserManagementDepartmentArrayOutputWithContext(context.Background())
}

func (i GetUserManagementDepartmentArray) ToGetUserManagementDepartmentArrayOutputWithContext(ctx context.Context) GetUserManagementDepartmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetUserManagementDepartmentArrayOutput)
}

type GetUserManagementDepartmentOutput struct{ *pulumi.OutputState }

func (GetUserManagementDepartmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserManagementDepartment)(nil)).Elem()
}

func (o GetUserManagementDepartmentOutput) ToGetUserManagementDepartmentOutput() GetUserManagementDepartmentOutput {
	return o
}

func (o GetUserManagementDepartmentOutput) ToGetUserManagementDepartmentOutputWithContext(ctx context.Context) GetUserManagementDepartmentOutput {
	return o
}

// (String) Additional information about the group
func (o GetUserManagementDepartmentOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserManagementDepartment) string { return v.Comments }).(pulumi.StringOutput)
}

// (Boolean) default: `false`
func (o GetUserManagementDepartmentOutput) Deleted() pulumi.BoolOutput {
	return o.ApplyT(func(v GetUserManagementDepartment) bool { return v.Deleted }).(pulumi.BoolOutput)
}

// The ID of the time window resource.
func (o GetUserManagementDepartmentOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetUserManagementDepartment) int { return v.Id }).(pulumi.IntOutput)
}

// (Number) Unique identfier for the identity provider (IdP)
func (o GetUserManagementDepartmentOutput) IdpId() pulumi.IntOutput {
	return o.ApplyT(func(v GetUserManagementDepartment) int { return v.IdpId }).(pulumi.IntOutput)
}

// User name. This appears when choosing users for policies.
func (o GetUserManagementDepartmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserManagementDepartment) string { return v.Name }).(pulumi.StringOutput)
}

type GetUserManagementDepartmentArrayOutput struct{ *pulumi.OutputState }

func (GetUserManagementDepartmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetUserManagementDepartment)(nil)).Elem()
}

func (o GetUserManagementDepartmentArrayOutput) ToGetUserManagementDepartmentArrayOutput() GetUserManagementDepartmentArrayOutput {
	return o
}

func (o GetUserManagementDepartmentArrayOutput) ToGetUserManagementDepartmentArrayOutputWithContext(ctx context.Context) GetUserManagementDepartmentArrayOutput {
	return o
}

func (o GetUserManagementDepartmentArrayOutput) Index(i pulumi.IntInput) GetUserManagementDepartmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetUserManagementDepartment {
		return vs[0].([]GetUserManagementDepartment)[vs[1].(int)]
	}).(GetUserManagementDepartmentOutput)
}

type GetUserManagementGroup struct {
	// (String) Additional information about the group
	Comments string `pulumi:"comments"`
	// The ID of the time window resource.
	Id int `pulumi:"id"`
	// (Number) Unique identfier for the identity provider (IdP)
	IdpId int `pulumi:"idpId"`
	// User name. This appears when choosing users for policies.
	Name string `pulumi:"name"`
}

// GetUserManagementGroupInput is an input type that accepts GetUserManagementGroupArgs and GetUserManagementGroupOutput values.
// You can construct a concrete instance of `GetUserManagementGroupInput` via:
//
//	GetUserManagementGroupArgs{...}
type GetUserManagementGroupInput interface {
	pulumi.Input

	ToGetUserManagementGroupOutput() GetUserManagementGroupOutput
	ToGetUserManagementGroupOutputWithContext(context.Context) GetUserManagementGroupOutput
}

type GetUserManagementGroupArgs struct {
	// (String) Additional information about the group
	Comments pulumi.StringInput `pulumi:"comments"`
	// The ID of the time window resource.
	Id pulumi.IntInput `pulumi:"id"`
	// (Number) Unique identfier for the identity provider (IdP)
	IdpId pulumi.IntInput `pulumi:"idpId"`
	// User name. This appears when choosing users for policies.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetUserManagementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserManagementGroup)(nil)).Elem()
}

func (i GetUserManagementGroupArgs) ToGetUserManagementGroupOutput() GetUserManagementGroupOutput {
	return i.ToGetUserManagementGroupOutputWithContext(context.Background())
}

func (i GetUserManagementGroupArgs) ToGetUserManagementGroupOutputWithContext(ctx context.Context) GetUserManagementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetUserManagementGroupOutput)
}

// GetUserManagementGroupArrayInput is an input type that accepts GetUserManagementGroupArray and GetUserManagementGroupArrayOutput values.
// You can construct a concrete instance of `GetUserManagementGroupArrayInput` via:
//
//	GetUserManagementGroupArray{ GetUserManagementGroupArgs{...} }
type GetUserManagementGroupArrayInput interface {
	pulumi.Input

	ToGetUserManagementGroupArrayOutput() GetUserManagementGroupArrayOutput
	ToGetUserManagementGroupArrayOutputWithContext(context.Context) GetUserManagementGroupArrayOutput
}

type GetUserManagementGroupArray []GetUserManagementGroupInput

func (GetUserManagementGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetUserManagementGroup)(nil)).Elem()
}

func (i GetUserManagementGroupArray) ToGetUserManagementGroupArrayOutput() GetUserManagementGroupArrayOutput {
	return i.ToGetUserManagementGroupArrayOutputWithContext(context.Background())
}

func (i GetUserManagementGroupArray) ToGetUserManagementGroupArrayOutputWithContext(ctx context.Context) GetUserManagementGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetUserManagementGroupArrayOutput)
}

type GetUserManagementGroupOutput struct{ *pulumi.OutputState }

func (GetUserManagementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserManagementGroup)(nil)).Elem()
}

func (o GetUserManagementGroupOutput) ToGetUserManagementGroupOutput() GetUserManagementGroupOutput {
	return o
}

func (o GetUserManagementGroupOutput) ToGetUserManagementGroupOutputWithContext(ctx context.Context) GetUserManagementGroupOutput {
	return o
}

// (String) Additional information about the group
func (o GetUserManagementGroupOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserManagementGroup) string { return v.Comments }).(pulumi.StringOutput)
}

// The ID of the time window resource.
func (o GetUserManagementGroupOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetUserManagementGroup) int { return v.Id }).(pulumi.IntOutput)
}

// (Number) Unique identfier for the identity provider (IdP)
func (o GetUserManagementGroupOutput) IdpId() pulumi.IntOutput {
	return o.ApplyT(func(v GetUserManagementGroup) int { return v.IdpId }).(pulumi.IntOutput)
}

// User name. This appears when choosing users for policies.
func (o GetUserManagementGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserManagementGroup) string { return v.Name }).(pulumi.StringOutput)
}

type GetUserManagementGroupArrayOutput struct{ *pulumi.OutputState }

func (GetUserManagementGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetUserManagementGroup)(nil)).Elem()
}

func (o GetUserManagementGroupArrayOutput) ToGetUserManagementGroupArrayOutput() GetUserManagementGroupArrayOutput {
	return o
}

func (o GetUserManagementGroupArrayOutput) ToGetUserManagementGroupArrayOutputWithContext(ctx context.Context) GetUserManagementGroupArrayOutput {
	return o
}

func (o GetUserManagementGroupArrayOutput) Index(i pulumi.IntInput) GetUserManagementGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetUserManagementGroup {
		return vs[0].([]GetUserManagementGroup)[vs[1].(int)]
	}).(GetUserManagementGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserManagementDepartmentInput)(nil)).Elem(), UserManagementDepartmentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserManagementDepartmentPtrInput)(nil)).Elem(), UserManagementDepartmentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserManagementGroupsInput)(nil)).Elem(), UserManagementGroupsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserManagementGroupsPtrInput)(nil)).Elem(), UserManagementGroupsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetUserManagementDepartmentInput)(nil)).Elem(), GetUserManagementDepartmentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetUserManagementDepartmentArrayInput)(nil)).Elem(), GetUserManagementDepartmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetUserManagementGroupInput)(nil)).Elem(), GetUserManagementGroupArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetUserManagementGroupArrayInput)(nil)).Elem(), GetUserManagementGroupArray{})
	pulumi.RegisterOutputType(UserManagementDepartmentOutput{})
	pulumi.RegisterOutputType(UserManagementDepartmentPtrOutput{})
	pulumi.RegisterOutputType(UserManagementGroupsOutput{})
	pulumi.RegisterOutputType(UserManagementGroupsPtrOutput{})
	pulumi.RegisterOutputType(GetUserManagementDepartmentOutput{})
	pulumi.RegisterOutputType(GetUserManagementDepartmentArrayOutput{})
	pulumi.RegisterOutputType(GetUserManagementGroupOutput{})
	pulumi.RegisterOutputType(GetUserManagementGroupArrayOutput{})
}
