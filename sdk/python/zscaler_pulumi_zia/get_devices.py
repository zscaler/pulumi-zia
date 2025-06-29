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
    'GetDevicesResult',
    'AwaitableGetDevicesResult',
    'get_devices',
    'get_devices_output',
]

@pulumi.output_type
class GetDevicesResult:
    """
    A collection of values returned by getDevices.
    """
    def __init__(__self__, description=None, device_group_type=None, device_model=None, hostname=None, id=None, name=None, os_type=None, os_version=None, owner_name=None, owner_user_id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if device_group_type and not isinstance(device_group_type, str):
            raise TypeError("Expected argument 'device_group_type' to be a str")
        pulumi.set(__self__, "device_group_type", device_group_type)
        if device_model and not isinstance(device_model, str):
            raise TypeError("Expected argument 'device_model' to be a str")
        pulumi.set(__self__, "device_model", device_model)
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        pulumi.set(__self__, "hostname", hostname)
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if os_type and not isinstance(os_type, str):
            raise TypeError("Expected argument 'os_type' to be a str")
        pulumi.set(__self__, "os_type", os_type)
        if os_version and not isinstance(os_version, str):
            raise TypeError("Expected argument 'os_version' to be a str")
        pulumi.set(__self__, "os_version", os_version)
        if owner_name and not isinstance(owner_name, str):
            raise TypeError("Expected argument 'owner_name' to be a str")
        pulumi.set(__self__, "owner_name", owner_name)
        if owner_user_id and not isinstance(owner_user_id, int):
            raise TypeError("Expected argument 'owner_user_id' to be a int")
        pulumi.set(__self__, "owner_user_id", owner_user_id)

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        (String) The device's description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="deviceGroupType")
    def device_group_type(self) -> builtins.str:
        """
        (String) The device group type. i.e ``ZCC_OS``, ``NON_ZCC``, ``CBI``
        """
        return pulumi.get(self, "device_group_type")

    @property
    @pulumi.getter(name="deviceModel")
    def device_model(self) -> builtins.str:
        """
        (String) The device model.
        """
        return pulumi.get(self, "device_model")

    @property
    @pulumi.getter
    def hostname(self) -> builtins.str:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def id(self) -> builtins.int:
        """
        (String) The unique identifer for the device group.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        (String) The device name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> builtins.str:
        """
        (String) The operating system (OS). ``ANY``, ``OTHER_OS``, ``IOS``, ``ANDROID_OS``, ``WINDOWS_OS``, ``MAC_OS``, ``LINUX``
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="osVersion")
    def os_version(self) -> builtins.str:
        """
        (String) The operating system version.
        """
        return pulumi.get(self, "os_version")

    @property
    @pulumi.getter(name="ownerName")
    def owner_name(self) -> builtins.str:
        """
        (String) The device owner's user name.
        """
        return pulumi.get(self, "owner_name")

    @property
    @pulumi.getter(name="ownerUserId")
    def owner_user_id(self) -> builtins.int:
        """
        (int) The unique identifier of the device owner (i.e., user).
        """
        return pulumi.get(self, "owner_user_id")


