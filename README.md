# provider-confluent

`provider-confluent` is a minimal [Crossplane](https://crossplane.io/) Provider
that is meant to be used as a template for implementing new Providers. It comes
with the following features that are meant to be refactored:

- A `ProviderConfig` type that only points to a credentials `Secret`.
- A `MyType` resource type that serves as an example managed resource.
- A managed resource controller that reconciles `MyType` objects and simply
  prints their configuration in its `Observe` method.


## Developing

### Prerequisites

- Linux/Unix development environment

- Install Confluent CLI:

> Note: Run the following commands in Bash

```
curl -sL --http1.1 https://cnfl.io/cli | sh -s -- -b /usr/local/bin <Confluent_CLI_Version>
```

Replace <Confluent_CLI_Version> with the correct version in the docker image /cluster/images/provider-confluent-controller/Dockerfile

- Define .env file using example

### Additional recommended steps
- Install [direnv](https://direnv.net/)
- Install [Minikube](https://minikube.sigs.k8s.io/)


### Running local development

Running the Confluent provider in local development environment can be either be using manual steps or automated steps.
The automated steps can be followed only when the additional recommend tools have been installed from the previous section.

Manual steps:
1) Make sure that you have access to a Kubernetes cluster and the correct context is used for Kubectl.

2) Ensure the list of environment variables in the .env file is exported
3) Run this command:
```console
make dev
```
This will build and run the confluent provider go code inside the configured Kubernetes cluster
Automated steps:
1) Run the script from this path ./scripts/setup.sh
This will do the followings:
- Start a Minikube instance and
- setup Crossplane and
- configure the required ProviderConfig for Confluent provider.

2) Execute this command:
```console
make dev
```
This will build and run the confluent provider go code inside the configured Minikube cluster

View the available make commands in the Makefile

## Useful commands

Need to build an amd64 Linux image but you're on say.. an M1 Mac?

```shell
make build BUILD_PLATFORMS=linux_amd64 PLATFORMS=linux_amd64
```

Want to skip lint during a `make build`?

```shell
make build nolint=1
```



