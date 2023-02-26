package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	port    = ":8795"
	apiPath = "/time"
)

func main() {
	http.HandleFunc(apiPath, getTime)

	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

func getTime(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Sorry, only GET request is supported.", http.StatusMethodNotAllowed)
		return
	}

	currentTime := time.Now()
	timeWithTimestamp := currentTime.Format(time.RFC3339)
	data := map[string]string{
		"time": timeWithTimestamp,
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Request handled successfully: method=%s, path=%s, remote_addr=%s", req.Method, req.URL.Path, req.RemoteAddr)
}
