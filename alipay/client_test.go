/*
 * Copyright (c) 2019. Qing Cheng Technology Co., Ltd.
 */

package alipay

import (
	"fmt"
	"github.com/jinlongchen/golang-utilities/config"
	"github.com/jinlongchen/golang-utilities/rand"
	"io/ioutil"
	"path"
	"runtime"
	"testing"
)

func TestClient_CreateTradeAppPay(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	cfg := config.NewConfig(path.Join(path.Dir(filename), "test.toml"))
	aliClient, err := NewClient(
		cfg.GetString("alipay.appID"),
		cfg.GetString("alipay.paternerID"),
		cfg.GetString("alipay.priKeyPath"),
		cfg.GetString("alipay.aliPublicKeyPath"),
	)
	if err != nil {
		panic(err)
	}
	respData, err := aliClient.CreateTradeAppPay(rand.GetShortTimestampRandString(),
		1, "测试", "测试1分钱",
		cfg.GetString("alipay.app.notifyURL"),
		cfg.GetString("alipay.app.returnURL"),
		)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(path.Join(path.Dir(filename), "resp.html"), respData, 0666)
	fmt.Println(string(respData))
}
