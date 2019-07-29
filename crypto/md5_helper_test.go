package crypto

import (
	"testing"
)

func BenchmarkString_GetMd5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		String("").GetMd5()
	}
}

func TestString_GetMd5String(t *testing.T) {
	println(String("").GetMd5String())
	println(len(String("").GetMd5String()))
}
