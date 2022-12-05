package crypto

import (
    "crypto/sha1"
    "encoding/hex"
)

func GetSha1Hash(s string) string {
    h := sha1.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}
