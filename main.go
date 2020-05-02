package main

import (
  "net/http"
  "fmt"
)

func main() {
http.HandleFunc("/", hello)

http.ListenAndServe(":",nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "yuppy")
}
