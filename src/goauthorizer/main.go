package main

import (
	"goauthorizer/db"
	"log"
	"net/http"
	"os"
)

func main() {
	dbi, err := db.ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer dbi.Close()

	SetupRoutes()
	listen := os.Getenv("SERVER_LISTEN")
	if listen == "" {
		log.Fatal("Setup enviroment variable SERVER_LISTEN in format IP:PORT")
	}
	log.Printf("Launched on %s !", listen)
    http.ListenAndServe(listen, nil)
}
