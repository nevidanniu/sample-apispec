FROM golang:1.23.4 as builder
ARG CONTROLLER_GEN_VER=v0.14.0
ARG CODE_GENERATORS_VER=v0.26.1
ARG SEMVER_VER=v1.1.0
ARG SETUP_ENVTEST_VER=v0.0.0-20230221102149-f6f37e6cc1ec


RUN for GENERATOR in "defaulter-gen" "client-gen" "lister-gen" "informer-gen" "deepcopy-gen" "openapi-gen" "applyconfiguration-gen"; do go install k8s.io/code-generator/cmd/${GENERATOR}@${CODE_GENERATORS_VER}; done
RUN go install sigs.k8s.io/controller-tools/cmd/controller-gen@${CONTROLLER_GEN_VER}
RUN go install github.com/maykonlf/semver-cli/cmd/semver@${SEMVER_VER}
RUN go install sigs.k8s.io/controller-runtime/tools/setup-envtest@${SETUP_ENVTEST_VER}


FROM golang:1.23.4
COPY --from=builder /go/bin /go/bin
WORKDIR /data