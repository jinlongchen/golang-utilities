package crypto

import (
    "encoding/base64"
    "testing"
)

func TestAESEncryptCBC(t *testing.T) {
    key := `59b61e4d2f5eadde9680f78ed9db9a4a`
    println(len(key))
    // keyBytes, err := hex.DecodeString(key)
    keyBytes, err := base64.StdEncoding.DecodeString(key)
    if err != nil {
        panic(err)
    }
    encrypted, err := AESEncryptCBC([]byte("123"), keyBytes[:16], keyBytes[:16])
    if err != nil {
        panic(err)
    }
    println(base64.StdEncoding.EncodeToString(encrypted))
    // dBIdIyaxb3M+2CcI6vB/4Q==
    // BQpBiscFSLg7JJGfUoimdQ==
}
