package converter

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jinlongchen/golang-utilities/json"
)

func AsInt(v interface{}, defaultValue int) int {
	if v != nil {
		switch v.(type) {
		case int:
			return v.(int)
		case int8:
			return int(v.(int8))
		case int16:
			return int(v.(int16))
		case int32:
			return int(v.(int32))
		case int64:
			return int(v.(int64))

		case uint:
			return int(v.(uint))
		case uint8:
			return int(v.(uint8))
		case uint16:
			return int(v.(uint16))
		case uint32:
			return int(v.(uint32))
		case uint64:
			return int(v.(uint64))

		case float64:
			return int(v.(float64))
		case float32:
			return int(v.(float32))

		case string:
			p, err := strconv.ParseInt(v.(string), 10, 32)
			if err != nil {
				return defaultValue
			}
			return int(p)
		default:
			return defaultValue
		}
	}
	return defaultValue
}
func AsInt32(v interface{}, defaultValue int32) int32 {
	if v != nil {
		switch v.(type) {
		case int:
			return int32(v.(int))
		case int8:
			return int32(v.(int8))
		case int16:
			return int32(v.(int16))
		case int32:
			return v.(int32)
		case int64:
			return int32(v.(int64))

		case uint:
			return int32(v.(uint))
		case uint8:
			return int32(v.(uint8))
		case uint16:
			return int32(v.(uint16))
		case uint32:
			return int32(v.(uint32))
		case uint64:
			return int32(v.(uint64))

		case float32:
			return int32(v.(float32))
		case float64:
			return int32(v.(float64))

		default:
			return defaultValue
		}
	}
	return defaultValue
}
func AsUInt32(v interface{}, defaultValue uint32) uint32 {
	if v != nil {
		switch v.(type) {
		case int:
			return uint32(v.(int))
		case int8:
			return uint32(v.(int8))
		case int16:
			return uint32(v.(int16))
		case int32:
			return uint32(v.(int32))
		case int64:
			return uint32(v.(int64))

		case uint:
			return uint32(v.(uint))
		case uint8:
			return uint32(v.(uint8))
		case uint16:
			return uint32(v.(uint16))
		case uint32:
			return v.(uint32)
		case uint64:
			return uint32(v.(uint64))

		case float32:
			return uint32(v.(float32))
		case float64:
			return uint32(v.(float64))

		default:
			return defaultValue
		}
	}
	return defaultValue
}
func AsInt64(v interface{}, defaultValue int64) int64 {
	if v != nil {
		switch v.(type) {
		case int:
			return int64(v.(int))
		case int8:
			return int64(v.(int8))
		case int16:
			return int64(v.(int16))
		case int32:
			return int64(v.(int32))
		case int64:
			return v.(int64)

		case uint:
			return int64(v.(uint))
		case uint8:
			return int64(v.(uint8))
		case uint16:
			return int64(v.(uint16))
		case uint32:
			return int64(v.(uint32))
		case uint64:
			return int64(v.(uint64))

		case float32:
			return int64(v.(float32))
		case float64:
			return int64(v.(float64))

		case string:
			ret, err := strconv.ParseInt(v.(string), 10, 64)
			if err != nil {
				return defaultValue
			}
			return ret
		default:
			return defaultValue
		}
	}
	return defaultValue
}
func AsFloat64(v interface{}, defaultValue float64) float64 {
	if v != nil {
		switch v.(type) {
		case int:
			return float64(v.(int))
		case int8:
			return float64(v.(int8))
		case int16:
			return float64(v.(int16))
		case int32:
			return float64(v.(int32))
		case int64:
			return float64(v.(int64))

		case uint:
			return float64(v.(uint))
		case uint8:
			return float64(v.(uint8))
		case uint16:
			return float64(v.(uint16))
		case uint32:
			return float64(v.(uint32))
		case uint64:
			return float64(v.(uint64))

		case float64:
			return v.(float64)
		case float32:
			return float64(v.(float32))

		case string:
			ret, err := strconv.ParseFloat(strings.Trim(v.(string), " \r\n"), 64)
			if err != nil {
				return defaultValue
			}
			return ret
		default:
			return defaultValue
		}
	}
	return defaultValue
}
func AsString(v interface{}, defaultValue string) string {
	if v != nil {
		switch v.(type) {
		case string:
			return v.(string)
		default:
			return fmt.Sprintf("%v", v)
		}
	}
	return defaultValue
}
func AsArray(v interface{}) []interface{} {
	if v != nil {
		switch v.(type) {
		case []interface{}:
			return v.([]interface{})
		default:
			return nil
		}
	}
	return nil
}
func AsStringSlice(v interface{}, defaultValue []string) []string {
	defer func() {
		if err := recover(); err != nil {
			println("[AsStringSlice]panic :%v", err)
		}
	}()
	if v != nil {
		switch v.(type) {
		case []string:
			return v.([]string)
		case []interface{}:
			x := v.([]interface{})
			r := make([]string, len(x))
			for key, value := range x {
				r[key] = AsString(value, "")
			}
			return r
		default:
			x := make([]interface{}, 0)
			x = ConvertToSlice(v, x)
			if x != nil {
				r := make([]string, len(x))
				for key, value := range x {
					r[key] = AsString(value, "")
				}
				return r
			}
			return defaultValue
		}
	}
	return defaultValue
}
func AsBool(v interface{}, defaultValue bool) bool {
	if v != nil {
		switch v.(type) {
		case bool:
			return v.(bool)
		default:
			return defaultValue
		}
	}
	return defaultValue
}
func AsMap(v interface{}) map[string]interface{} {
	if v != nil {
		switch v.(type) {
		case map[string]interface{}:
			return v.(map[string]interface{})
		default:
			return nil
		}
	}
	return nil
}
func AsMapSlice(v interface{}) []map[string]interface{} {
	if v != nil {
		switch v.(type) {
		case []interface{}:
			s := v.([]interface{})
			ret := make([]map[string]interface{}, len(s))
			for index, item := range s {
				if m, ok := item.(map[string]interface{}); ok {
					ret[index] = m
				} else {
					ret[index] = nil
				}
			}
			return ret
		case []map[string]interface{}:
			return v.([]map[string]interface{})
		default:
			return nil
		}
	}
	return nil
}
func AsDuration(v interface{}, defaultValue time.Duration) time.Duration {
	if v != nil {
		switch v.(type) {
		case int:
			return time.Duration(v.(int))
		case int8:
			return time.Duration(v.(int8))
		case int16:
			return time.Duration(v.(int16))
		case int32:
			return time.Duration(v.(int32))
		case int64:
			return time.Duration(v.(int64))
		case time.Duration:
			return v.(time.Duration)

		case uint:
			return time.Duration(v.(uint))
		case uint8:
			return time.Duration(v.(uint8))
		case uint16:
			return time.Duration(v.(uint16))
		case uint32:
			return time.Duration(v.(uint32))
		case uint64:
			return time.Duration(v.(uint64))

		case float32:
			return time.Duration(v.(float32))
		case float64:
			return time.Duration(v.(float64))

		default:
			return defaultValue
		}
	}
	return defaultValue
}

