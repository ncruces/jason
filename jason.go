// Package jason helps deal with dynamic JSON.
package jason

import "encoding/json"

// Object is a JSON object:
// an unordered set of name/value pairs.
//
// Example of an Object literal:
//   jason.Object[bool]{"abc": true}
type Object[T any] map[string]T

// Array is a JSON array:
// an ordered collection of values.
//
// Example of an Array literal:
//   jason.Array[int]{1, 2, 3}
type Array[T any] []T

// Number is an encoded JSON number.
//
// It has arbitrary precision, and is easily converted to a float64 or int64.
//
// Example of a Number literal:
//   jason.Number("10")
type Number = json.Number

// Value is an encoded JSON value.
//
// Like [json.RawMessage], it implements [json.Marshaler] and [json.Unmarshaler]
// and can be used to delay/precompute JSON decoding/encoding.
//
// Example of a Value literal:
//   jason.Value("false")
type Value []byte

// From marshals v into a Value, panics on error.
//
// Example of creating a Value from a time instant:
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
// Example of converting j into a time instant:
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
// Example of testing if j can be converted into a time instant:
//   if jason.Is[time.Time](j) { ... }
func Is[T any](j Value) bool {
	_, err := Maybe[T](j)
	return err == nil
}

// Maybe unmarshals j into a value of type T.
//
// Example of converting j into a time instant:
//   if v, err := jason.Maybe[time.Time](j); err == nil { ... }
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

// ValueObject is an Object of Value.
type ValueObject = Object[Value]

// ValueArray is an Array of Value.
type ValueArray = Array[Value]
