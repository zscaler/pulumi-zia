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
    'GetFirewallFilteringAppServicesResult',
    'AwaitableGetFirewallFilteringAppServicesResult',
    'get_firewall_filtering_app_services',
    'get_firewall_filtering_app_services_output',
]

@pulumi.output_type
class GetFirewallFilteringAppServicesResult:
    """
    A collection of values returned by getFirewallFilteringAppServices.
    """
    def __init__(__self__, id=None, name=None, name_l10n_tag=None):
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_l10n_tag and not isinstance(name_l10n_tag, bool):
            raise TypeError("Expected argument 'name_l10n_tag' to be a bool")
        pulumi.set(__self__, "name_l10n_tag", name_l10n_tag)

    @property
    @pulumi.getter
    def id(self) -> int:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameL10nTag")
    def name_l10n_tag(self) -> bool:
        return pulumi.get(self, "name_l10n_tag")


class AwaitableGetFirewallFilteringAppServicesResult(GetFirewallFilteringAppServicesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallFilteringAppServicesResult(
            id=self.id,
            name=self.name,
            name_l10n_tag=self.name_l10n_tag)


def get_firewall_filtering_app_services(name: Optional[str] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallFilteringAppServicesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zia:index/getFirewallFilteringAppServices:getFirewallFilteringAppServices', __args__, opts=opts, typ=GetFirewallFilteringAppServicesResult).value

    return AwaitableGetFirewallFilteringAppServicesResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        name_l10n_tag=pulumi.get(__ret__, 'name_l10n_tag'))


@_utilities.lift_output_func(get_firewall_filtering_app_services)
def get_firewall_filtering_app_services_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFirewallFilteringAppServicesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
