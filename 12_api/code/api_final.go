package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// // BaseURL is the base endpoint for the star wars API
// const BaseURL = "https://swapi.dev/api/"

// // Planet is a planet type
// type Planet struct {
// 	Name       string `json:"name"`
// 	Population string `json:"population"`
// 	Terrain    string `json:"terrain"`
// }

// // Person is a person type
// type Person struct {
// 	Name         string `json:"name"`
// 	HomeworldURL string `json:"homeworld"`
// 	Homeworld    Planet
// }

// func (p *Person) getHomeworld() {
// 	res, err := http.Get(p.HomeworldURL)
// 	if err != nil {
// 		log.Print("Error fetching homeworld", err)
// 	}

// 	var bytes []byte
// 	if bytes, err = ioutil.ReadAll(res.Body); err != nil {
// 		log.Print("Error reading response body", err)
// 	}

// 	json.Unmarshal(bytes, &p.Homeworld)
// }

// // AllPeople is a collection of Person types
// type AllPeople struct {
// 	People []Person `json:"results"`
// }

// func getPeople(w http.ResponseWriter, r *http.Request) {
// 	res, err := http.Get(BaseURL + "people")

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		log.Print("Failed to request star wars people")
// 	}
// 	fmt.Println(res)

// 	bytes, err := ioutil.ReadAll(res.Body)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		log.Print("Failed to parse request body")
// 	}

// 	var people AllPeople

// 	fmt.Println(string(bytes))

// 	if err := json.Unmarshal(bytes, &people); err != nil {
// 		fmt.Println("Error parsing json", err)
// 	}

// 	fmt.Println(people)

// 	for _, pers := range people.People {
// 		pers.getHomeworld()
// 		fmt.Println(pers)
// 	}

// }
// func main() {
// 	http.HandleFunc("/people", getPeople)
// 	fmt.Println("Serving on :8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
