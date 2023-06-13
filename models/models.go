package models

type Country struct {
	Name       string   `json:"name"`
	Capital    string   `json:"capital"`
	Languages  []string `json:"languages"`
	Population int      `json:"population"`
	Flag       string   `json:"flag"`
	Region     string   `json:"region"`
	Area       float64  `json:"area"`
}

// Variable model to store the country data
var CountriesData = []Country{}
