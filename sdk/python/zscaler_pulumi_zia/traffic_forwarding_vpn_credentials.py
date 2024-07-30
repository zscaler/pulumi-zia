# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TrafficForwardingVPNCredentialsArgs', 'TrafficForwardingVPNCredentials']

@pulumi.input_type
class TrafficForwardingVPNCredentialsArgs:
    def __init__(__self__, *,
                 comments: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 pre_shared_key: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TrafficForwardingVPNCredentials resource.
        :param pulumi.Input[str] fqdn: Fully Qualified Domain Name. Applicable only to `UFQDN` or `XAUTH` (or `HOSTED_MOBILE_USERS`) auth type.
        :param pulumi.Input[str] pre_shared_key: Pre-shared key. This is a required field for UFQDN and IP auth type.
        :param pulumi.Input[str] type: VPN authentication type (i.e., how the VPN credential is sent to the server). It is not modifiable after VpnCredential is created. The supported values are: `UFQDN` and `IP`
        """
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if fqdn is not None:
            pulumi.set(__self__, "fqdn", fqdn)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if pre_shared_key is not None:
            pulumi.set(__self__, "pre_shared_key", pre_shared_key)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter
    def fqdn(self) -> Optional[pulumi.Input[str]]:
        """
        Fully Qualified Domain Name. Applicable only to `UFQDN` or `XAUTH` (or `HOSTED_MOBILE_USERS`) auth type.
        """
        return pulumi.get(self, "fqdn")

    @fqdn.setter
    def fqdn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="preSharedKey")
    def pre_shared_key(self) -> Optional[pulumi.Input[str]]:
        """
        Pre-shared key. This is a required field for UFQDN and IP auth type.
        """
        return pulumi.get(self, "pre_shared_key")

    @pre_shared_key.setter
    def pre_shared_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pre_shared_key", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        VPN authentication type (i.e., how the VPN credential is sent to the server). It is not modifiable after VpnCredential is created. The supported values are: `UFQDN` and `IP`
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _TrafficForwardingVPNCredentialsState:
    def __init__(__self__, *,
                 comments: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 pre_shared_key: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vpn_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering TrafficForwardingVPNCredentials resources.
        :param pulumi.Input[str] fqdn: Fully Qualified Domain Name. Applicable only to `UFQDN` or `XAUTH` (or `HOSTED_MOBILE_USERS`) auth type.
        :param pulumi.Input[str] pre_shared_key: Pre-shared key. This is a required field for UFQDN and IP auth type.
        :param pulumi.Input[str] type: VPN authentication type (i.e., how the VPN credential is sent to the server). It is not modifiable after VpnCredential is created. The supported values are: `UFQDN` and `IP`
        """
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if fqdn is not None:
            pulumi.set(__self__, "fqdn", fqdn)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if pre_shared_key is not None:
            pulumi.set(__self__, "pre_shared_key", pre_shared_key)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vpn_id is not None:
            pulumi.set(__self__, "vpn_id", vpn_id)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter
    def fqdn(self) -> Optional[pulumi.Input[str]]:
        """
        Fully Qualified Domain Name. Applicable only to `UFQDN` or `XAUTH` (or `HOSTED_MOBILE_USERS`) auth type.
        """
        return pulumi.get(self, "fqdn")

    @fqdn.setter
    def fqdn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="preSharedKey")
    def pre_shared_key(self) -> Optional[pulumi.Input[str]]:
        """
        Pre-shared key. This is a required field for UFQDN and IP auth type.
        """
        return pulumi.get(self, "pre_shared_key")

    @pre_shared_key.setter
    def pre_shared_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pre_shared_key", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        VPN authentication type (i.e., how the VPN credential is sent to the server). It is not modifiable after VpnCredential is created. The supported values are: `UFQDN` and `IP`
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="vpnId")
    def vpn_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "vpn_id")

    @vpn_id.setter
    def vpn_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vpn_id", value)


