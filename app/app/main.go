package main

import (
    "io"
    "log"
    "net/http"
    "net/url"
    "strings"
)

func handleProxy(w http.ResponseWriter, r *http.Request) {
    // Extract the first path argument as the target URL
    // Format: ^()/(targetURL)/(.*)$
    pathParts := strings.SplitN(r.URL.Path, "/", 3)

    if len(pathParts) < 2 || pathParts[1] == "" {
        http.Error(w, "Target URL not specified", http.StatusNotFound)
        return
    }

    // URL decode the target URL
    targetURL, err := url.QueryUnescape(pathParts[1])
    if err != nil {
        http.Error(w, "Invalid target URL", http.StatusBadRequest)
        return
	}

    // Parse the target URL
    parsedURL, err := url.Parse(targetURL)
    if err != nil {
        http.Error(w, "Invalid target URL", http.StatusBadRequest)
        return
    }

    // Create a new request based on the incoming request
    proxyReq, err := http.NewRequest(r.Method, parsedURL.String(), r.Body)
    if err != nil {
        http.Error(w, "Failed to create request", http.StatusInternalServerError)
        return
    }

    // Copy headers from the original request
    for key, values := range r.Header {
        for _, value := range values {
            proxyReq.Header.Add(key, value)
        }
    }

    // Perform the request
    client := &http.Client{}
    resp, err := client.Do(proxyReq)
    if err != nil {
        http.Error(w, "Failed to perform request", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // Copy response headers and status code
    for key, values := range resp.Header {
        for _, value := range values {
            w.Header().Add(key, value)
        }
    }
    w.WriteHeader(resp.StatusCode)

    // Copy response body
    io.Copy(w, resp.Body)
}

func main() {
    http.HandleFunc("/", handleProxy)
    log.Println("Starting proxy server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
