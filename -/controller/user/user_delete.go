package user

import (
	"fmt"
	db "github.com/wellcode/LCWB/-/model/db"
	session "github.com/wellcode/LCWB/-/model/session"
	user "github.com/wellcode/LCWB/-/model/user"
	"net/http"
)

func UserDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session_val := session.CheckSession(r)
	if session_val != nil {
		if r.FormValue("id") == "" {
			http.Redirect(w, r, "/sign", 302)
		} else {
			id := r.FormValue("id")
			var msg string
			data_user := user.GetUserDetail(id)
			err := db.DeleteUser(id)
			if err == nil {
				msg = "User with email \"" + fmt.Sprint(data_user["email"]) + "\" was deleted!"
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
