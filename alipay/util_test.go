/*
 * Copyright (c) 2019. Qing Cheng Technology Co., Ltd.
 */

package alipay

import "testing"

func TestAliEscape(t *testing.T) {
    x := `{"out_trade_no":"32wahu02kyzokt06","product_code":"QUICK_MSECURITY_PAY","subject":"30元","timeout_express":"30m","total_amount":"0.01"}`
    x = aliEscape(x)
    println(x)
}
