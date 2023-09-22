FROM envoyproxy/envoy:v1.27-latest

RUN apt update && \
    apt install -y curl dnsutils && \
    rm -rf /var/lib/apt/lists

COPY . /envoy
WORKDIR /envoy

RUN chown -R envoy:envoy .
CMD envoy -c ./envoy.yaml --log-level info
