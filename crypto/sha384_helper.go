package crypto

import (
    "encoding/hex"

    "golang.org/x/crypto/sha3"
)

func GetSha384Hash(s string) string {
    h := sha3.New384()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}
