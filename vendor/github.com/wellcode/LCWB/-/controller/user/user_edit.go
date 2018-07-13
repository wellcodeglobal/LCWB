package user

import (
	//"fmt"
	db "github.com/wellcode/LCWB/-/model/db"
	session "github.com/wellcode/LCWB/-/model/session"
	"net/http"
	//"strings"
)

func UserEdit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session_val := session.CheckSession(r)
	if session_val != nil {
		if r.FormValue("id-edit") == "" {
			http.Redirect(w, r, "/sign", 302)
		} else {
			id := r.FormValue("id-edit")
			name := r.FormValue("name-edit")
			role := r.FormValue("role-edit")
			email := r.FormValue("email-edit")
			password := r.FormValue("password-edit")
			phone := r.FormValue("phone-edit")
			address := r.FormValue("address-edit")
			company := r.FormValue("company-edit")
			var msg string
			err := db.UpdateUser(id, name, phone, address, company, email, password, role)
			if err == nil {
				msg = "User with email \"" + email + "\" was updated!"
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
