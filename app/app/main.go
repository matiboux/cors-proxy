package main

import (
	"os"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func overrideCORSHeaders(w http.ResponseWriter, req *http.Request) {
	// Set allowed origins for CORS
	origin := req.Header.Get("Access-Control-Allow-Origin")
	if origin == "" {
		origin = "*" // Default to allow all origins
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Set allowed methods for CORS
	methods := req.Header.Get("Access-Control-Allow-Methods")
	if methods == "" {
		methods = "*" // Default to allow all methods
	}
	w.Header().Set("Access-Control-Allow-Methods", "*")

	// Set allowed headers for CORS
	headers := req.Header.Get("Access-Control-Allow-Headers")
	if headers == "" {
		headers = "*" // Default to allow all headers
	}
	w.Header().Set("Access-Control-Allow-Headers", "*")

	// Set allowed credentials for CORS
	credentials := req.Header.Get("Access-Control-Allow-Credentials")
	if credentials != "" {
		w.Header().Set("Access-Control-Allow-Credentials", credentials)
	}

	// Set max age for CORS
	maxAge := req.Header.Get("Access-Control-Max-Age")
	if maxAge != "" {
		w.Header().Set("Access-Control-Max-Age", maxAge)
	}

	// Set exposed headers for CORS
	jsHeaders := req.Header.Get("Access-Control-Expose-Headers")
	if jsHeaders != "" {
		w.Header().Set("Access-Control-Expose-Headers", jsHeaders)
	}
}

func handleProxy(w http.ResponseWriter, req *http.Request) {
	// Extract the first raw path argument as the target URL
	// Format: ^()/(targetURL)/(.*)$
	pathParts := strings.SplitN(req.URL.EscapedPath(), "/", 3)

	if len(pathParts) < 2 || pathParts[1] == "" {
		// Target URL not specified
		// Redirect to https://example.com
		http.Redirect(w, req, "https://example.com", http.StatusFound)
		log.Println("Redirect to https://example.com")
		return
	}

	// Decode the target URL
	targetURL, err := url.QueryUnescape(pathParts[1])
	if err != nil {
		http.Error(w, "Invalid target URL", http.StatusBadRequest)
		log.Printf("Invalid target URL: %v\n", err)
		return
	}

	// Remove trailing slash from target URL
	if strings.HasSuffix(targetURL, "/") {
		targetURL = targetURL[:len(targetURL)-1]
	}

	// Set the target path
	targetPath := "/"
	if len(pathParts) > 2 {
		targetPath += pathParts[2]
	}

	// Parse the target URL
	parsedURL, err := url.Parse(targetURL + targetPath)
	if err != nil {
		http.Error(w, "Invalid target URL", http.StatusInternalServerError)
		log.Printf("Invalid target URL: %v\n", err)
		return
	}

	// Handle CORS preflight requests
	if req.Method == http.MethodOptions {
		overrideCORSHeaders(w, req)
		w.WriteHeader(http.StatusOK)
		return
	}

	// Create target request based on incoming request
	targetReq, err := http.NewRequest(req.Method, parsedURL.String(), req.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		log.Printf("Failed to create request: %v\n", err)
		return
	}

	// Copy incoming headers into target request
	for key, values := range req.Header {
		for _, value := range values {
			targetReq.Header.Add(key, value)
		}
	}

	// Perform the target request
	client := &http.Client{}
	resp, err := client.Do(targetReq)
	if err != nil {
		http.Error(w, "Failed to perform request", http.StatusInternalServerError)
		log.Printf("Failed to perform request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Log the proxied request
	log.Printf("Proxied request to %s (%d)\n", parsedURL.String(), resp.StatusCode)

	// Copy response headers and status code
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Override CORS headers
	overrideCORSHeaders(w, req)

	w.WriteHeader(resp.StatusCode)

	// Copy response body
	io.Copy(w, resp.Body)
}

func main() {
	// Read the PORT environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080
	}

	http.HandleFunc("/", handleProxy)
	log.Println("Starting proxy server on :" + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