func ConvertToMap(s interface{}) map[string]interface{} {
	if s == nil {
		return nil
	}
	switch s.(type) {
	case map[string]interface{}:
		return s.(map[string]interface{})
	}
	data, err := json.Marshal(s)
	if err != nil {
		return nil
	}
	ret := map[string]interface{}{}
	err = json.Unmarshal(data, &ret)

	if err != nil {
		return nil
	}
	return ret
}

func ConvertToInterface(m interface{}, template interface{}) interface{} {
	data, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(data, template)
	if err != nil {
		return nil
	}
	return template
}

func ConvertToSlice(m interface{}, template []interface{}) []interface{} {
	data, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(data, &template)
	if err != nil {
		return nil
	}
	return template
}

func MustAsInt(v interface{}) int {
	if v != nil {
		switch v.(type) {
		case int:
			return v.(int)
		case int8:
			return int(v.(int8))
		case int16:
			return int(v.(int16))
		case int32:
			return int(v.(int32))
		case int64:
			return int(v.(int64))

		case uint:
			return int(v.(uint))
		case uint8:
			return int(v.(uint8))
		case uint16:
			return int(v.(uint16))
		case uint32:
			return int(v.(uint32))
		case uint64:
			return int(v.(uint64))

		case float64:
			return int(v.(float64))
		case float32:
			return int(v.(float32))
		default:
			panic("not int")
		}
	}
	panic("not int")
}
