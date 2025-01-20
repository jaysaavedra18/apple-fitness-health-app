// data/storage.go
package data

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"fitness/config"
	"fitness/models"
)

var (
	AllWorkouts []models.Workout
	AllMetrics  []models.Metric
)

func LoadCache(filename string) (*models.HealthData, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cache models.HealthData
	if err := json.Unmarshal(data, &cache); err != nil {
		return nil, err
	}

	AllWorkouts = append(AllWorkouts, cache.Data.Workouts...)
	AllMetrics = append(AllMetrics, cache.Data.Metrics...)

	return &cache, nil
}

// Load the new directory files into the program data
func LoadDirectory(directoryPath string, cacheLastUpdated string) (bool, string, error) {
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return false, cacheLastUpdated, err
	}

	dataWasUpdated := false
	latestFileDate := cacheLastUpdated
	cacheDate, err := time.Parse(config.DateFormat, cacheLastUpdated)
	if err != nil {
		return false, cacheLastUpdated, err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		// Extract date from filename
		re := regexp.MustCompile(config.DateRegexPattern)
		matches := re.FindAllString(file.Name(), -1)
		if len(matches) == 0 {
			continue
		}
		fileDate := matches[len(matches)-1]

		// Parse and compare dates
		currentFileDate, err := time.Parse(config.DateFormat, fileDate)
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

			var fileData models.HealthData
			if err := json.Unmarshal(content, &fileData); err != nil {
				continue
			}

			// Update our data collections
			AllWorkouts = append(AllWorkouts, fileData.Data.Workouts...)
			AllMetrics = append(AllMetrics, fileData.Data.Metrics...)
			dataWasUpdated = true

			// Keep track of the latest file date
			if currentFileDate.After(cacheDate) {
				latestFileDate = fileDate
			}
		}
	}

	return dataWasUpdated, latestFileDate, nil
}

func ImportData() {
	cache, err := LoadCache(config.CacheFilePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load cache: %v", err))
	}

	// Process directory and get update status
	wasUpdated, latestUpdate, err := LoadDirectory(config.ICloudDirPath, *cache.LastUpdated)
	if err != nil {
		panic(fmt.Sprintf("Failed to load directory: %v", err))
	}

	// Only write to cache if we found new data
	if wasUpdated {
		if err := WriteToCache(AllWorkouts, AllMetrics, &latestUpdate); err != nil {
			panic(fmt.Sprintf("Failed to write cache: %v", err))
		}
		fmt.Printf("Cache updated with data through: %s\n", latestUpdate)
	} else {
		// fmt.Println("No new data found, cache remains current")
	}
}

func WriteToCache(AllWorkouts []models.Workout, AllMetrics []models.Metric, lastUpdated *string) error {
	// Create the HealthData structure to match the original format
	healthData := models.HealthData{
		Data: models.DataCollection{
			Workouts: AllWorkouts,
			Metrics:  AllMetrics,
		},
		LastUpdated: lastUpdated,
	}

	// Marshal the HealthData structure into JSON
	data, err := json.MarshalIndent(healthData, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling data: %v", err)
	}

	// Write the JSON data to cache.json
	err = os.WriteFile(config.CacheFilePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	fmt.Printf("Data written to %s\n", config.CacheFilePath)
	return nil
}
