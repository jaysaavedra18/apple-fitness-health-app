// main.go
package main

import (
	"fitness/data"
	"fmt"
)

func main() {
	// Import data from cache and cloud drive
	data.ImportData()

	// Run the CLI Program
	// cli.StartCLI()

	// Run the RESTful API Server
	// api.RegisterRoutes()
	// fmt.Println("Starting server on :8080")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	fmt.Println("Error starting server:", err)
	// }

	w1, ok1 := data.FilterDate(data.AllWorkouts, "2024-12-01", true)
	w2, ok2 := data.FilterDate(data.AllWorkouts, "2023-12-25", false)

	fmt.Println("Filtered workouts by start date:", w1[0], ok1)
	fmt.Println("Filtered workouts by end date:", w2[0], ok2)

}
