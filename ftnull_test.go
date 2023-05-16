package ft_test

import (
	"encoding/json"
	"testing"
	"unicode"

	"github.com/guregu/null"
	"github.com/matryer/is"
	"github.com/mozey/ft"
)

func TestUnmarshalNString(t *testing.T) {
	is := is.New(t)

	type Data struct {
		String ft.NString `json:"string"`
	}
	d := Data{}

	// null
	b := []byte(`{"string": null}`)
	err := json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(false, d.String.Valid) // String must not be valid
	is.Equal("", d.String.String)   // Value must match

	b = []byte(`{}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(false, d.String.Valid) // String must not be valid
	is.Equal("", d.String.String)   // Value must match

	// string
	b = []byte(`{"string": "123"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.String.Valid)   // String must be valid
	is.Equal("123", d.String.String) // Value must match

	// int
	b = []byte(`{"string": 123}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.String.Valid)   // String must be valid
	is.Equal("123", d.String.String) // Value must match

	b = []byte(`{"string": 0}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.String.Valid) // String must be valid
	is.Equal("0", d.String.String) // Value must match

	b = []byte(`{"string": -123}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.String.Valid)    // String must be valid
	is.Equal("-123", d.String.String) // Value must match

	// float
	b = []byte(`{"string": 123.456}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.String.Valid)       // String must be valid
	is.Equal("123.456", d.String.String) // Value must match

	b = []byte(`{"string": -123.456}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.String.Valid)        // String must be valid
	is.Equal("-123.456", d.String.String) // Value must match

	// bool
	b = []byte(`{"string": true}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.String.Valid)    // String must be valid
	is.Equal("true", d.String.String) // Value must match

	b = []byte(`{"string": false}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.String.Valid)     // String must be valid
	is.Equal("false", d.String.String) // Value must match
}

func TestMarshalNString(t *testing.T) {
	is := is.New(t)

	type Data struct {
		String ft.NString `json:"string"`
	}
	d := Data{}

	// Unicode escape sequence
	d.String = ft.NStringFrom("[bla] bla \u0026 bla")
	b, err := json.Marshal(d)
	is.NoErr(err)
	is.Equal(`{"string":"[bla] bla \u0026 bla"}`, string(b))

	d.String = ft.NStringFrom("foo\u0002bar")
	b, err = json.Marshal(d)
	is.NoErr(err)
	is.Equal(`{"string":"foobar"}`, string(b))

	// Slashes
	d.String = ft.NStringFrom("bla bla\\bla bla")
	b, err = json.Marshal(d)
	is.NoErr(err)
	is.Equal(`{"string":"bla bla\\bla bla"}`, string(b))

	// JSON inside string
	d.String = ft.NStringFrom(`{"O":1}`)
	b, err = json.Marshal(d)
	is.NoErr(err)
	is.Equal(`{"string":"{\"O\":1}"}`, string(b))

	// Escaped JSON inside string
	d.String = ft.NStringFrom(`{\"O\":1}`)
	b, err = json.Marshal(d)
	is.NoErr(err)
	is.Equal(`{"string":"{\\\"O\\\":1}"}`, string(b))

	// Escape characters
	d.String = ft.NStringFrom("First line\nSecond line")
	b, err = json.Marshal(d)
	is.NoErr(err)
	is.Equal(`{"string":"First line\nSecond line"}`, string(b))

	// HTML
	d.String = ft.NStringFrom("<h3>Hello</h3>")
	b, err = json.Marshal(d)
	is.NoErr(err)
	is.Equal(`{"string":"\u003ch3\u003eHello\u003c/h3\u003e"}`, string(b))

	// Space characters
	d.String = ft.NStringFrom("Ascii: , Non-breaking:\u00A0, Tab:\t")
	b, err = json.Marshal(d)
	is.NoErr(err)
	// Note non-numeric chars in the unicode escape sequence is lowercased
	is.Equal(`{"string":"Ascii: , Non-breaking:\u00a0, Tab:\t"}`, string(b))
	// Tab character not preserved, it's not considered graphic?
	// See comments in ft.Clean and unicode.IsPrint
	is.Equal(false, unicode.IsGraphic([]rune("\t")[0]))
	is.Equal(false, unicode.IsPrint([]rune("\t")[0]))
	is.Equal("\t", ft.Clean("\t"))
}

func TestUnmarshalNInt(t *testing.T) {
	is := is.New(t)

	type Data struct {
		Int ft.NInt `json:"int"`
	}
	d := Data{}

	// null
	b := []byte(`{"int": null}`)
	err := json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(false, d.Int.Valid)    // Int must not be valid
	is.Equal(int64(0), d.Int.Int64) // Value must match

	b = []byte(`{}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(false, d.Int.Valid)    // Int must not be valid
	is.Equal(int64(0), d.Int.Int64) // Value must match

	// string
	b = []byte(`{"int": "123"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Int.Valid)       // Int must be valid
	is.Equal(int64(123), d.Int.Int64) // Value must match

	b = []byte(`{"int": "-123"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Int.Valid)        // Int must be valid
	is.Equal(int64(-123), d.Int.Int64) // Value must match

	b = []byte(`{"int": "abc"}`)
	err = json.Unmarshal(b, &d)
	is.Equal("strconv.ParseInt: parsing \"abc\": invalid syntax", err.Error())

	// int
	b = []byte(`{"int": -123}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Int.Valid)        // Int must be valid
	is.Equal(int64(-123), d.Int.Int64) // Value must match

	// float
	b = []byte(`{"int": -123.456}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Int.Valid)        // Int must be valid
	is.Equal(int64(-123), d.Int.Int64) // Value must match

	// bool
	b = []byte(`{"int": true}`)
	err = json.Unmarshal(b, &d)
	is.Equal("value is a bool", err.Error())

	// empty string
	b = []byte(`{"int": ""}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(false, d.Int.Valid)    // Int must not be valid
	is.Equal(int64(0), d.Int.Int64) // Value must match

	// space
	b = []byte(`{"int": "  "}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(false, d.Int.Valid)    // Int must not be valid
	is.Equal(int64(0), d.Int.Int64) // Value must match
}

func TestUnmarshalNFloat(t *testing.T) {
	is := is.New(t)

	type Data struct {
		Float ft.NFloat `json:"float"`
	}
	d := Data{}

	// null
	b := []byte(`{"float": null}`)
	err := json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(false, d.Float.Valid)        // Must not be valid
	is.Equal(float64(0), d.Float.Float64) // Value must match

	b = []byte(`{}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(false, d.Float.Valid)        // Must not be valid
	is.Equal(float64(0), d.Float.Float64) // Value must match

	// string
	b = []byte(`{"float": "1.618"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Float.Valid)             // Must be valid
	is.Equal(float64(1.618), d.Float.Float64) // Value must match

	b = []byte(`{"float": "-1.618"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Float.Valid)              // Must be valid
	is.Equal(float64(-1.618), d.Float.Float64) // Value must match

	b = []byte(`{"float": "abc"}`)
	err = json.Unmarshal(b, &d)
	is.Equal("strconv.ParseFloat: parsing \"abc\": invalid syntax", err.Error())

	// int
	b = []byte(`{"float": -1}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Float.Valid)          // Must be valid
	is.Equal(float64(-1), d.Float.Float64) // Value must match

	// float
	b = []byte(`{"float": -1.618}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Float.Valid)              // Int must be valid
	is.Equal(float64(-1.618), d.Float.Float64) // Value must match

	// bool
	b = []byte(`{"float": true}`)
	err = json.Unmarshal(b, &d)
	is.Equal("value is a bool", err.Error())
}

func TestUnmarshalNBool(t *testing.T) {
	is := is.New(t)

	type Data struct {
		Bool ft.NBool `json:"bool"`
	}
	d := Data{}

	// null
	b := []byte(`{"bool": null}`)
	err := json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(false, d.Bool.Valid) // Bool must not be valid
	is.Equal(false, d.Bool.Bool)  // Value must match

	b = []byte(`{}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(false, d.Bool.Valid) // Bool must not be valid
	is.Equal(false, d.Bool.Bool)  // Value must match

	// string
	b = []byte(`{"bool": "false"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Valid) // Bool must be valid
	is.Equal(false, d.Bool.Bool) // Value must match

	b = []byte(`{"bool": "0"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Valid) // Bool must be valid
	is.Equal(false, d.Bool.Bool) // Value must match

	b = []byte(`{"bool": ""}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Valid) // Bool must be valid
	is.Equal(false, d.Bool.Bool) // Value must match

	b = []byte(`{"bool": "true"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Valid) // Bool must be valid
	is.Equal(true, d.Bool.Bool)  // Value must match

	b = []byte(`{"bool": "1"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Valid) // Bool must be valid
	is.Equal(true, d.Bool.Bool)  // Value must match

	b = []byte(`{"bool": "abc"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Valid) // Bool must be valid
	is.Equal(true, d.Bool.Bool)  // Value must match

	// int
	b = []byte(`{"bool": 0}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Valid) // Bool must be valid
	is.Equal(false, d.Bool.Bool) // Value must match

	b = []byte(`{"bool": 1}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Valid) // Bool must be valid
	is.Equal(true, d.Bool.Bool)  // Value must match

	b = []byte(`{"bool": -1}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Valid) // Bool must be valid
	is.Equal(true, d.Bool.Bool)  // Value must match

	// float
	b = []byte(`{"bool": 1.23}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Valid) // Bool must be valid
	is.Equal(true, d.Bool.Bool)  // Value must match

	b = []byte(`{"bool": -1.23}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Valid) // Bool must be valid
	is.Equal(true, d.Bool.Bool)  // Value must match

	// bool
	b = []byte(`{"bool": false}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Valid) // Bool must be valid
	is.Equal(false, d.Bool.Bool) // Value must match

	b = []byte(`{"bool": true}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Valid) // Bool must be valid
	is.Equal(true, d.Bool.Bool)  // Value must match
}

func TestNMarshalToJSON(t *testing.T) {
	is := is.New(t)

	type Data struct {
		String ft.NString `json:"string"`
		Int    ft.NInt    `json:"int"`
		Bool   ft.NBool   `json:"bool"`
		Float  ft.NFloat  `json:"float"`
	}

	// Valid is empty (false) by default,
	// invalid values marshal to null
	d := Data{}
	d.String = ft.NString{}
	d.Int = ft.NInt{}
	d.Bool = ft.NBool{}
	d.Float = ft.NFloat{}
	b, err := json.Marshal(d)
	is.NoErr(err)
	is.Equal(`{"string":null,"int":null,"bool":null,"float":null}`, string(b))

	d = Data{}
	// Alternative instantiation, but the effect is the same as above
	d.String = ft.NString(null.String{})
	d.Int = ft.NInt(null.Int{})
	d.Bool = ft.NBool(null.Bool{})
	d.Float = ft.NFloat(null.Float{})
	b, err = json.Marshal(d)
	is.NoErr(err)
	is.Equal(`{"string":null,"int":null,"bool":null,"float":null}`, string(b))

	// Valid values marshal to the corresponding JSON value
	d = Data{}
	d.String = ft.NStringFrom("foo")
	d.Int = ft.NIntFrom(123)
	d.Bool = ft.NBoolFrom(true)
	d.Float = ft.NFloatFrom(1.618)
	b, err = json.Marshal(d)
	is.NoErr(err)
	is.Equal(`{"string":"foo","int":123,"bool":true,"float":1.618}`, string(b))

	// Empty values are valid
	d = Data{}
	d.String = ft.NStringFrom("")
	d.Int = ft.NIntFrom(0)
	d.Bool = ft.NBoolFrom(false)
	d.Float = ft.NFloatFrom(0)
	b, err = json.Marshal(d)
	is.NoErr(err)
	is.Equal(`{"string":"","int":0,"bool":false,"float":0}`, string(b))

	// Missing properties marshal to null.
	// This is different from un-marshal
	// where missing or null values are set to empty
	d = Data{}
	b, err = json.Marshal(d)
	is.NoErr(err)
	is.Equal(`{"string":null,"int":null,"bool":null,"float":null}`, string(b))

	// Missing properties are marshaled as null.
	// Marshal does not output the same as input to un-marshal
	err = json.Unmarshal([]byte("{}"), &d)
	is.NoErr(err)
	b, err = json.Marshal(d)
	is.NoErr(err)
	is.Equal(`{"string":null,"int":null,"bool":null,"float":null}`, string(b))

	// For base types the omitempty tag works as expected.
	// However, note the following as per README in guregu/null
	// "json's ",omitempty" struct tag does not work correctly right now.
	// It will never omit a null or empty String..."
	// https://github.com/guregu/null#bugs
	// https://github.com/golang/go/issues/11939
	type Data2 struct {
		// Note omitempty is not "omitnull",
		// both empty and null will be omitted
		StringEmptyTag string     `json:"string_empty,omitempty"`
		StringNullTag  string     `json:"string_null,omitempty"`
		String         ft.NString `json:"string,omitempty"`
		Int            ft.NInt    `json:"int,omitempty"`
		Bool           ft.NBool   `json:"bool,omitempty"`
		Float          ft.NFloat  `json:"float,omitempty"`
	}
	d2 := Data2{}
	err = json.Unmarshal([]byte("{}"), &d2)
	is.NoErr(err)
	b, err = json.Marshal(d2)
	is.NoErr(err)
	is.Equal(`{"string":null,"int":null,"bool":null,"float":null}`, string(b))
}
