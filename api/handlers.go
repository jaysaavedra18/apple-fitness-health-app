// api/handlers.go
package api

import (
	"encoding/json"
	"fitness/data"
	"fmt"
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
	// Handle workout query and errors fetching data
	var workoutType = r.URL.Query().Get("workoutType")
	fmt.Println("workoutType:", workoutType)
	workoutData, ok := data.FilterWorkoutData(workoutType)
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
