/*
 * Copyright (c) 2019. Jinlong Chen.
 */

package pg

import "testing"

func TestJsonbMap_To(t *testing.T) {
	p := &JsonbMap{
		"a":"123",
		"b":"456",
	}
	var b []byte

	p.To(&b)
	println(string(b))
}
