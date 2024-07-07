package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"
)

type PongResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := PongResponse{
		Message:   "pong",
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	addr := flag.String("addr", ":4141", "Port to run the health server on")
	flag.Parse()

	http.HandleFunc("/healthz", healthHandler)

	fmt.Println(fmt.Sprintf("Starting server on %v", *addr))

	if err := http.ListenAndServe(*addr, nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
