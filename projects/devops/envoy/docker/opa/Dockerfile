FROM openpolicyagent/opa:0.56.0-envoy-static

COPY . .

CMD ["run", "--server", "./policy.rego", "--config-file=./opa-config.yaml", "--log-level", "debug"]
