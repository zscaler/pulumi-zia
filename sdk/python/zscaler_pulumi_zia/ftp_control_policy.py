# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['FtpControlPolicyArgs', 'FtpControlPolicy']

@pulumi.input_type
class FtpControlPolicyArgs:
    def __init__(__self__, *,
                 ftp_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 ftp_over_http_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 url_categories: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a FtpControlPolicy resource.
        :param pulumi.Input[builtins.bool] ftp_enabled: Indicates whether to enable native FTP. When enabled, users can connect to native FTP sites and download files.
        :param pulumi.Input[builtins.bool] ftp_over_http_enabled: Indicates whether to enable FTP over HTTP.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] url_categories: List of URL categories for which rule must be applied
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] urls: Domains or URLs included for the FTP Control settings
        """
        if ftp_enabled is not None:
            pulumi.set(__self__, "ftp_enabled", ftp_enabled)
        if ftp_over_http_enabled is not None:
            pulumi.set(__self__, "ftp_over_http_enabled", ftp_over_http_enabled)
        if url_categories is not None:
            pulumi.set(__self__, "url_categories", url_categories)
        if urls is not None:
            pulumi.set(__self__, "urls", urls)

    @property
    @pulumi.getter(name="ftpEnabled")
    def ftp_enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Indicates whether to enable native FTP. When enabled, users can connect to native FTP sites and download files.
        """
        return pulumi.get(self, "ftp_enabled")

    @ftp_enabled.setter
    def ftp_enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "ftp_enabled", value)

    @property
    @pulumi.getter(name="ftpOverHttpEnabled")
    def ftp_over_http_enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Indicates whether to enable FTP over HTTP.
        """
        return pulumi.get(self, "ftp_over_http_enabled")

    @ftp_over_http_enabled.setter
    def ftp_over_http_enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "ftp_over_http_enabled", value)

    @property
    @pulumi.getter(name="urlCategories")
    def url_categories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of URL categories for which rule must be applied
        """
        return pulumi.get(self, "url_categories")

    @url_categories.setter
    def url_categories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "url_categories", value)

    @property
    @pulumi.getter
    def urls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Domains or URLs included for the FTP Control settings
        """
        return pulumi.get(self, "urls")

    @urls.setter
    def urls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "urls", value)


@pulumi.input_type
class _FtpControlPolicyState:
    def __init__(__self__, *,
                 ftp_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 ftp_over_http_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 url_categories: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering FtpControlPolicy resources.
        :param pulumi.Input[builtins.bool] ftp_enabled: Indicates whether to enable native FTP. When enabled, users can connect to native FTP sites and download files.
        :param pulumi.Input[builtins.bool] ftp_over_http_enabled: Indicates whether to enable FTP over HTTP.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] url_categories: List of URL categories for which rule must be applied
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] urls: Domains or URLs included for the FTP Control settings
        """
        if ftp_enabled is not None:
            pulumi.set(__self__, "ftp_enabled", ftp_enabled)
        if ftp_over_http_enabled is not None:
            pulumi.set(__self__, "ftp_over_http_enabled", ftp_over_http_enabled)
        if url_categories is not None:
            pulumi.set(__self__, "url_categories", url_categories)
        if urls is not None:
            pulumi.set(__self__, "urls", urls)

    @property
    @pulumi.getter(name="ftpEnabled")
    def ftp_enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Indicates whether to enable native FTP. When enabled, users can connect to native FTP sites and download files.
        """
        return pulumi.get(self, "ftp_enabled")

    @ftp_enabled.setter
    def ftp_enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "ftp_enabled", value)

    @property
    @pulumi.getter(name="ftpOverHttpEnabled")
    def ftp_over_http_enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Indicates whether to enable FTP over HTTP.
        """
        return pulumi.get(self, "ftp_over_http_enabled")

    @ftp_over_http_enabled.setter
    def ftp_over_http_enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "ftp_over_http_enabled", value)

    @property
    @pulumi.getter(name="urlCategories")
    def url_categories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of URL categories for which rule must be applied
        """
        return pulumi.get(self, "url_categories")

    @url_categories.setter
    def url_categories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "url_categories", value)

    @property
    @pulumi.getter
    def urls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Domains or URLs included for the FTP Control settings
        """
        return pulumi.get(self, "urls")

    @urls.setter
    def urls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "urls", value)


