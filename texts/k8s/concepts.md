## Kubernetes Concepts

## Context

This term applies only on the client side. It is made of access parameters to a K8s cluster containing the cluster, user and namespace. When this is set, all `kubectl` commands will run against the specified cluster

>kubectl help [command subcommand ...]

There is also [`kubectx`](https://github.com/ahmetb/kubectx) which is handy for changing context.

The configuration file can be found in `$HOME/.kube/config`

## Namespaces

Namespaces allow to group resources. Deleting a namespace will delete all child objects.

```
# set default namespace in config file
kubectl config set-context --current --namespace=kube-system
```
