package core

import (
	"text/template"

	"github.com/tabbed/sqlc-go/sdk"
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

var Foo = template.FuncMap{
	"lowerTitle": LowerTitle,         // sdk.LowerTitle,
	"comment":    DoubleSlashComment, // sdk.DoubleSlashComment,
	"imports":    Imports,
	"offset":     Offset,
}
