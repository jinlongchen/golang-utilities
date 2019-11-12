package datetime

import (
	"github.com/jinlongchen/golang-utilities/log"
	"testing"
)

func TestParseWithTimeZone(t *testing.T) {
	ti, err := ParseWithTimeZone("2006-01-02 15:04", "2019-11-12 20:13:14", TimeZoneShanghai)
	if err != nil {
		log.Errorf("parse err: %s", err.Error())
	}
	log.Infof("time: %s", ti.String())
}