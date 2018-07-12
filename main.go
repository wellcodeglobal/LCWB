package main

import (
	"fmt"
	"github.com/gorilla/mux"
	config "github.com/wellcode/LCWB/-/config"
	router "github.com/wellcode/LCWB/-/controller/router"
	database "github.com/wellcode/LCWB/-/model/db"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Running on " + config.Base_Port + " ...")
	database.Connect()
	addr, err := determineListenAddress()
	if err != nil {
		addr = config.Base_Port
	}
	r := mux.NewRouter()
	r.HandleFunc("/", router.PartialList)
	r.HandleFunc("/logout", router.Logout)
	r.HandleFunc("/home", router.Home)
	r.HandleFunc("/weblist", router.WebList)
	r.HandleFunc("/user/list", router.UserList)
	r.HandleFunc("/user/delete", router.UserDelete)
	r.HandleFunc("/user/edit", router.UserEdit)
	r.HandleFunc("/user/detail", router.UserDetail)
	r.HandleFunc("/user/create", router.UserCreate)
	r.HandleFunc("/web/list", router.WebList)
	r.HandleFunc("/web/detail", router.WebDetail)
	r.HandleFunc("/part/list", router.PartList)
	r.HandleFunc("/preview", router.Preview)
	r.HandleFunc("/edit/html", router.EditHTML)
	r.HandleFunc("/sign", router.SignIn)
	r.HandleFunc("/navbar/{type}/{pID}", router.Navbar)
	r.HandleFunc("/about/{pID}", router.About)
	r.HandleFunc("/form/{pID}", router.Form)
	r.HandleFunc("/footer/{pID}", router.Footer)
	r.HandleFunc("/login", router.Login)
	r.HandleFunc("/register", router.Register)
	r.HandleFunc("/dashboard", router.Dashboard)
	r.HandleFunc("/download", router.Download)
	r.HandleFunc("/create", router.Create)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./-/view/"))))
	r.PathPrefix("/download/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./-/file/archives/"))))
	http.ListenAndServe(addr, r)
}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
