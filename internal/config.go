package internal

// Config stores the configuration for the HTTP request and tracing
type Config struct {
	Endpoint   string
	RequestURL string
	Method     string
	Headers    []string
	Body       string
	Count      int
}
