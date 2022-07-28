// Package jason helps deal with dynamic JSON.
package jason

import "encoding/json"

// Object is a JSON object:
// an unordered set of name/value pairs.
//
// Example of an Object literal:
//   jason.Object{"abc": true}
type Object = map[string]any

// Array is a JSON array:
// an ordered collection of values.
//
// Example of an Array literal:
//   jason.Array{1, 2, 3}
type Array = []any

// Number is an encoded JSON number.
//
// It has arbitrary precision,
// and can be easily converted to a float64 or int64.
//
// Example of a Number literal:
//   jason.Number("10")
type Number = json.Number

// RawValue is a raw encoded JSON value.
//
// It implements [json.Marshaler] and [json.Unmarshaler],
// and can be used to delay/precompute JSON decoding/encoding.
//
// Example of a RawValue literal:
//   jason.RawValue("false")
type RawValue = json.RawMessage

// From marshals v into a RawValue, panics on error.
//
// Example of creating a RawValue from a time instant:
//   jason.From(time.Now())
func From(v any) RawValue {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

// ToA unmarshals j into a value of type T, panics on error.
//
// Example of converting j into a time instant:
//   jason.ToA[time.Time](j)
func ToA[T any](j RawValue) (v T) {
	err := json.Unmarshal(j, &v)
	if err != nil {
		panic(err)
	}
	return v
}

// AsA unmarshals j into a value of type T.
//
// Example of converting j into a time instant:
//   if v, err := jason.AsA[time.Time](j); err == nil { ... }
func AsA[T any](j RawValue) (v T, err error) {
	err = json.Unmarshal(j, &v)
	return v, err
}

// IsA checks whether j can be unmarshaled into type T.
//
// Example of testing whether j can be converted into a time instant:
//   if jason.IsA[time.Time](j) { ... }
func IsA[T any](j RawValue) bool {
	_, err := AsA[T](j)
	return err == nil
}

// RawObject is an object of RawValue's.
type RawObject = map[string]RawValue

// RawArray is an array of RawValue's.
type RawArray = []RawValue
