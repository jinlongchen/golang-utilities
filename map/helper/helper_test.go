package helper

import (
	"github.com/jinlongchen/golang-utilities/log"
	"testing"
	"encoding/json"
	"time"
)

func TestSetValue(t *testing.T) {
	filter := make(map[string]interface{})
	SetValue(filter, "applicant.phone", "123")
	data, _ := json.Marshal(filter)
	println(string(data))
	time.Sleep(time.Second)
}
type M0 struct {
	ByComplex string `json:"byComplex"`
}
type M1 struct {
	M00 map[string]interface{} `json:"byComplex"`
}
func TestGetValueAsString(t *testing.T) {
	m1 := &M1{
		M00: nil,
	}
	m := map[string]interface{}{
		"created_by": m1.M00,
	}
	log.DumpFormattedJson(m)
	v := GetValueAsString(m, "created_by.user.id", "")
	println(v == "")
}
