package model

type NationalityDate struct {
	Count   int       `json:"count"`
	Name    string    `json:"name"`
	Country []Country `json:"country"`
}

type Country struct {
	CountryId   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}
