# Envoy example

* Starts and envoy front proxy
* Runs two services
  * `service-example-1`
  * `service-example-2`

To access these services, use the following paths:

* `http://127.0.0.1:8080/service1`
* `http://127.0.0.1:8080/service2`

The admin console is also accessible at:

* `http://127.0.0.1:9901/`

To run everything, simply run:

* `docker compose build`
* `docker compose up`

## Testing

Test the OPA policy

* `opa test ./docker/opa/`

Test the requests to Envoy

* `curl -v 127.0.0.1:8080/service1`
* `curl -v 127.0.0.1:8080/service2`

## Relevant Docs

* [Envoy External Authorization](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/ext_authz_filter)
* [Policies for Envoy](https://www.openpolicyagent.org/docs/latest/envoy-primer/)
* [Envoy with OPA](https://www.openpolicyagent.org/docs/latest/envoy-introduction/)
