package text

import "strings"

type String string

func (d String) Format(fmt string, a ...interface{}) []byte {
	return []byte{}
}

// ToUpper converts the string to uppercase.
func (d String) ToUpper() String {
	return String(strings.ToUpper(string(d)))
}
