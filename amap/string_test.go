package amap

import (
	"encoding/json"
	"testing"
)

type amapTip struct {
	Name     String `json:"name" xml:"name"`
	District String `json:"district" xml:"district"`
	Adcode   String `json:"adcode" xml:"adcode"`
	Location String `json:"location" xml:"location"`
	Address  String `json:"address" xml:"address"`
	Typecode String `json:"typecode" xml:"typecode"`
	City     String `json:"city" xml:"city"`
	ID       String `json:"id" xml:"id"`
}
type amapInputtipsResponse struct {
	Status   String `json:"status" xml:"status"`
	Count    String `json:"count" xml:"count"`
	Info     String `json:"info" xml:"info"`
	Infocode String `json:"infocode" xml:"infocode"`
	Tips     []amapTip   `json:"tips" xml:"tips"`
}

func TestString_MarshalJSON(t *testing.T) {
	amapRespData := []byte( `{"status":[],"count":"10","info":"OK","infocode":"10000","tips":[{"id":"BV10069526","name":"天府广场(地铁站)","district":"四川省成都市青羊区","adcode":"510105","location":"104.065751,30.657453","address":"1号线;2号线","typecode":"150500","city":[]},{"id":"B001C7WEYU","name":"天府广场","district":"四川省成都市青羊区","adcode":"510105","location":"104.065851,30.657424","address":"人民南路一段86号","typecode":"110105","city":[]},{"id":"BV10061425","name":"天府广场东站(公交站)","district":"四川省成都市青羊区","adcode":"510105","location":"104.067459,30.657587","address":"16路;43路;45路;53路;61路;64路;夜间12路;夜间5路;夜间8路;机场专线2号线","typecode":"150700","city":[]},{"id":"B0FFHXMBBS","name":"天府广场(H口)","district":"四川省成都市青羊区","adcode":"510105","location":"104.065712,30.657763","address":"人民中路一段天府广场B1层附近","typecode":"120000","city":[]},{"id":"B001C80SKH","name":"城市之心","district":"四川省成都市青羊区","adcode":"510105","location":"104.065249,30.655953","address":"人民南路一段86号","typecode":"120201","city":[]},{"id":"B001C7X7HB","name":"成都天府广场停车场","district":"四川省成都市青羊区","adcode":"510105","location":"104.064416,30.657384","address":"人民东路天府广场B2层","typecode":"150904","city":[]},{"id":"B001C7USC2","name":"四川科技馆","district":"四川省成都市青羊区","adcode":"510105","location":"104.065787,30.659866","address":"人民中路一段16号","typecode":"140600","city":[]},{"id":"B001C7WER3","name":"成都城市名人酒店","district":"四川省成都市青羊区","adcode":"510105","location":"104.065452,30.654159","address":"人民南路一段122-124号","typecode":"100102","city":[]},{"id":"B001C7XWUG","name":"西御大厦B座","district":"四川省成都市青羊区","adcode":"510105","location":"104.062245,30.656988","address":"西御街8号","typecode":"120201","city":[]},{"id":"B0FFH33FVA","name":"天府广场(西入口)","district":"四川省成都市青羊区","adcode":"510105","location":"104.064305,30.657470","address":[],"typecode":"991500","city":[]}]}`)
	amapRespModel := &amapInputtipsResponse{}
	err := json.Unmarshal(amapRespData, amapRespModel)
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(amapRespModel)
	if err != nil {
		panic(err)
	}
	println("Status:", amapRespModel.Status.ToString())
	println("Status 2:", string(amapRespModel.Status))
	println(string(data))

	//resp := &amapInputtipsResponse{
	//	//Status:[]byte(`[]`),
	//}
	//str := `{"status":[],"count":"10","infocode":"10000","tips":[{"id":"BV10069526","name":"天府广场(地铁站)","district":"四川省成都市青羊区","adcode":"510105","location":"104.065751,30.657453","address":"1号线;2号线","typecode":"150500","city":[]},{"id":"B001C7WEYU","name":"天府广场","district":"四川省成都市青羊区","adcode":"510105","location":"104.065851,30.657424","address":"人民南路一段86号","typecode":"110105","city":[]},{"id":"BV10061425","name":"天府广场东站(公交站)","district":"四川省成都市青羊区","adcode":"510105","location":"104.067459,30.657587","address":"16路;43路;45路;53路;61路;64路;夜间12路;夜间5路;夜间8路;机场专线2号线","typecode":"150700","city":[]},{"id":"B0FFHXMBBS","name":"天府广场(H口)","district":"四川省成都市青羊区","adcode":"510105","location":"104.065712,30.657763","address":"人民中路一段天府广场B1层附近","typecode":"120000","city":"成都1"},{"id":"B001C80SKH","name":"城市之心","district":"四川省成都市青羊区","adcode":"510105","location":"104.065249,30.655953","address":"人民南路一段86号","typecode":"120201","city":[]},{"id":"B001C7X7HB","name":"成都天府广场停车场","district":"四川省成都市青羊区","adcode":"510105","location":"104.064416,30.657384","address":"人民东路天府广场B2层","typecode":"150904","city":[]},{"id":"B001C7USC2","name":"四川科技馆","district":"四川省成都市青羊区","adcode":"510105","location":"104.065787,30.659866","address":"人民中路一段16号","typecode":"140600","city":[]},{"id":"B001C7WER3","name":"成都城市名人酒店","district":"四川省成都市青羊区","adcode":"510105","location":"104.065452,30.654159","address":"人民南路一段122-124号","typecode":"100102","city":[]},{"id":"B001C7XWUG","name":"西御大厦B座","district":"四川省成都市青羊区","adcode":"510105","location":"104.062245,30.656988","address":"西御街8号","typecode":"120201","city":[]},{"id":"B0FFH33FVA","name":"天府广场(西入口)","district":"四川省成都市青羊区","adcode":"510105","location":"104.064305,30.657470","address":[],"typecode":"991500","city":[]}]}`
	//err := json.Unmarshal([]byte(str), resp)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("the string 1 is (%s)\n", string(resp.Status))
	//fmt.Printf("the status 2 is (%s)\n", resp.Status.ToString())
	//data, err := json.Marshal(resp)
	//if err != nil {
	//	panic(err)
	//}
	//println(string(data))
}

