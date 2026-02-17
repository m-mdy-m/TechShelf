SHELL := /bin/bash

GO            ?= go
PROJECT_DIR   := $(CURDIR)
MODULE        ?= $(shell $(GO) list -m 2>/dev/null)
BINARY        := shelf
CMD_DIR       ?= ./cmd
MAIN_PKG      ?= ./cmd
OUT_DIR       ?= bin
PREFIX        ?= /usr/local

VERSION       ?= $(shell git describe --tags --match 'v[0-9]*' --always --dirty 2>/dev/null || echo dev)
GIT_COMMIT    ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
BUILD_DATE    ?= $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

LDFLAGS       ?= -X 'github.com/m-mdy-m/TechShelf/internal/command.Version=$(VERSION)' \
                 -X 'github.com/m-mdy-m/TechShelf/internal/command.GitCommit=$(GIT_COMMIT)' \
                 -X 'github.com/m-mdy-m/TechShelf/internal/command.BuildDate=$(BUILD_DATE)' \
                 -s -w

.PHONY: help build test lint fmt clean install run version deps tidy build-all

help:
	@echo "TechShelf CLI"
	@echo "  make build      Build binary  →  bin/shelf"
	@echo "  make run        Run CLI"
	@echo "  make test       Run go test ./..."
	@echo "  make lint       Run go vet ./..."
	@echo "  make fmt        Run gofmt -w"
	@echo "  make tidy       Run go mod tidy"
	@echo "  make install    Install to $(PREFIX)/bin/shelf"
	@echo "  make version    Print build metadata"

build: deps
	@mkdir -p $(OUT_DIR)
	$(GO) build -ldflags "$(LDFLAGS)" -o $(OUT_DIR)/$(BINARY) $(MAIN_PKG)

run:
	$(GO) run $(MAIN_PKG)

test: deps
	$(GO) test ./...

lint: deps
	$(GO) vet ./...

fmt:
	@gofmt -w .

clean:
	@rm -rf $(OUT_DIR)

install: build
	@install -m 0755 $(OUT_DIR)/$(BINARY) $(PREFIX)/bin/$(BINARY)

version:
	@echo "binary    : $(BINARY)"
	@echo "version   : $(VERSION)"
	@echo "commit    : $(GIT_COMMIT)"
	@echo "build date: $(BUILD_DATE)"

deps:
	$(GO) mod download

tidy:
	$(GO) mod tidy

build-all: build
	GOOS=linux   GOARCH=amd64 $(GO) build -ldflags "$(LDFLAGS)" -o $(OUT_DIR)/$(BINARY)-linux-amd64   $(MAIN_PKG)
	GOOS=linux   GOARCH=arm64 $(GO) build -ldflags "$(LDFLAGS)" -o $(OUT_DIR)/$(BINARY)-linux-arm64   $(MAIN_PKG)
	GOOS=darwin  GOARCH=arm64 $(GO) build -ldflags "$(LDFLAGS)" -o $(OUT_DIR)/$(BINARY)-darwin-arm64  $(MAIN_PKG)
	GOOS=windows GOARCH=amd64 $(GO) build -ldflags "$(LDFLAGS)" -o $(OUT_DIR)/$(BINARY)-windows-amd64.exe $(MAIN_PKG)