CONTRACTS := IAVS IAVSDirectory


.PHONY: version
version: ## Print tool versions.
	@forge --version
	@abigen --version

.PHONY: build
build: version ## Build contracts.
	forge build

.PHONY: bindings
bindings: build ## Generate golang contract bindings.
	./gen-abi.sh $(CONTRACTS)

image:
	@echo "build docker image"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main_amd64 .
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -o main_arm64 .
	docker buildx build --platform linux/amd64,linux/arm64 --push -t chainbase-registry.ap-southeast-1.cr.aliyuncs.com/network/chainbase-node:${tag}  .
	rm ./main_amd64
	rm ./main_arm64