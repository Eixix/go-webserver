package main

import (
  "fmt"
  "net/http"
)

func handleAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "hello </br>")
}

func main() {
  fmt.Printf("Starting Server\n")
	http.HandleFunc("/", handleAll)
	http.ListenAndServe(":8080", nil)
}
