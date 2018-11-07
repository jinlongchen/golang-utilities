package config

import (
	"github.com/jinlongchen/golang-utilities/crypto"
	"github.com/jinlongchen/golang-utilities/log"
	"github.com/jinlongchen/golang-utilities/converter"
	"crypto/aes"
	"encoding/base64"
	"strings"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	"github.com/naoina/toml"
	"io/ioutil"
)

var (
	AesKeyKey string
)
type Config struct {
	cache     map[string]interface{}
	v         *viper.Viper
	aesKey    []byte
}

func NewConfig(path string) *Config {
	ret := &Config{
		cache:     make(map[string]interface{}),
		v:         viper.New(),
	}
	ret.v.SetConfigFile(path)
	err := ret.v.ReadInConfig()
	if err != nil {
		log.Errorf("read log file err:%s", err.Error())
	}
	ret.v.WatchConfig()
	ret.v.OnConfigChange(func(e fsnotify.Event) {
		log.Debugf("reload config")
		ret.cache = make(map[string]interface{})
	})
	return ret
}
func (cfg *Config) BindEnv(input ...string) error {
	return cfg.v.BindEnv(input...)
}
func (cfg *Config) SetDefault(key string, value interface{}) {
	cfg.v.SetDefault(key, value)
}
func (cfg *Config) GetString(path string) string {
	if val, ok := cfg.cache[path]; ok {
		return converter.AsString(val, "")
	}
	ret := cfg.v.GetString(path)
	if strings.HasPrefix(ret, "aes:") {
		ret = cfg.DecryptString(ret[4:])
	}
	cfg.cache[path] = ret
	return ret
}

func (cfg *Config) GetInt(path string) int {
	if val, ok := cfg.cache[path]; ok {
		return converter.AsInt(val, 0)
	}
	ret := cfg.v.GetInt(path)
	cfg.cache[path] = ret
	return ret
}

func (cfg *Config) GetInt64(path string) int64 {
	if val, ok := cfg.cache[path]; ok {
		return converter.AsInt64(val, 0)
	}
	ret := cfg.v.GetInt64(path)
	cfg.cache[path] = ret
	return ret
}

func (cfg *Config) GetBool(path string) bool {
	if val, ok := cfg.cache[path]; ok {
		return converter.AsBool(val, false)
	}
	ret := cfg.v.GetBool(path)
	cfg.cache[path] = ret
	return ret
}

func (cfg *Config) GetStringSlice(path string) []string {
	if val, ok := cfg.cache[path]; ok {
		if r, ok := val.([]string); ok {
			return r
		}
	}

	ret := cfg.v.GetStringSlice(path) //map_helper.GetValue(cfg.data, path)
	cfg.cache[path] = ret
	return ret
}

func (cfg *Config) GetStringMapString(path string) map[string]string {
	if val, ok := cfg.cache[path]; ok {
		if r, ok := val.(map[string]string); ok {
			return r
		}
	}
	ret := cfg.v.GetStringMapString(path) //map_helper.GetValue(cfg.data, path)
	cfg.cache[path] = ret
	return ret
}

func (cfg *Config) GetValue(path string) interface{} {
	if val, ok := cfg.cache[path]; ok {
		return val
	}
	ret := cfg.v.Get(path)
	cfg.cache[path] = ret
	return ret
}

func (cfg *Config) DecryptString(str string) string {
	if cfg.aesKey == nil {
		aesKey1 := cfg.GetString("crypto.aesKey")
		cfg.aesKey = crypto.String(aesKey1 + AesKeyKey).GetMd5()
	}

	eData, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Errorf("config DecryptString(%s) err:%s", str, err.Error())
		return ""
	}
	dData, err := crypto.AESDecryptCBC(eData, cfg.aesKey, cfg.aesKey[:aes.BlockSize])
	if err != nil {
		log.Errorf("config DecryptString(%s) err:%s", str, err.Error())
		return ""
	}

	return string(dData)
}

func (cfg *Config) EncryptString(str string) string {
	if cfg.aesKey == nil {
		aesKey1 := cfg.GetString("crypto.aesKey")
		cfg.aesKey = crypto.String(aesKey1 + AesKeyKey).GetMd5()
	}

	dData := []byte(str)
	eData, err := crypto.AESEncryptCBC(dData, cfg.aesKey, cfg.aesKey[:aes.BlockSize])
	if err != nil {
		log.Fatalf("config DecryptString(%s) err:%s", str, err.Error())
	}

	return base64.StdEncoding.EncodeToString(eData)
}
func (cfg *Config) Save(path string) error {
	data, err := toml.Marshal(cfg.v.AllSettings())
	if err != nil {
		return err
	}
	if path != "" {
		err = ioutil.WriteFile(path, data, 0666)
	}
	return err
}
