package crypto

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/jinlongchen/golang-utilities/errors"
)
var (
	ErrKeyMustBePEMEncoded = errors.New("Invalid Key: Key must be PEM encoded PKCS1 or PKCS8 private key")
	ErrNotRSAPrivateKey    = errors.New("Key is not a valid RSA private key")
)

func RSAEncrypt(keyPem []byte, data []byte) ([]byte, error) {
	var err error
	key, err := ParseRSAPublicKeyFromPEM(keyPem)
	if err != nil {
		return nil, err
	}
	klen := key.N.BitLen()/8 - 11
	if len(data) <= klen {
		var bb []byte
		if bb, err = rsa.EncryptPKCS1v15(rand.Reader, key, data); err != nil {
			return nil, err
		}
		return bb, nil
	}
	var buf bytes.Buffer
	var bb []byte
	for i, w, r := 0, 0, len(data); r > 0; i, r = i+w, r-w {
		if r <= klen {
			if bb, err = rsa.EncryptPKCS1v15(rand.Reader, key, data[i:]); err != nil {
				return nil, err
			}
			buf.Write(bb)
			w = r
		} else {
			if bb, err = rsa.EncryptPKCS1v15(rand.Reader, key, data[i:i+klen]); err != nil {
				return nil, err
			}
			buf.Write(bb)
			w = klen
		}
	}
	return buf.Bytes(), nil
}
func RSADecrypt(keyPem []byte, data []byte) ([]byte, error) {
	var err error

	key, err := ParseRSAPrivateKeyFromPEM(keyPem)
	if err != nil {
		return nil, err
	}

	klen := key.N.BitLen() / 8
	if len(data) <= klen {
		var bb []byte
		if bb, err = rsa.DecryptPKCS1v15(rand.Reader, key, data); err != nil {
			return nil, err
		}
		return bb, nil
	}
	var buf bytes.Buffer
	var bb []byte
	for i, w, r := 0, 0, len(data); r > 0; i, r = i+w, r-w {
		if r <= klen {
			if bb, err = rsa.DecryptPKCS1v15(rand.Reader, key, data[i:]); err != nil {
				return nil, err
			}
			buf.Write(bb)
			w = r
		} else {
			if bb, err = rsa.DecryptPKCS1v15(rand.Reader, key, data[i:i+klen]); err != nil {
				return nil, err
			}
			buf.Write(bb)
			w = klen
		}
	}
	return buf.Bytes(), nil
}

func ParseRSAPrivateKeyFromPEM(key []byte) (*rsa.PrivateKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, ErrKeyMustBePEMEncoded
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
			return nil, err
		}
	}

	var pkey *rsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PrivateKey); !ok {
		return nil, ErrNotRSAPrivateKey
	}

	return pkey, nil
}

func ParseRSAPublicKeyFromPEM(key []byte) (*rsa.PublicKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, ErrKeyMustBePEMEncoded
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
			parsedKey = cert.PublicKey
		} else {
			return nil, err
		}
	}

	var pkey *rsa.PublicKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PublicKey); !ok {
		return nil, ErrNotRSAPrivateKey
	}

	return pkey, nil
}
