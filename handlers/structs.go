package handlers

type Errors struct {
	Status  int
	Message string
}

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

type Concert struct {
	Id             int
	DatesLocations map[string][]string
}

type AllData struct {
	Main     Artist
	Concerts map[string][]string
}

type Place struct {
	Index []struct {
		Id             int
		DatesLocations map[string][]string
	}
}

type Location struct {
	Id        int
	Locations []string
	Dates     string
}

type SearchInput struct {
	Group               []Artist
	People              []string
	Created             []int
	Places              []string
	SearchInput         string
	FromCareerStartDate int
	ToCareerStartDate   int
	FromFirstAlbumDate  int
	ToFirstAlbumDate    int
	Member1             bool
	Member2             bool
	Member3             bool
	Member4             bool
	Member5             bool
	Member6             bool
	Member7             bool
	Member8             bool
	Location            string
	Suggestions         []string
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ContainsInt(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}

	return false
}
