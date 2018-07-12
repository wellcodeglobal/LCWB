package part

import (
	"fmt"
	html "github.com/wellcode/LCWB/-/model/html"
	session "github.com/wellcode/LCWB/-/model/session"
	user "github.com/wellcode/LCWB/-/model/user"
	"html/template"
	"net/http"
)

var t *template.Template

func PartList(w http.ResponseWriter, r *http.Request) {
	session_val := session.CheckSession(r)
	if session_val != nil {
		sesion_str := fmt.Sprint(session_val)
		data := user.GetUserData(sesion_str)
		if data["role"] == "Customer" {
			http.Redirect(w, r, "/sign", 302)
		} else {
			data["partlist"] = template.JS(html.GetAllHTMLPartList())
			data["parttypelist"] = html.GetHTMLPartType()
			t, _ = template.ParseFiles(
				"-/view/partial/layout.html",
				"-/view/partial/base/admin/admin_navbar.html",
				"-/view/partial/table/admin_part_list_table.html",
				"-/view/partial/CRUD/crud_part.html",
			)
		}
		t.ExecuteTemplate(w, "layout", data)
	} else {
		http.Redirect(w, r, "/sign", 302)
	}

}
