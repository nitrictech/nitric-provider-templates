binaries: deploybin

# build runtime binary directly into the deploy director so it can be embedded directly into the deployment engine binary
# We only build a linux amd64 binary here to be packaged for cloud runtimes with docker
# More binaries can be distributed in future if required
runtimebin:
	@echo Building Custom Runtime Server
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/runtime-custom -ldflags="-s -w -extldflags=-static" ./cmd/runtime

predeploybin: runtimebin
	@cp bin/runtime-custom deploy/runtime/runtime-custom

deploybin: predeploybin
	@echo Building Custom Deployment Server
	@CGO_ENABLED=0 go build -o bin/deploy-custom -ldflags="-s -w -extldflags=-static" -ldflags="-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=ignore" ./cmd/deploy
	@rm deploy/runtime/runtime-custom

.PHONY: install
install: deploybin
	@echo installing custom deployment server to ${HOME}/.nitric/providers/custom/pulumi-0.0.1
	@mkdir -p ${HOME}/.nitric/providers/custom/
	@rm -f ${HOME}/.nitric/providers/custom/pulumi-0.0.1
	@cp bin/deploy-custom ${HOME}/.nitric/providers/custom/pulumi-0.0.1