# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('zia')


class _ExportableConfig(types.ModuleType):
    @property
    def api_key(self) -> Optional[str]:
        return __config__.get('apiKey') or _utilities.get_env('ZIA_API_KEY')

    @property
    def api_token(self) -> Optional[str]:
        return __config__.get('apiToken') or _utilities.get_env('ZIA_SANDBOX_TOKEN')

    @property
    def password(self) -> Optional[str]:
        return __config__.get('password') or _utilities.get_env('ZIA_PASSWORD')

    @property
    def username(self) -> Optional[str]:
        return __config__.get('username') or _utilities.get_env('ZIA_USERNAME')

    @property
    def zia_cloud(self) -> Optional[str]:
        return __config__.get('ziaCloud') or _utilities.get_env('ZIA_CLOUD')

