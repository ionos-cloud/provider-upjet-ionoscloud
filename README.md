# Provider IonosCloud

`provider-upjet-ionoscloud` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
IonosCloud API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/ionos-cloud/provider-upjet-ionoscloud):
```
up ctp provider install ionos-cloud/provider-upjet-ionoscloud:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-upjet-ionoscloud
spec:
  package: ionos-cloud/provider-upjet-ionoscloud:v0.1.0
EOF
```

#### The API reference can be checked [here](https://doc.crds.dev/github.com/ionos-cloud/provider-upjet-ionoscloud).
### For more details on installation and configuration see the [Quickstart guide](docs/Quickstart.md)


## Developing

Run code-generation pipeline:
```console
make generate
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/ionos-cloud/provider-upjet-ionoscloud/issues).
