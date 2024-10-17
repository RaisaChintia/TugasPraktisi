package models

type Movie struct {
	ID    string `json:"MovieId"`
	Year  string `json:"ReleaseYear"`
	Title string `json:"MovieTitle"`
}
