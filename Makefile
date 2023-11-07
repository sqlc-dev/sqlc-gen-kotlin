all: sqlc-gen-kotlin sqlc-gen-kotlin.wasm

sqlc-gen-kotlin: bin
	cd plugin && go build -o ../bin/sqlc-gen-kotlin ./main.go

sqlc-gen-kotlin.wasm: bin
	cd plugin && GOOS=wasip1 GOARCH=wasm go build -o ../bin/sqlc-gen-kotlin.wasm main.go

bin:
	mkdir -p bin

