BROWSERTEST_VERSION = v0.3.5
LINT_VERSION = 1.41.1
GO_BIN = $(shell printf '%s/bin' "$$(go env GOPATH)")

.PHONY: all
all: lint test

.PHONY: lint-deps
lint-deps:
	@if ! which golangci-lint >/dev/null || [[ "$$(golangci-lint version 2>&1)" != *${LINT_VERSION}* ]]; then \
		curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "${GO_BIN}" v${LINT_VERSION}; \
	fi

.PHONY: lint
lint: lint-deps
	GOOS=js GOARCH=wasm "${GO_BIN}/golangci-lint" run

.PHONY: test-deps
test-deps:
	@if [ ! -f "${GO_BIN}/go_js_wasm_exec" ]; then \
		set -ex; \
		go install github.com/agnivade/wasmbrowsertest@${BROWSERTEST_VERSION}; \
		ln -s "${GO_BIN}/wasmbrowsertest" "${GO_BIN}/go_js_wasm_exec"; \
	fi

.PHONY: test
test: test-deps
	GOOS=js GOARCH=wasm go test -cover ./...
	go test -race ./...
