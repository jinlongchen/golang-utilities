package main

import (
    "fmt"
    "github.com/jinlongchen/golang-utilities/rand"
    "os"
    "strconv"
)

func main() {
    arg1 := "16"
    if len(os.Args) > 1 {
        arg1 = os.Args[1]
    }
    n, err := strconv.Atoi(arg1)
    if err != nil {
        n = 16
    }
    fmt.Println(rand.GetNonceString(n))
}
