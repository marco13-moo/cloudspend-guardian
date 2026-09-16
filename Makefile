.PHONY: build run test lint fmt-check verify

# Keep local and CI builds reproducible in restricted development environments.
GOCACHE ?= /tmp/cloudspend-guardian-go-cache
GOMODCACHE ?= /tmp/cloudspend-guardian-go-mod-cache
export GOCACHE
export GOMODCACHE

build:
	go build ./...

run:
	go run ./cmd/server

test:
	go test -race -cover ./...

lint:
	go vet ./...

fmt-check:
	@test -z "$$(gofmt -l .)" || (echo "Go files require formatting" && gofmt -d . && exit 1)

verify: fmt-check lint test build
