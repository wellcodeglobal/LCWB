package file

import (
	"fmt"
	"io/ioutil"
)

func WriteHTML(path, data string) {
	d1 := []byte(data)
	err := ioutil.WriteFile(path, d1, 0777)
	if err != nil {
		fmt.Println("Error - WriteFile : ", err)
	}
}
