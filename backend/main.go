package main

import (
	"encoding/json" // Import for JSON handling
	"fmt"
	"net/http"
)

// Message struct to hold our API response
type Message struct {
	Content string `json:"message"`
}

// handler for the root path, kept for basic check
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go Backend!")
}

// messageHandler for the /api/message endpoint
func messageHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins for now
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight OPTIONS request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Create a message
	msg := Message{Content: "Hello from Go Backend API!"}

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Encode the message to JSON and write to response
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Register handlers
	http.HandleFunc("/", rootHandler) // Keep root handler
	http.HandleFunc("/api/message", messageHandler)

	fmt.Println("Go backend server starting on port 8080")
	http.ListenAndServe(":8080", nil)
}
