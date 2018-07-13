package preview

import (
	//"fmt"
	session "github.com/wellcode/LCWB/-/model/session"
	//user "github.com/wellcode/LCWB/-/model/user"
	file "github.com/wellcode/LCWB/-/model/file"
	html "github.com/wellcode/LCWB/-/model/html"
	"html/template"
	"net/http"
)

var t *template.Template

func GenerateFile(web_id, web_name string) {
	path := "-/file/" + web_name
	file.CreatePreviewFolder(path, web_name)
	index, css, js := html.GetUserWebPartCode(web_id)
	file.CreateFile(path + "/index.html")
	file.CreateFile(path + "/css/style.css")
	file.CreateFile(path + "/js/script.js")

	//write to file
	file.WriteHTML(path+"/index.html", index)
	file.WriteHTML(path+"/css/style.css", css)
	file.WriteHTML(path+"/js/script.js", js)
}

func Preview(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session_val := session.CheckSession(r)
	if session_val != nil {
		web_id := r.FormValue("web_id")
		web_name := r.FormValue("web_name")
		GenerateFile(web_id, web_name)
	} else {
		http.Redirect(w, r, "/sign", 302)
	}

}
