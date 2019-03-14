What's kubectl?
===============

- Main entrypoint / CLI into a cluster
- Can be used to get data:

    kubectl get pods
    kubectl get pod foo
    kubectl describe pod foo

- Get logs:

    kubectl logs foo

- Exec into a container:

    kubectl exec foo

- Edit stuff:

    kubectl apply -f things.yaml
    kubectl edit pod foo

- Delete stuff:

    kubectl delete pod foo

- And more!

https://kubernetes.io/docs/tasks/tools/install-kubectl/
https://kubernetes.io/docs/reference/kubectl/overview/
