package web

import (
	"fmt"
	session "github.com/wellcode/LCWB/-/model/session"
	user "github.com/wellcode/LCWB/-/model/user"
	"html/template"
	"net/http"
)

func WebList(w http.ResponseWriter, r *http.Request) {
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
		}
		t.ExecuteTemplate(w, "layout", data)
	} else {
		http.Redirect(w, r, "/sign", 302)
	}

}
