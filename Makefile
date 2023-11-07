.PHONY: build test

build:
	go build ./...

test:
	go test ./...

all: bin/sqlc-gen-kotlin bin/sqlc-gen-kotlin.wasm

bin/sqlc-gen-kotlin: bin go.mod go.sum $(wildcard **/*.go)
	cd plugin && go build -o ../bin/sqlc-gen-kotlin ./main.go

bin/sqlc-gen-kotlin.wasm: bin/sqlc-gen-kotlin
	cd plugin && GOOS=wasip1 GOARCH=wasm go build -o ../bin/sqlc-gen-kotlin.wasm main.go

bin:
	mkdir -p bin

