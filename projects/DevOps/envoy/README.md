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
