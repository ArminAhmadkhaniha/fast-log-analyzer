# Fast Log Analyzer

A high-performance CLI tool written in Go to process large server log files.

## Problem
Processing gigabytes of server logs to find specific error patterns is slow with interpreted languages. This tool utilizes Go's concurrency to process logs in parallel.

## Features
- Reads large text files line-by-line (memory efficient).
- Uses Goroutines (Worker Pool pattern) to process lines concurrently.
- Filters for specific keywords (e.g., "ERROR", "INFO").
- Generates a summary report of error counts.

## Usage
go run main.go -file=server.log -workers=4

## Benchmarks & Lessons Learned

### Performance Results (10 Million Records)
- **Python (Sequential):** ~0.70s
- **Go (Concurrent Worker Pool):** ~2.05s

### Engineering Insight
Surprisingly, the Python script outperformed the concurrent Go implementation for this specific task. 

**Why?** The workload (simple string matching) was too lightweight to justify the overhead of Go's concurrency primitives (Channels & Context Switching). The cost of synchronizing Goroutines exceeded the time saved by parallel execution.

**Conclusion:** Go's worker pool pattern shines in **CPU-intensive** tasks (hashing, encryption) or **I/O-bound** tasks with high latency (network requests), but for simple sequential file reading, the overhead of concurrency can introduce latency.