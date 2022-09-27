# suppress output, run `make XXX V=` to be verbose
V := @

RELEASE=$(shell git describe --always --tags)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

OUT_DIR := ./build


default: all

.PHIONY: all
all:tidy
all:store

.PHONY: store
store: UACS_STORE_OUT := $(OUT_DIR)/uacs_store
store: UACS_STORE_MAIN := ./cmd/api
store:
	@echo BUILDING $(RULESRVOUT)
	$(V)go build  -ldflags "-s -w -X main.version=${RELEASE} -X main.buildTime=${BUILD_TIME}" -o $(UACS_STORE_OUT) $(UACS_STORE_MAIN)
	@echo DONE

.PHONY: store_linux
store_linux: export GOOS := linux
store_linux: export GOARCH := amd64
store_linux: store

#### gRPC generation
SRC = "./pkg/api"
DST = "."
.PHONY: gen-api
gen-api: API_SRC:= "./pkg/api"
gen-api: API_DEST:= "."
gen-api:
	protoc -I=. -I$(API_SRC) --go_out=$(API_DEST) --go_opt=paths=source_relative $(API_SRC)/store.proto
	protoc -I=. -I$(API_SRC) --go-grpc_out=$(API_DEST) --go-grpc_opt paths=source_relative $(API_SRC)/store.proto
	protoc -I=. -I$(API_SRC) --grpc-gateway_out=$(API_DEST)  --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative $(API_SRC)/store.proto
	protoc -I=. -I$(API_SRC) --openapiv2_out=$(API_DEST) --openapiv2_opt=logtostderr=true $(API_SRC)/store.proto

#### docker compose
.PHONY: compose-up
compose-up:
	docker compose -f scripts/deploy/docker-compose.yaml up -d


.PHONY: compose-down
compose-down:
	docker compose -f scripts/deploy/docker-compose.yaml down


## Настройка GOPRIVATE https://gist.github.com/MicahParks/1ba2b19c39d1e5fccc3e892837b10e21
GOPRIVATE="github.com/*"
.PHONY: tidy
tidy:
	$(V)GOPRIVATE=$(GOPRIVATE) go mod tidy -v

##### go-migrate
MIGRATE_PATH := scripts/migrate
.PHONY: migrate-up
migrate-up:
	migrate -database 'postgresql://pg:test@localhost:5432/icd11?sslmode=disable' -path $(MIGRATE_PATH) -verbose up

.PHONY: migrate-down
migrate-down:
	migrate -database 'postgresql://pg:test@localhost:5432/icd11?sslmode=disable' -path $(MIGRATE_PATH) -verbose down

NAME:=@
.PHONY: migrate-create
migrate-create:
	migrate create -ext sql -dir $(MIGRATE_PATH) -seq $(V)