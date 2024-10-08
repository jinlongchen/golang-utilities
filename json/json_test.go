/*
 * Copyright (c) 2019. Qing Cheng Technology Co., Ltd.
 */

package json

import (
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

type message struct {
	Code    string              `json:"code,omitempty"`
	Message string              `json:"message,omitempty"`
	Data    jsoniter.RawMessage `json:"data,omitempty"`
}

func TestShouldMarshal(t *testing.T) {
	res := string(ShouldMarshal(&message{
		Code: "123",
		Data: ShouldMarshal("456"),
	}))

	if res != `{"code":"123","data":"IjQ1NiI="}` {
		t.Fail()
	}
}

func TestShouldMarshal2(t *testing.T) {
	res := string(ShouldMarshal(&map[string]interface{}{}))
	if res != `{}` {
		t.Fail()
	}
}

type ImageUploadRequest struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Order      int    `json:"order"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	MutedColor string `json:"muted_color"`
	Cover      bool   `json:"cover"`
}

// go test -c -o json.test -gcflags all=-N -l github.com/jinlongchen/golang-utilities/json #gosetup
func TestUnmarshalV2(t *testing.T) {
	fmt.Printf("%v\n", "TestUnmarshalV2")
	if info, err := UnmarshalV2[ImageUploadRequest]([]byte(`{"type":"image","order":null,"width":720,"height":1280,"muted_color":null,"cover":null}`)); err == nil {
		fmt.Printf("%v\n", info)
		fmt.Printf("%v\n", info.Order)
		fmt.Printf("%v\n", string(ShouldMarshal(info)))
	} else {
		fmt.Printf("unmarshal err: %v\n", err)
	}
}
func TestShouldMarshal3(t *testing.T) {

}
