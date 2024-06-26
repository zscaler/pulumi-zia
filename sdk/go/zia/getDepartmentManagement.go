// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zia

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zia/sdk/go/zia/internal"
)

func GetDepartmentManagement(ctx *pulumi.Context, args *GetDepartmentManagementArgs, opts ...pulumi.InvokeOption) (*GetDepartmentManagementResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDepartmentManagementResult
	err := ctx.Invoke("zia:index/getDepartmentManagement:getDepartmentManagement", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDepartmentManagement.
type GetDepartmentManagementArgs struct {
	Name *string `pulumi:"name"`
}

// A collection of values returned by getDepartmentManagement.
type GetDepartmentManagementResult struct {
	Comments string  `pulumi:"comments"`
	Deleted  bool    `pulumi:"deleted"`
	Id       int     `pulumi:"id"`
	IdpId    int     `pulumi:"idpId"`
	Name     *string `pulumi:"name"`
}

func GetDepartmentManagementOutput(ctx *pulumi.Context, args GetDepartmentManagementOutputArgs, opts ...pulumi.InvokeOption) GetDepartmentManagementResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDepartmentManagementResult, error) {
			args := v.(GetDepartmentManagementArgs)
			r, err := GetDepartmentManagement(ctx, &args, opts...)
			var s GetDepartmentManagementResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDepartmentManagementResultOutput)
}

// A collection of arguments for invoking getDepartmentManagement.
type GetDepartmentManagementOutputArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetDepartmentManagementOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDepartmentManagementArgs)(nil)).Elem()
}

// A collection of values returned by getDepartmentManagement.
type GetDepartmentManagementResultOutput struct{ *pulumi.OutputState }

func (GetDepartmentManagementResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDepartmentManagementResult)(nil)).Elem()
}

func (o GetDepartmentManagementResultOutput) ToGetDepartmentManagementResultOutput() GetDepartmentManagementResultOutput {
	return o
}

func (o GetDepartmentManagementResultOutput) ToGetDepartmentManagementResultOutputWithContext(ctx context.Context) GetDepartmentManagementResultOutput {
	return o
}

func (o GetDepartmentManagementResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v GetDepartmentManagementResult) string { return v.Comments }).(pulumi.StringOutput)
}

func (o GetDepartmentManagementResultOutput) Deleted() pulumi.BoolOutput {
	return o.ApplyT(func(v GetDepartmentManagementResult) bool { return v.Deleted }).(pulumi.BoolOutput)
}

func (o GetDepartmentManagementResultOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetDepartmentManagementResult) int { return v.Id }).(pulumi.IntOutput)
}

func (o GetDepartmentManagementResultOutput) IdpId() pulumi.IntOutput {
	return o.ApplyT(func(v GetDepartmentManagementResult) int { return v.IdpId }).(pulumi.IntOutput)
}

func (o GetDepartmentManagementResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDepartmentManagementResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDepartmentManagementResultOutput{})
}
