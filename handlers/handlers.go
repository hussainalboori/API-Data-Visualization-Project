package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errorshandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		Errorshandler(w, http.StatusMethodNotAllowed)
		return
	}
	template, err := template.ParseFiles("./template/index.html")
	if err != nil {
		log.Println(err.Error())
		Errorshandler(w, http.StatusInternalServerError)
		return
	}
	band, errr := JsonArtists()

	if errr != nil {
		Errorshandler(w, http.StatusInternalServerError)
		return
	}
	
}

func Errorshandler(w http.ResponseWriter, status int) {
	template, err := template.ParseFiles("./template/error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var Result Errors
	Result.Status = status
	Result.Message = http.StatusText(status)
	w.WriteHeader(status)
	err = template.Execute(w, Result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
