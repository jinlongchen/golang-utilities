package converter

import (
	"fmt"
	"testing"
)

func TestConvertToMap(t *testing.T) {
	m := ConvertToMap("")
	fmt.Printf("%v", m)
}
func TestAsInt(t *testing.T) {
	i := AsInt("1025.0", 0)
	fmt.Printf("%v\n", i)
}