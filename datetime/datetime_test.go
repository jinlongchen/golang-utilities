package datetime

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestUTCTime_MarshalJSON(t *testing.T) {
	dt := UTCTime(time.Now())
	j, err := json.Marshal(dt)
	if err != nil {
		panic(err)
	}
	println(string(j))
}
func TestUTCTime_UnmarshalJSON(t *testing.T) {
	x := []byte(`"2019-08-20T07:52:03.341Z"`)
	dt := UTCTime{}
	err := json.Unmarshal(x, &dt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", time.Time(dt).Local())
}
