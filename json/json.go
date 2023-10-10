package json

import (
	"bytes"
	"encoding/json"
)

func MarshalToString(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	return string(data), err
}

func MarshalToBuffer(v interface{}) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	data, err := json.Marshal(v)
	if err != nil {
		return buf, err
	}
	_, err = buf.Write(data)
	return buf, err
}

func UnmarshalFromString(str string, v interface{}) error {
	return json.Unmarshal([]byte(str), v)
}

func UnmarshalFromStringV2[T any](str string) (T, error) {
	return UnmarshalV2[T]([]byte(str))
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func ShouldMarshal(v interface{}) []byte {
	ret, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return ret
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func UnmarshalV2[T any](data []byte) (T, error) {
	var t T
	err := json.Unmarshal(data, &t)
	if err != nil {
		return t, err
	}
	return t, err
}
