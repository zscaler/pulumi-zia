# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetFirewallFilteringNetworkServiceGroupsResult',
    'AwaitableGetFirewallFilteringNetworkServiceGroupsResult',
    'get_firewall_filtering_network_service_groups',
    'get_firewall_filtering_network_service_groups_output',
]

@pulumi.output_type
class GetFirewallFilteringNetworkServiceGroupsResult:
    """
    A collection of values returned by getFirewallFilteringNetworkServiceGroups.
    """
    def __init__(__self__, description=None, id=None, name=None, services=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if services and not isinstance(services, list):
            raise TypeError("Expected argument 'services' to be a list")
        pulumi.set(__self__, "services", services)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        (String)
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        (Number)
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        (String)
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def services(self) -> Sequence['outputs.GetFirewallFilteringNetworkServiceGroupsServiceResult']:
        """
        (Number) The ID of this resource.
        """
        return pulumi.get(self, "services")


class AwaitableGetFirewallFilteringNetworkServiceGroupsResult(GetFirewallFilteringNetworkServiceGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallFilteringNetworkServiceGroupsResult(
            description=self.description,
            id=self.id,
            name=self.name,
            services=self.services)


def get_firewall_filtering_network_service_groups(id: Optional[int] = None,
                                                  name: Optional[str] = None,
                                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallFilteringNetworkServiceGroupsResult:
    """
    Use the **zia_firewall_filtering_network_service_groups** data source to get information about a network service groups available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering network service rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zia as zia

    example = zia.get_firewall_filtering_network_service_groups(name="Corporate Custom SSH TCP_10022")
    ```


    :param int id: The ID of the ip source group to be exported.
    :param str name: The name of the ip source group to be exported.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zia:index/getFirewallFilteringNetworkServiceGroups:getFirewallFilteringNetworkServiceGroups', __args__, opts=opts, typ=GetFirewallFilteringNetworkServiceGroupsResult).value

    return AwaitableGetFirewallFilteringNetworkServiceGroupsResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        services=pulumi.get(__ret__, 'services'))


@_utilities.lift_output_func(get_firewall_filtering_network_service_groups)
def get_firewall_filtering_network_service_groups_output(id: Optional[pulumi.Input[Optional[int]]] = None,
                                                         name: Optional[pulumi.Input[Optional[str]]] = None,
                                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFirewallFilteringNetworkServiceGroupsResult]:
    """
    Use the **zia_firewall_filtering_network_service_groups** data source to get information about a network service groups available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering network service rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zia as zia

    example = zia.get_firewall_filtering_network_service_groups(name="Corporate Custom SSH TCP_10022")
    ```


    :param int id: The ID of the ip source group to be exported.
    :param str name: The name of the ip source group to be exported.
    """
    ...
