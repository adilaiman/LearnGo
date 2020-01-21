package main

import (
	"fmt"
	"net/http"
)

// create funcs for each page
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {
	// set up routes
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	// start
	fmt.Println("Server starting on port 3000")
	http.ListenAndServe(":3000", nil)
}
