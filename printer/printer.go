package printer

import (
	"fmt"
	"strings"
	"time"

	"fitness/config"
	"fitness/models"
	"fitness/utils"
)

// PrintHealthData is a generic function that handles printing both workout and metric data.
// It uses type switching to determine which specific print function to call.
// Parameters:
//   - data: interface{} that should be either []Workout or []models.Metric
//   - opts: PrintOptions containing formatting and filtering preferences
func PrintHealthData(data interface{}, opts PrintOptions) error {
	// Type switch to determine what kind of data we're dealing with
	switch v := data.(type) {
	case []models.Workout:
		return PrintWorkouts(v, opts)
	case []models.Metric:
		return PrintMetrics(v, opts)
	default:
		return fmt.Errorf("unsupported data type")
	}
}

// printWorkouts handles the formatting and display of workout data
// It applies filters, respects maximum items limit, and handles both detailed
// and compact display modes
func PrintWorkouts(workouts []models.Workout, opts PrintOptions) error {
	// Apply any filtering if a filter function is provided
	if opts.Filter != nil {
		var filtered []models.Workout
		for _, w := range workouts {
			if opts.Filter(w) {
				filtered = append(filtered, w)
			}
		}
		workouts = filtered
	}

	// Respect the MaxItems limit if set
	if opts.MaxItems > 0 && len(workouts) > opts.MaxItems {
		workouts = workouts[:opts.MaxItems]
	}

	// Use compact display mode if specified
	if opts.Compact {
		return PrintWorkoutsCompact(workouts, opts)
	}

	// Detailed display mode
	fmt.Println()
	fmt.Println("Workout Data:")
	fmt.Println(strings.Repeat("-", 80)) // Print a divider line

	// Iterate through workouts and print detailed information
	for i, w := range workouts {
		// Print divider between workouts
		if i > 0 {
			fmt.Println(strings.Repeat("-", 80))
		}

		// Print basic workout information
		fmt.Printf("Workout: %s\n", w.Name)
		fmt.Printf("ID: %s\n", w.ID)

		// Parse and format the start and end times from RFC3339 format
		start, err := time.Parse(config.TimeFormat, w.Start)
		if err != nil {
			fmt.Printf("Error parsing start time %q: %v\n", w.Start, err)
		}
		end, err := time.Parse(config.TimeFormat, w.End)
		if err != nil {
			fmt.Printf("Error parsing end time %q: %v\n", w.End, err)
		}
		fmt.Printf("Start: %s\n", start.Format(opts.TimeFormat))
		fmt.Printf("End: %s\n", end.Format(opts.TimeFormat))
		fmt.Printf("Duration: %.2f minutes\n", w.Duration)

		// Print optional fields if they exist
		// Each if statement checks for nil before attempting to print
		if w.Distance != nil {
			fmt.Printf("Distance: %.2f %s\n", w.Distance.Qty, w.Distance.Units)
		}
		if w.ActiveEnergyBurned != nil {
			fmt.Printf("Energy Burned: %.2f %s\n", w.ActiveEnergyBurned.Qty, w.ActiveEnergyBurned.Units)
		}
		if w.Intensity != nil {
			fmt.Printf("Intensity: %.2f %s\n", w.Intensity.Qty, w.Intensity.Units)
		}
		if w.Location != nil {
			fmt.Printf("Location: %s\n", *w.Location)
		}
		if w.Temperature != nil {
			fmt.Printf("Temperature: %.1f %s\n", w.Temperature.Qty, w.Temperature.Units)
		}
		fmt.Println()
	}

	return nil
}

// printWorkoutsCompact provides a condensed tabular view of workout data
// It shows only the most important fields in a space-efficient format
func PrintWorkoutsCompact(workouts []models.Workout, opts PrintOptions) error {
	// Print table header with fixed column widths
	fmt.Printf("%-20s %-19s %-8s %-10s %-10s\n",
		"Name", "Start", "Duration", "Distance", "Energy")
	fmt.Println(strings.Repeat("-", 72))

	// Print each workout as a single line
	for _, w := range workouts {
		// Parse the start time
		start, err := time.Parse(config.TimeFormat, w.Start)
		if err != nil {
			fmt.Printf("Error parsing start time %q: %v\n", w.Start, err)
		}

		// Handle optional distance field
		distance := "-" // Default value if nil
		if w.Distance != nil {
			distance = fmt.Sprintf("%.1f%s", w.Distance.Qty, w.Distance.Units)
		}

		// Handle optional energy field
		energy := "-" // Default value if nil
		if w.ActiveEnergyBurned != nil {
			energy = fmt.Sprintf("%.0f%s", w.ActiveEnergyBurned.Qty, w.ActiveEnergyBurned.Units)
		}

		// Print the row with fixed column widths
		fmt.Printf("%-20s %-19s %-8.1f %-10s %-10s\n",
			utils.Truncate(w.Name, 20),       // Truncate name if too long
			start.Format("2006-01-02 15:04"), // Format date
			w.Duration,
			distance,
			energy)
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
