package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	// Create a handler which responds to all HTTP requests with the contents of a given file system (static).
	fs := http.FileServer(http.Dir("./static"))
	// Register file server for all requests.
	// Strip off the "/static/" prefix from the request path before searching the file system.
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// All requests not picked up by the static file server should be handled by serveTemplates() function.
	http.HandleFunc("/", serveTemplates)

	log.Println("Listening on port 9100...")
	err := http.ListenAndServe(":9100", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveTemplates(w http.ResponseWriter, r *http.Request) {
	// Build paths to the layout file and the corresponding template file request.
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	// Bundle the requested template and layout into a template set and render a named template in the set.
	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}
