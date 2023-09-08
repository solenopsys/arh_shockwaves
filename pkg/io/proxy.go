package io

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func StartProxy(staticDir string, port string) {

	mux := http.NewServeMux()

	// Create a file server handler
	fileServer := http.FileServer(http.Dir(staticDir))

	targetURL, _ := url.Parse("http://solenopsys.org/")

	// Create a reverse proxy with custom Director function
	reverseProxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = targetURL.Scheme
			req.URL.Host = targetURL.Host
			req.URL.Path = "/dag"
			req.Host = targetURL.Host
		},
	}

	mux.Handle("/dag", reverseProxy)

	// Serve static files for any path
	mux.Handle("/", fileServer)

	// Start the server

	log.Fatal(http.ListenAndServe(":"+port, mux))
}

// LogTransport is a custom transport that logs the HTTP requests and responses
type LogTransport struct {
	Transport http.RoundTripper
}

// RoundTrip executes the HTTP request and logs the details
func (lt *LogTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Log the request details
	requestDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Println("Error logging request:", err)
	}
	log.Println("Request:")
	log.Println(string(requestDump))

	// Execute the request using the underlying transport
	response, err := lt.Transport.RoundTrip(req)
	if err != nil {
		log.Println("Error executing request:", err)
		return nil, err
	}

	// Log the response details
	responseDump, err := httputil.DumpResponse(response, true)
	if err != nil {
		log.Println("Error logging response:", err)
	}
	log.Println("Response:")
	log.Println(string(responseDump))

	return response, nil
}