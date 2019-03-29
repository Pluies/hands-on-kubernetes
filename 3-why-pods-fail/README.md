Why pods fail
=============

In this section, we'll look at common reasons why workloads fail.

Each folder contains a yaml file that you can apply in your namespace:

```
kubectl apply -f a-a-frame/a-frame.yaml
```

Two of them are special:

- a. `A-Frame` is _not_ broken! It's a basic sanity check that should apply cleanly and give you a Service accessible from within your cluster at: [http://a-frame.namespace.svc.cluster.local](http://a-frame.namespace.svc.cluster.local)
- z. `leaderboard` is a web frontend that sends http GETs to services to see if they work

All other folders will give you trouble :)

Useful commands
---------------

All the information needed to debug these cases can be found with the following commands:
- `kubectl get pods` to list the pods
- `kubectl get pods -l app=name` to list pods based on a label, as a Service does
- `kubectl describe pod $PODNAME` to get more info on a specific pod
- `kubectl logs -f $PODNAME` to tail the logs of a given pod
- `kubectl logs $PODNAME -p` to get the logs of the given pod before it crashed (`-p` is short for `--previous`)

The official documentation has some really well-written guides to debugging:
- [Debug Pods](https://kubernetes.io/docs/tasks/debug-application-cluster/debug-pod-replication-controller/)
- [Debug Services](https://kubernetes.io/docs/tasks/debug-application-cluster/debug-service/)
- [Determine the Reason for Pod Failure](https://kubernetes.io/docs/tasks/debug-application-cluster/determine-reason-pod-failure/)

Help/hints
----------

Each test is deemed to be working when it can serve HTTP traffic through its Service, so make sure that both Pods in the Deployment and the corresponding Service are working fine.

You can `kubectl exec` into the `leaderboard` pod to curl services from within the cluster.

As a last resort, the Go source code for each application is available under `src/`. It may give you hints about what's going wrong, so if you're stuck, do have a look! Note that this won't necessarily help you for all cases.
