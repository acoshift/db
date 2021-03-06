SHELL := /bin/bash

WAPPER  ?= all
DB_HOST ?= 127.0.0.1

TEST_FLAGS ?=

export DB_HOST
export WRAPPER

benchmark-lib:
	go test -v -benchtime=500ms -bench=. ./lib/...

benchmark-internal:
	go test -v -benchtime=500ms -bench=. ./internal/...

benchmark: benchmark-lib benchmark-internal

test-lib:
	go test -v ./lib/...

test-internal:
	go test -v ./internal/...

test-libs: test-lib test-internal

test-adapters: test-adapter-postgresql

reset-db:
	$(MAKE) -C postgresql reset-db

test: test-adapters test-libs

test-adapter-%:
	$(MAKE) -C $* test || exit 1;
