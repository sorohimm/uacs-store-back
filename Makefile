# suppress output, run `make XXX V=` to be verbose
V := @

RELEASE=$(shell git describe --always --tags)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

OUT_DIR := ./build


default: all

.PHIONY: all
all:tidy
all:product
all:order
all:auth

.PHONY: linux
linux: export GOOS := linux
linux: export GOARCH := amd64
linux:all

#### api service build
.PHONY: product
product: PRODUCT_OUT := $(OUT_DIR)/product
product: PRODUCT_MAIN := ./cmd/product
product:
	@echo BUILDING $(PRODUCT_OUT)
	$(V)go build  -ldflags "-s -w -X main.version=${RELEASE} -X main.buildTime=${BUILD_TIME}" -o $(PRODUCT_OUT) $(PRODUCT_MAIN)
	@echo DONE

#### store API service build for linux
.PHONY: product-linux
product-linux: export GOOS := linux
product-linux: export GOARCH := amd64
product-linux: product

#### orders service build
.PHONY: order
order: ORDER_OUT := $(OUT_DIR)/order
order: ORDER_MAIN := ./cmd/order
order:
	@echo BUILDING $(ORDER_OUT)
	$(V)go build  -ldflags "-s -w -X main.version=${RELEASE} -X main.buildTime=${BUILD_TIME}" -o $(ORDER_OUT) $(ORDER_MAIN)
	@echo DONE

#### auth service build for linux
.PHONY: order-linux
order-linux: export GOOS := linux
order-linux: export GOARCH := amd64
order-linux: order

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

#build:
#	GOARCH=wasm GOOS=js go build -o ./build/web/app.wasm ./cmd/frontend
#	go build -o ./build/frontend ./cmd/frontend

#### gRPC all proto generation
.PHONY: gen
gen: gen-product
gen: gen-order
gen: gen-cart
gen: gen-auth



#### gRPC product api generation
SRC = "./pkg/api"
DST = "."
.PHONY: gen-product
gen-product: PRODUCT_SRC:= "./pkg/api"
gen-product: PRODUCT_DEST:= "."
gen-product:
	protoc -I=. -I$(PRODUCT_SRC) --go_out=$(PRODUCT_DEST) --go_opt=paths=source_relative $(PRODUCT_SRC)/store.proto
	protoc -I=. -I$(PRODUCT_SRC) --go-grpc_out=$(PRODUCT_DEST) --go-grpc_opt paths=source_relative $(PRODUCT_SRC)/store.proto
	protoc -I=. -I$(PRODUCT_SRC) --grpc-gateway_out=$(PRODUCT_DEST)  --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative $(PRODUCT_SRC)/store.proto
	protoc -I=. -I$(PRODUCT_SRC) --openapiv2_out=$(PRODUCT_DEST) --openapiv2_opt=logtostderr=true $(PRODUCT_SRC)/store.proto


#### gRPC order api generation
SRC = "./pkg/api"
DST = "."
.PHONY: gen-order
gen-order: ORDER_SRC:= "./pkg/api"
gen-order: ORDER_DEST:= "."
gen-order:
	protoc -I=. -I$(ORDER_SRC) --go_out=$(ORDER_DEST) --go_opt=paths=source_relative $(ORDER_SRC)/order.proto
	protoc -I=. -I$(ORDER_SRC) --go-grpc_out=$(ORDER_DEST) --go-grpc_opt paths=source_relative $(ORDER_SRC)/order.proto
	protoc -I=. -I$(ORDER_SRC) --grpc-gateway_out=$(ORDER_DEST)  --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative $(ORDER_SRC)/order.proto
	protoc -I=. -I$(ORDER_SRC) --openapiv2_out=$(ORDER_DEST) --openapiv2_opt=logtostderr=true $(ORDER_SRC)/order.proto


#### gRPC cart api generation
SRC = "./pkg/api"
DST = "."
.PHONY: gen-cart
gen-cart: CART_SRC:= "./pkg/api"
gen-cart: CART_DEST:= "."
gen-cart:
	protoc -I=. -I$(CART_SRC) --go_out=$(CART_DEST) --go_opt=paths=source_relative $(CART_SRC)/cart.proto
	protoc -I=. -I$(CART_SRC) --go-grpc_out=$(CART_DEST) --go-grpc_opt paths=source_relative $(CART_SRC)/cart.proto
	protoc -I=. -I$(CART_SRC) --grpc-gateway_out=$(CART_DEST)  --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative $(CART_SRC)/cart.proto
	protoc -I=. -I$(CART_SRC) --openapiv2_out=$(CART_DEST) --openapiv2_opt=logtostderr=true $(CART_SRC)/cart.proto


#### gRPC auth api generation
SRC = "./pkg/api"
DST = "."
.PHONY: gen-auth
gen-auth: AUTH_SRC:= "./pkg/api"
gen-auth: AUTH_DEST:= "."
gen-auth:
	protoc -I=. -I$(AUTH_SRC) --go_out=$(AUTH_DEST) --go_opt=paths=source_relative $(AUTH_SRC)/auth.proto
	protoc -I=. -I$(AUTH_SRC) --go-grpc_out=$(AUTH_DEST) --go-grpc_opt paths=source_relative $(AUTH_SRC)/auth.proto
	protoc -I=. -I$(AUTH_SRC) --grpc-gateway_out=$(AUTH_DEST)  --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative $(AUTH_SRC)/auth.proto
	protoc -I=. -I$(AUTH_SRC) --openapiv2_out=$(AUTH_DEST) --openapiv2_opt=logtostderr=true $(AUTH_SRC)/auth.proto

REPO := "sorohimm"
AUTH := "uacs-auth"
STORE:= "uacs-product"
.PHONY: images
images: linux
images:
	docker image build -t ${REPO}/${AUTH}:${RELEASE} -t ${REPO}/${AUTH}:latest -f scripts/auth-server.Dockerfile .
	docker image build -t ${REPO}/${STORE}:${RELEASE} -t ${REPO}/${STORE}:latest -f scripts/product-server.Dockerfile .

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
