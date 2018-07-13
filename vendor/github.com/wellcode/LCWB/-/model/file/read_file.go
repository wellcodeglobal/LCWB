package file

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(pathS string) string {
	var files string
	file, _ := os.Open(pathS)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		files += fmt.Sprintln(scanner.Text())
	}
	return files
}
