from src.interfaces import *
from src.locale import localize


def check_env_variables(env_variables: dict) -> list[str]:
    errors = []
    if int(env_variables.get('LIBGL_MIPMAP', 3)) != 3:
        errors.append('mipmap.enabled')

    if not (0 <= int(env_variables.get('LIBGL_ES', 0)) <= 2):
        errors.append('libgles.wrong.value')

    if int(env_variables.get('LIBGL_GL', 0)) != 0:
        errors.append('libgl.custom.value')

    return errors


def check_java_arguments(java_args: str) -> bool:
    """
    Reimplementation of the java argument parser
    Warns if something was left unparsed or something
    """
    separators = ["-XX:-", "-XX:+", "-XX:", "--", "-D", "-X"]
    for separator in separators:
        while True:
            start_index = java_args.find(separator)
            if start_index == -1:
                break
            # Get the end of the current argument by checking the nearest separator
            end_index = -1
            for end_separator in separators:
                temp_end = java_args.find(end_separator, start_index + len(separator))
                if temp_end == -1:
                    continue
                if end_index == -1:
                    end_index = temp_end
                    continue
                end_index = min(end_index, temp_end)
            # Fallback
            if end_index == -1:
                end_index = len(java_args)

            # In this re-implementation, just to delete it
            extracted_argument = java_args[start_index:end_index]
            java_args = java_args.replace(extracted_argument, '')

            # Check whether two args aren't bundled together by mistake
            if extracted_argument.find('=') != extracted_argument.rfind('='):
                return False

    return not java_args


def check_branch(version: PojavLauncherVersion) -> bool:
    if version.branch != "v3_openjdk":
        return False

    return True

# TODO check whether the commit is the latest ?


def check_for_errors(log: str, parsed_dict: dict) -> list[str]:
    errors = []

    # User input checks
    errors += check_env_variables(parsed_dict['env_variables'])
    if not check_java_arguments(parsed_dict['java_arguments']):
        errors.append('java.args.unparsed')

    # Branch checks
    if not check_branch(parsed_dict['version']):
        errors.append('branch.not.mainline')

    # Runtime checks
    if parsed_dict['java_runtime'].source != JavaRuntimeSource.INTERNAL:
        errors.append('java.runtime.wrong.source')
    if parsed_dict['java_runtime'].type != JavaRuntimeType.JRE:
        errors.append('java.runtime.wrong.type')

    # Virgl + 32 bits combo check
    if parsed_dict['renderer'] == PojavRenderer.VIRGL and \
            (parsed_dict['architecture'] == Architecture.ARM_32 or parsed_dict['architecture'] == Architecture.X86_32):
        errors.append('renderer.unsupported.arch')

    # Corruption errors
    # Chrome
    if parsed_dict['minecraft_version'].type == VersionType.FORGE and 'java.lang.NoClassDefFoundError' in log:
        errors.append('install.forge.corrupted')

    # Shader errors
    if 'could not reload shaders' in log:
        errors.append('load.shader.fail')

    # TODO more verification

    return errors
