package config

import (
	"fmt"
	"github.com/jinlongchen/golang-utilities/map/helper"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"testing"
	"time"

	"github.com/jinlongchen/golang-utilities/converter"
)

func TestNewEtcdConfig(t *testing.T) {
	cfg := NewRemoteConfig("etcd", "http://192.168.2.42:2379", "/configs/a.toml", "toml")
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			println(cfg.GetInt("abc.def"))
		}
	}
}
func TestConfig_EncryptString2(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	tomlFile := path.Join(path.Dir(filename), "conf-file.toml")

	cfg := NewConfig(tomlFile)

	AesKeySalt = os.Getenv("GO_AES_SALT")

	println("code.qcse.com:", cfg.EncryptString("code.qcse.com"))
	println("https://code.qcse.com/:", cfg.EncryptString("https://code.qcse.com/"))
	println("MFrL8AXyOIH6uuBABwKKUGlAVoKVSIX4ObCa61GZtE2KZOm6J72fxO85wOSE99Oq:", cfg.EncryptString("MFrL8AXyOIH6uuBABwKKUGlAVoKVSIX4ObCa61GZtE2KZOm6J72fxO85wOSE99Oq"))
	println("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE1NDMyODk5MTN9.EV30v2VZQ1hiQhcTpCHf55UX7AfmzlMg5T2iD3neehs:", cfg.EncryptString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE1NDMyODk5MTN9.EV30v2VZQ1hiQhcTpCHf55UX7AfmzlMg5T2iD3neehs"))
}

func TestConfig_EncryptString(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	tomlFile := path.Join(path.Dir(filename), "conf-file.toml")

	cfg := NewConfig(tomlFile)

	AesKeySalt = os.Getenv("GO_AES_SALT")

	//println("server param:", cfg.EncryptString("ss://CHACHA20-IETF-POLY1305:981f08f5dcb43b1ac05f82632C65_d160127f2cc7@:11197"))
	//// local -> hk server
	println("ss://CHACHA20-IETF-POLY1305:981f08f5dcb43b1ac05f82632C65_76f070767cfd7eafb920@[kpc]:11196:",
		cfg.EncryptString("ss://CHACHA20-IETF-POLY1305:981f08f5dcb43b1ac05f82632C65_76f070767cfd7eafb920@[kpc]:11196"))
	//println("ss://CHACHA20-IETF-POLY1305:981f08f5dcb43b1ac05f82632C65_76f070767cfd7eafb920@:11197 = ",
	//	cfg.EncryptString("ss://CHACHA20-IETF-POLY1305:981f08f5dcb43b1ac05f82632C65_76f070767cfd7eafb920@:11197"))

	//println("ss://CHACHA20-IETF-POLY1305:981f08f5dcb43b1ac05f82632C65_d160127f2cc7@[35.182.130.31]:11197:", cfg.EncryptString("ss://CHACHA20-IETF-POLY1305:981f08f5dcb43b1ac05f82632C65_d160127f2cc7@[35.182.130.31]:11197"))
	//println("149.129.83.138:11195", cfg.EncryptString("149.129.83.138:11195"))
	println("149.129.83.138:11194", cfg.EncryptString("149.129.83.138:11194"))
	println("149.129.83.138:11195", cfg.EncryptString("149.129.83.138:11195"))

	//println("ss://CHACHA20-IETF-POLY1305:981f08f5dcb43b1ac05f82632C65_d160127f2cc7@[47.240.20.123]:11197:",
	//	cfg.EncryptString("ss://CHACHA20-IETF-POLY1305:981f08f5dcb43b1ac05f82632C65_d160127f2cc7@[47.240.20.123]:11197"))

	//println("ss://CHACHA20-IETF-POLY1305:981f08f5dcb43b1ac05f82632C65_d160127f2cc7@[47.52.232.133]:11197:", cfg.EncryptString("ss://CHACHA20-IETF-POLY1305:981f08f5dcb43b1ac05f82632C65_d160127f2cc7@[47.52.232.133]:11197"))
	//println("->> local server param kcp:", cfg.EncryptString("ss://CHACHA20-IETF-POLY1305:981f08f5dcb43b1ac05f82632C65_d160127f2cc7@[172.17.0.5]:11196"))
	//println("->> kcp server config:", cfg.EncryptString("172.17.0.2:11197"))
	//println("->> kcp server config ca:", cfg.EncryptString("35.182.130.31:11197"))
	//
	//println("->> ca 35.182.130.31:11195:", cfg.EncryptString("35.182.130.31:11195"))
	//println("->> ca 52.60.222.79:11195:", cfg.EncryptString("52.60.222.79:11195"))
	//println("->> jp 54.95.26.192:11195:", cfg.EncryptString("54.95.26.192:11195"))
	//println("->> ca 35.183.4.31:11195:", cfg.EncryptString("35.183.4.31:11195"))
	//
	//println("->> :4000:", cfg.EncryptString(":4000"))
	//println("->> kpc:11196:", cfg.EncryptString("kpc:11196")) // hHLayibyzDjn59jeqTc8QA==
	//println("->> cache2:11197:", cfg.EncryptString("cache2:11197")) // hHLayibyzDjn59jeqTc8QA==
	//
	//println("xxx:", cfg.GetString("file.minio.accessid"))
	//println("yyy:", cfg.GetString("file.minio.accesssecret"))
}

func TestConfig_DecryptString(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	tomlFile := path.Join(path.Dir(filename), "conf-file.toml")

	cfg := NewConfig(tomlFile)

	AesKeySalt = os.Getenv("GO_AES_SALT")

//ss://CHACHA20-IETF-POLY1305:981f08f5dcb43b1ac05f82632C65_d160127f2cc7@[kpc]:11196
//ss://CHACHA20-IETF-POLY1305:981f08f5dcb43b1ac05f82632C65_76f070767cfd7eafb920@:11197
	println(cfg.DecryptString(`Mv3GdsNAMkqgdRx+2pdZFZ/nhjolcIAYiprUkJ/SYA3amBxet6Rtp3mlnaT399TG4vbMmiIO3r+H8OeV2NJbOsCYeirTNZ2jQ9AGzmBjR48nGgfjIThS8I0V8Fy9Djxw`))
	println(cfg.DecryptString(`Mv3GdsNAMkqgdRx+2pdZFZ/nhjolcIAYiprUkJ/SYA3amBxet6Rtp3mlnaT399TG4vbMmiIO3r+H8OeV2NJbOiSpxnYwEGI6VA8G0j22Hw+iPJXZZAO+3XhmGuSxw93S`))
	println(cfg.DecryptString(`Mv3GdsNAMkqgdRx+2pdZFZ/nhjolcIAYiprUkJ/SYA3amBxet6Rtp3mlnaT399TGSal/299uOLtxHjbtaFB6d4x4ZDGErhpsB3RDrYuD6p3NnTlg10fCxMGlusP+dT09`))
	println(cfg.DecryptString(`K343GB5qBBxFc+CRfCwThFnLRax20m4tQ6GXnONH9VA=`))
	println(cfg.DecryptString(`Z5WHqYZtEQCNVRkGUZFt/6Klocyl17f6EI7ndcE2rUM=`))

}
func TestConfig_GetStringSlice(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	tomlFile := path.Join(path.Dir(filename), "conf-file.toml")

	cfg := NewConfig(tomlFile)
	fmt.Printf("xxx:%v\n", cfg.GetStringSlice("test.key2"))
	fmt.Printf("xxx2:%v\n", cfg.GetStringSlice("test.key2"))
}
func TestConfig_GetString(t *testing.T) {
	AesKeySalt = os.Getenv("GO_AES_SALT")

	_, filename, _, _ := runtime.Caller(0)
	tomlFile := path.Join(path.Dir(filename), "conf-file.toml")

	cfg := NewConfig(tomlFile)
	corpps := cfg.GetMapSlice("corpps")
	fmt.Printf("xxx2:%v\n", corpps)
	for _, corpp := range corpps {
		corpName := helper.GetValueAsString(corpp, "corpName", "")
		corpID := helper.GetValueAsString(corpp, "corpID", "")

		fmt.Printf("%s %s\n",corpID, corpName)
	}
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
