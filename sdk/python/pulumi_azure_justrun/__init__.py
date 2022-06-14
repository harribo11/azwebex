# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .containerapp import *
from .provider import *
from .webapp import *
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "azure-justrun",
  "mod": "index",
  "fqn": "pulumi_azure_justrun",
  "classes": {
   "azure-justrun:index:containerapp": "Containerapp",
   "azure-justrun:index:webapp": "Webapp"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "azure-justrun",
  "token": "pulumi:providers:azure-justrun",
  "fqn": "pulumi_azure_justrun",
  "class": "Provider"
 }
]
"""
)
