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
from . import outputs

__all__ = [
    'GetAdminUsersResult',
    'AwaitableGetAdminUsersResult',
    'get_admin_users',
    'get_admin_users_output',
]

@pulumi.output_type
class GetAdminUsersResult:
    """
    A collection of values returned by getAdminUsers.
    """
    def __init__(__self__, admin_scopes=None, comments=None, disabled=None, email=None, exec_mobile_app_tokens=None, id=None, is_auditor=None, is_exec_mobile_app_enabled=None, is_non_editable=None, is_password_expired=None, is_password_login_allowed=None, is_product_update_comm_enabled=None, is_security_report_comm_enabled=None, is_service_update_comm_enabled=None, login_name=None, pwd_last_modified_time=None, roles=None, username=None):
        if admin_scopes and not isinstance(admin_scopes, list):
            raise TypeError("Expected argument 'admin_scopes' to be a list")
        pulumi.set(__self__, "admin_scopes", admin_scopes)
        if comments and not isinstance(comments, str):
            raise TypeError("Expected argument 'comments' to be a str")
        pulumi.set(__self__, "comments", comments)
        if disabled and not isinstance(disabled, bool):
            raise TypeError("Expected argument 'disabled' to be a bool")
        pulumi.set(__self__, "disabled", disabled)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if exec_mobile_app_tokens and not isinstance(exec_mobile_app_tokens, list):
            raise TypeError("Expected argument 'exec_mobile_app_tokens' to be a list")
        pulumi.set(__self__, "exec_mobile_app_tokens", exec_mobile_app_tokens)
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)
        if is_auditor and not isinstance(is_auditor, bool):
            raise TypeError("Expected argument 'is_auditor' to be a bool")
        pulumi.set(__self__, "is_auditor", is_auditor)
        if is_exec_mobile_app_enabled and not isinstance(is_exec_mobile_app_enabled, bool):
            raise TypeError("Expected argument 'is_exec_mobile_app_enabled' to be a bool")
        pulumi.set(__self__, "is_exec_mobile_app_enabled", is_exec_mobile_app_enabled)
        if is_non_editable and not isinstance(is_non_editable, bool):
            raise TypeError("Expected argument 'is_non_editable' to be a bool")
        pulumi.set(__self__, "is_non_editable", is_non_editable)
        if is_password_expired and not isinstance(is_password_expired, bool):
            raise TypeError("Expected argument 'is_password_expired' to be a bool")
        pulumi.set(__self__, "is_password_expired", is_password_expired)
        if is_password_login_allowed and not isinstance(is_password_login_allowed, bool):
            raise TypeError("Expected argument 'is_password_login_allowed' to be a bool")
        pulumi.set(__self__, "is_password_login_allowed", is_password_login_allowed)
        if is_product_update_comm_enabled and not isinstance(is_product_update_comm_enabled, bool):
            raise TypeError("Expected argument 'is_product_update_comm_enabled' to be a bool")
        pulumi.set(__self__, "is_product_update_comm_enabled", is_product_update_comm_enabled)
        if is_security_report_comm_enabled and not isinstance(is_security_report_comm_enabled, bool):
            raise TypeError("Expected argument 'is_security_report_comm_enabled' to be a bool")
        pulumi.set(__self__, "is_security_report_comm_enabled", is_security_report_comm_enabled)
        if is_service_update_comm_enabled and not isinstance(is_service_update_comm_enabled, bool):
            raise TypeError("Expected argument 'is_service_update_comm_enabled' to be a bool")
        pulumi.set(__self__, "is_service_update_comm_enabled", is_service_update_comm_enabled)
        if login_name and not isinstance(login_name, str):
            raise TypeError("Expected argument 'login_name' to be a str")
        pulumi.set(__self__, "login_name", login_name)
        if pwd_last_modified_time and not isinstance(pwd_last_modified_time, int):
            raise TypeError("Expected argument 'pwd_last_modified_time' to be a int")
        pulumi.set(__self__, "pwd_last_modified_time", pwd_last_modified_time)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="adminScopes")
    def admin_scopes(self) -> Sequence['outputs.GetAdminUsersAdminScopeResult']:
        """
        (Set of Object) The admin's scope. Only applicable for the LOCATION_GROUP admin scope type, in which case this attribute gives the list of ID/name pairs of locations within the location group.
        """
        return pulumi.get(self, "admin_scopes")

    @property
    @pulumi.getter
    def comments(self) -> builtins.str:
        """
        (String) Additional information about the admin or auditor.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def disabled(self) -> builtins.bool:
        """
        (Boolean) Indicates whether or not the admin account is disabled.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def email(self) -> builtins.str:
        """
        (String) Admin or auditor's email address.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="execMobileAppTokens")
    def exec_mobile_app_tokens(self) -> Sequence['outputs.GetAdminUsersExecMobileAppTokenResult']:
        """
        (List of Object)
        """
        return pulumi.get(self, "exec_mobile_app_tokens")

    @property
    @pulumi.getter
    def id(self) -> builtins.int:
        """
        (Number) Identifier that uniquely identifies an entity
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isAuditor")
    def is_auditor(self) -> builtins.bool:
        """
        (Boolean) Indicates whether the user is an auditor. This attribute is subject to change.
        """
        return pulumi.get(self, "is_auditor")

    @property
    @pulumi.getter(name="isExecMobileAppEnabled")
    def is_exec_mobile_app_enabled(self) -> builtins.bool:
        """
        (Boolean) Indicates whether or not Executive Insights App access is enabled for the admin.
        """
        return pulumi.get(self, "is_exec_mobile_app_enabled")

    @property
    @pulumi.getter(name="isNonEditable")
    def is_non_editable(self) -> builtins.bool:
        """
        (Boolean) Indicates whether or not the admin can be edited or deleted.
        """
        return pulumi.get(self, "is_non_editable")

    @property
    @pulumi.getter(name="isPasswordExpired")
    def is_password_expired(self) -> builtins.bool:
        """
        (Boolean) Indicates whether or not an admin's password has expired.
        """
        return pulumi.get(self, "is_password_expired")

    @property
    @pulumi.getter(name="isPasswordLoginAllowed")
    def is_password_login_allowed(self) -> builtins.bool:
        """
        (Boolean) The default is true when SAML Authentication is disabled. When SAML Authentication is enabled, this can be set to false in order to force the admin to login via SSO only.
        """
        return pulumi.get(self, "is_password_login_allowed")

    @property
    @pulumi.getter(name="isProductUpdateCommEnabled")
    def is_product_update_comm_enabled(self) -> builtins.bool:
        """
        (Boolean) Communication setting for Product Update.
        """
        return pulumi.get(self, "is_product_update_comm_enabled")

    @property
    @pulumi.getter(name="isSecurityReportCommEnabled")
    def is_security_report_comm_enabled(self) -> builtins.bool:
        """
        (Boolean) Communication for Security Report is enabled.
        """
        return pulumi.get(self, "is_security_report_comm_enabled")

    @property
    @pulumi.getter(name="isServiceUpdateCommEnabled")
    def is_service_update_comm_enabled(self) -> builtins.bool:
        """
        (Boolean) Communication setting for Service Update.
        """
        return pulumi.get(self, "is_service_update_comm_enabled")

    @property
    @pulumi.getter(name="loginName")
    def login_name(self) -> builtins.str:
        return pulumi.get(self, "login_name")

    @property
    @pulumi.getter(name="pwdLastModifiedTime")
    def pwd_last_modified_time(self) -> builtins.int:
        return pulumi.get(self, "pwd_last_modified_time")

    @property
    @pulumi.getter
    def roles(self) -> Sequence['outputs.GetAdminUsersRoleResult']:
        """
        (Set of Object) Role of the admin. This is not required for an auditor.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter
    def username(self) -> builtins.str:
        return pulumi.get(self, "username")


class AwaitableGetAdminUsersResult(GetAdminUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAdminUsersResult(
            admin_scopes=self.admin_scopes,
            comments=self.comments,
            disabled=self.disabled,
            email=self.email,
            exec_mobile_app_tokens=self.exec_mobile_app_tokens,
            id=self.id,
            is_auditor=self.is_auditor,
            is_exec_mobile_app_enabled=self.is_exec_mobile_app_enabled,
            is_non_editable=self.is_non_editable,
            is_password_expired=self.is_password_expired,
            is_password_login_allowed=self.is_password_login_allowed,
            is_product_update_comm_enabled=self.is_product_update_comm_enabled,
            is_security_report_comm_enabled=self.is_security_report_comm_enabled,
            is_service_update_comm_enabled=self.is_service_update_comm_enabled,
            login_name=self.login_name,
            pwd_last_modified_time=self.pwd_last_modified_time,
            roles=self.roles,
            username=self.username)


def get_admin_users(id: Optional[builtins.int] = None,
                    login_name: Optional[builtins.str] = None,
                    username: Optional[builtins.str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAdminUsersResult:
    """
    * [Official documentation](https://help.zscaler.com/zia/about-administrators)
    * [API documentation](https://help.zscaler.com/zia/admin-role-management#/adminUsers-get)

    Use the **zia_admin_users** data source to get information about an admin user account created in the Zscaler Internet Access cloud or via the API. This data source can then be associated with a ZIA administrator role.

    ## Example Usage


    :param builtins.int id: The ID of the admin user to be exported.
    :param builtins.str login_name: The email address of the admin user to be exported.
    :param builtins.str username: The username of the admin user to be exported.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['loginName'] = login_name
    __args__['username'] = username
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zia:index/getAdminUsers:getAdminUsers', __args__, opts=opts, typ=GetAdminUsersResult).value

    return AwaitableGetAdminUsersResult(
        admin_scopes=pulumi.get(__ret__, 'admin_scopes'),
        comments=pulumi.get(__ret__, 'comments'),
        disabled=pulumi.get(__ret__, 'disabled'),
        email=pulumi.get(__ret__, 'email'),
        exec_mobile_app_tokens=pulumi.get(__ret__, 'exec_mobile_app_tokens'),
        id=pulumi.get(__ret__, 'id'),
        is_auditor=pulumi.get(__ret__, 'is_auditor'),
        is_exec_mobile_app_enabled=pulumi.get(__ret__, 'is_exec_mobile_app_enabled'),
        is_non_editable=pulumi.get(__ret__, 'is_non_editable'),
        is_password_expired=pulumi.get(__ret__, 'is_password_expired'),
        is_password_login_allowed=pulumi.get(__ret__, 'is_password_login_allowed'),
        is_product_update_comm_enabled=pulumi.get(__ret__, 'is_product_update_comm_enabled'),
        is_security_report_comm_enabled=pulumi.get(__ret__, 'is_security_report_comm_enabled'),
        is_service_update_comm_enabled=pulumi.get(__ret__, 'is_service_update_comm_enabled'),
        login_name=pulumi.get(__ret__, 'login_name'),
        pwd_last_modified_time=pulumi.get(__ret__, 'pwd_last_modified_time'),
        roles=pulumi.get(__ret__, 'roles'),
        username=pulumi.get(__ret__, 'username'))
def get_admin_users_output(id: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                           login_name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           username: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAdminUsersResult]:
    """
    * [Official documentation](https://help.zscaler.com/zia/about-administrators)
    * [API documentation](https://help.zscaler.com/zia/admin-role-management#/adminUsers-get)

    Use the **zia_admin_users** data source to get information about an admin user account created in the Zscaler Internet Access cloud or via the API. This data source can then be associated with a ZIA administrator role.

    ## Example Usage


    :param builtins.int id: The ID of the admin user to be exported.
    :param builtins.str login_name: The email address of the admin user to be exported.
    :param builtins.str username: The username of the admin user to be exported.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['loginName'] = login_name
    __args__['username'] = username
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('zia:index/getAdminUsers:getAdminUsers', __args__, opts=opts, typ=GetAdminUsersResult)
    return __ret__.apply(lambda __response__: GetAdminUsersResult(
        admin_scopes=pulumi.get(__response__, 'admin_scopes'),
        comments=pulumi.get(__response__, 'comments'),
        disabled=pulumi.get(__response__, 'disabled'),
        email=pulumi.get(__response__, 'email'),
        exec_mobile_app_tokens=pulumi.get(__response__, 'exec_mobile_app_tokens'),
        id=pulumi.get(__response__, 'id'),
        is_auditor=pulumi.get(__response__, 'is_auditor'),
        is_exec_mobile_app_enabled=pulumi.get(__response__, 'is_exec_mobile_app_enabled'),
        is_non_editable=pulumi.get(__response__, 'is_non_editable'),
        is_password_expired=pulumi.get(__response__, 'is_password_expired'),
        is_password_login_allowed=pulumi.get(__response__, 'is_password_login_allowed'),
        is_product_update_comm_enabled=pulumi.get(__response__, 'is_product_update_comm_enabled'),
        is_security_report_comm_enabled=pulumi.get(__response__, 'is_security_report_comm_enabled'),
        is_service_update_comm_enabled=pulumi.get(__response__, 'is_service_update_comm_enabled'),
        login_name=pulumi.get(__response__, 'login_name'),
        pwd_last_modified_time=pulumi.get(__response__, 'pwd_last_modified_time'),
        roles=pulumi.get(__response__, 'roles'),
        username=pulumi.get(__response__, 'username')))
