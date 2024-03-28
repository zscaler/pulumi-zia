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
    'GetActivationStatusResult',
    'AwaitableGetActivationStatusResult',
    'get_activation_status',
    'get_activation_status_output',
]

@pulumi.output_type
class GetActivationStatusResult:
    """
    A collection of values returned by getActivationStatus.
    """
    def __init__(__self__, id=None, status=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")


class AwaitableGetActivationStatusResult(GetActivationStatusResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetActivationStatusResult(
            id=self.id,
            status=self.status)


def get_activation_status(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetActivationStatusResult:
    """
    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_zia as zia

    activation = zia.get_activation_status()
    ```
    <!--End PulumiCodeChooser -->
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zia:index/getActivationStatus:getActivationStatus', __args__, opts=opts, typ=GetActivationStatusResult).value

    return AwaitableGetActivationStatusResult(
        id=pulumi.get(__ret__, 'id'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_activation_status)
def get_activation_status_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetActivationStatusResult]:
    """
    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_zia as zia

    activation = zia.get_activation_status()
    ```
    <!--End PulumiCodeChooser -->
    """
    ...
