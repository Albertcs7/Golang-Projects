package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HEllo World")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ABout page")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/about", about)
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
