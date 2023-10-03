package handlers

import (
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

var arr_cities [][]string

func Index(w http.ResponseWriter, r *http.Request) {
	searchInput := strings.TrimSpace(r.FormValue("searchInput"))
	searchInputLower := strings.ToLower(searchInput)
	careerStartDate, _ := strconv.Atoi(r.FormValue("careerStartDate"))
	firstAlbumDate, _ := strconv.Atoi(r.FormValue("firstAlbumDate"))
	member1 := strings.TrimSpace(r.FormValue("1member")) == "on"
	member2 := strings.TrimSpace(r.FormValue("2member")) == "on"
	member3 := strings.TrimSpace(r.FormValue("3member")) == "on"
	member4 := strings.TrimSpace(r.FormValue("4member")) == "on"
	member5 := strings.TrimSpace(r.FormValue("5member")) == "on"
	member6 := strings.TrimSpace(r.FormValue("6member")) == "on"
	member7 := strings.TrimSpace(r.FormValue("7member")) == "on"
	member8 := strings.TrimSpace(r.FormValue("8member")) == "on"
	location := r.FormValue("location")

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

	for _, v := range new_locations {
		if !Contains(res_locations, v) {
			res_locations = append(res_locations, v)
		}
	}

	var filteredBand []Artist
	for _, v := range band {
		a, _ := JsonLocation(v.Locations)
		searchLocation := false
		filterLocation := false
		for _, vv := range a.Locations {
			vv2 := strings.ReplaceAll(vv, "_", " ")
			vv2 = strings.ReplaceAll(vv2, "-", ", ")
			if strings.Contains(strings.ToLower(vv2), strings.ToLower(searchInput)) {
				searchLocation = true
			}
			if strings.ToLower(location) == vv2 || location == "" {
				filterLocation = true
			}
		}
		firstAlbum, _ := strconv.Atoi(strings.Split(v.FirstAlbum, "-")[2])
		if filterLocation && careerStartDate <= v.CreationDate && firstAlbumDate <= firstAlbum && ((!member1 && !member2 && !member3 && !member4 && !member5 && !member6 && !member7 && !member8) || (member1 && len(v.Members) == 1) || (member2 && len(v.Members) == 2) || (member3 && len(v.Members) == 3) || (member4 && len(v.Members) == 4) || (member5 && len(v.Members) == 5) || (member6 && len(v.Members) == 6) || (member7 && len(v.Members) == 7) || (member8 && len(v.Members) == 8)) {
			if searchLocation || strings.Contains(strings.ToLower(v.Name), searchInputLower) || strings.Contains(strings.ToLower(v.FirstAlbum), searchInputLower) || strings.Contains(strconv.FormatInt(int64(v.CreationDate), 10), searchInputLower) {
				filteredBand = append(filteredBand, v)
			} else {
				for _, member := range v.Members {
					if strings.Contains(strings.ToLower(member), searchInputLower) {
						filteredBand = append(filteredBand, v)
					}
				}
			}
		}
	}
	res := SearchInput{
		Group:           filteredBand,
		People:          members,
		Created:         created,
		Places:          res_locations,
		SearchInput:     searchInput,
		CareerStartDate: careerStartDate,
		FirstAlbumDate:  firstAlbumDate,
		Member1:         member1,
		Member2:         member2,
		Member3:         member3,
		Member4:         member4,
		Member5:         member5,
		Member6:         member6,
		Member7:         member7,
		Member8:         member8,
		Location:        location,
	}
	err = template.Execute(w, res)
	if err != nil {
		log.Println(err.Error())
		Errorshandler(w, http.StatusInternalServerError)
		return
	}
}

func Artists(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Path
	linkList := strings.Split(link, "/")
	id, err := strconv.Atoi(linkList[len(linkList)-1])
	if len(linkList) > 3 || linkList[1] != "artist" || (id <= 0 || id > 52) {
		Errorshandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		Errorshandler(w, http.StatusMethodNotAllowed)
		return
	}
	template, err := template.ParseFiles("./template/artist.html")
	if err != nil {
		log.Println(err.Error())
		Errorshandler(w, http.StatusInternalServerError)
		return
	}
	mainPage, err3 := JsonArtists()
	if err3 != nil {
		Errorshandler(w, http.StatusInternalServerError)
		return
	}
	concerts, err4 := JsonConcerts(strconv.Itoa(id))
	if err4 != nil {
		Errorshandler(w, http.StatusInternalServerError)
		return
	}
	MapData := map[string][]string{}

	for key, value := range concerts.DatesLocations {
		key = strings.ReplaceAll(key, "_", " ")
		key = strings.ReplaceAll(key, "-", ", ")
		MapData[key] = value
	}
	res := AllData{
		Main:     mainPage[id-1],
		Concerts: MapData,
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
