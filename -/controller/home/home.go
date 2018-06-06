package home

import (
	"fmt"
	"github.com/wellcode/LCWB/-/config"
	session "github.com/wellcode/LCWB/-/model/session"
	user "github.com/wellcode/LCWB/-/model/user"
	"html/template"
	"net/http"
)

var t *template.Template

func Home(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session_val := session.CheckSession(r)
	if session_val != nil {
		t, _ = template.ParseFiles(
			"-/view/home.html",
			"-/view/partial/table.html",
		)
		sesion_str := fmt.Sprint(session_val)
		data := user.GetUserData(sesion_str)
		data["BaseURL"] = config.Base_URL
		data["weblist"] = template.JS(user.GetUserWebList(sesion_str))
		t.ExecuteTemplate(w, "layout", data)
	} else {
		http.Redirect(w, r, "/sign", 302)
	}

}
