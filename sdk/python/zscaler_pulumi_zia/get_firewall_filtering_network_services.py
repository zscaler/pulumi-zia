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
    'GetFirewallFilteringNetworkServicesResult',
    'AwaitableGetFirewallFilteringNetworkServicesResult',
    'get_firewall_filtering_network_services',
    'get_firewall_filtering_network_services_output',
]

@pulumi.output_type
class GetFirewallFilteringNetworkServicesResult:
    """
    A collection of values returned by getFirewallFilteringNetworkServices.
    """
    def __init__(__self__, description=None, dest_tcp_ports=None, dest_udp_ports=None, id=None, is_name_l10n_tag=None, name=None, src_tcp_ports=None, src_udp_ports=None, tag=None, type=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dest_tcp_ports and not isinstance(dest_tcp_ports, list):
            raise TypeError("Expected argument 'dest_tcp_ports' to be a list")
        pulumi.set(__self__, "dest_tcp_ports", dest_tcp_ports)
        if dest_udp_ports and not isinstance(dest_udp_ports, list):
            raise TypeError("Expected argument 'dest_udp_ports' to be a list")
        pulumi.set(__self__, "dest_udp_ports", dest_udp_ports)
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)
        if is_name_l10n_tag and not isinstance(is_name_l10n_tag, bool):
            raise TypeError("Expected argument 'is_name_l10n_tag' to be a bool")
        pulumi.set(__self__, "is_name_l10n_tag", is_name_l10n_tag)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if src_tcp_ports and not isinstance(src_tcp_ports, list):
            raise TypeError("Expected argument 'src_tcp_ports' to be a list")
        pulumi.set(__self__, "src_tcp_ports", src_tcp_ports)
        if src_udp_ports and not isinstance(src_udp_ports, list):
            raise TypeError("Expected argument 'src_udp_ports' to be a list")
        pulumi.set(__self__, "src_udp_ports", src_udp_ports)
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        pulumi.set(__self__, "tag", tag)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        (String) (Optional) Enter additional notes or information. The description cannot exceed 10240 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="destTcpPorts")
    def dest_tcp_ports(self) -> Sequence['outputs.GetFirewallFilteringNetworkServicesDestTcpPortResult']:
        """
        (Required) The TCP destination port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        """
        return pulumi.get(self, "dest_tcp_ports")

    @property
    @pulumi.getter(name="destUdpPorts")
    def dest_udp_ports(self) -> Sequence['outputs.GetFirewallFilteringNetworkServicesDestUdpPortResult']:
        """
        The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        """
        return pulumi.get(self, "dest_udp_ports")

    @property
    @pulumi.getter
    def id(self) -> int:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isNameL10nTag")
    def is_name_l10n_tag(self) -> bool:
        """
        (Bool) - Default: false
        """
        return pulumi.get(self, "is_name_l10n_tag")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="srcTcpPorts")
    def src_tcp_ports(self) -> Sequence['outputs.GetFirewallFilteringNetworkServicesSrcTcpPortResult']:
        """
        (Optional) The TCP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service
        """
        return pulumi.get(self, "src_tcp_ports")

    @property
    @pulumi.getter(name="srcUdpPorts")
    def src_udp_ports(self) -> Sequence['outputs.GetFirewallFilteringNetworkServicesSrcUdpPortResult']:
        """
        The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        """
        return pulumi.get(self, "src_udp_ports")

    @property
    @pulumi.getter
    def tag(self) -> str:
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        (String) - Supported values are: `STANDARD`, `PREDEFINED` and `CUSTOM`
        """
        return pulumi.get(self, "type")


class AwaitableGetFirewallFilteringNetworkServicesResult(GetFirewallFilteringNetworkServicesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallFilteringNetworkServicesResult(
            description=self.description,
            dest_tcp_ports=self.dest_tcp_ports,
            dest_udp_ports=self.dest_udp_ports,
            id=self.id,
            is_name_l10n_tag=self.is_name_l10n_tag,
            name=self.name,
            src_tcp_ports=self.src_tcp_ports,
            src_udp_ports=self.src_udp_ports,
            tag=self.tag,
            type=self.type)


def get_firewall_filtering_network_services(id: Optional[int] = None,
                                            name: Optional[str] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallFilteringNetworkServicesResult:
    """
    The **zia_firewall_filtering_network_service** data source to get information about a network service available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering network service rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zia as zia

    example = zia.get_firewall_filtering_network_services(name="ICMP_ANY")
    ```


    :param int id: The ID of the application layer service to be exported.
    :param str name: Name of the application layer service that you want to control. It can include any character and spaces.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zia:index/getFirewallFilteringNetworkServices:getFirewallFilteringNetworkServices', __args__, opts=opts, typ=GetFirewallFilteringNetworkServicesResult).value

    return AwaitableGetFirewallFilteringNetworkServicesResult(
        description=pulumi.get(__ret__, 'description'),
        dest_tcp_ports=pulumi.get(__ret__, 'dest_tcp_ports'),
        dest_udp_ports=pulumi.get(__ret__, 'dest_udp_ports'),
        id=pulumi.get(__ret__, 'id'),
        is_name_l10n_tag=pulumi.get(__ret__, 'is_name_l10n_tag'),
        name=pulumi.get(__ret__, 'name'),
        src_tcp_ports=pulumi.get(__ret__, 'src_tcp_ports'),
        src_udp_ports=pulumi.get(__ret__, 'src_udp_ports'),
        tag=pulumi.get(__ret__, 'tag'),
        type=pulumi.get(__ret__, 'type'))


@_utilities.lift_output_func(get_firewall_filtering_network_services)
def get_firewall_filtering_network_services_output(id: Optional[pulumi.Input[Optional[int]]] = None,
                                                   name: Optional[pulumi.Input[Optional[str]]] = None,
                                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFirewallFilteringNetworkServicesResult]:
    """
    The **zia_firewall_filtering_network_service** data source to get information about a network service available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering network service rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zia as zia

    example = zia.get_firewall_filtering_network_services(name="ICMP_ANY")
    ```


    :param int id: The ID of the application layer service to be exported.
    :param str name: Name of the application layer service that you want to control. It can include any character and spaces.
    """
    ...
