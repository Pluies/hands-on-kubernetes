Deploy an application
=====================

Let's do it!

First, let's create our own namespace:

    kubectl create namespace $MY_NAME

And set it as the current namespace:

    kubectl config set "contexts."`kubectl config current-context`".namespace" $MY_NAME

(This will edit the kubeconfig file, we can also edit it manually)

Now let's create a Deployment:

    kubectl apply -f deployment.yaml

More info: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/

Questions
---------

Why a Deployment and not a Pod directly?

Why a Deployment and not a ReplicaSet?
