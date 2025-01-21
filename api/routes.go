// api/routes.go
package api

import (
	"fitness/data"
	"net/http"
)

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
	var workoutType = r.URL.Query().Get("workoutType")
	workoutData, ok := data.FilterWorkoutData(workoutType)
	if !ok || workoutData == nil {
		http.Error(w, "Error filtering workout data", http.StatusInternalServerError)
		return
	}
}

func UpdateWorkoutData(w http.ResponseWriter, r *http.Request) {
	// Update workout data
}
