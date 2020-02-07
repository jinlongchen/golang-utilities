/*
 * Copyright (c) 2020. Jinlong Chen.
 */

package sms

import (
	"fmt"
	"testing"
)
import "github.com/qiniu/api.v7/v7/auth"

func getManager() *Manager {
	ak := getBitsAppKey()
	sk := getBitsSecretKey()

	qiniuManager := NewManager(auth.New(ak, sk))

	return qiniuManager
}

func TestNewManager(t *testing.T) {
	tID1 := "1217103952852553728"

	qiniuManager := getManager()

	response, err := qiniuManager.SendMessage(MessagesRequest{
		TemplateID: tID1,
		Mobiles:    []string{"13183801710"},
		Parameters: map[string]interface{}{
			"module": "邮件",
			"reason": "被退信",
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("job id: %s\n", response.JobID)
	// 1219176816170766336
}

func TestNewManager2(t *testing.T) {
	jobID := "1219176816170766336"

	qiniuManager := getManager()

	response, err := qiniuManager.QueryMessage(QueryMessageRequest{
		JobID: jobID,
		Mobile: "13183801710",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("messages: %v\n", response)

}
