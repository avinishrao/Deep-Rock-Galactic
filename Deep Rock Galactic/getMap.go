package main

import (
	"net/http"
)

func main() {
	// Define the directory to serve files from (in this case, the "static" directory)
	fs := http.FileServer(http.Dir("static"))

	// Create a new HTTP server and route requests to the file server
	http.Handle("/", fs)

	// Start the server on a specified port
	port := ":8080"
	println("File server is running on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
