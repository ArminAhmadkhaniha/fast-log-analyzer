# Fast Log Analyzer

A high-performance CLI tool written in Go to process large server log files.

## Problem
Processing gigabytes of server logs to find specific error patterns is slow with interpreted languages. This tool utilizes Go's concurrency to process logs in parallel.

## Features
- Reads large text files line-by-line (memory efficient).
- Uses Goroutines (Worker Pool pattern) to process lines concurrently.
- Filters for specific keywords (e.g., "ERROR", "CRITICAL").
- Generates a summary report of error counts.

## Usage
go run main.go -file=server.log -workers=4