class TrafficForwardingVPNCredentials(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 pre_shared_key: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The **zia_traffic_forwarding_vpn_credentials** creates and manages VPN credentials that can be associated to locations. VPN is one way to route traffic from customer locations to the cloud. Site-to-site IPSec VPN credentials can be identified by the cloud through one of the following methods:

        * Common Name (CN) of IPSec Certificate
        * VPN User FQDN - requires VPN_SITE_TO_SITE subscription
        * VPN IP Address - requires VPN_SITE_TO_SITE subscription
        * Extended Authentication (XAUTH) or hosted mobile UserID - requires VPN_MOBILE subscription

        ## Example Usage

        ```python
        import pulumi
        import zscaler_pulumi_zia as zia

        ######### PASSWORDS IN THIS FILE ARE FAKE AND NOT USED IN PRODUCTION SYSTEMS #########
        # ZIA Traffic Forwarding - VPN Credentials (UFQDN)
        example = zia.TrafficForwardingVPNCredentials("example",
            comments="Example",
            fqdn="sjc-1-37@acme.com",
            pre_shared_key="*********************",
            type="UFQDN")
        ```

        ```python
        import pulumi
        import zscaler_pulumi_zia as zia

        example_traffic_forwarding_static_ip = zia.TrafficForwardingStaticIP("exampleTrafficForwardingStaticIP",
            ip_address="1.1.1.1",
            routable_ip=True,
            comment="Example",
            geo_override=True,
            latitude=-36.848461,
            longitude=174.763336)
        # ZIA Traffic Forwarding - VPN Credentials (IP)
        ######### PASSWORDS IN THIS FILE ARE FAKE AND NOT USED IN PRODUCTION SYSTEMS #########
        example_traffic_forwarding_vpn_credentials = zia.TrafficForwardingVPNCredentials("exampleTrafficForwardingVPNCredentials",
            type="IP",
            ip_address=example_traffic_forwarding_static_ip.ip_address,
            comments="Example",
            pre_shared_key="*********************",
            opts = pulumi.ResourceOptions(depends_on=[example_traffic_forwarding_static_ip]))
        ```

        > **NOTE** For VPN Credentials of Type `IP` a static IP resource must be created first.

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        **zia_traffic_forwarding_vpn_credentials** can be imported by using one of the following prefixes as the import ID:

        * `'IP'` - Imports all VPN Credentials of type IP

        ```sh
        $ pulumi import zia:index/trafficForwardingVPNCredentials:TrafficForwardingVPNCredentials example 'IP'
        ```

        * `'UFQDN'` - Imports all VPN Credentials of type UFQDN

        ```sh
        $ pulumi import zia:index/trafficForwardingVPNCredentials:TrafficForwardingVPNCredentials this 'UFQDN'
        ```

        * `UFQDN'` - Imports a VPN Credentials of type UFQDN containing a specific UFQDN address

        ```sh
        $ pulumi import zia:index/trafficForwardingVPNCredentials:TrafficForwardingVPNCredentials example 'testvpn@example.com'
        ```

        * `IP Address'` - Imports a VPN Credentials of type IP containing a specific IP address

        ```sh
        $ pulumi import zia:index/trafficForwardingVPNCredentials:TrafficForwardingVPNCredentials example '1.1.1.1'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fqdn: Fully Qualified Domain Name. Applicable only to `UFQDN` or `XAUTH` (or `HOSTED_MOBILE_USERS`) auth type.
        :param pulumi.Input[str] pre_shared_key: Pre-shared key. This is a required field for UFQDN and IP auth type.
        :param pulumi.Input[str] type: VPN authentication type (i.e., how the VPN credential is sent to the server). It is not modifiable after VpnCredential is created. The supported values are: `UFQDN` and `IP`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TrafficForwardingVPNCredentialsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The **zia_traffic_forwarding_vpn_credentials** creates and manages VPN credentials that can be associated to locations. VPN is one way to route traffic from customer locations to the cloud. Site-to-site IPSec VPN credentials can be identified by the cloud through one of the following methods:

        * Common Name (CN) of IPSec Certificate
        * VPN User FQDN - requires VPN_SITE_TO_SITE subscription
        * VPN IP Address - requires VPN_SITE_TO_SITE subscription
        * Extended Authentication (XAUTH) or hosted mobile UserID - requires VPN_MOBILE subscription

        ## Example Usage

        ```python
        import pulumi
        import zscaler_pulumi_zia as zia

        ######### PASSWORDS IN THIS FILE ARE FAKE AND NOT USED IN PRODUCTION SYSTEMS #########
        # ZIA Traffic Forwarding - VPN Credentials (UFQDN)
        example = zia.TrafficForwardingVPNCredentials("example",
            comments="Example",
            fqdn="sjc-1-37@acme.com",
            pre_shared_key="*********************",
            type="UFQDN")
        ```

        ```python
        import pulumi
        import zscaler_pulumi_zia as zia

        example_traffic_forwarding_static_ip = zia.TrafficForwardingStaticIP("exampleTrafficForwardingStaticIP",
            ip_address="1.1.1.1",
            routable_ip=True,
            comment="Example",
            geo_override=True,
            latitude=-36.848461,
            longitude=174.763336)
        # ZIA Traffic Forwarding - VPN Credentials (IP)
        ######### PASSWORDS IN THIS FILE ARE FAKE AND NOT USED IN PRODUCTION SYSTEMS #########
        example_traffic_forwarding_vpn_credentials = zia.TrafficForwardingVPNCredentials("exampleTrafficForwardingVPNCredentials",
            type="IP",
            ip_address=example_traffic_forwarding_static_ip.ip_address,
            comments="Example",
            pre_shared_key="*********************",
            opts = pulumi.ResourceOptions(depends_on=[example_traffic_forwarding_static_ip]))
        ```

        > **NOTE** For VPN Credentials of Type `IP` a static IP resource must be created first.

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        **zia_traffic_forwarding_vpn_credentials** can be imported by using one of the following prefixes as the import ID:

        * `'IP'` - Imports all VPN Credentials of type IP

        ```sh
        $ pulumi import zia:index/trafficForwardingVPNCredentials:TrafficForwardingVPNCredentials example 'IP'
        ```

        * `'UFQDN'` - Imports all VPN Credentials of type UFQDN

        ```sh
        $ pulumi import zia:index/trafficForwardingVPNCredentials:TrafficForwardingVPNCredentials this 'UFQDN'
        ```

        * `UFQDN'` - Imports a VPN Credentials of type UFQDN containing a specific UFQDN address

        ```sh
        $ pulumi import zia:index/trafficForwardingVPNCredentials:TrafficForwardingVPNCredentials example 'testvpn@example.com'
        ```

        * `IP Address'` - Imports a VPN Credentials of type IP containing a specific IP address

        ```sh
        $ pulumi import zia:index/trafficForwardingVPNCredentials:TrafficForwardingVPNCredentials example '1.1.1.1'
        ```

        :param str resource_name: The name of the resource.
        :param TrafficForwardingVPNCredentialsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TrafficForwardingVPNCredentialsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 pre_shared_key: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TrafficForwardingVPNCredentialsArgs.__new__(TrafficForwardingVPNCredentialsArgs)

            __props__.__dict__["comments"] = comments
            __props__.__dict__["fqdn"] = fqdn
            __props__.__dict__["ip_address"] = ip_address
            __props__.__dict__["pre_shared_key"] = None if pre_shared_key is None else pulumi.Output.secret(pre_shared_key)
            __props__.__dict__["type"] = type
            __props__.__dict__["vpn_id"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["preSharedKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(TrafficForwardingVPNCredentials, __self__).__init__(
            'zia:index/trafficForwardingVPNCredentials:TrafficForwardingVPNCredentials',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comments: Optional[pulumi.Input[str]] = None,
            fqdn: Optional[pulumi.Input[str]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            pre_shared_key: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            vpn_id: Optional[pulumi.Input[int]] = None) -> 'TrafficForwardingVPNCredentials':
        """
        Get an existing TrafficForwardingVPNCredentials resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fqdn: Fully Qualified Domain Name. Applicable only to `UFQDN` or `XAUTH` (or `HOSTED_MOBILE_USERS`) auth type.
        :param pulumi.Input[str] pre_shared_key: Pre-shared key. This is a required field for UFQDN and IP auth type.
        :param pulumi.Input[str] type: VPN authentication type (i.e., how the VPN credential is sent to the server). It is not modifiable after VpnCredential is created. The supported values are: `UFQDN` and `IP`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TrafficForwardingVPNCredentialsState.__new__(_TrafficForwardingVPNCredentialsState)

        __props__.__dict__["comments"] = comments
        __props__.__dict__["fqdn"] = fqdn
        __props__.__dict__["ip_address"] = ip_address
        __props__.__dict__["pre_shared_key"] = pre_shared_key
        __props__.__dict__["type"] = type
        __props__.__dict__["vpn_id"] = vpn_id
        return TrafficForwardingVPNCredentials(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def fqdn(self) -> pulumi.Output[Optional[str]]:
        """
        Fully Qualified Domain Name. Applicable only to `UFQDN` or `XAUTH` (or `HOSTED_MOBILE_USERS`) auth type.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="preSharedKey")
    def pre_shared_key(self) -> pulumi.Output[Optional[str]]:
        """
        Pre-shared key. This is a required field for UFQDN and IP auth type.
        """
        return pulumi.get(self, "pre_shared_key")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        VPN authentication type (i.e., how the VPN credential is sent to the server). It is not modifiable after VpnCredential is created. The supported values are: `UFQDN` and `IP`
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vpnId")
    def vpn_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "vpn_id")

