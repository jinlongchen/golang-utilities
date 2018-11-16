package aliapi

import (
	"testing"
	"runtime"
	"path"
	"io/ioutil"
	"github.com/jinlongchen/golang-utilities/config"
)

func TestOCRIDCard(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	data, err := ioutil.ReadFile(path.Join(path.Dir(filename), "idcard_sample.jpeg"))
	if err != nil {
		t.Fatal(err)
	}

	config.AesKeyKey = getAesKeyKey(t)
	cfg := config.NewConfig(path.Join(path.Dir(filename), "conf-file.toml"))
	helper := NewSAliApiHelper(cfg)

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

	config.AesKeyKey = getAesKeyKey(t)
	cfg := config.NewConfig(path.Join(path.Dir(filename), "conf-file.toml"))
	helper := NewSAliApiHelper(cfg)

	ret, err := helper.OCRIDCardBack(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("识别结果：%v", ret)
}

func getAesKeyKey(t *testing.T) string {
	_, filename, _, _ := runtime.Caller(0)
	data, err := ioutil.ReadFile(path.Join(path.Dir(filename), "aeskeykey.txt"))
	if err != nil {
		t.Fatal(err)
	}
	return string(data)

}
