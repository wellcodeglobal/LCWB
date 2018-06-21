package web

import (
	"fmt"
	"github.com/wellcode/LCWB/-/config"
	session "github.com/wellcode/LCWB/-/model/session"
	user "github.com/wellcode/LCWB/-/model/user"
	"html/template"
	"net/http"
)

var t *template.Template

func WebDetail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session_val := session.CheckSession(r)
	if session_val != nil {
		t, _ = template.ParseFiles(
			"-/view/home.html",
			"-/view/partial/web_detail.html",
		)
		sesion_str := fmt.Sprint(session_val)
		data := user.GetUserData(sesion_str)
		web_data := user.GetUserWebDetail(sesion_str, r.FormValue("web_id"))
		data["web_id"] = web_data["web_id"]
		data["base_cost"] = web_data["base_cost"]
		data["maintenance"] = web_data["maintenance"]
		data["service"] = web_data["service"]
		data["domain"] = web_data["domain"]
		data["BaseURL"] = config.Base_URL
		t.ExecuteTemplate(w, "layout", data)
	} else {
		http.Redirect(w, r, "/sign", 302)
	}

}
