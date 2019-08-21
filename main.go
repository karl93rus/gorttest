package main

import (
  "fmt"
  "net/http"
  "time"
)

func main() {
  port := ":5555"
  fmt.Println("Server started on port " + port)

  var logStr string
  logStr = "Server started on port " + port + " at " + time.Now().Format("2006-01-02 15:04:05") + "\n...\n...\n"
  writeLog(logStr)


  http.HandleFunc("/code/", getCodeByName)
  http.HandleFunc("/reload/", updateDb)

  http.ListenAndServe(port, nil)
}
