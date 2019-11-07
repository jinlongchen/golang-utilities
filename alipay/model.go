package alipay

type TradeAppPay struct {
	Body               string `json:"body" xml:"body"`
	BusinessParams     string `json:"business_params" xml:"business_params"`
	DisablePayChannels string `json:"disable_pay_channels" xml:"disable_pay_channels"`
	EnablePayChannels  string `json:"enable_pay_channels" xml:"enable_pay_channels"`
	GoodsType          string `json:"goods_type" xml:"goods_type"`
	MerchantOrderNo    string `json:"merchant_order_no" xml:"merchant_order_no"`
	OutTradeNo         string `json:"out_trade_no" xml:"out_trade_no"`
	PassbackParams     string `json:"passback_params" xml:"passback_params"`
	ProductCode        string `json:"product_code" xml:"product_code"`
	PromoParams        string `json:"promo_params" xml:"promo_params"`
	SpecifiedChannel   string `json:"specified_channel" xml:"specified_channel"`
	StoreID            string `json:"store_id" xml:"store_id"`
	Subject            string `json:"subject" xml:"subject"`
	TimeExpire         string `json:"time_expire" xml:"time_expire"`
	TimeoutExpress     string `json:"timeout_express" xml:"timeout_express"`
	TotalAmount        string `json:"total_amount" xml:"total_amount"`
	GoodsDetail        []struct {
		Quantity       int    `json:"quantity" xml:"quantity"`
		Price          int    `json:"price" xml:"price"`
		GoodsCategory  string `json:"goods_category" xml:"goods_category"`
		GoodsID        string `json:"goods_id" xml:"goods_id"`
		AlipayGoodsID  string `json:"alipay_goods_id" xml:"alipay_goods_id"`
		GoodsName      string `json:"goods_name" xml:"goods_name"`
		CategoriesTree string `json:"categories_tree" xml:"categories_tree"`
		Body           string `json:"body" xml:"body"`
		ShowURL        string `json:"show_url" xml:"show_url"`
	} `json:"goods_detail" xml:"goods_detail"`
	ExtUserInfo struct {
		Name          string `json:"name" xml:"name"`
		Mobile        string `json:"mobile" xml:"mobile"`
		CertType      string `json:"cert_type" xml:"cert_type"`
		CertNo        string `json:"cert_no" xml:"cert_no"`
		MinAge        string `json:"min_age" xml:"min_age"`
		FixBuyer      string `json:"fix_buyer" xml:"fix_buyer"`
		NeedCheckInfo string `json:"need_check_info" xml:"need_check_info"`
	} `json:"ext_user_info" xml:"ext_user_info"`
	AgreementSignParams struct {
		ExternalLogonID string `json:"external_logon_id" xml:"external_logon_id"`
		AccessParams    struct {
			Channel string `json:"channel" xml:"channel"`
		} `json:"access_params" xml:"access_params"`
		SubMerchant struct {
			SubMerchantID                 string `json:"sub_merchant_id" xml:"sub_merchant_id"`
			SubMerchantName               string `json:"sub_merchant_name" xml:"sub_merchant_name"`
			SubMerchantServiceName        string `json:"sub_merchant_service_name" xml:"sub_merchant_service_name"`
			SubMerchantServiceDescription string `json:"sub_merchant_service_description" xml:"sub_merchant_service_description"`
		} `json:"sub_merchant" xml:"sub_merchant"`
		PeriodRuleParams struct {
			PeriodType    string  `json:"period_type" xml:"period_type"`
			Period        int     `json:"period" xml:"period"`
			ExecuteTime   string  `json:"execute_time" xml:"execute_time"`
			SingleAmount  float64 `json:"single_amount" xml:"single_amount"`
			TotalAmount   int     `json:"total_amount" xml:"total_amount"`
			TotalPayments int     `json:"total_payments" xml:"total_payments"`
		} `json:"period_rule_params" xml:"period_rule_params"`
		PersonalProductCode string `json:"personal_product_code" xml:"personal_product_code"`
		SignScene           string `json:"sign_scene" xml:"sign_scene"`
		ExternalAgreementNo string `json:"external_agreement_no" xml:"external_agreement_no"`
	} `json:"agreement_sign_params" xml:"agreement_sign_params"`
	ExtendParams struct {
		CardType             string `json:"card_type" xml:"card_type"`
		SysServiceProviderID string `json:"sys_service_provider_id" xml:"sys_service_provider_id"`
		HbFqNum              string `json:"hb_fq_num" xml:"hb_fq_num"`
		HbFqSellerPercent    string `json:"hb_fq_seller_percent" xml:"hb_fq_seller_percent"`
		IndustryRefluxInfo   string `json:"industry_reflux_info" xml:"industry_reflux_info"`
	} `json:"extend_params" xml:"extend_params"`
}
type RefundResponse struct {
	Sign                      string `json:"sign" xml:"sign"`
	AlipayTradeRefundResponse struct {
		Code    string `json:"code" xml:"code"`
		Msg     string `json:"msg" xml:"msg"`
		SubCode string `json:"sub_code" xml:"sub_code"`
		SubMsg  string `json:"sub_msg" xml:"sub_msg"`

		BuyerLogonID string  `json:"buyer_logon_id" xml:"buyer_logon_id"`
		BuyerUserID  string  `json:"buyer_user_id" xml:"buyer_user_id"`
		FundChange   string  `json:"fund_change" xml:"fund_change"`
		GmtRefundPay string  `json:"gmt_refund_pay" xml:"gmt_refund_pay"`
		OpenID       string  `json:"open_id" xml:"open_id"`
		OutTradeNo   string  `json:"out_trade_no" xml:"out_trade_no"`
		RefundFee    float64 `json:"refund_fee" xml:"refund_fee"`
		StoreName    string  `json:"store_name" xml:"store_name"`
		TradeNo      string  `json:"trade_no" xml:"trade_no"`

		RefundDetailItemList []struct {
			Amount      int     `json:"amount" xml:"amount"`
			FundChannel string  `json:"fund_channel" xml:"fund_channel"`
			RealAmount  float64 `json:"real_amount" xml:"real_amount"`
		} `json:"refund_detail_item_list" xml:"refund_detail_item_list"`
	} `json:"alipay_trade_refund_response" xml:"alipay_trade_refund_response"`
}

