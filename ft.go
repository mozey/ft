package ft

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// String can be used to decode any JSON value to string
type String string

// MarshalJSON method with value receiver for String
// Method must not have a pointer receiver!
// See https://stackoverflow.com/a/21394657/639133
func (ns String) MarshalJSON() ([]byte, error) {
	// Control characters like "Start of Text" \u0002 breaks MarshalJSON
	// after calling Quote and wrapping the string with RawMessage.
	// First call Clean to remove non-graphic characters (except newline)
	return json.RawMessage(strconv.Quote(Clean(string(ns)))), nil
}

// UnmarshalJSON for String
func (ns *String) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		"", uint64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*ns = ""
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		*ns = String(s)
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		*ns = String(bArr[:])
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		*ns = String(bArr[:])
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		*ns = String(bArr[:])
		return
	}

	return
}

// Int can be used to decode any JSON value to int64.
// Strings that are not valid representation of a number will error.
// Boolean values will error
type Int int64

// MarshalJSON method for Int
func (ni Int) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(ni), 10)), nil
}

// UnmarshalJSON method for Int
func (ni *Int) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		"", int64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*ni = Int(0)
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		i, err2 := strconv.ParseInt(s, 10, 64)
		if err2 != nil {
			// Value is null if int could not be parsed from the string
			//*ni = Int(0) // This is not a good idea...
			return err2
		}
		*ni = Int(i)
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		*ni = Int(i)
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		*ni = Int(int64(f))
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		//*ni = Int(0) // This is not a good idea...
		return errors.WithStack(fmt.Errorf("value is a bool"))
	}

	return
}

// Float can be used to decode any JSON value to int64.
// Strings that are not valid representation of a number will error.
// Boolean values will error
type Float float64

// MarshalJSON method for Float
func (nf Float) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(nf), 'f', -1, 64)), nil
}

// UnmarshalJSON method for Float
func (nf *Float) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		"", int64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*nf = Float(0)
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		i, err2 := strconv.ParseFloat(s, 64)
		if err2 != nil {
			// Value is null if int could not be parsed from the string
			//*ni = Float(0) // This is not a good idea...
			return err2
		}
		*nf = Float(i)
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		*nf = Float(i)
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		*nf = Float(f)
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		//*nf = Float(0) // This is not a good idea...
		return errors.WithStack(fmt.Errorf("value is a bool"))
	}

	return
}

// Bool can be used to decode any JSON value to bool.
// Empty strings as well as "false" and "0" evaluate to false,
// all other strings are true.
// Numbers equal to 0 will evaluate to false,
// all other numbers are true.
type Bool bool

// MarshalJSON method for Bool
func (nb Bool) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatBool(bool(nb))), nil
}

// UnmarshalJSON method for Bool
func (nb *Bool) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		"", int64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*nb = false
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		compare := strings.ToLower(strings.TrimSpace(s))
		if compare == "false" || compare == "0" || compare == "" {
			*nb = false
			return
		}
		*nb = true
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		if i == 0 {
			*nb = false
			return
		}
		*nb = true
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		if f == 0 {
			*nb = false
			return
		}
		*nb = true
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		*nb = Bool(b)
		return
	}

	return
}
