package main

import (
  "fmt"
  "net/http"
  "database/sql"
  "strings"
  "time"
)

func getCodeByName(w http.ResponseWriter, req *http.Request) {
  country := req.URL.Path[len("/code/"):]
  country = strings.ToLower(country)
  if req.Method != "GET" {
    fmt.Fprintf(w, "Invalid request method. Terminating...")
    return
  }
  fmt.Println(req.Method, time.Now())

  connStr := "user=karl dbname=karl host=localhost sslmode=disable"
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()

  type Result struct {
    c string
    n string
  }
  var r Result

  res, err := db.Query("SELECT countries.country, numbers.number FROM numbers INNER JOIN countries ON numbers.ct = countries.ct AND  LOWER(countries.country)=LOWER($1)", country)
  if err != nil {
    fmt.Println(err)
  }
  defer res.Close()

  for res.Next() {
    res.Scan(&r.c, &r.n)
    fmt.Println(r.c, r.n)
  }

  if r.c == "" {
    w.WriteHeader(404)
    fmt.Fprintf(w, "Country not found...")
    fmt.Println("Country not found...")
    return
  }

  w.WriteHeader(200)
  fmt.Fprintf(w, "hello blya, %v, %v", r.c, r.n)
}
