package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)




func main() {



  http.HandleFunc("/time", func (rw http.ResponseWriter, r *http.Request) {

    rw.Header().Set("Content-Type", "application/json")
    currentDate := time.Now().Format(time.RFC3339)
    
    resp := make(map[string]string)
    resp["time"] = currentDate
    err := json.NewEncoder(rw).Encode(resp)
    if err != nil {
      log.Fatalf("Error happened in JSON marshal. Err: %s", err)
    }
  })
  http.ListenAndServe(":8795", nil)
}