type QueryOrderResponse struct {
	AlipayTradeQueryResponse struct {
		TradeNo        string `json:"trade_no" xml:"trade_no"`
		TradeStatus    string `json:"trade_status" xml:"trade_status"`
		SubCode        string `json:"sub_code" xml:"sub_code"`
		BuyerLogonID   string `json:"buyer_logon_id" xml:"buyer_logon_id"`
		BuyerUserID    string `json:"buyer_user_id" xml:"buyer_user_id"`
		Code           string `json:"code" xml:"code"`
		OutTradeNo     string `json:"out_trade_no" xml:"out_trade_no"`
		TotalAmount    string `json:"total_amount" xml:"total_amount"`
		SubMsg         string `json:"sub_msg" xml:"sub_msg"`
		PointAmount    string `json:"point_amount" xml:"point_amount"`
		SendPayDate    string `json:"send_pay_date" xml:"send_pay_date"`
		OpenID         string `json:"open_id" xml:"open_id"`
		ReceiptAmount  string `json:"receipt_amount" xml:"receipt_amount"`
		Msg            string `json:"msg" xml:"msg"`
		BuyerPayAmount string `json:"buyer_pay_amount" xml:"buyer_pay_amount"`
		InvoiceAmount  string `json:"invoice_amount" xml:"invoice_amount"`
	} `json:"alipay_trade_query_response" xml:"alipay_trade_query_response"`
	Sign string `json:"sign" xml:"sign"`
}
