# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSecuritySettingsResult',
    'AwaitableGetSecuritySettingsResult',
    'get_security_settings',
    'get_security_settings_output',
]

@pulumi.output_type
class GetSecuritySettingsResult:
    """
    A collection of values returned by getSecuritySettings.
    """
    def __init__(__self__, blacklist_urls=None, id=None, whitelist_urls=None):
        if blacklist_urls and not isinstance(blacklist_urls, list):
            raise TypeError("Expected argument 'blacklist_urls' to be a list")
        pulumi.set(__self__, "blacklist_urls", blacklist_urls)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if whitelist_urls and not isinstance(whitelist_urls, list):
            raise TypeError("Expected argument 'whitelist_urls' to be a list")
        pulumi.set(__self__, "whitelist_urls", whitelist_urls)

    @property
    @pulumi.getter(name="blacklistUrls")
    def blacklist_urls(self) -> Sequence[str]:
        return pulumi.get(self, "blacklist_urls")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="whitelistUrls")
    def whitelist_urls(self) -> Sequence[str]:
        return pulumi.get(self, "whitelist_urls")


class AwaitableGetSecuritySettingsResult(GetSecuritySettingsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecuritySettingsResult(
            blacklist_urls=self.blacklist_urls,
            id=self.id,
            whitelist_urls=self.whitelist_urls)


def get_security_settings(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecuritySettingsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zia:index/getSecuritySettings:getSecuritySettings', __args__, opts=opts, typ=GetSecuritySettingsResult).value

    return AwaitableGetSecuritySettingsResult(
        blacklist_urls=pulumi.get(__ret__, 'blacklist_urls'),
        id=pulumi.get(__ret__, 'id'),
        whitelist_urls=pulumi.get(__ret__, 'whitelist_urls'))


@_utilities.lift_output_func(get_security_settings)
def get_security_settings_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecuritySettingsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
