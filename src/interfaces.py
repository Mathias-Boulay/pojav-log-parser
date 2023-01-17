"""
File describing interfaces or dataclasses
"""

from dataclasses import dataclass
from enum import Enum


@dataclass
class Architecture(str, Enum):
    ARM_32 = 'ARM_32'
    ARM_64 = 'ARM_64'
    X86_32 = 'X86_32'
    X86_64 = 'X86_64'
    UNKNOWN = 'UNKNOWN'


@dataclass
class JavaRuntimeSource(str, Enum):
    INTERNAL = 'INTERNAL'
    EXTERNAL = 'EXTERNAL'
    UNKNOWN = 'UNKNOWN'


@dataclass
class JavaRuntimeType(str, Enum):
    JDK = 'JDK'
    JRE = 'JRE'
    UNKNOWN = 'UNKNOWN'


@dataclass
class JavaRuntime:
    version: int = 0
    source: JavaRuntimeSource = JavaRuntimeSource.UNKNOWN
    type: JavaRuntimeType = JavaRuntimeType.UNKNOWN


@dataclass
class PojavLauncherVersion:
    major_code: str
    commit_number: int
    commit_short: str
    branch: str


@dataclass
class VersionType(str, Enum):
    VANILLA = 'VANILLA'
    OPTIFINE = 'OPTIFINE'
    FORGE = 'FORGE'
    FABRIC = 'FABRIC'
    UNKNOWN = 'UNKNOWN'


@dataclass
class PojavRenderer(str, Enum):
    GL4ES_114 = 'GL4ES_114'
    GL4ES_115 = 'GL4ES_115'
    HOLY_GL4ES_114 = 'HOLY_GL4ES_114'
    ANGLE = 'ANGLE'
    VIRGL = 'VIRGL'
    ZINK = 'ZINK'
    UNKNOWN = 'UNKNOWN'


@dataclass
class BuildType(str, Enum):
    DEBUG = 'DEBUG'
    PRODUCTION = 'PRODUCTION'
    UNKNOWN = 'UNKNOWN'


@dataclass
class MinecraftVersion:
    name: str = 'UNKNOWN'
    type: VersionType = VersionType.UNKNOWN


