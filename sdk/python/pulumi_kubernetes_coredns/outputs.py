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
    'ReleaseStatus',
]

@pulumi.output_type
class ReleaseStatus(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "appVersion":
            suggest = "app_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ReleaseStatus. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ReleaseStatus.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ReleaseStatus.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 app_version: builtins.str,
                 chart: builtins.str,
                 name: builtins.str,
                 namespace: builtins.str,
                 revision: builtins.int,
                 status: builtins.str,
                 version: builtins.str):
        """
        :param builtins.str app_version: The version number of the application being deployed.
        :param builtins.str chart: The name of the chart.
        :param builtins.str name: Name is the name of the release.
        :param builtins.str namespace: Namespace is the kubernetes namespace of the release.
        :param builtins.int revision: Version is an int32 which represents the version of the release.
        :param builtins.str status: Status of the release.
        :param builtins.str version: A SemVer 2 conformant version string of the chart.
        """
        pulumi.set(__self__, "app_version", app_version)
        pulumi.set(__self__, "chart", chart)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "namespace", namespace)
        pulumi.set(__self__, "revision", revision)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="appVersion")
    def app_version(self) -> builtins.str:
        """
        The version number of the application being deployed.
        """
        return pulumi.get(self, "app_version")

    @property
    @pulumi.getter
    def chart(self) -> builtins.str:
        """
        The name of the chart.
        """
        return pulumi.get(self, "chart")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Name is the name of the release.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> builtins.str:
        """
        Namespace is the kubernetes namespace of the release.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def revision(self) -> builtins.int:
        """
        Version is an int32 which represents the version of the release.
        """
        return pulumi.get(self, "revision")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        Status of the release.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def version(self) -> builtins.str:
        """
        A SemVer 2 conformant version string of the chart.
        """
        return pulumi.get(self, "version")


