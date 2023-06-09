FROM openpolicyagent/opa:0.53.1-envoy-static

COPY . .

# CMD run --server ./policy.rego --config-file=./opa-config.yaml --log-level debug
CMD ["run", "--server", "./policy.rego", "--config-file=./opa-config.yaml", "--log-level", "debug"]
