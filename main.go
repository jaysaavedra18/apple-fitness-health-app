package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

type HealthData struct { // HealthData represents the root JSON structure
	Data        DataCollection `json:"data"`
	LastUpdated *string        `json:"lastUpdated"`
}
type DataCollection struct { // DataCollection holds workouts and metrics
	Workouts []Workout `json:"workouts"`
	Metrics  []Metric  `json:"metrics"`
}
type Measurement struct { // Measurement represents a generic quantity with units
	Units string  `json:"units"`
	Qty   float64 `json:"qty"`
}
type Workout struct { // Workout represents a single workout session
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
type MetricData struct { // MetricData represents a single data point in a metric
	Date string  `json:"date"`
	Qty  float64 `json:"qty"`
}
type Metric struct { // Metric represents a collection of measurements over time
	Name  string       `json:"name"`
	Data  []MetricData `json:"data"`
	Units string       `json:"units"`
}

// Global constants and variables used throughout the program
const (
	dateFormat       = "2006-01-02"
	timeFormat       = "2006-01-02 15:04:05 -0700"
	iCloudDirPath    = "/Users/saavedj/Library/Mobile Documents/com~apple~CloudDocs/health-data"
	cacheFilePath    = "cache.json"
	dateRegexPattern = `\d{4}-\d{2}-\d{2}`
)

var (
	allWorkouts []Workout
	allMetrics  []Metric
)

// Load the cache into the program data
func loadCache(filename string) (*HealthData, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cache HealthData
	if err := json.Unmarshal(data, &cache); err != nil {
		return nil, err
	}

	allWorkouts = append(allWorkouts, cache.Data.Workouts...)
	allMetrics = append(allMetrics, cache.Data.Metrics...)

	return &cache, nil
}

// Load the new directory files into the program data
func loadDirectory(directoryPath string, cacheLastUpdated string) (bool, string, error) {
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return false, cacheLastUpdated, err
	}

	dataWasUpdated := false
	latestFileDate := cacheLastUpdated
	cacheDate, err := time.Parse(dateFormat, cacheLastUpdated)
	if err != nil {
		return false, cacheLastUpdated, err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		// Extract date from filename
		re := regexp.MustCompile(dateRegexPattern)
		matches := re.FindAllString(file.Name(), -1)
		if len(matches) == 0 {
			continue
		}
		fileDate := matches[len(matches)-1]

		// Parse and compare dates
		currentFileDate, err := time.Parse(dateFormat, fileDate)
		if err != nil {
			continue
		}

		// Only process files newer than our cache
		if currentFileDate.After(cacheDate) {
			fmt.Printf("Processing new data from: %s\n", fileDate)

			// Read and parse file
			filePath := directoryPath + "/" + file.Name()
			content, err := os.ReadFile(filePath)
			if err != nil {
				continue
			}

			var fileData HealthData
			if err := json.Unmarshal(content, &fileData); err != nil {
				continue
			}

			// Update our data collections
			allWorkouts = append(allWorkouts, fileData.Data.Workouts...)
			allMetrics = append(allMetrics, fileData.Data.Metrics...)
			dataWasUpdated = true

			// Keep track of the latest file date
			if currentFileDate.After(cacheDate) {
				latestFileDate = fileDate
			}
		}
	}

	return dataWasUpdated, latestFileDate, nil
}

func importData() {
	cache, err := loadCache(cacheFilePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load cache: %v", err))
	}

	// Process directory and get update status
	wasUpdated, latestUpdate, err := loadDirectory(iCloudDirPath, *cache.LastUpdated)
	if err != nil {
		panic(fmt.Sprintf("Failed to load directory: %v", err))
	}

	// Only write to cache if we found new data
	if wasUpdated {
		if err := writeToCache(allWorkouts, allMetrics, &latestUpdate); err != nil {
			panic(fmt.Sprintf("Failed to write cache: %v", err))
		}
		fmt.Printf("Cache updated with data through: %s\n", latestUpdate)
	} else {
		// fmt.Println("No new data found, cache remains current")
	}
}

