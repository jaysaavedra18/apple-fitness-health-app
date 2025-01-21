// printer/options.go
package printer

// PrintOptions contains options for printing data
type PrintOptions struct {
	TimeFormat       string     // Format for time.Time values
	Filter           FilterFunc // Filter function to apply to data
	MaxItems         int        // Maximum number of items to display
	Compact          bool       // Whether to use compact display mode
	SortBy           string     // Field to sort results by
	IncludeFields    []string   // Fields to include in output
	ExcludeFields    []string   // Fields to exclude from output
	WorkoutsPerMonth bool       // Whether to show total workouts per month
}

// FilterFunc is a function type that filters data
type FilterFunc func(interface{}) bool

// DefaultPrintOptions returns a set of default PrintOptions
func DefaultPrintOptions() PrintOptions {
	return PrintOptions{
		TimeFormat: "2006-01-02 15:04:05",
		MaxItems:   0,
		Compact:    false,
	}
}
