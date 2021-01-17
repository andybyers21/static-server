package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
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

	// Return 404 if the requested template does not exist.
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return 404 if the requested template path is a directory.
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
