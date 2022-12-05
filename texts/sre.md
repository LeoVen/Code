# Site Reliability Engineering

[Source](https://sre.google/sre-book/service-level-objectives/)

## Terminology

* __SLI__: Service Level Indicators
* __SLO__: Service Level Objectives
* __SLA__: Service Level Agreements

### Indicators

Service level indicators can be availability, latency, CPU usage, memory usage, exceptions, successfull requests, etc. It is important to choose the correct ones, without aggregating too few or too many.

### Objectives

A target value that is expected from an SLI. For example, wanting to have request latency under 100 milliseconds or that a processing of a 1GB file should take no longer than 5 minutes to return results to the front-end.

### Agreements

These are agreements about the consequences of meeting or missing SLOs. Consequences can be financial or others defined by both parties. SLAs are closely tied to business and product decisions. E.g. Google Search doesn't have an SLA as it doesn't sign a contract with its users of when SLOs aren't met.
