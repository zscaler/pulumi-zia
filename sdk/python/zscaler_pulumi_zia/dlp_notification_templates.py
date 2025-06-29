# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['DLPNotificationTemplatesArgs', 'DLPNotificationTemplates']

@pulumi.input_type
class DLPNotificationTemplatesArgs:
    def __init__(__self__, *,
                 html_message: pulumi.Input[builtins.str],
                 plain_text_message: pulumi.Input[builtins.str],
                 subject: pulumi.Input[builtins.str],
                 attach_content: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 tls_enabled: Optional[pulumi.Input[builtins.bool]] = None):
        """
        The set of arguments for constructing a DLPNotificationTemplates resource.
        :param pulumi.Input[builtins.str] html_message: The template for the HTML message body that must be displayed in the DLP notification email
        :param pulumi.Input[builtins.str] plain_text_message: The template for the plain text UTF-8 message body that must be displayed in the DLP notification email
        :param pulumi.Input[builtins.str] subject: The Subject line that is displayed within the DLP notification email
        :param pulumi.Input[builtins.bool] attach_content: f set to true, the content that is violation is attached to the DLP notification email
        :param pulumi.Input[builtins.str] name: The DLP notification template name
        :param pulumi.Input[builtins.bool] tls_enabled: If set to true, TLS will be enabled
        """
        pulumi.set(__self__, "html_message", html_message)
        pulumi.set(__self__, "plain_text_message", plain_text_message)
        pulumi.set(__self__, "subject", subject)
        if attach_content is not None:
            pulumi.set(__self__, "attach_content", attach_content)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tls_enabled is not None:
            pulumi.set(__self__, "tls_enabled", tls_enabled)

    @property
    @pulumi.getter(name="htmlMessage")
    def html_message(self) -> pulumi.Input[builtins.str]:
        """
        The template for the HTML message body that must be displayed in the DLP notification email
        """
        return pulumi.get(self, "html_message")

    @html_message.setter
    def html_message(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "html_message", value)

    @property
    @pulumi.getter(name="plainTextMessage")
    def plain_text_message(self) -> pulumi.Input[builtins.str]:
        """
        The template for the plain text UTF-8 message body that must be displayed in the DLP notification email
        """
        return pulumi.get(self, "plain_text_message")

    @plain_text_message.setter
    def plain_text_message(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "plain_text_message", value)

    @property
    @pulumi.getter
    def subject(self) -> pulumi.Input[builtins.str]:
        """
        The Subject line that is displayed within the DLP notification email
        """
        return pulumi.get(self, "subject")

    @subject.setter
    def subject(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "subject", value)

    @property
    @pulumi.getter(name="attachContent")
    def attach_content(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        f set to true, the content that is violation is attached to the DLP notification email
        """
        return pulumi.get(self, "attach_content")

    @attach_content.setter
    def attach_content(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "attach_content", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The DLP notification template name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If set to true, TLS will be enabled
        """
        return pulumi.get(self, "tls_enabled")

    @tls_enabled.setter
    def tls_enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "tls_enabled", value)


