package web

import (
	"fmt"
	"github.com/wellcode/LCWB/-/config"
	html "github.com/wellcode/LCWB/-/model/html"
	session "github.com/wellcode/LCWB/-/model/session"
	user "github.com/wellcode/LCWB/-/model/user"
	"html/template"
	"net/http"
)

var t *template.Template

func WebDetail(w http.ResponseWriter, r *http.Request) {
	web_id := r.URL.Query()["web"][0]
	user_id := r.URL.Query()["user"][0]
	if web_id == "" && user_id == "" {
		http.Redirect(w, r, "/home", 302)
	} else {
		session_val := session.CheckSession(r)
		if session_val != nil {
			sesion_str := fmt.Sprint(session_val)
			data := user.GetUserData(sesion_str)
			var web_data map[string]interface{}
			if data["role"] == "Admin" {
				web_data = user.GetUserWebDetail("Admin", web_id)
				t, _ = template.ParseFiles(
					"-/view/partial/layout.html",
					"-/view/partial/base/admin/admin_navbar.html",
					"-/view/partial/detail/web_detail.html",
				)
				user_data := user.GetUserDetail(user_id)
				data["user_name"] = user_data["name"]
				data["user_phone"] = user_data["phone"]
				data["user_email"] = user_data["email"]
				data["arr_part"] = html.GetUserWebPartAdmin(web_id)
			} else {
				web_data = user.GetUserWebDetail(sesion_str, web_id)
				t, _ = template.ParseFiles(
					"-/view/partial/layout.html",
					"-/view/partial/base/customer/customer_navbar.html",
					"-/view/partial/detail/web_detail.html",
				)
				//Detail untuk setiap Part
				data["arr_part"] = html.GetUserWebPart(sesion_str, web_id)
				//data["previous"] = user.GetUserHTMLPart(sesion_str, r.FormValue("web_id"), "1")
			}
			if web_data == nil {
				http.Redirect(w, r, "/sign", 302)
			} else {
				data["web_id"] = web_data["web_id"]
				data["base_cost"] = web_data["base_cost"]
				data["maintenance"] = web_data["maintenance"]
				data["service"] = web_data["service"]
				data["domain"] = web_data["domain"]
				data["BaseURL"] = config.Base_URL
				t.ExecuteTemplate(w, "layout", data)
			}

		} else {
			http.Redirect(w, r, "/sign", 302)
		}
	}
}
