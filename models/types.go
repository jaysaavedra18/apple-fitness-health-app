// models/types.go
package models

type HealthData struct {
	Data        DataCollection `json:"data"`
	LastUpdated *string        `json:"lastUpdated"`
}

type DataCollection struct {
	Workouts []Workout `json:"workouts"`
	Metrics  []Metric  `json:"metrics"`
}

type Measurement struct {
	Units string  `json:"units"`
	Qty   float64 `json:"qty"`
}

type Workout struct {
	ID                 string       `json:"id"`
	Name               string       `json:"name"`
	Start              string       `json:"start"`
	End                string       `json:"end"`
	Duration           float64      `json:"duration"`
	Distance           *Measurement `json:"distance,omitempty"`
	ActiveEnergyBurned *Measurement `json:"activeEnergyBurned,omitempty"`
	Intensity          *Measurement `json:"intensity,omitempty"`
	Location           *string      `json:"location,omitempty"`
	Humidity           *struct {
		Units string `json:"units"`
		Qty   int64  `json:"qty"`
	} `json:"humidity,omitempty"`
	Temperature *Measurement `json:"temperature,omitempty"`
	LapLength   *Measurement `json:"lapLength,omitempty"`
}

type MetricData struct {
	Date string  `json:"date"`
	Qty  float64 `json:"qty"`
}

type Metric struct {
	Name  string       `json:"name"`
	Data  []MetricData `json:"data"`
	Units string       `json:"units"`
}
