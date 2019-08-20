package datetime

import (
	"github.com/jinlongchen/golang-utilities/json"
	"strings"
	"time"
)

type UTCTime time.Time

func (dt UTCTime) MarshalJSON() ([]byte, error) {
	tmp := time.Time(dt)
	return json.ShouldMarshal(tmp.UTC().Format(time.RFC3339)), nil
}

func (dt *UTCTime) UnmarshalJSON(p []byte) error {
	t, err := time.Parse(time.RFC3339, strings.Replace(
		string(p),
		"\"",
		"",
		-1,
	))
	if err != nil {
		return err
	}
	*dt = UTCTime(t)

	return nil
}
