package wechat

import (
	"encoding/json"
	"github.com/jinlongchen/golang-utilities/http"
	"net/url"
)

func (wx *Wechat) GetWxAcodeUnlimit(scene, page string, appId, appSecret string) ([]byte, error) {
	accessToken, err := wx.GetAccessTokenByClient(appId, appSecret)

	if err != nil {
		return nil, err
	}

	requestURL, _ := url.Parse("https://api.weixin.qq.com/wxa/getwxacodeunlimit")
	parameters := requestURL.Query()
	parameters.Set("access_token", accessToken.AccessToken)
	requestURL.RawQuery = parameters.Encode()

	jData, err := json.Marshal(&struct {
		Scene string `json:"scene"`
		Page  string `json:"page"`
	}{
		Scene: scene,
		Page:  page,
	})
	if err != nil {
		jData = []byte("{}")
	}

	respData, err := http.PostData(
		requestURL.String(),
		"application/json;charset=UTF-8",
		jData,
	)

	return respData, err
}
