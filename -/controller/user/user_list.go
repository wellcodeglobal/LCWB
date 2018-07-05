package web

import (
	"fmt"
	session "github.com/wellcode/LCWB/-/model/session"
	user "github.com/wellcode/LCWB/-/model/user"
	"html/template"
	"net/http"
)

func UserList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session_val := session.CheckSession(r)
	if session_val != nil {
		sesion_str := fmt.Sprint(session_val)
		data := user.GetUserData(sesion_str)
		if data["role"] == "Customer" {
			http.Redirect(w, r, "/sign", 302)
		} else {
			data["userlist"] = template.JS(user.GetUserList())
			t, _ = template.ParseFiles(
				"-/view/partial/layout.html",
				"-/view/partial/base/admin/admin_navbar.html",
				"-/view/partial/table/admin_user_list_table.html",
			)
		}
		t.ExecuteTemplate(w, "layout", data)
	} else {
		http.Redirect(w, r, "/sign", 302)
	}

}
