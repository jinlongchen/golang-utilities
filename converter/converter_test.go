package converter

import (
	"fmt"
	"testing"
)

func TestConvertToMap(t *testing.T) {
	m := ConvertToMap("")
	fmt.Printf("%v", m)
}