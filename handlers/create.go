package handlers

import (
	"encoding/json"
	"golang/models"
	"net/http"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// Insert to database
	query := "INSERT INTO netflix (id, Year, Title) VALUES (?, ?, ?)"
	_, err := models.DB.Exec(query, movie.ID ,movie.Year, movie.Title)

	if err != nil {
		http.Error(w, "Error inserting movie into database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}
