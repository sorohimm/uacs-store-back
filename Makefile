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
store: UACS_STORE_MAIN := ./cmd/store
store:
	@echo BUILDING $(RULESRVOUT)
	$(V)go build  -ldflags "-s -w -X main.version=${RELEASE} -X main.buildTime=${BUILD_TIME}" -o $(UACS_STORE_OUT) $(UACS_STORE_MAIN)
	@echo DONE

.PHONY: store_linux
store_linux: export GOOS := linux
store_linux: export GOARCH := amd64
store_linux: store

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
