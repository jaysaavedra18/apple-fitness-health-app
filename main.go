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
	fmt.Println(data.FilterWorkoutData("Outdoor Run"))                        // Test 1 workout filter
	fmt.Println(data.FilterWorkoutData("Outdoor Run, Indoor Run"))            // Test 2 workout filter
	fmt.Println(data.FilterWorkoutData("Outdoor Run, Indoor Run, Pool Swim")) // Test 3 workout filter
	fmt.Println(data.FilterWorkoutData("Sky Dive"))                           // Test no match workout filter
}
