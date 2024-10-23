package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Define a struct to represent the data model (e.g., Item)
type Item struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
}

// Create a slice to act as an in-memory database
var items = []Item{}
var nextID = 1
var mu sync.Mutex // Mutex for thread-safe access

// Get all items (GET /items)
func getItems(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

// Get a specific item by ID (GET /items/{id})
func getItem(w http.ResponseWriter, r *http.Request) {
	idString, err := strconv.Atoi(r.PathValue("id"))
    w.Header().Set("Content-Type", "application/json")

    if err != nil {
        http.Error(w, "Invalid item ID", http.StatusBadRequest)
        return
    }
    for _, item := range items {
        if item.ID == idString {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    http.Error(w, "Item not found", http.StatusNotFound)
}

// Create a new item (POST /items)
func createItem(w http.ResponseWriter, r *http.Request) {
    var newItem Item
    err := json.NewDecoder(r.Body).Decode(&newItem)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    mu.Lock()
    newItem.ID = nextID
    nextID++
    items = append(items, newItem)
    mu.Unlock()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newItem)
}

// Update an item by ID (PUT /items/{id})
func updateItem(w http.ResponseWriter, r *http.Request) {
	
    id, err := strconv.Atoi(r.URL.Path[len("/items/"):])
    if err != nil {
        http.Error(w, "Invalid item ID", http.StatusBadRequest)
        return
    }

    var updatedItem Item
    err = json.NewDecoder(r.Body).Decode(&updatedItem)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    mu.Lock()
    defer mu.Unlock()

    for i, item := range items {
        if item.ID == id {
            updatedItem.ID = id
            items[i] = updatedItem
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(updatedItem)
            return
        }
    }
    http.Error(w, "Item not found", http.StatusNotFound)
}

// Delete an item by ID (DELETE /items/{id})
func deleteItem(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Path[len("/items/"):])
    if err != nil {
        http.Error(w, "Invalid item ID", http.StatusBadRequest)
        return
    }

    mu.Lock()
    defer mu.Unlock()

    for i, item := range items {
        if item.ID == id {
            items = append(items[:i], items[i+1:]...) // Remove item from the slice
            w.WriteHeader(http.StatusNoContent)      // Return 204 No Content
            return
        }
    }
    http.Error(w, "Item not found", http.StatusNotFound)
}

func main() {
    // Define the routes for the API
    http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            getItems(w, r)
        } else if r.Method == http.MethodPost {
            createItem(w, r)
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    http.HandleFunc("/items/{id}", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            getItem(w, r)
        } else if r.Method == http.MethodPut {
            updateItem(w, r)
        } else if r.Method == http.MethodDelete {
            deleteItem(w, r)
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    fmt.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
