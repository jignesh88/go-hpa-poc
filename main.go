package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	// Set up handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/load", loadHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/ready", readyHandler)

	// Get port from env or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Basic info page
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello from HPA Demo App!\n")
	fmt.Fprintf(w, "Hostname: %s\n", hostname)
	fmt.Fprintf(w, "CPU Cores: %d\n\n", runtime.NumCPU())
	fmt.Fprintf(w, "Try /load?duration=10&cores=1 to generate CPU load\n")
	fmt.Fprintf(w, "Health endpoint: /health\n")
	fmt.Fprintf(w, "Readiness endpoint: /ready\n")
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse parameters for the load test
	durationStr := r.URL.Query().Get("duration")
	coresStr := r.URL.Query().Get("cores")

	duration := 10 // default 10 seconds
	cores := 1     // default 1 core

	if durationStr != "" {
		if d, err := strconv.Atoi(durationStr); err == nil && d > 0 && d <= 300 {
			duration = d
		}
	}

	if coresStr != "" {
		if c, err := strconv.Atoi(coresStr); err == nil && c > 0 && c <= runtime.NumCPU() {
			cores = c
		}
	}

	// Log the load test
	log.Printf("Starting load test: %d cores for %d seconds", cores, duration)
	fmt.Fprintf(w, "Starting load test: %d cores for %d seconds\n", cores, duration)

	// Start the load in a goroutine so we can return immediately
	for i := 0; i < cores; i++ {
		go generateLoad(duration)
	}

	fmt.Fprintf(w, "Load test started! Check your metrics to see CPU usage increase.\n")
}

func generateLoad(duration int) {
	// End time for the load test
	endTime := time.Now().Add(time.Duration(duration) * time.Second)

	// Generate CPU load by calculating square roots
	for time.Now().Before(endTime) {
		for i := 0; i < 1000000; i++ {
			math.Sqrt(float64(i))
		}
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Simple health check endpoint
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "healthy")
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	// Simple readiness check endpoint
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ready")
}