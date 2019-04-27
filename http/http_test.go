package http
//
//import "testing"
//
//func TestGetJSON(t *testing.T) {
//	info := &UserInfoResult{}
//	err := GetJSON(`https://api.weixin.qq.com/sns/userinfo?access_token=16_Xtgu18-uW4r1zadDM_SCsGPEFprqijskkT1wjm8rNFIeg6De6xDsKrf8awGNANGbHzDQ0xtT0KkL6EImiPpwaw&lang=zh_CN&openid=okq2F0m-iDoEEeZJApDKW4Xu6vpU`, info)
//	if err != nil {
//		t.Fatal(err)
//	}
//	println(info.Nickname)
//	println(info.HeadImgURL)
//}
//type UserInfoResult struct {
//	Errcode int `json:"errcode,omitempty" xml:"errcode,omitempty" bson:"-"`
//	Errmsg  string `json:"errmsg,omitempty" xml:"errmsg,omitempty" bson:"-"`
//
//	City       string `json:"city" bson:"city" xml:"city"`
//	Country    string `json:"country" bson:"country" xml:"country"`
//	HeadImgURL string `json:"headimgurl" bson:"headImgURL" xml:"headimgurl"`
//	Language   string `json:"language" bson:"language" xml:"language"`
//	Nickname   string `json:"nickname" bson:"nickname" xml:"nickname"`
//	Openid     string `json:"openid" bson:"openId" xml:"openid"`
//	Privilege  []string `json:"privilege" bson:"privilege" xml:"privilege"`
//	Province   string `json:"province" bson:"province" xml:"province"`
//	Sex        int `json:"sex" bson:"sex" xml:"sex"`
//}
