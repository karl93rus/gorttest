package main

import (
  "fmt"
  "net/http"
  _ "github.com/lib/pq"
)

func updateDb(w http.ResponseWriter, req *http.Request) {
  fmt.Println(req.Method)
  fmt.Println(req.URL)
  fmt.Println(req.Body)
  if req.Method != "GET" {
    fmt.Fprintf(w, "Invalid request method...")
    return
  }
  connStr := "user=karl dbname=karl host=localhost sslmode=disable"

  fmt.Fprintf(w, "Updating...........")

  getCountries(connStr)
  getCodes(connStr)

  fmt.Fprintf(w, "Updated")
}
