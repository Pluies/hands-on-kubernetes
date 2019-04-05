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
| Pod crashing | If it happens infrequently, merely an increase in the Restart Count. If too frequent, Kubernetes will back off restarting the container and show the status as `CrashLoopBackoff` | Use `restart: Always` on your container definition so that Kubernetes will restart them if they crash (it is the default) |
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

Control plane
-------------

The term _control plane_ refers to Kubernetes master nodes and the datastore backing Kubernetes, etcd.

We run Kubernetes on the AWS-managed EKS, so all of this is transparent to us and managed by Kubernetes (which is great, because configuring masters, maintaining HA on etcd, and backups, are not very fun at all).

DNS
---

In order to provide internal DNS facilities that resolves such addresses as [my-service.namespace.svc.cluster.local](my-service.namespace.svc.cluster.local), Kubernetes has DNS resolvers in-cluster.

These are deployed as a standard Deployment of [core-dns](https://coredns.io/) instances that resolve Kubernetes addresses themselves and forward other requests to upstream DNS servers.

Being a Deployment, CoreDNS is mnonitored like any other Kubernetes workloads.

Out-of-cluster
==============

We run Kubernetes on EKS, AWS integrations to provide functionality to Kubernetes.

Autoscaling Groups
------------------

Node groups are defined as AWS Autoscaling groups. Autoscaling groups will automatically detect bad hardware or failed AWS Status Checks and boot that instance out of the ASG. A new instance will then come back and rejoin the cluster.

ELBs
----

Traffic arrives from the "external world" into the cluster through standard AWS ELBs that are created by a `Service` of type `LoadBalancer`.

If the Service is deleted in Kubernetes, then the ELB will be deleted as well.

ExternalDNS
-----------

In order not to have to update upstream to point to a new ELB when this happens, we use [ExternalDNS](https://github.com/kubernetes-incubator/external-dns).

ExternalDNS lets us add labels to a `Service` and will configure DNS records in Route53 to point to the ELB created by Kubernetes, ensuring that even if the Service gets deleted and recreated, exteranl traffic will still be able to be routed to the cluster.

Monitoring & Alerting pipeline
==============================

All our alerts and monitoring conditions are defined in Prometheus.

Prometheus scrapes targets from within Kubernetes, and stores the metrics. It also raises alerts based on conditions we define.

Alerts are then picked up by AlertManager, whose role is to send notifications to various channels (Slack, PagerDuty, DeadsManSnitch...). It also allows us to silence alerts.

We use Grafana as a front-end to visualise data stored in Kubernetes.
