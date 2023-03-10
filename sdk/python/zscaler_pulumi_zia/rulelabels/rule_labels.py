# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['RuleLabelsArgs', 'RuleLabels']

@pulumi.input_type
class RuleLabelsArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RuleLabels resource.
        :param pulumi.Input[str] description: The rule label description.
        :param pulumi.Input[str] name: The name of the devices to be created.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The rule label description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the devices to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _RuleLabelsState:
    def __init__(__self__, *,
                 created_bies: Optional[pulumi.Input[Sequence[pulumi.Input['RuleLabelsCreatedByArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 last_modified_bies: Optional[pulumi.Input[Sequence[pulumi.Input['RuleLabelsLastModifiedByArgs']]]] = None,
                 last_modified_time: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 referenced_rule_count: Optional[pulumi.Input[int]] = None,
                 rule_label_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering RuleLabels resources.
        :param pulumi.Input[Sequence[pulumi.Input['RuleLabelsCreatedByArgs']]] created_bies: The admin that created the rule label. This is a read-only field. Ignored by PUT requests.
        :param pulumi.Input[str] description: The rule label description.
        :param pulumi.Input[Sequence[pulumi.Input['RuleLabelsLastModifiedByArgs']]] last_modified_bies: The admin that modified the rule label last. This is a read-only field. Ignored by PUT requests.
        :param pulumi.Input[int] last_modified_time: Timestamp when the rule lable was last modified. This is a read-only field. Ignored by PUT and DELETE requests.
        :param pulumi.Input[str] name: The name of the devices to be created.
        """
        if created_bies is not None:
            pulumi.set(__self__, "created_bies", created_bies)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if last_modified_bies is not None:
            pulumi.set(__self__, "last_modified_bies", last_modified_bies)
        if last_modified_time is not None:
            pulumi.set(__self__, "last_modified_time", last_modified_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if referenced_rule_count is not None:
            pulumi.set(__self__, "referenced_rule_count", referenced_rule_count)
        if rule_label_id is not None:
            pulumi.set(__self__, "rule_label_id", rule_label_id)

    @property
    @pulumi.getter(name="createdBies")
    def created_bies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RuleLabelsCreatedByArgs']]]]:
        """
        The admin that created the rule label. This is a read-only field. Ignored by PUT requests.
        """
        return pulumi.get(self, "created_bies")

    @created_bies.setter
    def created_bies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RuleLabelsCreatedByArgs']]]]):
        pulumi.set(self, "created_bies", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The rule label description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="lastModifiedBies")
    def last_modified_bies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RuleLabelsLastModifiedByArgs']]]]:
        """
        The admin that modified the rule label last. This is a read-only field. Ignored by PUT requests.
        """
        return pulumi.get(self, "last_modified_bies")

    @last_modified_bies.setter
    def last_modified_bies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RuleLabelsLastModifiedByArgs']]]]):
        pulumi.set(self, "last_modified_bies", value)

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> Optional[pulumi.Input[int]]:
        """
        Timestamp when the rule lable was last modified. This is a read-only field. Ignored by PUT and DELETE requests.
        """
        return pulumi.get(self, "last_modified_time")

    @last_modified_time.setter
    def last_modified_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "last_modified_time", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the devices to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="referencedRuleCount")
    def referenced_rule_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "referenced_rule_count")

    @referenced_rule_count.setter
    def referenced_rule_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "referenced_rule_count", value)

    @property
    @pulumi.getter(name="ruleLabelId")
    def rule_label_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "rule_label_id")

    @rule_label_id.setter
    def rule_label_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rule_label_id", value)


class RuleLabels(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The **zia_rule_labels** resource allows the creation and management of rule labels in the Zscaler Internet Access cloud or via the API. This resource can then be associated with resources such as: Firewall Rules and URL filtering rules

        ## Example Usage

        ```python
        import pulumi
        import zscaler_pulumi_zia as zia

        # ZIA Rule Labels Resource
        example = zia.rule_labels.RuleLabels("example", description="Example")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The rule label description.
        :param pulumi.Input[str] name: The name of the devices to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RuleLabelsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The **zia_rule_labels** resource allows the creation and management of rule labels in the Zscaler Internet Access cloud or via the API. This resource can then be associated with resources such as: Firewall Rules and URL filtering rules

        ## Example Usage

        ```python
        import pulumi
        import zscaler_pulumi_zia as zia

        # ZIA Rule Labels Resource
        example = zia.rule_labels.RuleLabels("example", description="Example")
        ```

        :param str resource_name: The name of the resource.
        :param RuleLabelsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RuleLabelsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RuleLabelsArgs.__new__(RuleLabelsArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["created_bies"] = None
            __props__.__dict__["last_modified_bies"] = None
            __props__.__dict__["last_modified_time"] = None
            __props__.__dict__["referenced_rule_count"] = None
            __props__.__dict__["rule_label_id"] = None
        super(RuleLabels, __self__).__init__(
            'zia:RuleLabels/ruleLabels:RuleLabels',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_bies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleLabelsCreatedByArgs']]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            last_modified_bies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleLabelsLastModifiedByArgs']]]]] = None,
            last_modified_time: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            referenced_rule_count: Optional[pulumi.Input[int]] = None,
            rule_label_id: Optional[pulumi.Input[int]] = None) -> 'RuleLabels':
        """
        Get an existing RuleLabels resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleLabelsCreatedByArgs']]]] created_bies: The admin that created the rule label. This is a read-only field. Ignored by PUT requests.
        :param pulumi.Input[str] description: The rule label description.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleLabelsLastModifiedByArgs']]]] last_modified_bies: The admin that modified the rule label last. This is a read-only field. Ignored by PUT requests.
        :param pulumi.Input[int] last_modified_time: Timestamp when the rule lable was last modified. This is a read-only field. Ignored by PUT and DELETE requests.
        :param pulumi.Input[str] name: The name of the devices to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RuleLabelsState.__new__(_RuleLabelsState)

        __props__.__dict__["created_bies"] = created_bies
        __props__.__dict__["description"] = description
        __props__.__dict__["last_modified_bies"] = last_modified_bies
        __props__.__dict__["last_modified_time"] = last_modified_time
        __props__.__dict__["name"] = name
        __props__.__dict__["referenced_rule_count"] = referenced_rule_count
        __props__.__dict__["rule_label_id"] = rule_label_id
        return RuleLabels(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdBies")
    def created_bies(self) -> pulumi.Output[Sequence['outputs.RuleLabelsCreatedBy']]:
        """
        The admin that created the rule label. This is a read-only field. Ignored by PUT requests.
        """
        return pulumi.get(self, "created_bies")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The rule label description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lastModifiedBies")
    def last_modified_bies(self) -> pulumi.Output[Sequence['outputs.RuleLabelsLastModifiedBy']]:
        """
        The admin that modified the rule label last. This is a read-only field. Ignored by PUT requests.
        """
        return pulumi.get(self, "last_modified_bies")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> pulumi.Output[int]:
        """
        Timestamp when the rule lable was last modified. This is a read-only field. Ignored by PUT and DELETE requests.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the devices to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="referencedRuleCount")
    def referenced_rule_count(self) -> pulumi.Output[int]:
        return pulumi.get(self, "referenced_rule_count")

    @property
    @pulumi.getter(name="ruleLabelId")
    def rule_label_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "rule_label_id")

