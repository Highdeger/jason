package jason

import (
	"bytes"
	"encoding/json"
	"strings"
)

// Array is representation of a JSON array.
type Array interface {
	Minify(bool) string
	Format(bool, string, string) string
	MinifyBytes(bool) []byte
	FormatBytes(bool, string, string) []byte
	Len() int
	String(string) Array
	Int(int) Array
	Int8(int8) Array
	Int16(int16) Array
	Int32(int32) Array
	Int64(int64) Array
	Uint(uint) Array
	Uint8(uint8) Array
	Uint16(uint16) Array
	Uint32(uint32) Array
	Uint64(uint64) Array
	Float32(float32) Array
	Float64(float64) Array
	Bool(bool) Array
	Null() Array
	Object(Object) Array
	Array(Array) Array
}

// array is underlying struct for JSON arrays.
type array []any

//
// static methods
//

// NewArray creates a new empty JSON array.
func NewArray() Array {
	return &array{}
}

//
// public methods
//

// Minify returns the string representation of the JSON array,
// flatten and without whitespaces.
func (r *array) Minify(htmlSafe bool) string {
	return r.Format(htmlSafe, "", "")
}

// Format returns the string representation of the JSON array,
// indented and multi-line if the object is non-empty.
func (r *array) Format(htmlSafe bool, prefix, indent string) string {
	return strings.TrimSpace(r.encode(htmlSafe, prefix, indent).String())
}

// MinifyBytes returns the byte-array representation of the JSON array,
// flatten and without whitespaces.
func (r *array) MinifyBytes(htmlSafe bool) []byte {
	return r.FormatBytes(htmlSafe, "", "")
}

// FormatBytes returns the byte-array representation of the JSON array,
// indented and multi-line if the object is non-empty.
func (r *array) FormatBytes(htmlSafe bool, prefix, indent string) []byte {
	return bytes.TrimSpace(r.encode(htmlSafe, prefix, indent).Bytes())
}

// Len returns the count of elements.
func (r *array) Len() int {
	return len(*r)
}

// String adds a string.
func (r *array) String(v string) Array {
	r.put(v)
	return r
}

// Int adds an integer number (int).
func (r *array) Int(v int) Array {
	r.put(v)
	return r
}

// Int8 adds an integer number (int8).
func (r *array) Int8(v int8) Array {
	r.put(v)
	return r
}

// Int16 adds an integer number (int16).
func (r *array) Int16(v int16) Array {
	r.put(v)
	return r
}

// Int32 adds an integer number (int32).
func (r *array) Int32(v int32) Array {
	r.put(v)
	return r
}

// Int64 adds an integer number (int64).
func (r *array) Int64(v int64) Array {
	r.put(v)
	return r
}

// Uint adds a non-negative integer number (uint).
func (r *array) Uint(v uint) Array {
	r.put(v)
	return r
}

// Uint8 adds a non-negative integer number (uint8).
func (r *array) Uint8(v uint8) Array {
	r.put(v)
	return r
}

// Uint16 adds a non-negative integer number (uint16).
func (r *array) Uint16(v uint16) Array {
	r.put(v)
	return r
}

// Uint32 adds a non-negative integer number (uint32).
func (r *array) Uint32(v uint32) Array {
	r.put(v)
	return r
}

// Uint64 adds a non-negative integer number (uint64).
func (r *array) Uint64(v uint64) Array {
	r.put(v)
	return r
}

// Float32 adds a decimal number (float32).
func (r *array) Float32(v float32) Array {
	r.put(v)
	return r
}

// Float64 adds a decimal number (float64).
func (r *array) Float64(v float64) Array {
	r.put(v)
	return r
}

// Bool adds a boolean.
func (r *array) Bool(v bool) Array {
	r.put(v)
	return r
}

// Null adds a null value.
func (r *array) Null() Array {
	r.put(nil)
	return r
}

// Object adds a JSON object.
func (r *array) Object(v Object) Array {
	r.put(v)
	return r
}

// Array adds an JSON array.
func (r *array) Array(v Array) Array {
	r.put(v)
	return r
}

//
// private methods
//

// put appends a new element to the JSON array.
func (r *array) put(v any) {
	*r = append(*r, v)
}

// encode encodes the JSON array into a byte-array buffer.
func (r *array) encode(escapeHTML bool, prefix, indent string) *bytes.Buffer {
	w := &bytes.Buffer{}
	e := json.NewEncoder(w)
	e.SetEscapeHTML(escapeHTML)
	e.SetIndent(prefix, indent)

	err := e.Encode(r)
	if err != nil {
		panic(err)
	}

	return w
}
