package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

type Root struct {
	Data        Data    `json:"data"`
	LastUpdated *string `json:"lastUpdated"`
}

type Data struct {
	Workouts []Workout `json:"workouts"`
	Metrics  []Metric  `json:"metrics"`
}
type Workout struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Start    string    `json:"start"`
	End      string    `json:"end"`
	Duration float64   `json:"duration"`
	Distance *struct { // Optionals use pointers
		Units string  `json:"units"`
		Qty   float64 `json:"qty"`
	} `json:"distance,omitempty"`
	ActiveEnergyBurned *struct {
		Units string  `json:"units"`
		Qty   float64 `json:"qty"`
	} `json:"activeEnergyBurned,omitempty"`
	Intensity *struct {
		Units string  `json:"units"`
		Qty   float64 `json:"qty"`
	} `json:"intensity,omitempty"`
	Location *string `json:"location,omitempty"`
	Humidity *struct {
		Units string `json:"units"`
		Qty   int64  `json:"qty"`
	} `json:"humidity,omitempty"`
	Temperature *struct {
		Units string  `json:"units"`
		Qty   float64 `json:"qty"`
	} `json:"temperature,omitempty"`
	LapLength *struct {
		Units string  `json:"units"`
		Qty   float64 `json:"qty"`
	} `json:"lapLength,omitempty"`
}
type Metric struct {
	Name string `json:"name"`
	Data []struct {
		Date string  `json:"date"`
		Qty  float64 `json:"qty"`
	} `json:"data"`
	Units string `json:"units"`
}

func main() {
	// Open the directory (iCloud Drive)
	dirPath := "/Users/saavedj/Library/Mobile Documents/com~apple~CloudDocs/health-data"

	files, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile("cache.json")
	if err != nil {
		panic(err)
	}

	var root Root
	err = json.Unmarshal(data, &root)
	if err != nil {
		panic(err)
	}

	lastUpdated := *root.LastUpdated

	// Aggregate all the data
	var allWorkouts []Workout
	var allMetrics []Metric

	// Sift through the files for json files
	var thisDate string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			// Extract the date from the filename
			re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
			matches := re.FindAllString(file.Name(), -1)
			if len(matches) > 0 {
				thisDate = matches[len(matches)-1]
				fmt.Println("Extracted Date:", thisDate)
			} else {
				fmt.Println("No date found in filename")
			}

			// Compare file date to last cache update
			cacheUpdateDate, err := time.Parse("2006-01-02", lastUpdated)
			if err != nil {
				panic(err)
			}
			currentFileDate, err := time.Parse("2006-01-02", thisDate)
			if err != nil {
				panic(err)
			}

			// Compare the dates
			if cacheUpdateDate.Before(currentFileDate) {
				fmt.Printf("The file date %s is more recent. Updating the cache.\n", currentFileDate)
				lastUpdated = currentFileDate.Format("2006-01-02")
				fmt.Println("Last Updated:", lastUpdated)
			} else if cacheUpdateDate.After(currentFileDate) {
				fmt.Printf("The file date %s is older. No update needed.\n", currentFileDate)
				continue
			} else {
				fmt.Println("The file date is the same as the local cache date. No update needed.")
			}

			// Read file data
			filePath := dirPath + "/" + file.Name()
			content, err := os.ReadFile(filePath)
			if err != nil {
				panic(err)
			}

			var root Root
			err = json.Unmarshal(content, &root)
			if err != nil {
				panic(err)
			}
			// Aggregate the workouts and metrics
			allWorkouts = append(allWorkouts, root.Data.Workouts...)
			allMetrics = append(allMetrics, root.Data.Metrics...)

		}
	}

	writeToCache(allWorkouts, allMetrics, &lastUpdated)

}

func writeToCache(allWorkouts []Workout, allMetrics []Metric, lastUpdated *string) error {
	// Create the Root structure to match the original format
	root := Root{
		Data: Data{
			Workouts: allWorkouts,
			Metrics:  allMetrics,
		},
		LastUpdated: lastUpdated,
	}

	// Marshal the Root structure into JSON
	data, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling data: %v", err)
	}

	// Write the JSON data to cache.json
	err = os.WriteFile("cache.json", data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	fmt.Println("Data written to cache.json")
	return nil
}
