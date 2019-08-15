# Purpose

The e2e tests are used to test end to end functionality of the ims-kaas api.  The cms-ssh api is the primary endpoint
for ims-kaas functionality, and the cluster-manager-api(cma) api provides an end user endpoint to include callbacks
and access to other providers.  There are scripts here to test both api endpoints (cma, and ims-kaas).

# Developer Context

When making changes to ims-kaas, the developer and possibly reviewer should run the `full-test-ims-kaas.sh` script which will e2e test the ims-kaas api.

# How to run `full-test-ims-kaas.sh`

```bash
# create a k8s cluster to use for ims-kaas
minikube start
kubectl create clusterrolebinding superpowers --clusterrole=cluster-admin --user=system:serviceaccount:kube-system:default
kubectl create rolebinding superpowers --clusterrole=cluster-admin --user=system:serviceaccount:kube-system:default
kubectl apply -f crd

# build and run ims-kaas
cd $GOPATH/src/github.com/samsung-cnct/ims-kaas
go1.12.4 build -o ims-kaas cmd/ims-kaas/main.go
IMS_KAAS_MAAS_API_URL=http://192.168.2.24:5240/MAAS IMS_KAAS_MAAS_API_KEY=<your maas key> ./ims-kaas --logtostderr

# run the e2e script
cd $GOPATH/src/github.com/samsung-cnct/ims-kaas/test/e2e
source ./testEnvs.sh
./full-test-ims-kaas.sh

# Look at test output and final status
full-test-ims-kaas PASSED

```

# Testing from the cma api

The script `full-test.sh` is an end to end test of the cma api.
Normally `full-test.sh` would be used on an already installed CMS.

# How to run `full-test.sh`

Follow instructions to install [cluster-manager-api](https://github.com/samsung-cnct/cluster-manager-api)
and set helpers.ssh.enabled=true.
Note: cert-manager and nginx-ingress should be installed per instructions before installing
cluster-manager-api.

```bash
# install cma
cd $GOPATH/src/github.com/samsung-cnct/cluster-manager-api
helm install deployments/helm/cluster-manager-api --name cma --set helpers.ssh.enabled=true

# install ims-kaas
cd $GOPATH/src/github.com/samsung-cnct/ims-kaas
vi deployments/helm/ims-kaas/values.yaml  - to add maas key
helm install --name ims-kaas deployments/helm/ims-kaas/

# make sure cluster-manager-api service is accessible (NodePort, or ingress)
kubectl get svc

# run full-test.sh
cd $GOPATH/src/github.com/samsung-cnct/ims-kaas/test/e2e
vi testEnvs.sh - modify CLUSTER_API to point to cma service
source ./testEnvs.sh
./full-test.sh

# Look at test output and final status
full-test-cma PASSED

````

# CI Pipeline Context

TBD.  CI needs access to the MaaS environment in order to test end to end provisioning.

# Deprecated

This set of tests using the Cluster Manager API pipeline through CMA SSH, which will install a managed cluster on pre-provisioned hosts. In order for this specific set of tests to work, you *must* have at least two IP addresses (machines, vms, etc) running BEFORE these tests can be performed.

# How to run

1.  populate the environment with required VARs (defaults are provided)  See all of the environmentals needed in full-test.sh
2.  execute `full-test.sh`

## Sequence of the tests
(FIX UP - WIP)
1.  create a client cluster via a parent ims-kaas helper:
    `create-cluster.sh`
2.  get the kubeconfig for the client cluster from the parent cluster
    K8S API `get-kubeconfig()` in `full-test.sh`
3.  create a simple system in the client cluster (using nginx-ingress)
4.  verify the simple system functions
5.  tear down the client cluster: `delete-cluster.sh`
