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
    'FirewallFilteringNetworkServicesDestTcpPortArgs',
    'FirewallFilteringNetworkServicesDestUdpPortArgs',
    'FirewallFilteringNetworkServicesSrcTcpPortArgs',
    'FirewallFilteringNetworkServicesSrcUdpPortArgs',
    'FirewallFilteringRuleAppServiceGroupsArgs',
    'FirewallFilteringRuleAppServicesArgs',
    'FirewallFilteringRuleDepartmentsArgs',
    'FirewallFilteringRuleDestIpGroupsArgs',
    'FirewallFilteringRuleGroupsArgs',
    'FirewallFilteringRuleLabelsArgs',
    'FirewallFilteringRuleLastModifiedByArgs',
    'FirewallFilteringRuleLocationGroupsArgs',
    'FirewallFilteringRuleLocationsArgs',
    'FirewallFilteringRuleNwApplicationGroupsArgs',
    'FirewallFilteringRuleNwServiceGroupsArgs',
    'FirewallFilteringRuleNwServicesArgs',
    'FirewallFilteringRuleSrcIpGroupsArgs',
    'FirewallFilteringRuleTimeWindowsArgs',
    'FirewallFilteringRuleUsersArgs',
    'FirewallFilteringServiceGroupsServiceArgs',
]

@pulumi.input_type
class FirewallFilteringNetworkServicesDestTcpPortArgs:
    def __init__(__self__, *,
                 end: Optional[pulumi.Input[int]] = None,
                 start: Optional[pulumi.Input[int]] = None):
        if end is not None:
            pulumi.set(__self__, "end", end)
        if start is not None:
            pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def end(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "end")

    @end.setter
    def end(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "end", value)

    @property
    @pulumi.getter
    def start(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start", value)


@pulumi.input_type
class FirewallFilteringNetworkServicesDestUdpPortArgs:
    def __init__(__self__, *,
                 end: Optional[pulumi.Input[int]] = None,
                 start: Optional[pulumi.Input[int]] = None):
        if end is not None:
            pulumi.set(__self__, "end", end)
        if start is not None:
            pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def end(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "end")

    @end.setter
    def end(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "end", value)

    @property
    @pulumi.getter
    def start(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start", value)


@pulumi.input_type
class FirewallFilteringNetworkServicesSrcTcpPortArgs:
    def __init__(__self__, *,
                 end: Optional[pulumi.Input[int]] = None,
                 start: Optional[pulumi.Input[int]] = None):
        if end is not None:
            pulumi.set(__self__, "end", end)
        if start is not None:
            pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def end(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "end")

    @end.setter
    def end(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "end", value)

    @property
    @pulumi.getter
    def start(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start", value)


@pulumi.input_type
class FirewallFilteringNetworkServicesSrcUdpPortArgs:
    def __init__(__self__, *,
                 end: Optional[pulumi.Input[int]] = None,
                 start: Optional[pulumi.Input[int]] = None):
        if end is not None:
            pulumi.set(__self__, "end", end)
        if start is not None:
            pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def end(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "end")

    @end.setter
    def end(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "end", value)

    @property
    @pulumi.getter
    def start(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start", value)


@pulumi.input_type
class FirewallFilteringRuleAppServiceGroupsArgs:
    def __init__(__self__, *,
                 ids: pulumi.Input[Sequence[pulumi.Input[int]]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "ids", value)


@pulumi.input_type
class FirewallFilteringRuleAppServicesArgs:
    def __init__(__self__, *,
                 ids: pulumi.Input[Sequence[pulumi.Input[int]]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "ids", value)


@pulumi.input_type
class FirewallFilteringRuleDepartmentsArgs:
    def __init__(__self__, *,
                 ids: pulumi.Input[Sequence[pulumi.Input[int]]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "ids", value)


@pulumi.input_type
class FirewallFilteringRuleDestIpGroupsArgs:
    def __init__(__self__, *,
                 ids: pulumi.Input[Sequence[pulumi.Input[int]]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "ids", value)


@pulumi.input_type
class FirewallFilteringRuleGroupsArgs:
    def __init__(__self__, *,
                 ids: pulumi.Input[Sequence[pulumi.Input[int]]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "ids", value)


@pulumi.input_type
class FirewallFilteringRuleLabelsArgs:
    def __init__(__self__, *,
                 ids: pulumi.Input[Sequence[pulumi.Input[int]]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "ids", value)


@pulumi.input_type
class FirewallFilteringRuleLastModifiedByArgs:
    def __init__(__self__, *,
                 extensions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 id: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] id: Identifier that uniquely identifies an entity
        """
        if extensions is not None:
            pulumi.set(__self__, "extensions", extensions)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def extensions(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "extensions")

    @extensions.setter
    def extensions(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "extensions", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class FirewallFilteringRuleLocationGroupsArgs:
    def __init__(__self__, *,
                 ids: pulumi.Input[Sequence[pulumi.Input[int]]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "ids", value)


@pulumi.input_type
class FirewallFilteringRuleLocationsArgs:
    def __init__(__self__, *,
                 ids: pulumi.Input[Sequence[pulumi.Input[int]]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "ids", value)


@pulumi.input_type
class FirewallFilteringRuleNwApplicationGroupsArgs:
    def __init__(__self__, *,
                 ids: pulumi.Input[Sequence[pulumi.Input[int]]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "ids", value)


@pulumi.input_type
class FirewallFilteringRuleNwServiceGroupsArgs:
    def __init__(__self__, *,
                 ids: pulumi.Input[Sequence[pulumi.Input[int]]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "ids", value)


@pulumi.input_type
class FirewallFilteringRuleNwServicesArgs:
    def __init__(__self__, *,
                 ids: pulumi.Input[Sequence[pulumi.Input[int]]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "ids", value)


@pulumi.input_type
class FirewallFilteringRuleSrcIpGroupsArgs:
    def __init__(__self__, *,
                 ids: pulumi.Input[Sequence[pulumi.Input[int]]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "ids", value)


@pulumi.input_type
class FirewallFilteringRuleTimeWindowsArgs:
    def __init__(__self__, *,
                 ids: pulumi.Input[Sequence[pulumi.Input[int]]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "ids", value)


@pulumi.input_type
class FirewallFilteringRuleUsersArgs:
    def __init__(__self__, *,
                 ids: pulumi.Input[Sequence[pulumi.Input[int]]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[int]]] ids: Identifier that uniquely identifies an entity
        """
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "ids", value)


@pulumi.input_type
class FirewallFilteringServiceGroupsServiceArgs:
    def __init__(__self__, *,
                 ids: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "ids", value)


