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
Certainly! Here's an addition to the README that includes a scenario for running Jaeger with `docker-compose up -d`.

---

## Running with Docker Compose and Jaeger 🐳🔍

You can enhance your tracing experience by running `tmcurl` alongside Jaeger, a powerful tracing system, using Docker Compose. This setup allows you to visualize and manage trace data collected by `tmcurl`.

1. **Add Jaeger Service to Docker Compose**: Include the Jaeger service in your `docker-compose.yml` file to collect and visualize traces. A sample configuration for Jaeger is provided below.

```yaml
services:
  jaeger:
    image: jaegertracing/all-in-one:1.53
    container_name: jaegertracing
    ports:
      - "6831:6831/udp"
      - "6832:6832/udp"
      - "5778:5778"
      - "16686:16686"
      - "4317:4317"
      - "4318:4318"
      - "14250:14250"
      - "14268:14268"
      - "14269:14269"
      - "9411:9411"
    environment:
      - COLLECTOR_ZIPKIN_HOST_PORT=:9411
      - COLLECTOR_OTLP_ENABLED=true
```

This configuration starts Jaeger within a container, exposing necessary ports for collecting traces and displaying them through the Jaeger UI.

2. **Run Docker Compose**: To start `tmcurl` and Jaeger, run the following command in your terminal:

```sh
docker-compose up -d
```

This command launches all services defined in your `docker-compose.yml` file in detached mode, freeing up your terminal.

3. **Access Jaeger UI**: Once Jaeger is running, you can access its UI to view and analyze trace data. By default, the Jaeger UI is available at `http://localhost:16686`.

4. **Send Traces from tmcurl**: Ensure `tmcurl` is configured to send traces to Jaeger by setting the OTLP exporter endpoint to Jaeger's OTLP port (default: `localhost:4317` for gRPC). Traces collected by `tmcurl` will then be visible in the Jaeger UI.

This setup provides an integrated environment for sending HTTP requests with `tmcurl`, collecting traces, and visualizing them through Jaeger's powerful UI, enhancing the observability of your applications.

### Options 🎛

- `--endpoint` (string): The OTLP exporter endpoint (default: "localhost:4317").
- `--url` (string): The URL for the HTTP request.
- `--method` (string): The HTTP method to use (default: "GET").
- `--header` (string array): HTTP request headers.
- `--body` (string): The HTTP request body.
- `--count` (int): The number of times to send the request (default: 1).
- `--concurrency` (int): The number of requests to be made in parallel (default: 1).
