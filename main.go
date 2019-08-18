package main

import (
  "fmt"
  "net/http"
)

func main() {
  fmt.Println("Start")

  http.HandleFunc("/code/", getCodeByName)

  http.ListenAndServe(":5555", nil)
}
