package rand

import (
	"bytes"
	"encoding/binary"
	"net"
	"testing"
	"time"
)

func TestGetRandInt(t *testing.T) {
	for i := 0; i < 10; i++ {
		println(GetRandInt(0, 2))
	}
}
func TestGetShortTimestampSequenceNo2(t *testing.T) {
	println(GetShortTimestampSequenceNo())
	println(GetShortTimestampSequenceNo())
	println(GetShortTimestampSequenceNo())
	println(GetShortTimestampSequenceNo())
	println(GetShortTimestampSequenceNo())
}
func TestGetNormalTimestampRandString(t *testing.T) {
	t.Log(GetNormalTimestampRandString())
}
func TestParseShortTimestampRandString(t *testing.T) {
	//2m2hs502kyv66909
	epochTime, _, _ := ParseShortTimestampRandString("31wwxy02kyqil404")
	println(epochTime)
	println(time.Unix(epochTime, 0).String())
}
func TestGetShortTimestampRandString(t *testing.T) {
	time2 := time.Date(2019, 2, 23, 21, 32, 14, 0, time.Local)
	randStr2 := GetRandStringWithTimestamp(time2.Unix())
	t.Log("randStr2:", randStr2)

	time0 := time.Date(2085, 6, 8, 13, 45, 35, 0, time.Local)
	randStr := GetRandStringWithTimestamp(time0.Unix())
	t.Log(randStr)

	epochTime, _, _ := ParseShortTimestampRandString(randStr)
	t.Log(epochTime)

	time1 := time.Unix(epochTime, 0)

	if time0.Year() != time1.Year() &&
		time0.Month() != time1.Month() &&
		time0.Day() != time1.Day() &&
		time0.Hour() != time1.Hour() &&
		time0.Minute() != time1.Minute() &&
		time0.Second() != time1.Second() {
		t.Fatal("not eq:", time1.String())
	}
}
func TestGetRandStringWithTimestamp(t *testing.T) {
	t.Log(GetRandStringWithTimestamp(time.Now().Unix()))
	t.Log(GetRandStringWithTimestamp(time.Date(2085, 6, 8, 13, 45, 35, 0, time.Local).Unix()))
}
func TestGetShortTimestampRandString2(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(GetShortTimestampRandString())
		//5kxSz18pHe014
		//1awhw59zovg014
	}
}
func TestWrite62N(t *testing.T) {
	buf := new(bytes.Buffer)
	//timeStamp1 := time.Date(2085, 6, 8, 13,45,35,0,time.Local).Unix()- 1466035200
	timeStamp1 := time.Date(2074, 9, 16, 13, 20, 24, 0, time.Local).Unix() - 1466035200
	write36N(buf, uint64(timeStamp1))
	//write36N(buf, uint64(255))
	t.Log(buf.String())
}
func BenchmarkGetShortTimestampRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetShortTimestampRandString()
	}
}
func TestIpAddr(t *testing.T) {
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				ip4 := ipnet.IP.To4()
				if ip4 != nil {
					t.Log(a.String())
					addr := uint64(binary.BigEndian.Uint32(ip4))
					t.Log(ip4.String()) //0xAC 14 0A 03
					t.Log(addr)
				}
			}
		}
	}
}
