package file

import (
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ConcatFile(pathS []string) (template.HTML, string, string) {
	var html template.HTML
	var css, js string
	for i := 0; i < len(pathS); i++ {
		filepath.Walk(pathS[i], func(path string, f os.FileInfo, _ error) error {
			if !f.IsDir() {
				r, err := regexp.MatchString(".html", f.Name())
				if err == nil && r {
					html_code := ReadFile(path)
					html_code = strings.Replace(html_code, "<link rel=\"stylesheet\" href=\"css/style.css\">", "", -1)
					html_code = strings.Replace(html_code, "<script type=\"text/javascript\" src=\"/js/js.js\"></script>", "", -1)
					html += "\n" + template.HTML(html_code)
				}
				r, err = regexp.MatchString(".css", f.Name())
				if err == nil && r {
					content := strings.Split(ReadFile(path), "\n\n")
					for i := 0; i < len(content); i++ {
						if !strings.Contains(css, content[i]) {
							css += "\n" + content[i]
						}
					}
				}
				r, err = regexp.MatchString(".js", f.Name())
				if err == nil && r {
					js += "\n" + ReadFile(path)
				}

			}
			return nil
		})
	}
	return html, css, js
}
