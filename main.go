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
	fmt.Println("Running on " + config.Base_URL + " ...")
	database.Connect()
	addr, err := determineListenAddress()
	if err != nil {
		addr = config.Base_Port
	}
	r := mux.NewRouter()
	r.HandleFunc("/", router.PartialList)
	r.HandleFunc("/home", router.Home)
	r.HandleFunc("/web/detail", router.WebDetail)
	r.HandleFunc("/preview", router.Preview)
	r.HandleFunc("/sign", router.SignIn)
	r.HandleFunc("/navbar/{type}/{pID}", router.Navbar)
	r.HandleFunc("/about/{pID}", router.About)
	r.HandleFunc("/form/{pID}", router.Form)
	r.HandleFunc("/footer/{pID}", router.Footer)
	r.HandleFunc("/login", router.Login)
	r.HandleFunc("/register", router.Register)
	r.HandleFunc("/dashboard", router.Dashboard)
	r.HandleFunc("/create", router.Create)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./-/view/"))))
	http.Handle("/assets/", r)
	http.ListenAndServe(addr, r)
}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
