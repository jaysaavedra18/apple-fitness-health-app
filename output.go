type RootDataMetricsData struct {
	Date string `json:"date"`
	Qty int64 `json:"qty"`
	Source string `json:"source"`
}

type RootDataMetrics struct {
	Units string `json:"units"`
	Data []RootDataMetricsData `json:"data"`
	Name string `json:"name"`
}

type RootDataWorkoutsRoute struct {
	Course float64 `json:"course"`
	Altitude float64 `json:"altitude"`
	Timestamp string `json:"timestamp"`
	Speedaccuracy float64 `json:"speedAccuracy"`
	Speed float64 `json:"speed"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Courseaccuracy float64 `json:"courseAccuracy"`
	Verticalaccuracy float64 `json:"verticalAccuracy"`
	Horizontalaccuracy float64 `json:"horizontalAccuracy"`
}

type RootDataWorkoutsMetadata struct {
}

type RootDataWorkoutsDistance struct {
	Units string `json:"units"`
	Qty float64 `json:"qty"`
}

type RootDataWorkoutsActiveenergyburned struct {
	Units string `json:"units"`
	Qty float64 `json:"qty"`
}

type RootDataWorkoutsIntensity struct {
	Qty float64 `json:"qty"`
	Units string `json:"units"`
}

type RootDataWorkoutsTemperature struct {
	Units string `json:"units"`
	Qty float64 `json:"qty"`
}

type RootDataWorkoutsHumidity struct {
	Qty int64 `json:"qty"`
	Units string `json:"units"`
}

type RootDataWorkoutsActiveenergy struct {
	Date string `json:"date"`
	Units string `json:"units"`
	Source string `json:"source"`
	Qty float64 `json:"qty"`
}

type RootDataWorkoutsHeartraterecovery struct {
	Date string `json:"date"`
	Units string `json:"units"`
	Max int64 `json:"Max"`
	Min int64 `json:"Min"`
	Source string `json:"source"`
	Avg int64 `json:"Avg"`
}

type RootDataWorkoutsHeartratedata struct {
	Avg float64 `json:"Avg"`
	Date string `json:"date"`
	Units string `json:"units"`
	Source string `json:"source"`
	Max int64 `json:"Max"`
	Min int64 `json:"Min"`
}

type RootDataWorkoutsStepcount struct {
	Date string `json:"date"`
	Qty float64 `json:"qty"`
	Source string `json:"source"`
	Units string `json:"units"`
}

type RootDataWorkoutsWalkingandrunningdistance struct {
	Source string `json:"source"`
	Units string `json:"units"`
	Qty float64 `json:"qty"`
	Date string `json:"date"`
}

type RootDataWorkouts struct {
	Name string `json:"name"`
	Walkingandrunningdistance []RootDataWorkoutsWalkingandrunningdistance `json:"walkingAndRunningDistance"`
	Stepcount []RootDataWorkoutsStepcount `json:"stepCount"`
	Heartratedata []RootDataWorkoutsHeartratedata `json:"heartRateData"`
	Heartraterecovery []RootDataWorkoutsHeartraterecovery `json:"heartRateRecovery"`
	Activeenergy []RootDataWorkoutsActiveenergy `json:"activeEnergy"`
	Location string `json:"location"`
	Id string `json:"id"`
	Humidity RootDataWorkoutsHumidity `json:"humidity"`
	Temperature RootDataWorkoutsTemperature `json:"temperature"`
	Intensity RootDataWorkoutsIntensity `json:"intensity"`
	Start string `json:"start"`
	Activeenergyburned RootDataWorkoutsActiveenergyburned `json:"activeEnergyBurned"`
	Distance RootDataWorkoutsDistance `json:"distance"`
	Metadata RootDataWorkoutsMetadata `json:"metadata"`
	Duration float64 `json:"duration"`
	End string `json:"end"`
	Route []RootDataWorkoutsRoute `json:"route"`
}

type RootData struct {
	Workouts []RootDataWorkouts `json:"workouts"`
	Metrics []RootDataMetrics `json:"metrics"`
}

type Root struct {
	Data RootData `json:"data"`
}
