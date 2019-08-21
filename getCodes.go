package main

import (
  "fmt"
  "net/http"
  "database/sql"
  "encoding/json"
  "io/ioutil"
  _ "github.com/lib/pq"
)

func getCodes(connStr string) error {
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()

  var r interface{}

  resp, err := http.Get("http://country.io/phone.json")
  if err != nil {
    fmt.Println(err)
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  json.Unmarshal(body, &r)
  data := r.(map[string]interface{})
  queryStr := "INSERT INTO numbers (ct, number) VALUES ($1, $2) ON CONFLICT (ct) DO NOTHING"
  for k, v := range data {
    _, err := db.Exec(queryStr, k, v)
    if err != nil {
      fmt.Println(err)
      return err
    }
    // fmt.Println(k, v)
  }

  fmt.Println("Codes updated\n")

  return nil
}
