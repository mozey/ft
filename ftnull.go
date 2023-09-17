package ft

import (
	"encoding/json"
	"strconv"
	"strings"
	"unicode"

	"github.com/guregu/null"
	"github.com/pkg/errors"
)

// NString can be used to decode any JSON value to string
type NString null.String

func NStringFrom(fs string) NString {
	return NString(null.StringFrom(fs))
}

// Clean removes non-graphic characters from the given string, see
// https://github.com/icza/gox/blob/master/stringsx/stringsx.go#L9 and
// https://stackoverflow.com/a/58994297/639133 However, above function also
// removes newlines, that is not desired?
func Clean(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) ||
			// Preserving some non-graphic characters
			string(r) == "\n" || string(r) == "\t" {
			return r
		}
		return -1
	}, s)
}

// MarshalJSON method with value receiver for String
// Method must not have a pointer receiver!
// See https://stackoverflow.com/a/21394657/639133
func (fs NString) MarshalJSON() ([]byte, error) {
	if !fs.Valid {
		return []byte(`null`), nil
	}
	// Control characters like "Start of Text" \u0002 breaks MarshalJSON
	// after calling Quote and wrapping the string with RawMessage.
	// First call Clean to remove non-graphic characters (except newline)
	return json.RawMessage(strconv.Quote(Clean(fs.String))), nil
}

// UnmarshalJSON for String
func (fs *NString) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		string(""), uint64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*fs = NString(null.String{})
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		*fs = NString(null.StringFrom(s))
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		*fs = NString(null.StringFrom(string(bArr[:])))
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		*fs = NString(null.StringFrom(string(bArr[:])))
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		*fs = NString(null.StringFrom(string(bArr[:])))
		return
	}

	return
}

// TextMarshaler and TextUnmarshaler interfaces
// must be implemented for custom types to be used as JSON map keys
// https://pkg.go.dev/encoding#TextMarshaler

func (fs NString) MarshalText() (text []byte, err error) {
	if !fs.Valid {
		// TODO Consider marshalling as empty value instead?
		return text, errors.Errorf("invalid ft.NString")
	}
	return []byte(fs.String), nil
}

func (fs *NString) UnmarshalText(text []byte) error {
	return fs.UnmarshalJSON(text)
}

// NInt can be used to decode any JSON value to int64.
// Strings that are not valid representation of a number will error.
// Boolean values will error
type NInt null.Int

func NIntFrom(fi int64) NInt {
	return NInt(null.IntFrom(fi))
}

// NIntFromString returns an Int for the given string
func NIntFromString(s string) (NInt, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return NInt{}, err
	}
	return NIntFrom(i), nil
}

// MarshalJSON method for Int
func (fi NInt) MarshalJSON() ([]byte, error) {
	if !fi.Valid {
		return []byte(`null`), nil
	}
	return []byte(strconv.FormatInt(fi.Int64, 10)), nil
}

// UnmarshalJSON method for Int
func (fi *NInt) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		string(""), int64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*fi = NInt(null.Int{})
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		if strings.TrimSpace(s) == "" {
			// Empty string parses as null
			*fi = NInt(null.Int{})
			return
		}
		i, err2 := strconv.ParseInt(s, 10, 64)
		if err2 != nil {
			return err2
		}
		*fi = NInt(null.IntFrom(i))
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		*fi = NInt(null.IntFrom(i))
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		*fi = NInt(null.IntFrom(int64(f)))
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		return errors.Errorf("value is a bool")
	}

	return
}

func (fi NInt) MarshalText() (text []byte, err error) {
	if !fi.Valid {
		return text, errors.Errorf("invalid ft.NInt")
	}
	return []byte(strconv.FormatInt(fi.Int64, 10)), nil
}

func (fi *NInt) UnmarshalText(text []byte) error {
	return fi.UnmarshalJSON(text)
}

// NFloat can be used to decode any JSON value to int64.
// Strings that are not valid representation of a number will error.
// Boolean values will error
type NFloat null.Float

func NFloatFrom(ff float64) NFloat {
	return NFloat(null.FloatFrom(ff))
}

// MarshalJSON method for Float
func (fi NFloat) MarshalJSON() ([]byte, error) {
	if !fi.Valid {
		return []byte(`null`), nil
	}
	return []byte(strconv.FormatFloat(fi.Float64, 'f', -1, 64)), nil
}

// UnmarshalJSON method for Float
func (fi *NFloat) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		string(""), int64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*fi = NFloat(null.Float{})
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		i, err2 := strconv.ParseFloat(s, 64)
		if err2 != nil {
			return err2
		}
		*fi = NFloat(null.FloatFrom(i))
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		*fi = NFloat(null.FloatFrom(float64(i)))
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		*fi = NFloat(null.FloatFrom(f))
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		return errors.Errorf("value is a bool")
	}

	return
}

func (ff NFloat) MarshalText() (text []byte, err error) {
	if !ff.Valid {
		return text, errors.Errorf("invalid ft.NFloat")
	}
	return []byte(strconv.FormatFloat(ff.Float64, 'f', -1, 64)), nil
}

func (ff *NFloat) UnmarshalText(text []byte) error {
	return ff.UnmarshalJSON(text)
}

// NBool can be used to decode any JSON value to bool.
// Empty strings as well as "false" and "0" evaluate to false,
// all other strings are true.
// Numbers equal to 0 will evaluate to false,
// all other numbers are true.
type NBool null.Bool

func NBoolFrom(fb bool) NBool {
	return NBool(null.BoolFrom(fb))
}

// MarshalJSON method for Bool
func (fb NBool) MarshalJSON() ([]byte, error) {
	if !fb.Valid {
		return []byte(`null`), nil
	}
	return []byte(strconv.FormatBool(fb.Bool)), nil
}

// UnmarshalJSON method for Bool
func (fb *NBool) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		string(""), int64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*fb = NBool(null.Bool{})
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		compare := strings.ToLower(strings.TrimSpace(s))
		if compare == "false" || compare == "0" || compare == "" {
			*fb = NBool(null.BoolFrom(false))
			return
		}
		*fb = NBool(null.BoolFrom(true))
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		if i == 0 {
			*fb = NBool(null.BoolFrom(false))
			return
		}
		*fb = NBool(null.BoolFrom(true))
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		if f == 0 {
			*fb = NBool(null.BoolFrom(false))
			return
		}
		*fb = NBool(null.BoolFrom(true))
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		*fb = NBool(null.BoolFrom(b))
		return
	}

	return
}

func (fb NBool) MarshalText() (text []byte, err error) {
	if !fb.Valid {
		return text, errors.Errorf("invalid ft.NBool")
	}
	return []byte(strconv.FormatBool(fb.Bool)), nil
}

func (fb *NBool) UnmarshalText(text []byte) error {
	return fb.UnmarshalJSON(text)
}
