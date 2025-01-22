// main.go
package main

import (
	"fitness/cli"
	"fitness/data"
)

func main() {
	// Import data from cache and cloud drive
	data.ImportData()

	// Run the CLI Program
	cli.StartCLI()

	// Run the RESTful API Server
	// api.RegisterRoutes()
	// fmt.Println("Starting server on :8080")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	fmt.Println("Error starting server:", err)
	// }

}
