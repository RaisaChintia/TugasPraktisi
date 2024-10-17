package main

import (
	"fmt"
	"golang/models"
	"golang/router"
	"log"
	"net/http"
)

func main() {
	// Connect to database
	models.ConnectDatabase()

	// Routers
	r := router.InitializeRouter()

	// Start server
	fmt.Println("Starting server on the port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
