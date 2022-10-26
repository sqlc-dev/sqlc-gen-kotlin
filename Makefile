all: sqlc-gen-kotlin sqlc-gen-kotlin.wasm

sqlc-gen-kotlin:
	cd plugin && go build -o ~/bin/sqlc-gen-kotlin ./main.go

sqlc-gen-kotlin.wasm:
	cd plugin && tinygo build -o sqlc-gen-kotlin.wasm -gc=leaking -scheduler=none -wasm-abi=generic -target=wasi main.go
	openssl sha256 plugin/sqlc-gen-kotlin.wasm

