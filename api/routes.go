// api/routes.go
package api

import "net/http"

// Register API endpoints and their respective handlers
func RegisterRoutes() {
	// Register the workout data handler
	http.HandleFunc("/workouts", HandleWorkoutData)
}
