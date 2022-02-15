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
)

var siteTemplate *template.Template

type Artists struct {
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
	ID             int               `json:"id"`
	DatesLocations map[string]string `json:"dateslocations"`
}

func main() {
	url := [4]string{
		"https://groupietrackers.herokuapp.com/api/artists",
		"https://groupietrackers.herokuapp.com/api/locations",
		"https://groupietrackers.herokuapp.com/api/dates",
		"https://groupietrackers.herokuapp.com/api/relation",
	}
	siteTemplate = template.Must(template.ParseFiles("main/main.html"))

	ArtistStruct := artistUnmarshler(url[0])
	locationStruct := locationUnmarshler(url[1])
	// dataStruct := datesUnmarshler(url[2])
	//  relationStruct := relationUnmarshler(url[3])
	fmt.Println(ArtistStruct[5].Name)
	fmt.Println(locationStruct[5].Locations)
	// fmt.Println(dataStruct[5].Dates)
	//  fmt.Println(relationStruct)

	mux := http.NewServeMux()
	mux.HandleFunc("./main", home)

	mux.HandleFunc("/artist", artist)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("500 Internal server Error\n", err)
	}
	// for k, v := range relationStruct{
	// 	switch c:= v.(type){
	// 	case string:
	// 		fmt.Printf("item in &q is a string containing %q\n", k, c)
	// 	case float64:
	// 		fmt.Printf("Looks like item %q is a number, specifically %f\n",k, c)
	// 	default:
	// 		fmt.Printf("no idea on the type that %q is but it may be %T\n",k,c)
	// 	}
	// }

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
func relationUnmarshler(link string) map[string]interface{} {

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

	var x map[string]interface{} // creates a map with string key and interface value for the key:value pair identification names are uknown in json

	json.Unmarshal(body, &x) // turns the json into a slice of bytes before unmatshling into interface
	return x

}
func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "hello")
	writer.Write([]byte("hello"))
}

func artist(writer http.ResponseWriter, request *http.Request) {
	// writer.Write([]byte("hello world\n"))
	artistOutput := artistUnmarshler("https://groupietrackers.herokuapp.com/api/artists")
	index := 5

	// var titleName = [5]string{
		// "Name", "Image", "Members", "Creation Date", "First Album"}
	// json.NewEncoder(writer).Encode(artistOutput) // is this needed???
	// writer.Write([]byte(artistOutput[index].Name))
	// fmt.Fprintln(writer, artistOutput[index].Image)
	// fmt.Fprintln(writer, artistOutput[index].Members)
	// fmt.Fprintln(writer, artistOutput[index].CreationDate)
	// fmt.Fprintln(writer, artistOutput[index].FirstAlbum, "\n")
	writer.Header().Set("Content-Type", "text/html") // this tells the program to expect html files and to artistOutput files as html

	// for _, title := range titleName{
	// 	siteTemplate.Execute(writer, title )
	// }

	siteTemplate.Execute(writer, artistOutput[index].Name)
	siteTemplate.Execute(writer, artistOutput[index].Image)
	siteTemplate.Execute(writer, artistOutput[index].Members)
	siteTemplate.Execute(writer, artistOutput[index].CreationDate)
	siteTemplate.Execute(writer, artistOutput[index].FirstAlbum)
		// fmt.Fprintf()
}
