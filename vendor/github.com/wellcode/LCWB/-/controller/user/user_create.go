package user

import (
	//"fmt"
	db "github.com/wellcode/LCWB/-/model/db"
	session "github.com/wellcode/LCWB/-/model/session"
	//"html/template"
	"net/http"
	"strings"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session_val := session.CheckSession(r)
	if session_val != nil {
		if r.FormValue("name") == "" {
			http.Redirect(w, r, "/sign", 302)
		} else {
			name := r.FormValue("name")
			email := r.FormValue("email")
			password := r.FormValue("password")
			phone := r.FormValue("phone")
			address := r.FormValue("address")
			company := r.FormValue("company")
			role := r.FormValue("role")
			err, msg := db.InsertUser(name, email, phone, address, company, role, password)
			check := strings.Contains(msg, "Error")
			if err == nil {
				if check {
					w.WriteHeader(401)
				} else {
					w.WriteHeader(200)
				}
			} else {
				w.WriteHeader(401)
			}
			w.Write([]byte(msg))
		}

	} else {
		http.Redirect(w, r, "/sign", 302)
	}
}
