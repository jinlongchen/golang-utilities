package file

import (
	"errors"
	"os"
)

// Exists https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
func Exists(filename string) (bool, error) {
	if _, err := os.Stat(filename); err == nil {
		// exists
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		// does *not* exist
		return false, nil
	} else {
		// Schrodinger: file may or may not exist. See err for details.

		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence

		return false, err
	}
}
