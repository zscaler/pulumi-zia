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
    'GetAdminRolesResult',
    'AwaitableGetAdminRolesResult',
    'get_admin_roles',
    'get_admin_roles_output',
]

@pulumi.output_type
class GetAdminRolesResult:
    """
    A collection of values returned by getAdminRoles.
    """
    def __init__(__self__, admin_acct_access=None, analysis_access=None, dashboard_access=None, id=None, is_auditor=None, is_non_editable=None, logs_limit=None, name=None, permissions=None, policy_access=None, rank=None, report_access=None, role_type=None, username_access=None):
        if admin_acct_access and not isinstance(admin_acct_access, str):
            raise TypeError("Expected argument 'admin_acct_access' to be a str")
        pulumi.set(__self__, "admin_acct_access", admin_acct_access)
        if analysis_access and not isinstance(analysis_access, str):
            raise TypeError("Expected argument 'analysis_access' to be a str")
        pulumi.set(__self__, "analysis_access", analysis_access)
        if dashboard_access and not isinstance(dashboard_access, str):
            raise TypeError("Expected argument 'dashboard_access' to be a str")
        pulumi.set(__self__, "dashboard_access", dashboard_access)
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)
        if is_auditor and not isinstance(is_auditor, bool):
            raise TypeError("Expected argument 'is_auditor' to be a bool")
        pulumi.set(__self__, "is_auditor", is_auditor)
        if is_non_editable and not isinstance(is_non_editable, bool):
            raise TypeError("Expected argument 'is_non_editable' to be a bool")
        pulumi.set(__self__, "is_non_editable", is_non_editable)
        if logs_limit and not isinstance(logs_limit, str):
            raise TypeError("Expected argument 'logs_limit' to be a str")
        pulumi.set(__self__, "logs_limit", logs_limit)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if permissions and not isinstance(permissions, list):
            raise TypeError("Expected argument 'permissions' to be a list")
        pulumi.set(__self__, "permissions", permissions)
        if policy_access and not isinstance(policy_access, str):
            raise TypeError("Expected argument 'policy_access' to be a str")
        pulumi.set(__self__, "policy_access", policy_access)
        if rank and not isinstance(rank, int):
            raise TypeError("Expected argument 'rank' to be a int")
        pulumi.set(__self__, "rank", rank)
        if report_access and not isinstance(report_access, str):
            raise TypeError("Expected argument 'report_access' to be a str")
        pulumi.set(__self__, "report_access", report_access)
        if role_type and not isinstance(role_type, str):
            raise TypeError("Expected argument 'role_type' to be a str")
        pulumi.set(__self__, "role_type", role_type)
        if username_access and not isinstance(username_access, str):
            raise TypeError("Expected argument 'username_access' to be a str")
        pulumi.set(__self__, "username_access", username_access)

    @property
    @pulumi.getter(name="adminAcctAccess")
    def admin_acct_access(self) -> str:
        """
        (String)
        """
        return pulumi.get(self, "admin_acct_access")

    @property
    @pulumi.getter(name="analysisAccess")
    def analysis_access(self) -> str:
        """
        (String)
        """
        return pulumi.get(self, "analysis_access")

    @property
    @pulumi.getter(name="dashboardAccess")
    def dashboard_access(self) -> str:
        """
        (String) Dashboard access permission. Supported values are: `NONE`, `READ_ONLY`
        """
        return pulumi.get(self, "dashboard_access")

    @property
    @pulumi.getter
    def id(self) -> int:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isAuditor")
    def is_auditor(self) -> bool:
        """
        (Boolean) Indicates whether this is an auditor role.
        """
        return pulumi.get(self, "is_auditor")

    @property
    @pulumi.getter(name="isNonEditable")
    def is_non_editable(self) -> bool:
        """
        (Boolean) Indicates whether or not this admin user is editable/deletable.
        """
        return pulumi.get(self, "is_non_editable")

    @property
    @pulumi.getter(name="logsLimit")
    def logs_limit(self) -> str:
        """
        (String) Log range limit. Returned values are: `UNRESTRICTED`, `MONTH_1`, `MONTH_2`, `MONTH_3`, `MONTH_4`, `MONTH_5`, `MONTH_6`
        """
        return pulumi.get(self, "logs_limit")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def permissions(self) -> Sequence[str]:
        """
        (List of String) List of functional areas to which this role has access. This attribute is subject to change.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="policyAccess")
    def policy_access(self) -> str:
        """
        (String) Policy access permission. Returned values are: `NONE`, `READ_ONLY`,`READ_WRITE`
        """
        return pulumi.get(self, "policy_access")

    @property
    @pulumi.getter
    def rank(self) -> int:
        """
        (Number) Admin rank of this admin role. This is applicable only when admin rank is enabled in the advanced settings. Default value is 7 (the lowest rank). The assigned admin rank determines the roles or admin users this user can manage, and which rule orders this admin can access.
        """
        return pulumi.get(self, "rank")

    @property
    @pulumi.getter(name="reportAccess")
    def report_access(self) -> str:
        """
        (String) Report access permission. Returned values are: `NONE`, `READ_ONLY`,`READ_WRITE`
        """
        return pulumi.get(self, "report_access")

    @property
    @pulumi.getter(name="roleType")
    def role_type(self) -> str:
        """
        (String) The admin role type. ()This attribute is subject to change.) Supported values are:  `ORG_ADMIN`, `EXEC_INSIGHT`, `EXEC_INSIGHT_AND_ORG_ADMIN`, `SDWAN`
        """
        return pulumi.get(self, "role_type")

    @property
    @pulumi.getter(name="usernameAccess")
    def username_access(self) -> str:
        """
        (String) Username access permission. When set to NONE, the username will be obfuscated. Supported values are: `NONE|READ_ONLY`
        """
        return pulumi.get(self, "username_access")


class AwaitableGetAdminRolesResult(GetAdminRolesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAdminRolesResult(
            admin_acct_access=self.admin_acct_access,
            analysis_access=self.analysis_access,
            dashboard_access=self.dashboard_access,
            id=self.id,
            is_auditor=self.is_auditor,
            is_non_editable=self.is_non_editable,
            logs_limit=self.logs_limit,
            name=self.name,
            permissions=self.permissions,
            policy_access=self.policy_access,
            rank=self.rank,
            report_access=self.report_access,
            role_type=self.role_type,
            username_access=self.username_access)


def get_admin_roles(name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAdminRolesResult:
    """
    Use the **zia_admin_roles** data source to get information about an admin role created in the Zscaler Internet Access cloud or via the API. This data source can then be associated with a ZIA administrator account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zia as zia

    example = zia.get_admin_roles(name="Super Admin")
    ```


    :param str name: The name of the Admin role to be exported.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zia:index/getAdminRoles:getAdminRoles', __args__, opts=opts, typ=GetAdminRolesResult).value

    return AwaitableGetAdminRolesResult(
        admin_acct_access=pulumi.get(__ret__, 'admin_acct_access'),
        analysis_access=pulumi.get(__ret__, 'analysis_access'),
        dashboard_access=pulumi.get(__ret__, 'dashboard_access'),
        id=pulumi.get(__ret__, 'id'),
        is_auditor=pulumi.get(__ret__, 'is_auditor'),
        is_non_editable=pulumi.get(__ret__, 'is_non_editable'),
        logs_limit=pulumi.get(__ret__, 'logs_limit'),
        name=pulumi.get(__ret__, 'name'),
        permissions=pulumi.get(__ret__, 'permissions'),
        policy_access=pulumi.get(__ret__, 'policy_access'),
        rank=pulumi.get(__ret__, 'rank'),
        report_access=pulumi.get(__ret__, 'report_access'),
        role_type=pulumi.get(__ret__, 'role_type'),
        username_access=pulumi.get(__ret__, 'username_access'))


@_utilities.lift_output_func(get_admin_roles)
def get_admin_roles_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAdminRolesResult]:
    """
    Use the **zia_admin_roles** data source to get information about an admin role created in the Zscaler Internet Access cloud or via the API. This data source can then be associated with a ZIA administrator account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zia as zia

    example = zia.get_admin_roles(name="Super Admin")
    ```


    :param str name: The name of the Admin role to be exported.
    """
    ...