class AwaitableGetDevicesResult(GetDevicesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDevicesResult(
            description=self.description,
            device_group_type=self.device_group_type,
            device_model=self.device_model,
            hostname=self.hostname,
            id=self.id,
            name=self.name,
            os_type=self.os_type,
            os_version=self.os_version,
            owner_name=self.owner_name,
            owner_user_id=self.owner_user_id)


def get_devices(device_group_type: Optional[builtins.str] = None,
                device_model: Optional[builtins.str] = None,
                hostname: Optional[builtins.str] = None,
                id: Optional[builtins.int] = None,
                name: Optional[builtins.str] = None,
                os_type: Optional[builtins.str] = None,
                os_version: Optional[builtins.str] = None,
                owner_name: Optional[builtins.str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDevicesResult:
    """
    * [Official documentation](https://help.zscaler.com/zia/device-groups#/deviceGroups-get)
    * [API documentation](https://help.zscaler.com/zia/device-groups#/deviceGroups-get)

    Use the **zia_devices** data source to get information about a device in the Zscaler Internet Access cloud or via the API. This data source can then be associated with resources such as: URL Filtering Rules

    ## Example Usage


    :param builtins.str device_group_type: (String) The device group type. i.e ``ZCC_OS``, ``NON_ZCC``, ``CBI``
    :param builtins.str device_model: (String) The device model.
    :param builtins.int id: The unique identifer for the devices.
    :param builtins.str name: The name of the devices to be exported.
    :param builtins.str os_type: (String) The operating system (OS). ``ANY``, ``OTHER_OS``, ``IOS``, ``ANDROID_OS``, ``WINDOWS_OS``, ``MAC_OS``, ``LINUX``
    :param builtins.str os_version: (String) The operating system version.
    :param builtins.str owner_name: (String) The device owner's user name.
    """
    __args__ = dict()
    __args__['deviceGroupType'] = device_group_type
    __args__['deviceModel'] = device_model
    __args__['hostname'] = hostname
    __args__['id'] = id
    __args__['name'] = name
    __args__['osType'] = os_type
    __args__['osVersion'] = os_version
    __args__['ownerName'] = owner_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zia:index/getDevices:getDevices', __args__, opts=opts, typ=GetDevicesResult).value

    return AwaitableGetDevicesResult(
        description=pulumi.get(__ret__, 'description'),
        device_group_type=pulumi.get(__ret__, 'device_group_type'),
        device_model=pulumi.get(__ret__, 'device_model'),
        hostname=pulumi.get(__ret__, 'hostname'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        os_type=pulumi.get(__ret__, 'os_type'),
        os_version=pulumi.get(__ret__, 'os_version'),
        owner_name=pulumi.get(__ret__, 'owner_name'),
        owner_user_id=pulumi.get(__ret__, 'owner_user_id'))
def get_devices_output(device_group_type: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                       device_model: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                       hostname: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                       id: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                       name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                       os_type: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                       os_version: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                       owner_name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDevicesResult]:
    """
    * [Official documentation](https://help.zscaler.com/zia/device-groups#/deviceGroups-get)
    * [API documentation](https://help.zscaler.com/zia/device-groups#/deviceGroups-get)

    Use the **zia_devices** data source to get information about a device in the Zscaler Internet Access cloud or via the API. This data source can then be associated with resources such as: URL Filtering Rules

    ## Example Usage


    :param builtins.str device_group_type: (String) The device group type. i.e ``ZCC_OS``, ``NON_ZCC``, ``CBI``
    :param builtins.str device_model: (String) The device model.
    :param builtins.int id: The unique identifer for the devices.
    :param builtins.str name: The name of the devices to be exported.
    :param builtins.str os_type: (String) The operating system (OS). ``ANY``, ``OTHER_OS``, ``IOS``, ``ANDROID_OS``, ``WINDOWS_OS``, ``MAC_OS``, ``LINUX``
    :param builtins.str os_version: (String) The operating system version.
    :param builtins.str owner_name: (String) The device owner's user name.
    """
    __args__ = dict()
    __args__['deviceGroupType'] = device_group_type
    __args__['deviceModel'] = device_model
    __args__['hostname'] = hostname
    __args__['id'] = id
    __args__['name'] = name
    __args__['osType'] = os_type
    __args__['osVersion'] = os_version
    __args__['ownerName'] = owner_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('zia:index/getDevices:getDevices', __args__, opts=opts, typ=GetDevicesResult)
    return __ret__.apply(lambda __response__: GetDevicesResult(
        description=pulumi.get(__response__, 'description'),
        device_group_type=pulumi.get(__response__, 'device_group_type'),
        device_model=pulumi.get(__response__, 'device_model'),
        hostname=pulumi.get(__response__, 'hostname'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        os_type=pulumi.get(__response__, 'os_type'),
        os_version=pulumi.get(__response__, 'os_version'),
        owner_name=pulumi.get(__response__, 'owner_name'),
        owner_user_id=pulumi.get(__response__, 'owner_user_id')))
