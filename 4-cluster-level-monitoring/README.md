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
| Pod crashing | If it happens infrequently, merely an increase in the Restart Count. If too frequent, Kubernetes will back off restarting the container and show the status as `CrashLoopBackoff` | Use `restart: Always` so that your Pods |
| git diff | Show file differences that haven't been staged | |

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

