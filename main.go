package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a handler which responds to all HTTP requests with the contents of a given file system (static).
	fs := http.FileServer(http.Dir("./static"))
	// Register file server for all requests (on port 9100 in this case).
	http.Handle("/", fs)

	log.Println("Listening on port 9100...")
	err := http.ListenAndServe(":9100", nil)
	if err != nil {
		log.Fatal(err)
	}
}
