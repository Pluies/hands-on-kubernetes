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

Help/hints
----------

Each test is deemed to be working when it can serve HTTP traffic through its Service, so make sure that both Pods in the Deployment and the corresponding Service are working fine.

You can `kubectl exec` into the `leaderboard` pod to curl services from within the cluster.

As a last resort, the Go source code for each application is available under `src/`. It may give you hints about what's going wrong, so if you're stuck, do have a look! Note that this won't necessarily help you for all cases.
