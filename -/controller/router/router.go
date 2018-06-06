package router

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	config "github.com/wellcode/LCWB/-/config"
	create "github.com/wellcode/LCWB/-/controller/create"
	home "github.com/wellcode/LCWB/-/controller/home"
	preview "github.com/wellcode/LCWB/-/controller/preview"
	sign "github.com/wellcode/LCWB/-/controller/sign"
	file "github.com/wellcode/LCWB/-/model/file"
	html_exe "github.com/wellcode/LCWB/-/model/html"
)

var t *template.Template

func Home(w http.ResponseWriter, r *http.Request) {
	home.Home(w, r)
}

func Preview(w http.ResponseWriter, r *http.Request) {
	preview.Preview(w, r)
}

func PartialList(w http.ResponseWriter, r *http.Request) {
	t, _ = template.ParseFiles(
		"-/view/all.html",
	)
	data := map[string]interface{}{
		"BaseURL": config.Base_URL,
	}
	t.ExecuteTemplate(w, "layout", data)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	sign.SignIn(w, r)
}

func Create(w http.ResponseWriter, r *http.Request) {
	param_step := r.URL.Query()["step"]
	step_str := "0"
	if param_step != nil {
		step_str = param_step[0]
	}
	step, _ := strconv.Atoi(step_str)
	create.Create(w, r, step)
}

func Navbar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	types := vars["type"]
	id := vars["pID"]
	var folder []string
	var path string
	if types == "top" {
		folder = file.FolderList("partial_html/Navigation Bar/Top")
		path = "partial_html/Navigation Bar/Top/" + id
	} else if types == "left" {
		folder = file.FolderList("partial_html/Navigation Bar/Left")
		path = "partial_html/Navigation Bar/Left/" + id
	}
	title := "NAVBAR - " + strings.ToUpper(types)
	html_exe.KatalogTemplate(title, r.URL.String(), path, folder, w)
}

func About(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["pID"]
	var folder []string
	var path string
	folder = file.FolderList("partial_html/About/")
	path = "partial_html/About/" + id
	title := "About - " + id
	html_exe.KatalogTemplate(title, r.URL.String(), path, folder, w)
}

func Form(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["pID"]
	var folder []string
	var path string
	folder = file.FolderList("partial_html/Form/")
	path = "partial_html/Form/" + id
	title := "Form - " + id
	html_exe.KatalogTemplate(title, r.URL.String(), path, folder, w)
}

func Footer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["pID"]
	var folder []string
	var path string
	folder = file.FolderList("partial_html/Footer/")
	path = "partial_html/Footer/" + id
	title := "Footer - " + id
	html_exe.KatalogTemplate(title, r.URL.String(), path, folder, w)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id := vars["pID"]
	var folder []string
	var path string
	folder = file.FolderList("partial_html/Login/")
	fmt.Sprintln(123)
	path = "Login"
	title := "Login"
	html_exe.KatalogTemplate(title, r.URL.String(), path, folder, w)
}

func Register(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id := vars["pID"]
	var folder []string
	var path string
	folder = file.FolderList("partial_html/Register/")
	fmt.Sprintln(123)
	path = "Register"
	title := "Register"
	html_exe.KatalogTemplate(title, r.URL.String(), path, folder, w)
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id := vars["pID"]
	var folder []string
	var path string
	folder = file.FolderList("partial_html/Dashboard/")
	fmt.Sprintln(123)
	path = "Dashboard"
	title := "Dashboard"
	html_exe.KatalogTemplate(title, r.URL.String(), path, folder, w)
}
