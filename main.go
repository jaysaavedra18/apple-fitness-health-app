// main.go
package main

import (
	"fitness/cli"
	"fitness/data"
	"fitness/printer"
	"fmt"
	"os"
)

func main() {
	flags := cli.ParseFlags()
	opts := cli.CreatePrintOptions(flags)

	// Import data
	data.ImportData()
	fmt.Println()

	// Determine which data to display
	var err error
	switch flags.DataType {
	case "workouts":
		err = printer.PrintHealthData(data.AllWorkouts, opts)
	case "metrics":
		err = printer.PrintHealthData(data.AllMetrics, opts)
	default:
		fmt.Fprintf(os.Stderr, "Invalid data type: %s\n", flags.DataType)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
