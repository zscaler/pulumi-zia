# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SecuritySettingsArgs', 'SecuritySettings']

@pulumi.input_type
class SecuritySettingsArgs:
    def __init__(__self__, *,
                 blacklist_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 whitelist_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SecuritySettings resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] blacklist_urls: URLs on the denylist for your organization. Allow up to 25000 URLs.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] whitelist_urls: Allowlist URLs whose contents will not be scanned. Allows up to 255 URLs.
        """
        if blacklist_urls is not None:
            pulumi.set(__self__, "blacklist_urls", blacklist_urls)
        if whitelist_urls is not None:
            pulumi.set(__self__, "whitelist_urls", whitelist_urls)

    @property
    @pulumi.getter(name="blacklistUrls")
    def blacklist_urls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        URLs on the denylist for your organization. Allow up to 25000 URLs.
        """
        return pulumi.get(self, "blacklist_urls")

    @blacklist_urls.setter
    def blacklist_urls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "blacklist_urls", value)

    @property
    @pulumi.getter(name="whitelistUrls")
    def whitelist_urls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Allowlist URLs whose contents will not be scanned. Allows up to 255 URLs.
        """
        return pulumi.get(self, "whitelist_urls")

    @whitelist_urls.setter
    def whitelist_urls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "whitelist_urls", value)


@pulumi.input_type
class _SecuritySettingsState:
    def __init__(__self__, *,
                 blacklist_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 whitelist_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering SecuritySettings resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] blacklist_urls: URLs on the denylist for your organization. Allow up to 25000 URLs.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] whitelist_urls: Allowlist URLs whose contents will not be scanned. Allows up to 255 URLs.
        """
        if blacklist_urls is not None:
            pulumi.set(__self__, "blacklist_urls", blacklist_urls)
        if whitelist_urls is not None:
            pulumi.set(__self__, "whitelist_urls", whitelist_urls)

    @property
    @pulumi.getter(name="blacklistUrls")
    def blacklist_urls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        URLs on the denylist for your organization. Allow up to 25000 URLs.
        """
        return pulumi.get(self, "blacklist_urls")

    @blacklist_urls.setter
    def blacklist_urls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "blacklist_urls", value)

    @property
    @pulumi.getter(name="whitelistUrls")
    def whitelist_urls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Allowlist URLs whose contents will not be scanned. Allows up to 255 URLs.
        """
        return pulumi.get(self, "whitelist_urls")

    @whitelist_urls.setter
    def whitelist_urls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "whitelist_urls", value)


class SecuritySettings(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blacklist_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 whitelist_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a SecuritySettings resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] blacklist_urls: URLs on the denylist for your organization. Allow up to 25000 URLs.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] whitelist_urls: Allowlist URLs whose contents will not be scanned. Allows up to 255 URLs.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SecuritySettingsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SecuritySettings resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SecuritySettingsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecuritySettingsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blacklist_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 whitelist_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecuritySettingsArgs.__new__(SecuritySettingsArgs)

            __props__.__dict__["blacklist_urls"] = blacklist_urls
            __props__.__dict__["whitelist_urls"] = whitelist_urls
        super(SecuritySettings, __self__).__init__(
            'zia:index/securitySettings:SecuritySettings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            blacklist_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            whitelist_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'SecuritySettings':
        """
        Get an existing SecuritySettings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] blacklist_urls: URLs on the denylist for your organization. Allow up to 25000 URLs.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] whitelist_urls: Allowlist URLs whose contents will not be scanned. Allows up to 255 URLs.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecuritySettingsState.__new__(_SecuritySettingsState)

        __props__.__dict__["blacklist_urls"] = blacklist_urls
        __props__.__dict__["whitelist_urls"] = whitelist_urls
        return SecuritySettings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="blacklistUrls")
    def blacklist_urls(self) -> pulumi.Output[Sequence[str]]:
        """
        URLs on the denylist for your organization. Allow up to 25000 URLs.
        """
        return pulumi.get(self, "blacklist_urls")

    @property
    @pulumi.getter(name="whitelistUrls")
    def whitelist_urls(self) -> pulumi.Output[Sequence[str]]:
        """
        Allowlist URLs whose contents will not be scanned. Allows up to 255 URLs.
        """
        return pulumi.get(self, "whitelist_urls")

