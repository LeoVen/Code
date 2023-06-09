package envoy.authz

import input.attributes.request.http as http_request

default allow = false

allow {
    is_request_valid
}

is_request_valid {
    http_request.method == "GET"
    startswith(http_request.path, "/service")
}
