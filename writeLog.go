package main

import (
  "fmt"
  "os"
)

func writeLog(log string) {
  file, err := os.OpenFile("mainlog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    fmt.Println("Error opening file...")
  }
  defer file.Close()
  file.Write([]byte(log + "\n"))
}
