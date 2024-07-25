#
# Credit:
#   This makefile was adapted from: https://github.com/vincentbernat/hellogopher/blob/feature/glide/Makefile
#
# Go environment
export GOPATH?=$(shell go env GOPATH)
BINDIR=$(CURDIR)/bin
# Build info
BINARY_NAME=sriovdp
BUILDDIR=$(CURDIR)/build
REPO_PATH=$(ORG_PATH)/$(PACKAGE)
BASE=$(GOPATH)/src/$(REPO_PATH)
PKGS = $(or $(PKG),$(shell cd $(BASE) && env GOPATH=$(GOPATH) go list ./... | grep -v "^$(PACKAGE)/vendor/"))
GOFILES = $(shell find . -name *.go | grep -vE "(\/vendor\/)|(_test.go)")
# Test artifacts and settings
TIMEOUT = 15
COVERAGE_DIR = $(CURDIR)/test/coverage
COVERAGE_MODE = atomic
COVERAGE_PROFILE = $(COVERAGE_DIR)/cover.out
# Docker image
DOCKERFILE?=$(CURDIR)/images/Dockerfile
TAG?=ghcr.io/k8snetworkplumbingwg/sriov-network-device-plugin
# Docker arguments - To pass proxy for Docker invoke it as 'make image HTTP_POXY=http://192.168.0.1:8080'
DOCKERARGS=
ifdef HTTP_PROXY
	DOCKERARGS += --build-arg http_proxy=$(HTTP_PROXY)
endif
ifdef HTTPS_PROXY
	DOCKERARGS += --build-arg https_proxy=$(HTTPS_PROXY)
endif

GO_BUILD_OPTS ?=
GO_LDFLAGS ?=
GO_FLAGS ?=
GO_TAGS ?=-tags no_openssl

ifdef STATIC
	GO_BUILD_OPTS+= CGO_ENABLED=0
	GO_LDFLAGS+= -extldflags \"-static\"
	GO_FLAGS+= -a
endif

V = 0
Q = $(if $(filter 1,$V),,@)

.PHONY: all
all: lint build test

.PHONY: create-dirs
create-dirs: $(BINDIR) $(BUILDDIR) $(COVERAGE_DIR)

$(BINDIR) $(BUILDDIR) $(COVERAGE_DIR): ; $(info Creating directory $@...)
	$Q mkdir -p $@

.PHONY: build
build: | $(BUILDDIR) ; $(info Building $(BINARY_NAME)...) @ ## Build SR-IOV Network device plugin
	$Q cd $(CURDIR)/cmd/$(BINARY_NAME) && $(GO_BUILD_OPTS) go build -ldflags '$(GO_LDFLAGS)' $(GO_FLAGS) -o $(BUILDDIR)/$(BINARY_NAME) $(GO_TAGS) -v
	$(info Done!)

GOLANGCI_LINT = $(BINDIR)/golangci-lint
GOLANGCI_LINT_VERSION ?= v1.55.2
$(GOLANGCI_LINT): | $(BINDIR) ; $(info  installing golangci-lint...)
	$Q[ -f $(GOLANGCI_LINT) ] || { \
	set -e ;\
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell dirname $(GOLANGCI_LINT)) $(GOLANGCI_LINT_VERSION) ;\
	}

MOCKERY = $(BINDIR)/mockery
$(MOCKERY): | $(BINDIR) ; $(info  installing mockery...)
	$(call go-install-tool,$(MOCKERY),github.com/vektra/mockery/v2@latest)

