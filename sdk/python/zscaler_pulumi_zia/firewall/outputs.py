# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'FirewallFilteringNetworkServicesDestTcpPort',
    'FirewallFilteringNetworkServicesDestUdpPort',
    'FirewallFilteringNetworkServicesSrcTcpPort',
    'FirewallFilteringNetworkServicesSrcUdpPort',
    'FirewallFilteringRuleAppServiceGroups',
    'FirewallFilteringRuleAppServices',
    'FirewallFilteringRuleDepartments',
    'FirewallFilteringRuleDestIpGroups',
    'FirewallFilteringRuleGroups',
    'FirewallFilteringRuleLabels',
    'FirewallFilteringRuleLastModifiedBy',
    'FirewallFilteringRuleLocationGroups',
    'FirewallFilteringRuleLocations',
    'FirewallFilteringRuleNwApplicationGroups',
    'FirewallFilteringRuleNwServiceGroups',
    'FirewallFilteringRuleNwServices',
    'FirewallFilteringRuleSrcIpGroups',
    'FirewallFilteringRuleTimeWindows',
    'FirewallFilteringRuleUsers',
    'FirewallFilteringServiceGroupsService',
    'GetFirewallFilteringNetworkServiceGroupsServiceResult',
    'GetFirewallFilteringNetworkServicesDestTcpPortResult',
    'GetFirewallFilteringNetworkServicesDestUdpPortResult',
    'GetFirewallFilteringNetworkServicesSrcTcpPortResult',
    'GetFirewallFilteringNetworkServicesSrcUdpPortResult',
    'GetFirewallFilteringRuleAppServiceResult',
    'GetFirewallFilteringRuleAppServiceGroupResult',
    'GetFirewallFilteringRuleDepartmentResult',
    'GetFirewallFilteringRuleGroupResult',
    'GetFirewallFilteringRuleLabelResult',
    'GetFirewallFilteringRuleLastModifiedByResult',
    'GetFirewallFilteringRuleLocationResult',
    'GetFirewallFilteringRuleLocationGroupResult',
    'GetFirewallFilteringRuleNwApplicationGroupResult',
    'GetFirewallFilteringRuleNwServiceResult',
    'GetFirewallFilteringRuleNwServiceGroupResult',
    'GetFirewallFilteringRuleTimeWindowResult',
    'GetFirewallFilteringRuleUserResult',
]

@pulumi.output_type
class FirewallFilteringNetworkServicesDestTcpPort(dict):
    def __init__(__self__, *,
                 end: Optional[int] = None,
                 start: Optional[int] = None):
        if end is not None:
            pulumi.set(__self__, "end", end)
        if start is not None:
            pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def end(self) -> Optional[int]:
        return pulumi.get(self, "end")

    @property
    @pulumi.getter
    def start(self) -> Optional[int]:
        return pulumi.get(self, "start")


@pulumi.output_type
class FirewallFilteringNetworkServicesDestUdpPort(dict):
    def __init__(__self__, *,
                 end: Optional[int] = None,
                 start: Optional[int] = None):
        if end is not None:
            pulumi.set(__self__, "end", end)
        if start is not None:
            pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def end(self) -> Optional[int]:
        return pulumi.get(self, "end")

    @property
    @pulumi.getter
    def start(self) -> Optional[int]:
        return pulumi.get(self, "start")


@pulumi.output_type
class FirewallFilteringNetworkServicesSrcTcpPort(dict):
    def __init__(__self__, *,
                 end: Optional[int] = None,
                 start: Optional[int] = None):
        if end is not None:
            pulumi.set(__self__, "end", end)
        if start is not None:
            pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def end(self) -> Optional[int]:
        return pulumi.get(self, "end")

    @property
    @pulumi.getter
    def start(self) -> Optional[int]:
        return pulumi.get(self, "start")


@pulumi.output_type
class FirewallFilteringNetworkServicesSrcUdpPort(dict):
    def __init__(__self__, *,
                 end: Optional[int] = None,
                 start: Optional[int] = None):
        if end is not None:
            pulumi.set(__self__, "end", end)
        if start is not None:
            pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def end(self) -> Optional[int]:
        return pulumi.get(self, "end")

    @property
    @pulumi.getter
    def start(self) -> Optional[int]:
        return pulumi.get(self, "start")


@pulumi.output_type
class FirewallFilteringRuleAppServiceGroups(dict):
    def __init__(__self__, *,
                 ids: Sequence[int]):
        """
        :param Sequence[int] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Sequence[int]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")


@pulumi.output_type
class FirewallFilteringRuleAppServices(dict):
    def __init__(__self__, *,
                 ids: Sequence[int]):
        """
        :param Sequence[int] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Sequence[int]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")


