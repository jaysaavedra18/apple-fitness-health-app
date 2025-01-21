// data/filter.go

package data

import (
	"fitness/models"
	"strings"
)

func FilterWorkoutData(workoutType string) ([]models.Workout, bool) {
	// Split the workout type into workout names if multiple present
	targetNames := strings.Split(workoutType, ",")
	for i, name := range targetNames {
		targetNames[i] = strings.TrimSpace(name)
	}

	// Filter the workout data based on the workout type
	var filteredWorkouts []models.Workout
	for _, workout := range AllWorkouts { // Check each workout for match
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
