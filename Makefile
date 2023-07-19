PHONY: ci, info build test quality lint security static-analysis

SHELL=bash
PROJECT_ROOT:=$(shell git rev-parse --show-toplevel)
GIT_BRANCH:=$(shell git rev-parse --abbrev-ref HEAD | sed "s*/*-*g")
GO_PROJECT_NAME:=$(shell go mod edit -json | jq -r .Module.Path | xargs basename)
MAIN_DIR:=$(PROJECT_ROOT)

LIB_DIR:=$(PROJECT_ROOT)/middleware

MAKE_LIB:=$(PROJECT_ROOT)/scripts
-include $(MAKE_LIB)/help.mk $(MAKE_LIB)/print.mk

ci: info build test quality security ## Run the same steps as the build pipeline
	$(call dump_header, $(DEBUG_DEFAULT_HDR_WIDTH), "$@", $(DEBUG_DEFAULT_HDR_CHAR), $(DEBUG_DEFAULT_HDR_FG))

pre-reqs:
	$(call dump_header, $(DEBUG_DEFAULT_HDR_WIDTH), "shadow tool install", $(DEBUG_DEFAULT_HDR_CHAR), $(DEBUG_DEFAULT_HDR_FG))
	go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	$(call dump_header, $(DEBUG_DEFAULT_HDR_WIDTH), "staticcheck install", $(DEBUG_DEFAULT_HDR_CHAR), $(DEBUG_DEFAULT_HDR_FG))
	go install honnef.co/go/tools/cmd/staticcheck@2023.1.3

info: ## get tool version info
	$(call dump_header, $(DEBUG_DEFAULT_HDR_WIDTH), "$@", $(DEBUG_DEFAULT_HDR_CHAR), $(DEBUG_DEFAULT_HDR_FG))
	go version

build: ## build the code
	$(call dump_header, $(DEBUG_DEFAULT_HDR_WIDTH), "$@", $(DEBUG_DEFAULT_HDR_CHAR), $(DEBUG_DEFAULT_HDR_FG))
	go build $(GO_BUILD_FLAGS) $(MAIN_DIR)/...

build-examples:
	go build $(GO_BUILD_FLAGS) $(MAIN_DIR)/cmd/examples/chi/chi.go
	go build $(GO_BUILD_FLAGS) $(MAIN_DIR)/cmd/examples/gin/gin.go
	go build $(GO_BUILD_FLAGS) $(MAIN_DIR)/cmd/examples/stdlib/stdlib.go

test: ## Run the unit tests
	$(call dump_header, $(DEBUG_DEFAULT_HDR_WIDTH), "$@", $(DEBUG_DEFAULT_HDR_CHAR), $(DEBUG_DEFAULT_HDR_FG))
	go test -cover -failfast ./...

test-race: ## Run the unit tests, check for race conditions
	go test --short -cover ./... -v -race

quality: lint static-analysis ## TODO
	$(call dump_header, $(DEBUG_DEFAULT_HDR_WIDTH), "$@", $(DEBUG_DEFAULT_HDR_CHAR), $(DEBUG_DEFAULT_HDR_FG))

lint: ## TODO
	$(call dump_header, $(DEBUG_DEFAULT_HDR_WIDTH), "$@", $(DEBUG_DEFAULT_HDR_CHAR), $(DEBUG_DEFAULT_HDR_FG))

security: ## TODO
	$(call dump_header, $(DEBUG_DEFAULT_HDR_WIDTH), "$@", $(DEBUG_DEFAULT_HDR_CHAR), $(DEBUG_DEFAULT_HDR_FG))

static-analysis: ## Perform static analysis on the code
	$(call dump_header, $(DEBUG_DEFAULT_HDR_WIDTH), "$@", $(DEBUG_DEFAULT_HDR_CHAR), $(DEBUG_DEFAULT_HDR_FG))
	go vet ./...
	shadow ./...

