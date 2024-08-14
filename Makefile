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