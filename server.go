package main

import (
	"net/http"
)

func main() {
  http.HandleFunc("/", nil)
  http.ListenAndServe(":8000", nil)
}