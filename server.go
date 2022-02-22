package main

import (
	"net/http"
)

func handleRequest(rw http.ResponseWriter, r *http.Request) {
    rw.Write([]byte("Hello!"))
}

func main() {
  http.HandleFunc("/", handleRequest)
  http.ListenAndServe(":8000", nil)
}