package text

import "testing"

func TestPrintf(t *testing.T) {
	Printf("{{ id:{0:d} }},123\n", 1)
}
