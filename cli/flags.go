// cli/flags.go
package cli

import (
	"fitness/models"
	"fitness/printer"
	"flag"
	"fmt"
	"os"
	"strings"
)

type CLIFlags struct {
	MaxItems    int
	Compact     bool
	TimeFormat  string
	FilterType  string
	FilterValue string
	SortBy      string
	SortDesc    bool
	DataType    string
	Include     string
	Exclude     string
}

func ParseFlags() CLIFlags {
	flags := CLIFlags{}

	// Basic display options
	flag.IntVar(&flags.MaxItems, "n", 0, "Maximum number of items to display (0 for all)")
	flag.BoolVar(&flags.Compact, "c", false, "Use compact display mode")
	flag.StringVar(&flags.TimeFormat, "time-format", "2006-01-02 15:04:05", "Time format string")

	// Filtering options
	flag.StringVar(&flags.FilterType, "f", "", "Filter type (name, distance, duration, energy)")
	flag.StringVar(&flags.FilterValue, "v", "", "Filter value")

	// Sorting options
	flag.StringVar(&flags.SortBy, "sort", "", "Sort by field (name, date, duration, distance, energy)")
	flag.BoolVar(&flags.SortDesc, "desc", false, "Sort in descending order")

	// Data type selection
	flag.StringVar(&flags.DataType, "type", "workouts", "Data type to display (workouts or metrics)")

	// Field inclusion/exclusion
	flag.StringVar(&flags.Include, "i", "", "Include only specific fields (comma-separated)")
	flag.StringVar(&flags.Exclude, "x", "", "Exclude specific fields (comma-separated)")

	// Add custom usage message
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
	}

	flag.Parse()
	return flags
}

func CreateFilterFunction(flags CLIFlags) printer.FilterFunc {
	if flags.FilterType == "" || flags.FilterValue == "" {
		return nil
	}

	return func(v interface{}) bool {
		if w, ok := v.(models.Workout); ok {
			switch flags.FilterType {
			case "name":
				return strings.Contains(strings.ToLower(w.Name),
					strings.ToLower(flags.FilterValue))
			case "distance":
				if w.Distance != nil {
					val := w.Distance.Qty
					return fmt.Sprintf("%.1f", val) == flags.FilterValue
				}
			case "duration":
				return fmt.Sprintf("%.1f", w.Duration) == flags.FilterValue
			case "energy":
				if w.ActiveEnergyBurned != nil {
					val := w.ActiveEnergyBurned.Qty
					return fmt.Sprintf("%.1f", val) == flags.FilterValue
				}
			}
		}
		return false
	}
}

func CreatePrintOptions(flags CLIFlags) printer.PrintOptions {
	opts := printer.DefaultPrintOptions()
	opts.TimeFormat = flags.TimeFormat
	opts.MaxItems = flags.MaxItems
	opts.Compact = flags.Compact
	opts.Filter = CreateFilterFunction(flags)

	// Handle include/exclude fields
	if flags.Include != "" {
		opts.IncludeFields = strings.Split(flags.Include, ",")
		for i := range opts.IncludeFields {
			opts.IncludeFields[i] = strings.TrimSpace(opts.IncludeFields[i])
		}
	}
	if flags.Exclude != "" {
		opts.ExcludeFields = strings.Split(flags.Exclude, ",")
		for i := range opts.ExcludeFields {
			opts.ExcludeFields[i] = strings.TrimSpace(opts.ExcludeFields[i])
		}
	}

	return opts
}
