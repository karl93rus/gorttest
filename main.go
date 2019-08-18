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

func main() {
  connStr := "user=karl dbname=karl host=localhost sslmode=disable"

  getCountries(connStr)
  getCodes(connStr)
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()

  var (
    ct string
    cct string
    countries string
    number string
  )

  rows, err := db.Query("SELECT ct, country FROM countries")
  for rows.Next() {
    rows.Scan(&ct, &countries)
    fmt.Println(ct, countries)
  }

  fmt.Println("\n\n\n\n")

  rowss, err := db.Query("SELECT ct, number FROM numbers")
  for rowss.Next() {
    rowss.Scan(&cct, &number)
    fmt.Println(cct, number)
  }
}
