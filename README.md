## Usage

```yaml
version: '2'
plugins:
- name: kt
  wasm:
    url: https://downloads.sqlc.dev/plugin/sqlc-gen-kotlin_1.1.0.wasm
    sha256: 57890144f4effed4fe71855418b87ad26d53dc5ed2030a66ad77e5e3a93d77fb
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
