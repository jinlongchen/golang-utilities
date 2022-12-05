package crypto

import (
    "encoding/base64"
    "encoding/hex"
    "testing"
)

func TestEncryptDes(t *testing.T) {
    data, err := DesEncrypt([]byte("123"), []byte("12345678"))
    if err != nil {
        t.Fatal(err)
    }
    // LDiFUdf0iew=
    t.Log(base64.StdEncoding.EncodeToString(data))
}

func TestDecryptDes(t *testing.T) {
    // data, err := base64.StdEncoding.DecodeString("4PYjz7uBgHM=")
    data, err := hex.DecodeString("a16c7426acc1572e5b37cf651ba02e113346f2702cfbaffda26fea17ebd4f94bc539805b5cf41995e479f6f2d8eefb7b9f8601e16246e70ab324759c918494726a281f54e06239a7769ea4a455635154ee462a6b1a35681234d2d645124a07d1726e5fd21a6a2aeeafd553ff9873f11b561f384a1d5d81388c10b09293bc90c53a2079088c8fe17ee612602a8095b8cb0d6ab7c718f6b421c0408a61a5bfb551a002b1c6feb0ee414121c7224812839afb10ba9f45e4cedef7e7cff189d64bb2620aa4ac6589a3f08026abaabb1c0d9127dcbb0c10750b7ba9b67c8be8472f8512d2ccc6694469bc585b9a042545e57c165e0bdf9fa99332")
    if err != nil {
        t.Fatal(err)
    }
    // data, _ = base64.StdEncoding.DecodeString("LDiFUdf0iew=")
    data, err = DesDecrypt(data, []byte("12345678"))
    if err != nil {
        t.Fatal(err)
    }
    t.Log(string(data))
}
