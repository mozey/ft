# mozey/ft

Flexible types<sup>[1]</sup> for un-marshalling any JSON<sup>[2]</sup> value.

This package builds on
- `null.*` types in [guregu/null](https://github.com/guregu/null)
- `sql.Null*` types in [database/sql](https://pkg.go.dev/database/sql)

Numeric types (Int, Float) wrap 64bit base types, e.g. int64 and float64.


## Tests

See tests for usage examples
```bash
cd ${PRO_PATH}
git clone https://github.com/mozey/ft.git
cd ${PRO_PATH}/ft
go clean -testcache && go test -v ./...
```


## Reference

[1] Not quite the same concept as [Flexible typing in SQLite](https://www.sqlite.org/flextypegood.html). Note that *"as of SQLite version 3.37.0 (2021-11-27), SQLite supports this development style [Rigid Type Enforcement] using [STRICT tables](https://www.sqlite.org/stricttables.html)"*

[Previously](https://github.com/mozey/fuzzy) this repo used the term *"fuzzy types"*. It was renamed, and the term replaced with *"flexible types"*. To avoid confusion with [fuzzing](https://go.dev/security/fuzz/), a new feature added to the *"standard toolchain beginning in Go 1.18"*

[2] [JSON spec](https://www.json.org/json-en.html)

