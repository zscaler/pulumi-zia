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
    'GetTrafficForwardingGRETunnelInfoResult',
    'AwaitableGetTrafficForwardingGRETunnelInfoResult',
    'get_traffic_forwarding_gre_tunnel_info',
    'get_traffic_forwarding_gre_tunnel_info_output',
]

@pulumi.output_type
class GetTrafficForwardingGRETunnelInfoResult:
    """
    A collection of values returned by getTrafficForwardingGRETunnelInfo.
    """
    def __init__(__self__, gre_enabled=None, gre_range_primary=None, gre_range_secondary=None, gre_tunnel_ip=None, id=None, ip_address=None, primary_gw=None, secondary_gw=None, tun_id=None):
        if gre_enabled and not isinstance(gre_enabled, bool):
            raise TypeError("Expected argument 'gre_enabled' to be a bool")
        pulumi.set(__self__, "gre_enabled", gre_enabled)
        if gre_range_primary and not isinstance(gre_range_primary, str):
            raise TypeError("Expected argument 'gre_range_primary' to be a str")
        pulumi.set(__self__, "gre_range_primary", gre_range_primary)
        if gre_range_secondary and not isinstance(gre_range_secondary, str):
            raise TypeError("Expected argument 'gre_range_secondary' to be a str")
        pulumi.set(__self__, "gre_range_secondary", gre_range_secondary)
        if gre_tunnel_ip and not isinstance(gre_tunnel_ip, str):
            raise TypeError("Expected argument 'gre_tunnel_ip' to be a str")
        pulumi.set(__self__, "gre_tunnel_ip", gre_tunnel_ip)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_address and not isinstance(ip_address, str):
            raise TypeError("Expected argument 'ip_address' to be a str")
        pulumi.set(__self__, "ip_address", ip_address)
        if primary_gw and not isinstance(primary_gw, str):
            raise TypeError("Expected argument 'primary_gw' to be a str")
        pulumi.set(__self__, "primary_gw", primary_gw)
        if secondary_gw and not isinstance(secondary_gw, str):
            raise TypeError("Expected argument 'secondary_gw' to be a str")
        pulumi.set(__self__, "secondary_gw", secondary_gw)
        if tun_id and not isinstance(tun_id, int):
            raise TypeError("Expected argument 'tun_id' to be a int")
        pulumi.set(__self__, "tun_id", tun_id)

    @property
    @pulumi.getter(name="greEnabled")
    def gre_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "gre_enabled")

    @property
    @pulumi.getter(name="greRangePrimary")
    def gre_range_primary(self) -> str:
        """
        (String)
        """
        return pulumi.get(self, "gre_range_primary")

    @property
    @pulumi.getter(name="greRangeSecondary")
    def gre_range_secondary(self) -> str:
        """
        (String)
        """
        return pulumi.get(self, "gre_range_secondary")

    @property
    @pulumi.getter(name="greTunnelIp")
    def gre_tunnel_ip(self) -> str:
        """
        (String) The start of the internal IP address in /29 CIDR range
        """
        return pulumi.get(self, "gre_tunnel_ip")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[str]:
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="primaryGw")
    def primary_gw(self) -> str:
        """
        (String)
        """
        return pulumi.get(self, "primary_gw")

    @property
    @pulumi.getter(name="secondaryGw")
    def secondary_gw(self) -> str:
        """
        (String)
        """
        return pulumi.get(self, "secondary_gw")

    @property
    @pulumi.getter(name="tunId")
    def tun_id(self) -> int:
        """
        (Number)
        """
        return pulumi.get(self, "tun_id")


class AwaitableGetTrafficForwardingGRETunnelInfoResult(GetTrafficForwardingGRETunnelInfoResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTrafficForwardingGRETunnelInfoResult(
            gre_enabled=self.gre_enabled,
            gre_range_primary=self.gre_range_primary,
            gre_range_secondary=self.gre_range_secondary,
            gre_tunnel_ip=self.gre_tunnel_ip,
            id=self.id,
            ip_address=self.ip_address,
            primary_gw=self.primary_gw,
            secondary_gw=self.secondary_gw,
            tun_id=self.tun_id)


def get_traffic_forwarding_gre_tunnel_info(gre_enabled: Optional[bool] = None,
                                           ip_address: Optional[str] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTrafficForwardingGRETunnelInfoResult:
    """
    The **zia_traffic_forwarding_gre_tunnel_info** data source to get information about provisioned GRE tunnel information created in the Zscaler Internet Access portal.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zia as zia

    example = zia.get_traffic_forwarding_gre_tunnel_info(ip_address="1.1.1.1")
    ```


    :param bool gre_enabled: Displays only ip addresses with GRE tunnel enabled
           
           > **NOTE** `ip_address` is the public IP address (Static IP) associated with the GRE Tunnel
    :param str ip_address: Filter based on an IP address range.
    """
    __args__ = dict()
    __args__['greEnabled'] = gre_enabled
    __args__['ipAddress'] = ip_address
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zia:index/getTrafficForwardingGRETunnelInfo:getTrafficForwardingGRETunnelInfo', __args__, opts=opts, typ=GetTrafficForwardingGRETunnelInfoResult).value

    return AwaitableGetTrafficForwardingGRETunnelInfoResult(
        gre_enabled=pulumi.get(__ret__, 'gre_enabled'),
        gre_range_primary=pulumi.get(__ret__, 'gre_range_primary'),
        gre_range_secondary=pulumi.get(__ret__, 'gre_range_secondary'),
        gre_tunnel_ip=pulumi.get(__ret__, 'gre_tunnel_ip'),
        id=pulumi.get(__ret__, 'id'),
        ip_address=pulumi.get(__ret__, 'ip_address'),
        primary_gw=pulumi.get(__ret__, 'primary_gw'),
        secondary_gw=pulumi.get(__ret__, 'secondary_gw'),
        tun_id=pulumi.get(__ret__, 'tun_id'))


@_utilities.lift_output_func(get_traffic_forwarding_gre_tunnel_info)
def get_traffic_forwarding_gre_tunnel_info_output(gre_enabled: Optional[pulumi.Input[Optional[bool]]] = None,
                                                  ip_address: Optional[pulumi.Input[Optional[str]]] = None,
                                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTrafficForwardingGRETunnelInfoResult]:
    """
    The **zia_traffic_forwarding_gre_tunnel_info** data source to get information about provisioned GRE tunnel information created in the Zscaler Internet Access portal.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zia as zia

    example = zia.get_traffic_forwarding_gre_tunnel_info(ip_address="1.1.1.1")
    ```


    :param bool gre_enabled: Displays only ip addresses with GRE tunnel enabled
           
           > **NOTE** `ip_address` is the public IP address (Static IP) associated with the GRE Tunnel
    :param str ip_address: Filter based on an IP address range.
    """
    ...
