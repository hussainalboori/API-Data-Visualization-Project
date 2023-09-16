package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./web/static"))))
	log.Println("Connect to our website throught http://Localhost:4000")
	http.ListenAndServe(":4000", nil)
}
