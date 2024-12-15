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

// ToJSON converts an object to its JSON string representation.
func ToJSON(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// FromJSON parses a JSON string into an object.
func FromJSON(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}

// IsValidJSON checks if a string is a valid JSON.
func IsValidJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}
