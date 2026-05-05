# 🚀 URL Checker (Concurrent Health Monitoring System in Go)

A high-performance, configurable URL health checker built in Go using **goroutines, channels, worker pools, and context-based cancellation**.

This project simulates a **real-world concurrent monitoring system**, capable of validating large sets of URLs efficiently with structured logging and graceful shutdown.

---

## ✨ Key Features

* ⚡ Concurrent URL checking using worker pool pattern
* 📂 Reads URLs from external file (`-filename`)
* ⚙️ Config-driven system using JSON (`-config`)
* 🧵 Efficient concurrency with goroutines & channels
* ⏱ Latency measurement for each request
* 🔁 Automatic redirect handling
* 🔒 Detects SSL/TLS and network failures
* ❌ Handles invalid URLs and unreachable hosts
* 📊 Categorized responses:

  * 2xx → Success
  * 3xx → Redirect
  * 4xx → Client Error
  * 5xx → Server Error
* 🪵 Structured logging using `slog`
* 📁 Logs written to:

  * Console
  * Output file (`output-<filename>.txt`)
* 🛑 Graceful shutdown using OS signals (`SIGINT`, `SIGTERM`)
* 🔄 Context-based cancellation for safe worker termination

---

## 🧠 Concepts Demonstrated

* Goroutines & concurrency primitives
* Worker Pool pattern
* Buffered vs unbuffered channels
* Context cancellation (`context.WithCancel`)
* Graceful shutdown handling
* Channel lifecycle management
* Structured logging (`slog`)
* File I/O (streaming input)
* Real-world concurrent system design

---

## 🏗️ Project Structure

```text
.
├── config.json
├── go.mod
├── internal
│   ├── model
│   │   ├── config.go
│   │   └── result.go
│   ├── urlchecker
│   │   ├── urlchecker.go
│   │   └── urlchecker_test.go
│   └── worker
│       └── worker.go
├── main.go
└── readme.md
```

---

## ⚙️ How It Works

1. Config is loaded from `config.json`
2. URLs are streamed from input file
3. URLs are pushed into `url channel`
4. Worker pool processes URLs concurrently
5. Each worker:

   * sends HTTP request
   * measures latency
   * classifies response
6. Results are sent to `result channel`
7. Main goroutine consumes results and logs output

---

## 🔄 Data Flow Architecture

```text
          +----------------------+
          |   URL File Input     |
          +----------+-----------+
                     ↓
          +----------------------+
          |     URL Channel      |
          +----------+-----------+
                     ↓
        +-------------------------+
        |   Worker Pool (N)       |
        |  ┌───────────────────┐  |
        |  |   URL Checker     |  |
        |  └───────────────────┘  |
        +-----------+-------------+
                    ↓
          +----------------------+
          |   Result Channel     |
          +----------+-----------+
                     ↓
          +----------------------+
          |   Logger (slog)      |
          | Console + File Output|
          +----------------------+
```

---

## ▶️ Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/rranand/URL-Checker.git
cd URL-Checker
```

---

### 2. Install dependencies

```bash
go mod tidy
```

---

### 3. Prepare input file

Example: `urls.txt`

```text
https://google.com
https://github.com
https://amazon.com
https://invalid-url.test
```

---

### 4. Prepare config file

Example: `config.json`

```json
{
  "timeout_duration" : 5,
  "max_idle_conn" : 90,
  "max_idle_conn_per_host" : 10,
  "idle_conn_timeout_duration" : 20,
  "no_of_workers": 10
}
```

---

### 5. Run the project

```bash
go run main.go -filename=urls.txt -config=config.json
```

---

## 📁 Output

Logs are written to:

```
output-urls.txt
```

And also printed to console.

---

## 📊 Sample Output

```text
INFO health check worker_id=3 url=https://google.com status=HEALTHY latency=477ms
INFO health check worker_id=5 url=https://amazon.com status=HEALTHY latency=1.7s
INFO health check worker_id=2 url=https://invalid-url.test status=DOWN latency=120ms
```

---

## 🧵 Concurrency Design Insights

### Worker Pool

* Fixed number of workers prevents overload
* Efficient parallel processing

### Channels

* `urlChan` → distributes work
* `resChan` → collects results

### Context Cancellation

* Stops all workers safely
* Triggered by:

  * error during ingestion
  * OS signals

---

## 🛑 Graceful Shutdown

* Listens for:

  * `SIGINT` (Ctrl+C)
  * `SIGTERM`
* Cancels context → stops workers → completes processing
* Ensures no data loss

---

## ⚠️ Design Considerations

* Avoids premature program exit
* Prevents race conditions between:

  * worker completion
  * result consumption
* Ensures connection reuse by draining HTTP response body
* Handles network instability gracefully

---

## 🧪 Test Coverage

Includes real-world scenarios:

* ✅ Valid URLs
* 🔁 Redirect chains
* ❌ Client errors (4xx)
* 💥 Server errors (5xx)
* ⏱ Slow responses / timeouts
* 🔒 SSL/TLS failures
* 🚫 Invalid domains

---

## 🚀 Future Improvements

* Retry with exponential backoff
* Rate limiting
* Metrics (Prometheus / Grafana)
* CLI enhancements
* Distributed worker nodes
* JSON logging mode

---

## 📄 License

MIT License

---

## 👨‍💻 Author

[@rranand](https://github.com/rranand)
