# Kubernetes Components

[Source](https://kubernetes.io/docs/concepts/overview/components/)

## Hierarchy

A cluster contains multiple nodes that can be individual machines or VMs that can run containerized applications where pods are hosted (a group of containers).

A Control Plane runs across multiple computers and a cluster runs with multiple nodes.

* [Nodes](https://kubernetes.io/docs/concepts/architecture/nodes/)
* [Pods](https://kubernetes.io/docs/concepts/workloads/pods/)
* [Containers](https://kubernetes.io/docs/concepts/containers/)

## Components

* ControlPlane
  * kube-apiserver
    * Exposes the Kubernetes API
  * etcd
    * Key value store for all cluster data.
  * kube-scheduler
    * Watches for newly created Pods with no assigned node, and selects a node for them to run on
  * kube-controller-manager
    * Runs [controller](https://kubernetes.io/docs/concepts/architecture/controller/) processes
  * cloud-controller-manager
    * Embeds cloud-specific control logic
* Node
  * kubelet
    * Agent that runs on each node. Makes sure that containers are running in a Pod
  * kube-proxy
    * Network proxy that runs on each Node, implementing Kubernetes [Service](https://kubernetes.io/docs/concepts/services-networking/service/)
  * Container Runtime
    * Runs containers
* Addons (cluster-level features)
  * DNS
  * WebUI (Dashboard)
  * Container Resource Monitoring
  * Cluster-level Logging
