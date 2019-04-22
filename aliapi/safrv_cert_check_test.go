package aliapi

import (
	"testing"
	"runtime"
	"path"
	"github.com/jinlongchen/golang-utilities/config"
	"github.com/jinlongchen/golang-utilities/log"
)

func TestAliApiHelper_CheckIdCardNameMatch(t *testing.T) {
	log.Config("test", "debug", log.FormatJSON)
	_, filename, _, _ := runtime.Caller(0)

	config.AesKeyKey = getAesKeyKey(t)
	cfg := config.NewConfig(path.Join(path.Dir(filename), "conf-file.toml"))

	helper := NewAliApiHelper(cfg)

	ret, err := helper.CheckIdCardNameMatch("6540****0386", "梁素华")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("识别结果：%v", string(ret))
}
