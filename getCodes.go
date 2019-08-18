package main

import (
  "fmt"
  "net/http"
  "database/sql"
  "encoding/json"
  "io/ioutil"
  _ "github.com/lib/pq"
)

func getCodes(connStr string) {
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()

  resp, err := http.Get("http://country.io/phone.json")
  if err != nil {
    fmt.Println(err)
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  var r interface{}

  json.Unmarshal(body, &r)
  data := r.(map[string]interface{})
  for k, v := range data {
    db.Query("INSERT INTO numbers (ct, number) VALUES ($1, $2) ON CONFLICT DO NOTHING/UPDATE", k, v)
  }
}
