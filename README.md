# tmcurl 🚀

`tmcurl` is a CLI tool that traces HTTP requests and measures response times, leveraging the power of OpenTelemetry 🌐. It provides insightful performance data for your HTTP requests.

## Features 🌟

- Trace HTTP requests with OpenTelemetry. 🕵️‍♂️
- Measure response times. ⏱
- Customize request methods, headers, and body. 🛠
- Execute requests multiple times. 🔁
- Configure OTLP exporter endpoint. 📡

## Installation 📦

`tmcurl` is written in Go. Ensure you have Go installed and then run the following command:

```sh
go get github.com/nwiizo/tmcurl
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
