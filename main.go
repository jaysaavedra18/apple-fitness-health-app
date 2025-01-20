package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Root struct {
	Data Data `json:"data"`
}

type Data struct {
	Workouts []Workout `json:"workouts"`
	Metrics  []Metric  `json:"metrics"`
}
type Workout struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Start    string  `json:"start"`
	End      string  `json:"end"`
	Duration float64 `json:"duration"`
	Distance struct {
		Units string  `json:"units"`
		Qty   float64 `json:"qty"`
	} `json:"distance"`
	ActiveEnergyBurned struct {
		Units string  `json:"units"`
		Qty   float64 `json:"qty"`
	} `json:"activeEnergyBurned"`
	Intensity struct {
		Units string  `json:"units"`
		Qty   float64 `json:"qty"`
	} `json:"intensity"`
	Location string `json:"location"`
	Humidity struct {
		Units string `json:"units"`
		Qty   int64  `json:"qty"`
	} `json:"humidity"`
	Temperature struct {
		Units string  `json:"units"`
		Qty   float64 `json:"qty"`
	} `json:"temperature"`
	LapLength struct {
		Units string  `json:"units"`
		Qty   float64 `json:"qty"`
	} `json:"lapLength"`
}
type Metric struct {
	Data []struct {
		Date string  `json:"date"`
		Qty  float64 `json:"qty"`
	} `json:"data"`
}

func main() {
	// Open the directory (iCloud Drive)
	dirPath := "/Users/saavedj/Library/Mobile Documents/com~apple~CloudDocs/health-data"

	files, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}
	// Sift through the files for json files
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			filePath := dirPath + "/" + file.Name()
			fmt.Printf("File: %s\n", filePath)

			content, err := os.ReadFile(filePath)
			if err != nil {
				panic(err)
			}

			var root Root
			err = json.Unmarshal(content, &root)
			if err != nil {
				panic(err)
			}

			for _, workout := range root.Data.Workouts {
				// Do something with all the workout data
				fmt.Printf("Workout: %+v\n", workout)
			}
			for _, metric := range root.Data.Metrics {
				// Do something with all the metrics data
				fmt.Printf("Metric: %+v\n", metric)
			}

		}
	}

}
