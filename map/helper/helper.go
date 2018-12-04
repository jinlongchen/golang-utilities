package helper

import (
	"strings"
	"github.com/jinlongchen/golang-utilities/converter"
)

func SetValue(m map[string]interface{}, path string, val interface{}) {
	if path == "" {
		return
	}

	index := strings.Index(path, ".")
	if index > 0 {
		child := m[path[:index]]
		if child == nil {
			return
		}
		switch child.(type) {
		case map[string]interface{}:
			SetValue(child.(map[string]interface{}), path[index+1:], val)
		default:
			return
		}
	} else {
		m[path] = val
	}
}
func GetValue(m map[string]interface{}, path string) interface{} {
	if path == "" {
		return m
	}

	index := strings.Index(path, ".")
	if index > 0 {
		child := m[path[:index]]
		if child == nil {
			return nil
		}
		switch child.(type) {
		case map[string]interface{}:
			return GetValue(child.(map[string]interface{}), path[index+1:])
		default:
			return nil
		}
	} else {
		return m[path]
	}
}

func GetValueAsInt(m map[string]interface{}, path string, defaultValue int) int {
	if path == "" {
		return converter.AsInt(m, defaultValue)
	}

	index := strings.Index(path, ".")
	if index > 0 {
		child := m[path[:index]]
		if child == nil {
			return defaultValue
		}
		switch child.(type) {
		case map[string]interface{}:
			return GetValueAsInt(child.(map[string]interface{}), path[index+1:], defaultValue)
		default:
			return defaultValue
		}
	} else {
		return converter.AsInt(m[path], defaultValue)
	}
}
func GetValueAsInt32(m map[string]interface{}, path string, defaultValue int32) int32 {
	if path == "" {
		return converter.AsInt32(m, defaultValue)
	}

	index := strings.Index(path, ".")
	if index > 0 {
		child := m[path[:index]]
		if child == nil {
			return defaultValue
		}
		switch child.(type) {
		case map[string]interface{}:
			return GetValueAsInt32(child.(map[string]interface{}), path[index+1:], defaultValue)
		default:
			return defaultValue
		}
	} else {
		return converter.AsInt32(m[path], defaultValue)
	}
}
func GetValueAsInt64(m map[string]interface{}, path string, defaultValue int64) int64 {
	if path == "" {
		return converter.AsInt64(m, defaultValue)
	}

	index := strings.Index(path, ".")
	if index > 0 {
		child := m[path[:index]]
		if child == nil {
			return defaultValue
		}
		switch child.(type) {
		case map[string]interface{}:
			return GetValueAsInt64(child.(map[string]interface{}), path[index+1:], defaultValue)
		default:
			return defaultValue
		}
	} else {
		return converter.AsInt64(m[path], defaultValue)
	}
}
func GetValueAsString(m map[string]interface{}, path string, defaultValue string) string {
	if m == nil {
		return defaultValue
	}
	if path == "" {
		return converter.AsString(m, defaultValue)
	}

	index := strings.Index(path, ".")
	if index > 0 {
		child := m[path[:index]]
		if child == nil {
			return defaultValue
		}
		switch child.(type) {
		case map[string]interface{}:
			return GetValueAsString(child.(map[string]interface{}), path[index+1:], defaultValue)
		default:
			return converter.AsString(child, defaultValue)
		}
	} else {
		return converter.AsString(m[path], defaultValue)
	}
}
func GetValueAsBool(m map[string]interface{}, path string, defaultValue bool) bool {
	if path == "" {
		return converter.AsBool(m, defaultValue)
	}

	index := strings.Index(path, ".")
	if index > 0 {
		child := m[path[:index]]
		if child == nil {
			return defaultValue
		}
		switch child.(type) {
		case map[string]interface{}:
			return GetValueAsBool(child.(map[string]interface{}), path[index+1:], defaultValue)
		default:
			return converter.AsBool(child, defaultValue)
		}
	} else {
		return converter.AsBool(m[path], defaultValue)
	}
}
func GetValueAsFloat64(m map[string]interface{}, path string, defaultValue float64) float64 {
	if path == "" {
		return converter.AsFloat64(m, defaultValue)
	}

	index := strings.Index(path, ".")
	if index > 0 {
		child := m[path[:index]]
		if child == nil {
			return defaultValue
		}
		switch child.(type) {
		case map[string]interface{}:
			return GetValueAsFloat64(child.(map[string]interface{}), path[index+1:], defaultValue)
		default:
			return defaultValue
		}
	} else {
		return converter.AsFloat64(m[path], defaultValue)
	}
}
func GetValueAsMap(m map[string]interface{}, path string, defaultValue map[string]interface{}) map[string]interface{} {
	if path == "" {
		return m
	}

	index := strings.Index(path, ".")
	if index > 0 {
		child := m[path[:index]]
		if child == nil {
			return defaultValue
		}
		switch child.(type) {
		case map[string]interface{}:
			return GetValueAsMap(child.(map[string]interface{}), path[index+1:], defaultValue)
		default:
			return defaultValue
		}
	} else {
		child := m[path]
		switch child.(type) {
		case map[string]interface{}:
			return child.(map[string]interface{})
		default:
			return defaultValue
		}
	}
}
