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

.PHONY: build-account
build-account: fmt
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) \
		go build -ldflags "-s -w" \
		-trimpath -o bin/account-app cmd/account/server.go

.PHONY: build-address
build-address: fmt
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) \
		go build -ldflags "-s -w" \
		-trimpath -o bin/address-app cmd/address/server.go

.PHONY: serve-account-dev
serve-account-dev: build-account
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) \
		go run cmd/account/server.go

.PHONY: serve-address-dev
serve-address-dev: build-address
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) \
		go run cmd/address/server.go

.PHONY: fmt
fmt:
	gofmt -l -w .

.PHONY: lint
lint:
	staticcheck ./...

.PHONY: test
test:
	GO111MODULE=on CGO_ENABLED=0 go test -v ./... -coverprofile=coverage.out

.PHONY: install-migrate
install-migrate:
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2

.PHONY: migrate-account-up
migrate-account-up: install-migrate
	./sh/migrate.sh account up

.PHONY: migrate-account-down
migrate-account-down: install-migrate
	./sh/migrate.sh account down

.PHONY: migrate-address-up
migrate-address-up: install-migrate
	./sh/migrate.sh address up

.PHONY: migrate-address-down
migrate-address-down: install-migrate
	./sh/migrate.sh address down

.PHONY: install-sqlboiler
install-sqlboiler:
	go install github.com/volatiletech/sqlboiler/v4@v4.14.2
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.14.2

.PHONY: generate-account-model
generate-account-model: install-sqlboiler
	./sh/sqlboiler.sh account

.PHONY: generate-address-model
generate-address-model: install-sqlboiler
	./sh/sqlboiler.sh address
