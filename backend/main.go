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

}
