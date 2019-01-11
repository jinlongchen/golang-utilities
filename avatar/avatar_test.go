package avatar

import (
	"bytes"
	"github.com/o1egl/govatar"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
	"runtime/pprof"
)

func TestGenerateFromUsername(t *testing.T) {
	img ,err :=  randomAvatar(time.Now().Unix())
	if err != nil {
		t.Fatal(err.Error())
	}
	var w bytes.Buffer
	err = png.Encode(&w, img)

	if err != nil {
		t.Fatal(err.Error())
	}

	ret := w.Bytes()

	err = ioutil.WriteFile("/Users/chenjinlong/work/go/src/github.com/jinlongchen/golang-utilities/avatar/test/test.png", ret, 0666)
	if err != nil {
		t.Fatal(err.Error())
	}
	govatar.GenerateFromUsername(govatar.MALE, "i")
}
func BenchmarkGenerateFromUsername(b *testing.B) {
	f, err := os.Create("/Users/chenjinlong/work/go/src/github.com/jinlongchen/golang-utilities/avatar/test/t.prof")
	if err != nil {
		log.Fatalln(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < b.N; i++ {
		//img ,err :=  govatar.GenerateFromUsername(govatar.MALE, fmt.Sprintf("%d", i)) //20	  57610322 ns/op
		img ,err :=  randomAvatar(time.Now().Unix())									//50	  22600245 ns/op
																						//200	   6216791 ns/op
																						//300	   4559934 ns/op
		if err != nil {
			b.Fatal(err.Error())
		}
		var w bytes.Buffer
		err = png.Encode(&w, img)

		if err != nil {
			b.Fatal(err.Error())
		}

		ret := w.Bytes()

		err = ioutil.WriteFile("/Users/chenjinlong/work/go/src/github.com/jinlongchen/golang-utilities/avatar/test/test.png", ret, 0666)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}