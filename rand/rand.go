package rand

import (
	mrand "math/rand"
	"time"
	"sync/atomic"
	"net"
	"bytes"
	"math/bits"
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

/*
2018-11-01 00:00:00: 18neo0
2018-11-02 00:00:00: 18p9c0
2018-11-03 00:00:00: 18r400
2018-11-04 00:00:00: 18syo0
2018-11-05 00:00:00: 18utc0
2018-11-06 00:00:00: 18wo00
2018-11-07 00:00:00: 18yio0
2018-11-08 00:00:00: 190dc0
2018-11-09 00:00:00: 192800
2018-11-10 00:00:00: 1942o0
2018-11-11 00:00:00: 195xc0
2018-11-12 00:00:00: 197s00
2018-11-13 00:00:00: 199mo0
2018-11-14 00:00:00: 19bhc0
2018-11-15 00:00:00: 19dc00
2018-11-16 00:00:00: 19f6o0
2018-11-17 00:00:00: 19h1c0
2018-11-18 00:00:00: 19iw00
2018-11-19 00:00:00: 19kqo0
2018-11-20 00:00:00: 19mlc0
2018-11-21 00:00:00: 19og00
2018-11-22 00:00:00: 19qao0
2018-11-23 00:00:00: 19s5c0
2018-11-24 00:00:00: 19u000
2018-11-25 00:00:00: 19vuo0
2018-11-26 00:00:00: 1bh9c0

2018-11-01 00:00:00: 18neo0
2018-11-01 01:00:00: 18nhg0
2018-11-01 02:00:00: 18nk80
2018-11-01 03:00:00: 18nn00
2018-11-01 04:00:00: 18nps0
2018-11-01 05:00:00: 18nsk0
2018-11-01 06:00:00: 18nvc0
2018-11-01 07:00:00: 18ny40
2018-11-01 08:00:00: 18o0w0
2018-11-01 09:00:00: 18o3o0
2018-11-01 10:00:00: 18o6g0
2018-11-01 11:00:00: 18o980
2018-11-01 12:00:00: 18oc00
2018-11-01 13:00:00: 18oes0
2018-11-01 14:00:00: 18ohk0
2018-11-01 15:00:00: 18okc0
2018-11-01 16:00:00: 18on40
2018-11-01 17:00:00: 18opw0
2018-11-01 18:00:00: 18oso0
2018-11-01 19:00:00: 18ovg0
2018-11-01 20:00:00: 18oy80
2018-11-01 21:00:00: 18p100
2018-11-01 22:00:00: 18p3s0
2018-11-01 23:00:00: 18p6k0
2018-11-01 24:00:00: 18p9c0
2018-11-01 25:00:00: 18pc40
2018-11-01 26:00:00: 18pew0
*/
func GetShortTimestampRandString() string {
	timeStamp := time.Now().Unix() - 1466006400
	sn := atomic.AddUint64(&seqNo, 1)

	//strconv.FormatInt(timeStamp, 36)
	buf := new(bytes.Buffer)
	write36N(buf, uint64(timeStamp))
	buf.WriteByte('-')
	write36N(buf, sn)
	write36N(buf, uint64(address))
	write36N(buf, uint64(r.Int31n(0xFF)))
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

const digits = "0123456789abcdefghijklmnopqrstuvwxyz"

func write36N(buffer *bytes.Buffer, u uint64) {
	base := 36

	var buf [64 + 1]byte
	i := len(buf)

	if isPowerOfTwo(base) {
		shift := uint(bits.TrailingZeros(uint(base))) & 31
		b := uint64(base)
		m := uint(base) - 1
		for u >= b {
			i--
			buf[i] = digits[uint(u)&m]
			u >>= shift
		}
		i--
		buf[i] = digits[uint(u)]
	} else {
		b := uint64(base)
		for u >= b {
			i--
			q := u / b
			buf[i] = digits[uint(u-q*b)]
			u = q
		}
		i--
		buf[i] = digits[uint(u)]
	}

	for k := i; k < 65; k++ {
		buffer.WriteByte(buf[k])
	}
}

func isPowerOfTwo(x int) bool {
	return x&(x-1) == 0
}
