package handlers

import (
	"countryapi/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type CountryJson struct {
	Countries []models.Country
}

var RecieveCountryData CountryJson

// GetCountryData() read and returns data from a given file.
// @Param {fileLocation} url path of the file to be read
func GetCountryData(fileLocation string) (err error) {
	file, err := os.ReadFile(fileLocation)
	if err == nil {
		if err := json.Unmarshal([]byte(string(file)), &RecieveCountryData); err != nil {
			log.Fatal(err.Error())
		}
	}
	return
}

// greet() returns a welcome message when server is access successfully.
func greet() string {
	message := fmt.Sprintf("Good Day 😇, its now %v from our side. \n You can access the countries data endpoint using /api/countries", time.Now().Format(time.Kitchen))
	return message
}

func Greetings(w http.ResponseWriter, r *http.Request) {
	//Handle unknown url routing paths
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte(greet()))
}

// SetCountryData(), return a json format of the countries data.
func SetCountryData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method Not Allowd", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.CountriesData)
}
