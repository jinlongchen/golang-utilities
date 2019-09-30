package config

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"testing"
	"time"

	"github.com/jinlongchen/golang-utilities/converter"
)

func TestNewEtcdConfig(t *testing.T) {
	cfg := NewRemoteConfig("etcd", "http://192.168.2.42:2379", "/configs/a.toml")
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			println(cfg.GetInt("abc.def"))
		}
	}
}
func TestConfig_EncryptString(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	tomlFile := path.Join(path.Dir(filename), "conf-file.toml")

	cfg := NewConfig(tomlFile)

	//AesKeyKey = getProductionAesKeyKey(t)

	println("xxx:", cfg.GetString("file.minio.accessid"))
	println("yyy:", cfg.GetString("file.minio.accesssecret"))
}

func TestConfig_GetStringSlice(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	tomlFile := path.Join(path.Dir(filename), "conf-file.toml")

	cfg := NewConfig(tomlFile)
	fmt.Printf("xxx:%v\n", cfg.GetStringSlice("test.key2"))
	fmt.Printf("xxx2:%v\n", cfg.GetStringSlice("test.key2"))
}
func TestConfig_GetBool(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	tomlFile := path.Join(path.Dir(filename), "conf-file.toml")

	cfg := NewConfig(tomlFile)
	fmt.Printf("xxx:%v\n", cfg.GetBool("test.key3"))
	fmt.Printf("xxx2:%v\n", cfg.GetBool("test.key3"))
}
func TestConfig_GetInt(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	tomlFile := path.Join(path.Dir(filename), "conf-file.toml")

	cfg := NewConfig(tomlFile)
	fmt.Printf("xxx:%v\n", cfg.GetInt("test.key4"))
	fmt.Printf("xxx2:%v\n", cfg.GetInt("test.key4"))
}
func TestConfig_GetDuration(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	tomlFile := path.Join(path.Dir(filename), "conf-file.toml")

	cfg := NewConfig(tomlFile)
	fmt.Printf("xxx:%v\n", cfg.GetDuration("test.key5"))
	fmt.Printf("xxx2:%v\n", cfg.GetDuration("test.key5"))
}
func TestConfig_GetValue(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	tomlFile := path.Join(path.Dir(filename), "conf-file.toml")

	cfg := NewConfig(tomlFile)
	origins1 := converter.AsStringSlice(cfg.GetValue("http.header.origins"), []string{})
	origins2 := converter.AsStringSlice(cfg.GetValue("http.header.origins"), []string{})
	fmt.Printf("xxx:%v\n", origins1)
	fmt.Printf("xxx2:%v\n", origins2)

}
func TestConfig_GetStringMapString(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	tomlFile := path.Join(path.Dir(filename), "conf-file.toml")

	cfg := NewConfig(tomlFile)
	fmt.Printf("xxx :%v\n", cfg.GetStringMapString("cache.redis.addresses"))
	fmt.Printf("xxx2:%v\n", cfg.GetStringMapString("cache.redis.addresses"))
}
func getTestingAesKeyKey(t *testing.T) string {
	_, filename, _, _ := runtime.Caller(0)
	data, err := ioutil.ReadFile(path.Join(path.Dir(filename), "aeskeykey.txt"))
	if err != nil {
		t.Fatal(err)
	}
	return string(data)

}

func getProductionAesKeyKey(t *testing.T) string {
	_, filename, _, _ := runtime.Caller(0)
	data, err := ioutil.ReadFile(path.Join(path.Dir(filename), "aeskeykey_yijiu.txt"))
	if err != nil {
		t.Fatal(err)
	}
	return string(data)

}
