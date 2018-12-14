package rand

import (
	mrand "math/rand"
	"time"
	"sync/atomic"
	"net"
	"bytes"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	r       = mrand.New(mrand.NewSource(time.Now().UnixNano()))
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
	timeStamp := time.Now().Unix() - 1466035200
	sn := atomic.AddUint64(&seqNo, 1)

	buf := new(bytes.Buffer)
	write35N(buf, uint64(timeStamp))
	write35N(buf, uint64(address))
	buf.WriteByte('0')
	write35N(buf, sn)
	return buf.String()
}

func GetNormalTimestampRandString() string {
	now := time.Now()
	year, month, day := now.Date()
	hour, min, sec := now.Clock()

	sn := atomic.AddUint64(&seqNo, 1)

	buf := new(bytes.Buffer)

	writeN(buf, uint64(year%100), 2)
	writeN(buf, uint64(int(month)), 2)
	writeN(buf, uint64(day), 2)
	writeN(buf, uint64(hour), 2)
	writeN(buf, uint64(min), 2)
	writeN(buf, uint64(sec), 2)
	buf.WriteByte('-')
	writeN(buf, sn, 2)
	buf.WriteByte('-')
	writeN(buf, uint64(address), 2)
	writeN(buf, uint64(r.Intn(0xFF)), 1)

	return buf.String()
}

func GetRandInt64(min int64, max int64) int64 {
	return min + r.Int63n(max-min)
}
func writeN(buffer *bytes.Buffer, x uint64, width int) {
	u := uint(x)
	if x < 0 {
		buffer.WriteByte('-')
		u = uint(-x)
	}

	var buf [20]byte
	i := len(buf)
	for u >= 10 {
		i--
		q := u / 10
		buf[i] = byte('0' + u - q*10)
		u = q
	}
	i--
	buf[i] = byte('0' + u)

	for w := len(buf) - i; w < width; w++ {
		buffer.WriteByte('0')
	}

	for k := i; k < 20; k++ {
		buffer.WriteByte(buf[k])
	}
}

const digits = "123456789abcdefghijklmnopqrstuvwxyz"

func write35N(buffer *bytes.Buffer, u uint64) {
	base := 35

	var buf [64 + 1]byte
	i := len(buf)

	b := uint64(base)
	for u >= b {
		i--
		q := u / b
		buf[i] = digits[uint(u-q*b)]
		u = q
	}
	i--
	buf[i] = digits[uint(u)]

	for k := i; k < 65; k++ {
		buffer.WriteByte(buf[k])
	}
}
