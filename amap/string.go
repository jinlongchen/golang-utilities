package amap

import (
	"bytes"
	"errors"
	"github.com/jinlongchen/golang-utilities/json"
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
	if len(m) == 0 {
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
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	*m = []byte(str) //append((*m)[0:0], data...)
	return nil
}

func (m String) ToString() string {
	if m == nil {
		return ""
	}
	if bytes.Compare(sbArray, m) == 0 {
		return ""
	}
	if len(m) < 1 {
		return ""
	}
	var ret string
	err := json.Unmarshal(m, &ret)
	if err != nil {
		return ""
	}
	return ret
}
