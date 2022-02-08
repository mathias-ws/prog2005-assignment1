package model

//UniversityInfo defines the fields taken from the university info api.
type UniversityInfo struct {
	Name     string   `json:"name"`
	Webpages []string `json:"web_pages"`
	Country  string   `json:"country"`
	Isocode  string   `json:"alpha_two_code"`
}

// CountryApi defines the fields taken from the countries api.
type CountryApi struct {
	Maps      map[string]string `json:"maps"`
	Languages map[string]string `json:"languages"`
	BordersTo []string          `json:"borders"`
}

// University Struct that defines the fields of the response given by the api.
type University struct {
	Name      string            `json:"name"`
	Country   string            `json:"country"`
	Isocode   string            `json:"isocode"`
	Webpages  []string          `json:"webpages"`
	Languages map[string]string `json:"languages"`
	Map       string            `json:"map"`
}
