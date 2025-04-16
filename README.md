# build code-generator
```bash
docker buildx build --progress=plain  -f docker/Dockerfile -t code-generator:latest .
```


# generate deepcopy
```bash
docker run -ti -v $(pwd):/data code-generator:latest
for PKG in core ldap; do PKG_VERSION=v1alpha1 make pkg=$PKG openapigen; make generate pkg=$PKG; done
```

# generate client-go
## create client-go structure
```bash
go mod init github.com/nevidanniu/sample-apispec/client-go
go get -x github.com/nevidanniu/sample-apispec
```

## update api ver & generate client-go
```bash
docker run -ti -v $(pwd):/data code-generator:latest
cd client-go
rm -rf applyconfigurations ssp listers informers

# should be tag, but sha commit also works
# 16d3e97 is sha commit https://github.com/nevidanniu/sample-apispec/commit/16d3e97b4371a4c51591ae652e77858d7d547977
go get -x -u github.com/nevidanniu/sample-apispec@16d3e97

# from go.mod
export API_TAG=v0.0.0-20250416105522-16d3e97b4371
/go/bin/applyconfiguration-gen --input-dirs $( paste -d, -s "/go/pkg/mod/github.com/nevidanniu/sample-apispec@$API_TAG/modules.txt" )  --output-base "."  --output-package "github.com/nevidanniu/sample-apispec/client-go/applyconfigurations" --trim-path-prefix "github.com/nevidanniu/sample-apispec/client-go" --go-header-file hack/boilerplate.go.txt -v 2

/go/bin/client-gen --clientset-name "ssp"  --input-base "github.com/nevidanniu/sample-apispec" --input $( paste -d, -s "/go/pkg/mod/github.com/nevidanniu/sample-apispec@$DRONE_TAG/groups.txt" ) --output-base "." --output-package "github.com/nevidanniu/sample-apispec/client-go" --trim-path-prefix "github.com/nevidanniu/sample-apispec/client-go" --apply-configuration-package "github.com/nevidanniu/sample-apispec/client-go/applyconfigurations" --go-header-file hack/boilerplate.go.txt -v 2

/go/bin/lister-gen --input-dirs $( paste -d, -s "/go/pkg/mod/github.com/nevidanniu/sample-apispec@$DRONE_TAG/modules.txt" ) --output-base "." --trim-path-prefix "github.com/nevidanniu/sample-apispec/client-go" --output-package "github.com/nevidanniu/sample-apispec/client-go/listers" --go-header-file hack/boilerplate.go.txt -v 2

/go/bin/informer-gen --input-dirs $( paste -d, -s "/go/pkg/mod/github.com/nevidanniu/sample-apispec@$DRONE_TAG/modules.txt" ) --output-base "." --trim-path-prefix "github.com/nevidanniu/sample-apispec/client-go" --versioned-clientset-package github.com/nevidanniu/sample-apispec/client-go/ssp --listers-package github.com/nevidanniu/sample-apispec/client-go/listers --output-package "github.com/nevidanniu/sample-apispec/client-go/informers" --go-header-file hack/boilerplate.go.txt -v 2
```

