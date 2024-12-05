package file

import (
	"io/ioutil"
	"os"
)

// ReadFile reads the content of the file at the given path.
func ReadFile(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// WriteFile writes the given content to the file at the given path.
func WriteFile(path string, content string) error {
	return ioutil.WriteFile(path, []byte(content), 0644)
}

// DeleteFile deletes the file at the given path.
func DeleteFile(path string) error {
	return os.Remove(path)
}
