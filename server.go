package main

import (
	"net/http"
	"time"
)

func handleRequest(rw http.ResponseWriter, r *http.Request) {
	currentDate := time.Now().Format(time.RFC3339)
    rw.Write([]byte(currentDate))
}

func main() {
  http.HandleFunc("/", handleRequest)
  http.ListenAndServe(":8000", nil)
}