package main

import (
	"countryapi/handlers"
	"countryapi/models"
	"log"
	"net/http"
)

func init() {
	if err := handlers.GetCountryData("./country_data.json"); err != nil {
		log.Fatal(err.Error())
	} else {
		models.CountriesData = append(models.CountriesData, handlers.RecieveCountryData.Countries...)
	}
}

func main() {
	http.HandleFunc("/api/countries", handlers.SetCountryData)
	http.ListenAndServe(":2300", nil)
}
