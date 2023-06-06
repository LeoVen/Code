FROM envoyproxy/envoy:v1.26-latest

RUN apt update && \
    apt install -y curl && \
    rm -rf /var/lib/apt/lists

COPY . /envoy
WORKDIR /envoy

RUN chown -R envoy:envoy .
CMD envoy -c ./envoy.yaml
