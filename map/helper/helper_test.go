package helper

import (
	"github.com/jinlongchen/golang-utilities/json"
	"github.com/jinlongchen/golang-utilities/log"

	"testing"
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
func TestExists(t *testing.T) {
	data := []byte(`{"id":"","member":{"id":"e8332630ae4a9c0ddc4d6ac0bc6a53d8","type":""},"type":"","title":"","message":"","amount":0,"balance":null,"payment_provider":"","time":0,"created_on":0,"created_by":""}`)
	filter := make(map[string]interface{})
	err := json.Unmarshal(data, &filter)
	if err != nil {
		panic(err)
	}
	type X struct {
		Y string `json:"y"`
	}
	filter["abc"] = []X {X{
		Y: "y",
	}}
	println(string(json.ShouldMarshal(X{
		Y: "y",
	})))
	exists3 := Exists(filter, "abc.y")
	exists := Exists(filter, "member.kk")
	exists2 := Exists(filter, "amount.range.min")
	println(exists, exists2, exists3)
}
