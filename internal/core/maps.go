package core

import (
	"github.com/sqlc-dev/plugin-sdk-go/sdk"
)

func DoubleSlashComment(f string) string {
	return sdk.DoubleSlashComment(f)
}

func LowerTitle(f string) string {
	return sdk.LowerTitle(f)
}

var DefaultImporter *Importer

func Imports(filename string) [][]string {
	if DefaultImporter == nil {
		return nil
	}
	return DefaultImporter.Imports(filename)
}
