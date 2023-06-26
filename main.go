package main

import (
	"countryapi/handlers"
	"countryapi/models"
	"github.com/rs/cors"
	"log"
	"net/http"
)

// append data from country.json file to a struct type to be serve on the http server
func init() {
	if err := handlers.GetCountryData("country_data.json"); err != nil {
		log.Fatal(err.Error())
	} else {
		models.CountriesData = append(models.CountriesData, handlers.RecieveCountryData.Countries...)
	}
}

func main() {
	mux := http.NewServeMux()

	cors := cors.New(cors.Options{
		// Add all allow origins to the AllowedOrigins slice
		AllowedOrigins: []string{"http://127.0.0.1:5501", "http://127.0.0.1:5500"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET"},
	})

	mux.HandleFunc("/", handlers.Greetings)
	mux.HandleFunc("/api/countries", handlers.SetCountryData)

	log.Print("Starting server on port :4000")
	err := http.ListenAndServe(":4000", cors.Handler(mux))
	log.Fatal(err)
}
