/*
 * Copyright (c) 2019. Qing Cheng Technology Co., Ltd.
 */

package alipay

import (
	"fmt"
	"github.com/jinlongchen/golang-utilities/config"
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

	//sign, err := crypto.RSA256Sign(aliClient.PrivateKey, []byte(`app_id=` + cfg.GetString("alipay.appID") + `&biz_content={"timeout_express":"30m","seller_id":"","product_code":"QUICK_MSECURITY_PAY","total_amount":"0.01","subject":"1","body":"我是测试数据","out_trade_no":"TZHAXSOGRGQQRTA"}&charset=utf-8&method=alipay.trade.app.pay&sign_type=RSA2&timestamp=2019-11-16 18:46:16&version=1.0`))
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(sign)

	respData, err := aliClient.CreateTradeAppPay("TZHAXSOGRGQQRTA",
		1, "1", "我是测试数据",
		"",
		"",
		)
	if err != nil {
		panic(err)
	}
	fmt.Println(respData)
}
