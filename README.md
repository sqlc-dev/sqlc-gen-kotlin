## Usage

```yaml
version: '2'
plugins:
- name: kt
  wasm:
    url: https://downloads.sqlc.dev/plugin/sqlc-gen-kotlin_1.2.0.wasm
    sha256: 22b437ecaea66417bbd3b958339d9868ba89368ce542c936c37305acf373104b
sql:
- schema: src/main/resources/authors/postgresql/schema.sql
  queries: src/main/resources/authors/postgresql/query.sql
  engine: postgresql
  codegen:
  - out: src/main/kotlin/com/example/authors/postgresql
    plugin: kt
    options:
      package: com.example.authors.postgresql
```
