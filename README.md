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
go get -x -u github.com/nevidanniu/sample-apispec
```