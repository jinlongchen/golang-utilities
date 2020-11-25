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
func TestClient_VerifyTradeAppPayNotify(t *testing.T) {
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
	err = aliClient.VerifyTradeAppPayNotify(`gmt_create=2019-11-17+03%3A56%3A52&charset=utf-8&seller_email=chenjinlong%40qcse.com&subject=%E8%BD%BB%E7%A8%8B%E5%87%BA%E8%A1%8C%E5%85%85%E5%80%BC+30%E5%85%83&sign=VKO17JR4wNtTcDxx0R6i85mZxOZzseDPpm5x7bpBfRu3zreVmt1kJgBgzYD9ws6nVNxrfOqzEZ3ozovXeY%2FxERiKPWS60E%2Fgy3xe5Jco%2FyoK8VPgG3gDgT3A%2FApX%2BBVGn6sKEMQs3ZrMP2hQZfWJizGGIdyQmbD7iFE%2FBWtLCBHZ8u%2FeDrhU84e5P6KJsf%2FT4%2BiiJz0I8ZIoUyXs%2BCS75a0%2BJH2EK6jz8O5G2478mwDqmSG8v%2FxbaERCSzzDnzj0C%2FxUC02dzsv1wiMjamm8qJzSUAZMY4SgBmv9nJdTIn%2FYaHIVSy6lUmXGZGwQh%2BpMp3AanJ%2BgB6AvNKpy80Z%2BJQ%3D%3D&body=%E8%BD%BB%E7%A8%8B%E5%87%BA%E8%A1%8C%E5%85%85%E5%80%BC+30%E5%85%83&buyer_id=2088702140287087&invoice_amount=0.01&notify_id=2019111700222035653087080570858517&fund_bill_list=%5B%7B%22amount%22%3A%220.01%22%2C%22fundChannel%22%3A%22PCREDIT%22%7D%5D&notify_type=trade_status_sync&trade_status=TRADE_SUCCESS&receipt_amount=0.01&app_id=2019110768980114&buyer_pay_amount=0.01&sign_type=RSA2&seller_id=2088631274749102&gmt_payment=2019-11-17+03%3A56%3A52&notify_time=2019-11-17+03%3A56%3A53&version=1.0&out_trade_no=32wltt02kz6sl805&total_amount=0.01&trade_no=2019111722001487080541566585&auth_app_id=2019110768980114&buyer_logon_id=131****1710&point_amount=0.00`)
	if err != nil {
		panic(err)
	}
	fmt.Println("OK")
}
