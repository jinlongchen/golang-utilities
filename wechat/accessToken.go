package wechat

import (
	"errors"
	"github.com/jinlongchen/golang-utilities/http"
	"github.com/jinlongchen/golang-utilities/log"
	"net/url"
	"time"
)

type AccessTokenResult struct {
	Errcode int    `json:"errcode,omitempty" bson:"-"`
	Errmsg  string `json:"errmsg,omitempty" bson:"-"`

	AccessToken string `json:"access_token,omitempty" bson:"accessToken,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty" bson:"expiresIn,omitempty"`
	RfreshToken string `json:"rfresh_token,omitempty" bson:"rfreshToken,omitempty"`

	OpenId string `json:"openid,omitempty" bson:"openid,omitempty"`
	Scope  string `json:"scope,omitempty" bson:"scope,omitempty"`
}

func (wx *Wechat) GetAccessTokenByCode(code string) (*AccessTokenResult, error) {
	requestURL, _ := url.Parse("https://api.weixin.qq.com/sns/oauth2/access_token")
	parameters := requestURL.Query()

	parameters.Set("appid", wx.config.GetString("wechat.appId"))
	parameters.Set("secret", wx.config.GetString("wechat.appSecret"))
	parameters.Set("code", code)
	parameters.Set("grant_type", "authorization_code")

	requestURL.RawQuery = parameters.Encode()

	ret := &AccessTokenResult{}

	err := http.GetJSON(requestURL.String(), ret)

	if err != nil {
		return nil, err
	}
	if ret.Errcode != 0 {
		return nil, errors.New(ret.Errmsg)
	}
	return ret, nil
}

func (wx *Wechat) GetAccessTokenByClient(appId, appSecret string) (*AccessTokenResult, error) {
	ret := &AccessTokenResult{}
	cacheKey := wx.config.GetString("application.name") + ".wx_access_token." + appId
	err := wx.cache.Get(cacheKey, ret)
	if err == nil && ret != nil {
		return ret, nil
	}
	if err != nil {
		log.Errorf("%s GetAccessTokenByClient appid err:%s", wx.config.GetString("application.name"), err.Error())
	}

	getTokenURL, _ := url.Parse("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential")
	parameters := getTokenURL.Query()

	log.Debugf("%s GetAccessTokenByClient appid:%s", wx.config.GetString("application.name"), wx.config.GetString("wechat.appId"))

	parameters.Set("appid", appId)
	parameters.Set("secret", appSecret)

	getTokenURL.RawQuery = parameters.Encode()

	err = http.GetJSON(getTokenURL.String(), ret)
	if err != nil {
		return nil, err
	}
	if ret.Errcode != 0 {
		return nil, errors.New(ret.Errmsg)
	}
	if ret.ExpiresIn > 30 {
		ret.ExpiresIn = ret.ExpiresIn - 30
	}
	_ = wx.cache.Set(cacheKey, ret, time.Second*time.Duration(ret.ExpiresIn))
	return ret, nil
}
