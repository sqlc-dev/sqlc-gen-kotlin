package kotlin

type Config struct {
	Package                     string   `json:"package"`
	EmitExactTableNames         bool     `json:"emit_exact_table_names"`
	InflectionExcludeTableNames []string `json:"inflection_exclude_table_names"`
}
