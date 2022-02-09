# provider-confluent

`provider-confluent` is a minimal [Crossplane](https://crossplane.io/) Provider
that is meant to be used as a template for implementing new Providers. It comes
with the following features that are meant to be refactored:

- A `ProviderConfig` type that only points to a credentials `Secret`.
- A `MyType` resource type that serves as an example managed resource.
- A managed resource controller that reconciles `MyType` objects and simply
  prints their configuration in its `Observe` method.

## Developing

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build image:

```console
make image
```

Push image:

```console
make push
```

Build binary:

```console
make build
```


## Useful commands

Need to build an amd64 Linux image but you're on say.. an M1 Mac?

```shell
make build BUILD_PLATFORMS=linux_amd64 PLATFORMS=linux_amd64
```

Want to skip lint during a `make build`?

```shell
make build nolint=1
```
