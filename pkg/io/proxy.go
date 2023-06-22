package io

import (
	"log"
	"net/http"
)

func StartProxy(staticDir string) {

	// Create a file server handler
	fileServer := http.FileServer(http.Dir(staticDir))

	// Serve static files for any path
	http.Handle("/", fileServer)

	// Start the server
	log.Println("Static file server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
