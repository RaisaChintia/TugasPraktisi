package handlers

import (
	"encoding/json"
	"golang/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// Delete from database
	query := "DELETE FROM netflix WHERE id = ?"
	_, err := models.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Error deleting movie from database", http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("Movie deleted successfully")
}
