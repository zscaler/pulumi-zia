# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FirewallFilteringDestinationGroupsArgs', 'FirewallFilteringDestinationGroups']

@pulumi.input_type
class FirewallFilteringDestinationGroupsArgs:
    def __init__(__self__, *,
                 addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 countries: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_categories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallFilteringDestinationGroups resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] addresses: Destination IP addresses within the group
        :param pulumi.Input[Sequence[pulumi.Input[str]]] countries: Destination countries for which the rule is applicable. If not set, the rule is not restricted to specific destination
               countries.
        :param pulumi.Input[str] description: Additional information about the destination IP group
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_categories: List of URL categories for which rule must be applied
        :param pulumi.Input[str] name: Destination IP group name
        :param pulumi.Input[str] type: Destination IP group type (i.e., the group can contain destination IP addresses or FQDNs)
        """
        if addresses is not None:
            pulumi.set(__self__, "addresses", addresses)
        if countries is not None:
            pulumi.set(__self__, "countries", countries)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_categories is not None:
            pulumi.set(__self__, "ip_categories", ip_categories)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Destination IP addresses within the group
        """
        return pulumi.get(self, "addresses")

    @addresses.setter
    def addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "addresses", value)

    @property
    @pulumi.getter
    def countries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Destination countries for which the rule is applicable. If not set, the rule is not restricted to specific destination
        countries.
        """
        return pulumi.get(self, "countries")

    @countries.setter
    def countries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "countries", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Additional information about the destination IP group
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipCategories")
    def ip_categories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of URL categories for which rule must be applied
        """
        return pulumi.get(self, "ip_categories")

    @ip_categories.setter
    def ip_categories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_categories", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Destination IP group name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Destination IP group type (i.e., the group can contain destination IP addresses or FQDNs)
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _FirewallFilteringDestinationGroupsState:
    def __init__(__self__, *,
                 addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 countries: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[int]] = None,
                 ip_categories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallFilteringDestinationGroups resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] addresses: Destination IP addresses within the group
        :param pulumi.Input[Sequence[pulumi.Input[str]]] countries: Destination countries for which the rule is applicable. If not set, the rule is not restricted to specific destination
               countries.
        :param pulumi.Input[str] description: Additional information about the destination IP group
        :param pulumi.Input[int] group_id: Unique identifer for the destination IP group
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_categories: List of URL categories for which rule must be applied
        :param pulumi.Input[str] name: Destination IP group name
        :param pulumi.Input[str] type: Destination IP group type (i.e., the group can contain destination IP addresses or FQDNs)
        """
        if addresses is not None:
            pulumi.set(__self__, "addresses", addresses)
        if countries is not None:
            pulumi.set(__self__, "countries", countries)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if ip_categories is not None:
            pulumi.set(__self__, "ip_categories", ip_categories)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Destination IP addresses within the group
        """
        return pulumi.get(self, "addresses")

    @addresses.setter
    def addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "addresses", value)

    @property
    @pulumi.getter
    def countries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Destination countries for which the rule is applicable. If not set, the rule is not restricted to specific destination
        countries.
        """
        return pulumi.get(self, "countries")

    @countries.setter
    def countries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "countries", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Additional information about the destination IP group
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[int]]:
        """
        Unique identifer for the destination IP group
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="ipCategories")
    def ip_categories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of URL categories for which rule must be applied
        """
        return pulumi.get(self, "ip_categories")

    @ip_categories.setter
    def ip_categories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_categories", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Destination IP group name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Destination IP group type (i.e., the group can contain destination IP addresses or FQDNs)
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class FirewallFilteringDestinationGroups(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 countries: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_categories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The **zia_firewall_filtering_destination_groups** resource allows the creation and management of ZIA Cloud Firewall IP destination groups in the Zscaler Internet Access. This resource can then be associated with a ZIA cloud firewall filtering rule.

        ## Example Usage

        ```python
        import pulumi
        import zscaler_pulumi_zia as zia

        # IP Destination Group of Type DSTN_FQDN
        dstn_fqdn = zia.FirewallFilteringDestinationGroups("dstnFqdn",
            addresses=[
                "test1.acme.com",
                "test2.acme.com",
                "test3.acme.com",
            ],
            description="Example Destination FQDN",
            type="DSTN_FQDN")
        ```

        ```python
        import pulumi
        import zscaler_pulumi_zia as zia

        # IP Destination Group of Type DSTN_IP
        example_dstn_ip = zia.FirewallFilteringDestinationGroups("exampleDstnIp",
            addresses=[
                "3.217.228.0-3.217.231.255",
                "3.235.112.0-3.235.119.255",
                "52.23.61.0-52.23.62.25",
                "35.80.88.0-35.80.95.255",
            ],
            description="Example Destination IP",
            type="DSTN_IP")
        ```

        ```python
        import pulumi
        import zscaler_pulumi_zia as zia

        # IP Destination Group of Type DSTN_DOMAIN
        example_dstn_domain = zia.FirewallFilteringDestinationGroups("exampleDstnDomain",
            addresses=[
                "acme.com",
                "acme1.com",
            ],
            description="Example Destination Domain",
            type="DSTN_DOMAIN")
        ```

        ```python
        import pulumi
        import zscaler_pulumi_zia as zia

        # IP Destination Group of Type DSTN_OTHER
        example_dstn_other = zia.FirewallFilteringDestinationGroups("exampleDstnOther",
            countries=["COUNTRY_CA"],
            description="Example Destination Other",
            ip_categories=[
                "CUSTOM_01",
                "CUSTOM_02",
            ],
            type="DSTN_OTHER")
        ```

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        **zia_firewall_filtering_destination_groups** can be imported by using `<GROUP_ID>` or `<GROUP_NAME>` as the import ID.

        For example:

        ```sh
        $ pulumi import zia:index/firewallFilteringDestinationGroups:FirewallFilteringDestinationGroups example <group_id>
        ```

        or

        ```sh
        $ pulumi import zia:index/firewallFilteringDestinationGroups:FirewallFilteringDestinationGroups example <group_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] addresses: Destination IP addresses within the group
        :param pulumi.Input[Sequence[pulumi.Input[str]]] countries: Destination countries for which the rule is applicable. If not set, the rule is not restricted to specific destination
               countries.
        :param pulumi.Input[str] description: Additional information about the destination IP group
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_categories: List of URL categories for which rule must be applied
        :param pulumi.Input[str] name: Destination IP group name
        :param pulumi.Input[str] type: Destination IP group type (i.e., the group can contain destination IP addresses or FQDNs)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FirewallFilteringDestinationGroupsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The **zia_firewall_filtering_destination_groups** resource allows the creation and management of ZIA Cloud Firewall IP destination groups in the Zscaler Internet Access. This resource can then be associated with a ZIA cloud firewall filtering rule.

        ## Example Usage

        ```python
        import pulumi
        import zscaler_pulumi_zia as zia

        # IP Destination Group of Type DSTN_FQDN
        dstn_fqdn = zia.FirewallFilteringDestinationGroups("dstnFqdn",
            addresses=[
                "test1.acme.com",
                "test2.acme.com",
                "test3.acme.com",
            ],
            description="Example Destination FQDN",
            type="DSTN_FQDN")
        ```

        ```python
        import pulumi
        import zscaler_pulumi_zia as zia

        # IP Destination Group of Type DSTN_IP
        example_dstn_ip = zia.FirewallFilteringDestinationGroups("exampleDstnIp",
            addresses=[
                "3.217.228.0-3.217.231.255",
                "3.235.112.0-3.235.119.255",
                "52.23.61.0-52.23.62.25",
                "35.80.88.0-35.80.95.255",
            ],
            description="Example Destination IP",
            type="DSTN_IP")
        ```

        ```python
        import pulumi
        import zscaler_pulumi_zia as zia

        # IP Destination Group of Type DSTN_DOMAIN
        example_dstn_domain = zia.FirewallFilteringDestinationGroups("exampleDstnDomain",
            addresses=[
                "acme.com",
                "acme1.com",
            ],
            description="Example Destination Domain",
            type="DSTN_DOMAIN")
        ```

        ```python
        import pulumi
        import zscaler_pulumi_zia as zia

        # IP Destination Group of Type DSTN_OTHER
        example_dstn_other = zia.FirewallFilteringDestinationGroups("exampleDstnOther",
            countries=["COUNTRY_CA"],
            description="Example Destination Other",
            ip_categories=[
                "CUSTOM_01",
                "CUSTOM_02",
            ],
            type="DSTN_OTHER")
        ```

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        **zia_firewall_filtering_destination_groups** can be imported by using `<GROUP_ID>` or `<GROUP_NAME>` as the import ID.

        For example:

        ```sh
        $ pulumi import zia:index/firewallFilteringDestinationGroups:FirewallFilteringDestinationGroups example <group_id>
        ```

        or

        ```sh
        $ pulumi import zia:index/firewallFilteringDestinationGroups:FirewallFilteringDestinationGroups example <group_name>
        ```

        :param str resource_name: The name of the resource.
        :param FirewallFilteringDestinationGroupsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallFilteringDestinationGroupsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 countries: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_categories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallFilteringDestinationGroupsArgs.__new__(FirewallFilteringDestinationGroupsArgs)

            __props__.__dict__["addresses"] = addresses
            __props__.__dict__["countries"] = countries
            __props__.__dict__["description"] = description
            __props__.__dict__["ip_categories"] = ip_categories
            __props__.__dict__["name"] = name
            __props__.__dict__["type"] = type
            __props__.__dict__["group_id"] = None
        super(FirewallFilteringDestinationGroups, __self__).__init__(
            'zia:index/firewallFilteringDestinationGroups:FirewallFilteringDestinationGroups',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            countries: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            group_id: Optional[pulumi.Input[int]] = None,
            ip_categories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'FirewallFilteringDestinationGroups':
        """
        Get an existing FirewallFilteringDestinationGroups resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] addresses: Destination IP addresses within the group
        :param pulumi.Input[Sequence[pulumi.Input[str]]] countries: Destination countries for which the rule is applicable. If not set, the rule is not restricted to specific destination
               countries.
        :param pulumi.Input[str] description: Additional information about the destination IP group
        :param pulumi.Input[int] group_id: Unique identifer for the destination IP group
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_categories: List of URL categories for which rule must be applied
        :param pulumi.Input[str] name: Destination IP group name
        :param pulumi.Input[str] type: Destination IP group type (i.e., the group can contain destination IP addresses or FQDNs)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallFilteringDestinationGroupsState.__new__(_FirewallFilteringDestinationGroupsState)

        __props__.__dict__["addresses"] = addresses
        __props__.__dict__["countries"] = countries
        __props__.__dict__["description"] = description
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["ip_categories"] = ip_categories
        __props__.__dict__["name"] = name
        __props__.__dict__["type"] = type
        return FirewallFilteringDestinationGroups(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        Destination IP addresses within the group
        """
        return pulumi.get(self, "addresses")

    @property
    @pulumi.getter
    def countries(self) -> pulumi.Output[Sequence[str]]:
        """
        Destination countries for which the rule is applicable. If not set, the rule is not restricted to specific destination
        countries.
        """
        return pulumi.get(self, "countries")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Additional information about the destination IP group
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[int]:
        """
        Unique identifer for the destination IP group
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="ipCategories")
    def ip_categories(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of URL categories for which rule must be applied
        """
        return pulumi.get(self, "ip_categories")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Destination IP group name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Destination IP group type (i.e., the group can contain destination IP addresses or FQDNs)
        """
        return pulumi.get(self, "type")

