ifndef GOARCH 
	GOARCH=$(shell go env GOARCH)
endif

ifndef GOOS 
	GOOS := $(shell go env GOOS)
endif

ifndef GOROOT
	GOROOT=$(shell go env GOROOT)
endif

ROOT_PACKAGE := github.com/simply-app/simply-console

.PHONY: tidy
tidy:
	GO111MODULE=on go mod tidy

.PHONY: build
build: fmt
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) \
		go build -ldflags "-s -w" \
		-trimpath -o bin/app cmd/account/server.go

.PHONY: serve-dev
serve-dev: build
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) \
		go run cmd/account/server.go

.PHONY: fmt
fmt:
	gofmt -l -w .

.PHONY: lint
lint:
	staticcheck ./...

.PHONY: test
test:
	GO111MODULE=on CGO_ENABLED=0 go test -v ./... -coverprofile=coverage.out
