# Prometheus Monitoring System

## Pull Metrics vs. Publishing Metrics

Prometheus scrapes metrics, while also allowing for metrics to be pushed through the Prometheus Pushgateway.

## Prometheus Server

* Data Retrieval Worker
  * Pulls metric data from services
* Time Series Database
  * Sotres metric data
* HTTP Server
  * Accepts queries over the data
  * Prometheus Web UI
  * Grafana

## Targets and Metrics

* Targets are what Prometheus monitors (a server, a service, an application, a database, etc)
* Units are what is monitored for each target (CPU status, memory usage, disk space, request duration, etc)

## Data Model

All data in Prometheus is stored as a [time series](https://en.wikipedia.org/wiki/Time_series) which are streams of timestamped values belonging to the same metric and the same set of labeled dimensions.

Every time series is uniquely identified by its __metric name__ and optional key-value pairs called __labels__.

There are four core __metric types__. These do not diferentiate in storage, only at the client.

* __Counter__: a value that can only increase;
* __Gauge__: a value that can increase or decrease;
* __Histogram__: samples observations and separates them in buckets;
* __Summary__: samples observations, while also providing a count and a sum of these.

## Configuration Example

```yaml
# How often Prometheus will scrape its targets
global:
  scrape_interval: 15s
  evaluation_interval: 15s

# Rules for aggregating metric values or creating alerts when conditions met
rule_files:
  # - "first.rules"
  # - "second.rules"

# What resources Prometheus monitors
scrape_configs:
  - job_name: prometheus
    static_configs:
      - targets: ['localhost:9090']
```

## Characteristics

| Pros | Cons |
| ---- | ---- |
| reliable | difficult to scale |
| stand-alone and self-containing | limits monitoring |
| works, even if other parts of the infrastructure break | |
| no extensive set-up needed | |
| less complex | |
