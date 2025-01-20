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
	if err := json.Unmarshal(data, &cache); err != nil {
		return nil, err
	}

	allWorkouts = append(allWorkouts, cache.Data.Workouts...)
	allMetrics = append(allMetrics, cache.Data.Metrics...)

	return &cache, nil
}

// Load the new directory files into the program data
func loadDirectory(directoryPath string, cacheLastUpdated string) (bool, string, error) {
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return false, cacheLastUpdated, err
	}

	dataWasUpdated := false
	latestFileDate := cacheLastUpdated
	cacheDate, err := time.Parse(dateFormat, cacheLastUpdated)
	if err != nil {
		return false, cacheLastUpdated, err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		// Extract date from filename
		re := regexp.MustCompile(dateRegexPattern)
		matches := re.FindAllString(file.Name(), -1)
		if len(matches) == 0 {
			continue
		}
		fileDate := matches[len(matches)-1]

		// Parse and compare dates
		currentFileDate, err := time.Parse(dateFormat, fileDate)
		if err != nil {
			continue
		}

		// Only process files newer than our cache
		if currentFileDate.After(cacheDate) {
			fmt.Printf("Processing new data from: %s\n", fileDate)

			// Read and parse file
			filePath := directoryPath + "/" + file.Name()
			content, err := os.ReadFile(filePath)
			if err != nil {
				continue
			}

			var fileData HealthData
			if err := json.Unmarshal(content, &fileData); err != nil {
				continue
			}

			// Update our data collections
			allWorkouts = append(allWorkouts, fileData.Data.Workouts...)
			allMetrics = append(allMetrics, fileData.Data.Metrics...)
			dataWasUpdated = true

			// Keep track of the latest file date
			if currentFileDate.After(cacheDate) {
				latestFileDate = fileDate
			}
		}
	}

	return dataWasUpdated, latestFileDate, nil
}

func importData() {
	cache, err := loadCache(cacheFilePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load cache: %v", err))
	}

	// Process directory and get update status
	wasUpdated, latestUpdate, err := loadDirectory(iCloudDirPath, *cache.LastUpdated)
	if err != nil {
		panic(fmt.Sprintf("Failed to load directory: %v", err))
	}

	// Only write to cache if we found new data
	if wasUpdated {
		if err := writeToCache(allWorkouts, allMetrics, &latestUpdate); err != nil {
			panic(fmt.Sprintf("Failed to write cache: %v", err))
		}
		fmt.Printf("Cache updated with data through: %s\n", latestUpdate)
	} else {
		fmt.Println("No new data found, cache remains current")
	}
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

func main() {
	importData()
}
