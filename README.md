# application-aware-controller

This repository implements a simple controller for watching Foo resources as
defined with a CustomResourceDefinition (CRD).

**Note:** go-get or vendor this package as `k8s.io/application-aware-controller`.

This particular example demonstrates how to perform basic operations such as:

* How to register a new custom resource (custom resource type) of type `Foo` using a CustomResourceDefinition.
* How to create/get/list instances of your new resource type `Foo`.
* How to setup a controller on resource handling create/update/delete events.

It makes use of the generators in [k8s.io/code-generator](https://github.com/kubernetes/code-generator)
to generate a typed client, informers, listers and deep-copy functions. You can
do this yourself using the `./hack/update-codegen.sh` script.

The `update-codegen` script will automatically generate the following files &
directories:

* `pkg/apis/appawarecontroller/v1/zz_generated.deepcopy.go`
* `pkg/generated/`

Changes should not be made to these files manually, and when creating your own
controller based off of this implementation you should not copy these files and
instead run the `update-codegen` script to generate your own.

## Details

The ahpa controller uses [client-go library](https://github.com/kubernetes/client-go/tree/master/tools/cache) extensively.
The details of interaction points of the ahpa controller with various mechanisms from this library are
explained [here](docs/controller-client-go.md).

## Fetch application-aware-controller and its dependencies

Like the rest of Kubernetes, application-aware-controller has used
[godep](https://github.com/tools/godep) and `$GOPATH` for years and is
now adopting go 1.11 modules.  There are thus two alternative ways to
go about fetching this demo and its dependencies.

### Fetch with godep

When NOT using go 1.11 modules, you can use the following commands.

```sh
go get -d k8s.io/application-aware-controller
cd $GOPATH/src/k8s.io/application-aware-controller
godep restore
```

### When using go 1.11 modules

When using go 1.11 modules (`GO111MODULE=on`), issue the following
commands --- starting in whatever working directory you like.

```sh
git clone https://github.com/nkwangjun/application-aware-controller.git
cd application-aware-controller
```

Note, however, that if you intend to
generate code then you will also need the
code-generator repo to exist in an old-style location.  One easy way
to do this is to use the command `go mod vendor` to create and
populate the `vendor` directory.

### A Note on kubernetes/kubernetes

If you are developing Kubernetes according to
https://github.com/kubernetes/community/blob/master/contributors/guide/github-workflow.md
then you already have a copy of this demo in
`kubernetes/staging/src/k8s.io/application-aware-controller` and its dependencies
--- including the code generator --- are in usable locations
(valid for all Go versions).

## Purpose

This is an example of how to build a kube-like controller with a single type.

## Running

**Prerequisite**: Since the application-aware-controller uses `apps/v1` deployments, the Kubernetes cluster version should be greater than 1.9.

```sh
# assumes you have a working kubeconfig, not required if operating in-cluster
go build -o application-aware-controller .
./application-aware-controller -kubeconfig=$HOME/.kube/config

# create a CustomResourceDefinition
kubectl create -f artifacts/examples/crd-app-aware-hpa.yaml

# create a custom resource of type ahpa
kubectl create -f artifacts/examples/example-deployment-appaware-hpa.yaml

# check deployments created through the custom resource
kubectl get ahpas
```
