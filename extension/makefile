# Edit these values to match your extension e.g. NAMESPACE := myorg, NAME := myextension
# The resulting provider name will be ${NAMESPACE}/${NAME}, e.g. myorg/myextension
NAMESPACE := example
NAME := extaws

# This is the local development version number, we recommend 0.0.1.
# This will only be used to install the extension locally for testing.
DEV_VERSION := 0.0.1

# === Don't modify below this line, unless you know what you're doing ===
FQN := ${NAMESPACE}/${NAME}

RUNTIME_BINARY := bin/${NAMESPACE}-${NAME}
RUNTIME_EMBED := cmd/deploy/runtime-bin

binaries: deploybin

# build the runtime binary
runtimebin:
	@echo Building ${FQN} runtime server
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${RUNTIME_BINARY} -ldflags="-s -w -extldflags=-static" ./cmd/runtime

# moves the runtime binary so it can be embedded into the deployment engine binary
predeploybin: runtimebin
	@cp ${RUNTIME_BINARY} ${RUNTIME_EMBED}

# build the deployment engine binary
deploybin: predeploybin
	@echo Building ${FQN} deployment server
	@CGO_ENABLED=0 go build -o bin/deploy-extension -ldflags="-s -w -extldflags=-static" -ldflags="-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=ignore" ./cmd/deploy
	@rm ${RUNTIME_EMBED}

# build and install the provider locally for testing
install: deploybin
	@echo installing ${FQN} to ${HOME}/.nitric/providers/${NAMESPACE}/${NAME}-${DEV_VERSION}
	@mkdir -p ${HOME}/.nitric/providers/${NAMESPACE}/
	@rm -f ${HOME}/.nitric/providers/${NAMESPACE}/${NAME}-${DEV_VERSION}
	@cp bin/deploy-extension ${HOME}/.nitric/providers/${NAMESPACE}/${NAME}-${DEV_VERSION}
	@echo "Installed successfully\n\nUpdate your stack file to use this provider.\n provider: ${FQN}@${DEV_VERSION}"

clean:
	@rm -rf bin

.PHONY: binaries runtimebin predeploybin deploybin install clean