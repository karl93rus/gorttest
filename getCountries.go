package main

import (
  "fmt"
  "net/http"
  "database/sql"
  "encoding/json"
  "io/ioutil"
  _ "github.com/lib/pq"
)

func getCountries(connStr string) error {
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    fmt.Println(err)
    return err
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
  queryStr := "INSERT INTO countries (ct, country) VALUES ($1, $2) ON CONFLICT (ct) DO NOTHING"
  for k, v := range data {
    _, err := db.Exec(queryStr, k, v)
    if err != nil {
      fmt.Println(err)
      return err
    }
    // fmt.Println(k, v)
  }

  fmt.Println("Countries updated\n")

  return nil
}
