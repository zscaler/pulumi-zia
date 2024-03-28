# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 api_key: Optional[pulumi.Input[str]] = None,
                 api_token: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 zia_cloud: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        """
        if api_key is None:
            api_key = _utilities.get_env('ZIA_API_KEY')
        if api_key is not None:
            pulumi.set(__self__, "api_key", api_key)
        if api_token is None:
            api_token = _utilities.get_env('ZIA_SANDBOX_TOKEN')
        if api_token is not None:
            pulumi.set(__self__, "api_token", api_token)
        if password is None:
            password = _utilities.get_env('ZIA_PASSWORD')
        if password is not None:
            pulumi.set(__self__, "password", password)
        if username is None:
            username = _utilities.get_env('ZIA_USERNAME')
        if username is not None:
            pulumi.set(__self__, "username", username)
        if zia_cloud is None:
            zia_cloud = _utilities.get_env('ZIA_CLOUD')
        if zia_cloud is not None:
            pulumi.set(__self__, "zia_cloud", zia_cloud)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "api_key")

    @api_key.setter
    def api_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_key", value)

    @property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "api_token")

    @api_token.setter
    def api_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_token", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="ziaCloud")
    def zia_cloud(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "zia_cloud")

    @zia_cloud.setter
    def zia_cloud(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zia_cloud", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 api_token: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 zia_cloud: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the zia package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the zia package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 api_token: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 zia_cloud: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if api_key is None:
                api_key = _utilities.get_env('ZIA_API_KEY')
            __props__.__dict__["api_key"] = None if api_key is None else pulumi.Output.secret(api_key)
            if api_token is None:
                api_token = _utilities.get_env('ZIA_SANDBOX_TOKEN')
            __props__.__dict__["api_token"] = None if api_token is None else pulumi.Output.secret(api_token)
            if password is None:
                password = _utilities.get_env('ZIA_PASSWORD')
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            if username is None:
                username = _utilities.get_env('ZIA_USERNAME')
            __props__.__dict__["username"] = username
            if zia_cloud is None:
                zia_cloud = _utilities.get_env('ZIA_CLOUD')
            __props__.__dict__["zia_cloud"] = zia_cloud
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["apiKey", "apiToken", "password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Provider, __self__).__init__(
            'zia',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "api_token")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "username")

    @property
    @pulumi.getter(name="ziaCloud")
    def zia_cloud(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "zia_cloud")

