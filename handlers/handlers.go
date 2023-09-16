package handlers

import (
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strings"
)

var arr_cities [][]string

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
	var members []string
	var created []int
	for _, v := range band {
		for _, vv := range v.Members {
			if !Contains(members, v.Name) {
				members = append(members, vv)
			}
		}
		if !ContainsInt(created, v.CreationDate) {
			created = append(created, v.CreationDate)
		}
	}
	locations, err2 := JsonLocations()
	if err2 != nil {
		Errorshandler(w, http.StatusInternalServerError)
	}
	var new_locations, res_locations []string
	var arr [][]reflect.Value
	for _, v := range locations.Index {
		arr = append(arr, reflect.ValueOf(v.DatesLocations).MapKeys())
	}

	for _, v := range arr {
		var array []string
		for _, vv := range v {
			vv2 := strings.ReplaceAll(vv.String(), "_", " ")
			vv2 = strings.ReplaceAll(vv2, "-", ", ")
			array = append(array, vv2)
			vv2 = strings.Title(vv2)
			replacer := strings.NewReplacer("Uk", "UK", "Usa", "USA")
			res_vv := replacer.Replace(vv2)

			new_locations = append(new_locations, res_vv)

		}
		arr_cities = append(arr_cities, array)

	}

	// adding cities without duplicates to res_locations - array
	for _, v := range new_locations {
		if !Contains(res_locations, v) {
			res_locations = append(res_locations, v)
		}
	}

	res := SearchInput{
		Group:   band,
		People:  members,
		Created: created,
		Places:  res_locations,
	}
	err = template.Execute(w, res)
	if err != nil {
		log.Println(err.Error())
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
