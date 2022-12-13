package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetJSON(t *testing.T) {
	info := &UserInfoResult{}
	err := GetJSON(`https://api.weixin.qq.com/sns/userinfo?access_token=16_Xtgu18-uW4r1zadDM_SCsGPEFprqijskkT1wjm8rNFIeg6De6xDsKrf8awGNANGbHzDQ0xtT0KkL6EImiPpwaw&lang=zh_CN&openid=akq2F0m-iDoEEeZJApDKW4Xu6vpU`, info)
	if err != nil {
		t.Fatal(err)
	}
	println(info.Nickname)
	println(info.HeadImgURL)
}

func TestGetDataWithHeaders(t *testing.T) {
	_, resp, err := GetDataWithHeaders(`https://api.weixin.qq.com/sns/userinfo?access_token=16_Xtgu18-uW4r1zadDM_SCsGPEFprqijskkT1wjm8rNFIeg6De6xDsKrf8awGNANGbHzDQ0xtT0KkL6EImiPpwaw&lang=zh_CN&openid=akq2F0m-iDoEEeZJApDKW4Xu6vpU`, http.Header{
		"Accept-Encoding": []string{"gzip"},
	})
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", string(resp))
	}
}

func TestGetDataProxy(t *testing.T) {
	data, err := GetData("https://www.google.com")
	if err == nil {
		t.Fail()
		return
	}
	data, err = GetDataProxy("https://www.google.com", "http://127.0.0.1:9900")
	if err != nil {
		t.Fail()
		return
	}
	fmt.Println(string(data))
}

type UserInfoResult struct {
	Errcode int    `json:"errcode,omitempty" xml:"errcode,omitempty" bson:"-"`
	Errmsg  string `json:"errmsg,omitempty" xml:"errmsg,omitempty" bson:"-"`

	City       string   `json:"city" bson:"city" xml:"city"`
	Country    string   `json:"country" bson:"country" xml:"country"`
	HeadImgURL string   `json:"headimgurl" bson:"headImgURL" xml:"headimgurl"`
	Language   string   `json:"language" bson:"language" xml:"language"`
	Nickname   string   `json:"nickname" bson:"nickname" xml:"nickname"`
	Openid     string   `json:"openid" bson:"openId" xml:"openid"`
	Privilege  []string `json:"privilege" bson:"privilege" xml:"privilege"`
	Province   string   `json:"province" bson:"province" xml:"province"`
	Sex        int      `json:"sex" bson:"sex" xml:"sex"`
}
