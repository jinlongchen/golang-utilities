package crypto

import (
	"testing"
)

func BenchmarkString_GetMd5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		String("").GetMd5()
	}
}

