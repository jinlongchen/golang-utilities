package wechat

import (
	"encoding/xml"
	gu_xml "github.com/jinlongchen/golang-utilities/xml"
)

type NotifyData struct {
	XMLName xml.Name `xml:"xml"`
	// Coupon_batch_id_$n   string `xml:"coupon_batch_id_$n,omitempty"`
	// Coupon_count   string `xml:"coupon_count,omitempty"`
	// Coupon_fee   string `xml:"coupon_fee,omitempty"`
	// Coupon_fee_$n   string `xml:"coupon_fee_$n,omitempty"`
	// Coupon_id_$n   string `xml:"coupon_id_$n,omitempty"`
	AppId         gu_xml.CData `xml:"appid"` //must
	Attach        gu_xml.CData `xml:"attach,omitempty"`
	BankType      gu_xml.CData `xml:"bank_type"`
	CashFee       gu_xml.CData `xml:"cash_fee"`
	CashFeeType   gu_xml.CData `xml:"cash_fee_type,omitempty"`
	DeviceInfo    gu_xml.CData `xml:"device_info,omitempty"` //maybe
	ErrCode       gu_xml.CData `xml:"err_code,omitempty"`
	ErrCodeDes    gu_xml.CData `xml:"err_code_des,omitempty"`
	FeeType       gu_xml.CData `xml:"fee_type,omitempty"`
	IsSubscribe   gu_xml.CData `xml:"is_subscribe"`
	MchId         gu_xml.CData `xml:"mch_id"`    //must
	NonceStr      gu_xml.CData `xml:"nonce_str"` //must
	OpenId        gu_xml.CData `xml:"openid"`
	OutTradeNo    gu_xml.CData `xml:"out_trade_no"`
	ResultCode    gu_xml.CData `xml:"result_code"`
	ReturnCode    gu_xml.CData `xml:"return_code"`
	ReturnMsg     gu_xml.CData `xml:"return_msg"`
	Sign          gu_xml.CData `xml:"sign"`     //must
	TimeEnd       gu_xml.CData `xml:"time_end"` //支付完成时间
	TotalFee      gu_xml.CData `xml:"total_fee"`
	TradeType     gu_xml.CData `xml:"trade_type"`
	TransactionId gu_xml.CData `xml:"transaction_id"`

	XML string `xml:"-"` //结果串
}

type NotifyResponse struct {
	XMLName    xml.Name     `xml:"xml"`
	ReturnCode gu_xml.CData `xml:"return_code"`
	ReturnMsg  gu_xml.CData `xml:"return_msg"`
	XML        string       `xml:"-"` //最终请求串
}

func (v *NotifyResponse) Xml() error {
	data, err := xml.Marshal(v)
	v.XML = string(data)

	return err

}
