// printer/options.go
package printer

type PrintOptions struct {
	TimeFormat    string
	Filter        FilterFunc
	MaxItems      int
	Compact       bool
	SortBy        string
	IncludeFields []string
	ExcludeFields []string
}

type FilterFunc func(interface{}) bool

func DefaultPrintOptions() PrintOptions {
	return PrintOptions{
		TimeFormat: "2006-01-02 15:04:05",
		MaxItems:   0,
		Compact:    false,
	}
}
