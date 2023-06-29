module github.com/sqlc-dev/sqlc-gen-kotlin

go 1.19

require (
	buf.build/gen/go/sqlc/sqlc/protocolbuffers/go v1.30.0-20230621221448-196413f69ab3.1
	github.com/jinzhu/inflection v1.0.0
	github.com/sqlc-dev/sqlc-go v1.18.1
	github.com/tabbed/sqlc-go v1.18.0
)

require google.golang.org/protobuf v1.30.0 // indirect
