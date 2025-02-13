package utils

import (
    "log"
    "net/http"
)

// ValidateInput checks if the input is valid.
func ValidateInput(input string) bool {
    return input != ""
}

// LogRequest logs the details of an HTTP request.
func LogRequest(r *http.Request) {
    log.Printf("Received %s request for %s", r.Method, r.URL.Path)
}