@pulumi.output_type
class FirewallFilteringRuleDepartments(dict):
    def __init__(__self__, *,
                 ids: Sequence[int]):
        """
        :param Sequence[int] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Sequence[int]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")


@pulumi.output_type
class FirewallFilteringRuleDestIpGroups(dict):
    def __init__(__self__, *,
                 ids: Sequence[int]):
        """
        :param Sequence[int] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Sequence[int]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")


@pulumi.output_type
class FirewallFilteringRuleGroups(dict):
    def __init__(__self__, *,
                 ids: Sequence[int]):
        """
        :param Sequence[int] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Sequence[int]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")


@pulumi.output_type
class FirewallFilteringRuleLabels(dict):
    def __init__(__self__, *,
                 ids: Sequence[int]):
        """
        :param Sequence[int] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Sequence[int]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")


@pulumi.output_type
class FirewallFilteringRuleLastModifiedBy(dict):
    def __init__(__self__, *,
                 extensions: Optional[Mapping[str, str]] = None,
                 id: Optional[int] = None):
        """
        :param int id: Identifier that uniquely identifies an entity
        """
        if extensions is not None:
            pulumi.set(__self__, "extensions", extensions)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def extensions(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class FirewallFilteringRuleLocationGroups(dict):
    def __init__(__self__, *,
                 ids: Sequence[int]):
        """
        :param Sequence[int] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Sequence[int]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")


@pulumi.output_type
class FirewallFilteringRuleLocations(dict):
    def __init__(__self__, *,
                 ids: Sequence[int]):
        """
        :param Sequence[int] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Sequence[int]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")


@pulumi.output_type
class FirewallFilteringRuleNwApplicationGroups(dict):
    def __init__(__self__, *,
                 ids: Sequence[int]):
        """
        :param Sequence[int] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Sequence[int]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")


@pulumi.output_type
class FirewallFilteringRuleNwServiceGroups(dict):
    def __init__(__self__, *,
                 ids: Sequence[int]):
        """
        :param Sequence[int] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Sequence[int]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")


@pulumi.output_type
class FirewallFilteringRuleNwServices(dict):
    def __init__(__self__, *,
                 ids: Sequence[int]):
        """
        :param Sequence[int] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Sequence[int]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")


@pulumi.output_type
class FirewallFilteringRuleSrcIpGroups(dict):
    def __init__(__self__, *,
                 ids: Sequence[int]):
        """
        :param Sequence[int] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Sequence[int]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")


@pulumi.output_type
class FirewallFilteringRuleTimeWindows(dict):
    def __init__(__self__, *,
                 ids: Sequence[int]):
        """
        :param Sequence[int] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Sequence[int]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")


@pulumi.output_type
class FirewallFilteringRuleUsers(dict):
    def __init__(__self__, *,
                 ids: Sequence[int]):
        """
        :param Sequence[int] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Sequence[int]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")


@pulumi.output_type
class FirewallFilteringServiceGroupsService(dict):
    def __init__(__self__, *,
                 ids: Sequence[int]):
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Sequence[int]:
        return pulumi.get(self, "ids")


@pulumi.output_type
class GetFirewallFilteringNetworkServiceGroupsServiceResult(dict):
    def __init__(__self__, *,
                 description: str,
                 id: int,
                 is_name_l10n_tag: bool,
                 name: Optional[str] = None):
        """
        :param str description: (String)
        :param int id: The ID of the ip source group to be exported.
        :param bool is_name_l10n_tag: (Bool) - Default: false
        :param str name: The name of the ip source group to be exported.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "is_name_l10n_tag", is_name_l10n_tag)
        if name is not None:
            pulumi.set(__self__, "name", name)

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
        The ID of the ip source group to be exported.
        """
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
    def name(self) -> Optional[str]:
        """
        The name of the ip source group to be exported.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetFirewallFilteringNetworkServicesDestTcpPortResult(dict):
    def __init__(__self__, *,
                 end: int,
                 start: int):
        """
        :param int end: (Number)
        :param int start: (Number)
        """
        pulumi.set(__self__, "end", end)
        pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def end(self) -> int:
        """
        (Number)
        """
        return pulumi.get(self, "end")

    @property
    @pulumi.getter
    def start(self) -> int:
        """
        (Number)
        """
        return pulumi.get(self, "start")


@pulumi.output_type
class GetFirewallFilteringNetworkServicesDestUdpPortResult(dict):
    def __init__(__self__, *,
                 end: int,
                 start: int):
        """
        :param int end: (Number)
        :param int start: (Number)
        """
        pulumi.set(__self__, "end", end)
        pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def end(self) -> int:
        """
        (Number)
        """
        return pulumi.get(self, "end")

    @property
    @pulumi.getter
    def start(self) -> int:
        """
        (Number)
        """
        return pulumi.get(self, "start")


@pulumi.output_type
class GetFirewallFilteringNetworkServicesSrcTcpPortResult(dict):
    def __init__(__self__, *,
                 end: int,
                 start: int):
        """
        :param int end: (Number)
        :param int start: (Number)
        """
        pulumi.set(__self__, "end", end)
        pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def end(self) -> int:
        """
        (Number)
        """
        return pulumi.get(self, "end")

    @property
    @pulumi.getter
    def start(self) -> int:
        """
        (Number)
        """
        return pulumi.get(self, "start")


@pulumi.output_type
class GetFirewallFilteringNetworkServicesSrcUdpPortResult(dict):
    def __init__(__self__, *,
                 end: int,
                 start: int):
        """
        :param int end: (Number)
        :param int start: (Number)
        """
        pulumi.set(__self__, "end", end)
        pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def end(self) -> int:
        """
        (Number)
        """
        return pulumi.get(self, "end")

    @property
    @pulumi.getter
    def start(self) -> int:
        """
        (Number)
        """
        return pulumi.get(self, "start")


@pulumi.output_type
class GetFirewallFilteringRuleAppServiceResult(dict):
    def __init__(__self__, *,
                 extensions: Mapping[str, str],
                 id: int,
                 name: str):
        """
        :param Mapping[str, str] extensions: (Map of String)
        :param int id: Unique identifier for the Firewall Filtering policy rule
        :param str name: Name of the Firewall Filtering policy rule
        """
        pulumi.set(__self__, "extensions", extensions)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def extensions(self) -> Mapping[str, str]:
        """
        (Map of String)
        """
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Unique identifier for the Firewall Filtering policy rule
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the Firewall Filtering policy rule
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetFirewallFilteringRuleAppServiceGroupResult(dict):
    def __init__(__self__, *,
                 extensions: Mapping[str, str],
                 id: int,
                 name: str):
        """
        :param Mapping[str, str] extensions: (Map of String)
        :param int id: Unique identifier for the Firewall Filtering policy rule
        :param str name: Name of the Firewall Filtering policy rule
        """
        pulumi.set(__self__, "extensions", extensions)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def extensions(self) -> Mapping[str, str]:
        """
        (Map of String)
        """
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Unique identifier for the Firewall Filtering policy rule
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the Firewall Filtering policy rule
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetFirewallFilteringRuleDepartmentResult(dict):
    def __init__(__self__, *,
                 extensions: Mapping[str, str],
                 id: int,
                 name: str):
        """
        :param Mapping[str, str] extensions: (Map of String)
        :param int id: Unique identifier for the Firewall Filtering policy rule
        :param str name: Name of the Firewall Filtering policy rule
        """
        pulumi.set(__self__, "extensions", extensions)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def extensions(self) -> Mapping[str, str]:
        """
        (Map of String)
        """
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Unique identifier for the Firewall Filtering policy rule
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the Firewall Filtering policy rule
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetFirewallFilteringRuleGroupResult(dict):
    def __init__(__self__, *,
                 extensions: Mapping[str, str],
                 id: int,
                 name: str):
        """
        :param Mapping[str, str] extensions: (Map of String)
        :param int id: Unique identifier for the Firewall Filtering policy rule
        :param str name: Name of the Firewall Filtering policy rule
        """
        pulumi.set(__self__, "extensions", extensions)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def extensions(self) -> Mapping[str, str]:
        """
        (Map of String)
        """
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Unique identifier for the Firewall Filtering policy rule
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the Firewall Filtering policy rule
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetFirewallFilteringRuleLabelResult(dict):
    def __init__(__self__, *,
                 extensions: Mapping[str, str],
                 id: int,
                 name: str):
        """
        :param Mapping[str, str] extensions: (Map of String)
        :param int id: Unique identifier for the Firewall Filtering policy rule
        :param str name: Name of the Firewall Filtering policy rule
        """
        pulumi.set(__self__, "extensions", extensions)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def extensions(self) -> Mapping[str, str]:
        """
        (Map of String)
        """
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Unique identifier for the Firewall Filtering policy rule
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the Firewall Filtering policy rule
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetFirewallFilteringRuleLastModifiedByResult(dict):
    def __init__(__self__, *,
                 extensions: Mapping[str, str],
                 id: int,
                 name: str):
        """
        :param Mapping[str, str] extensions: (Map of String)
        :param int id: Unique identifier for the Firewall Filtering policy rule
        :param str name: Name of the Firewall Filtering policy rule
        """
        pulumi.set(__self__, "extensions", extensions)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def extensions(self) -> Mapping[str, str]:
        """
        (Map of String)
        """
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Unique identifier for the Firewall Filtering policy rule
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the Firewall Filtering policy rule
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetFirewallFilteringRuleLocationResult(dict):
    def __init__(__self__, *,
                 extensions: Mapping[str, str],
                 id: int,
                 name: str):
        """
        :param Mapping[str, str] extensions: (Map of String)
        :param int id: Unique identifier for the Firewall Filtering policy rule
        :param str name: Name of the Firewall Filtering policy rule
        """
        pulumi.set(__self__, "extensions", extensions)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def extensions(self) -> Mapping[str, str]:
        """
        (Map of String)
        """
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Unique identifier for the Firewall Filtering policy rule
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the Firewall Filtering policy rule
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetFirewallFilteringRuleLocationGroupResult(dict):
    def __init__(__self__, *,
                 extensions: Mapping[str, str],
                 id: int,
                 name: str):
        """
        :param Mapping[str, str] extensions: (Map of String)
        :param int id: Unique identifier for the Firewall Filtering policy rule
        :param str name: Name of the Firewall Filtering policy rule
        """
        pulumi.set(__self__, "extensions", extensions)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def extensions(self) -> Mapping[str, str]:
        """
        (Map of String)
        """
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Unique identifier for the Firewall Filtering policy rule
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the Firewall Filtering policy rule
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetFirewallFilteringRuleNwApplicationGroupResult(dict):
    def __init__(__self__, *,
                 extensions: Mapping[str, str],
                 id: int,
                 name: str):
        """
        :param Mapping[str, str] extensions: (Map of String)
        :param int id: Unique identifier for the Firewall Filtering policy rule
        :param str name: Name of the Firewall Filtering policy rule
        """
        pulumi.set(__self__, "extensions", extensions)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def extensions(self) -> Mapping[str, str]:
        """
        (Map of String)
        """
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Unique identifier for the Firewall Filtering policy rule
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the Firewall Filtering policy rule
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetFirewallFilteringRuleNwServiceResult(dict):
    def __init__(__self__, *,
                 extensions: Mapping[str, str],
                 id: int,
                 name: str):
        """
        :param Mapping[str, str] extensions: (Map of String)
        :param int id: Unique identifier for the Firewall Filtering policy rule
        :param str name: Name of the Firewall Filtering policy rule
        """
        pulumi.set(__self__, "extensions", extensions)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def extensions(self) -> Mapping[str, str]:
        """
        (Map of String)
        """
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Unique identifier for the Firewall Filtering policy rule
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the Firewall Filtering policy rule
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetFirewallFilteringRuleNwServiceGroupResult(dict):
    def __init__(__self__, *,
                 extensions: Mapping[str, str],
                 id: int,
                 name: str):
        """
        :param Mapping[str, str] extensions: (Map of String)
        :param int id: Unique identifier for the Firewall Filtering policy rule
        :param str name: Name of the Firewall Filtering policy rule
        """
        pulumi.set(__self__, "extensions", extensions)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def extensions(self) -> Mapping[str, str]:
        """
        (Map of String)
        """
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Unique identifier for the Firewall Filtering policy rule
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the Firewall Filtering policy rule
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetFirewallFilteringRuleTimeWindowResult(dict):
    def __init__(__self__, *,
                 extensions: Mapping[str, str],
                 id: int,
                 name: str):
        """
        :param Mapping[str, str] extensions: (Map of String)
        :param int id: Unique identifier for the Firewall Filtering policy rule
        :param str name: Name of the Firewall Filtering policy rule
        """
        pulumi.set(__self__, "extensions", extensions)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def extensions(self) -> Mapping[str, str]:
        """
        (Map of String)
        """
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Unique identifier for the Firewall Filtering policy rule
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the Firewall Filtering policy rule
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetFirewallFilteringRuleUserResult(dict):
    def __init__(__self__, *,
                 extensions: Mapping[str, str],
                 id: int,
                 name: str):
        """
        :param Mapping[str, str] extensions: (Map of String)
        :param int id: Unique identifier for the Firewall Filtering policy rule
        :param str name: Name of the Firewall Filtering policy rule
        """
        pulumi.set(__self__, "extensions", extensions)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def extensions(self) -> Mapping[str, str]:
        """
        (Map of String)
        """
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Unique identifier for the Firewall Filtering policy rule
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the Firewall Filtering policy rule
        """
        return pulumi.get(self, "name")


