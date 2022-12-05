package number

import (
    "testing"

    "github.com/jinlongchen/golang-utilities/json"
)

func TestFloat64p2_MarshalJSON(t *testing.T) {
    var a Float64p2
    a = 123.456789
    println(string(json.ShouldMarshal(a)))
}
