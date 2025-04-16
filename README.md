# build code-generator
```bash
docker buildx build --progress=plain  -f docker/Dockerfile -t code-generator:latest .
```


# generate deepcopy
```bash
docker run -ti -v $(pwd):/data code-generator:latest
for PKG in core ldap; do PKG_VERSION=v1alpha1 make pkg=$PKG openapigen; make generate pkg=$PKG; done
```