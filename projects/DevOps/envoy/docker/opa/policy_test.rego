package envoy.authz

import future.keywords

test_allowed if {
	allow with input.attributes.request.http as {"path": ["/service1"], "method": "GET"}
}

test_wrong_method if {
	not allow with input.attributes.request.http as {"path": ["/service1"], "method": "POST"}
}

test_wrong_path if {
	not allow with input.attributes.request.http as {"path": ["/something"], "method": "GET"}
}
