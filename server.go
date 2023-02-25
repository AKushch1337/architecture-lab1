package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", getTime)

	if err := http.ListenAndServe(":8795", nil); err != nil {
		log.Fatal(err)
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
