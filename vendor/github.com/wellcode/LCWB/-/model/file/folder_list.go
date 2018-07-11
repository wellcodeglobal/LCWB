package file

import (
	"os"
	"path/filepath"
	"strings"
)

func FolderList(pathS string) []string {
	var list []string
	parent_folder_split := strings.Split(pathS, "/")
	parent_folder := ""
	if len(parent_folder_split) > 2 {
		parent_folder = parent_folder_split[1]
	}else{
		parent_folder = parent_folder_split[0]
	}
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if f.IsDir() {
			if f.Name() != "css" && f.Name() != "js" && f.Name() != parent_folder  && f.Name()!= "images"{
				list = append(list, f.Name())
			}
		}
		return nil
	})
	return list
}
