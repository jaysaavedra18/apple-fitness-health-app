package printer

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"fitness/config"
	"fitness/models"
	"fitness/utils"
)

// Print health data to the console
func PrintHealthData(data interface{}, opts PrintOptions) error {
	// Type switch to determine what kind of data we're dealing with
	switch v := data.(type) {
	case []models.Workout: // Use workouts print if we have workout
		return PrintWorkouts(v, opts)
	case []models.Metric: // Use metrics print if we have metrics
		return PrintMetrics(v, opts)
	default: // Return an error if we don't recognize the data type
		return fmt.Errorf("unsupported data type")
	}
}

// PrintWorkouts function with exclude field support
func PrintWorkouts(workouts []models.Workout, opts PrintOptions) error {
	// Create a map of excluded fields for faster lookup
	excludedFields := make(map[string]bool)
	for _, field := range opts.ExcludeFields {
		excludedFields[strings.ToLower(field)] = true
	}

	// Filter the workouts if a filter function is provided
	if opts.Filter != nil {
		var filtered []models.Workout
		for _, w := range workouts {
			if opts.Filter(w) {
				filtered = append(filtered, w)
			}
		}
		workouts = filtered
	}

	// Limit the number of items displayed if specified
	if opts.MaxItems > 0 && len(workouts) > opts.MaxItems {
		workouts = workouts[:opts.MaxItems]
	}

	// Use compact mode if specified
	if opts.Compact {
		return PrintWorkoutsCompact(workouts, opts)
	}

	// Detailed display mode
	fmt.Println()
	fmt.Println("Workout Data:")
	fmt.Println(strings.Repeat("-", 80))

	// Print each workout data with formatted fields
	for i, w := range workouts {
		if i > 0 {
			fmt.Println(strings.Repeat("-", 80))
		}

		// Only print fields that aren't excluded
		if !excludedFields["name"] {
			fmt.Printf("Workout: %s\n", w.Name)
		}
		if !excludedFields["id"] {
			fmt.Printf("ID: %s\n", w.ID)
		}

		// Handle start/end times
		if !excludedFields["start"] && !excludedFields["time"] {
			start, err := time.Parse(config.TimeFormat, w.Start)
			if err == nil {
				fmt.Printf("Start: %s\n", start.Format(opts.TimeFormat))
			}
		}
		if !excludedFields["end"] && !excludedFields["time"] {
			end, err := time.Parse(config.TimeFormat, w.End)
			if err == nil {
				fmt.Printf("End: %s\n", end.Format(opts.TimeFormat))
			}
		}
		// Handle duration, distance, energy, intensity, location, and temperature
		if !excludedFields["duration"] {
			fmt.Printf("Duration: %s\n", utils.FormatTime(w.Duration))
		}
		if !excludedFields["distance"] && w.Distance != nil {
			fmt.Printf("Distance: %.2f %s\n", w.Distance.Qty, w.Distance.Units)
		}
		if !excludedFields["energy"] && w.ActiveEnergyBurned != nil {
			fmt.Printf("Energy Burned: %.2f %s\n", w.ActiveEnergyBurned.Qty, w.ActiveEnergyBurned.Units)
		}
		if !excludedFields["intensity"] && w.Intensity != nil {
			fmt.Printf("Intensity: %.2f %s\n", w.Intensity.Qty, w.Intensity.Units)
		}
		if !excludedFields["location"] && w.Location != nil {
			fmt.Printf("Location: %s\n", *w.Location)
		}
		if !excludedFields["temperature"] && w.Temperature != nil {
			fmt.Printf("Temperature: %.1f %s\n", w.Temperature.Qty, w.Temperature.Units)
		}
		fmt.Println()
	}

	return nil
}

