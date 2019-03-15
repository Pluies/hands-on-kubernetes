# YNAP Custom

This is where any general generic Kubernetes objects specifically required by other charts that don't make sense to be created in helm preinstall hooks (eg storageclass) live. 

Also will contain our custom cleanup job to prune non-master branch helm releases after a set period of time