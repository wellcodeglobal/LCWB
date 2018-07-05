package file

import (
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetFileList(ext, pathS string) ([]string, []template.HTML, []string) {
	var files []string
	var html []template.HTML
	var arr_path []string
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				html_code := ReadFile(path)
				html_code = strings.Replace(html_code, "<link rel=\"stylesheet\" href=\"css/style.css\">", "", -1)
				html_code = strings.Replace(html_code, "<script type=\"text/javascript\" src=\"/js/js.js\"></script>", "", -1)
				html = append(html, template.HTML(html_code))
			}
		} else {
			if f.Name() != "css" && f.Name() != "js" && f.Name()!="images"{
				files = append(files, f.Name())
				arr_path = append(arr_path, path)
			}
		}
		return nil
	})
	return files, html, arr_path[1:]
}
