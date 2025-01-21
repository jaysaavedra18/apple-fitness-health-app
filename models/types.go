// models/types.go
package models

// HealthData is the top-level struct that contains all health data
type HealthData struct {
	Data        DataCollection `json:"data"`        // Collection of workout and metric data
	LastUpdated *string        `json:"lastUpdated"` // Timestamp of the last update
}

// DataCollection contains all workout and metric data
type DataCollection struct {
	Workouts []Workout `json:"workouts"` // Collection of workout data
	Metrics  []Metric  `json:"metrics"`  // Collection of metric data
}

// Measurement is a generic struct that represents a quantity with units
type Measurement struct {
	Units string  `json:"units"` // Units of the measurement
	Qty   float64 `json:"qty"`   // Quantity of the measurement
}

// Workout represents a single workout entry
type Workout struct {
	ID                 string       `json:"id"`                           // Unique identifier for the workout
	Name               string       `json:"name"`                         // Name of the workout
	Start              string       `json:"start"`                        // Start time of the workout
	End                string       `json:"end"`                          // End time of the workout
	Duration           float64      `json:"duration"`                     // Duration of the workout in seconds
	Distance           *Measurement `json:"distance,omitempty"`           // Distance covered during the workout
	ActiveEnergyBurned *Measurement `json:"activeEnergyBurned,omitempty"` // Energy burned during the workout
	Intensity          *Measurement `json:"intensity,omitempty"`          // Intensity level of the workout
	Location           *string      `json:"location,omitempty"`           // Location of the workout
	Humidity           *struct {    // Humidity data for the workout
		Units string  `json:"units"`
		Qty   float64 `json:"qty"`
	} `json:"humidity,omitempty"`
	Temperature *Measurement `json:"temperature,omitempty"` // Temperature during the workout
	LapLength   *Measurement `json:"lapLength,omitempty"`   // Length of each lap during the workout
}

// MetricData represents a single data point for a metric
type MetricData struct {
	Date string  `json:"date"` // Date of the data point
	Qty  float64 `json:"qty"`  // Quantity of the data point
}

// Metric represents a single metric entry
type Metric struct {
	Name  string       `json:"name"`  // Name of the metric
	Data  []MetricData `json:"data"`  // Collection of data points for the metric
	Units string       `json:"units"` // Units of the metric
}
