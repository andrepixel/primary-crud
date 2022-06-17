package entity

type Game struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	DateOfExist string `json:"date_of_exist,omitempty"`
	Company     string `json:"company,omitempty"`
	ReleaseYear string `json:"release_year,omitempty"`
}
