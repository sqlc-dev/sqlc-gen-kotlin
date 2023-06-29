package main

import (
	"github.com/sqlc-dev/sqlc-go/codegen"

	kotlin "github.com/sqlc-dev/sqlc-gen-kotlin/internal"
)

func main() {
	codegen.Run(kotlin.Generate)
}
