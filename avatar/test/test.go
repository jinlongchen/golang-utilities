package main

import (
	"bytes"
	"fmt"
	"github.com/jinlongchen/golang-utilities/avatar"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("/Users/chenjinlong/work/go/src/github.com/jinlongchen/golang-utilities/avatar/test/t.prof")
	if err != nil {
		log.Fatalln(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 10000; i++ {

		fileName := fmt.Sprintf("/Users/chenjinlong/work/go/src/github.com/jinlongchen/golang-utilities/avatar/test/png/%d.png", i)
		//if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		//}
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			img, err := avatar.GenerateFromUsername(fmt.Sprintf("%d", i)) //50	  22600245 ns/op
			if err != nil {
				log.Fatal(err.Error())
			}
			var w bytes.Buffer
			err = png.Encode(&w, img)

			if err != nil {
				log.Fatal(err.Error())
			}

			data = w.Bytes()
			err = ioutil.WriteFile(fileName, data, 0666)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}
	//for i := 0; i < 1000; i++ {
	//	//img ,err :=  govatar.GenerateFromUsername(govatar.MALE, fmt.Sprintf("%d", i)) //20	  57610322 ns/op
	//	img, err := avatar.GenerateFromUsername(fmt.Sprintf("%d", i)) //50	  22600245 ns/op
	//	//200	   6216791 ns/op
	//	//300	   4559934 ns/op
	//	if err != nil {
	//		log.Fatal(err.Error())
	//	}
	//	var w bytes.Buffer
	//	err = png.Encode(&w, img)
	//
	//	if err != nil {
	//		log.Fatal(err.Error())
	//	}
	//
	//	ret := w.Bytes()
	//
	//	err = ioutil.WriteFile("/Users/chenjinlong/work/go/src/github.com/jinlongchen/golang-utilities/avatar/test/test.png", ret, 0666)
	//	if err != nil {
	//		log.Fatal(err.Error())
	//	}
	//}
}
