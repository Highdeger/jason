package jason

import (
	"testing"
)

func TestObject(t *testing.T) {
	obj := NewObject().
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
		Object("k15", NewObject().
			String("k15-1", "value").
			Float32("k15-2", 6.6)).
		Bool("k16", true).
		Null("k17").
		Array("k18", NewArray().
			String("value").
			Bool(true))

	expectedLen := 17
	outputLen := obj.Len()
	if outputLen != expectedLen {
		t.Fatalf("object len is incorrect: expected %d, got %d", expectedLen, outputLen)
	}

	expectedMinified := `{"k1":"value","k10":6,"k11":6,"k12":6.6,"k13":6.6,"k14":{"k14-1":"value","k14-2":6.6},"k15":true,"k16":null,"k17":["value",true],"k2":6,"k3":6,"k4":6,"k5":6,"k6":6,"k7":6,"k8":6,"k9":6}`
	outputMinified := obj.Minify(true)
	if outputMinified != expectedMinified {
		t.Fatalf("object minified incorrectly:\n\texpected %s\n\tgot %s", expectedMinified, outputMinified)
	}

	expectedFormated := `{
-  "k1": "value",
-  "k10": 6,
-  "k11": 6,
-  "k12": 6.6,
-  "k13": 6.6,
-  "k14": {
-    "k14-1": "value",
-    "k14-2": 6.6
-  },
-  "k15": true,
-  "k16": null,
-  "k17": [
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
-}`
	outputFormated := obj.Format(true, "-", "  ")
	if outputFormated != expectedFormated {
		t.Fatalf("object formated incorrectly:\n\texpected %s\n\tgot %s", expectedMinified, outputMinified)
	}
}
