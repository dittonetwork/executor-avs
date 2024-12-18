# ------------------------------------------------------------------------------
# Variables
# ------------------------------------------------------------------------------
OPERATOR := operator

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOLINT_PATH := $(GOBIN)/golangci-lint

GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

BUILD_MODE?=debug
ifeq ($(BUILD_MODE),debug)
	BUILD_FLAGS := -gcflags="all=-N -l"
	CGO_ENABLED := 1
else
	BUILD_FLAGS := -ldflags="-s -w"
	CGO_ENABLED := 0
endif

.PHONY: toolchain
toolchain:
	go install github.com/vektra/mockery/v2@v2.43.2
	go install github.com/mailru/easyjson/...@v0.7.7
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2

# ------------------------------------------------------------------------------
# Functions
# ------------------------------------------------------------------------------
define go_build
@echo " > Building $(1) in $(BUILD_MODE) mode..."
@CGO_ENABLED=$(CGO_ENABLED) GOBIN=$(GOBIN) GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(BUILD_FLAGS) -o ./bin/$(1) ./cmd/$(1)
endef

# ------------------------------------------------------------------------------
# Operator
# ------------------------------------------------------------------------------
.PHONY: operator
operator: go-build-operator go-run-operator

.PHONY: go-build-operator
go-build-operator:
	$(call go_build,$(OPERATOR))

.PHONY: go-run-operator
go-run-operator:
	@echo " > Running $(OPERATOR)"
	@-$(GOBIN)/$(OPERATOR)

# ------------------------------------------------------------------------------
# Code check
# ------------------------------------------------------------------------------
.PHONY: build
build: go-build-operator

.PHONY: test
test: unit-tests

.PHONY: unit-tests
unit-tests:
	@echo "  >  Running unit tests"
	go clean -testcache
	go test -coverprofile=coverage -cover -v ./cmd... ./pkg...

.PHONY: coverage-tests
coverage-tests:
	go tool cover -html=coverage -o cover_out.html
	open cover_out.html

.PHONY: auto-tests
auto-tests:
	@echo "  >  Running auto tests"
	go clean -testcache
	go test -cover -v ./test...

.PHONY: lint
lint: go-lint-install go-lint

GOLINT_VERSION := 1.62.2
.PHONY: go-lint-install
go-lint-install:
	$(eval CURRENT_VERSION := $(shell $(GOLINT_PATH) --version 2>/dev/null | grep -oP 'version \K[0-9.]+'))
	@if [ "$(CURRENT_VERSION)" != "$(GOLINT_VERSION)" ]; then \
		echo "  >  Installing or updating golangci-lint to version $(GOLINT_VERSION)"; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v$(GOLINT_VERSION); \
		echo "Installation complete. Version now: $$($(GOLINT_PATH) --version)"; \
	else \
		echo "golangci-lint is already installed at version $(GOLINT_VERSION)"; \
	fi


.PHONY: go-lint
go-lint:
	@echo "  >  Running golint"
	@-$(GOLINT_PATH) run ./...

.PHONY: abi-gen
abi-gen:
	go generate ./cmd/operator/internal/contracts

.PHONY: mock-gen
mock-gen:
	go generate ./cmd/operator/internal/adapters/dittoentrypoint/...
	go generate ./cmd/operator/internal/services/...
