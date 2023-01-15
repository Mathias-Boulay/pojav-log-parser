import functions_framework
import re
from interfaces import *


def parse_version(log: str) -> PojavLauncherVersion:
    pattern = 'Info: Launcher version: (.*?)-(.*?)-(.*?)-(.*)$'
    results = re.search(pattern, log, re.MULTILINE)
    return PojavLauncherVersion(
        results.group(1), results.group(2),
        results.group(3), results.group(4))


def parse_architecture(log: str) -> str:
    pattern = 'Architecture: (.*)$'
    results = re.search(pattern, log, re.MULTILINE)
    return results.group(1)


def parse_java_arguments(log: str) -> str:
    pattern = 'Info: Custom Java arguments: (.*)$'
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
    except IndexError | TypeError:
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


def parse_log(content_log: str) -> dict:
    print(content_log)
    return {
        'renderer': parse_renderer(content_log),
        'env_variables': parse_env_variables(content_log)
    }


@functions_framework.http
def handle_request(request):
    """
        HTTP Cloud Function that declares a variable.
        Args:
            request (flask.Request): The request object.
            <https://flask.palletsprojects.com/en/latest/api/>
        Returns:
            The response text, or any set of values that can be turned into a
            Response object using `make_response`
            <https://flask.palletsprojects.com/en/latest/api/>.
        """
    return parse_log(request.data.decode('utf-8'))

