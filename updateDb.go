package main

import (
  "fmt"
  "net/http"
  _ "github.com/lib/pq"
  "os"
  "time"
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

  fmt.Fprintf(w, "Updating...........\n")

  getCountries(connStr)
  getCodes(connStr)

  fmt.Fprintf(w, "Updated")

  var logStr string
  logStr = "UPDATE DATABASE " + req.Method + " from " + req.RemoteAddr + " on " + string(req.URL.Path) + " at " + time.Now().Format("2006-01-02 15:04:05")
  file, err := os.OpenFile("mainlog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    fmt.Println("Error opening file...")
  }
  defer file.Close()
  file.Write([]byte(logStr + "\n"))
}
