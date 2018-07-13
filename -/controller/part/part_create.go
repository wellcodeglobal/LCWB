package part

import (
	"fmt"
	db "github.com/wellcode/LCWB/-/model/db"
	session "github.com/wellcode/LCWB/-/model/session"
	//"html/template"
	"net/http"
	"strings"
)

func PartCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session_val := session.CheckSession(r)
	name := r.FormValue("name")
	price := r.FormValue("price")
	types := r.FormValue("types")
	html := r.FormValue("html")
	css := r.FormValue("css")
	js := r.FormValue("js")
	if session_val != nil {
		if r.FormValue("name") == "" || price == "" || types == "" || html == "" || css == "" {
			http.Redirect(w, r, "/sign", 302)
		} else {
			err, msg := db.InsertHTMLPart(name, price, types, html, css, js)
			fmt.Println(msg)
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
