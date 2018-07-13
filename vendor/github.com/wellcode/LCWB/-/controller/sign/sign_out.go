package sign

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

func SignOut(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("%SESSION%2104%"))
	sessionToken, _ := store.Get(r, "session-token")
	session_val := sessionToken.Values["token"]
	sesion_str := fmt.Sprint(session_val)
	if sesion_str != "" {
		sessionToken.Values["token"] = ""
		sessions.Save(r, w)
	}
	http.Redirect(w, r, "/sign", 302)
}
