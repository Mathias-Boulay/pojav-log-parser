import re

from src.checker import check_for_errors
from src.interfaces import *
from src.locale import localize


def parse_version(log: str) -> PojavLauncherVersion:
    pattern = r'Info: Launcher version: (.*?)-(.*?)-(.*?)-(\S*)'
    results = re.search(pattern, log, re.MULTILINE)
    return PojavLauncherVersion(
        results.group(1), results.group(2),
        results.group(3), results.group(4))


def parse_minecraft_version(log: str) -> MinecraftVersion:
    version_name_pattern = r'Selected Minecraft version: (.*)'
    analyzed_version = MinecraftVersion('UNKNOWN', VersionType.VANILLA)

    regex_version_name = re.search(version_name_pattern, log)
    if regex_version_name:
        analyzed_version.name = regex_version_name.group(1).lower()

    if 'forge' in analyzed_version.name:
        analyzed_version.type = VersionType.FORGE
    if 'optifine' in analyzed_version.name:
        analyzed_version.type = VersionType.OPTIFINE
    if 'fabric' in analyzed_version.name:
        analyzed_version.type = VersionType.FABRIC

    return analyzed_version


def parse_architecture(log: str) -> Architecture:
    pattern = 'Architecture: (.*)$'
    result = re.search(pattern, log, re.MULTILINE).group(1)
    result = result.strip()
    if 'arm64' in result:
        return Architecture.ARM_64
    if 'arm' in result:
        return Architecture.ARM_32

    # TODO detect x86 architecture !
    return Architecture.UNKNOWN


def parse_java_arguments(log: str) -> str:
    pattern = 'Info: Custom Java arguments:.*?\"(.*)\"'
    results = re.search(pattern, log, re.MULTILINE)
    return results.group(1)


def parse_env_variables(log: str) -> dict:
    pattern = 'Added custom env: (.*?)=(.*)$'
    results = re.findall(pattern, log, re.MULTILINE)
    # Build a dict with all the variables
    env_dictionary = {}
    for env_tuple in results:
        env_dictionary[env_tuple[0]] = env_tuple[1]

    return env_dictionary


def parse_java_runtime(java_path: str) -> JavaRuntime:
    java_path = java_path.casefold()
    if java_path.endswith('/internal'):
        return JavaRuntime(8, JavaRuntimeSource.INTERNAL, JavaRuntimeType.JRE)
    if java_path.endswith('/internal17'):
        return JavaRuntime(17, JavaRuntimeSource.INTERNAL, JavaRuntimeType.JRE)

    runtime = JavaRuntime()
    runtime.source = JavaRuntimeSource.EXTERNAL
    if 'jdk' in java_path:
        runtime.type = JavaRuntimeType.JDK
    if 'jre' in java_path:
        runtime.type = JavaRuntimeType.JRE

    pattern = r'(?:jdk|jre)(\d*)'
    results = re.search(pattern, java_path)
    try:
        runtime.version = int(results.group(1))
    except (IndexError, TypeError) as error:
        print('Error during java runtime parsing')

    return runtime


def parse_renderer(log: str) -> PojavRenderer:
    # The pattern is a log line printed by gl4es
    pattern = r'v1\.1\.(\d) built on '
    result = re.search(pattern, log)
    if result:
        gl4es_version = int(result.group(1))
        if gl4es_version == 5:
            return PojavRenderer.GL4ES_115
        if gl4es_version == 4:
            if 'GLSL 320 es supported' in log:
                return PojavRenderer.HOLY_GL4ES_114
            else:
                return PojavRenderer.GL4ES_114

    if 'VirGL: libvirgl_test_server' in log:
        return PojavRenderer.VIRGL

    if 'POJAV_RENDERER=opengles3_desktopgl_angle_vulkan' in log:
        return PojavRenderer.ANGLE

    return PojavRenderer.UNKNOWN


def parse_build_type(log: str) -> BuildType:
    if 'net.kdt.pojavlaunch.debug' in log:
        return BuildType.DEBUG
    if 'net.kdt.pojavlaunch' in log:
        return BuildType.PRODUCTION
    return BuildType.UNKNOWN


def parse_log(content_log: str, should_localize: bool=False) -> dict:
    json_result = {
        'version': parse_version(content_log),
        'minecraft_version': parse_minecraft_version(content_log),
        'renderer': parse_renderer(content_log),
        'env_variables': parse_env_variables(content_log),
        'build_type': parse_build_type(content_log),
        'architecture': parse_architecture(content_log),
        'java_arguments': parse_java_arguments(content_log)
    }
    json_result['java_runtime'] = parse_java_runtime(json_result['env_variables']['JAVA_HOME'])

    # errors
    if not should_localize:
        json_result['errors'] = check_for_errors(content_log, json_result)
    else:
        json_result['errors'] = localize(check_for_errors(content_log, json_result))

    return json_result
