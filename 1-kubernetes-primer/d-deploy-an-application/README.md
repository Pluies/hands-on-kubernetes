Deploy an application
=====================

Let's do it!

First, let's create our own namespace:

    kubectl create namespace $MY_NAME

And set it as the current namespace:

    kubectl config set "contexts."`kubectl config current-context`".namespace" $MY_NAME

(This will edit the kubeconfig file to set the namespace – it can also be edited manually)

Now let's create a Deployment:

    kubectl apply -f deployment.yaml

More info: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/

See what happened
-----------------

```
kubectl get deployments
```

```
kubectl get pods
```

```
kubectl port-forward deployment/nginx 8080:80
```
-> And open http://127.0.0.1:8080

Questions
---------

Why a Deployment and not a Pod directly?

What happens when a Pod gets deleted?
