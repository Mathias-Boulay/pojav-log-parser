import functions_framework
from src.parser import *
from src.checker import *
import jsons


def parse_log(content_log: str, should_localize: bool = False) -> dict:
    json_result = {
        'version': parse_version(content_log),
        'minecraft_version': parse_minecraft_version(content_log),
        'minecraft_username': parse_minecraft_username(content_log),
        'renderer': parse_renderer(content_log),
        'env_variables': parse_env_variables(content_log),
        'build_type': parse_build_type(content_log),
        'architecture': parse_architecture(content_log),
        'java_arguments': parse_java_arguments(content_log)
    }
    json_result['java_runtime'] = parse_java_runtime(json_result['env_variables']['JAVA_HOME'])
    json_result['dumped_shaders'] = parse_shader_dump(content_log, json_result['env_variables'])

    # errors
    if not should_localize:
        json_result['errors'] = check_for_errors(content_log, json_result)
    else:
        json_result['errors'] = localize(check_for_errors(content_log, json_result))

    return json_result


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
    return jsons.dumps(parse_log(request.data.decode('utf-8'), request.args.get('should_localize', False)))

