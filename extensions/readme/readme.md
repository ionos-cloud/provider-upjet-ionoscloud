# Quickstart

This guide walks through the process to install Upbound Universal Crossplane and the Ionoscloud provider.

## Install the Up command-line
Download and install the Upbound `up` command-line.

```shell
curl -sL "https://cli.upbound.io" | sh
sudo mv up /usr/local/bin/
```

Verify the version of `up` with `up --version`

```shell
$ up version --client
Client:
Version:     v0.33.0
```

More information about the Up command-line is available in the [Upbound Up
documentation](https://docs.upbound.io/cli/).

## Install Upbound Universal Crossplane
Install Upbound Universal Crossplane (UXP) with the Up command-line `up uxp
install` command.

```shell
$ up uxp install
UXP 1.17.1-up.1 installed
```

Verify all UXP pods are `Running` with `kubectl get pods`

```shell
kubectl get pods
NAME                                       READY   STATUS    RESTARTS   AGE
crossplane-77ff754998-4l8xb                1/1     Running   0          21s
crossplane-rbac-manager-79b8bdd6d8-ml6ft   1/1     Running   0          21s
```

Find more information in the [Upbound UXP
documentation](https://docs.upbound.io/uxp/).


## Install the Ionoscloud provider

Install the Ionoscloud provider into the Kubernetes cluster with a Kubernetes
configuration file.

```shell
cat <<EOF | kubectl apply -f -
          apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
name: provider-upjet-ionoscloud
spec:
package: xpkg.upbound.io/ionos-cloud/provider-upjet-ionoscloud:v0.1.4
EOF
```


Verify again all pods are `Running` with `kubectl get pods`

```shell
kubectl get pods
NAME                                       READY   STATUS    RESTARTS   AGE
crossplane-7b9585ddfd-6rzm6                              1/1     Running   0          7m19s
crossplane-rbac-manager-699c59799d-mwtkl                 1/1     Running   0          7m19s
provider-upjet-ionoscloud-1f720e145fea-8994d7b58-w7xxb   1/1     Running   0          4m32s

```

Apply this configuration with `kubectl apply -f`.

After installing the provider, verify the install with `kubectl get providers`.

```shell
NAME                        INSTALLED   HEALTHY   PACKAGE                            AGE
provider-upjet-ionoscloud   True        True      xpkg.upbound.io/ionos-cloud/provider-upjet-ionoscloud:v0.1.4   5m35s
```

## Create a Kubernetes secret for Ionoscloud

### Using a Ionoscloud token
```shell
kubectl -n upbound-system create secret generic ionoscloud-secret --from-literal=credentials="{\"token\":\"${IONOS_TOKEN}\"}"
```

### With Object Storage access
```shell
kubectl -n upbound-system create secret generic ionoscloud-secret --from-literal=credentials="{\"token\":\"${IONOS_TOKEN}\",\"s3_access_key\":\"${IONOS_S3_ACCESS_KEY}\",\"s3_secret_key\":\"${IONOS_S3_SECRET_KEY}\"}"
```

### Using the Ionoscloud username and password(not recommended)
```shell
kubectl -n upbound-system create secret generic ionoscloud-secret --from-literal=credentials="{\"user\":\"${IONOS_USERNAME}\",\"password\":\"${IONOS_PASSWORD}\"}"
```

### Fields that can be set using the credentials secret:
```yaml
{
"token": "string",
"user": "string",
"password": "string",
"s3_access_key": "string",
"s3_secret_key": "string",
"s3_region": "string",
"endpoint": "string",
"contract_number": "string",
}

View the secret with `kubectl describe secret`
```shell
kubectl describe secret ionoscloud-secret -n upbound-system
Name:         ionoscloud-secret
Namespace:    upbound-system
Labels:       <none>
  Annotations:  <none>

    Type:  Opaque

    Data
    ====
    credentials:  1274 bytes
    ```
    _Note:_ the size may be larger or smaller depending on which credentials you supplied

    ## Create a ProviderConfig
    Create a `ProviderConfig` Kubernetes configuration file to attach the Ionoscloud
    credentials to the provider.

    ```shell
    cat <<EOF | kubectl apply -f -
              apiVersion: upjet-ionoscloud.ionoscloud.io/v1beta1
    kind: ProviderConfig
    metadata:
    name: default
    spec:
    credentials:
    source: Secret
    secretRef:
    namespace: upbound-system
    name: ionoscloud-secret
    key: credentials
    EOF
    ```

    The `spec.secretRef` describes the parameters of the secret to use.
    * `namespace` is the Kubernetes namespace the secret is in.
    * `name` is the name of the Kubernetes `secret` object.
    * `key` is the `Data` field from `kubectl describe secret`.

    Apply this configuration with `kubectl apply -f`.

    Verify the `ProviderConfig` with `kubectl describe providerconfigs`.

    ```shell
    kubectl describe providerconfigs
    Name:         default
    Namespace:
    API Version:  upjet-ionoscloud.ionoscloud.io/v1beta1
    Kind:         ProviderConfig
    # Output truncated
    Spec:
    Credentials:
    Secret Ref:
    Key:        credentials
    Name:       ionoscloud-secret
    Namespace:  upbound-system
    Source:       Secret
    ```

    **Note:** the `ProviderConfig` install fails and Kubernetes returns an error if
    the `Provider` isn't installed.

    ```shell
    kubectl apply -f providerconfig.yml
    error: resource mapping not found for name: "default" namespace: "" from "providerconfig.yml": no matches for kind "ProviderConfig" in version "upjet-ionoscloud.ionoscloud.io/v1beta1"
    ensure CRDs are installed first
    ```

    ## Create a managed resource
    Create a managed resource to verify the provider is functioning.

    This example creates a Ionoscloud datacenter.

    Either from a file:
    ```shell
    kubectl apply -f examples/datacenter.yaml
    ```

    Or directly from the command line:

    ```shell
    cat <<EOF | kubectl apply -f -
              apiVersion: compute.ionoscloud.io/v1alpha1
    kind: Datacenter
    metadata:
    name: example
    spec:
    forProvider:
    description: Example datacenter description
    location: de/fra
    name: Datacenter Example
    secAuthProtection: false
    EOF
    ```

    Use `kubectl get datacenters` to verify datacenter creation.

    ```shell
    kubectl get datacenters
    NAME      SYNCED   READY   EXTERNAL-NAME                          AGE
    example   True     True    5bc1d7a9-2b1d-488c-b16f-8b946017ecde   5m3s
    ```
    The datacenter creation is finalized when the values `READY` and `SYNCED` are `True`.

    ## Delete the managed resource
    ```shell
    kubectl delete datacenters.compute.ionoscloud.io example
    datacenter.compute.ionoscloud.io "example" deleted

    kubectl get datacenters
    No resources found
    ```

    ## Troubleshooting
    Check events to see if there are any issues with the provider installation or managed resources.
    ```shell
    kubectl get events
    ```


    If you encounter issues, check the logs of the provider pod with `kubectl logs`.

    ```shell
    kubectl -n crossplane-system logs  provider-upjet-ionoscloud-1f720e145fea-8994d7b58-gh57n
    ```

    For more information on troubleshooting, see the [Crossplane troubleshooting guide](https://docs.crossplane.io/latest/guides/troubleshoot-crossplane/).