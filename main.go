package main

import (
	// "encoding/json"
	// "errors"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var homeTemplate *template.Template
var siteTemplate *template.Template

type Full struct {
	Artists   Artists
	Locations Locations
	Dates     Dates
	Relation  Relation
}
type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationdate"`
	FirstAlbum   string   `json:"firstalbum"`
}

type NewArtists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationdate"`
	FirstAlbum   string   `json:"firstalbum"`
}
type Loc struct {
	Index []Locations
}
type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}
type Dat struct {
	Index []Dates
}
type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
type Rel struct {
	Index []Relation
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"dateslocations"`
}

func main() {
	// url := [4]string{
	// 	"https://groupietrackers.herokuapp.com/api/artists",
	// 	"https://groupietrackers.herokuapp.com/api/locations",
	// 	"https://groupietrackers.herokuapp.com/api/dates",
	// 	"https://groupietrackers.herokuapp.com/api/relation",
	// }
	homeTemplate = template.Must(template.ParseFiles("main/home.html"))
	siteTemplate = template.Must(template.ParseFiles("main/artists.html"))

	// ArtistStruct := artistUnmarshler(url[0])
	// locationStruct := locationUnmarshler(url[1])
	// dataStruct := datesUnmarshler(url[2])
	// relationStruct := relationUnmarshler(url[3])
	// fmt.Println(ArtistStruct)
	// fmt.Println(locationStruct[5].Locations)
	// fmt.Println(dataStruct[5].Dates)
	// fmt.Println(relationStruct[1])

	mux := http.NewServeMux()
	mux.HandleFunc("/main", home)
	mux.HandleFunc("/style", css) // this handles css extension so html template can use it
	mux.HandleFunc("/artist/style", css)
	mux.HandleFunc("/artist/", artist)
	fmt.Printf("Starting server at port 8080\n\t -----------\nhttp://localhost:8080/main\n")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("500 Internal server Error\n", err)
	}

}
func artistUnmarshler(link string) []Artists {
	response, err := http.Get(link)
	//var artist Artists

	if err != nil {
		log.Fatal(err, "no response from request")
	}
	defer response.Body.Close()

	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
	var result []Artists
	if err3 := json.Unmarshal(body, &result); err3 != nil {
		log.Fatal(err3, "can not unmarshal JSON\n")
	}

	return result
}
func css(writer http.ResponseWriter, r *http.Request) {
	http.ServeFile(writer, r, "./main/style.css") // tells html to look for css file in current directory/main/style.css
}

func locationUnmarshler(link string) []Locations {
	locationResponse, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer locationResponse.Body.Close()
	body, err := ioutil.ReadAll(locationResponse.Body)

	var index Loc                                        // hold a struct referecing a struct
	var location Locations                               // holds Location struct fields
	if err := json.Unmarshal(body, &index); err != nil { //unmarshals index element from locations
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &location); err != nil { // unmarshals other elements from locations
		log.Fatal(err)
	}
	return index.Index
}

func datesUnmarshler(link string) []Dates {
	locationResponse, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer locationResponse.Body.Close()
	body, err := ioutil.ReadAll(locationResponse.Body)

	var index Dat                                        // hold a struct referecing a struct
	var date Dates                                       // holds Location struct fields
	if err := json.Unmarshal(body, &index); err != nil { //unmarshals index element from locations
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &date); err != nil { // unmarshals other elements from locations
		log.Fatal(err)
	}
	return index.Index
}

//relationUnmarshler unmarshals json map file into [string]interface{}
func relationUnmarshler(link string) []Relation {

	relationResponse, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer relationResponse.Body.Close()
	body, err := ioutil.ReadAll(relationResponse.Body)
	// var rel Rel                                        // hold a struct referecing a struct
	// var relation Relation                                       // holds Location struct fields
	// if err := json.Unmarshal(body, &rel); err != nil { //unmarshals index element from locations
	// 	log.Fatal(err)
	// }
	// if err := json.Unmarshal(body, &relation); err != nil { // unmarshals other elements from locations
	// 	log.Fatal(err)
	// }
	// // mymap := make(map[string]string)
	var rel Rel
	var x []Relation // creates a map with string key and interface value for the key:value pair identification names are uknown in json
	if err = json.Unmarshal(body, &rel); err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &x) // turns the json into a slice of bytes before unmatshling into interface
	return rel.Index

}
func home(writer http.ResponseWriter, request *http.Request) {
	artistOutput := artistUnmarshler("https://groupietrackers.herokuapp.com/api/artists")
	writer.Header().Set("Content-Type", "text/html") // this tells the program to expect html files and to artistOutput files as html

	homeTemplate.Execute(writer, artistOutput)

	// keys := request.URL.Query()["image"]
	// fmt.Println(keys)
	// fmt.Println(request.Method)
	// s := strings.Trim(request.RequestURI, "/artist")

}

func artist(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/artist/style" {
		return
	}
	url := [4]string{
		"https://groupietrackers.herokuapp.com/api/artists",
		"https://groupietrackers.herokuapp.com/api/locations",
		"https://groupietrackers.herokuapp.com/api/dates",
		"https://groupietrackers.herokuapp.com/api/relation",
	}
	// writer.Write([]byte("hello world\n"))
	artistOutput := artistUnmarshler(url[0])




	locationStruct := locationUnmarshler(url[1])
	dataStruct := datesUnmarshler(url[2])
	relationStruct := relationUnmarshler(url[3])
	// index := 5
	writer.Header().Set("Content-Type", "text/html") // this tells the program to expect html files and to artistOutput files as html

	if err := request.ParseForm(); err != nil {
		return
	}
	request.ParseForm()
	r := request.URL.Path

	idstr := strings.Trim(r, "/artist/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
	}
	id -= 1

	// s:= request.URL.Path

	// fmt.Println(len(s), "s", s)
	// fmt.Println(s)
	var a Full
	a.Artists = artistOutput[id] 
	a.Locations = locationStruct[id]
	a.Dates = dataStruct[id]
	a.Relation = relationStruct[id]

	fmt.Println(a)
	siteTemplate.Execute(writer, a)
	fmt.Println(artistOutput[id].Name)

}
