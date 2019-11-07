package alipay

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/jinlongchen/golang-utilities/crypto"
	"github.com/jinlongchen/golang-utilities/http"
	"github.com/jinlongchen/golang-utilities/json"
	"github.com/jinlongchen/golang-utilities/log"
	"io/ioutil"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	AppID        string
	SellerID     string
	PrivateKey   *rsa.PrivateKey
	AliPublicKey *rsa.PublicKey
}

func NewClient(appID, sellerID, priKeyPath, aliPublicKeyPath string) (*Client, error) {
	priKey, err := loadPrivateKey(priKeyPath)
	if err != nil {
		log.Fatalf("LoadPrivateKey err:%s", err.Error())
		return nil, err
	}
	aliPublicKey, err := loadPublicKey(aliPublicKeyPath)
	if err != nil {
		log.Fatalf("LoadPublicKey err:%s", err.Error())
		return nil, err
	}

	return &Client{
		AppID:        appID,
		SellerID:     sellerID,
		PrivateKey:   priKey,
		AliPublicKey: aliPublicKey,
	}, nil
}

// alipay.trade.app.pay(app支付接口2.0)
func (client *Client) CreateTradeAppPay(orderID string, fee int, subject string, desc string, notifyUrl, returnUrl string) (string, error) {
	reqUrl, err := client.getOpenApiBizRequestUrl(
		"alipay.trade.app.pay",
		returnUrl,
		notifyUrl,
		json.ShouldMarshal(&TradeAppPay{
			OutTradeNo:  orderID,
			Subject:     subject,
			Body:        desc,
			TotalAmount: strconv.FormatFloat(float64(fee)/100, 'f', 2, 32),
			ProductCode: "QUICK_MSECURITY_PAY",
		}),
	)

	if err != nil {
		log.Errorf("err:%s", err.Error())
	}

	return reqUrl, err
}

//退款
func (client *Client) Refund(orderID string, fee int, reason string, operator string, notifyUrl, returnUrl string) error {
	reqMap := map[string]string{
		"out_trade_no":   orderID,
		"out_request_no": fmt.Sprintf("refund_%s", orderID),
		"refund_amount":  strconv.FormatFloat(float64(fee)/100, 'f', 2, 32),
		"refund_reason":  reason,
		"operator_id":    operator,
	}

	reqUrl, err := client.getOpenApiBizRequestUrl(
		"alipay.trade.refund",
		notifyUrl,
		returnUrl,
		json.ShouldMarshal(reqMap),
	)

	if err != nil {
		log.Errorf("getOpenApiBizRequestUrl err:%s", err.Error())
		return err
	}

	data, err := http.GetData(reqUrl)
	if err != nil {
		log.Errorf("GetDataSkipTls err:%s", err.Error())
		return err
	}

	refundResponse := &RefundResponse{}
	err = json.Unmarshal(data, refundResponse)
	if err != nil {
		return err
	}
	if refundResponse.AlipayTradeRefundResponse.Code != "10000" {
		return errors.New(refundResponse.AlipayTradeRefundResponse.Msg)
	}
	return nil
}

func (client *Client) QueryOrder(orderID string) (queryOrderResponse *QueryOrderResponse, err error) {
	reqMap := map[string]string{
		"out_trade_no": orderID,
	}

	reqUrl, err := client.getOpenApiBizRequestUrl(
		"alipay.trade.query",
		"",
		"",
		json.ShouldMarshal(reqMap),
	)

	if err != nil {
		log.Errorf("getOpenApiBizRequestUrl err:%s", err.Error())
		return nil, err
	}

	data, err := http.GetData(reqUrl)
	if err != nil {
		log.Errorf("GetDataSkipTls err:%s", err.Error())
		return nil, err
	}

	queryOrderResponse = &QueryOrderResponse{}
	err = json.Unmarshal(data, queryOrderResponse)
	if err != nil {
		return nil, err
	}
	if queryOrderResponse.AlipayTradeQueryResponse.Code != "10000" {
		return nil, errors.New(queryOrderResponse.AlipayTradeQueryResponse.Msg)
	} else {
		return queryOrderResponse, nil
	}
}

//RSA方式的验签（新版本的手机端支付）
func (client *Client) VerifyTradeWapPay(plainText, sign string) error {
	return crypto.RSA256Verify(client.AliPublicKey, []byte(plainText), sign)
}

//RSA签名（新版本的手机端支付）
func (client *Client) SignTradeWapPay(plainText string) (string, error) {
	return crypto.RSA256Sign(client.PrivateKey, []byte(plainText))
}

func (client *Client) getOpenApiBizRequestUrl(method, notifyUrl, returnUrl string, bizData []byte) (string, error) {
	var (
		err error
	)
	gatewayUrl := "https://openapi.alipay.com/gateway.do"
	var query = Params{
		"app_id":      client.AppID,
		"biz_content": string(bizData),
		"charset":     "utf-8",
		"format":      "JSON",
		"method":      method,
		"notify_url":  notifyUrl,
		"return_url":  returnUrl,
		"sign_type":   "RSA2",
		"timestamp":   time.Now().Format("2006-01-02 15:04:05"),
		"version":     "1.0",
	}
	if returnUrl != "" {
		query["return_url"] = returnUrl
	}
	if notifyUrl != "" {
		query["notify_url"] = notifyUrl
	}

	query["sign"], err = crypto.RSA256Sign(client.PrivateKey, []byte(query.Encode(false)))
	if err != nil {
		return "", err
	}

	reqUrl, _ := url.Parse(gatewayUrl)
	reqUrl.RawQuery = map2Values(query).Encode()

	return reqUrl.String(), nil
}

func map2Values(i map[string]string) (values url.Values) {
	values = url.Values{}
	for key, value := range i {
		values.Set(key, value)
	}
	return
}

func loadPublicKey(publicKeyPath string) (*rsa.PublicKey, error) {
	certPEMBlock, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}

	return crypto.ParseRSAPublicKeyFromPEM(certPEMBlock)
}

func loadPrivateKey(priKeyPath string) (*rsa.PrivateKey, error) {
	certPEMBlock, err := ioutil.ReadFile(priKeyPath)
	if err != nil {
		return nil, err
	}

	return crypto.ParseRSAPrivateKeyFromPEM(certPEMBlock)
}
