package jason

import (
	"bytes"
	"encoding/json"
	"strings"
)

// Object is representation of a JSON object.
type Object interface {
	Minify(bool) string
	Format(bool, string, string) string
	MinifyBytes(bool) []byte
	FormatBytes(bool, string, string) []byte
	Len() int
	String(string, string) Object
	Int(string, int) Object
	Int8(string, int8) Object
	Int16(string, int16) Object
	Int32(string, int32) Object
	Int64(string, int64) Object
	Uint(string, uint) Object
	Uint8(string, uint8) Object
	Uint16(string, uint16) Object
	Uint32(string, uint32) Object
	Uint64(string, uint64) Object
	Float32(string, float32) Object
	Float64(string, float64) Object
	Bool(string, bool) Object
	Null(string) Object
	Object(string, Object) Object
	Array(string, Array) Object
}

// object is underlying struct for JSON objects.
type object map[string]any

//
// static methods
//

// NewObject creates a new empty JSON Object.
func NewObject() Object {
	return &object{}
}

//
// public methods
//

// Minify returns the string representation of the JSON object,
// flatten and without whitespaces.
func (r *object) Minify(htmlSafe bool) string {
	return r.Format(htmlSafe, "", "")
}

// Format returns the string representation of the JSON object,
// indented and multi-line if the object is non-empty.
func (r *object) Format(htmlSafe bool, prefix, indent string) string {
	return strings.TrimSpace(r.encode(htmlSafe, prefix, indent).String())
}

// MinifyBytes returns the byte-array representation of the JSON object,
// flatten and without whitespaces.
func (r *object) MinifyBytes(htmlSafe bool) []byte {
	return r.FormatBytes(htmlSafe, "", "")
}

// FormatBytes returns the byte-array representation of the JSON object,
// indented and multi-line if the object is non-empty.
func (r *object) FormatBytes(htmlSafe bool, prefix, indent string) []byte {
	return bytes.TrimSpace(r.encode(htmlSafe, prefix, indent).Bytes())
}

// Len returns the count of elements.
func (r *object) Len() int {
	return len(*r)
}

// String adds a string.
func (r *object) String(k string, v string) Object {
	r.set(k, v)
	return r
}

// Int adds an integer number (int).
func (r *object) Int(k string, v int) Object {
	r.set(k, v)
	return r
}

// Int8 adds an integer number (int8).
func (r *object) Int8(k string, v int8) Object {
	r.set(k, v)
	return r
}

// Int16 adds an integer number (int16).
func (r *object) Int16(k string, v int16) Object {
	r.set(k, v)
	return r
}

// Int32 adds an integer number (int32).
func (r *object) Int32(k string, v int32) Object {
	r.set(k, v)
	return r
}

// Int64 adds an integer number (int64).
func (r *object) Int64(k string, v int64) Object {
	r.set(k, v)
	return r
}

// Uint adds a non-negative integer number (uint).
func (r *object) Uint(k string, v uint) Object {
	r.set(k, v)
	return r
}

// Uint8 adds a non-negative integer number (uint8).
func (r *object) Uint8(k string, v uint8) Object {
	r.set(k, v)
	return r
}

// Uint16 adds a non-negative integer number (uint16).
func (r *object) Uint16(k string, v uint16) Object {
	r.set(k, v)
	return r
}

// Uint32 adds a non-negative integer number (uint32).
func (r *object) Uint32(k string, v uint32) Object {
	r.set(k, v)
	return r
}

// Uint64 adds a non-negative integer number (uint64).
func (r *object) Uint64(k string, v uint64) Object {
	r.set(k, v)
	return r
}

// Float32 adds a decimal number (float32).
func (r *object) Float32(k string, v float32) Object {
	r.set(k, v)
	return r
}

// Float64 adds a decimal number (float64).
func (r *object) Float64(k string, v float64) Object {
	r.set(k, v)
	return r
}

// Bool adds a boolean.
func (r *object) Bool(k string, v bool) Object {
	r.set(k, v)
	return r
}

// Null adds a null value.
func (r *object) Null(k string) Object {
	r.set(k, nil)
	return r
}

// Object adds a JSON object.
func (r *object) Object(k string, v Object) Object {
	r.set(k, v)
	return r
}

// Array adds an JSON array.
func (r *object) Array(k string, v Array) Object {
	r.set(k, v)
	return r
}

//
// private methods
//

// set sets a value for the specified key of the JSON object.
func (r *object) set(k string, v any) {
	(*r)[k] = v
}

// encode encodes the JSON object into a byte-array buffer.
func (r *object) encode(escapeHTML bool, prefix, indent string) *bytes.Buffer {
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
