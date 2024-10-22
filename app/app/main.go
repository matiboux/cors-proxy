package main

import (
    "os"
    "io"
    "log"
    "net/http"
    "net/url"
    "strings"
)

func overrideCORSHeaders(w http.ResponseWriter) {
    // Allow CORS from all origins, methods, and headers
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")
}

func handleProxy(w http.ResponseWriter, req *http.Request) {
    // Extract the first raw path argument as the target URL
    // Format: ^()/(targetURL)/(.*)$
    pathParts := strings.SplitN(req.URL.RawPath, "/", 3)

    if len(pathParts) < 2 || pathParts[1] == "" {
        // Target URL not specified
        // Redirect to https://example.com
        http.Redirect(w, req, "https://example.com", http.StatusFound)
        return
    }

    // Decode the target URL
    targetURL, err := url.QueryUnescape(pathParts[1])
    if err != nil {
        http.Error(w, "Invalid target URL", http.StatusBadRequest)
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
        http.Error(w, "Invalid target URL", http.StatusBadRequest)
        return
    }

    // Handle CORS preflight requests
    if req.Method == http.MethodOptions {
        overrideCORSHeaders(w)
        w.WriteHeader(http.StatusOK)
        return
    }

    // Create target request based on incoming request
    targetReq, err := http.NewRequest(req.Method, parsedURL.String(), req.Body)
    if err != nil {
        http.Error(w, "Failed to create request", http.StatusInternalServerError)
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
        return
    }
    defer resp.Body.Close()

    // Copy response headers and status code
    for key, values := range resp.Header {
        for _, value := range values {
            w.Header().Add(key, value)
        }
    }

    // Override CORS headers
    overrideCORSHeaders(w)

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
