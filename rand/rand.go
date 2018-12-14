package rand

import (
	"encoding/binary"
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
	address uint64
)

func init() {
	address = uint64(getIpAddr())
	//address = getMacAddr()
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
	write62N(buf, uint64(timeStamp))
	write62N(buf, uint64(address))
	buf.WriteByte('0')
	write62N(buf, sn)
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

const digits = "_123456789abcdefghijklmnopqrstuvwxyz"

func write62N(buffer *bytes.Buffer, u uint64) {
	base := 36

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

func getMacAddr() (addr uint64) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				b := i.HardwareAddr
				addr = uint64(b[5])<<16 | uint64(b[4])<<24 |
					uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
				break
			}
		}
	}
	return
}
func getIpAddr() (addr uint32) {
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				println(ipnet.IP.String())
				ip4 := ipnet.IP.To4()
				if ip4 != nil {
					addr = binary.BigEndian.Uint32(ip4)
				}
			}
		}
	}
	return
}