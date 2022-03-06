package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetSha256Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
