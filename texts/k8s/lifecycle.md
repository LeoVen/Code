# Kubernetes POD Lifecycle

Sources

[Video](https://www.youtube.com/watch?v=wlYESb124xM)

[Docs](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/)

A pod is the atomic unit of scheduling and may contain one or more containers

Lifecycle:

* 1- Pending
  * Hasn't been scheduled to a machine yet
  * If stuck here, resources are not available
  * If stuck here, a bad image may have been provided (e.g. 404 on the creating step)
* 2- Creating
  * The host is pulling the containers from the registry
* 3- Running
  * The program is running
* 4- Crashloop Backoff
  * If the container is crashing too many times

## Lifecycle Integration

* Common routes for lifecycle integration:
  * `/health`
  * `/ready`
  * `/postStart`
  * `/preStop`
* Init container: provide some initialization before the other containers inside a POD starts
  * Schema migration
  * File downloads
