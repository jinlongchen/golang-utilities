package filepath

import "testing"

func TestExpandUser(t *testing.T) {
	s := `$HOME/work/go/src/qcse.com/jtb/main/data/`
	s = ExpandUser(s)
	println(s)
}
