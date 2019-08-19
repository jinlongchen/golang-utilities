package amap

import (
	"bytes"
	"errors"
)

var (
	sbArray = []byte(`[]`)
)
type String []byte

// MarshalJSON returns m as the JSON encoding of m.
func (m String) MarshalJSON() ([]byte, error) {
	println("MarshalJSON:", string(m))
	if m == nil {
		return []byte(`""`), nil
	}
	if bytes.Compare(sbArray, m) == 0 {
		return []byte(`""`), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *String) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("amap.String: UnmarshalJSON on nil pointer")
	}
	if bytes.Compare(sbArray, data) == 0 {
		*m = []byte("")
		return nil
	}
	*m = append((*m)[0:0], data...)
	return nil
}

func (m String) ToString() string {
	if m == nil {
		return ""
	}
	if bytes.Compare(sbArray, m) == 0 {
		return ""
	}
	return string(m)
}
