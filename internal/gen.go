package kotlin

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"strings"

	easyjson "github.com/mailru/easyjson"
	plugin "github.com/tabbed/sqlc-go/codegen"

	"github.com/tabbed/sqlc-gen-kotlin/internal/core"
	"github.com/tabbed/sqlc-gen-kotlin/internal/tmpl"
)

func Generate(ctx context.Context, req *plugin.Request) (*plugin.Response, error) {
	var conf core.Config
	if len(req.PluginOptions) > 0 {
		if err := easyjson.Unmarshal(req.PluginOptions, &conf); err != nil {
			return nil, err
		}
	}

	enums := core.BuildEnums(req)
	structs := core.BuildDataClasses(conf, req)
	queries, err := core.BuildQueries(req, structs)
	if err != nil {
		return nil, err
	}

	i := &core.Importer{
		Settings:    req.Settings,
		Enums:       enums,
		DataClasses: structs,
		Queries:     queries,
	}

	core.DefaultImporter = i

	tctx := core.KtTmplCtx{
		Settings:    req.Settings,
		Q:           `"""`,
		Package:     conf.Package,
		Queries:     queries,
		Enums:       enums,
		DataClasses: structs,
		SqlcVersion: req.SqlcVersion,
	}

	output := map[string]string{}

	execute := func(name string, f func(io.Writer, core.KtTmplCtx) error) error {
		var b bytes.Buffer
		w := bufio.NewWriter(&b)
		tctx.SourceName = name
		err := f(w, tctx)
		w.Flush()
		if err != nil {
			return err
		}
		if !strings.HasSuffix(name, ".kt") {
			name += ".kt"
		}
		output[name] = core.KtFormat(b.String())
		return nil
	}

	if err := execute("Models.kt", tmpl.KtModels); err != nil {
		return nil, err
	}
	if err := execute("Queries.kt", tmpl.KtIface); err != nil {
		return nil, err
	}
	if err := execute("QueriesImpl.kt", tmpl.KtSQL); err != nil {
		return nil, err
	}

	resp := plugin.CodeGenResponse{}

	for filename, code := range output {
		resp.Files = append(resp.Files, &plugin.File{
			Name:     filename,
			Contents: []byte(code),
		})
	}

	return &resp, nil
}
