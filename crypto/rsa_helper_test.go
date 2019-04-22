package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

func TestRSASign(t *testing.T) {
	rsaKey , _ := rsa.GenerateKey(rand.Reader, 2048)
	s, err := RSASign(rsaKey, []byte("123"))
	if err != nil {
		t.Fatal(err.Error())
	}
	err = RSAVerify(rsaKey.Public().(*rsa.PublicKey), []byte("1231"), s)
	if err != nil {
		t.Fatal(err.Error())
	}
}