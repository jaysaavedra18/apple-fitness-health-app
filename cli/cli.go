package cli

import (
	"fitness/data"
	"fitness/printer"
	"fmt"
	"os"
)

// Start the command line interface
func StartCLI() {
	fmt.Println()

	// Parse command line flags
	flags := ParseFlags()
	opts := CreatePrintOptions(flags)

	// Determine which data to display
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
}
