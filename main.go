package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

type HealthData struct { // HealthData represents the root JSON structure
	Data        DataCollection `json:"data"`
	LastUpdated *string        `json:"lastUpdated"`
}
type DataCollection struct { // DataCollection holds workouts and metrics
	Workouts []Workout `json:"workouts"`
	Metrics  []Metric  `json:"metrics"`
}
type Measurement struct { // Measurement represents a generic quantity with units
	Units string  `json:"units"`
	Qty   float64 `json:"qty"`
}
type Workout struct { // Workout represents a single workout session
	ID                 string       `json:"id"`
	Name               string       `json:"name"`
	Start              string       `json:"start"`
	End                string       `json:"end"`
	Duration           float64      `json:"duration"`
	Distance           *Measurement `json:"distance,omitempty"`
	ActiveEnergyBurned *Measurement `json:"activeEnergyBurned,omitempty"`
	Intensity          *Measurement `json:"intensity,omitempty"`
	Location           *string      `json:"location,omitempty"`
	Humidity           *struct {
		Units string `json:"units"`
		Qty   int64  `json:"qty"`
	} `json:"humidity,omitempty"`
	Temperature *Measurement `json:"temperature,omitempty"`
	LapLength   *Measurement `json:"lapLength,omitempty"`
}
type MetricData struct { // MetricData represents a single data point in a metric
	Date string  `json:"date"`
	Qty  float64 `json:"qty"`
}
type Metric struct { // Metric represents a collection of measurements over time
	Name  string       `json:"name"`
	Data  []MetricData `json:"data"`
	Units string       `json:"units"`
}

// Global constants and variables used throughout the program
const (
	dateFormat       = "2006-01-02"
	iCloudDirPath    = "/Users/saavedj/Library/Mobile Documents/com~apple~CloudDocs/health-data"
	cacheFilePath    = "cache.json"
	dateRegexPattern = `\d{4}-\d{2}-\d{2}`
)

var (
	allWorkouts []Workout
	allMetrics  []Metric
)

// Load the cache into the program data
func loadCache(filename string) (*HealthData, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cache HealthData
	err = json.Unmarshal(data, &cache)
	if err != nil {
		return nil, err
	}
	return &cache, nil
}

func main() {
	cache, err := loadCache(cacheFilePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load cache: %v", err))
	}
	lastUpdated := *cache.LastUpdated

	// Sift through the files for json files
	files, err := os.ReadDir(iCloudDirPath)
	if err != nil {
		panic(err)
	}
	var thisDate string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			// Extract the date from the filename
			re := regexp.MustCompile(dateRegexPattern)
			matches := re.FindAllString(file.Name(), -1)
			if len(matches) > 0 {
				thisDate = matches[len(matches)-1]
				fmt.Println("Extracted Date:", thisDate)
			} else {
				fmt.Println("No date found in filename")
			}

			// Compare file date to last cache update
			cacheUpdateDate, err := time.Parse(dateFormat, lastUpdated)
			if err != nil {
				panic(err)
			}
			currentFileDate, err := time.Parse(dateFormat, thisDate)
			if err != nil {
				panic(err)
			}

			// Compare the dates
			if cacheUpdateDate.Before(currentFileDate) {
				fmt.Printf("The file date %s is more recent. Updating the cache.\n", currentFileDate)
				lastUpdated = currentFileDate.Format(dateFormat)
			} else if cacheUpdateDate.After(currentFileDate) {
				fmt.Printf("The file date %s is older. No update needed.\n", currentFileDate)
				continue
			} else {
				fmt.Println("The file date is the same as the local cache date. No update needed.")
			}

			// Read file data
			filePath := iCloudDirPath + "/" + file.Name()
			content, err := os.ReadFile(filePath)
			if err != nil {
				panic(err)
			}

			var healthData HealthData
			err = json.Unmarshal(content, &healthData)
			if err != nil {
				panic(err)
			}
			// Aggregate the workouts and metrics
			allWorkouts = append(allWorkouts, healthData.Data.Workouts...)
			allMetrics = append(allMetrics, healthData.Data.Metrics...)

		}
	}

	writeToCache(allWorkouts, allMetrics, &lastUpdated)
}

func writeToCache(allWorkouts []Workout, allMetrics []Metric, lastUpdated *string) error {
	// Create the HealthData structure to match the original format
	healthData := HealthData{
		Data: DataCollection{
			Workouts: allWorkouts,
			Metrics:  allMetrics,
		},
		LastUpdated: lastUpdated,
	}

	// Marshal the HealthData structure into JSON
	data, err := json.MarshalIndent(healthData, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling data: %v", err)
	}

	// Write the JSON data to cache.json
	err = os.WriteFile(cacheFilePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	fmt.Printf("Data written to %s\n", cacheFilePath)
	return nil
}
