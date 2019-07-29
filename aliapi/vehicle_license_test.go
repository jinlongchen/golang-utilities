package aliapi

import (
	"github.com/jinlongchen/golang-utilities/config"
	"io/ioutil"
	"path"
	"runtime"
	"testing"
)

func TestAliApiHelper_OCRVehicleLicenseFace(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	data, err := ioutil.ReadFile(path.Join(path.Dir(filename), "vehicle_lic_sample.jpg"))
	if err != nil {
		t.Fatal(err)
	}

	config.AesKeyKey = getAesKeyKey(t)
	cfg := config.NewConfig(path.Join(path.Dir(filename), "conf-file.toml"))

	helper := NewAliApiHelper(cfg)

	ret, err := helper.OCRVehicleLicenseFace(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("识别结果：%v", ret)
}

func TestAliApiHelper_OCRVehicleLicenseBack(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	data, err := ioutil.ReadFile(path.Join(path.Dir(filename), "vehicle_lic_sample2.jpg"))
	if err != nil {
		t.Fatal(err)
	}

	config.AesKeyKey = getAesKeyKey(t)
	cfg := config.NewConfig(path.Join(path.Dir(filename), "conf-file.toml"))

	helper := NewAliApiHelper(cfg)

	ret, err := helper.OCRVehicleLicenseBack(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("识别结果：%v", ret)
}
