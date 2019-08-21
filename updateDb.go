package main

import (
  "fmt"
  "net/http"
  _ "github.com/lib/pq"
  "time"
)

func updateDb(w http.ResponseWriter, req *http.Request) {
  fmt.Println(req.Method, req.URL, req.Body)
  if req.Method != "POST" {
    fmt.Println("Invalid request method...")
    fmt.Fprintf(w, "Invalid request method...")
    return
  }
  connStr := "user=karl dbname=karl host=localhost sslmode=disable"

  fmt.Println("Updating...")
  fmt.Fprintf(w, "Updating...........\n")

  gerr := getCountries(connStr)
  if gerr != nil {
    fmt.Fprintf(w, "Error occured while connecting with database...")
    return
  }
  getCodes(connStr)

  fmt.Fprintf(w, "Updated")

  var logStr string
  logStr = "UPDATE DATABASE " + req.Method + " from " + req.RemoteAddr + " on " + string(req.URL.Path) + " at " + time.Now().Format("2006-01-02 15:04:05")
  writeLog(logStr)
}
