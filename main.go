// main.go
package main

import (
	"fitness/data"
)

func main() {
	// Import data from cache and cloud drive
	data.ImportData()

	// Run the CLI Program
	// cli.StartCLI()

	// Run the RESTful API Server

}
