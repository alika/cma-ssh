# ims-kaas
[![CircleCI](https://circleci.com/gh/samsung-cnct/ims-kaas/tree/master.svg?style=shield&circle-token=43855a16154d0ded4dda8b3fa08f3dd4b2298a69)](https://circleci.com/gh/samsung-cnct/ims-kaas/tree/master)

# What is this?

`ims-kaas` is a k8s operator which manages the lifecycle of Kubernetes "managed"
clusters (i.e. `CnctCluster` resources), machinesets (`CnctMachineSet`), and machines (`CnctMachine`). Currently,
this tool instantiates a managed cluster using the MaaS API.

In order for this tool to work, one must manage some pre-requisites first:

1. Set up a bare-metal pool of [MaaS](maas_website) managed servers. This
project was developed using a lab of NUCs running on maas 2.5. How to set
up maas is beyond the scope of this documentation. However, the documentation
on the [MaaS](maas_website) website is comprehensive and simple. NOTE! The
global [user_data](maas_user_data) configuration outlined in the next step
MUST be applied before any Kubernetes hosts are instantiated.

1. Before your pool of servers is set up in maas, the maas controller has to have
some global [user_data](maas_user_data) configuration copied to it for this
project to properly function; namely swap will become re-enabled on machines
after a reboot unless you apply this user_data change. Follow the above
instructions and then move on.

1. Be sure and create a maas API key and know how to obtain for later in this
document. You can [generate an API Key][generate-an-api-key] using the MAAS GUI.

## ims-kaas usage

### Building ims-kaas

There are a few steps you should run to get your development environment set up.

```bash
make go1.12.7
```

Now whenever you want to build you should run.

```bash
make
```

If you need to regenerate files you can do it with the following commands.

```bash
# run the protobuf compiler and code generator
make protoc
# create crd and rbac
make manifests
```

If you want to test a clean build (no deps installed) or you love long build times.

```bash
make clean-test
```

**Note:** If you have issues building, there is an outstanding issue in which
you need to build this project in the appropriate `$GOPATH` location. Specifically
`$GOPATH/src/github.com/samsung-cnct/ims-kaas`.

### Running ims-kaas

ims-kaas can be run locally if you have access to the Maas Lab (i.e. login to the Maas lab VPN if you are remote).
We also have a helm chart to install ims-kaas into a kubernetes cluster

#### To run ims-kaas locally for development
```bash
export IMS_KAAS_MAAS_API_URL=http://192.168.2.24:5240/MAAS/
export IMS_KAAS_MAAS_API_VERSION=2
export IMS_KAAS_MAAS_API_KEY=<your maas api key>

minikube start
make
./ims-kaas
```
#### To install ims-kaas in a kubernetes cluster

The Kubernetes cluster on which the `ims-kaas` is installed must
have network access to a MAAS server. Within the CNCT lab this
means you must be in the Seattle office or logged onto the VPN.

To test `ims-kaas` you can use `kind` and `helm`. You'll need to
obtain a copy of your maas apikey. For example:

```bash
kind create cluster
export KUBECONFIG="$(kind get kubeconfig-path --name="1")"

# If you're using helm < 3
# install helm tiller plugin
helm tiller install
helm tiller start-ci
export HELM_HOST=localhost:44134
helm install --name ims-kaas deployments/helm/ims-kaas/ --set maas.apiKey=<maas api key>

# --OR--

# if you're using helm >= 3
helm install ims-kaas deployments/helm/ims-kaas/ --set maas.apiKey=<maas api key>

kubectl get pods --watch
```

## Creating kubernetes clusters with ims-kaas using kubectl

Either kubectl or the Swagger UI REST interface can be used to create Kubernetes clusters with ims-kaas.  This section will focus on using kubectl.

A cluster definition consists of three kinds of Kubernetes Custom Resource Definitions (CRDs):
- [cnctcluster CRD](https://github.com/samsung-cnct/ims-kaas/blob/master/crd/cluster_v1alpha1_cnctcluster.yaml),
- [cnctmachine CRD](https://github.com/samsung-cnct/ims-kaas/blob/master/crd/cluster_v1alpha1_cnctmachine.yaml),
- [cnctmachineset CRD](https://github.com/samsung-cnct/ims-kaas/blob/master/crd/cluster_v1alpha1_cnctmachineset.yaml)

To create a single cluster apply these manifests to define:
- one [cnctcluster resource](https://github.com/samsung-cnct/ims-kaas/blob/master/samples/cluster/cluster_v1alpha1_cluster.yaml),
- one [cnctmachine resource](https://github.com/samsung-cnct/ims-kaas/blob/master/samples/cluster/cluster_v1alpha1_machine.yaml) to define the master node,
- one or more [cnctmachineset resources](https://github.com/samsung-cnct/ims-kaas/blob/master/samples/cluster/cluster_v1alpha1_machineset.yaml) to define each worker node pool.

### Namespace per cluster

The resources for a single cluster definition must be in the same namespace.
You cannot define two clusters in the same namespace, each cluster requires its own namespace.
The preferred naming option is to use unique cluster names and define a namespace that matches the cluster name.

### Node Pool
The cnctmachineset resource defines each worker node pool.  The `replicas` field specifies the number of worker nodes in the pool.
The ims-kaas machineset controller is similar to a kubernetes replicaSet controller.  It creates the specified number of cnctmachines, and labels them with the specified labels.  The specified label for a cnctmachineset node pool is what is used to label the cnctmachines that are in the same pool.  This label should be unique per node pool within a cluster - we recommend using the cnctmachineset name for the label.  The replica count of cnctmachines in the node pool can be scaled up and down.

#### Deleting Nodes in a Node Pool
Scaling down the replica count will delete cnctmachines labeled with the node pool specified label in a random order.

#### Deleting a Specific Node in a Node Pool
If a specific cnctmachine in the worker node pool needs to be deleted, you must add an annotation to the cnctmachine with this format: `cluster.k8s.io/delete-machine: true`, and then scale down the replica count in the corresponding (ownerReference) cnctmachineset.

```bash
kubectl annotate cnctmachine <cnctmachine name> cluster.k8s.io/delete-machine=true -n <cluster name>
```
### Example using samples for a cluster named cluster1

Create a namespace for the cluster definition resources:

```bash
kubectl create namespace cluster1
```

Copy the resource samples to your own cluster dir and modify them:

```bash
mkdir ~/cluster1
cp samples/cluster/cluster_v1alpha1_cluster.yaml ~/cluster1/cluster.yaml
cp samples/cluster/cluster_v1alpha1_machine.yaml ~/cluster1/machines.yaml
cp samples/cluster/cluster_v1alpha1_machineset.yaml ~/cluster1/machineset.yaml

vi ~/cluster1/cluster.yaml
#Modify the name, namespace, and optionally kubernetes version

vi ~/cluster1/machine.yaml
# Modify the name, namespace and instanceType (to match MaaS tags desired)

vi ~/cluster1/machineset.yaml
# Add a nodepool name, the replica count, and the label that is unique for this nodepool
# tip: use the nodepool name for the label
```
Using kubectl, apply a cluster manifest, and one or more machine manifests to
create a kubernetes cluster:

```bash
kubectl apply -f ~/cluster1/cluster.yaml
kubectl apply -f ~/cluster1/machine.yaml
kubectl apply -f ~/cluster1/machineset.yaml
```

## How instanceType is mapped to MaaS machine tags

[MaaS tags](https://docs.maas.io/2.5/en/nodes-tags) can be used to:
- select hardware reserved for use by ims-kaas,
- select hardware for masters or workers, and
- select hardware for specific workloads (e.g. those requiring GPUs, etc.)

### Define MaaS tags on MaaS machines before using ims-kaas

User defined MaaS tags would be assigned to MaaS machines using the MaaS cli or
MaaS UI before running ims-kaas. The machine spec [instanceType](https://github.com/samsung-cnct/ims-kaas/blob/master/samples/cluster/cluster_v1alpha1_machine.yaml#L15)
field is used to map a single instanceType string to a MaaS tag.  If no MaaS
tags have been defined, the instanceType field can be passed in as an empty
string so that any MaaS machine will be chosen.

## Retrieving the kubeconfig for the cluster

A secret named `cluster-private-key` is defined in the namespace of the cluster.

To retrieve the kubeconfig:
```bash
# If you're using Linux `base64` then use `-d` not `-D`
kubectl get secret cluster-private-key -ojson -n <namespace> | \
  jq -r '.data["kubernetes.kubeconfig"]' | \
  base64 -D > kubeconfig-<clustername>
```
To use the kubeconfig:
```bash
kubectl get nodes --kubeconfig kubeconfig-<clustername>
````

## Deleting the cluster or individual machines

To delete the cluster:
```bash
kubectl delete cnctcluster <cluster name> -n <namespace>
```

To delete a single machine in the cluster:
```bash
kubectl delete cnctmachine <machine name> -n <namespace>
```

# Deprecated

The instructions below are deprecated as we move towards a cloud-init approach
to configuration instead of ssh.

## Overview

The ims-kaas repo provides a helper API for [cluster-manager-api](https://github.com/samsung-cnct/cluster-manager-api)
by utilizing ssh to interact with virtual machines for kubernetes cluster
create, upgrade, add node, and delete.

### Getting started

See [Protocol Documentation](https://github.com/samsung-cnct/ims-kaas/blob/master/docs/api-generated/api.md)
- [open api in swagger ui](http://petstore.swagger.io/?url=https://raw.githubusercontent.com/samsung-cnct/ims-kaas/master/assets/generated/swagger/api.swagger.json)
- [open api in swagger editor](https://editor.swagger.io/?url=https://raw.githubusercontent.com/samsung-cnct/ims-kaas/master/assets/generated/swagger/api.swagger.json)


### Requirements
- Kubernetes 1.10+

### Deployment
The default way to deploy ims-kaas is by the provided helm chart located in the
`deployment/helm/ims-kaas` directory.

#### install via [helm](https://helm.sh/docs/using_helm/#quickstart)
1. Locate the private IP of a k8s node that ims-kaas is going to be deployed on
and will be used as the `install.bootstrapIp`.
1. Locate the nginx proxy used by the airgap environment to be used as the
`install.airgapProxyIp`.
1. Install helm chart passing in the above values:
    ```bash
    helm install deployments/helm/ims-kaas --name ims-kaas --set install.bootstrapIp="ip from step 1" --set install.airgapProxyIp="ip of step 2"
    ```
    *alternatively you can update `values.yaml` with IPs

### Utilizes:
- [kubebuilder](https://github.com/kubernetes-sigs/kubebuilder)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)
- [kustomize]()

## Build
#### one time setup of tools
- mac osx:
`make -f build/Makefile install-tools-darwin`

- linux:
`make -f build/Makefile install-tools-linux`

#### To generate code and binary:
- mac osx:
`make -f build/Makefile darwin`

- linux:
`make -f build/Makefile linux`

CRDs are generated in `./crd`
RBAC is generated in `./rbac`

Helm chart under `./deployments/helm/ims-kaas` gets updated with the right CRDs and RBAC

## Testing with Azure

Requirements:
- docker
- [opctl](https://opctl.io/docs/getting-started/opctl.html)
- [azure cli](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest)
    - install on mac osx: `brew install azure-cli`

#### Setup steps:
1. create the ssh key pair (requires rsa and 2048 bit)
    **no password**
    ```bash
    ssh-keygen -t rsa -b 2048 -f id_rsa
    ```

2. create `args.yml` file
    ```bash
    touch .opspec/args.yml
    ```
    add inputs:
    ```$xslt
      subscriptionId: <azure subscription id>
      loginId: <azure service principal id (must have permission to edit user permissions in subscription>
      loginSecret: <azure service principal secret>
      loginTenantId: <azure active directory id>
      sshKeyValue: <path to public key from step 1>
      sshPrivateKey: <path to private key from step 1>
      clusterAccountId: <azure service principal for in cluster resources (ex: load balancer creation)>
      clusterAccountSecret: <azure service principal secret>
      rootPassword: <root password for client vm>
      name: <prefix name to give to all resources> (ex: zaptest01)
    ```

3. from root directory of repo run
    ```bash
    opctl run build-azure
    ```
    first run takes 10/15 minutes. *this can be run multiple times

4. to get kubeconfig for central cluster:
    - login to azure via cli:
        ```bash
        az login
        ```
    - get kubeconfig from aks cluster:
        ```bash
        az aks get-credentials -n <name> -g <name>-group
        ```
        *replace with name from args.yml (step 2)

5. install bootstrap and connect to proxy:
    ```bash
    helm install deployments/helm/ims-kaas --name ims-kaas \
    --set install.operator=false \
    --set images.bootstrap.tag=0.1.17-local \
    --set install.bootstrapIp=10.240.0.6 \
    --set install.airgapProxyIp=10.240.0.7
    ```
    * check bootstrap latest tag at [quay.io](https://quay.io/repository/samsung_cnct/ims-kaas-bootstrap?tab=tags)
    * bootstrapIP is any node private ip (most likely: 10.240.0.4 thru .6)
    * to get airgapProxyIp run:
    ```bash
    az vm show -g <name>-group -n <name>-proxy -d --query publicIps --out tsv
    ```
6. locally start operator
    ```bash
    CMA_BOOTSTRAP_IP=10.240.0.6 CMA_NEXUS_PROXY_IP=10.240.0.7 ./ims-kaas
    ```

#### creating additional azure vm for testing clusters:
* to create additional vms:
```bash
opctl run create-vm
```
* this will create a new vm and provide the name/public ip

* TODO: return private IP also

#### cleanup azure:
* TODO: create azure-delete op.

* currently requires manually deleting resources / resource group manually in
the azure portal or cli

* resource group will be named `<name>-group` from `args.yml` file.

[generate-an-api-key]: https://docs.maas.io/2.1/en/manage-account#api-key
[packer_tool]: https://packer.io/downloads.html
[maas_website]: https://maas.io
[maas_user_data]: https://github.com/notjames/ims-kaas/tree/master/build/maas_deployment
