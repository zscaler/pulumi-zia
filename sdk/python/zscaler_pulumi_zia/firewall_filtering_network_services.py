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
from ._inputs import *

__all__ = ['FirewallFilteringNetworkServicesArgs', 'FirewallFilteringNetworkServices']

@pulumi.input_type
class FirewallFilteringNetworkServicesArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 dest_tcp_ports: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesDestTcpPortArgs']]]] = None,
                 dest_udp_ports: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesDestUdpPortArgs']]]] = None,
                 is_name_l10n_tag: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 src_tcp_ports: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesSrcTcpPortArgs']]]] = None,
                 src_udp_ports: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesSrcUdpPortArgs']]]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallFilteringNetworkServices resource.
        :param pulumi.Input[str] description: Description of the service
        :param pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesDestTcpPortArgs']]] dest_tcp_ports: The TCP destination port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesDestUdpPortArgs']]] dest_udp_ports: The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        :param pulumi.Input[bool] is_name_l10n_tag: (Optional
        :param pulumi.Input[str] name: Name of the service
        :param pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesSrcTcpPortArgs']]] src_tcp_ports: The TCP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service
        :param pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesSrcUdpPortArgs']]] src_udp_ports: The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        :param pulumi.Input[str] tag: The following values are supported: `"ICMP_ANY`, `"UDP_ANY"`, `"TCP_ANY"`, `"OTHER_NETWORK_SERVICE"`, `"DNS"`, `"NETBIOS"`, `"FTP"`, `"GNUTELLA"`, `"H_323"`, `"HTTP"`, `"HTTPS"`, `"IKE"`, `"IMAP"`, `"ILS"`, `"IKE_NAT"`, `"IRC"`, `"LDAP"`, `"QUIC"`, `"TDS"`, `"NETMEETING"`, `"NFS"`, `"NTP"`, `"SIP"`, `"SNMP"`, `"SMB"`, `"SMTP"`, `"SSH"`, `"SYSLOG"`, `"TELNET"`, `"TRACEROUTE"`, `"POP3"`, `"PPTP"`, `"RADIUS"`, `"REAL_MEDIA"`, `"RTSP"`, `"VNC"`, `"WHOIS"`, `"KERBEROS_SEC"`, `"TACACS"`, `"SNMPTRAP"`, `"NMAP"`, `"RSYNC"`, `"L2TP"`, `"HTTP_PROXY"`, `"PC_ANYWHERE"`, `"MSN"`, `"ECHO"`, `"AIM"`, `"IDENT"`, `"YMSG"`, `"SCCP"`, `"MGCP_UA"`, `"MGCP_CA"`, `"VDO_LIVE"`, `"OPENVPN"`, `"TFTP"`, `"FTPS_IMPLICIT"`, `"ZSCALER_PROXY_NW_SERVICES"`, `"GRE_PROTOCOL"`, `"ESP_PROTOCOL"`, `"DHCP"`
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dest_tcp_ports is not None:
            pulumi.set(__self__, "dest_tcp_ports", dest_tcp_ports)
        if dest_udp_ports is not None:
            pulumi.set(__self__, "dest_udp_ports", dest_udp_ports)
        if is_name_l10n_tag is not None:
            pulumi.set(__self__, "is_name_l10n_tag", is_name_l10n_tag)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if src_tcp_ports is not None:
            pulumi.set(__self__, "src_tcp_ports", src_tcp_ports)
        if src_udp_ports is not None:
            pulumi.set(__self__, "src_udp_ports", src_udp_ports)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the service
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="destTcpPorts")
    def dest_tcp_ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesDestTcpPortArgs']]]]:
        """
        The TCP destination port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        """
        return pulumi.get(self, "dest_tcp_ports")

    @dest_tcp_ports.setter
    def dest_tcp_ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesDestTcpPortArgs']]]]):
        pulumi.set(self, "dest_tcp_ports", value)

    @property
    @pulumi.getter(name="destUdpPorts")
    def dest_udp_ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesDestUdpPortArgs']]]]:
        """
        The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        """
        return pulumi.get(self, "dest_udp_ports")

    @dest_udp_ports.setter
    def dest_udp_ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesDestUdpPortArgs']]]]):
        pulumi.set(self, "dest_udp_ports", value)

    @property
    @pulumi.getter(name="isNameL10nTag")
    def is_name_l10n_tag(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional
        """
        return pulumi.get(self, "is_name_l10n_tag")

    @is_name_l10n_tag.setter
    def is_name_l10n_tag(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_name_l10n_tag", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the service
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="srcTcpPorts")
    def src_tcp_ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesSrcTcpPortArgs']]]]:
        """
        The TCP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service
        """
        return pulumi.get(self, "src_tcp_ports")

    @src_tcp_ports.setter
    def src_tcp_ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesSrcTcpPortArgs']]]]):
        pulumi.set(self, "src_tcp_ports", value)

    @property
    @pulumi.getter(name="srcUdpPorts")
    def src_udp_ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesSrcUdpPortArgs']]]]:
        """
        The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        """
        return pulumi.get(self, "src_udp_ports")

    @src_udp_ports.setter
    def src_udp_ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesSrcUdpPortArgs']]]]):
        pulumi.set(self, "src_udp_ports", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[str]]:
        """
        The following values are supported: `"ICMP_ANY`, `"UDP_ANY"`, `"TCP_ANY"`, `"OTHER_NETWORK_SERVICE"`, `"DNS"`, `"NETBIOS"`, `"FTP"`, `"GNUTELLA"`, `"H_323"`, `"HTTP"`, `"HTTPS"`, `"IKE"`, `"IMAP"`, `"ILS"`, `"IKE_NAT"`, `"IRC"`, `"LDAP"`, `"QUIC"`, `"TDS"`, `"NETMEETING"`, `"NFS"`, `"NTP"`, `"SIP"`, `"SNMP"`, `"SMB"`, `"SMTP"`, `"SSH"`, `"SYSLOG"`, `"TELNET"`, `"TRACEROUTE"`, `"POP3"`, `"PPTP"`, `"RADIUS"`, `"REAL_MEDIA"`, `"RTSP"`, `"VNC"`, `"WHOIS"`, `"KERBEROS_SEC"`, `"TACACS"`, `"SNMPTRAP"`, `"NMAP"`, `"RSYNC"`, `"L2TP"`, `"HTTP_PROXY"`, `"PC_ANYWHERE"`, `"MSN"`, `"ECHO"`, `"AIM"`, `"IDENT"`, `"YMSG"`, `"SCCP"`, `"MGCP_UA"`, `"MGCP_CA"`, `"VDO_LIVE"`, `"OPENVPN"`, `"TFTP"`, `"FTPS_IMPLICIT"`, `"ZSCALER_PROXY_NW_SERVICES"`, `"GRE_PROTOCOL"`, `"ESP_PROTOCOL"`, `"DHCP"`
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _FirewallFilteringNetworkServicesState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 dest_tcp_ports: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesDestTcpPortArgs']]]] = None,
                 dest_udp_ports: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesDestUdpPortArgs']]]] = None,
                 is_name_l10n_tag: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_service_id: Optional[pulumi.Input[int]] = None,
                 src_tcp_ports: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesSrcTcpPortArgs']]]] = None,
                 src_udp_ports: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesSrcUdpPortArgs']]]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallFilteringNetworkServices resources.
        :param pulumi.Input[str] description: Description of the service
        :param pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesDestTcpPortArgs']]] dest_tcp_ports: The TCP destination port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesDestUdpPortArgs']]] dest_udp_ports: The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        :param pulumi.Input[bool] is_name_l10n_tag: (Optional
        :param pulumi.Input[str] name: Name of the service
        :param pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesSrcTcpPortArgs']]] src_tcp_ports: The TCP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service
        :param pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesSrcUdpPortArgs']]] src_udp_ports: The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        :param pulumi.Input[str] tag: The following values are supported: `"ICMP_ANY`, `"UDP_ANY"`, `"TCP_ANY"`, `"OTHER_NETWORK_SERVICE"`, `"DNS"`, `"NETBIOS"`, `"FTP"`, `"GNUTELLA"`, `"H_323"`, `"HTTP"`, `"HTTPS"`, `"IKE"`, `"IMAP"`, `"ILS"`, `"IKE_NAT"`, `"IRC"`, `"LDAP"`, `"QUIC"`, `"TDS"`, `"NETMEETING"`, `"NFS"`, `"NTP"`, `"SIP"`, `"SNMP"`, `"SMB"`, `"SMTP"`, `"SSH"`, `"SYSLOG"`, `"TELNET"`, `"TRACEROUTE"`, `"POP3"`, `"PPTP"`, `"RADIUS"`, `"REAL_MEDIA"`, `"RTSP"`, `"VNC"`, `"WHOIS"`, `"KERBEROS_SEC"`, `"TACACS"`, `"SNMPTRAP"`, `"NMAP"`, `"RSYNC"`, `"L2TP"`, `"HTTP_PROXY"`, `"PC_ANYWHERE"`, `"MSN"`, `"ECHO"`, `"AIM"`, `"IDENT"`, `"YMSG"`, `"SCCP"`, `"MGCP_UA"`, `"MGCP_CA"`, `"VDO_LIVE"`, `"OPENVPN"`, `"TFTP"`, `"FTPS_IMPLICIT"`, `"ZSCALER_PROXY_NW_SERVICES"`, `"GRE_PROTOCOL"`, `"ESP_PROTOCOL"`, `"DHCP"`
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dest_tcp_ports is not None:
            pulumi.set(__self__, "dest_tcp_ports", dest_tcp_ports)
        if dest_udp_ports is not None:
            pulumi.set(__self__, "dest_udp_ports", dest_udp_ports)
        if is_name_l10n_tag is not None:
            pulumi.set(__self__, "is_name_l10n_tag", is_name_l10n_tag)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_service_id is not None:
            pulumi.set(__self__, "network_service_id", network_service_id)
        if src_tcp_ports is not None:
            pulumi.set(__self__, "src_tcp_ports", src_tcp_ports)
        if src_udp_ports is not None:
            pulumi.set(__self__, "src_udp_ports", src_udp_ports)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the service
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="destTcpPorts")
    def dest_tcp_ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesDestTcpPortArgs']]]]:
        """
        The TCP destination port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        """
        return pulumi.get(self, "dest_tcp_ports")

    @dest_tcp_ports.setter
    def dest_tcp_ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesDestTcpPortArgs']]]]):
        pulumi.set(self, "dest_tcp_ports", value)

    @property
    @pulumi.getter(name="destUdpPorts")
    def dest_udp_ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesDestUdpPortArgs']]]]:
        """
        The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        """
        return pulumi.get(self, "dest_udp_ports")

    @dest_udp_ports.setter
    def dest_udp_ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesDestUdpPortArgs']]]]):
        pulumi.set(self, "dest_udp_ports", value)

    @property
    @pulumi.getter(name="isNameL10nTag")
    def is_name_l10n_tag(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional
        """
        return pulumi.get(self, "is_name_l10n_tag")

    @is_name_l10n_tag.setter
    def is_name_l10n_tag(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_name_l10n_tag", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the service
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkServiceId")
    def network_service_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "network_service_id")

    @network_service_id.setter
    def network_service_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "network_service_id", value)

    @property
    @pulumi.getter(name="srcTcpPorts")
    def src_tcp_ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesSrcTcpPortArgs']]]]:
        """
        The TCP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service
        """
        return pulumi.get(self, "src_tcp_ports")

    @src_tcp_ports.setter
    def src_tcp_ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesSrcTcpPortArgs']]]]):
        pulumi.set(self, "src_tcp_ports", value)

    @property
    @pulumi.getter(name="srcUdpPorts")
    def src_udp_ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesSrcUdpPortArgs']]]]:
        """
        The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        """
        return pulumi.get(self, "src_udp_ports")

    @src_udp_ports.setter
    def src_udp_ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallFilteringNetworkServicesSrcUdpPortArgs']]]]):
        pulumi.set(self, "src_udp_ports", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[str]]:
        """
        The following values are supported: `"ICMP_ANY`, `"UDP_ANY"`, `"TCP_ANY"`, `"OTHER_NETWORK_SERVICE"`, `"DNS"`, `"NETBIOS"`, `"FTP"`, `"GNUTELLA"`, `"H_323"`, `"HTTP"`, `"HTTPS"`, `"IKE"`, `"IMAP"`, `"ILS"`, `"IKE_NAT"`, `"IRC"`, `"LDAP"`, `"QUIC"`, `"TDS"`, `"NETMEETING"`, `"NFS"`, `"NTP"`, `"SIP"`, `"SNMP"`, `"SMB"`, `"SMTP"`, `"SSH"`, `"SYSLOG"`, `"TELNET"`, `"TRACEROUTE"`, `"POP3"`, `"PPTP"`, `"RADIUS"`, `"REAL_MEDIA"`, `"RTSP"`, `"VNC"`, `"WHOIS"`, `"KERBEROS_SEC"`, `"TACACS"`, `"SNMPTRAP"`, `"NMAP"`, `"RSYNC"`, `"L2TP"`, `"HTTP_PROXY"`, `"PC_ANYWHERE"`, `"MSN"`, `"ECHO"`, `"AIM"`, `"IDENT"`, `"YMSG"`, `"SCCP"`, `"MGCP_UA"`, `"MGCP_CA"`, `"VDO_LIVE"`, `"OPENVPN"`, `"TFTP"`, `"FTPS_IMPLICIT"`, `"ZSCALER_PROXY_NW_SERVICES"`, `"GRE_PROTOCOL"`, `"ESP_PROTOCOL"`, `"DHCP"`
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class FirewallFilteringNetworkServices(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dest_tcp_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesDestTcpPortArgs']]]]] = None,
                 dest_udp_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesDestUdpPortArgs']]]]] = None,
                 is_name_l10n_tag: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 src_tcp_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesSrcTcpPortArgs']]]]] = None,
                 src_udp_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesSrcUdpPortArgs']]]]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The **zia_firewall_filtering_network_service** resource allows the creation and management of ZIA Cloud Firewall IP network services in the Zscaler Internet Access. This resource can then be associated with a ZIA cloud firewall filtering rule and network service group resources.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import zscaler_pulumi_zia as zia

        example = zia.FirewallFilteringNetworkServices("example",
            description="example",
            dest_tcp_ports=[
                zia.FirewallFilteringNetworkServicesDestTcpPortArgs(
                    start=5000,
                ),
                zia.FirewallFilteringNetworkServicesDestTcpPortArgs(
                    start=5001,
                ),
                zia.FirewallFilteringNetworkServicesDestTcpPortArgs(
                    end=5005,
                    start=5003,
                ),
            ],
            src_tcp_ports=[
                zia.FirewallFilteringNetworkServicesSrcTcpPortArgs(
                    start=5000,
                ),
                zia.FirewallFilteringNetworkServicesSrcTcpPortArgs(
                    start=5001,
                ),
                zia.FirewallFilteringNetworkServicesSrcTcpPortArgs(
                    end=5005,
                    start=5002,
                ),
            ],
            type="CUSTOM")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        **zia_firewall_filtering_network_service** can be imported by using `<SERVICE_ID>` or `<SERVICE_NAME>` as the import ID.

        For example:

        ```sh
        $ pulumi import zia:index/firewallFilteringNetworkServices:FirewallFilteringNetworkServices example <service_id>
        ```

        or

        ```sh
        $ pulumi import zia:index/firewallFilteringNetworkServices:FirewallFilteringNetworkServices example <service_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the service
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesDestTcpPortArgs']]]] dest_tcp_ports: The TCP destination port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesDestUdpPortArgs']]]] dest_udp_ports: The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        :param pulumi.Input[bool] is_name_l10n_tag: (Optional
        :param pulumi.Input[str] name: Name of the service
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesSrcTcpPortArgs']]]] src_tcp_ports: The TCP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesSrcUdpPortArgs']]]] src_udp_ports: The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        :param pulumi.Input[str] tag: The following values are supported: `"ICMP_ANY`, `"UDP_ANY"`, `"TCP_ANY"`, `"OTHER_NETWORK_SERVICE"`, `"DNS"`, `"NETBIOS"`, `"FTP"`, `"GNUTELLA"`, `"H_323"`, `"HTTP"`, `"HTTPS"`, `"IKE"`, `"IMAP"`, `"ILS"`, `"IKE_NAT"`, `"IRC"`, `"LDAP"`, `"QUIC"`, `"TDS"`, `"NETMEETING"`, `"NFS"`, `"NTP"`, `"SIP"`, `"SNMP"`, `"SMB"`, `"SMTP"`, `"SSH"`, `"SYSLOG"`, `"TELNET"`, `"TRACEROUTE"`, `"POP3"`, `"PPTP"`, `"RADIUS"`, `"REAL_MEDIA"`, `"RTSP"`, `"VNC"`, `"WHOIS"`, `"KERBEROS_SEC"`, `"TACACS"`, `"SNMPTRAP"`, `"NMAP"`, `"RSYNC"`, `"L2TP"`, `"HTTP_PROXY"`, `"PC_ANYWHERE"`, `"MSN"`, `"ECHO"`, `"AIM"`, `"IDENT"`, `"YMSG"`, `"SCCP"`, `"MGCP_UA"`, `"MGCP_CA"`, `"VDO_LIVE"`, `"OPENVPN"`, `"TFTP"`, `"FTPS_IMPLICIT"`, `"ZSCALER_PROXY_NW_SERVICES"`, `"GRE_PROTOCOL"`, `"ESP_PROTOCOL"`, `"DHCP"`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FirewallFilteringNetworkServicesArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The **zia_firewall_filtering_network_service** resource allows the creation and management of ZIA Cloud Firewall IP network services in the Zscaler Internet Access. This resource can then be associated with a ZIA cloud firewall filtering rule and network service group resources.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import zscaler_pulumi_zia as zia

        example = zia.FirewallFilteringNetworkServices("example",
            description="example",
            dest_tcp_ports=[
                zia.FirewallFilteringNetworkServicesDestTcpPortArgs(
                    start=5000,
                ),
                zia.FirewallFilteringNetworkServicesDestTcpPortArgs(
                    start=5001,
                ),
                zia.FirewallFilteringNetworkServicesDestTcpPortArgs(
                    end=5005,
                    start=5003,
                ),
            ],
            src_tcp_ports=[
                zia.FirewallFilteringNetworkServicesSrcTcpPortArgs(
                    start=5000,
                ),
                zia.FirewallFilteringNetworkServicesSrcTcpPortArgs(
                    start=5001,
                ),
                zia.FirewallFilteringNetworkServicesSrcTcpPortArgs(
                    end=5005,
                    start=5002,
                ),
            ],
            type="CUSTOM")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        **zia_firewall_filtering_network_service** can be imported by using `<SERVICE_ID>` or `<SERVICE_NAME>` as the import ID.

        For example:

        ```sh
        $ pulumi import zia:index/firewallFilteringNetworkServices:FirewallFilteringNetworkServices example <service_id>
        ```

        or

        ```sh
        $ pulumi import zia:index/firewallFilteringNetworkServices:FirewallFilteringNetworkServices example <service_name>
        ```

        :param str resource_name: The name of the resource.
        :param FirewallFilteringNetworkServicesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallFilteringNetworkServicesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dest_tcp_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesDestTcpPortArgs']]]]] = None,
                 dest_udp_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesDestUdpPortArgs']]]]] = None,
                 is_name_l10n_tag: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 src_tcp_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesSrcTcpPortArgs']]]]] = None,
                 src_udp_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesSrcUdpPortArgs']]]]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallFilteringNetworkServicesArgs.__new__(FirewallFilteringNetworkServicesArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["dest_tcp_ports"] = dest_tcp_ports
            __props__.__dict__["dest_udp_ports"] = dest_udp_ports
            __props__.__dict__["is_name_l10n_tag"] = is_name_l10n_tag
            __props__.__dict__["name"] = name
            __props__.__dict__["src_tcp_ports"] = src_tcp_ports
            __props__.__dict__["src_udp_ports"] = src_udp_ports
            __props__.__dict__["tag"] = tag
            __props__.__dict__["type"] = type
            __props__.__dict__["network_service_id"] = None
        super(FirewallFilteringNetworkServices, __self__).__init__(
            'zia:index/firewallFilteringNetworkServices:FirewallFilteringNetworkServices',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            dest_tcp_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesDestTcpPortArgs']]]]] = None,
            dest_udp_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesDestUdpPortArgs']]]]] = None,
            is_name_l10n_tag: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_service_id: Optional[pulumi.Input[int]] = None,
            src_tcp_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesSrcTcpPortArgs']]]]] = None,
            src_udp_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesSrcUdpPortArgs']]]]] = None,
            tag: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'FirewallFilteringNetworkServices':
        """
        Get an existing FirewallFilteringNetworkServices resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the service
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesDestTcpPortArgs']]]] dest_tcp_ports: The TCP destination port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesDestUdpPortArgs']]]] dest_udp_ports: The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        :param pulumi.Input[bool] is_name_l10n_tag: (Optional
        :param pulumi.Input[str] name: Name of the service
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesSrcTcpPortArgs']]]] src_tcp_ports: The TCP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallFilteringNetworkServicesSrcUdpPortArgs']]]] src_udp_ports: The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        :param pulumi.Input[str] tag: The following values are supported: `"ICMP_ANY`, `"UDP_ANY"`, `"TCP_ANY"`, `"OTHER_NETWORK_SERVICE"`, `"DNS"`, `"NETBIOS"`, `"FTP"`, `"GNUTELLA"`, `"H_323"`, `"HTTP"`, `"HTTPS"`, `"IKE"`, `"IMAP"`, `"ILS"`, `"IKE_NAT"`, `"IRC"`, `"LDAP"`, `"QUIC"`, `"TDS"`, `"NETMEETING"`, `"NFS"`, `"NTP"`, `"SIP"`, `"SNMP"`, `"SMB"`, `"SMTP"`, `"SSH"`, `"SYSLOG"`, `"TELNET"`, `"TRACEROUTE"`, `"POP3"`, `"PPTP"`, `"RADIUS"`, `"REAL_MEDIA"`, `"RTSP"`, `"VNC"`, `"WHOIS"`, `"KERBEROS_SEC"`, `"TACACS"`, `"SNMPTRAP"`, `"NMAP"`, `"RSYNC"`, `"L2TP"`, `"HTTP_PROXY"`, `"PC_ANYWHERE"`, `"MSN"`, `"ECHO"`, `"AIM"`, `"IDENT"`, `"YMSG"`, `"SCCP"`, `"MGCP_UA"`, `"MGCP_CA"`, `"VDO_LIVE"`, `"OPENVPN"`, `"TFTP"`, `"FTPS_IMPLICIT"`, `"ZSCALER_PROXY_NW_SERVICES"`, `"GRE_PROTOCOL"`, `"ESP_PROTOCOL"`, `"DHCP"`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallFilteringNetworkServicesState.__new__(_FirewallFilteringNetworkServicesState)

        __props__.__dict__["description"] = description
        __props__.__dict__["dest_tcp_ports"] = dest_tcp_ports
        __props__.__dict__["dest_udp_ports"] = dest_udp_ports
        __props__.__dict__["is_name_l10n_tag"] = is_name_l10n_tag
        __props__.__dict__["name"] = name
        __props__.__dict__["network_service_id"] = network_service_id
        __props__.__dict__["src_tcp_ports"] = src_tcp_ports
        __props__.__dict__["src_udp_ports"] = src_udp_ports
        __props__.__dict__["tag"] = tag
        __props__.__dict__["type"] = type
        return FirewallFilteringNetworkServices(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the service
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="destTcpPorts")
    def dest_tcp_ports(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallFilteringNetworkServicesDestTcpPort']]]:
        """
        The TCP destination port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        """
        return pulumi.get(self, "dest_tcp_ports")

    @property
    @pulumi.getter(name="destUdpPorts")
    def dest_udp_ports(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallFilteringNetworkServicesDestUdpPort']]]:
        """
        The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        """
        return pulumi.get(self, "dest_udp_ports")

    @property
    @pulumi.getter(name="isNameL10nTag")
    def is_name_l10n_tag(self) -> pulumi.Output[Optional[bool]]:
        """
        (Optional
        """
        return pulumi.get(self, "is_name_l10n_tag")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the service
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkServiceId")
    def network_service_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "network_service_id")

    @property
    @pulumi.getter(name="srcTcpPorts")
    def src_tcp_ports(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallFilteringNetworkServicesSrcTcpPort']]]:
        """
        The TCP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service
        """
        return pulumi.get(self, "src_tcp_ports")

    @property
    @pulumi.getter(name="srcUdpPorts")
    def src_udp_ports(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallFilteringNetworkServicesSrcUdpPort']]]:
        """
        The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.
        """
        return pulumi.get(self, "src_udp_ports")

    @property
    @pulumi.getter
    def tag(self) -> pulumi.Output[str]:
        """
        The following values are supported: `"ICMP_ANY`, `"UDP_ANY"`, `"TCP_ANY"`, `"OTHER_NETWORK_SERVICE"`, `"DNS"`, `"NETBIOS"`, `"FTP"`, `"GNUTELLA"`, `"H_323"`, `"HTTP"`, `"HTTPS"`, `"IKE"`, `"IMAP"`, `"ILS"`, `"IKE_NAT"`, `"IRC"`, `"LDAP"`, `"QUIC"`, `"TDS"`, `"NETMEETING"`, `"NFS"`, `"NTP"`, `"SIP"`, `"SNMP"`, `"SMB"`, `"SMTP"`, `"SSH"`, `"SYSLOG"`, `"TELNET"`, `"TRACEROUTE"`, `"POP3"`, `"PPTP"`, `"RADIUS"`, `"REAL_MEDIA"`, `"RTSP"`, `"VNC"`, `"WHOIS"`, `"KERBEROS_SEC"`, `"TACACS"`, `"SNMPTRAP"`, `"NMAP"`, `"RSYNC"`, `"L2TP"`, `"HTTP_PROXY"`, `"PC_ANYWHERE"`, `"MSN"`, `"ECHO"`, `"AIM"`, `"IDENT"`, `"YMSG"`, `"SCCP"`, `"MGCP_UA"`, `"MGCP_CA"`, `"VDO_LIVE"`, `"OPENVPN"`, `"TFTP"`, `"FTPS_IMPLICIT"`, `"ZSCALER_PROXY_NW_SERVICES"`, `"GRE_PROTOCOL"`, `"ESP_PROTOCOL"`, `"DHCP"`
        """
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "type")

