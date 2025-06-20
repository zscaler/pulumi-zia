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
    'GetDLPWebRulesResult',
    'AwaitableGetDLPWebRulesResult',
    'get_dlp_web_rules',
    'get_dlp_web_rules_output',
]

@pulumi.output_type
class GetDLPWebRulesResult:
    """
    A collection of values returned by getDLPWebRules.
    """
    def __init__(__self__, access_control=None, action=None, cloud_applications=None, departments=None, description=None, dlp_download_scan_enabled=None, dlp_engines=None, excluded_departments=None, excluded_groups=None, excluded_users=None, external_auditor_email=None, file_types=None, groups=None, id=None, included_domain_profiles=None, labels=None, last_modified_bies=None, last_modified_time=None, location_groups=None, locations=None, match_only=None, min_size=None, name=None, order=None, parent_rule=None, protocols=None, rank=None, severity=None, source_ip_groups=None, state=None, sub_rules=None, time_windows=None, url_categories=None, users=None, without_content_inspection=None, workload_groups=None, zcc_notifications_enabled=None, zscaler_incident_receiver=None):
        if access_control and not isinstance(access_control, str):
            raise TypeError("Expected argument 'access_control' to be a str")
        pulumi.set(__self__, "access_control", access_control)
        if action and not isinstance(action, str):
            raise TypeError("Expected argument 'action' to be a str")
        pulumi.set(__self__, "action", action)
        if cloud_applications and not isinstance(cloud_applications, list):
            raise TypeError("Expected argument 'cloud_applications' to be a list")
        pulumi.set(__self__, "cloud_applications", cloud_applications)
        if departments and not isinstance(departments, list):
            raise TypeError("Expected argument 'departments' to be a list")
        pulumi.set(__self__, "departments", departments)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dlp_download_scan_enabled and not isinstance(dlp_download_scan_enabled, bool):
            raise TypeError("Expected argument 'dlp_download_scan_enabled' to be a bool")
        pulumi.set(__self__, "dlp_download_scan_enabled", dlp_download_scan_enabled)
        if dlp_engines and not isinstance(dlp_engines, list):
            raise TypeError("Expected argument 'dlp_engines' to be a list")
        pulumi.set(__self__, "dlp_engines", dlp_engines)
        if excluded_departments and not isinstance(excluded_departments, list):
            raise TypeError("Expected argument 'excluded_departments' to be a list")
        pulumi.set(__self__, "excluded_departments", excluded_departments)
        if excluded_groups and not isinstance(excluded_groups, list):
            raise TypeError("Expected argument 'excluded_groups' to be a list")
        pulumi.set(__self__, "excluded_groups", excluded_groups)
        if excluded_users and not isinstance(excluded_users, list):
            raise TypeError("Expected argument 'excluded_users' to be a list")
        pulumi.set(__self__, "excluded_users", excluded_users)
        if external_auditor_email and not isinstance(external_auditor_email, str):
            raise TypeError("Expected argument 'external_auditor_email' to be a str")
        pulumi.set(__self__, "external_auditor_email", external_auditor_email)
        if file_types and not isinstance(file_types, list):
            raise TypeError("Expected argument 'file_types' to be a list")
        pulumi.set(__self__, "file_types", file_types)
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)
        if included_domain_profiles and not isinstance(included_domain_profiles, list):
            raise TypeError("Expected argument 'included_domain_profiles' to be a list")
        pulumi.set(__self__, "included_domain_profiles", included_domain_profiles)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if last_modified_bies and not isinstance(last_modified_bies, list):
            raise TypeError("Expected argument 'last_modified_bies' to be a list")
        pulumi.set(__self__, "last_modified_bies", last_modified_bies)
        if last_modified_time and not isinstance(last_modified_time, int):
            raise TypeError("Expected argument 'last_modified_time' to be a int")
        pulumi.set(__self__, "last_modified_time", last_modified_time)
        if location_groups and not isinstance(location_groups, list):
            raise TypeError("Expected argument 'location_groups' to be a list")
        pulumi.set(__self__, "location_groups", location_groups)
        if locations and not isinstance(locations, list):
            raise TypeError("Expected argument 'locations' to be a list")
        pulumi.set(__self__, "locations", locations)
        if match_only and not isinstance(match_only, bool):
            raise TypeError("Expected argument 'match_only' to be a bool")
        pulumi.set(__self__, "match_only", match_only)
        if min_size and not isinstance(min_size, int):
            raise TypeError("Expected argument 'min_size' to be a int")
        pulumi.set(__self__, "min_size", min_size)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if order and not isinstance(order, int):
            raise TypeError("Expected argument 'order' to be a int")
        pulumi.set(__self__, "order", order)
        if parent_rule and not isinstance(parent_rule, int):
            raise TypeError("Expected argument 'parent_rule' to be a int")
        pulumi.set(__self__, "parent_rule", parent_rule)
        if protocols and not isinstance(protocols, list):
            raise TypeError("Expected argument 'protocols' to be a list")
        pulumi.set(__self__, "protocols", protocols)
        if rank and not isinstance(rank, int):
            raise TypeError("Expected argument 'rank' to be a int")
        pulumi.set(__self__, "rank", rank)
        if severity and not isinstance(severity, str):
            raise TypeError("Expected argument 'severity' to be a str")
        pulumi.set(__self__, "severity", severity)
        if source_ip_groups and not isinstance(source_ip_groups, list):
            raise TypeError("Expected argument 'source_ip_groups' to be a list")
        pulumi.set(__self__, "source_ip_groups", source_ip_groups)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if sub_rules and not isinstance(sub_rules, list):
            raise TypeError("Expected argument 'sub_rules' to be a list")
        pulumi.set(__self__, "sub_rules", sub_rules)
        if time_windows and not isinstance(time_windows, list):
            raise TypeError("Expected argument 'time_windows' to be a list")
        pulumi.set(__self__, "time_windows", time_windows)
        if url_categories and not isinstance(url_categories, list):
            raise TypeError("Expected argument 'url_categories' to be a list")
        pulumi.set(__self__, "url_categories", url_categories)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)
        if without_content_inspection and not isinstance(without_content_inspection, bool):
            raise TypeError("Expected argument 'without_content_inspection' to be a bool")
        pulumi.set(__self__, "without_content_inspection", without_content_inspection)
        if workload_groups and not isinstance(workload_groups, list):
            raise TypeError("Expected argument 'workload_groups' to be a list")
        pulumi.set(__self__, "workload_groups", workload_groups)
        if zcc_notifications_enabled and not isinstance(zcc_notifications_enabled, bool):
            raise TypeError("Expected argument 'zcc_notifications_enabled' to be a bool")
        pulumi.set(__self__, "zcc_notifications_enabled", zcc_notifications_enabled)
        if zscaler_incident_receiver and not isinstance(zscaler_incident_receiver, bool):
            raise TypeError("Expected argument 'zscaler_incident_receiver' to be a bool")
        pulumi.set(__self__, "zscaler_incident_receiver", zscaler_incident_receiver)

    @property
    @pulumi.getter(name="accessControl")
    def access_control(self) -> builtins.str:
        return pulumi.get(self, "access_control")

    @property
    @pulumi.getter
    def action(self) -> builtins.str:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="cloudApplications")
    def cloud_applications(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "cloud_applications")

    @property
    @pulumi.getter
    def departments(self) -> Sequence['outputs.GetDLPWebRulesDepartmentResult']:
        return pulumi.get(self, "departments")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dlpDownloadScanEnabled")
    def dlp_download_scan_enabled(self) -> builtins.bool:
        return pulumi.get(self, "dlp_download_scan_enabled")

    @property
    @pulumi.getter(name="dlpEngines")
    def dlp_engines(self) -> Sequence['outputs.GetDLPWebRulesDlpEngineResult']:
        return pulumi.get(self, "dlp_engines")

    @property
    @pulumi.getter(name="excludedDepartments")
    def excluded_departments(self) -> Sequence['outputs.GetDLPWebRulesExcludedDepartmentResult']:
        return pulumi.get(self, "excluded_departments")

    @property
    @pulumi.getter(name="excludedGroups")
    def excluded_groups(self) -> Sequence['outputs.GetDLPWebRulesExcludedGroupResult']:
        return pulumi.get(self, "excluded_groups")

    @property
    @pulumi.getter(name="excludedUsers")
    def excluded_users(self) -> Sequence['outputs.GetDLPWebRulesExcludedUserResult']:
        return pulumi.get(self, "excluded_users")

    @property
    @pulumi.getter(name="externalAuditorEmail")
    def external_auditor_email(self) -> builtins.str:
        return pulumi.get(self, "external_auditor_email")

    @property
    @pulumi.getter(name="fileTypes")
    def file_types(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "file_types")

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.GetDLPWebRulesGroupResult']:
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.int]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includedDomainProfiles")
    def included_domain_profiles(self) -> Sequence['outputs.GetDLPWebRulesIncludedDomainProfileResult']:
        return pulumi.get(self, "included_domain_profiles")

    @property
    @pulumi.getter
    def labels(self) -> Sequence['outputs.GetDLPWebRulesLabelResult']:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lastModifiedBies")
    def last_modified_bies(self) -> Sequence['outputs.GetDLPWebRulesLastModifiedByResult']:
        return pulumi.get(self, "last_modified_bies")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> builtins.int:
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter(name="locationGroups")
    def location_groups(self) -> Sequence['outputs.GetDLPWebRulesLocationGroupResult']:
        return pulumi.get(self, "location_groups")

    @property
    @pulumi.getter
    def locations(self) -> Sequence['outputs.GetDLPWebRulesLocationResult']:
        return pulumi.get(self, "locations")

    @property
    @pulumi.getter(name="matchOnly")
    def match_only(self) -> builtins.bool:
        return pulumi.get(self, "match_only")

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> builtins.int:
        return pulumi.get(self, "min_size")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def order(self) -> builtins.int:
        return pulumi.get(self, "order")

    @property
    @pulumi.getter(name="parentRule")
    def parent_rule(self) -> builtins.int:
        return pulumi.get(self, "parent_rule")

    @property
    @pulumi.getter
    def protocols(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "protocols")

    @property
    @pulumi.getter
    def rank(self) -> builtins.int:
        return pulumi.get(self, "rank")

    @property
    @pulumi.getter
    def severity(self) -> builtins.str:
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter(name="sourceIpGroups")
    def source_ip_groups(self) -> Sequence['outputs.GetDLPWebRulesSourceIpGroupResult']:
        return pulumi.get(self, "source_ip_groups")

    @property
    @pulumi.getter
    def state(self) -> builtins.str:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="subRules")
    def sub_rules(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "sub_rules")

    @property
    @pulumi.getter(name="timeWindows")
    def time_windows(self) -> Sequence['outputs.GetDLPWebRulesTimeWindowResult']:
        return pulumi.get(self, "time_windows")

    @property
    @pulumi.getter(name="urlCategories")
    def url_categories(self) -> Sequence['outputs.GetDLPWebRulesUrlCategoryResult']:
        return pulumi.get(self, "url_categories")

    @property
    @pulumi.getter
    def users(self) -> Sequence['outputs.GetDLPWebRulesUserResult']:
        return pulumi.get(self, "users")

    @property
    @pulumi.getter(name="withoutContentInspection")
    def without_content_inspection(self) -> builtins.bool:
        return pulumi.get(self, "without_content_inspection")

    @property
    @pulumi.getter(name="workloadGroups")
    def workload_groups(self) -> Sequence['outputs.GetDLPWebRulesWorkloadGroupResult']:
        return pulumi.get(self, "workload_groups")

    @property
    @pulumi.getter(name="zccNotificationsEnabled")
    def zcc_notifications_enabled(self) -> builtins.bool:
        return pulumi.get(self, "zcc_notifications_enabled")

    @property
    @pulumi.getter(name="zscalerIncidentReceiver")
    def zscaler_incident_receiver(self) -> builtins.bool:
        return pulumi.get(self, "zscaler_incident_receiver")


