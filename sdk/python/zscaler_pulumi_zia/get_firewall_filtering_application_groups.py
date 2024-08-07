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
    'GetFirewallFilteringApplicationGroupsResult',
    'AwaitableGetFirewallFilteringApplicationGroupsResult',
    'get_firewall_filtering_application_groups',
    'get_firewall_filtering_application_groups_output',
]

@pulumi.output_type
class GetFirewallFilteringApplicationGroupsResult:
    """
    A collection of values returned by getFirewallFilteringApplicationGroups.
    """
    def __init__(__self__, description=None, id=None, name=None, network_applications=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_applications and not isinstance(network_applications, list):
            raise TypeError("Expected argument 'network_applications' to be a list")
        pulumi.set(__self__, "network_applications", network_applications)

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
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkApplications")
    def network_applications(self) -> Sequence[str]:
        """
        (List of String)
        """
        return pulumi.get(self, "network_applications")


class AwaitableGetFirewallFilteringApplicationGroupsResult(GetFirewallFilteringApplicationGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallFilteringApplicationGroupsResult(
            description=self.description,
            id=self.id,
            name=self.name,
            network_applications=self.network_applications)


def get_firewall_filtering_application_groups(id: Optional[int] = None,
                                              name: Optional[str] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallFilteringApplicationGroupsResult:
    """
    Use the **zia_firewall_filtering_network_application_groups** data source to get information about a network application group available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering network application rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zia as zia

    example = zia.get_firewall_filtering_application_groups(name="example")
    ```


    :param int id: The ID of the ip source group resource.
    :param str name: The name of the ip source group to be exported.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zia:index/getFirewallFilteringApplicationGroups:getFirewallFilteringApplicationGroups', __args__, opts=opts, typ=GetFirewallFilteringApplicationGroupsResult).value

    return AwaitableGetFirewallFilteringApplicationGroupsResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        network_applications=pulumi.get(__ret__, 'network_applications'))


@_utilities.lift_output_func(get_firewall_filtering_application_groups)
def get_firewall_filtering_application_groups_output(id: Optional[pulumi.Input[Optional[int]]] = None,
                                                     name: Optional[pulumi.Input[Optional[str]]] = None,
                                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFirewallFilteringApplicationGroupsResult]:
    """
    Use the **zia_firewall_filtering_network_application_groups** data source to get information about a network application group available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering network application rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zia as zia

    example = zia.get_firewall_filtering_application_groups(name="example")
    ```


    :param int id: The ID of the ip source group resource.
    :param str name: The name of the ip source group to be exported.
    """
    ...
