package main

import (
	"goauthorizer/db"
	"goauthorizer/web"
	"log"
	"net/http"
)

func main() {
	dbi, err := db.ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer dbi.Close()
	log.Printf("Привет! %s \n", dbi)

	http.HandleFunc("/", web.MainPage)
    http.ListenAndServe(":8080", nil)

}
