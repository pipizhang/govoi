SHELL := /bin/bash

help: ## This help message
	@echo "Usage: make [target]"
	@echo "Commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY:test
test: ## Run go test
	@(scripts/test)

benchmark: ## Run go benchmark
	@exec go test -bench=. -v -run=none

.PHONY:install
install: ## Install dependencies
	@exec glide install
