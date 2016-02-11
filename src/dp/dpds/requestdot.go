package dpds

type RequestDot struct {
	Dot
}

// Initializes the context.
func (rd *RequestDot) Init(queryParams map[string][]string, currentSubRoute string, routeComplete string, result string) {
}

// Get result
func (rd *RequestDot) GetResult() string {
	return ""
}
