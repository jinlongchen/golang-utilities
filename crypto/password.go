package crypto

import (
    "math/rand"
    "time"
)

var (
    letters = []rune("~!@#$%^&*(),./';<>?:abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    r       = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func GeneratePassword(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[r.Intn(len(letters))]
    }
    return string(b)
}
