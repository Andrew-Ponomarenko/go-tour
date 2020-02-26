package main

import (
  "log"
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/", fs)

  log.Println("responding...")
  http.ListenAndServe(":9999", nil)
}