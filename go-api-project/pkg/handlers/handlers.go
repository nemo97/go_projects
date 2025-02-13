package handlers

import (
    "net/http"
    "encoding/json"
)

type Response struct {
    Message string `json:"message"`
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    response := Response{Message: "GET request successful"}
    json.NewEncoder(w).Encode(response)
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    response := Response{Message: "POST request successful"}
    json.NewEncoder(w).Encode(response)
}