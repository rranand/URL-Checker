# 🚀 URL Checker (Go Concurrency Project)

A high-performance URL health checker built in Go using goroutines, channels, and worker pools.

This project demonstrates practical concurrency patterns in Go while validating the availability, status, and behavior of multiple URLs in parallel.

---

## ✨ Features

* ⚡ Concurrent URL checking using worker pools
* 🔁 Handles redirects automatically
* 📊 Supports multiple HTTP status categories:
  * 2xx (Success)
  * 4xx (Client errors)
  * 5xx (Server errors)
* ⏱ Handles delayed responses
* 🔒 Detects SSL certificate issues
* ❌ Handles invalid URLs and unreachable hosts
* 🧵 Efficient synchronization with channels & WaitGroups

---

## 🧠 Concepts Covered

This project is great for learning:

* Goroutines
* Channels (buffered)
* Worker Pool pattern
* Synchronization using `sync.WaitGroup`
* Handling race conditions
* Concurrent system design

---

## 🏗️ Project Structure

```
.
├── go.mod
├── internal
│   ├── model
│   │   └── result.go
│   ├── urlchecker
│   │   ├── config.go
│   │   ├── urlchecker.go
│   │   └── urlchecker_test.go
│   └── worker
│       └── worker.go
├── main.go
└── readme.md
```

---

## ⚙️ How It Works

1. A list of URLs is defined in `main.go`
2. Worker pool is initialized
3. URLs are sent into `url channel`
4. Workers pick URLs and perform health checks
5. Results are pushed into `result channel`
6. Main goroutine prints results continuously

---

## 🔄 Data Flow

```
URLs → url channel → Workers → result channel → Output
```

---

## ▶️ Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/rranand/URL-Checker.git
cd URL-Checker
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the project

```bash
go run main.go
```

---

## ⚡ Configuration

You can modify these values in `main.go`:

```go
workerPoolSize := 10
buffer := workerPoolSize * 2
```

### 🔍 Explanation

* **workerPoolSize** → number of concurrent workers
* **buffer** → size of channel buffer

---

## 🧵 Concurrency Design Insights

### ✅ Buffer ≥ Worker Pool

* Maximizes throughput
* Workers stay busy

### ⚠️ Buffer < Worker Pool

* Some workers idle
* Reduced efficiency

---

## ⚠️ Important Note

The program **does NOT rely only on WaitGroup**.

Instead, it:

* Listens to `result channel`
* Ensures all results are printed before exit

This avoids race conditions between:

* Worker completion
* Output processing

---

## 🧪 Test Cases Included

The project tests multiple real-world scenarios:

* ✅ Valid URLs
* 🔁 Redirect chains
* ❌ 4xx errors
* 💥 5xx errors
* ⏱ Delayed responses
* 🔒 SSL issues
* 🚫 Invalid domains

---

## 📄 License

MIT License

---

## 👨‍💻 Author

[@rranand](https://github.com/rranand)