class AwaitableGetDLPWebRulesResult(GetDLPWebRulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDLPWebRulesResult(
            access_control=self.access_control,
            action=self.action,
            cloud_applications=self.cloud_applications,
            departments=self.departments,
            description=self.description,
            dlp_download_scan_enabled=self.dlp_download_scan_enabled,
            dlp_engines=self.dlp_engines,
            excluded_departments=self.excluded_departments,
            excluded_groups=self.excluded_groups,
            excluded_users=self.excluded_users,
            external_auditor_email=self.external_auditor_email,
            file_types=self.file_types,
            groups=self.groups,
            id=self.id,
            included_domain_profiles=self.included_domain_profiles,
            labels=self.labels,
            last_modified_bies=self.last_modified_bies,
            last_modified_time=self.last_modified_time,
            location_groups=self.location_groups,
            locations=self.locations,
            match_only=self.match_only,
            min_size=self.min_size,
            name=self.name,
            order=self.order,
            parent_rule=self.parent_rule,
            protocols=self.protocols,
            rank=self.rank,
            severity=self.severity,
            source_ip_groups=self.source_ip_groups,
            state=self.state,
            sub_rules=self.sub_rules,
            time_windows=self.time_windows,
            url_categories=self.url_categories,
            users=self.users,
            without_content_inspection=self.without_content_inspection,
            workload_groups=self.workload_groups,
            zcc_notifications_enabled=self.zcc_notifications_enabled,
            zscaler_incident_receiver=self.zscaler_incident_receiver)


def get_dlp_web_rules(id: Optional[builtins.int] = None,
                      name: Optional[builtins.str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDLPWebRulesResult:
    """
    * [Official documentation](https://help.zscaler.com/zia/configuring-dlp-policy-rules-content-inspection#Rules)
    * [API documentation](https://help.zscaler.com/zia/data-loss-prevention#/webDlpRules-get)

    Use the **zia_dlp_web_rules** data source to get information about a ZIA DLP Web Rules in the Zscaler Internet Access cloud or via the API.

    ## Example Usage
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zia:index/getDLPWebRules:getDLPWebRules', __args__, opts=opts, typ=GetDLPWebRulesResult).value

    return AwaitableGetDLPWebRulesResult(
        access_control=pulumi.get(__ret__, 'access_control'),
        action=pulumi.get(__ret__, 'action'),
        cloud_applications=pulumi.get(__ret__, 'cloud_applications'),
        departments=pulumi.get(__ret__, 'departments'),
        description=pulumi.get(__ret__, 'description'),
        dlp_download_scan_enabled=pulumi.get(__ret__, 'dlp_download_scan_enabled'),
        dlp_engines=pulumi.get(__ret__, 'dlp_engines'),
        excluded_departments=pulumi.get(__ret__, 'excluded_departments'),
        excluded_groups=pulumi.get(__ret__, 'excluded_groups'),
        excluded_users=pulumi.get(__ret__, 'excluded_users'),
        external_auditor_email=pulumi.get(__ret__, 'external_auditor_email'),
        file_types=pulumi.get(__ret__, 'file_types'),
        groups=pulumi.get(__ret__, 'groups'),
        id=pulumi.get(__ret__, 'id'),
        included_domain_profiles=pulumi.get(__ret__, 'included_domain_profiles'),
        labels=pulumi.get(__ret__, 'labels'),
        last_modified_bies=pulumi.get(__ret__, 'last_modified_bies'),
        last_modified_time=pulumi.get(__ret__, 'last_modified_time'),
        location_groups=pulumi.get(__ret__, 'location_groups'),
        locations=pulumi.get(__ret__, 'locations'),
        match_only=pulumi.get(__ret__, 'match_only'),
        min_size=pulumi.get(__ret__, 'min_size'),
        name=pulumi.get(__ret__, 'name'),
        order=pulumi.get(__ret__, 'order'),
        parent_rule=pulumi.get(__ret__, 'parent_rule'),
        protocols=pulumi.get(__ret__, 'protocols'),
        rank=pulumi.get(__ret__, 'rank'),
        severity=pulumi.get(__ret__, 'severity'),
        source_ip_groups=pulumi.get(__ret__, 'source_ip_groups'),
        state=pulumi.get(__ret__, 'state'),
        sub_rules=pulumi.get(__ret__, 'sub_rules'),
        time_windows=pulumi.get(__ret__, 'time_windows'),
        url_categories=pulumi.get(__ret__, 'url_categories'),
        users=pulumi.get(__ret__, 'users'),
        without_content_inspection=pulumi.get(__ret__, 'without_content_inspection'),
        workload_groups=pulumi.get(__ret__, 'workload_groups'),
        zcc_notifications_enabled=pulumi.get(__ret__, 'zcc_notifications_enabled'),
        zscaler_incident_receiver=pulumi.get(__ret__, 'zscaler_incident_receiver'))
def get_dlp_web_rules_output(id: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                             name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDLPWebRulesResult]:
    """
    * [Official documentation](https://help.zscaler.com/zia/configuring-dlp-policy-rules-content-inspection#Rules)
    * [API documentation](https://help.zscaler.com/zia/data-loss-prevention#/webDlpRules-get)

    Use the **zia_dlp_web_rules** data source to get information about a ZIA DLP Web Rules in the Zscaler Internet Access cloud or via the API.

    ## Example Usage
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('zia:index/getDLPWebRules:getDLPWebRules', __args__, opts=opts, typ=GetDLPWebRulesResult)
    return __ret__.apply(lambda __response__: GetDLPWebRulesResult(
        access_control=pulumi.get(__response__, 'access_control'),
        action=pulumi.get(__response__, 'action'),
        cloud_applications=pulumi.get(__response__, 'cloud_applications'),
        departments=pulumi.get(__response__, 'departments'),
        description=pulumi.get(__response__, 'description'),
        dlp_download_scan_enabled=pulumi.get(__response__, 'dlp_download_scan_enabled'),
        dlp_engines=pulumi.get(__response__, 'dlp_engines'),
        excluded_departments=pulumi.get(__response__, 'excluded_departments'),
        excluded_groups=pulumi.get(__response__, 'excluded_groups'),
        excluded_users=pulumi.get(__response__, 'excluded_users'),
        external_auditor_email=pulumi.get(__response__, 'external_auditor_email'),
        file_types=pulumi.get(__response__, 'file_types'),
        groups=pulumi.get(__response__, 'groups'),
        id=pulumi.get(__response__, 'id'),
        included_domain_profiles=pulumi.get(__response__, 'included_domain_profiles'),
        labels=pulumi.get(__response__, 'labels'),
        last_modified_bies=pulumi.get(__response__, 'last_modified_bies'),
        last_modified_time=pulumi.get(__response__, 'last_modified_time'),
        location_groups=pulumi.get(__response__, 'location_groups'),
        locations=pulumi.get(__response__, 'locations'),
        match_only=pulumi.get(__response__, 'match_only'),
        min_size=pulumi.get(__response__, 'min_size'),
        name=pulumi.get(__response__, 'name'),
        order=pulumi.get(__response__, 'order'),
        parent_rule=pulumi.get(__response__, 'parent_rule'),
        protocols=pulumi.get(__response__, 'protocols'),
        rank=pulumi.get(__response__, 'rank'),
        severity=pulumi.get(__response__, 'severity'),
        source_ip_groups=pulumi.get(__response__, 'source_ip_groups'),
        state=pulumi.get(__response__, 'state'),
        sub_rules=pulumi.get(__response__, 'sub_rules'),
        time_windows=pulumi.get(__response__, 'time_windows'),
        url_categories=pulumi.get(__response__, 'url_categories'),
        users=pulumi.get(__response__, 'users'),
        without_content_inspection=pulumi.get(__response__, 'without_content_inspection'),
        workload_groups=pulumi.get(__response__, 'workload_groups'),
        zcc_notifications_enabled=pulumi.get(__response__, 'zcc_notifications_enabled'),
        zscaler_incident_receiver=pulumi.get(__response__, 'zscaler_incident_receiver')))
