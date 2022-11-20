package map_util

import (
    "testing"

    "github.com/jinlongchen/golang-utilities/json"
)

func TestGetValueAsMapSlice(t *testing.T) {
    m := make(map[string]interface{})

    json.Unmarshal([]byte(`{
    "a": [
{
 "b":1
},
{
 "c":2
}
]
`), &m)
    GetValueAsMapSlice(m, "a")
}
