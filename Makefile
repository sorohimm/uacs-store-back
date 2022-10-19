# suppress output, run `make XXX V=` to be verbose
V := @

RELEASE=$(shell git describe --always --tags)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

OUT_DIR := ./build


default: all

.PHIONY: all
all:tidy
all:store
all:auth

#### store API service build
.PHONY: store
store: UACS_STORE_OUT := $(OUT_DIR)/uacs-store
store: UACS_STORE_MAIN := ./cmd/api
store:
	@echo BUILDING $(UACS_STORE_OUT)
	$(V)go build  -ldflags "-s -w -X main.version=${RELEASE} -X main.buildTime=${BUILD_TIME}" -o $(UACS_STORE_OUT) $(UACS_STORE_MAIN)
	@echo DONE

#### store API service build for linux
.PHONY: store-linux
store-linux: export GOOS := linux
store-linux: export GOARCH := amd64
store-linux: store

#### auth service build
.PHONY: auth
auth: UACS_AUTH_OUT := $(OUT_DIR)/uacs-auth
auth: UACS_AUTH_MAIN := ./cmd/auth
auth:
	@echo BUILDING $(UACS_AUTH_OUT)
	$(V)go build  -ldflags "-s -w -X main.version=${RELEASE} -X main.buildTime=${BUILD_TIME}" -o $(UACS_AUTH_OUT) $(UACS_AUTH_MAIN)
	@echo DONE

#### auth service build for linux
.PHONY: auth-linux
auth-linux: export GOOS := linux
auth-linux: export GOARCH := amd64
auth-linux: auth

#.PHONY: wasm
#wasm: WASM_OUT := $(OUT_DIR)/frontend
#wasm: WASM_MAIN := ./cmd/frontend
#wasm: export GOARCH := wasm
#wasm: export GOOS := js go build -o ./build/web/app.wasm ./cmd/frontend
#wasm:
#	@echo BUILDING $(WASM_OUT)
#	$(V)go build -ldflags "-s -w -X main.version=${RELEASE} -X main.buildTime=${BUILD_TIME}" -o $(WASM_OUT) $(WASM_MAIN)
#	@echo DONE

build:
	GOARCH=wasm GOOS=js go build -o ./build/web/app.wasm ./cmd/frontend
	go build -o ./build/frontend ./cmd/frontend

#### gRPC store api generation
SRC = "./pkg/api"
DST = "."
.PHONY: gen-api
gen-api: API_SRC:= "./pkg/api"
gen-api: API_DEST:= "."
gen-api:
	protoc -I=. -I$(API_SRC) --go_out=$(API_DEST) --go_opt=paths=source_relative $(API_SRC)/store.proto
	protoc -I=. -I$(API_SRC) --go-grpc_out=$(API_DEST) --go-grpc_opt paths=source_relative $(API_SRC)/store.proto
	protoc -I=. -I$(API_SRC) --grpc-gateway_out=$(API_DEST)  --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative $(API_SRC)/store.proto
	protoc -I=. -I$(API_SRC) --openapiv2_out=$(API_DEST) --openapiv2_opt=logtostderr=true $(API_SRC)/store.proto#### gRPC generation

#### gRPC auth api generation
SRC = "./pkg/auth"
DST = "."
.PHONY: gen-auth
gen-auth: AUTH_SRC:= "./pkg/auth"
gen-auth: AUTH_DEST:= "."
gen-auth:
	protoc -I=. -I$(AUTH_SRC) --go_out=$(AUTH_DEST) --go_opt=paths=source_relative $(AUTH_SRC)/auth.proto
	protoc -I=. -I$(AUTH_SRC) --go-grpc_out=$(AUTH_DEST) --go-grpc_opt paths=source_relative $(AUTH_SRC)/auth.proto
	protoc -I=. -I$(AUTH_SRC) --grpc-gateway_out=$(AUTH_DEST)  --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative $(AUTH_SRC)/auth.proto
	protoc -I=. -I$(AUTH_SRC) --openapiv2_out=$(AUTH_DEST) --openapiv2_opt=logtostderr=true $(AUTH_SRC)/auth.proto

#### docker compose up
.PHONY: compose-up
compose-up:
	docker compose -f scripts/deploy/local/docker-compose.yml up -d

#### docker compose down
.PHONY: compose-down
compose-down:
	docker compose -f scripts/deploy/local/docker-compose.yml down

.PHONY: fmt
fmt:
	$(V)gofumpt -l -w .

#### GOPRIVATE setup https://gist.github.com/MicahParks/1ba2b19c39d1e5fccc3e892837b10e21
GOPRIVATE="github.com/*"
.PHONY: tidy
tidy:
	$(V)GOPRIVATE=$(GOPRIVATE) go mod tidy -v

.PHONY: lint
lint:
	$(V)golangci-lint run --config scripts/.golangci.yml

.PHONY: test
test: GO_TEST_FLAGS += -race
call testtest: GO_TEST_FLAGS += -count=1
test:
	$(V)go test $(GO_TEST_FLAGS) --tags=$(GO_TEST_TAGS) ./...


.PHONY: fulltest
fulltest: GO_TEST_TAGS += integration
fulltest: test

##### go-migrate
MIGRATE_PATH := scripts/migrate
.PHONY: migrate-up
migrate-up:
	migrate -database 'postgresql://pg:test@localhost:5432/uacs?sslmode=disable' -path $(MIGRATE_PATH) -verbose up

.PHONY: migrate-down
migrate-down:
	migrate -database 'postgresql://pg:test@localhost:5432/uacs?sslmode=disable' -path $(MIGRATE_PATH) -verbose down

NAME:=@
.PHONY: migrate-create
migrate-create:
	migrate create -ext sql -dir $(MIGRATE_PATH) -seq $(V)