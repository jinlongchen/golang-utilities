package datetime

import (
	"github.com/araddon/dateparse"
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
	t, err := dateparse.ParseAny(strings.Replace(
		string(p),
		"\"",
		"",
		-1,
	))
	if err != nil {
		t, err = time.Parse(time.RFC3339, strings.Replace(
			string(p),
			"\"",
			"",
			-1,
		))
		if err != nil {
			t, err = time.Parse("2006-01-02T15:04Z07:00", strings.Replace(
				string(p),
				"\"",
				"",
				-1,
			))
		}
	}
	if err != nil {
		return err
	}
	*dt = UTCTime(t)

	return nil
}