TEST_TARGETS := test-default test-verbose test-race
.PHONY: $(TEST_TARGETS) test
test-verbose: ARGS=-v            ## Run tests in verbose mode with coverage reporting
test-race:    ARGS=-race         ## Run tests with race detector
$(TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
$(TEST_TARGETS): test
test: ; $(info  running $(NAME:%=% )tests...) @ ## Run tests
	$Q go test -timeout $(TIMEOUT)s $(ARGS) $(PKGS)

test-xml: fmt lint | $(BASE) $(GO2XUNIT) ; $(info  running $(NAME:%=% )tests...) @ ## Run tests with xUnit output
	$Q cd $(BASE) && 2>&1 go test -timeout 20s -v $(TESTPKGS) | tee test/tests.output
	$(GO2XUNIT) -fail -input test/tests.output -output test/tests.xml

.PHONY: test-coverage test-coverage-tools
test-coverage-tools: | $(GOCOVMERGE) $(GOCOV) $(GOCOVXML) $(GCOV2LCOV)
test-coverage: COVERAGE_DIR := $(CURDIR)/test/coverage
test-coverage: fmt lint test-coverage-tools | $(BASE) ; $(info  Running coverage tests...) @ ## Run coverage tests
	$Q mkdir -p $(COVERAGE_DIR)/coverage
	$Q cd $(BASE) && for pkg in $(TESTPKGS); do \
		go test \
			-coverpkg=$$(go list -f '{{ join .Deps "\n" }}' $$pkg | \
					grep '^$(PACKAGE)/' | grep -v '^$(PACKAGE)/vendor/' | \
					tr '\n' ',')$$pkg \
			-covermode=$(COVERAGE_MODE) \
			-coverprofile="$(COVERAGE_DIR)/coverage/`echo $$pkg | tr "/" "-"`.cover" $$pkg ;\
	 done
	$Q $(GOCOVMERGE) $(COVERAGE_DIR)/coverage/*.cover > $(COVERAGE_PROFILE)
	$Q go tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	$Q $(GOCOV) convert $(COVERAGE_PROFILE) | $(GOCOVXML) > $(COVERAGE_XML)
	$Q $(GCOV2LCOV) -infile $(COVERAGE_PROFILE) -outfile $(COVERAGE_DIR)/lcov.info

.PHONY: lint
lint: | $(BASE) $(GOLINT) ; $(info  Running golint...) @ ## Run golint on all source files
	$Q cd $(BASE) && ret=0 && for pkg in $(PKGS); do \
		test -z "$$($(GOLINT) $$pkg | tee /dev/stderr)" || ret=1 ; \
	 done ; exit $$ret

.PHONY: fmt
fmt: ; $(info  running gofmt...) @ ## Run gofmt on all source files
	@ret=0 && for d in $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		gofmt -l -w $$d/*.go || ret=$$? ; \
	 done ; exit $$ret

.PHONY: deps-update
deps-update: ; $(info  Updating dependencies...) @ ## Update dependencies
	@go mod tidy && go mod vendor

.PHONY: image
image: ; $(info Building Docker image...) @ ## Build SR-IOV Network device plugin docker image
	$Q docker build -t $(TAG) -f $(DOCKERFILE)  $(CURDIR) $(DOCKERARGS)

.PHONY: clean
clean: ; $(info  Cleaning...) @ ## Cleanup everything
	@go clean --modcache --cache --testcache
	@rm -rf $(BUILDDIR)
	@rm -rf $(BINDIR)
	@rm -rf test/

.PHONY: generate-mocks
generate-mocks: | $(MOCKERY) ; $(info generating mocks...) @ ## Generate mocks
	$Q $(MOCKERY) --name=".*" --dir=pkg/types --output=pkg/types/mocks --recursive=false --log-level=debug
	$Q $(MOCKERY) --name=".*" --dir=pkg/utils --output=pkg/utils/mocks --recursive=false --log-level=debug
	$Q $(MOCKERY) --name=".*" --dir=pkg/cdi --output=pkg/cdi/mocks --recursive=false --log-level=debug

.PHONY: help
help: ; @ ## Display this help message
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'


# go-install-tool will 'go install' any package $2 and install it to $1.
define go-install-tool
$Q[ -f $(1) ] || { \
set -e ;\
echo "Downloading $(2)" ;\
GOBIN=$(BINDIR) go install $(2) ;\
}
endef
