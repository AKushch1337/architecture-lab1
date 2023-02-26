package main

import (
	"encoding/json"
	"fmt"
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
	currentTime := time.Now()
	timeWithTimestamp := currentTime.Format(time.RFC3339)
	data := map[string]string{
		"time": timeWithTimestamp,
	}
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	if req.Method == "GET" {
		fmt.Fprintf(w, string(b))
	} else {
		fmt.Fprintf(w, "Sorry, only GET request is supported.")
	}
}
