package ft

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// String can be used to decode any JSON value to string
type String struct {
	String string
	Valid  bool // Always true
}

func StringFrom(fs string) String {
	return String{String: fs, Valid: true}
}

// MarshalJSON method with value receiver for String
// Method must not have a pointer receiver!
// See https://stackoverflow.com/a/21394657/639133
func (fs String) MarshalJSON() ([]byte, error) {
	// Control characters like "Start of Text" \u0002 breaks MarshalJSON
	// after calling Quote and wrapping the string with RawMessage.
	// First call Clean to remove non-graphic characters (except newline)
	return json.RawMessage(strconv.Quote(Clean(fs.String))), nil
}

// UnmarshalJSON for String
func (fs *String) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		"", uint64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*fs = StringFrom("")
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		*fs = StringFrom(s)
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		*fs = StringFrom(string(bArr[:]))
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		*fs = StringFrom(string(bArr[:]))
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		*fs = StringFrom(string(bArr[:]))
		return
	}

	return
}

// Int can be used to decode any JSON value to int64.
// Strings that are not valid representation of a number will error.
// Boolean values will error
type Int struct {
	Int   int64
	Valid bool // Always true
}

func IntFrom(fi int64) Int {
	return Int{Int: fi, Valid: true}
}

// MarshalJSON method for Int
func (fi Int) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(fi.Int, 10)), nil
}

// UnmarshalJSON method for Int
func (fi *Int) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		"", int64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*fi = IntFrom(0)
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		i, err2 := strconv.ParseInt(s, 10, 64)
		if err2 != nil {
			// Value is null if int could not be parsed from the string
			//*fi = Int(0) // This is not a good idea...
			return err2
		}
		*fi = IntFrom(i)
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		*fi = IntFrom(i)
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		*fi = IntFrom(int64(f))
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		//*fi = Int(0) // This is not a good idea...
		return errors.WithStack(fmt.Errorf("value is a bool"))
	}

	return
}

// Float can be used to decode any JSON value to int64.
// Strings that are not valid representation of a number will error.
// Boolean values will error
type Float struct {
	Float float64
	Valid bool // Always true
}

func FloatFrom(ff float64) Float {
	return Float{Float: ff, Valid: true}
}

// MarshalJSON method for Float
func (ff Float) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(ff.Float, 'f', -1, 64)), nil
}

// UnmarshalJSON method for Float
func (ff *Float) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		"", int64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*ff = FloatFrom(0)
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		i, err2 := strconv.ParseFloat(s, 64)
		if err2 != nil {
			// Value is null if int could not be parsed from the string
			//*fi = Float(0) // This is not a good idea...
			return err2
		}
		*ff = FloatFrom(i)
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		*ff = FloatFrom(float64(i))
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		*ff = FloatFrom(f)
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		//*ff = Float(0) // This is not a good idea...
		return errors.WithStack(fmt.Errorf("value is a bool"))
	}

	return
}

// Bool can be used to decode any JSON value to bool.
// Empty strings as well as "false" and "0" evaluate to false,
// all other strings are true.
// Numbers equal to 0 will evaluate to false,
// all other numbers are true.
type Bool struct {
	Bool  bool
	Valid bool // Always true
}

func BoolFrom(fb bool) Bool {
	return Bool{Bool: fb, Valid: true}
}

// MarshalJSON method for Bool
func (fb Bool) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatBool(fb.Bool)), nil
}

// UnmarshalJSON method for Bool
func (fb *Bool) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		"", int64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*fb = BoolFrom(false)
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		compare := strings.ToLower(strings.TrimSpace(s))
		if compare == "false" || compare == "0" || compare == "" {
			*fb = BoolFrom(false)
			return
		}
		*fb = BoolFrom(true)
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		if i == 0 {
			*fb = BoolFrom(false)
			return
		}
		*fb = BoolFrom(true)
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		if f == 0 {
			*fb = BoolFrom(false)
			return
		}
		*fb = BoolFrom(true)
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		*fb = BoolFrom(b)
		return
	}

	return
}
