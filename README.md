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
```bash
docker run -ti -v $(pwd):/data code-generator:latest
cd client-go

```

# create client-go structure
```bash
go mod init github.com/nevidanniu/sample-apispec/client-go
# ad7dbfb is sha commit https://github.com/nevidanniu/sample-apispec/commit/ad7dbfb1911e61c61a9f457ea9987a43c86e1864
go get -x -u github.com/nevidanniu/sample-apispec@ad7dbfb

export API_TAG=v0.0.0-20250416103906-ad7dbfb1911e
/go/bin/applyconfiguration-gen --input-dirs $( paste -d, -s "/go/pkg/mod/github.com/nevidanniu/sample-apispec@$API_TAG/modules.txt" )  --output-base "."  --output-package "github.com/nevidanniu/sample-apispec/client-go/applyconfigurations" --trim-path-prefix "github.com/nevidanniu/sample-apispec/client-go" --go-header-file hack/boilerplate.go.txt -v 2
```