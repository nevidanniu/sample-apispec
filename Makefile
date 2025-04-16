# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec
CONTROLLER_GEN = /go/bin/controller-gen
OPENAPI_GEN = /go/bin/openapi-gen
REPO=github.com/nevidanniu/sample-apispec

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Clientset-gen
clean-client-go:
	rm -rf client-go/* || true


manifests: ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
ifdef pkg
	$(CONTROLLER_GEN) crd paths="./${pkg}/..." output:crd:dir=./manifests/crd
else
	$(CONTROLLER_GEN) crd paths="./..." output:crd:dir=./manifests/crd
endif

generate: ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
ifdef pkg
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./${pkg}/..."
else
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."
endif

PKG_VERSION ?= v1alpha1
K8SAPIS ?= k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/version
openapigen: ## Generate openapi spec.
ifdef pkg
	$(OPENAPI_GEN) -i ${K8SAPIS},${REPO}/${pkg}/${PKG_VERSION} -p ${PKG_VERSION} -O zz_generated.openapi --go-header-file ./hack/boilerplate.go.txt -o ${pkg} -v5
else
	$(error Missing pkg variable. Specify pkg and version (optional as PKG_VERSION env), e.g. 'PKG_VERSION=v1alpha make pkg=core openapigen')
endif
