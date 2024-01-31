package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptrace"
	"strings"
	"sync"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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
	var totalDuration time.Duration
	var successfulRequests int
	semaphore := make(chan struct{}, config.Concurrency)

	// 一回の実行で一つの親スパンを作成
	ctx, parentSpan := otel.Tracer("tmcurl").Start(context.Background(), "tmcurl-request")
	defer parentSpan.End()

	for i := 0; i < config.Count; i++ {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(i int) {
			defer wg.Done()
			bodyReader := strings.NewReader(config.Body)
			req, err := http.NewRequest(config.Method, config.RequestURL, bodyReader)
			if err != nil {
				fmt.Printf("Error creating request %d: %v\n", i, err)
				<-semaphore
				return
			}

			for _, h := range config.Headers {
				split := strings.SplitN(h, ":", 2)
				if len(split) == 2 {
					req.Header.Add(strings.TrimSpace(split[0]), strings.TrimSpace(split[1]))
				}
			}

			spanName := fmt.Sprintf("%s %s - request-%d", req.Method, req.URL.Path, i)
			traceCtx, span := otel.Tracer("tmcurl").
				Start(ctx, spanName, trace.WithSpanKind(trace.SpanKindClient))

			// otelhttptraceを使用してHTTPトレースを追加
			traceCtx = httptrace.WithClientTrace(traceCtx, otelhttptrace.NewClientTrace(traceCtx))
			req = req.WithContext(traceCtx)

			start := time.Now()
			resp, err := client.Do(req)
			duration := time.Since(start)
			if err != nil {
				fmt.Printf("Error sending request %d: %v\n", i, err)
				span.RecordError(err)
			} else {
				span.SetAttributes(
					attribute.Int("http.status_code", resp.StatusCode),
					attribute.String("http.url", req.URL.String()),
					attribute.String("http.method", req.Method),
				)
				resp.Body.Close()
				successfulRequests++
			}
			span.End()
			totalDuration += duration
			<-semaphore
		}(i)
	}

	wg.Wait()
	avgDuration := totalDuration / time.Duration(successfulRequests)
	fmt.Printf(
		"Total requests: %d, Successful: %d, Average response time: %v\n",
		config.Count,
		successfulRequests,
		avgDuration,
	)
}
