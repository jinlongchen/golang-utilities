package main

import (
    "encoding/hex"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"

    "github.com/jinlongchen/golang-utilities/compress"
)

func main() {
    str2 := "	categoryCount = []int{"
    str := `package avatar

var (
	avatarBinData = [][]string{
`
    lastCategory := ""
    count := 0
    root := "/Users/chenjinlong/work/go/src/github.com/jinlongchen/golang-utilities/avatar/gen2/"
    _ = filepath.Walk(root, func(sub string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }
        if !strings.HasSuffix(sub, ".png") {
            return nil
        }
        data, err := ioutil.ReadFile(sub)
        if err != nil {
            log.Fatalln(err.Error())
        }
        oLen := len(data)
        data, err = compress.Gzip(data)
        if err != nil {
            log.Fatalln(err.Error())
        }
        log.Println("olen:", oLen, " newLen:", len(data))
        path, _ := filepath.Rel(root, sub)
        parts := strings.Split(path, "/")
        category := parts[0]
        // name := strings.Replace(parts[1], "@3x.png", "", -1)

        if lastCategory == "" {
            count = 1
            str += `		{ // ` + category + `
` + "			`" + hex.EncodeToString(data) + "`," + " // " + path + `
`
        } else if lastCategory == category {
            count++
            str += "			`" + hex.EncodeToString(data) + "`, // " + path + `
`
        } else {
            str2 += fmt.Sprintf(`
		%d,//%s`, count, lastCategory)
            count = 1
            str += `		},
		{ // ` + category + `
` + "			`" + hex.EncodeToString(data) + "`, // " + path + `
`
        }

        lastCategory = category
        log.Println(path)
        return nil
    })

    str += `		},
	}
`
    str2 += fmt.Sprintf(`
		%d,//%s`, count, lastCategory)

    str2 += `
	}
)
`
    _ = ioutil.WriteFile("/Users/chenjinlong/work/go/src/github.com/jinlongchen/golang-utilities/avatar/bindata.go", []byte(str+str2), 0666)
    println(str + str2)
}
