package banner

import (
	"github.com/arsham/figurine"
	"os"
)

func Print(s string) {
	figurine.Write(os.Stdout, s, "Banner.flf")
}
