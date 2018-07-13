package part

import (
	//"fmt"
	db "github.com/wellcode/LCWB/-/model/db"
	session "github.com/wellcode/LCWB/-/model/session"
	"net/http"
)

func PartDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session_val := session.CheckSession(r)
	id := r.FormValue("id")
	name := r.FormValue("name")
	if session_val != nil {
		if id == "" || name == "" {
			http.Redirect(w, r, "/sign", 302)
		} else {
			var msg string
			err := db.DeletePart(id)
			if err == nil {
				msg = "Part with name \"" + name + "\" was deleted!"
				w.WriteHeader(200)
			} else {
				msg = "Something wrong. Please try again later!"
				w.WriteHeader(401)
			}
			w.Write([]byte(msg))
		}

	} else {
		http.Redirect(w, r, "/sign", 302)
	}
}
