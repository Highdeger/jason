package jason

import (
	"testing"
)

func TestArray(t *testing.T) {
	arr := NewArray().
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
		Object(NewObject().
			String("k1", "value").
			Float32("k2", 6.6)).
		Bool(true).
		Null().
		Array(NewArray().
			String("value").
			Bool(true))

	expectedLen := 17
	outputLen := arr.Len()
	if outputLen != expectedLen {
		t.Fatalf("array len is incorrect: expected %d, got %d", expectedLen, outputLen)
	}

	expectedMinified := `["value",6,6,6,6,6,6,6,6,6,6,6.6,6.6,{"k1":"value","k2":6.6},true,null,["value",true]]`
	outputMinified := arr.Minify(true)
	if outputMinified != expectedMinified {
		t.Fatalf("array minified incorrectly:\n\texpected %s\n\tgot %s", expectedMinified, outputMinified)
	}

	expectedFormated := `[
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
-]`
	outputFormated := arr.Format(true, "-", "  ")
	if outputFormated != expectedFormated {
		t.Fatalf("array formated incorrectly:\n\texpected %s\n\tgot %s", expectedFormated, outputFormated)
	}
}
