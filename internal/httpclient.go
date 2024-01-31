package internal

import (
	"context"
	"fmt"
	"net/http"
	"strings"
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

	for i := 0; i < config.Count; i++ {
		bodyReader := strings.NewReader(config.Body)
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

		ctx, span := otel.Tracer("tmcurl").Start(req.Context(), "http-request")
		req = req.WithContext(ctx)

		start := time.Now()
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error sending request: %v\n", err)
			span.End()
			continue
		}
		resp.Body.Close()
		span.End()

		duration := time.Since(start)
		fmt.Printf("Response time for request %d: %v\n", i+1, duration)
	}
}
