Expose an application
=====================

It's as simple as a Service:

    kubectl apply -f service.yaml

AWS integration then creates an ELB. In GKE, it would create a Google Load Balancer.

Let's check it works

More info: https://kubernetes.io/docs/concepts/services-networking/service/
