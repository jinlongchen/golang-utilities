package crypto

import (
	"crypto/cipher"
	"crypto/aes"
	"bytes"
	"fmt"
	"errors"
)

var (
	EmptyIV = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
)

func AESEncryptCBC(src []byte, key []byte, iv []byte) (dst []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	src, _ = Pkcs7Pad(src, aes.BlockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	dst = make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return
}

func AESDecryptCBC(src []byte, key []byte, iv []byte) (dst []byte, err error) {
	if len(src)%aes.BlockSize != 0 {
		return nil, errors.New("src is not correct")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	dst = make([]byte, len(src))

	blockMode.CryptBlocks(dst, src)
	//log.Println("dst:", dst)

	dst, err = Pkcs7Unpad(dst, block.BlockSize())

	return
}

func Pkcs7Pad(data []byte, blocklen int) ([]byte, error) {
	if blocklen <= 0 {
		return nil, fmt.Errorf("invalid blocklen %d", blocklen)
	}
	padlen := 1
	for ((len(data) + padlen) % blocklen) != 0 {
		padlen = padlen + 1
	}

	pad := bytes.Repeat([]byte{byte(padlen)}, padlen)
	return append(data, pad...), nil
}

func Pkcs7Unpad(data []byte, blocklen int) ([]byte, error) {
	if blocklen <= 0 {
		return nil, fmt.Errorf("invalid blocklen %d", blocklen)
	}
	if len(data)%blocklen != 0 || len(data) == 0 {
		return nil, fmt.Errorf("invalid data len %d", len(data))
	}
	padlen := int(data[len(data)-1])
	if padlen > blocklen || padlen == 0 {
		return nil, fmt.Errorf("invalid padding")
	}
	// check padding
	pad := data[len(data)-padlen:]
	for i := 0; i < padlen; i++ {
		if pad[i] != byte(padlen) {
			return nil, fmt.Errorf("invalid padding")
		}
	}

	return data[:len(data)-padlen], nil
}

func Pkcs5Pad(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func Pkcs5Unpad(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
