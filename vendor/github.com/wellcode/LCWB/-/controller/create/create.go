package create

import (
	"fmt"
	config "github.com/wellcode/LCWB/-/config"
	database "github.com/wellcode/LCWB/-/model/db"
	file "github.com/wellcode/LCWB/-/model/file"
	"html/template"
	"net/http"
	"strings"
)

var t *template.Template

func Create(w http.ResponseWriter, r *http.Request, step int) {
	data := map[string]interface{}{
		"BaseURL": config.Base_URL,
	}
	if r.Method == "POST" {
		//save di db
		filename := r.FormValue("templatename")
		id, err := database.InsertWeb(500000, 100000, 25000, filename)
		if err {
			err = database.InsertUserWeb(1, id)
			folder_path := ""

			if step == 1 {
				data["Title"] = "NavBar"
				data["type"] = "1"
				folder_path = "Navigation Bar/Top/"
			} else if step == 2 {
				data["Title"] = "NavBar"
				data["type"] = "2"
				folder_path = "Navigation Bar/Top/"
			} else if step == 3 {

			} else if step == 4 {
			}

			if step < 3 {
				array_title, array_html, array_path := file.GetFileList(".html", folder_path)
				data["array_title"] = array_title
				data["array_html"] = array_html
				data["array_path"] = array_path
				data["Track"] = "track"
				data["step"] = step + 1
				t, _ = template.ParseFiles(
					"-/view/form/form_list.html",
				)
				t.ExecuteTemplate(w, "layout", data)
			} else {
				track_split := strings.Split("track", "!@#$%^&*()_+")
				name := track_split[0]
				arr_path := track_split[1 : len(track_split)-1]
				html, css, js := file.ConcatFile(arr_path)
				fmt.Sprint(html, css, js, name)
			}
		} else {
			fmt.Println("--- INSERT WEB ERR ---")
		}

	} else {
		t, _ = template.ParseFiles(
			"-/view/form/form.html",
		)
		t.ExecuteTemplate(w, "layout", data)
	}
}
