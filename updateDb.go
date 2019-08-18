package main

import (
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)

func updateDb() {
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
    rowss.Scan(&ct, &number)
    fmt.Println(ct, number)
  }
}
