// main.go
package main

import (
	"fitness/data"
	"fitness/printer"
)

func main() {
	// cli.StartCLI()

	data.ImportData()
	printer.PrintCustom(data.AllWorkouts)
}
