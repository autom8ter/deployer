clean: ## format and vet project
	@go mod tidy
	@go fmt ./...
	@go vet ./...

test: ## run all tests
	@go test -v ./...

install: ## install deployer
	@go install ./...

.PHONY: help
help:   ## show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'