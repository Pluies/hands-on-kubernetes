Expose an application
=====================

It's as simple as a Service:

    kubectl apply -f service.yaml

AWS integration then creates an ELB.

In GKE, it would create Google's flavour of a load balancer.

Let's check it works:

    kubectl describe service nginx

More info: https://kubernetes.io/docs/concepts/services-networking/service/
