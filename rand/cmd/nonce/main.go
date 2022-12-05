package main

import (
    "fmt"
    "os"
    "strconv"

    "github.com/jinlongchen/golang-utilities/rand"
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
