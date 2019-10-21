/*
 * Copyright (c) 2019. Qing Cheng Technology Co., Ltd.
 */

package amap

import (
	"fmt"
	"testing"
)

func TestDistance(t *testing.T) {
	amapKey := ""
	distance, duration, err := Distance(
		104.073830,
		30.494952,
		104.065810,
		30.657468,
		amapKey,
	)
	fmt.Println(distance, duration.Minutes(), err)
}
