package rand

import (
	mrand "math/rand"
	"time"
	"fmt"
	"sync/atomic"
	"net"
)

var (
	letters  = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	r        = mrand.New(mrand.NewSource(time.Now().UnixNano()))
	seqNo   uint64
	address uint8
)

func init() {
	address = uint8(r.Intn(255))
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				ip4 := ipnet.IP.To4()
				if ip4 != nil {
					address = uint8([]byte(ip4)[3])
				}
			}
		}
	}
}

func GetNonceString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func GetShortTimestampRandString() string {
	timeStamp := time.Now().Unix() - 1466006400
	atomic.AddUint64(&seqNo, 1)

	resIdStr := fmt.Sprintf("%x%x%x%x",
		timeStamp,
		seqNo,
		address,
		0x12 + r.Int31n(0xFF),
	)

	return resIdStr
}
func GetNormalTimestampRandString() string {
	now := time.Now()
	atomic.AddUint64(&seqNo, 1)

	resIdStr := fmt.Sprintf("%02d%02d%02d%02d%02d%02d-%05d-%x%02X",
		now.Year() - 2000,
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
		seqNo,
		address,
		0x12 + r.Int31n(0xED),
	)

	return resIdStr
}

func GetRandInt64(min int64, max int64) int64 {
	return min + r.Int63n(max - min)
}
