package file

import (
	"io/ioutil"
)

func WriteFile(file, path string) {
	ioutil.WriteFile(path, []byte(""), 0644)
	ioutil.WriteFile(path, []byte(file), 0644)
}
