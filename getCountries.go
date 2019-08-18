package main

import (
  "fmt"
  "net/http"
  "database/sql"
  "encoding/json"
  "io/ioutil"
  _ "github.com/lib/pq"
)

func getCountries(connStr string) {
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()

  var i interface{}

  resp, err := http.Get("http://country.io/names.json")
  if err != nil {
    fmt.Println(err)
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  json.Unmarshal(body, &i)
  data := i.(map[string]interface{})
  for k, v := range data {
    db.Query("INSERT INTO countries (ct, country) VALUES ($1, $2) ON CONFLICT DO NOTHING/UPDATE", k, v)
  }
}
