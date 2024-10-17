package router

import (
	"golang/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func auntenticmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || user != "raisaadmin" || pass != "raisax123" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)		
			return
		}
		next.ServeHTTP(w, r)
	})	
}
func InitializeRouter() *mux.Router {
	r := mux.NewRouter()

	// Middleware
	r.Use(auntenticmiddleware)

	// Routers
	r.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies", handlers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", handlers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", handlers.DeleteMovie).Methods("DELETE")

	return r
}
