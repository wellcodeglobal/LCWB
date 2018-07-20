package html

import (
	config "github.com/wellcode/LCWB/-/config"
	file "github.com/wellcode/LCWB/-/model/file"
	"html/template"
	"net/http"
	"strings"
)

var t *template.Template

const (
	BaseURL    = config.Base_URL
	PreviewCSS = "-/view/css/preview.css"
	PreviewJS  = "-/view/js/preview.js"
)

func KatalogTemplate(title, url, path string, folderList []string, w http.ResponseWriter) {
	html, css, js := file.ConcatFile([]string{path})
	file.WriteFile(css, PreviewCSS)
	file.WriteFile(js, PreviewJS)
	url_split := strings.Split(url, "/")
	url_str := ""
	if len(url_split) < 4 {
		url_str = url_split[1]
	} else {
		url_str = url_split[1] + "/" + url_split[2]
	}
	data := map[string]interface{}{
		"BaseURL":  BaseURL,
		"Title":    title,
		"list":     folderList,
		"url":      url_str,
		"html_str": html,
		"css_str":  css,
		"js_str":   js,
	}
	t, _ = template.ParseFiles(
		"-/view/list.html",
	)
	t.ExecuteTemplate(w, "layout", data)
}
