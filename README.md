Here’s a **~100-word README** you can paste directly:

---

# Go CLI Server with Graceful Shutdown

This project is a lightweight Go-based CLI tool that starts an HTTP server and shuts it down gracefully. It demonstrates proper signal handling using `os/signal`, `context`, and `http.Server` to ensure in-flight requests complete before exit.

## Features

* CLI command to start the server
* Configurable port via flags
* Graceful shutdown on SIGINT / SIGTERM
* Clean resource release and logging

## Usage

```bash
go run main.go --port 8080
```

Press `Ctrl+C` to stop the server safely without dropping active connections.

## Use Case

Ideal for learning production-grade server lifecycle management in Go.
