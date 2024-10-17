package handlers

import (
	"encoding/json"
	"golang/models"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// Update to database
	query := "UPDATE netflix SET id = ?, Year = ?, Title = ? WHERE id = ?"
	_, err := models.DB.Exec(query, movie.ID, movie.Year, movie.Title, id)
	if err != nil {
		http.Error(w, "Error updating movie into database \n"+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(movie)
}
