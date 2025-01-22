// data/filter.go
// Filters for workout data

package data

import (
	"fitness/config"
	"fitness/models"
	"fmt"
	"strings"
	"time"
)

func FilterWorkout(workouts []models.Workout, workoutType string) ([]models.Workout, bool) {
	// If workout type is empty, return all workouts
	if workoutType == "" {
		return workouts, true
	}

	// Split the workout type into workout names if multiple present
	targetNames := strings.Split(workoutType, ",")
	for i, name := range targetNames {
		targetNames[i] = strings.TrimSpace(name)
	}

	// Filter the workout data based on the workout type
	var filteredWorkouts []models.Workout
	for _, workout := range workouts { // Check each workout for match
		for _, name := range targetNames { // Check each target name for match
			if strings.EqualFold(workout.Name, name) {
				filteredWorkouts = append(filteredWorkouts, workout)
				break
			}
		}
	}

	// Return the filtered workouts and a boolean indicating if any were found
	return filteredWorkouts, len(filteredWorkouts) > 0
}

func FilterCalories(workouts []models.Workout, calorieThreshold float64) ([]models.Workout, bool) {
	// If calorie threshold is zero, return all workouts
	if calorieThreshold == 0 {
		return workouts, true
	}

	// Filter the workout data based on the calorie threshold
	var filteredWorkouts []models.Workout
	for _, workout := range workouts {
		if workout.ActiveEnergyBurned != nil && workout.ActiveEnergyBurned.Qty >= calorieThreshold {
			filteredWorkouts = append(filteredWorkouts, workout)
		}
	}

	// Return the filtered workouts and a boolean indicating if any were found
	return filteredWorkouts, len(filteredWorkouts) > 0
}

func FilterDate(workouts []models.Workout, queryDate string, isStartDate bool) ([]models.Workout, bool) {
	var filteredWorkouts []models.Workout
	// If queryDate is empty, return all workouts
	if queryDate == "" {
		return workouts, true
	}

	// Parse the queryDate string into a time.Time object
	providedDate, err := time.Parse(config.DateFormat, queryDate)
	if err != nil {
		fmt.Println("Error parsing query queryDate:", err)
		return nil, false
	}

	for _, workout := range workouts {
		//  Parse the start queryDate out of the workout
		workoutDate, err := time.Parse(config.TimeFormat, workout.Start)
		if err != nil {
			continue
		}

		// Filter the workout data based on the queryDate
		if isStartDate && !workoutDate.Before(providedDate) {
			// Filter out everything before the queryDate
			filteredWorkouts = append(filteredWorkouts, workout)
		} else if !isStartDate && !workoutDate.After(providedDate) {
			// Filter out everything after the queryDate
			filteredWorkouts = append(filteredWorkouts, workout)
		}
	}

	// Return the filtered workouts and a boolean indicating if any were found
	return filteredWorkouts, len(filteredWorkouts) > 0
}
