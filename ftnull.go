package ft

import (
	"encoding/json"
	"strconv"
	"strings"
	"unicode"

	"github.com/guregu/null"
	"github.com/pkg/errors"
)

// NullString can be used to decode any JSON value to string
type NullString null.String

func StringFrom(fs string) NullString {
	return NullString(null.StringFrom(fs))
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
func (fs NullString) MarshalJSON() ([]byte, error) {
	if !fs.Valid {
		return []byte(`null`), nil
	}
	// Control characters like "Start of Text" \u0002 breaks MarshalJSON
	// after calling Quote and wrapping the string with RawMessage.
	// First call Clean to remove non-graphic characters (except newline)
	return json.RawMessage(strconv.Quote(Clean(fs.String))), nil
}

// UnmarshalJSON for String
func (fs *NullString) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		string(""), uint64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*fs = NullString(null.String{})
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		*fs = NullString(null.StringFrom(s))
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		*fs = NullString(null.StringFrom(string(bArr[:])))
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		*fs = NullString(null.StringFrom(string(bArr[:])))
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		*fs = NullString(null.StringFrom(string(bArr[:])))
		return
	}

	return
}

// NullInt can be used to decode any JSON value to int64.
// Strings that are not valid representation of a number will error.
// Boolean values will error
type NullInt null.Int

func IntFrom(fi int64) NullInt {
	return NullInt(null.IntFrom(fi))
}

// IntFromString returns an Int for the given string
func IntFromString(s string) (NullInt, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return NullInt{}, err
	}
	return IntFrom(i), nil
}

// MarshalJSON method for Int
func (fi NullInt) MarshalJSON() ([]byte, error) {
	if !fi.Valid {
		return []byte(`null`), nil
	}
	return []byte(strconv.FormatInt(fi.Int64, 10)), nil
}

// UnmarshalJSON method for Int
func (fi *NullInt) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		string(""), int64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*fi = NullInt(null.Int{})
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		if strings.TrimSpace(s) == "" {
			// Empty string parses as null
			*fi = NullInt(null.Int{})
			return
		}
		i, err2 := strconv.ParseInt(s, 10, 64)
		if err2 != nil {
			return err2
		}
		*fi = NullInt(null.IntFrom(i))
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		*fi = NullInt(null.IntFrom(i))
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		*fi = NullInt(null.IntFrom(int64(f)))
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		return errors.Errorf("value is a bool")
	}

	return
}

// NullFloat can be used to decode any JSON value to int64.
// Strings that are not valid representation of a number will error.
// Boolean values will error
type NullFloat null.Float

func FloatFrom(ff float64) NullFloat {
	return NullFloat(null.FloatFrom(ff))
}

// MarshalJSON method for Float
func (fi NullFloat) MarshalJSON() ([]byte, error) {
	if !fi.Valid {
		return []byte(`null`), nil
	}
	return []byte(strconv.FormatFloat(fi.Float64, 'f', -1, 64)), nil
}

// UnmarshalJSON method for Float
func (fi *NullFloat) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		string(""), int64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*fi = NullFloat(null.Float{})
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		i, err2 := strconv.ParseFloat(s, 64)
		if err2 != nil {
			return err2
		}
		*fi = NullFloat(null.FloatFrom(i))
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		*fi = NullFloat(null.FloatFrom(float64(i)))
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		*fi = NullFloat(null.FloatFrom(f))
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		return errors.Errorf("value is a bool")
	}

	return
}

// NullBool can be used to decode any JSON value to bool.
// Empty strings as well as "false" and "0" evaluate to false,
// all other strings are true.
// Numbers equal to 0 will evaluate to false,
// all other numbers are true.
type NullBool null.Bool

func BoolFrom(fb bool) NullBool {
	return NullBool(null.BoolFrom(fb))
}

// MarshalJSON method for Bool
func (fb NullBool) MarshalJSON() ([]byte, error) {
	if !fb.Valid {
		return []byte(`null`), nil
	}
	return []byte(strconv.FormatBool(fb.Bool)), nil
}

// UnmarshalJSON method for Bool
func (fb *NullBool) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		string(""), int64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*fb = NullBool(null.Bool{})
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		compare := strings.ToLower(strings.TrimSpace(s))
		if compare == "false" || compare == "0" || compare == "" {
			*fb = NullBool(null.BoolFrom(false))
			return
		}
		*fb = NullBool(null.BoolFrom(true))
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		if i == 0 {
			*fb = NullBool(null.BoolFrom(false))
			return
		}
		*fb = NullBool(null.BoolFrom(true))
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		if f == 0 {
			*fb = NullBool(null.BoolFrom(false))
			return
		}
		*fb = NullBool(null.BoolFrom(true))
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		*fb = NullBool(null.BoolFrom(b))
		return
	}

	return
}
