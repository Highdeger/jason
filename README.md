# jason

[![Release](https://img.shields.io/badge/release-0.1.0-blue)](https://github.com/highdeger/jason/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/highdeger/jason.svg)](https://pkg.go.dev/github.com/highdeger/jason)
[![Go Report Card](https://goreportcard.com/badge/github.com/highdeger/jason)](https://goreportcard.com/report/github.com/highdeger/jason)

Tiny and fail-safe JSON representation in pure Go.

# Examples
## Usage of JSON Object and JSON Array
```go
package main

import (
	"fmt"

	"github.com/highdeger/jason"
)

func main() {
	jsonObject := jason.NewObject().
		String("k1", "value").
		Int("k2", 6).
		Int8("k3", 6).
		Int16("k4", 6).
		Int32("k5", 6).
		Int64("k6", 6).
		Uint("k7", 6).
		Uint8("k8", 6).
		Uint16("k9", 6).
		Uint32("k10", 6).
		Uint64("k11", 6).
		Uint64("k12", 6).
		Float32("k13", 6.6).
		Float64("k14", 6.6).
		Object("k15", jason.NewObject().
			String("k15-1", "value").
			Float32("k15-2", 6.6)).
		Bool("k16", true).
		Null("k17").
		Array("k18", jason.NewArray().
			String("value").
			Bool(true))

	jsonArray := jason.NewArray().
		String("value").
		Int(6).
		Int8(6).
		Int16(6).
		Int32(6).
		Int64(6).
		Uint(6).
		Uint8(6).
		Uint16(6).
		Uint32(6).
		Uint64(6).
		Float32(6.6).
		Float64(6.6).
		Object(jason.NewObject().
			String("k1", "value").
			Float32("k2", 6.6)).
		Bool(true).
		Null().
		Array(jason.NewArray().
			String("value").
			Bool(true))

	fmt.Println("jsonObject.Minify:")
	fmt.Println(jsonObject.Minify(true))
	fmt.Println()
	fmt.Println("jsonObject.Format:")
	fmt.Println(jsonObject.Format(true, "-", "  "))
	fmt.Println()
	fmt.Println("jsonArray.Minify:")
	fmt.Println(jsonArray.Minify(true))
	fmt.Println()
	fmt.Println("jsonArray.Format:")
	fmt.Println(jsonArray.Format(true, "-", "  "))
}
```
output:
```
jsonObject.Minify:
{"k1":"value","k10":6,"k11":6,"k12":6,"k13":6.6,"k14":6.6,"k15":{"k15-1":"value","k15-2":6.6},"k16":true,"k17":null,"k18":["value",true],"k2":6,"k3":6,"k4":6,"k5":6,"k6":6,"k7":6,"k8":6,"k9":6}

jsonObject.Format:
{
-  "k1": "value",
-  "k10": 6,
-  "k11": 6,
-  "k12": 6,
-  "k13": 6.6,
-  "k14": 6.6,
-  "k15": {
-    "k15-1": "value",
-    "k15-2": 6.6
-  },
-  "k16": true,
-  "k17": null,
-  "k18": [
-    "value",
-    true
-  ],
-  "k2": 6,
-  "k3": 6,
-  "k4": 6,
-  "k5": 6,
-  "k6": 6,
-  "k7": 6,
-  "k8": 6,
-  "k9": 6
-}

jsonArray.Minify:
["value",6,6,6,6,6,6,6,6,6,6,6.6,6.6,{"k1":"value","k2":6.6},true,null,["value",true]]

jsonArray.Format:
[
-  "value",
-  6,
-  6,
-  6,
-  6,
-  6,
-  6,
-  6,
-  6,
-  6,
-  6,
-  6.6,
-  6.6,
-  {
-    "k1": "value",
-    "k2": 6.6
-  },
-  true,
-  null,
-  [
-    "value",
-    true
-  ]
-]
```
output if object and array are empty:
```
emptyObject.Minify:
{}

emptyObject.Format:
{}

emptyArray.Minify:
[]

emptyArray.Format:
[]
```
## Comparison of Safe and Not-Safe HTML Encoding
```go
package main

import (
	"fmt"

	"github.com/highdeger/jason"
)

func main() {
	jsonObject := jason.NewObject().
		String("key <of> element", "value <of> element")

	fmt.Println("jsonObject.Minify:")
	fmt.Printf("html-not-safe => %s\n", jsonObject.Minify(false))
	fmt.Printf("html-safe => %s\n", jsonObject.Minify(true))
	fmt.Println()
	fmt.Println("jsonObject.Format:")
	fmt.Printf("html-not-safe => %s\n", jsonObject.Format(false, "-", "  "))
	fmt.Printf("html-safe => %s\n", jsonObject.Format(true, "-", "  "))
}
```
output:
```
jsonObject.Minify:
html-not-safe => {"key <of> element":"value <of> element"}
html-safe => {"key \u003cof\u003e element":"value \u003cof\u003e element"}

jsonObject.Format:
html-not-safe => {
-  "key <of> element": "value <of> element"
-}
html-safe => {
-  "key \u003cof\u003e element": "value \u003cof\u003e element"
-}
```