func TestString_MarshalJSON2(t *testing.T) {
	amapRespData:= []byte(`{"status":"1","count":"10","info":"OK","infocode":"10000","tips":[{"id":"BV10069526","name":"天府广场(地铁站)","district":"四川省成都市青羊区","adcode":"510105","location":"104.065751,30.657453","address":"1号线;2号线","typecode":"150500","city":[]},{"id":"B001C7WEYU","name":"天府广场","district":"四川省成都市青羊区","adcode":"510105","location":"104.065851,30.657424","address":"人民南路一段86号","typecode":"110105","city":[]},{"id":"BV10061425","name":"天府广场东站(公交站)","district":"四川省成都市青羊区","adcode":"510105","location":"104.067459,30.657587","address":"16路;43路;45路;53路;61路;64路;夜间12路;夜间5路;夜间8路;机场专线2号线","typecode":"150700","city":[]},{"id":"B0FFHXMBBS","name":"天府广场(H口)","district":"四川省成都市青羊区","adcode":"510105","location":"104.065712,30.657763","address":"人民中路一段天府广场B1层附近","typecode":"120000","city":[]},{"id":"B001C80SKH","name":"城市之心","district":"四川省成都市青羊区","adcode":"510105","location":"104.065249,30.655953","address":"人民南路一段86号","typecode":"120201","city":[]},{"id":"B001C7X7HB","name":"成都天府广场停车场","district":"四川省成都市青羊区","adcode":"510105","location":"104.064416,30.657384","address":"人民东路天府广场B2层","typecode":"150904","city":[]},{"id":"B001C7USC2","name":"四川科技馆","district":"四川省成都市青羊区","adcode":"510105","location":"104.065787,30.659866","address":"人民中路一段16号","typecode":"140600","city":[]},{"id":"B001C7WER3","name":"成都城市名人酒店","district":"四川省成都市青羊区","adcode":"510105","location":"104.065452,30.654159","address":"人民南路一段122-124号","typecode":"100102","city":[]},{"id":"B001C7XWUG","name":"西御大厦B座","district":"四川省成都市青羊区","adcode":"510105","location":"104.062245,30.656988","address":"西御街8号","typecode":"120201","city":[]},{"id":"B0FFH33FVA","name":"天府广场(西入口)","district":"四川省成都市青羊区","adcode":"510105","location":"104.064305,30.657470","address":[],"typecode":"991500","city":[]}]}`)
	amapRespModel := &amapInputtipsResponse{}
	err := json.Unmarshal(amapRespData, amapRespModel)
	if err != nil {
		panic(err)
	}

	println("status:", string(amapRespModel.Status))
	println("status:", amapRespModel.Status.ToString())

}
