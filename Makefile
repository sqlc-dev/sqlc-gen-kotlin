all: sqlc-gen-kotlin sqlc-gen-kotlin.wasm

sqlc-gen-kotlin:
	cd plugin && go build -o ~/bin/sqlc-gen-kotlin ./main.go

sqlc-gen-kotlin.wasm:
	cd plugin && GOOS=wasip1 GOARCH=wasm go build -o sqlc-gen-kotlin.wasm main.go
	openssl sha256 plugin/sqlc-gen-kotlin.wasm

