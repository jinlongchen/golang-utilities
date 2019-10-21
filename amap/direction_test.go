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
	//http://restapi.amap.com/v3/distance?origins=104.073830,30.494952&destination=104.065810,30.657468&output=json&key=8568cc10c01a137923e2885a297dbf9f&type=1
	distance, duration, err := Distance(
		104.073830,
		30.494952,
		104.065810,
		30.657468,
		amapKey,
	)
	fmt.Println(distance, duration.Minutes(), err)
}
