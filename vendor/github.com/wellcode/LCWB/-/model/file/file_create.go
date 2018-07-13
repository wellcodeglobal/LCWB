package file

import (
	"fmt"
	"os"
)

func CreateFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Cannot create file", err)
	}
	defer file.Close()
}
