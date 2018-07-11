package edithtml

import (
	"fmt"
	"github.com/wellcode/LCWB/-/config"
	database "github.com/wellcode/LCWB/-/model/db"
	html_part "github.com/wellcode/LCWB/-/model/html"
	session "github.com/wellcode/LCWB/-/model/session"
	user "github.com/wellcode/LCWB/-/model/user"
	"html/template"
	"net/http"
	"strconv"
)

var t *template.Template

func EditHTML(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.FormValue("html_id") != "" {
		session_val := session.CheckSession(r)
		if session_val != nil {
			t, _ = template.ParseFiles(
				"-/view/partial/layout.html",
				"-/view/partial/base/customer/customer_navbar.html",
				"-/view/partial/edit/edit_html.html",
			)
			session_str := fmt.Sprint(session_val)
			data := user.GetUserData(session_str)
			html_id := r.FormValue("html_id")
			types := r.FormValue("type")
			types_int, _ := strconv.Atoi(types)
			previous, html_part_list := html_part.GetHTMLPartListType(session_str, html_id, types)
			if types_int != 1 {
				new_part := r.FormValue("new_part")
				database.UpdateUserHTMLPart(html_id, new_part, r.FormValue("previous"))
			}
			data["previous"] = previous
			data["BaseURL"] = config.Base_URL
			data["html_id"] = html_id
			data["type"] = types_int + 1

			data["html_part_list"] = html_part_list

			if types_int < 6 {
				data["submit_link"] = template.HTML("edit/html")
			} else {
				data["submit_link"] = template.HTML("weblist")
			}
			t.ExecuteTemplate(w, "layout", data)
		} else {
			http.Redirect(w, r, "/sign", 302)
		}
	} else {
		http.Redirect(w, r, "/sign", 302)
	}
}
