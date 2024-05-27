# ------------------------------------------------------------------------------
# Variables
# ------------------------------------------------------------------------------
OPERATOR := operator

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

# ------------------------------------------------------------------------------
# Operator
# ------------------------------------------------------------------------------
.PHONY: operator
operator: go-build-operator go-run-operator

.PHONY: go-build-operator
go-build-operator:
	@echo " > Building $(OPERATOR) binary"
	GOBIN=$(GOBIN) go build -o ./bin/$(OPERATOR) ./cmd/$(OPERATOR)

.PHONY: go-run-operator
go-run-operator:
	@echo " > Running $(OPERATOR)"
	@-$(GOBIN)/$(OPERATOR)

# ------------------------------------------------------------------------------
# Code check
# ------------------------------------------------------------------------------
.PHONY: build
build:
	@echo "  >  Building $(OPERATOR) binary..."
	go build -o ./bin/$(OPERATOR) ./cmd/$(OPERATOR)

.PHONY: unit-tests
unit-tests:
	@echo "  >  Running unit tests"
	go clean -testcache
	go test -coverprofile=coverage -cover -v ./cmd... ./internal... ./pkg...

.PHONY: unit-tests-coverage
unit-tests-coverage: unit-tests
	go tool cover -html=coverage -o cover_out.html
	open cover_out.html

.PHONY: auto-tests
auto-tests:
	@echo "  >  Running auto tests"
	go clean -testcache
	go test -cover -v ./test...

.PHONY: lint
lint: go-lint-install go-lint

.PHONY: go-lint-install
go-lint-install:
ifeq (,$(shell which golangci-lint))
	@echo "  >  Installing golint"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- v1.57.2
endif

.PHONY: go-lint
go-lint:
	@echo "  >  Running golint"
	golangci-lint run ./...

abi-gen:
	abigen --abi=./contracts/abi/dittoentrypoint/DittoEntryPoint.json --pkg=dittoentrypoint --out=./contracts/gen/dittoentrypoint/dittoentrypoint.go
