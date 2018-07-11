package home

import (
	"fmt"
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
		sesion_str := fmt.Sprint(session_val)
		data := user.GetUserData(sesion_str)
		if data["role"] == "Customer" {
			data["weblist"] = template.JS(user.GetUserWebList(sesion_str))
			t, _ = template.ParseFiles(
				"-/view/partial/layout.html",
				"-/view/partial/base/customer/customer_navbar.html",
				"-/view/partial/table/customer_web_table.html",
			)
		} else {
			data["userlist"] = template.JS(user.GetUserList())
			data["rolelist"] = user.GetUserRole()
			t, _ = template.ParseFiles(
				"-/view/partial/layout.html",
				"-/view/partial/base/admin/admin_navbar.html",
				"-/view/partial/table/admin_user_list_table.html",
				"-/view/partial/CRUD/crud_user.html",
			)
		}
		t.ExecuteTemplate(w, "layout", data)
	} else {
		http.Redirect(w, r, "/sign", 302)
	}
}
