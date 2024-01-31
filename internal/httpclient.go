package internal

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// TraceAndTimeRequests sends HTTP requests as per the provided configuration,
// traces them, and measures response times.
func TraceAndTimeRequests(config Config) {
	client := &http.Client{}

	for i := 0; i < config.Count; i++ {
		var bodyReader io.Reader
		if config.Body != "" {
			bodyReader = strings.NewReader(config.Body)
		}

		req, err := http.NewRequest(config.Method, config.RequestURL, bodyReader)
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			continue
		}

		for _, h := range config.Headers {
			split := strings.SplitN(h, ":", 2)
			if len(split) == 2 {
				req.Header.Add(strings.TrimSpace(split[0]), strings.TrimSpace(split[1]))
			}
		}

		start := time.Now()
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error sending request: %v\n", err)
			continue
		}
		resp.Body.Close()

		duration := time.Since(start)
		fmt.Printf("Response time for request %d: %v\n", i+1, duration)
	}
}
