package handlers

import (
	"encoding/json"
	"golang/models"
	"net/http"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := models.DB.Query("SELECT id, Year, Title FROM netflix")
	if err != nil {
		http.Error(w, "Error Querying database", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.ID, &movie.Year, &movie.Title); err != nil {
			http.Error(w, "Error scanning rows", http.StatusInternalServerError)
			return
		}
		movies = append(movies, movie)
	}

	json.NewEncoder(w).Encode(movies)
}
