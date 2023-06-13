package handlers

import (
	"countryapi/models"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type CountryJson struct {
	Countries []models.Country
}

var RecieveCountryData CountryJson

// get countries data from a file
func GetCountryData(fileLocation string) (err error) {
	file, err := os.ReadFile(fileLocation)
	if err == nil {
		if err := json.Unmarshal([]byte(string(file)), &RecieveCountryData); err != nil {
			log.Fatal(err.Error())
		}
	}
	return
}

func SetCountryData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.CountriesData)

	log.Printf("Listeneing on: %v", r.Host)
}
