package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write the HTML response
	fmt.Fprintf(w, "<html><body><h1>Hello, From test webapp !</h1></body></html>")
}

func main() {
	// Define a handler function for the root route ("/")
	http.HandleFunc("/", helloHandler)

	// Start the web server on port 8080
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
