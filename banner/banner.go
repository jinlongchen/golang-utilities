package banner

import (
	"github.com/arsham/figurine/figurine"
	_ "github.com/jinlongchen/golang-utilities/banner/statik"
	"os"
)

func Print(s string) {
	_ = figurine.Write(os.Stdout, s, "Bloody.flf")
}
