/*
 * Copyright (c) 2020. Jinlong Chen.
 */

package text

import (
	"testing"
)

func TestSprintf(t *testing.T) {
	println(Sprintf(`{1:s} {1:s}`, "0", "1"))
}
