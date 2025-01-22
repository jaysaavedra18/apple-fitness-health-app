package cli

import (
	"fitness/api"
	"fitness/data"
	"fitness/printer"
	"fmt"
	"net/http"
	"os"
)

// Start the command line interface
func StartCLI() {
	fmt.Println()

	// Parse command line flags
	flags := ParseFlags()
	opts := CreatePrintOptions(flags)

	// If the server flag is not set, print workout data
	if !opts.Server {
		var err error
		switch flags.DataType {
		case "workouts": // Print workout data
			err = printer.PrintHealthData(data.AllWorkouts, opts)
		case "metrics": // Print metric data
			err = printer.PrintHealthData(data.AllMetrics, opts)
		default: // Invalid data type
			fmt.Fprintf(os.Stderr, "Invalid data type: %s\n", flags.DataType)
			os.Exit(1)
		}
		// Handle any errors that occurred during printing
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println()
	} else {
		// Run the RESTful API Server
		api.RegisterRoutes()
		fmt.Println("Starting server on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("Error starting server:", err)
		}
	}
}