@pulumi.input_type
class _DLPNotificationTemplatesState:
    def __init__(__self__, *,
                 attach_content: Optional[pulumi.Input[builtins.bool]] = None,
                 html_message: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 plain_text_message: Optional[pulumi.Input[builtins.str]] = None,
                 subject: Optional[pulumi.Input[builtins.str]] = None,
                 template_id: Optional[pulumi.Input[builtins.int]] = None,
                 tls_enabled: Optional[pulumi.Input[builtins.bool]] = None):
        """
        Input properties used for looking up and filtering DLPNotificationTemplates resources.
        :param pulumi.Input[builtins.bool] attach_content: f set to true, the content that is violation is attached to the DLP notification email
        :param pulumi.Input[builtins.str] html_message: The template for the HTML message body that must be displayed in the DLP notification email
        :param pulumi.Input[builtins.str] name: The DLP notification template name
        :param pulumi.Input[builtins.str] plain_text_message: The template for the plain text UTF-8 message body that must be displayed in the DLP notification email
        :param pulumi.Input[builtins.str] subject: The Subject line that is displayed within the DLP notification email
        :param pulumi.Input[builtins.int] template_id: The unique identifier for a DLP notification template
        :param pulumi.Input[builtins.bool] tls_enabled: If set to true, TLS will be enabled
        """
        if attach_content is not None:
            pulumi.set(__self__, "attach_content", attach_content)
        if html_message is not None:
            pulumi.set(__self__, "html_message", html_message)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if plain_text_message is not None:
            pulumi.set(__self__, "plain_text_message", plain_text_message)
        if subject is not None:
            pulumi.set(__self__, "subject", subject)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)
        if tls_enabled is not None:
            pulumi.set(__self__, "tls_enabled", tls_enabled)

    @property
    @pulumi.getter(name="attachContent")
    def attach_content(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        f set to true, the content that is violation is attached to the DLP notification email
        """
        return pulumi.get(self, "attach_content")

    @attach_content.setter
    def attach_content(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "attach_content", value)

    @property
    @pulumi.getter(name="htmlMessage")
    def html_message(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The template for the HTML message body that must be displayed in the DLP notification email
        """
        return pulumi.get(self, "html_message")

    @html_message.setter
    def html_message(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "html_message", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The DLP notification template name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="plainTextMessage")
    def plain_text_message(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The template for the plain text UTF-8 message body that must be displayed in the DLP notification email
        """
        return pulumi.get(self, "plain_text_message")

    @plain_text_message.setter
    def plain_text_message(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "plain_text_message", value)

    @property
    @pulumi.getter
    def subject(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Subject line that is displayed within the DLP notification email
        """
        return pulumi.get(self, "subject")

    @subject.setter
    def subject(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "subject", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The unique identifier for a DLP notification template
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "template_id", value)

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If set to true, TLS will be enabled
        """
        return pulumi.get(self, "tls_enabled")

    @tls_enabled.setter
    def tls_enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "tls_enabled", value)


@pulumi.type_token("zia:index/dLPNotificationTemplates:DLPNotificationTemplates")
class DLPNotificationTemplates(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attach_content: Optional[pulumi.Input[builtins.bool]] = None,
                 html_message: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 plain_text_message: Optional[pulumi.Input[builtins.str]] = None,
                 subject: Optional[pulumi.Input[builtins.str]] = None,
                 tls_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        """
        * [Official documentation](https://help.zscaler.com/zia/about-dlp-notification-templates)
        * [API documentation](https://help.zscaler.com/zia/data-loss-prevention#/dlpNotificationTemplates-get)

        The **zia_dlp_notification_templates** resource allows the creation and management of ZIA DLP Notification Templates in the Zscaler Internet Access cloud or via the API.

        ## Example Usage

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        **zia_dlp_notification_templates** can be imported by using `<TEMPLATE ID>` or `<TEMPLATE NAME>` as the import ID.

        For example:

        ```sh
        $ pulumi import zia:index/dLPNotificationTemplates:DLPNotificationTemplates example <template_id>
        ```

        or

        ```sh
        $ pulumi import zia:index/dLPNotificationTemplates:DLPNotificationTemplates example <template_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] attach_content: f set to true, the content that is violation is attached to the DLP notification email
        :param pulumi.Input[builtins.str] html_message: The template for the HTML message body that must be displayed in the DLP notification email
        :param pulumi.Input[builtins.str] name: The DLP notification template name
        :param pulumi.Input[builtins.str] plain_text_message: The template for the plain text UTF-8 message body that must be displayed in the DLP notification email
        :param pulumi.Input[builtins.str] subject: The Subject line that is displayed within the DLP notification email
        :param pulumi.Input[builtins.bool] tls_enabled: If set to true, TLS will be enabled
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DLPNotificationTemplatesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [Official documentation](https://help.zscaler.com/zia/about-dlp-notification-templates)
        * [API documentation](https://help.zscaler.com/zia/data-loss-prevention#/dlpNotificationTemplates-get)

        The **zia_dlp_notification_templates** resource allows the creation and management of ZIA DLP Notification Templates in the Zscaler Internet Access cloud or via the API.

        ## Example Usage

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.

        Visit

        **zia_dlp_notification_templates** can be imported by using `<TEMPLATE ID>` or `<TEMPLATE NAME>` as the import ID.

        For example:

        ```sh
        $ pulumi import zia:index/dLPNotificationTemplates:DLPNotificationTemplates example <template_id>
        ```

        or

        ```sh
        $ pulumi import zia:index/dLPNotificationTemplates:DLPNotificationTemplates example <template_name>
        ```

        :param str resource_name: The name of the resource.
        :param DLPNotificationTemplatesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DLPNotificationTemplatesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attach_content: Optional[pulumi.Input[builtins.bool]] = None,
                 html_message: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 plain_text_message: Optional[pulumi.Input[builtins.str]] = None,
                 subject: Optional[pulumi.Input[builtins.str]] = None,
                 tls_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DLPNotificationTemplatesArgs.__new__(DLPNotificationTemplatesArgs)

            __props__.__dict__["attach_content"] = attach_content
            if html_message is None and not opts.urn:
                raise TypeError("Missing required property 'html_message'")
            __props__.__dict__["html_message"] = html_message
            __props__.__dict__["name"] = name
            if plain_text_message is None and not opts.urn:
                raise TypeError("Missing required property 'plain_text_message'")
            __props__.__dict__["plain_text_message"] = plain_text_message
            if subject is None and not opts.urn:
                raise TypeError("Missing required property 'subject'")
            __props__.__dict__["subject"] = subject
            __props__.__dict__["tls_enabled"] = tls_enabled
            __props__.__dict__["template_id"] = None
        super(DLPNotificationTemplates, __self__).__init__(
            'zia:index/dLPNotificationTemplates:DLPNotificationTemplates',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attach_content: Optional[pulumi.Input[builtins.bool]] = None,
            html_message: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            plain_text_message: Optional[pulumi.Input[builtins.str]] = None,
            subject: Optional[pulumi.Input[builtins.str]] = None,
            template_id: Optional[pulumi.Input[builtins.int]] = None,
            tls_enabled: Optional[pulumi.Input[builtins.bool]] = None) -> 'DLPNotificationTemplates':
        """
        Get an existing DLPNotificationTemplates resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] attach_content: f set to true, the content that is violation is attached to the DLP notification email
        :param pulumi.Input[builtins.str] html_message: The template for the HTML message body that must be displayed in the DLP notification email
        :param pulumi.Input[builtins.str] name: The DLP notification template name
        :param pulumi.Input[builtins.str] plain_text_message: The template for the plain text UTF-8 message body that must be displayed in the DLP notification email
        :param pulumi.Input[builtins.str] subject: The Subject line that is displayed within the DLP notification email
        :param pulumi.Input[builtins.int] template_id: The unique identifier for a DLP notification template
        :param pulumi.Input[builtins.bool] tls_enabled: If set to true, TLS will be enabled
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DLPNotificationTemplatesState.__new__(_DLPNotificationTemplatesState)

        __props__.__dict__["attach_content"] = attach_content
        __props__.__dict__["html_message"] = html_message
        __props__.__dict__["name"] = name
        __props__.__dict__["plain_text_message"] = plain_text_message
        __props__.__dict__["subject"] = subject
        __props__.__dict__["template_id"] = template_id
        __props__.__dict__["tls_enabled"] = tls_enabled
        return DLPNotificationTemplates(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attachContent")
    def attach_content(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        f set to true, the content that is violation is attached to the DLP notification email
        """
        return pulumi.get(self, "attach_content")

    @property
    @pulumi.getter(name="htmlMessage")
    def html_message(self) -> pulumi.Output[builtins.str]:
        """
        The template for the HTML message body that must be displayed in the DLP notification email
        """
        return pulumi.get(self, "html_message")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The DLP notification template name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="plainTextMessage")
    def plain_text_message(self) -> pulumi.Output[builtins.str]:
        """
        The template for the plain text UTF-8 message body that must be displayed in the DLP notification email
        """
        return pulumi.get(self, "plain_text_message")

    @property
    @pulumi.getter
    def subject(self) -> pulumi.Output[builtins.str]:
        """
        The Subject line that is displayed within the DLP notification email
        """
        return pulumi.get(self, "subject")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Output[builtins.int]:
        """
        The unique identifier for a DLP notification template
        """
        return pulumi.get(self, "template_id")

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        If set to true, TLS will be enabled
        """
        return pulumi.get(self, "tls_enabled")

