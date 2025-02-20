# mozey/ft

Flexible types<sup>[[1](https://github.com/mozey/ft?tab=readme-ov-file#1)]</sup> for un-marshaling any JSON<sup>[[2](https://github.com/mozey/ft?tab=readme-ov-file#2)]</sup> value.


## Why?

This lib is useful when API clients send inconsistent JSON.

For example, a webhook endpoint might want to accept POST requests with body:

- This is what the API expects
```json
{
    "sku": "123",
    "price": 9.99
}
```

- This will also un-marshal, sku is cast to a string, and price to a float64. Errors if price is not a valid number
```json
{
    "sku": 123,
    "price": "9.99"
}
```

Validation can then be done after un-marshaling


## Usage

This package builds on
- `null.*` types in [guregu/null](https://github.com/guregu/null)
- `sql.Null*` types in [database/sql](https://pkg.go.dev/database/sql)

Numeric types (Int, Float) wrap 64bit base types, e.g. int64 and float64.

For each [basic type](https://go.dev/tour/basics/11), this package implements two flexible types<sup>[[3](https://github.com/mozey/ft?tab=readme-ov-file#3)]</sup>, for example
- **required** struct fields may use **ft.String**, it supports flexible types when un-marshaling. 
- **optional** fields may use **ft.NString**, it supports both flexible types and allows NULL. Think of the N-prefix as meaning *"allows null"*, or *"not-required"*
```go
type Data struct {
    Required ft.String `json:"required"`
    Optional ft.NString `json:"optional"`
}
```

The data type above can be used to un-marshal the following JSON
```json
{
    "required": 123,
    "optional": null
}
```

NULL is not considered valid, for example
```go
b := []byte(`{"required": 123,"optional": null}`)
d := Data{}
err := json.Unmarshal(b, &d)
// ft.String doesn't have the Valid property
fmt.Println(d.Required.String) // "123"
fmt.Println(d.Optional.Valid)  // false
fmt.Println(d.Optional.String) // ""
```

Valid is set if the value is defined (not null), don't use it for e.g. form validation. Rather create a method on the data type
```go
// Valid validates all the required properties in the Data type
func (d *Data) Valid() bool {
    return d.Required.String == "123"
}
```

Flexible types can be used with the [templating packages](https://gobyexample.com/text-templates). The data type above could be used like this
```go
t1 := template.Must(template.New("t1").Parse(
    `{{if not .Optional.Valid }}Required = {{.Required.String}} {{end}}`))
t1.Execute(os.Stdout, d) // Required = 123
```


## Tests

See tests for more usage examples
```bash
git clone https://github.com/mozey/ft.git && cd ft
go clean -testcache && go test -v ./...
```


## References

### [1] 

Not quite the same concept as [Flexible typing in SQLite](https://www.sqlite.org/flextypegood.html). Note that *"as of SQLite version 3.37.0 (2021-11-27), SQLite supports this development style [Rigid Type Enforcement] using [STRICT tables](https://www.sqlite.org/stricttables.html)"*

[Previously](https://github.com/mozey/fuzzy) this repo used the term *"fuzzy types"*. It was renamed, and the term replaced with *"flexible types"*. To avoid confusion with [fuzzing](https://go.dev/security/fuzz/), a new feature added to the *"standard toolchain beginning in Go 1.18"*

### [2] 

[JSON spec](https://www.json.org/json-en.html)

### [3] 

This repo does not use [generics](https://go.dev/doc/tutorial/generics), and will compile with older versions of Go

