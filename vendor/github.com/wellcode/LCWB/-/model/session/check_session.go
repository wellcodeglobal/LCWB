package session

import (
	"fmt"
	"github.com/gorilla/sessions"
	database "github.com/wellcode/LCWB/-/model/db"
	"net/http"
)

func CheckSession(r *http.Request) interface{} {
	var store = sessions.NewCookieStore([]byte("%SESSION%2104%"))
	sessionToken, _ := store.Get(r, "session-token")
	session_val := sessionToken.Values["token"]
	rows, err := database.ExecuteQuery("SELECT * FROM t_user WHERE session_token = $1;", session_val)
	if len(rows) > 0 && err == nil {
		return sessionToken.Values["token"]
	} else {
		if err != nil {
			fmt.Println("ERR : CheckSession - ", err)
		}
		return nil
	}

}
