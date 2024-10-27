package text

import (
	"math"
	"strings"
)

// Mask masks the input string based on its length.
func Mask(str string) string {
	l := len(str)
	if l == 0 {
		return ""
	} else if l < 4 {
		return "***"
	}

	s := int(math.Sqrt(float64(l)))
	if s < 3 {
		s = 3
	}
	prelen := len(str) / s
	postlen := (len(str) - prelen) / (s - 1)
	masklen := len(str) - prelen - postlen

	for masklen < 3 {
		postlen--
		masklen = len(str) - prelen - postlen
		if masklen >= 3 {
			break
		}
		prelen--
		masklen = len(str) - prelen - postlen
	}

	return str[:prelen] + strings.Repeat("*", masklen) + str[len(str)-postlen:]
}
