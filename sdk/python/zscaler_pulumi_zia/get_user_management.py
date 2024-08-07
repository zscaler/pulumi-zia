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

__all__ = [
    'GetUserManagementResult',
    'AwaitableGetUserManagementResult',
    'get_user_management',
    'get_user_management_output',
]

@pulumi.output_type
class GetUserManagementResult:
    """
    A collection of values returned by getUserManagement.
    """
    def __init__(__self__, admin_user=None, auth_methods=None, comments=None, departments=None, email=None, groups=None, id=None, is_auditor=None, name=None, temp_auth_email=None, type=None):
        if admin_user and not isinstance(admin_user, bool):
            raise TypeError("Expected argument 'admin_user' to be a bool")
        pulumi.set(__self__, "admin_user", admin_user)
        if auth_methods and not isinstance(auth_methods, list):
            raise TypeError("Expected argument 'auth_methods' to be a list")
        pulumi.set(__self__, "auth_methods", auth_methods)
        if comments and not isinstance(comments, str):
            raise TypeError("Expected argument 'comments' to be a str")
        pulumi.set(__self__, "comments", comments)
        if departments and not isinstance(departments, list):
            raise TypeError("Expected argument 'departments' to be a list")
        pulumi.set(__self__, "departments", departments)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)
        if is_auditor and not isinstance(is_auditor, str):
            raise TypeError("Expected argument 'is_auditor' to be a str")
        pulumi.set(__self__, "is_auditor", is_auditor)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if temp_auth_email and not isinstance(temp_auth_email, str):
            raise TypeError("Expected argument 'temp_auth_email' to be a str")
        pulumi.set(__self__, "temp_auth_email", temp_auth_email)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="adminUser")
    def admin_user(self) -> bool:
        """
        (String) True if this user is an Admin user. readOnly: `true` default: `false`
        """
        return pulumi.get(self, "admin_user")

    @property
    @pulumi.getter(name="authMethods")
    def auth_methods(self) -> Optional[Sequence[str]]:
        """
        (String) Type of authentication method to be enabled. Supported values are: ``BASIC`` and ``DIGEST``
        """
        return pulumi.get(self, "auth_methods")

    @property
    @pulumi.getter
    def comments(self) -> str:
        """
        (String) Additional information about the group
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def departments(self) -> Sequence['outputs.GetUserManagementDepartmentResult']:
        """
        (String) Department a user belongs to
        """
        return pulumi.get(self, "departments")

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        (Required) User email consists of a user name and domain name. It does not have to be a valid email address, but it must be unique and its domain must belong to the organization
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.GetUserManagementGroupResult']:
        """
        (String) List of Groups a user belongs to. Groups are used in policies.
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        (Number) Unique identfier for the group
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isAuditor")
    def is_auditor(self) -> str:
        return pulumi.get(self, "is_auditor")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        (String) Group name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="tempAuthEmail")
    def temp_auth_email(self) -> str:
        """
        (String) Temporary Authentication Email. If you enabled one-time tokens or links, enter the email address to which the Zscaler service sends the tokens or links. If this is empty, the service will send the email to the User email.
        """
        return pulumi.get(self, "temp_auth_email")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        (String) User type. Provided only if this user is not an end user. The supported types are:
        * `SUPERADMIN`
        * `ADMIN`
        * `AUDITOR`
        * `GUEST`
        * `REPORT_USER`
        * `UNAUTH_TRAFFIC_DEFAULT`
        """
        return pulumi.get(self, "type")


class AwaitableGetUserManagementResult(GetUserManagementResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserManagementResult(
            admin_user=self.admin_user,
            auth_methods=self.auth_methods,
            comments=self.comments,
            departments=self.departments,
            email=self.email,
            groups=self.groups,
            id=self.id,
            is_auditor=self.is_auditor,
            name=self.name,
            temp_auth_email=self.temp_auth_email,
            type=self.type)


def get_user_management(auth_methods: Optional[Sequence[str]] = None,
                        id: Optional[int] = None,
                        name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserManagementResult:
    """
    Use the **zia_user_management** data source to get information about a user account that may have been created in the Zscaler Internet Access portal or via API. This data source can then be associated with a ZIA cloud firewall filtering rule, and URL filtering rules.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zia as zia

    adam_ashcroft = zia.get_user_management(name="Adam Ashcroft")
    ```


    :param Sequence[str] auth_methods: (String) Type of authentication method to be enabled. Supported values are: ``BASIC`` and ``DIGEST``
    :param int id: The ID of the time window resource.
    :param str name: User name. This appears when choosing users for policies.
    """
    __args__ = dict()
    __args__['authMethods'] = auth_methods
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zia:index/getUserManagement:getUserManagement', __args__, opts=opts, typ=GetUserManagementResult).value

    return AwaitableGetUserManagementResult(
        admin_user=pulumi.get(__ret__, 'admin_user'),
        auth_methods=pulumi.get(__ret__, 'auth_methods'),
        comments=pulumi.get(__ret__, 'comments'),
        departments=pulumi.get(__ret__, 'departments'),
        email=pulumi.get(__ret__, 'email'),
        groups=pulumi.get(__ret__, 'groups'),
        id=pulumi.get(__ret__, 'id'),
        is_auditor=pulumi.get(__ret__, 'is_auditor'),
        name=pulumi.get(__ret__, 'name'),
        temp_auth_email=pulumi.get(__ret__, 'temp_auth_email'),
        type=pulumi.get(__ret__, 'type'))


@_utilities.lift_output_func(get_user_management)
def get_user_management_output(auth_methods: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                               id: Optional[pulumi.Input[Optional[int]]] = None,
                               name: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserManagementResult]:
    """
    Use the **zia_user_management** data source to get information about a user account that may have been created in the Zscaler Internet Access portal or via API. This data source can then be associated with a ZIA cloud firewall filtering rule, and URL filtering rules.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zia as zia

    adam_ashcroft = zia.get_user_management(name="Adam Ashcroft")
    ```


    :param Sequence[str] auth_methods: (String) Type of authentication method to be enabled. Supported values are: ``BASIC`` and ``DIGEST``
    :param int id: The ID of the time window resource.
    :param str name: User name. This appears when choosing users for policies.
    """
    ...
