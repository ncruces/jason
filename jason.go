// Package jason helps deal with dynamic JSON.
package jason

import "encoding/json"

// Object is a dynamic JSON object:
// an unordered set of name/value pairs.
//
// Example of a JSON object literal:
//   jason.Object{"abc": true}
type Object = map[string]any

// Array is a dynamic JSON array:
// an ordered collection of values.
//
// Example of a JSON array literal:
//   jason.Array{10, true}
type Array = []any

// Number is an encoded JSON number.
//
// It has arbitrary precision, and is easily converted to a float64 or int64.
//
// Example of a JSON number literal:
//   jason.Number("10")
type Number = json.Number

// Value is an encoded JSON value.
//
// Like [json.RawMessage], it implements [json.Marshaler] and [json.Unmarshaler]
// and can be used to delay/precompute JSON decoding/encoding.
//
// Example of a JSON value literal:
//   jason.Value("false")
type Value []byte

// From marshals v into a [Value], panics on error.
//
// Example:
//   jason.From(time.Now())
func From(v any) Value {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

// To unmarshals j into a value of type T, panics on error.
//
// Example:
//   jason.To[time.Time](j)
func To[T any](j Value) (v T) {
	err := json.Unmarshal(j, &v)
	if err != nil {
		panic(err)
	}
	return v
}

// Is checks if j can unmarshal into type T.
//
// Example:
//   if jason.Is[time.Time](j) { ... }
func Is[T any](j Value) bool {
	_, err := Maybe[T](j)
	return err == nil
}

// Maybe unmarshals j into a value of type T.
//
// Example:
//   v, err = jason.Maybe[time.Time](j)
func Maybe[T any](j Value) (v T, err error) {
	err = json.Unmarshal(j, &v)
	return v, err
}

// MarshalJSON returns j as the JSON encoding of j.
func (j Value) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

// UnmarshalJSON sets *j to a copy of data.
func (j *Value) UnmarshalJSON(data []byte) error {
	*j = append((*j)[0:0], data...)
	return nil
}

// String returns j as a string.
func (j Value) String() string {
	b, _ := j.MarshalJSON()
	return string(b)
}

// GoString returns a Go representation of j.
func (j Value) GoString() string {
	return "jason.Value(`" + j.String() + "`)"
}
