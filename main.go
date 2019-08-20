package main

import (
  "fmt"
  "net/http"
  "os"
  "time"
)

func main() {
  port := ":5555"
  fmt.Println("Server started on port " + port)

  var logStr string
  logStr = "Server started on port " + port + " at " + time.Now().Format("2006-01-02 15:04:05")
  file, err := os.OpenFile("mainlog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    fmt.Println("Error opening file...")
  }
  defer file.Close()
  file.Write([]byte(logStr + "\n"))


  http.HandleFunc("/code/", getCodeByName)
  http.HandleFunc("/reload/", updateDb)

  http.ListenAndServe(port, nil)
}
