package sign

import (
	"github.com/gorilla/sessions"
	config "github.com/wellcode/LCWB/-/config"
	session "github.com/wellcode/LCWB/-/model/session"
	signin "github.com/wellcode/LCWB/-/model/sign"
	"html/template"
	"net/http"
)

var t *template.Template

func SignIn(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("%SESSION%2104%"))
	sessionToken, _ := store.Get(r, "session-token")
	if session.CheckSession(r) != nil {
		http.Redirect(w, r, "/home", 302)
	} else {
		if r.Method == "POST" {
			r.ParseForm()
			if len(r.FormValue("email")) == 0 {
				http.Redirect(w, r, "/sign", 302)
				return
			}
			email := r.FormValue("email")
			if len(r.FormValue("password")) == 0 {
				http.Redirect(w, r, "/sign", 302)
				return
			}
			password := r.FormValue("password")
			isOK, session_token := signin.Login(email, password)
			if isOK {
				sessionToken.Values["token"] = session_token
				sessionToken.Options = &sessions.Options{
					Path:     "/",
					MaxAge:   config.SESSION_TOKEN_TIME,
					HttpOnly: true,
				}
				sessions.Save(r, w)
				http.Redirect(w, r, "/home", 302)
			} else {
				http.Redirect(w, r, "/sign", 302)
			}
		} else {
			t, _ = template.ParseFiles(
				"-/view/login.html",
			)
			data := map[string]interface{}{
				"BaseURL": config.Base_URL,
			}
			t.ExecuteTemplate(w, "layout", data)

		}
	}

}
