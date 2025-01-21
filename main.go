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
	workouts1, ok1 := data.FilterWorkoutData("Outdoor Run")
	fmt.Printf("Test 1 - Data:\n %v, Success: %t\n", workouts1, ok1)

	// Test 2: Filter for "Outdoor Run, Indoor Run"
	workouts2, ok2 := data.FilterWorkoutData("Outdoor Run, Indoor Run")
	fmt.Printf("Test 2 - Data:\n %v, Success: %t\n", workouts2, ok2)

	// Test 3: Filter for "Outdoor Run, Indoor Run, Pool Swim"
	workouts3, ok3 := data.FilterWorkoutData("Outdoor Run, Indoor Run, Pool Swim")
	fmt.Printf("Test 3 - Data:\n %v, Success: %t\n", workouts3, ok3)

	// Test 4: Filter for "Sky Dive" (no match)
	workouts4, ok4 := data.FilterWorkoutData("Sky Dive, outdoor Run")
	fmt.Printf("Test 4 - Data:\n %v, Success: %t\n", workouts4, ok4)
}
