package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Hello from Go!")
}

func main() {
  http.HandleFunc("/", handler )

  fmt.Println("Server started at localhost:8080")
  err := http.ListenAndServe(":8080", nil)

  if err != nil {
    panic(err)
  }
}


