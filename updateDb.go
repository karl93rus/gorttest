package main

import (
  "fmt"
  //"database/sql"
  "net/http"
  _ "github.com/lib/pq"
)

func updateDb(w http.ResponseWriter, req *http.Request) {
  connStr := "user=karl dbname=karl host=localhost sslmode=disable"

  fmt.Fprintf(w, "Updating...........")

  getCountries(connStr)
  getCodes(connStr)

  fmt.Fprintf(w, "Updated")

  // db, err := sql.Open("postgres", connStr)
  // if err != nil {
  //   fmt.Println(err)
  // }
  // defer db.Close()

  // var (
  //   ct string
  //   countries string
  //   number string
  // )

  // rows, err := db.Query("SELECT ct, country FROM countries")
  // for rows.Next() {
  //   rows.Scan(&ct, &countries)
  //   fmt.Println(ct, countries)
  // }

  // fmt.Println("\n\n")

  // rowss, err := db.Query("SELECT ct, number FROM numbers")
  // for rowss.Next() {
  //   rowss.Scan(&ct, &number)
  //   fmt.Println(ct, number)
  // }
}
