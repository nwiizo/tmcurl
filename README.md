# tmcurl 🚀

`tmcurl` is a CLI tool that traces HTTP requests and measures response times, leveraging the power of OpenTelemetry 🌐. It provides insightful performance data for your HTTP requests and detailed tracing information, including DNS lookup, connection setup, and more.

## Features 🌟

- Trace HTTP requests with OpenTelemetry, including detailed lifecycle events like DNS lookup and TCP connection. 🕵️‍♂️
- Measure response times and aggregate results. ⏱
- Customize request methods, headers, and body. 🛠
- Execute requests multiple times and in parallel. 🔁
- Configure OTLP exporter endpoint for distributed tracing. 📡

## Installation 📦

`tmcurl` is written in Go. Ensure you have Go installed and then run the following command:

```sh
go install github.com/nwiizo/tmcurl
```

## Usage 🚀

Use the `tmcurl` command to send and trace HTTP requests:

```sh
tmcurl trace --url "https://3-shake.com" --method "GET"
```

### Options 🎛

- `--endpoint` (string): The OTLP exporter endpoint (default: "localhost:4317").
- `--url` (string): The URL for the HTTP request.
- `--method` (string): The HTTP method to use (default: "GET").
- `--header` (string array): HTTP request headers.
- `--body` (string): The HTTP request body.
- `--count` (int): The number of times to send the request (default: 1).
- `--concurrency` (int): The number of requests to be made in parallel (default: 1).
