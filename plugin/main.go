package main

import (
	"github.com/tabbed/sqlc-go/codegen"

	kotlin "github.com/tabbed/sqlc-gen-kotlin/internal"
)

func main() {
	codegen.Run(kotlin.Generate)
}
