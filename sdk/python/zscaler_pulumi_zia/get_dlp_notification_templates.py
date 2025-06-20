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

__all__ = [
    'GetDLPNotificationTemplatesResult',
    'AwaitableGetDLPNotificationTemplatesResult',
    'get_dlp_notification_templates',
    'get_dlp_notification_templates_output',
]

@pulumi.output_type
class GetDLPNotificationTemplatesResult:
    """
    A collection of values returned by getDLPNotificationTemplates.
    """
    def __init__(__self__, attach_content=None, html_message=None, id=None, name=None, plain_text_message=None, subject=None, tls_enabled=None):
        if attach_content and not isinstance(attach_content, bool):
            raise TypeError("Expected argument 'attach_content' to be a bool")
        pulumi.set(__self__, "attach_content", attach_content)
        if html_message and not isinstance(html_message, str):
            raise TypeError("Expected argument 'html_message' to be a str")
        pulumi.set(__self__, "html_message", html_message)
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if plain_text_message and not isinstance(plain_text_message, str):
            raise TypeError("Expected argument 'plain_text_message' to be a str")
        pulumi.set(__self__, "plain_text_message", plain_text_message)
        if subject and not isinstance(subject, str):
            raise TypeError("Expected argument 'subject' to be a str")
        pulumi.set(__self__, "subject", subject)
        if tls_enabled and not isinstance(tls_enabled, bool):
            raise TypeError("Expected argument 'tls_enabled' to be a bool")
        pulumi.set(__self__, "tls_enabled", tls_enabled)

    @property
    @pulumi.getter(name="attachContent")
    def attach_content(self) -> builtins.bool:
        return pulumi.get(self, "attach_content")

    @property
    @pulumi.getter(name="htmlMessage")
    def html_message(self) -> builtins.str:
        return pulumi.get(self, "html_message")

    @property
    @pulumi.getter
    def id(self) -> builtins.int:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="plainTextMessage")
    def plain_text_message(self) -> builtins.str:
        return pulumi.get(self, "plain_text_message")

    @property
    @pulumi.getter
    def subject(self) -> builtins.str:
        return pulumi.get(self, "subject")

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> builtins.bool:
        return pulumi.get(self, "tls_enabled")


class AwaitableGetDLPNotificationTemplatesResult(GetDLPNotificationTemplatesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDLPNotificationTemplatesResult(
            attach_content=self.attach_content,
            html_message=self.html_message,
            id=self.id,
            name=self.name,
            plain_text_message=self.plain_text_message,
            subject=self.subject,
            tls_enabled=self.tls_enabled)


def get_dlp_notification_templates(id: Optional[builtins.int] = None,
                                   name: Optional[builtins.str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDLPNotificationTemplatesResult:
    """
    * [Official documentation](https://help.zscaler.com/zia/about-dlp-notification-templates)
    * [API documentation](https://help.zscaler.com/zia/data-loss-prevention#/dlpNotificationTemplates-get)

    Use the **zia_dlp_notification_templates** data source to get information about a ZIA DLP Notification Templates in the Zscaler Internet Access cloud or via the API.

    ## Example Usage
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zia:index/getDLPNotificationTemplates:getDLPNotificationTemplates', __args__, opts=opts, typ=GetDLPNotificationTemplatesResult).value

    return AwaitableGetDLPNotificationTemplatesResult(
        attach_content=pulumi.get(__ret__, 'attach_content'),
        html_message=pulumi.get(__ret__, 'html_message'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        plain_text_message=pulumi.get(__ret__, 'plain_text_message'),
        subject=pulumi.get(__ret__, 'subject'),
        tls_enabled=pulumi.get(__ret__, 'tls_enabled'))
def get_dlp_notification_templates_output(id: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                                          name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDLPNotificationTemplatesResult]:
    """
    * [Official documentation](https://help.zscaler.com/zia/about-dlp-notification-templates)
    * [API documentation](https://help.zscaler.com/zia/data-loss-prevention#/dlpNotificationTemplates-get)

    Use the **zia_dlp_notification_templates** data source to get information about a ZIA DLP Notification Templates in the Zscaler Internet Access cloud or via the API.

    ## Example Usage
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('zia:index/getDLPNotificationTemplates:getDLPNotificationTemplates', __args__, opts=opts, typ=GetDLPNotificationTemplatesResult)
    return __ret__.apply(lambda __response__: GetDLPNotificationTemplatesResult(
        attach_content=pulumi.get(__response__, 'attach_content'),
        html_message=pulumi.get(__response__, 'html_message'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        plain_text_message=pulumi.get(__response__, 'plain_text_message'),
        subject=pulumi.get(__response__, 'subject'),
        tls_enabled=pulumi.get(__response__, 'tls_enabled')))
