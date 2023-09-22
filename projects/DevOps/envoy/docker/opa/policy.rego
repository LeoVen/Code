package envoy.authz

import future.keywords
import input.attributes.request.http as http_request

default allow := false

allow if {
	print(input)
	is_request_valid
}

is_request_valid if {
	http_request.method == "GET"
	startswith(http_request.path[0], "/service")
}
