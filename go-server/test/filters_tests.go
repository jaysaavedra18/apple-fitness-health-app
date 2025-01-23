// test/data_test.go

package test

import (
	"fitness/data"
	"fitness/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var workoutData []models.Workout

func init() {
	// Create mock data for testing
	workoutData = []models.Workout{
		{
			ID:                 "1",
			Name:               "Outdoor Run",
			Duration:           1800, // Duration in seconds (30 minutes)
			Distance:           &models.Measurement{Units: "mi", Qty: 5.0},
			Start:              "2021-01-01T07:00:00Z",
			End:                "2021-01-01T07:30:00Z",
			ActiveEnergyBurned: &models.Measurement{Qty: 350.0, Units: "kcal"}, // Added calories
		},
		{
			ID:                 "2",
			Name:               "Indoor Run",
			Duration:           2700, // Duration in seconds (45 minutes)
			Distance:           &models.Measurement{Units: "mi", Qty: 7.5},
			Start:              "2021-01-02T07:00:00Z",
			End:                "2021-01-02T07:45:00Z",
			ActiveEnergyBurned: &models.Measurement{Qty: 250.0, Units: "kcal"}, // Added calories
		},
		{
			ID:                 "3",
			Name:               "Pool Swim",
			Duration:           3600, // Duration in seconds (60 minutes)
			Distance:           &models.Measurement{Units: "mi", Qty: 1.0},
			Start:              "2021-01-03T07:00:00Z",
			End:                "2021-01-03T08:00:00Z",
			ActiveEnergyBurned: &models.Measurement{Qty: 1400.0, Units: "kcal"}, // Added calories
		},
		{
			ID:                 "4",
			Name:               "Outdoor Run",
			Duration:           1500, // Duration in seconds (25 minutes)
			Distance:           &models.Measurement{Units: "mi", Qty: 4.0},
			Start:              "2021-01-04T07:00:00Z",
			End:                "2021-01-04T07:25:00Z",
			ActiveEnergyBurned: &models.Measurement{Qty: 300.0, Units: "kcal"}, // Added calories
		},
		{
			ID:                 "5",
			Name:               "Indoor Run",
			Duration:           2400, // Duration in seconds (40 minutes)
			Distance:           &models.Measurement{Units: "mi", Qty: 6.0},
			Start:              "2021-01-05T07:00:00Z",
			End:                "2021-01-05T07:40:00Z",
			ActiveEnergyBurned: &models.Measurement{Qty: 270.0, Units: "kcal"}, // Added calories
		},
		{
			ID:                 "6",
			Name:               "Pool Swim",
			Duration:           3300, // Duration in seconds (55 minutes)
			Distance:           &models.Measurement{Units: "mi", Qty: 0.5},
			Start:              "2021-01-06T07:00:00Z",
			End:                "2021-01-06T07:55:00Z",
			ActiveEnergyBurned: &models.Measurement{Qty: 350.0, Units: "kcal"}, // Added calories
		},
	}

}

func TestFilterWorkout(t *testing.T) {
	// Test 1: Filter for "Outdoor Run"
	workouts1, ok1 := data.FilterWorkout(workoutData, "Outdoor Run")
	assert.NotEmpty(t, workouts1, "Expected workouts to be returned, but got none.")
	assert.True(t, ok1, "Expected ok to be true when matches are found.")

	// Test 2: Filter for "Outdoor Run, Indoor Run"
	workouts2, ok2 := data.FilterWorkout(workoutData, "Outdoor Run, Indoor Run")
	assert.NotEmpty(t, workouts2, "Expected workouts to be returned, but got none.")
	assert.True(t, ok2, "Expected ok to be true when matches are found.")

	// Test 3: Filter for "Outdoor Run, Indoor Run, Pool Swim"
	workouts3, ok3 := data.FilterWorkout(workoutData, "Outdoor Run, Indoor Run, Pool Swim")
	assert.NotEmpty(t, workouts3, "Expected workouts to be returned, but got none.")
	assert.True(t, ok3, "Expected ok to be true when matches are found.")

	// Test 4: Filter for "Sky Dive" (no match) and "Outdoor Run" (match)
	workouts4, ok4 := data.FilterWorkout(workoutData, "Sky Dive, outdoor Run")
	assert.NotEmpty(t, workouts4, "Expected workouts to be returned for 'Outdoor Run', but got none.")
	assert.True(t, ok4, "Expected ok to be true even when only 'Outdoor Run' matches.")

	// Test 5: Filter for "Sky Dive" (no match)
	workouts5, ok5 := data.FilterWorkout(workoutData, "Sky Dive")
	assert.Empty(t, workouts5, "Expected no workouts to match, but got some.")
	assert.False(t, ok5, "Expected ok to be false when no matches are found.")

}
func TestFilterCalories(t *testing.T) {
	// Test 1: Filter for workouts with calories above a certain threshold (e.g., 300)
	workouts1, ok1 := data.FilterCalories(workoutData, 300)
	assert.NotEmpty(t, workouts1, "Expected workouts to be returned, but got none.")
	assert.True(t, ok1, "Expected ok to be true when matches are found.")
	// Ensure the filtered workouts have calories >= 300
	for _, workout := range workouts1 {
		assert.GreaterOrEqual(t, workout.ActiveEnergyBurned.Qty, 300.0, "Expected calories to be greater than or equal to 300.")
	}

	// Test 2: Filter for workouts with calories above a high threshold (e.g., 1000)
	workouts2, ok2 := data.FilterCalories(workoutData, 1000)
	assert.NotEmpty(t, workouts2, "Expected workouts to be returned, but got none.")
	assert.True(t, ok2, "Expected ok to be true when matches are found.")
	// Ensure the filtered workouts have calories < 1000
	for _, workout := range workouts2 {
		assert.GreaterOrEqual(t, workout.ActiveEnergyBurned.Qty, 1000.0, "Expected calories to be less than 1000.")
	}

	// Test 3: Filter for workouts with no calories above a very high threshold (e.g., 10000)
	workouts3, ok3 := data.FilterCalories(workoutData, 10000)
	assert.Empty(t, workouts3, "Expected no workouts to match, but got some.")
	assert.False(t, ok3, "Expected ok to be false when no matches are found.")

	// Test 4: Filter for workouts with no threshold (e.g., null calories)
	workouts5, ok5 := data.FilterCalories(workoutData, 0)
	assert.NotEmpty(t, workouts5, "Expected workouts to be returned, but got none.")
	assert.True(t, ok5, "Expected ok to be true when workouts with missing calorie data are included.")
	// Ensure the filtered workouts are returned even if they don't have calorie data
	for _, workout := range workouts5 {
		if workout.ActiveEnergyBurned == nil {
			assert.GreaterOrEqual(t, workout.ActiveEnergyBurned, 0, "Expected no ActiveEnergyBurned for this workout.")
		}
	}
}
