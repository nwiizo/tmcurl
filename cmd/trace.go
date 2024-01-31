/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/nwiizo/tmcurl/internal"
	"github.com/spf13/cobra"
)

var (
	endpoint   string
	requestURL string
	method     string
	headers    []string
	body       string
	count      int
	// その他のフラグ変数...
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace and time HTTP requests",
	Long: `tmcurl trace command sends HTTP requests and traces them using OpenTelemetry.
You can customize the request with various options like setting headers, request method, and body. 
This command also measures the response time for each request and supports repeating the request multiple times.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := internal.Config{
			Endpoint:   endpoint,
			RequestURL: requestURL,
			Method:     method,
			Headers:    headers,
			Body:       body,
			Count:      count,
		}
		internal.TraceAndTimeRequests(config)
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)

	traceCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "localhost:4317", "OTLP exporter endpoint")
	traceCmd.Flags().StringVarP(&requestURL, "url", "u", "", "URL to send the HTTP request to")
	traceCmd.Flags().StringVarP(&method, "method", "m", "GET", "HTTP method to use")
	traceCmd.Flags().StringArrayVarP(&headers, "header", "H", []string{}, "HTTP headers to include in the request")
	traceCmd.Flags().StringVarP(&body, "body", "b", "", "HTTP request body")
	traceCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of times to send the request")
	// その他のフラグの初期化...
}
