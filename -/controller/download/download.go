package download

import (
	"fmt"
	preview "github.com/wellcode/LCWB/-/controller/preview"
	file "github.com/wellcode/LCWB/-/model/file"
	session "github.com/wellcode/LCWB/-/model/session"
	user "github.com/wellcode/LCWB/-/model/user"
	//"io"
	"net/http"
	"os"
)

func Download(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session_val := session.CheckSession(r)
	if session_val != nil {
		sesion_str := fmt.Sprint(session_val)
		web_id := r.FormValue("web_id")
		web_name := r.FormValue("web_name")
		data := user.GetUserData(sesion_str)
		check := true
		if data["role"] != "Admin" {
			if !user.CheckUserWeb(sesion_str, web_id) {
				check = false
			}
		}
		if check {
			preview.GenerateFile(web_id, web_name)
			if web_id == "" && web_name == "" {
				http.Redirect(w, r, "/sign", 302)
			} else {
				filename := web_name + ".zip"
				path := "-/file/" + web_name
				output := "-/file/archieves/" + filename
				file.ZipFile(path, output)
				DownloadFile(w, r, output, filename)
			}
		} else {
			http.Redirect(w, r, "/web/list", 302)
		}
	} else {
		http.Redirect(w, r, "/sign", 302)
	}
}

func DownloadFile(w http.ResponseWriter, r *http.Request, path, filename string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), f)
}
