Cluster-level monitoring
=================

What can go wrong?

In-cluster
==========

Pods
----

As we saw last week, Pods can fail to start, can crash, can become unresponsive, can run out of resources...

Thankfully, Kubernetes is pretty good at handling Pods failure cases.

| Failure case | Symptoms | Remediation |
| --- | --- | --- |
| Pod crashing | If it happens infrequently, merely an increase in the Restart Count.
If too frequent, Kubernetes will back off restarting the container and show the status as `CrashLoopBackoff` | Use `restart: Always` on your container definition so that Kubernetes will restart them if they crash (it is the default) |
| Pod becomes unresponsive | If the pod has a `readinessProbe` defined, this probe will start failing and traffic from a Service will stop being directed to the pod, preventing users from hitting an unresponsive pod. The Pod status will show as not ready. If all pods in a service become unresponsive, the whole service will start failing. | Define `readinessProbe` to ensure traffic always goes to healthy pods, and `livenessProbe` to ensure pods get restarted if they're unresponsive. |
| New release of a service is broken | Kubernetes will try and bring up new Pods by doing a rolling update, and, noticing these pods are not coming up healthy, will stop the deployment. Some healthy pods will be left from the previous ReplicaSet, so existing traffic should not be impacted. | Monitor deployments to ensure they go through properly. Monitor ReplicaSet to ensure the desired number of replicas is equal to the actual number of pods. |

Nodes
-----

Kubernetes nodes are standard EC2 instances, so we have the classic trifecta of CPU, RAM, disk.

Once again, Kubernetes should handle these failure modes by itself, but we can also help by configuring [resource requests and limits](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/) appropriately.

| Failure case | Symptoms | Remediation |
| --- | --- | --- |
| Node Out of Memory | The Node will log an alert saying it went out of memory and starts evicting Pods. If memory reaches its maximum before kubelet notices, the Linux OOM killer kicks in and starts killing processes, which can be fairly disruptive. | Use `memory` requests and limits so that workloads are not allowed to consume all memory on a node. |
| Node maxing out CPU | Workloads will become slow as the CPU is being hogged by "something else". This can lead to pods crashing if they cannot serve LivenessProbes, and in extreme cases, the entire Node becoming unresponsive and Pods showing up as `LostNode`. | Use `cpu` requests and limits so that workloads are not allowed to consume all cpu on a node |
| Node running out of disk space | Kubernetes will evict Pods consuming too much ephemeral storage to reclaim the disk space. | Use `ephemeralStorage` requests and limits so that workloads are not allowed to consume all ephemeral storage on a node |

DNS: in-cluster DNS

h2. Out-of-cluster

AWS integration:
ELBs
Autoscaling
ExternalDNS

