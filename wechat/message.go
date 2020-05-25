/*
 * Copyright (c) 2020. Jinlong Chen.
 */

package wechat

import (
	"fmt"
	"github.com/jinlongchen/golang-utilities/errors"
	"github.com/jinlongchen/golang-utilities/http"
	"github.com/jinlongchen/golang-utilities/json"
	"github.com/jinlongchen/golang-utilities/log"
)

// SubscribeMessageData 订阅消息模板数据
type SubscribeMessageData map[string]SubscribeMessageDataValue

type SubscribeMessageDataValue struct {
	Value string `json:"value"`
}
type SubscribeMessage struct {
	ToUser     string               `json:"touser"`
	TemplateID string               `json:"template_id"`
	Page       string               `json:"page,omitempty"`
	Data       SubscribeMessageData `json:"data"`
}
type SubscribeMessageResponse struct {
	ErrCode int    `json:"errcode"` // 	错误码
	ErrMSG  string `json:"errmsg"`  // 	错误描述
}

func (wx *Wechat) SendSubscribeMessage(
	appID, appSecret string,
	templateID string,
	toOpenID string,
	page string,
	data SubscribeMessageData,
) (*SubscribeMessageResponse, error) {
	accessToken,err := wx.GetAccessTokenByClient(
		appID, //wx.config.GetString("wechat.appId"),
		appSecret, //wx.config.GetString("wechat.appSecret"),
	)
	if err != nil {
		return nil ,err
	}
	if accessToken == nil || accessToken.AccessToken == "" {
		return nil, errors.New("GetAccessTokenByClient error")
	}
	sendMsgUrl := fmt.Sprintf(`https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s`, accessToken.AccessToken)

	sendM := &SubscribeMessage{
		ToUser:     toOpenID,
		TemplateID: templateID,
		Page:       page,
		Data:       data,
	}

	log.Infof("send subscribe msg: %s", sendMsgUrl)
	log.Infof("send subscribe msg: %s", string(json.ShouldMarshal(sendM)))
	httpData, err := http.PostData(sendMsgUrl, "application/json", json.ShouldMarshal(sendM))
	log.Infof("send subscribe msg: %v", string(httpData))
	if err != nil {
		return nil, err
	}
	ret := &SubscribeMessageResponse{}
	err = json.Unmarshal(httpData, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
