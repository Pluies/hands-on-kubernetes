What's a kubeconfig file?
=========================

A kubeconfig file tells kubectl and other tools how to connect to one or more clusters.

Defaults to `~/.kube/config`

It contains:
- A list of servers with:
 - The URL of the APIserver, so you know who to talk to
 - The Certificate Authority of the APIserver, to ensure authenticity and prevent MITM
- Authentication info so that the API server knows who you are
- Contexts to group api endpoint + authentication + namespace into one

In practice:

```yaml
apiVersion: v1
kind: Config
preferences: {}

clusters:
- name: dev-cluster
  cluster:
    certificate-authority-data: LS0tLS1CRUd......=

users:
- name: jane-c-developer
  user:
    token: eyJhbGciOiJSUzI1NiIsImtpZCI....

contexts:
- name: dev-context
  context:
    cluster: dev-cluster
    user: jane-c-developer
    namespace: dev-ns
```

https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/
