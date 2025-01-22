package utils

import (
	"fitness/config"
	"fitness/models"
	"time"
)

func CalculateWorkoutsPerMonth(workouts []models.Workout) map[string]int {
	// Initialize a map to store total workouts per month
	workoutsPerMonth := make(map[string]int)

	// Iterate over the workouts
	for _, workout := range workouts {
		// Parse the start date
		startTime, err := time.Parse(config.TimeFormat, workout.Start)
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

func CalculateDistancePerWorkout(workouts []models.Workout) map[string]float64 {
	// Initialize a map to store total distances per workout type
	distancePerWorkout := make(map[string]float64)

	// Iterate over the workouts
	for _, workout := range workouts {
		// Check if the workout has a distance
		if workout.Distance != nil {
			// Aggregate the distance by workout name
			distancePerWorkout[workout.Name] += workout.Distance.Qty
		}
	}

	return distancePerWorkout
}

func CalculateDistancePerWeek(workouts []models.Workout) map[string]float64 {
	// Initialize a map to store total distance per week
	distancePerWeek := make(map[string]float64)

	// Iterate over the workouts
	for _, workout := range workouts {
		// Check if the workout has distance data
		if workout.Distance != nil {
			// Parse the start date
			startTime, err := time.Parse(config.TimeFormat, workout.Start)
			if err != nil {
				// Skip if the date parsing fails
				continue
			}

			// Get the date for the start of the week (Monday)
			// _, week := startTime.ISOWeek()

			// Find the start of the week (Monday)
			weekStart := startTime.AddDate(0, 0, -int(startTime.Weekday()-time.Monday))

			// Format the date as "MM-DD-YYYY"
			weekOf := weekStart.Format(config.DateFormat)

			// Aggregate the distance for the week
			distancePerWeek[weekOf] += workout.Distance.Qty
		}
	}

	return distancePerWeek
}

func CalculateEnergyPerWeek(workouts []models.Workout) map[string]float64 {
	// Initialize a map to store total energy burned per week
	energyPerWeek := make(map[string]float64)

	// Iterate over the workouts
	for _, workout := range workouts {
		// Check if the workout has energy burned data
		if workout.ActiveEnergyBurned != nil {
			// Parse the start date
			startTime, err := time.Parse(config.TimeFormat, workout.Start)
			if err != nil {
				// Skip if the date parsing fails
				continue
			}

			// Get the date for the start of the week (Monday)
			// _, week := startTime.ISOWeek()

			// Find the start of the week (Monday)
			weekStart := startTime.AddDate(0, 0, -int(startTime.Weekday()-time.Monday))

			// Format the date as "MM-DD-YYYY"
			weekOf := weekStart.Format(config.DateFormat)

			// Aggregate the energy burned for the week
			energyPerWeek[weekOf] += workout.ActiveEnergyBurned.Qty
		}
	}

	return energyPerWeek
}
