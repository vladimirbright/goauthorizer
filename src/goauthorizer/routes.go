package main

import (
	"goauthorizer/web"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	log.Println("Setup routes")
	log.Println("Override /assets/ to your webserver!!!")

	r := mux.NewRouter()
	r.HandleFunc("/", web.MainPage)
	r.HandleFunc("/signin/", web.LoginPage)
	r.HandleFunc("/signup/", web.RegisterPage)

	http.Handle("/", r)

	// Overrite this location with normal webserver!
	http.Handle("/assets/",
				http.StripPrefix("/assets/",
								 http.FileServer(http.Dir("src/goauthorizer/assets"))))

	return r
}
