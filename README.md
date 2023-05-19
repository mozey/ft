# mozey/ft

Flexible types<sup>[1]</sup> for un-marshaling any JSON<sup>[2]</sup> value.

This package builds on
- `null.*` types in [guregu/null](https://github.com/guregu/null)
- `sql.Null*` types in [database/sql](https://pkg.go.dev/database/sql)

Numeric types (Int, Float) wrap 64bit base types, e.g. int64 and float64.

For each [basic type](https://go.dev/tour/basics/11), this package implements two flexible types<sup>[3]</sup>, for example
- **required** struct fields may use **ft.String**, it supports flexible types when un-marshaling. 
- **optional** fields may use **ft.NString**, it supports both flexible types and allows NULL. Think of the N-prefix as meaning *"allows null"*, or *"not-required"*
```go
type data struct {
    Required ft.String `json:"required"`
    Optional ft.NString `json:"optional"`
}
```
Above data struct can be used to un-marshal the following JSON
```json
{
    "required": 123,
    "optional": null
}
```
NULL is not considered valid, for example
```go
b := []byte(`{"required": 123,"optional": null}`)
d := data{}
err = json.Unmarshal(b, &d)
fmt.Println(d.Required.Valid)  // true
fmt.Println(d.Required.String) // "123"
fmt.Println(d.Optional.Valid)  // false
fmt.Println(d.Optional.String) // ""
```

Both custom types in this package, allowing NULL or not, embed structs with the same fields. For example
```go
// String
s := ft.StringForm(123)
fmt.Println(string(s.Valid)) // true
fmt.Println(s.String)        // 123
// NString
s := ft.NStringForm(123)
fmt.Println(string(s.Valid)) // true
fmt.Println(s.String)        // 123
```

After un-marshaling, the **Valid** field may be toggled by your custom validation code.

Flexible types can be used with the [templating packages](https://gobyexample.com/text-templates). Either of the types above, **String** or **NString**, could be used in this text template example
```go
t1 = template.Must(t1.Parse(`{{if .Valid }} The value is .String {{end}}`))
t1.Execute(os.Stdout, s) // The value is 123
```


## Tests

See tests for more usage examples
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

[3] This repo does not use [generics](https://go.dev/doc/tutorial/generics), and will compile with older versions of Go

