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

__all__ = [
    'GetFirewallFilteringDestinationGroupsResult',
    'AwaitableGetFirewallFilteringDestinationGroupsResult',
    'get_firewall_filtering_destination_groups',
    'get_firewall_filtering_destination_groups_output',
]

@pulumi.output_type
class GetFirewallFilteringDestinationGroupsResult:
    """
    A collection of values returned by getFirewallFilteringDestinationGroups.
    """
    def __init__(__self__, addresses=None, countries=None, description=None, id=None, ip_categories=None, name=None, type=None):
        if addresses and not isinstance(addresses, list):
            raise TypeError("Expected argument 'addresses' to be a list")
        pulumi.set(__self__, "addresses", addresses)
        if countries and not isinstance(countries, list):
            raise TypeError("Expected argument 'countries' to be a list")
        pulumi.set(__self__, "countries", countries)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)
        if ip_categories and not isinstance(ip_categories, list):
            raise TypeError("Expected argument 'ip_categories' to be a list")
        pulumi.set(__self__, "ip_categories", ip_categories)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def addresses(self) -> Sequence[builtins.str]:
        """
        (List of String) Destination IP addresses within the group
        """
        return pulumi.get(self, "addresses")

    @property
    @pulumi.getter
    def countries(self) -> Sequence[builtins.str]:
        """
        (List of String) Destination IP address counties. You can identify destinations based on the location of a server.
        """
        return pulumi.get(self, "countries")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        (String) Additional information about the destination IP group
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> builtins.int:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipCategories")
    def ip_categories(self) -> Sequence[builtins.str]:
        """
        (List of String) Destination IP address URL categories. You can identify destinations based on the URL category of the domain. See list of all IP Categories [Here](https://help.zscaler.com/zia/firewall-policies#/ipDestinationGroups-get)
        * !> **WARNING:** The `ip_categories` attribute only accepts custom URL categories.
        """
        return pulumi.get(self, "ip_categories")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        (String) Destination IP group type (i.e., the group can contain destination IP addresses or FQDNs)
        """
        return pulumi.get(self, "type")


class AwaitableGetFirewallFilteringDestinationGroupsResult(GetFirewallFilteringDestinationGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallFilteringDestinationGroupsResult(
            addresses=self.addresses,
            countries=self.countries,
            description=self.description,
            id=self.id,
            ip_categories=self.ip_categories,
            name=self.name,
            type=self.type)


def get_firewall_filtering_destination_groups(id: Optional[builtins.int] = None,
                                              name: Optional[builtins.str] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallFilteringDestinationGroupsResult:
    """
    * [Official documentation](https://help.zscaler.com/zia/firewall-policies#/ipDestinationGroups-post)
    * [API documentation](https://help.zscaler.com/zia/firewall-policies#/ipDestinationGroups-post)

    Use the **zia_firewall_filtering_destination_groups** data source to get information about IP destination groups option available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering rule.

    ## Example Usage


    :param builtins.int id: The ID of the destination group resource.
    :param builtins.str name: The name of the destination group to be exported.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zia:index/getFirewallFilteringDestinationGroups:getFirewallFilteringDestinationGroups', __args__, opts=opts, typ=GetFirewallFilteringDestinationGroupsResult).value

    return AwaitableGetFirewallFilteringDestinationGroupsResult(
        addresses=pulumi.get(__ret__, 'addresses'),
        countries=pulumi.get(__ret__, 'countries'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        ip_categories=pulumi.get(__ret__, 'ip_categories'),
        name=pulumi.get(__ret__, 'name'),
        type=pulumi.get(__ret__, 'type'))
def get_firewall_filtering_destination_groups_output(id: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                                                     name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFirewallFilteringDestinationGroupsResult]:
    """
    * [Official documentation](https://help.zscaler.com/zia/firewall-policies#/ipDestinationGroups-post)
    * [API documentation](https://help.zscaler.com/zia/firewall-policies#/ipDestinationGroups-post)

    Use the **zia_firewall_filtering_destination_groups** data source to get information about IP destination groups option available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering rule.

    ## Example Usage


    :param builtins.int id: The ID of the destination group resource.
    :param builtins.str name: The name of the destination group to be exported.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('zia:index/getFirewallFilteringDestinationGroups:getFirewallFilteringDestinationGroups', __args__, opts=opts, typ=GetFirewallFilteringDestinationGroupsResult)
    return __ret__.apply(lambda __response__: GetFirewallFilteringDestinationGroupsResult(
        addresses=pulumi.get(__response__, 'addresses'),
        countries=pulumi.get(__response__, 'countries'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        ip_categories=pulumi.get(__response__, 'ip_categories'),
        name=pulumi.get(__response__, 'name'),
        type=pulumi.get(__response__, 'type')))
