package rand

import (
	"testing"
	"time"
	"strconv"
	"bytes"
)

func TestGetNormalTimestampRandString(t *testing.T) {
	t.Log(GetNormalTimestampRandString())
}
func TestGetShortTimestampRandString(t *testing.T) {
	timeStamp1 := time.Date(2018, 11, 1, 1,0,0,0,time.Local).Unix()- 1466035200
	timeStamp2 := time.Date(2018, 11, 1, 2,0,0,0,time.Local).Unix()- 1466035200
	timeStamp3 := time.Date(2018, 11, 1, 3,0,0,0,time.Local).Unix()- 1466035200
	timeStamp4 := time.Date(2018, 11, 1, 4,0,0,0,time.Local).Unix()- 1466035200
	timeStamp5 := time.Date(2018, 11, 1, 5,0,0,0,time.Local).Unix()- 1466035200
	timeStamp6 := time.Date(2018, 11, 1, 6,0,0,0,time.Local).Unix()- 1466035200
	timeStamp7 := time.Date(2018, 11, 1, 7,0,0,0,time.Local).Unix()- 1466035200
	timeStamp8 := time.Date(2018, 11, 1, 8,0,0,0,time.Local).Unix()- 1466035200
	timeStamp9 := time.Date(2018, 11, 1, 9,0,0,0,time.Local).Unix()- 1466035200
	timeStamp10 := time.Date(2018, 11, 1, 10,0,0,0,time.Local).Unix()- 1466035200
	timeStamp11 := time.Date(2018, 11, 1, 11,0,0,0,time.Local).Unix()- 1466035200
	timeStamp12 := time.Date(2018, 11, 1, 12,0,0,0,time.Local).Unix()- 1466035200
	timeStamp13 := time.Date(2018, 11, 1, 13,0,0,0,time.Local).Unix()- 1466035200
	timeStamp14 := time.Date(2018, 11, 1, 14,0,0,0,time.Local).Unix()- 1466035200
	timeStamp15 := time.Date(2018, 11, 1, 15,0,0,0,time.Local).Unix()- 1466035200
	timeStamp16 := time.Date(2018, 11, 1, 16,0,0,0,time.Local).Unix()- 1466035200
	timeStamp17 := time.Date(2018, 11, 1, 17,0,0,0,time.Local).Unix()- 1466035200
	timeStamp18 := time.Date(2018, 11, 1, 18,0,0,0,time.Local).Unix()- 1466035200
	timeStamp19 := time.Date(2018, 11, 1, 19,0,0,0,time.Local).Unix()- 1466035200
	timeStamp20 := time.Date(2018, 11, 1, 20,0,0,0,time.Local).Unix()- 1466035200
	timeStamp21 := time.Date(2018, 11, 1, 21,0,0,0,time.Local).Unix()- 1466035200
	timeStamp22 := time.Date(2018, 11, 1, 22,0,0,0,time.Local).Unix()- 1466035200
	timeStamp23 := time.Date(2018, 11, 1, 23,0,0,0,time.Local).Unix()- 1466035200
	timeStamp24 := time.Date(2018, 11, 1, 0,0,0,0,time.Local).Unix()- 1466035200
	timeStamp25 := time.Date(2018, 11, 1, 0,0,0,0,time.Local).Unix()- 1466035200
	timeStamp26 := time.Date(2018, 11, 1, 0,0,0,0,time.Local).Unix()- 1466035200
	timeStamp27 := time.Date(2085, 6, 8, 13,45,35,0,time.Local).Unix()- 1466035200

	println("timeStamp1:", strconv.FormatInt(timeStamp1, 36))
	println("timeStamp2:", strconv.FormatInt(timeStamp2, 36))
	println("timeStamp3:", strconv.FormatInt(timeStamp3, 36))
	println("timeStamp4:", strconv.FormatInt(timeStamp4, 36))
	println("timeStamp5:", strconv.FormatInt(timeStamp5, 36))
	println("timeStamp6:", strconv.FormatInt(timeStamp6, 36))
	println("timeStamp7:", strconv.FormatInt(timeStamp7, 36))
	println("timeStamp8:", strconv.FormatInt(timeStamp8, 36))
	println("timeStamp9:", strconv.FormatInt(timeStamp9, 36))
	println("timeStamp10:", strconv.FormatInt(timeStamp10, 36))
	println("timeStamp11:", strconv.FormatInt(timeStamp11, 36))
	println("timeStamp12:", strconv.FormatInt(timeStamp12, 36))
	println("timeStamp13:", strconv.FormatInt(timeStamp13, 36))
	println("timeStamp14:", strconv.FormatInt(timeStamp14, 36))
	println("timeStamp15:", strconv.FormatInt(timeStamp15, 36))
	println("timeStamp16:", strconv.FormatInt(timeStamp16, 36))
	println("timeStamp17:", strconv.FormatInt(timeStamp17, 36))
	println("timeStamp18:", strconv.FormatInt(timeStamp18, 36))
	println("timeStamp19:", strconv.FormatInt(timeStamp19, 36))
	println("timeStamp20:", strconv.FormatInt(timeStamp20, 36))
	println("timeStamp21:", strconv.FormatInt(timeStamp21, 36))
	println("timeStamp22:", strconv.FormatInt(timeStamp22, 36))
	println("timeStamp23:", strconv.FormatInt(timeStamp23, 36))
	println("timeStamp24:", strconv.FormatInt(timeStamp24, 36))
	println("timeStamp25:", strconv.FormatInt(timeStamp25, 36))
	println("timeStamp26:", strconv.FormatInt(timeStamp26, 36))
	println("timeStamp27:", strconv.FormatInt(timeStamp27, 36))
	//	"zzzzzz" := time.Date(2085, 6, 8, 13,45,35,0,time.Local).Unix()- 1466035200

	t.Log(GetShortTimestampRandString())
}
func TestGetShortTimestampRandString2(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(GetShortTimestampRandString())
	}
}
func TestWrite62N(t *testing.T) {
	buf := new(bytes.Buffer)
	timeStamp1 := time.Date(2018, 11, 1, 1, 0, 0, 0, time.Local).Unix() - 1466006400
	write62N(buf, uint64(timeStamp1))   //62:54GC4
	t.Log(buf.String())

}
func BenchmarkGetShortTimestampRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetShortTimestampRandString()
	}
}
