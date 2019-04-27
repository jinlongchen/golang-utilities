package crypto

import "testing"

func TestNewKey(t *testing.T) {
	key := NewKey([]byte("1"), []byte("123456"))
	t.Log(key)
}
