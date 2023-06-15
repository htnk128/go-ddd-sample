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

.PHONY: fmt
fmt:
	gofmt -l -w .

.PHONY: lint
lint:
	staticcheck ./...

.PHONY: test
test:
	GO111MODULE=on CGO_ENABLED=0 go test -v ./... -coverprofile=coverage.out

.PHONY: build
build: fmt
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) \
		go build -tags timetzdata -ldflags "-s -w" -trimpath -o bin/$(app) cmd/$(app)/server.go

.PHONY: up-all
up-all:
	@docker-compose up -d

.PHONY: up
up:
	@docker-compose up -d go-ddd-sample-$(app)

.PHONY: ps
ps:
	@docker-compose ps

.PHONY: stop
stop:
	@docker-compose stop

.PHONY: rm
rm:
	@docker-compose rm -f -s -v

.PHONY: down
down:
	@docker-compose down --rmi all --volumes

.PHONY: install-migrate
install-migrate:
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2

.PHONY: migrate-up
migrate-up: install-migrate
	./sh/migrate.sh $(app) up

.PHONY: migrate-down
migrate-down: install-migrate
	./sh/migrate.sh $(app) down

.PHONY: install-sqlboiler
install-sqlboiler:
	go install github.com/volatiletech/sqlboiler/v4@v4.14.2
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.14.2

.PHONY: generate-model
generate-model: install-sqlboiler
	./sh/sqlboiler.sh $(app)
