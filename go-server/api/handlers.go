// api/handlers.go
package api

import (
	"encoding/json"
	"fitness/data"
	"fitness/models"
	"fmt"
	"net/http"
	"strconv"
)

// Each handler corresponds to a different endpoint in the API
// Define API endpoints and map them to their respective handlers
func HandleWorkoutData(w http.ResponseWriter, r *http.Request) {
	// Handle workout data
	switch r.Method {
	case http.MethodGet:
		GetWorkoutData(w, r)
	case http.MethodPatch:
		UpdateWorkoutData(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetWorkoutData(w http.ResponseWriter, r *http.Request) {
	// Create a copy of the data from AllWorkouts to avoid modifying the original data
	workoutData := append([]models.Workout(nil), data.AllWorkouts...)
	ok := true

	// Get the workout query parameter from the request
	var workout = r.URL.Query().Get("workout")
	// fmt.Println("workout:", workout)
	if workout != "" {
		workoutData, ok = data.FilterWorkout(workoutData, workout)
		if !ok || workoutData == nil {
			http.Error(w, "Error filtering workout data by workout name", http.StatusInternalServerError)
			return
		}
	}

	// Get the calories threshold query parameter from the request
	var calories = r.URL.Query().Get("calories")
	// fmt.Println("calories:", calories)
	if calories != "" {
		caloriesParsed, err := strconv.ParseFloat(calories, 64)
		if err != nil {
			fmt.Println("Error parsing string to float:", err)
			http.Error(w, "Error parsing calories threshold", http.StatusBadRequest)
			return
		}
		// Filter the workout data based on the parsed calorie threshold
		workoutData, ok = data.FilterCalories(workoutData, caloriesParsed)
		if !ok || workoutData == nil {
			http.Error(w, "Error filtering workout data by calorie threshold", http.StatusInternalServerError)
			return
		}
	}

	// Get the date query parameter from the request
	var start = r.URL.Query().Get("start")
	var end = r.URL.Query().Get("end")
	// fmt.Println("start:", start)
	// fmt.Println("end:", end)
	if start != "" {
		// Filter the workout data based on the start date
		workoutData, ok = data.FilterDate(workoutData, start, true)
		if !ok || workoutData == nil {
			http.Error(w, "Error filtering workout data by start date", http.StatusInternalServerError)
			return
		}
	}
	if end != "" {
		// Filter the workout data based on the end date
		workoutData, ok = data.FilterDate(workoutData, end, false)
		if !ok || workoutData == nil {
			http.Error(w, "Error filtering workout data by end date", http.StatusInternalServerError)
			return
		}
	}

	// Set response header and return the filtered workout data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workoutData)

}

func UpdateWorkoutData(w http.ResponseWriter, r *http.Request) {
	// Update workout data
}
