package internal

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

// TraceAndTimeRequests sends HTTP requests as per the provided configuration,
// traces them, and measures response times.
func TraceAndTimeRequests(config Config) {
	// OpenTelemetryのトレーサープロバイダーをセットアップ
	tp, err := setupTracer(config.Endpoint)
	if err != nil {
		fmt.Printf("Error setting up tracer: %v\n", err)
		return
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			fmt.Printf("Error shutting down tracer: %v\n", err)
		}
	}()

	client := &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, config.Concurrency)

	for i := 0; i < config.Count; i++ {
		wg.Add(1)
		semaphore <- struct{}{} // Acquire a slot

		go func(i int) {
			defer wg.Done()
			bodyReader := strings.NewReader(config.Body)
			req, err := http.NewRequest(config.Method, config.RequestURL, bodyReader)
			if err != nil {
				fmt.Printf("Error creating request: %v\n", err)
				<-semaphore // Release the slot
				return
			}

			for _, h := range config.Headers {
				split := strings.SplitN(h, ":", 2)
				if len(split) == 2 {
					req.Header.Add(strings.TrimSpace(split[0]), strings.TrimSpace(split[1]))
				}
			}

			ctx, span := otel.Tracer("tmcurl").
				Start(req.Context(), fmt.Sprintf("http-request-%d", i))
			req = req.WithContext(ctx)

			start := time.Now()
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("Error sending request %d: %v\n", i, err)
				span.End()
				<-semaphore // Release the slot
				return
			}
			resp.Body.Close()
			span.End()

			duration := time.Since(start)
			fmt.Printf("Response time for request %d: %v\n", i+1, duration)
			<-semaphore // Release the slot
		}(i)
	}

	wg.Wait() // Wait for all requests to complete
}
