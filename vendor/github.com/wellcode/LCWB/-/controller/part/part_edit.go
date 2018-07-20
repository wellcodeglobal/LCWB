package part

import (
	//"fmt"
	db "github.com/wellcode/LCWB/-/model/db"
	session "github.com/wellcode/LCWB/-/model/session"
	"net/http"
	//"strings"
)

func PartEdit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session_val := session.CheckSession(r)
	if session_val != nil {
		if r.FormValue("id-edit") == "" {
			http.Redirect(w, r, "/sign", 302)
		} else {
			id := r.FormValue("id-edit")
			name := r.FormValue("name-edit")
			price := r.FormValue("price-edit")
			html := r.FormValue("html-edit")
			css := r.FormValue("css-edit")
			js := r.FormValue("js-edit")
			types := r.FormValue("types-edit")
			var msg string
			err := db.UpdateHTMLPart(id, name, price, types, html, css, js)
			if err == nil {
				msg = "Part with name \"" + name + "\" was updated!"
				w.WriteHeader(200)
			} else {
				msg = "Something error! Please try again."
				w.WriteHeader(401)
			}
			w.Write([]byte(msg))
		}

	} else {
		http.Redirect(w, r, "/sign", 302)
	}
}
