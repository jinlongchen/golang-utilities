/*
 * Copyright (c) 2019. Qing Cheng Technology Co., Ltd.
 */

package json

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

type message struct {
	Code    string              `json:"code,omitempty"`
	Message string              `json:"message,omitempty"`
	Data    jsoniter.RawMessage `json:"data,omitempty"`
}

func TestShouldMarshal(t *testing.T) {
	fmt.Println(string(ShouldMarshal(&message{
		Code: "123",
		Data: ShouldMarshal("456"),
	})))
}

func TestShouldMarshal2(t *testing.T) {
	fmt.Println(string(ShouldMarshal(&map[string]interface{}{
	})))
}
