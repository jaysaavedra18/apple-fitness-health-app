package api

import (
	"fitness/data"
	"fmt"
	"net/http"
)

func StartServer() {
	// Import data from the JSON file
	data.ImportData()

	// Run the RESTful API Server
	RegisterRoutes()
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