// PrintWorkoutsCompact with exclude field support
func PrintWorkoutsCompact(workouts []models.Workout, opts PrintOptions) error {
	// Create excluded fields map
	excludedFields := make(map[string]bool)
	for _, field := range opts.ExcludeFields {
		excludedFields[strings.ToLower(field)] = true
	}

	// Determine which columns to show
	showName := !excludedFields["name"]
	showStart := !excludedFields["start"] && !excludedFields["time"]
	showDuration := !excludedFields["duration"]
	showDistance := !excludedFields["distance"]
	showEnergy := !excludedFields["energy"]

	// Print headers incl. name, start, duration, distance, and energy
	var headers []string
	if showName {
		headers = append(headers, fmt.Sprintf("%-20s", "Name"))
	}
	if showStart {
		headers = append(headers, fmt.Sprintf("%-19s", "Start"))
	}
	if showDuration {
		headers = append(headers, fmt.Sprintf("%-8s", "Duration"))
	}
	if showDistance {
		headers = append(headers, fmt.Sprintf("%-10s", "Distance"))
	}
	if showEnergy {
		headers = append(headers, fmt.Sprintf("%-10s", "Energy"))
	}
	// Print padded headers and separator
	fmt.Println(strings.Join(headers, " "))
	fmt.Println(strings.Repeat("-", len(strings.Join(headers, " "))))

	// Print each workout
	for _, w := range workouts {
		// Only print fields that aren't excluded
		var fields []string

		// Handle name, start, duration, distance, and energy
		if showName {
			fields = append(fields, fmt.Sprintf("%-20s", utils.Truncate(w.Name, 20)))
		}
		if showStart {
			start, err := time.Parse(config.TimeFormat, w.Start)
			startStr := "-"
			if err == nil {
				startStr = start.Format("2006-01-02 15:04")
			}
			fields = append(fields, fmt.Sprintf("%-19s", startStr))
		}
		if showDuration {
			fields = append(fields, fmt.Sprintf("%-8s", utils.FormatTime(w.Duration)))
		}
		if showDistance {
			distance := "-"
			if w.Distance != nil {
				distance = fmt.Sprintf("%.1f%s", w.Distance.Qty, w.Distance.Units)
			}
			fields = append(fields, fmt.Sprintf("%-10s", distance))
		}
		if showEnergy {
			energy := "-"
			if w.ActiveEnergyBurned != nil {
				energy = fmt.Sprintf("%.0f%s", w.ActiveEnergyBurned.Qty, w.ActiveEnergyBurned.Units)
			}
			fields = append(fields, fmt.Sprintf("%-10s", energy))
		}
		// Print padded fields
		fmt.Println(strings.Join(fields, " "))
	}
	return nil
}

// printMetrics handles the display of metric data
// It follows similar patterns to printWorkouts for filtering and limiting
func PrintMetrics(metrics []models.Metric, opts PrintOptions) error {
	// Apply filtering if specified
	if opts.Filter != nil {
		var filtered []models.Metric
		for _, m := range metrics {
			if opts.Filter(m) {
				filtered = append(filtered, m)
			}
		}
		metrics = filtered
	}

	// Apply maximum items limit if specified
	if opts.MaxItems > 0 && len(metrics) > opts.MaxItems {
		metrics = metrics[:opts.MaxItems]
	}

	// Print each metric and its data points
	for i, m := range metrics {
		if i > 0 {
			fmt.Println()
		}
		fmt.Printf("Metric: %s (%s)\n", m.Name, m.Units)
		fmt.Println(strings.Repeat("-", 40))

		// Print each data point with formatted date
		for _, d := range m.Data {
			date, _ := time.Parse(time.RFC3339, d.Date)
			fmt.Printf("%s: %.2f\n", date.Format(opts.TimeFormat), d.Qty)
		}
	}
	return nil
}

// PrintCustom flags incl. workoutsPerMonth
func PrintCustom(workouts []models.Workout, opts PrintOptions) {

	// If flag is present print the total workouts per month
	if opts.WorkoutsPerMonth {
		// Calculate total workouts per month
		workoutsPerMonth := utils.CalculateWorkoutsPerMonth(workouts)

		// Create a slice to hold the months (for sorting)
		var months []string
		for month := range workoutsPerMonth {
			months = append(months, month)
		}

		// Sort months based on desc flag
		if opts.SortDesc {
			sort.Sort(sort.Reverse(sort.StringSlice(months))) // Sort descending
		} else {
			sort.Strings(months) // Sort ascending
		}

		// Print the total workouts per month
		fmt.Println()
		fmt.Println("Workouts per Month:")
		fmt.Println(strings.Repeat("-", 40))
		for _, month := range months {
			fmt.Printf("%s: %d\n", month, workoutsPerMonth[month])
		}
		fmt.Println()
	}

	// If flag is present print the distance per workout
	if opts.DistancePerWorkout {
		// Calculate distance per workout
		distancePerWorkout := utils.CalculateDistancePerWorkout(workouts)

		// Print the distance per workout
		fmt.Println()
		fmt.Println("Distance per Workout:")
		fmt.Println(strings.Repeat("-", 40))
		fmt.Printf("%-20s %-20s\n", "Workout", "Distance")
		fmt.Println(strings.Repeat("-", 40))
		for workoutName, totalDistance := range distancePerWorkout {
			fmt.Printf("%-20s %-7.2f miles\n", workoutName, totalDistance)
		}
		fmt.Println()
	}

	// If flag is present print the total energy burned per week
	if opts.EnergyPerWeek {
		// Calculate total energy burned per week
		energyPerWeek := utils.CalculateEnergyPerWeek(workouts)
		// Create a slice to hold the weeks (for sorting)
		var weeks []string
		for week := range energyPerWeek {
			weeks = append(weeks, week)
		}
		// Sort weeks in order based on desc flag
		if opts.SortDesc {
			sort.Sort(sort.Reverse(sort.StringSlice(weeks)))
		} else {
			sort.Sort(sort.StringSlice(weeks))
		}

		// Print the energy burned per week in the sorted order
		fmt.Println()
		fmt.Println("Energy Burned per Week:")
		fmt.Println(strings.Repeat("-", 40))
		for _, week := range weeks {
			fmt.Printf("Week Of %s: %.2f kcal\n", week, energyPerWeek[week])
		}
		fmt.Println()

	}
}