@pulumi.type_token("zia:index/ftpControlPolicy:FtpControlPolicy")
class FtpControlPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ftp_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 ftp_over_http_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 url_categories: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        * [Official documentation](https://help.zscaler.com/zia/about-ftp-control)
        * [API documentation](https://help.zscaler.com/zia/ftp-control-policy#/ftpSettings-get)

        The **zia_ftp_control_policy** resource allows you to update FTP control Policy. To learn more see [Configuring the FTP Control Policy](https://help.zscaler.com/zia/configuring-ftp-control-policy)

        ## Example Usage

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        **zia_ftp_control_policy** can be imported by using `ftp_control` as the import ID.

        For example:

        ```sh
        $ pulumi import zia:index/ftpControlPolicy:FtpControlPolicy this "ftp_control"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] ftp_enabled: Indicates whether to enable native FTP. When enabled, users can connect to native FTP sites and download files.
        :param pulumi.Input[builtins.bool] ftp_over_http_enabled: Indicates whether to enable FTP over HTTP.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] url_categories: List of URL categories for which rule must be applied
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] urls: Domains or URLs included for the FTP Control settings
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FtpControlPolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [Official documentation](https://help.zscaler.com/zia/about-ftp-control)
        * [API documentation](https://help.zscaler.com/zia/ftp-control-policy#/ftpSettings-get)

        The **zia_ftp_control_policy** resource allows you to update FTP control Policy. To learn more see [Configuring the FTP Control Policy](https://help.zscaler.com/zia/configuring-ftp-control-policy)

        ## Example Usage

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        **zia_ftp_control_policy** can be imported by using `ftp_control` as the import ID.

        For example:

        ```sh
        $ pulumi import zia:index/ftpControlPolicy:FtpControlPolicy this "ftp_control"
        ```

        :param str resource_name: The name of the resource.
        :param FtpControlPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FtpControlPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ftp_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 ftp_over_http_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 url_categories: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FtpControlPolicyArgs.__new__(FtpControlPolicyArgs)

            __props__.__dict__["ftp_enabled"] = ftp_enabled
            __props__.__dict__["ftp_over_http_enabled"] = ftp_over_http_enabled
            __props__.__dict__["url_categories"] = url_categories
            __props__.__dict__["urls"] = urls
        super(FtpControlPolicy, __self__).__init__(
            'zia:index/ftpControlPolicy:FtpControlPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ftp_enabled: Optional[pulumi.Input[builtins.bool]] = None,
            ftp_over_http_enabled: Optional[pulumi.Input[builtins.bool]] = None,
            url_categories: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            urls: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None) -> 'FtpControlPolicy':
        """
        Get an existing FtpControlPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] ftp_enabled: Indicates whether to enable native FTP. When enabled, users can connect to native FTP sites and download files.
        :param pulumi.Input[builtins.bool] ftp_over_http_enabled: Indicates whether to enable FTP over HTTP.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] url_categories: List of URL categories for which rule must be applied
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] urls: Domains or URLs included for the FTP Control settings
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FtpControlPolicyState.__new__(_FtpControlPolicyState)

        __props__.__dict__["ftp_enabled"] = ftp_enabled
        __props__.__dict__["ftp_over_http_enabled"] = ftp_over_http_enabled
        __props__.__dict__["url_categories"] = url_categories
        __props__.__dict__["urls"] = urls
        return FtpControlPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ftpEnabled")
    def ftp_enabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Indicates whether to enable native FTP. When enabled, users can connect to native FTP sites and download files.
        """
        return pulumi.get(self, "ftp_enabled")

    @property
    @pulumi.getter(name="ftpOverHttpEnabled")
    def ftp_over_http_enabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Indicates whether to enable FTP over HTTP.
        """
        return pulumi.get(self, "ftp_over_http_enabled")

    @property
    @pulumi.getter(name="urlCategories")
    def url_categories(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        List of URL categories for which rule must be applied
        """
        return pulumi.get(self, "url_categories")

    @property
    @pulumi.getter
    def urls(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        Domains or URLs included for the FTP Control settings
        """
        return pulumi.get(self, "urls")

