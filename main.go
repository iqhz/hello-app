package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Golang App on Kubernetes Cluster RKE2 - Keren")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}