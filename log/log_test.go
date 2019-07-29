package log

import "testing"

func TestDumpKeyValue(t *testing.T) {
	type X struct {
		A string `json:"a"`
		B string `json:"b"`
	}

	DumpKeyValue(&X{
		A: "aaaa",
		B: "bbb",
	})
}
