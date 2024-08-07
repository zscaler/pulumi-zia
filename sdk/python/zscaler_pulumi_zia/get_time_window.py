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
    'GetTimeWindowResult',
    'AwaitableGetTimeWindowResult',
    'get_time_window',
    'get_time_window_output',
]

@pulumi.output_type
class GetTimeWindowResult:
    """
    A collection of values returned by getTimeWindow.
    """
    def __init__(__self__, day_of_weeks=None, end_time=None, id=None, name=None, start_time=None):
        if day_of_weeks and not isinstance(day_of_weeks, list):
            raise TypeError("Expected argument 'day_of_weeks' to be a list")
        pulumi.set(__self__, "day_of_weeks", day_of_weeks)
        if end_time and not isinstance(end_time, int):
            raise TypeError("Expected argument 'end_time' to be a int")
        pulumi.set(__self__, "end_time", end_time)
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if start_time and not isinstance(start_time, int):
            raise TypeError("Expected argument 'start_time' to be a int")
        pulumi.set(__self__, "start_time", start_time)

    @property
    @pulumi.getter(name="dayOfWeeks")
    def day_of_weeks(self) -> Sequence[str]:
        """
        (String). The supported values are:
        * `ANY` - (String)
        * `NONE` - (String)
        * `EVERYDAY` - (String)
        * `SUN` - (String)
        * `MON` - (String)
        * `TUE` - (String)
        * `WED` - (String)
        * `THU` - (String)
        * `FRI` - (String)
        * `SAT` - (String)
        """
        return pulumi.get(self, "day_of_weeks")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> int:
        """
        (String)
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def id(self) -> int:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> int:
        """
        (String)
        """
        return pulumi.get(self, "start_time")


class AwaitableGetTimeWindowResult(GetTimeWindowResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTimeWindowResult(
            day_of_weeks=self.day_of_weeks,
            end_time=self.end_time,
            id=self.id,
            name=self.name,
            start_time=self.start_time)


def get_time_window(name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTimeWindowResult:
    """
    Use the **zia_firewall_filtering_time_window** data source to get information about a time window option available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zia as zia

    work_hours = zia.get_time_window(name="Work hours")
    ```

    ```python
    import pulumi
    import pulumi_zia as zia

    weekends = zia.get_time_window(name="Weekends")
    ```

    ```python
    import pulumi
    import pulumi_zia as zia

    off_hours = zia.get_time_window(name="Off hours")
    ```


    :param str name: The name of the time window to be exported.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zia:index/getTimeWindow:getTimeWindow', __args__, opts=opts, typ=GetTimeWindowResult).value

    return AwaitableGetTimeWindowResult(
        day_of_weeks=pulumi.get(__ret__, 'day_of_weeks'),
        end_time=pulumi.get(__ret__, 'end_time'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        start_time=pulumi.get(__ret__, 'start_time'))


@_utilities.lift_output_func(get_time_window)
def get_time_window_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTimeWindowResult]:
    """
    Use the **zia_firewall_filtering_time_window** data source to get information about a time window option available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zia as zia

    work_hours = zia.get_time_window(name="Work hours")
    ```

    ```python
    import pulumi
    import pulumi_zia as zia

    weekends = zia.get_time_window(name="Weekends")
    ```

    ```python
    import pulumi
    import pulumi_zia as zia

    off_hours = zia.get_time_window(name="Off hours")
    ```


    :param str name: The name of the time window to be exported.
    """
    ...
