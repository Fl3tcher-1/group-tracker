package main

import (
	// "encoding/json"
	// "errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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

	artistStruct := artistUnmarshler(url[0])
	locationStruct := locationUnmarshler(url[1])
	fmt.Println(artistStruct[5].Name)
	fmt.Println(locationStruct[5].Locations)

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
		log.Fatal(err3, "can not unmarshal JSON")
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
