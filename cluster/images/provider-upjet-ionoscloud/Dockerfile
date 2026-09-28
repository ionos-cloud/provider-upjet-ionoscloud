# Builder stage: assembles the provider binary, OpenTofu, and the native
# Terraform provider plugin. Needs a shell/unzip, so it stays off the final
# (distroless, shell-less) runtime image.
FROM harbor.infra.cluster.ionos.com/imagefactory-public/alpine:3.24@sha256:abef013291641fbe7099aa145d0a64e9561e443596ad4bf9bff9acd93f64be6d AS builder
RUN apk --no-cache add unzip

ARG TARGETOS
ARG TARGETARCH

ENV USER_ID=65532

## Provider-dependent configuration
ARG OPENTOFU_VERSION
ARG TERRAFORM_PROVIDER_SOURCE
ARG TERRAFORM_PROVIDER_VERSION
ARG TERRAFORM_PROVIDER_DOWNLOAD_NAME
ARG TERRAFORM_NATIVE_PROVIDER_BINARY
ARG TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX
## End of - Provider-dependent configuration

ENV PLUGIN_DIR=/terraform/provider-mirror/registry.opentofu.org/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_PROVIDER_VERSION}/${TARGETOS}_${TARGETARCH}

RUN mkdir -p ${PLUGIN_DIR}

ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider
ADD https://github.com/opentofu/opentofu/releases/download/v${OPENTOFU_VERSION}/tofu_${OPENTOFU_VERSION}_${TARGETOS}_${TARGETARCH}.zip /tmp
ADD ${TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX}/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip /tmp

RUN chmod +x /usr/local/bin/provider \
  && unzip /tmp/tofu_${OPENTOFU_VERSION}_${TARGETOS}_${TARGETARCH}.zip -d /usr/local/bin \
  && chmod +x /usr/local/bin/tofu \
  && ln -s /usr/local/bin/tofu /usr/local/bin/terraform \
  && rm /tmp/tofu_${OPENTOFU_VERSION}_${TARGETOS}_${TARGETARCH}.zip \
  && unzip /tmp/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip -d ${PLUGIN_DIR} \
  && chmod +x ${PLUGIN_DIR}/* \
  && rm /tmp/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip \
  && chown -R ${USER_ID}:${USER_ID} /terraform /usr/local/bin/provider /usr/local/bin/tofu

# Runtime stage: distroless, non-root by default (uid 65532), no shell.
FROM harbor.infra.cluster.ionos.com/imagefactory-public/distroless/static:nonroot@sha256:e072559f9ec00a3a83cdec2036bbfe6b7d0f442e7fcc07a4fe0f067c97593234

ARG TARGETOS
ARG TARGETARCH
ARG OPENTOFU_VERSION
ARG TERRAFORM_PROVIDER_SOURCE
ARG TERRAFORM_PROVIDER_VERSION
ARG TERRAFORM_NATIVE_PROVIDER_BINARY

ENV USER_ID=65532
ENV PLUGIN_DIR=/terraform/provider-mirror/registry.opentofu.org/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_PROVIDER_VERSION}/${TARGETOS}_${TARGETARCH}
ENV TF_CLI_CONFIG_FILE=/terraform/.terraformrc
ENV TF_FORK=0

# Provider controller needs these environment variables at runtime
ENV TERRAFORM_VERSION=${OPENTOFU_VERSION}
ENV TERRAFORM_PROVIDER_SOURCE=${TERRAFORM_PROVIDER_SOURCE}
ENV TERRAFORM_PROVIDER_VERSION=${TERRAFORM_PROVIDER_VERSION}
ENV TERRAFORM_NATIVE_PROVIDER_PATH=${PLUGIN_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}

COPY --from=builder --chown=${USER_ID}:${USER_ID} /usr/local/bin/provider /usr/local/bin/provider
COPY --from=builder --chown=${USER_ID}:${USER_ID} /usr/local/bin/tofu /usr/local/bin/tofu
COPY --from=builder --chown=${USER_ID}:${USER_ID} /usr/local/bin/terraform /usr/local/bin/terraform
COPY --from=builder --chown=${USER_ID}:${USER_ID} ${PLUGIN_DIR} ${PLUGIN_DIR}
COPY --chown=${USER_ID}:${USER_ID} terraformrc.hcl ${TF_CLI_CONFIG_FILE}

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