func writeToCache(allWorkouts []Workout, allMetrics []Metric, lastUpdated *string) error {
	// Create the HealthData structure to match the original format
	healthData := HealthData{
		Data: DataCollection{
			Workouts: allWorkouts,
			Metrics:  allMetrics,
		},
		LastUpdated: lastUpdated,
	}

	// Marshal the HealthData structure into JSON
	data, err := json.MarshalIndent(healthData, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling data: %v", err)
	}

	// Write the JSON data to cache.json
	err = os.WriteFile(cacheFilePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	fmt.Printf("Data written to %s\n", cacheFilePath)
	return nil
}

// PrintOptions configures how data should be printed
type PrintOptions struct {
	TimeFormat    string     // Format for displaying times
	Filter        FilterFunc // Optional filter function
	MaxItems      int        // Maximum items to display (0 for all)
	Compact       bool       // If true, use more compact output
	SortBy        string     // Field to sort by
	IncludeFields []string   // Specific fields to include (empty for all)
	ExcludeFields []string   // Fields to exclude
}

// FilterFunc is a generic filter function type
type FilterFunc func(interface{}) bool

// DefaultPrintOptions returns standard print options
func DefaultPrintOptions() PrintOptions {
	return PrintOptions{
		TimeFormat: "2006-01-02 15:04:05",
		MaxItems:   0,
		Compact:    false,
	}
}

// PrintHealthData is a generic function that handles printing both workout and metric data.
// It uses type switching to determine which specific print function to call.
// Parameters:
//   - data: interface{} that should be either []Workout or []Metric
//   - opts: PrintOptions containing formatting and filtering preferences
func PrintHealthData(data interface{}, opts PrintOptions) error {
	// Type switch to determine what kind of data we're dealing with
	switch v := data.(type) {
	case []Workout:
		return printWorkouts(v, opts)
	case []Metric:
		return printMetrics(v, opts)
	default:
		return fmt.Errorf("unsupported data type")
	}
}

// printWorkouts handles the formatting and display of workout data
// It applies filters, respects maximum items limit, and handles both detailed
// and compact display modes
func printWorkouts(workouts []Workout, opts PrintOptions) error {
	// Apply any filtering if a filter function is provided
	if opts.Filter != nil {
		var filtered []Workout
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
		return printWorkoutsCompact(workouts, opts)
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
		start, err := time.Parse(timeFormat, w.Start)
		if err != nil {
			fmt.Printf("Error parsing start time %q: %v\n", w.Start, err)
		}
		end, err := time.Parse(timeFormat, w.End)
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
func printWorkoutsCompact(workouts []Workout, opts PrintOptions) error {
	// Print table header with fixed column widths
	fmt.Printf("%-20s %-19s %-8s %-10s %-10s\n",
		"Name", "Start", "Duration", "Distance", "Energy")
	fmt.Println(strings.Repeat("-", 72))

	// Print each workout as a single line
	for _, w := range workouts {
		// Parse the start time
		start, err := time.Parse(timeFormat, w.Start)
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
			truncate(w.Name, 20),             // Truncate name if too long
			start.Format("2006-01-02 15:04"), // Format date
			w.Duration,
			distance,
			energy)
	}
	return nil
}

// printMetrics handles the display of metric data
// It follows similar patterns to printWorkouts for filtering and limiting
func printMetrics(metrics []Metric, opts PrintOptions) error {
	// Apply filtering if specified
	if opts.Filter != nil {
		var filtered []Metric
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

// truncate is a helper function that shortens strings that exceed a specified length
// It adds "..." to the end of truncated strings
func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-3] + "..." // Leave room for "..."
}

type CLIFlags struct {
	maxItems    int
	compact     bool
	timeFormat  string
	filterType  string
	filterValue string
	sortBy      string
	sortDesc    bool
	dataType    string
	include     string
	exclude     string
}

func parseFlags() CLIFlags {
	flags := CLIFlags{}

	// Basic display options
	flag.IntVar(&flags.maxItems, "n", 0, "Maximum number of items to display (0 for all)")
	flag.BoolVar(&flags.compact, "c", false, "Use compact display mode")
	flag.StringVar(&flags.timeFormat, "time-format", "2006-01-02 15:04:05", "Time format string")

	// Filtering options
	flag.StringVar(&flags.filterType, "f", "", "Filter type (name, distance, duration, energy)")
	flag.StringVar(&flags.filterValue, "v", "", "Filter value")

	// Sorting options
	flag.StringVar(&flags.sortBy, "sort", "", "Sort by field (name, date, duration, distance, energy)")
	flag.BoolVar(&flags.sortDesc, "desc", false, "Sort in descending order")

	// Data type selection
	flag.StringVar(&flags.dataType, "type", "workouts", "Data type to display (workouts or metrics)")

	// Field inclusion/exclusion
	flag.StringVar(&flags.include, "i", "", "Include only specific fields (comma-separated)")
	flag.StringVar(&flags.exclude, "x", "", "Exclude specific fields (comma-separated)")

	// Add custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of Health Data Printer:\n")
		fmt.Fprintf(os.Stderr, "  health-printer [options]\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  health-printer -n 10 -c                    # Show 10 items in compact mode\n")
		fmt.Fprintf(os.Stderr, "  health-printer -f name -v \"Pool Swim\"      # Show only Pool Swim workouts\n")
		fmt.Fprintf(os.Stderr, "  health-printer -sort duration -desc        # Sort by duration descending\n")
		fmt.Fprintf(os.Stderr, "  health-printer -i \"name,duration,distance\" # Show only specific fields\n")
	}

	flag.Parse()
	return flags
}

func createFilterFunction(flags CLIFlags) FilterFunc {
	if flags.filterType == "" || flags.filterValue == "" {
		return nil
	}

	return func(v interface{}) bool {
		if w, ok := v.(Workout); ok {
			switch flags.filterType {
			case "name":
				return strings.Contains(strings.ToLower(w.Name),
					strings.ToLower(flags.filterValue))
			case "distance":
				if w.Distance != nil {
					val := w.Distance.Qty
					return fmt.Sprintf("%.1f", val) == flags.filterValue
				}
			case "duration":
				return fmt.Sprintf("%.1f", w.Duration) == flags.filterValue
			case "energy":
				if w.ActiveEnergyBurned != nil {
					val := w.ActiveEnergyBurned.Qty
					return fmt.Sprintf("%.1f", val) == flags.filterValue
				}
			}
		}
		return false
	}
}

func createPrintOptions(flags CLIFlags) PrintOptions {
	opts := DefaultPrintOptions()
	opts.TimeFormat = flags.timeFormat
	opts.MaxItems = flags.maxItems
	opts.Compact = flags.compact
	opts.Filter = createFilterFunction(flags)

	// Handle include/exclude fields
	if flags.include != "" {
		opts.IncludeFields = strings.Split(flags.include, ",")
		for i := range opts.IncludeFields {
			opts.IncludeFields[i] = strings.TrimSpace(opts.IncludeFields[i])
		}
	}
	if flags.exclude != "" {
		opts.ExcludeFields = strings.Split(flags.exclude, ",")
		for i := range opts.ExcludeFields {
			opts.ExcludeFields[i] = strings.TrimSpace(opts.ExcludeFields[i])
		}
	}

	return opts
}

func main() {
	flags := parseFlags()
	opts := createPrintOptions(flags)

	// Import data
	importData()
	fmt.Println()

	// Determine which data to display
	var err error
	switch flags.dataType {
	case "workouts":
		err = PrintHealthData(allWorkouts, opts)
	case "metrics":
		err = PrintHealthData(allMetrics, opts)
	default:
		fmt.Fprintf(os.Stderr, "Invalid data type: %s\n", flags.dataType)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
