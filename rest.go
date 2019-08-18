package main

import (
  "fmt"
  "net/http"
  "database/sql"
)

func getCodeByName(w http.ResponseWriter, req *http.Request) {
  // country := req.URL.Path[len("/code/"):]

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

  // res, err := db.Query("SELECT country, number FROM numbers INNER JOIN countries ON countries.ct = numbers.ct AND countries.country = $1", country)
  res, err := db.Query("SELECT * FROM countries INNER JOIN numbers ON countries.ct = numbers.ct AND country = 'France';")
  if err != nil {
    fmt.Println(err)
  }
  defer res.Close()

  fmt.Println(res)

  for res.Next() {
    res.Scan(&r.c)
    fmt.Println(r.c)
  }

  fmt.Fprintf(w, "hello, %v", r.c)
}
