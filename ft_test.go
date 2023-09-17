package ft_test

import (
	"bytes"
	"encoding/json"
	"testing"
	"text/template"

	"github.com/matryer/is"
	"github.com/mozey/ft"
)

func TestUnmarshalString(t *testing.T) {
	is := is.New(t)

	type Data struct {
		String ft.String `json:"string"`
	}
	d := Data{}

	// null
	b := []byte(`{"string": null}`)
	err := json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal("", d.String.String) // Value must match

	b = []byte(`{}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal("", d.String.String) // Value must match

	// string
	b = []byte(`{"string": "123"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal("123", d.String.String) // Value must match

	// int
	b = []byte(`{"string": 123}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal("123", d.String.String) // Value must match

	b = []byte(`{"string": 0}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal("0", d.String.String) // Value must match

	b = []byte(`{"string": -123}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal("-123", d.String.String) // Value must match

	// float
	b = []byte(`{"string": 123.456}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal("123.456", d.String.String) // Value must match

	b = []byte(`{"string": -123.456}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal("-123.456", d.String.String) // Value must match

	// bool
	b = []byte(`{"string": true}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal("true", d.String.String) // Value must match

	b = []byte(`{"string": false}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal("false", d.String.String) // Value must match
}

func TestUnmarshalInt(t *testing.T) {
	is := is.New(t)

	type Data struct {
		Int ft.Int `json:"int"`
	}
	d := Data{}

	// null
	b := []byte(`{"int": null}`)
	err := json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(int64(0), d.Int.Int64) // Value must match

	b = []byte(`{}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(int64(0), d.Int.Int64) // Value must match

	// string
	b = []byte(`{"int": "123"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(int64(123), d.Int.Int64) // Value must match

	b = []byte(`{"int": "-123"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(int64(-123), d.Int.Int64) // Value must match

	b = []byte(`{"int": "abc"}`)
	err = json.Unmarshal(b, &d)
	is.Equal("strconv.ParseInt: parsing \"abc\": invalid syntax", err.Error())

	// int
	b = []byte(`{"int": -123}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(int64(-123), d.Int.Int64) // Value must match

	// float
	b = []byte(`{"int": -123.456}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(int64(-123), d.Int.Int64) // Value must match

	// bool
	b = []byte(`{"int": true}`)
	err = json.Unmarshal(b, &d)
	is.Equal("value is a bool", err.Error())
}

func TestUnmarshalFloat(t *testing.T) {
	is := is.New(t)

	type Data struct {
		Float ft.Float `json:"float"`
	}
	d := Data{}

	// null
	b := []byte(`{"float": null}`)
	err := json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(float64(0), d.Float.Float64) // Value must match

	b = []byte(`{}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(float64(0), d.Float.Float64) // Value must match

	// string
	b = []byte(`{"float": "1.618"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(float64(1.618), d.Float.Float64) // Value must match

	b = []byte(`{"float": "-1.618"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(float64(-1.618), d.Float.Float64) // Value must match

	b = []byte(`{"float": "abc"}`)
	err = json.Unmarshal(b, &d)
	is.Equal("strconv.ParseFloat: parsing \"abc\": invalid syntax", err.Error())

	// int
	b = []byte(`{"float": -1}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(float64(-1), d.Float.Float64) // Value must match

	// float
	b = []byte(`{"float": -1.618}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(float64(-1.618), d.Float.Float64) // Value must match

	// bool
	b = []byte(`{"float": true}`)
	err = json.Unmarshal(b, &d)
	is.Equal("value is a bool", err.Error())
}

func TestUnmarshalBool(t *testing.T) {
	is := is.New(t)

	type Data struct {
		Bool ft.Bool `json:"bool"`
	}
	d := Data{}

	// null
	b := []byte(`{"bool": null}`)
	err := json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(false, d.Bool.Bool) // Value must match

	b = []byte(`{}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(false, d.Bool.Bool) // Value must match

	// string
	b = []byte(`{"bool": "false"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(false, d.Bool.Bool) // Value must match

	b = []byte(`{"bool": "0"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(false, d.Bool.Bool) // Value must match

	b = []byte(`{"bool": ""}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(false, d.Bool.Bool) // Value must match

	b = []byte(`{"bool": "true"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Bool) // Value must match

	b = []byte(`{"bool": "1"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Bool) // Value must match

	b = []byte(`{"bool": "abc"}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Bool) // Value must match

	// int
	b = []byte(`{"bool": 0}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(false, d.Bool.Bool) // Value must match

	b = []byte(`{"bool": 1}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Bool) // Value must match

	b = []byte(`{"bool": -1}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Bool) // Value must match

	// float
	b = []byte(`{"bool": 1.23}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Bool) // Value must match

	b = []byte(`{"bool": -1.23}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Bool) // Value must match

	// bool
	b = []byte(`{"bool": false}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(false, d.Bool.Bool) // Value must match

	b = []byte(`{"bool": true}`)
	err = json.Unmarshal(b, &d)
	is.NoErr(err)
	is.Equal(true, d.Bool.Bool) // Value must match
}

func TestMarshalToJSON(t *testing.T) {
	is := is.New(t)

	type Data struct {
		String ft.String `json:"string"`
		Int    ft.Int    `json:"int"`
		Bool   ft.Bool   `json:"bool"`
		Float  ft.Float  `json:"float"`
	}

	// Valid values
	d := Data{}
	d.String = ft.StringFrom("foo")
	d.Int = ft.IntFrom(123)
	d.Bool = ft.BoolFrom(true)
	d.Float = ft.FloatFrom(1.618)
	b, err := json.Marshal(d)
	is.NoErr(err)
	is.Equal(`{"string":"foo","int":123,"bool":true,"float":1.618}`, string(b))

	// Empty value is empty
	d = Data{}
	b, err = json.Marshal(d)
	is.NoErr(err)
	is.Equal(`{"string":"","int":0,"bool":false,"float":0}`, string(b))

	// Missing properties are marshaled as empty.
	// Note that Marshal does not output the same as input to Unmarshal
	d = Data{}
	err = json.Unmarshal([]byte("{}"), &d)
	is.NoErr(err)
	b, err = json.Marshal(d)
	is.NoErr(err)
	is.Equal(`{"string":"","int":0,"bool":false,"float":0}`, string(b))

	// omitempty, see comments in TestNMarshalToJSON
	// https://github.com/mozey/ft/blob/main/ftnull_test.go#L450
	type Data2 struct {
		// Note omitempty is not "omitnull",
		// both empty and null will be omitted
		StringEmptyTag string    `json:"string_empty,omitempty"`
		StringNullTag  string    `json:"string_null,omitempty"`
		String         ft.String `json:"string,omitempty"`
		Int            ft.Int    `json:"int,omitempty"`
		Bool           ft.Bool   `json:"bool,omitempty"`
		Float          ft.Float  `json:"float,omitempty"`
	}
	d2 := Data2{}
	err = json.Unmarshal([]byte("{}"), &d2)
	is.NoErr(err)
	b, err = json.Marshal(d2)
	is.NoErr(err)
	is.Equal(`{"string":"","int":0,"bool":false,"float":0}`, string(b))
}

func TestTextTemplate(t *testing.T) {
	is := is.New(t)

	type data struct {
		Name        ft.String
		Number      ft.Int
		Active      ft.Bool
		Measurement ft.Float
	}

	tpl := template.Must(template.New("tpl").Parse(`
Name: {{.Name.String}}
Number: {{.Number.Int64}}
Active: {{.Active.Bool}}
Measurement: {{.Measurement.Float64}}`))

	buf := bytes.NewBufferString("")
	d := data{
		Name:        ft.StringFrom("John Doe"),
		Number:      ft.IntFrom(123),
		Active:      ft.BoolFrom(true),
		Measurement: ft.FloatFrom(1.618),
	}
	err := tpl.Execute(buf, d)
	is.NoErr(err)

	is.Equal(buf.String(), `
Name: John Doe
Number: 123
Active: true
Measurement: 1.618`)
}

// TestMapKeys verifies custom types implement encoding.TextMarshaler
func TestMapKeys(t *testing.T) {
	is := is.New(t)

	type data struct {
		StringMap map[ft.String]bool
		IntMap    map[ft.Int]bool
		BoolMap   map[ft.Bool]bool
		FloatMap  map[ft.Float]bool
	}
	d := data{
		StringMap: make(map[ft.String]bool),
		IntMap:    make(map[ft.Int]bool),
		BoolMap:   make(map[ft.Bool]bool),
		FloatMap:  make(map[ft.Float]bool),
	}
	d.StringMap[ft.StringFrom("foo")] = true
	d.IntMap[ft.IntFrom(123)] = true
	d.BoolMap[ft.BoolFrom(true)] = true
	d.FloatMap[ft.FloatFrom(1.618)] = true

	b, err := json.Marshal(d)
	is.NoErr(err)

	is.Equal(string(b), `{"StringMap":{"foo":true},"IntMap":{"123":true},"BoolMap":{"true":true},"FloatMap":{"1.618":true}}`)

	d = data{StringMap: make(map[ft.String]bool)}
	d.StringMap[ft.String{String: ""}] = true
	_, err = json.Marshal(d)
	is.NoErr(err) // Empty value is allowed

	d = data{IntMap: make(map[ft.Int]bool)}
	d.IntMap[ft.Int{Int64: 0}] = true
	_, err = json.Marshal(d)
	is.NoErr(err) // Empty value is allowed

	d = data{FloatMap: make(map[ft.Float]bool)}
	d.FloatMap[ft.Float{Float64: 0}] = true
	_, err = json.Marshal(d)
	is.NoErr(err) // Empty value is allowed

	d = data{BoolMap: make(map[ft.Bool]bool)}
	d.BoolMap[ft.Bool{Bool: false}] = true
	_, err = json.Marshal(d)
	is.NoErr(err) // Empty value is allowed

	b = []byte(`{"StringMap":{"foo":true},"IntMap":{"123":true},"BoolMap":{"true":true},"FloatMap":{"1.618":true}}`)
	d = data{}
	err = json.Unmarshal(b, &d)
	is.NoErr(err)

	compare, err := json.Marshal(d)
	is.NoErr(err)
	is.Equal(string(b), string(compare)) // Unmarshal and Marshal doesn't match
}
