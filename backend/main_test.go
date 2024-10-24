package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetItems(t *testing.T) {
    // Create a request to pass to our handler
    req, err := http.NewRequest("GET", "/items", nil)
    if err != nil {
        t.Fatalf("Could not create request: %v", err)
    }

    // Create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()

    // Call the handler with the ResponseRecorder and the request
    handler := http.HandlerFunc(getItems)
    handler.ServeHTTP(rr, req)

    // Check if the status code is 200 OK
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check if the Content-Type is application/json
    if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
        t.Errorf("Handler returned wrong Content-Type: got %v want %v", contentType, "application/json")
    }

    // Check if the response body contains the expected JSON
    expected, _ := json.Marshal(items)
    if strings.TrimSpace(rr.Body.String()) != string(expected) {
        t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
    }
}
