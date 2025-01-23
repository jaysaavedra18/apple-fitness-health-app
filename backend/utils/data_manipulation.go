package utils

import (
	"fitness/config"
	"fitness/models"
	"time"
)

func CalculateWorkoutsPerMonth(workouts []models.Workout) map[string]int {
	workoutsPerMonth := make(map[string]int)
	for _, workout := range workouts {
		if startTime, err := time.Parse(config.TimeFormat, workout.Start); err == nil {
			workoutsPerMonth[startTime.Format("2006-01")]++
		}
	}

	return workoutsPerMonth
}

func CalculateDistancePerWorkout(workouts []models.Workout) map[string]float64 {
	distancePerWorkout := make(map[string]float64)
	for _, workout := range workouts {
		if workout.Distance != nil {
			distancePerWorkout[workout.Name] += workout.Distance.Qty
		}
	}
	return distancePerWorkout
}

func CalculateDistancePerWeek(workouts []models.Workout) map[string]float64 {
	return aggregateByWeek(workouts, func(w models.Workout) float64 {
		if w.Distance != nil {
			return w.Distance.Qty
		}
		return 0
	})
}

func CalculateEnergyPerWeek(workouts []models.Workout) map[string]float64 {
	return aggregateByWeek(workouts, func(w models.Workout) float64 {
		if w.ActiveEnergyBurned != nil {
			return w.ActiveEnergyBurned.Qty
		}
		return 0
	})
}

func aggregateByWeek(workouts []models.Workout, getValue func(models.Workout) float64) map[string]float64 {
	result := make(map[string]float64)
	for _, workout := range workouts {
		if startTime, err := time.Parse(config.TimeFormat, workout.Start); err == nil {
			weekStart := startTime.AddDate(0, 0, -int(startTime.Weekday()-time.Monday))
			weekOf := weekStart.Format(config.DateFormat)
			result[weekOf] += getValue(workout)
		}
	}
	return result
}
