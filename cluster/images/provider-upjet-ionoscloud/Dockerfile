# All resources are reconciled in-process (Terraform Plugin SDK / Framework
# no-fork architecture, see config/external_name.go: CLIReconciledExternalNameConfigs
# is empty), so neither the OpenTofu/Terraform CLI nor the native provider plugin
# is needed at runtime. Only the provider binary is shipped.
FROM harbor.infra.cluster.ionos.com/imagefactory-public/distroless/static:nonroot@sha256:e072559f9ec00a3a83cdec2036bbfe6b7d0f442e7fcc07a4fe0f067c97593234

ARG TARGETOS
ARG TARGETARCH

ENV USER_ID=65532

# No TERRAFORM_* env vars: the versions reported in User-Agents are derived
# by the binary itself (see cmd/provider/main.go).

COPY --chown=${USER_ID}:${USER_ID} bin/${TARGETOS}_${TARGETARCH}/provider /usr/local/bin/provider

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
