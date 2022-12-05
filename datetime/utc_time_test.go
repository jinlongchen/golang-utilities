package datetime

import (
    "fmt"
    "testing"
    "time"

    "github.com/jinlongchen/golang-utilities/json"
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
func TestUTCTime_BeforeEpoch(t *testing.T) {
    t1 := UTCTime(time.Date(1, 1, 1, 0, 0, 0, 0, time.Local))
    fmt.Println(t1.BeforeEpoch())
}

func TestUTCTime_AfterEpoch(t *testing.T) {
    t1 := UTCTime(time.Date(1, 1, 1, 0, 0, 0, 0, time.Local))
    fmt.Println(t1.AfterEpoch())
    fmt.Println(time.Unix(0, 0).String())
}
