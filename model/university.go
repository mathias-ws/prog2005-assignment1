package model

// University Struct that defines the fields of a university.
type University struct {
	Name      string   `json:"name"`
	Country   string   `json:"country"`
	Isocode   string   `json:"isocode"`
	Webpages  []string `json:"webpages"`
	Languages []string `json:"languages"`
	Map       string   `json:"map"`
}
