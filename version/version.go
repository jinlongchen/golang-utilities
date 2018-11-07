package version

import (
	"fmt"
)

var (
	build string
)

func String(app string, binary string) string {
	return fmt.Sprintf("%s v%s (build %s)", app, binary, build)
}
