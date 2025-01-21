// Package cli handles command-line interface functionality
package cli

import (
	"fitness/models"
	"fitness/printer"
	"flag"
	"fmt"
	"os"
	"strings"
)

// CLIFlags stores all command-line flags that can be passed to the application
type CLIFlags struct {
	MaxItems              int    // Maximum number of items to display (0 means show all)
	Compact               bool   // Whether to use compact display mode
	TimeFormat            string // Format string for displaying timestamps
	FilterType            string // Type of filter to apply (name, distance, etc.)
	FilterValue           string // Value to filter by
	SortBy                string // Field to sort results by
	SortDesc              bool   // Whether to sort in descending order
	DataType              string // Type of data to display (workouts/metrics)
	Include               string // Comma-separated list of fields to include
	Exclude               string // Comma-separated list of fields to exclude
	TotalWorkoutsPerMonth bool   // Whether to show total workouts per month
}

// ParseFlags sets up and processes all command-line flags
// Returns: A CLIFlags struct containing all parsed flag values
func ParseFlags() CLIFlags {
	flags := CLIFlags{}

	// Define display-related flags
	flag.IntVar(&flags.MaxItems, "n", 0, "Maximum number of items to display (0 for all)")
	flag.BoolVar(&flags.Compact, "c", false, "Use compact display mode")
	flag.StringVar(&flags.TimeFormat, "time-format", "2006-01-02 15:04:05", "Time format string")

	// Define filtering flags
	flag.StringVar(&flags.FilterType, "f", "", "Filter type (name, distance, duration, energy)")
	flag.StringVar(&flags.FilterValue, "value", "", "Filter value")

	// Define sorting flags
	flag.StringVar(&flags.SortBy, "sort", "", "Sort by field (name, date, duration, distance, energy)")
	flag.BoolVar(&flags.SortDesc, "desc", false, "Sort in descending order")

	// Define data selection flags
	flag.StringVar(&flags.DataType, "type", "workouts", "Data type to display (workouts or metrics)")

	// Define field selection flags
	flag.StringVar(&flags.Include, "i", "", "Include only specific fields (comma-separated)")
	flag.StringVar(&flags.Exclude, "x", "", "Exclude specific fields (comma-separated)")

	// Set up custom usage message with examples
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of Health Fitness Data Printer:\n")
		fmt.Fprintf(os.Stderr, "  fitness [options]\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  fitness -n 10 -c                    # Show 10 items in compact mode\n")
		fmt.Fprintf(os.Stderr, "  fitness -f name -v \"Pool Swim\"      # Show only Pool Swim workouts\n")
		fmt.Fprintf(os.Stderr, "  fitness -sort duration -desc        # Sort by duration descending\n")
		fmt.Fprintf(os.Stderr, "  fitness -i \"name,duration,distance\" # Show only specific fields\n")
		fmt.Println()
	}

	// Define custom flags incl. total workouts per month
	flag.BoolVar(&flags.TotalWorkoutsPerMonth, "total-workouts-per-month", false, "Show total workouts per month")

	// Parse the flags
	flag.Parse()
	return flags
}

// CreateFilterFunction creates a filter function based on the provided flags
// Returns: A FilterFunc that returns true if an item should be included in the output
func CreateFilterFunction(flags CLIFlags) printer.FilterFunc {
	// If no filter type or value is specified, return nil (no filtering)
	if flags.FilterType == "" || flags.FilterValue == "" {
		return nil
	}

	// Return a function that filters Workout objects based on the specified criteria
	// `ok` will be true if the cast is successful
	// `w` will contain the Workout object if the cast is successful
	// `value` is the interface{} type that we are trying to cast
	return func(value interface{}) bool {
		// Try to cast the interface to a Workout type
		if workout, ok := value.(models.Workout); ok {
			switch flags.FilterType {
			case "name":
				// Case-insensitive substring match for workout names
				return strings.Contains(strings.ToLower(workout.Name),
					strings.ToLower(flags.FilterValue))
			case "distance":
				// Match exact distance value if available
				if workout.Distance != nil {
					val := workout.Distance.Qty
					return fmt.Sprintf("%.1f", val) == flags.FilterValue
				}
			case "duration":
				// Match exact duration value
				return fmt.Sprintf("%.1f", workout.Duration) == flags.FilterValue
			case "energy":
				// Match exact energy value if available
				if workout.ActiveEnergyBurned != nil {
					val := workout.ActiveEnergyBurned.Qty
					return fmt.Sprintf("%.1f", val) == flags.FilterValue
				}
			}
		}
		return false
	}
}

// CreatePrintOptions creates a PrintOptions struct based on the provided flags
// Returns: A PrintOptions struct configured according to the command-line flags
func CreatePrintOptions(flags CLIFlags) printer.PrintOptions {
	// Start with default print options
	opts := printer.DefaultPrintOptions()

	// Apply basic display options
	opts.TimeFormat = flags.TimeFormat
	opts.MaxItems = flags.MaxItems
	opts.Compact = flags.Compact
	opts.Filter = CreateFilterFunction(flags)

	// Process included fields if specified
	if flags.Include != "" {
		// Split comma-separated field list and trim whitespace
		opts.IncludeFields = strings.Split(flags.Include, ",")
		for i := range opts.IncludeFields {
			opts.IncludeFields[i] = strings.TrimSpace(opts.IncludeFields[i])
		}
	}

	// Process excluded fields if specified
	if flags.Exclude != "" {
		// Split comma-separated field list and trim whitespace
		opts.ExcludeFields = strings.Split(flags.Exclude, ",")
		for i := range opts.ExcludeFields {
			opts.ExcludeFields[i] = strings.TrimSpace(opts.ExcludeFields[i])
		}
	}

	return opts
}
