import functions_framework

import jsons

from src.parser import parse_log


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
    return jsons.dumps(parse_log(request.data.decode('utf-8')))

