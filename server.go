package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func handleRequest(rw http.ResponseWriter, r *http.Request) {

  rw.Header().Set("Content-Type", "application/json")
  currentDate := time.Now().Format(time.RFC3339)
  
  resp := make(map[string]string)
  
  resp["time"] = currentDate

  err := json.NewEncoder(rw).Encode(resp)
  if err != nil {
    log.Fatalf("Error happened in while encoding. Err: %s", err)
  }
}


func main() {



  http.HandleFunc("/time", handleRequest) 
  http.ListenAndServe(":8795", nil)
}