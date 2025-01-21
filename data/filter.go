// data/filter.go
// Filters for workout data

package data

import (
	"fitness/models"
	"strings"
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
	return nil, false
}
