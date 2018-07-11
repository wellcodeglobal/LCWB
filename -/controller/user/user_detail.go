package user

import (
	"fmt"
	session "github.com/wellcode/LCWB/-/model/session"
	user "github.com/wellcode/LCWB/-/model/user"
	"html/template"
	"net/http"
)

var t *template.Template

func UserDetail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session_val := session.CheckSession(r)
	if session_val != nil {
		sesion_str := fmt.Sprint(session_val)
		data := user.GetUserData(sesion_str)
		if r.FormValue("user_id") == "" {
			http.Redirect(w, r, "/sign", 302)
		} else {
			user_id := r.FormValue("user_id")
			user_data := user.GetUserDetail(user_id)
			data["weblist"] = template.JS(user.GetUserWebListId(user_id))
			data["user_id"] = user_data["id"]
			data["user_name"] = user_data["name"]
			data["user_phone"] = user_data["phone"]
			data["user_email"] = user_data["email"]
			data["user_address"] = user_data["address"]
			data["user_role"] = user_data["role"]
			data["user_company"] = user_data["company"]
			t, _ = template.ParseFiles(
				"-/view/partial/layout.html",
				"-/view/partial/base/admin/admin_navbar.html",
				"-/view/partial/detail/user_detail.html",
			)
			t.ExecuteTemplate(w, "layout", data)
		}

	} else {
		http.Redirect(w, r, "/sign", 302)
	}

}
