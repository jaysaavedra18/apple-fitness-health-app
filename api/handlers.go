// api/handlers.go
package api

import (
	"encoding/json"
	"fitness/data"
	"fitness/models"
	"net/http"
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
	// Initialize the workout data to be filtered through
	var workoutData []models.Workout

	// Get the workout query parameter from the request
	var workout = r.URL.Query().Get("workout")
	workoutData, ok := data.FilterWorkoutData(workout)

	if !ok || workoutData == nil {
		http.Error(w, "Error filtering workout data", http.StatusInternalServerError)
		return
	}

	// Set response header and return the filtered workout data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workoutData)

}

func UpdateWorkoutData(w http.ResponseWriter, r *http.Request) {
	// Update workout data
}
