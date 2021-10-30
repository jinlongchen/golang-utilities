package aliapi

import (
	"github.com/jinlongchen/golang-utilities/config"
	"io/ioutil"
	"path"
	"runtime"
	"testing"
)

func TestOCRIDCard(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	data, err := ioutil.ReadFile(path.Join(path.Dir(filename), "idcard_sample.jpeg"))
	if err != nil {
		t.Fatal(err)
	}

	config.AesKeySalt = getAesKeySalt(t)
	cfg := config.NewConfig(path.Join(path.Dir(filename), "conf-file.toml"))
	helper := NewAliApiHelper(cfg)

	ret, err := helper.OCRIDCardFace(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("识别结果：%v", ret)
}

func TestOCRIDCard2(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	data, err := ioutil.ReadFile(path.Join(path.Dir(filename), "idcard_sample2.jpg"))
	if err != nil {
		t.Fatal(err)
	}

	config.AesKeySalt = getAesKeySalt(t)
	cfg := config.NewConfig(path.Join(path.Dir(filename), "conf-file.toml"))
	helper := NewAliApiHelper(cfg)

	ret, err := helper.OCRIDCardBack(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("识别结果：%v", ret)
}

func getAesKeySalt(t *testing.T) string {
	_, filename, _, _ := runtime.Caller(0)
	data, err := ioutil.ReadFile(path.Join(path.Dir(filename), "aeskeysalt.txt"))
	if err != nil {
		t.Fatal(err)
	}
	return string(data)

}
