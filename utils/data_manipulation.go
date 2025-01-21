package utils

import (
	"fitness/models"
	"time"
)

func CalculateWorkoutsPerMonth(workouts []models.Workout) map[string]int {
	// Initialize a map to store total workouts per month
	workoutsPerMonth := make(map[string]int)

	// Iterate over the workouts
	for _, workout := range workouts {
		// Parse the start date
		startTime, err := time.Parse("2006-01-02 15:04:05 -0700", workout.Start)
		if err != nil {
			// Handle error, e.g., skip the workout if parsing fails
			continue
		}

		// Extract year and month in "YYYY-MM" format
		yearMonth := startTime.Format("2006-01")

		// Increment the count for the corresponding month
		workoutsPerMonth[yearMonth]++
	}

	return workoutsPerMonth
}
