package main

import (
	"encoding/json"
	"log"
	"net/http"
	"restuarent-manager-go/controllers"
	"restuarent-manager-go/middleware"
	"restuarent-manager-go/models"
	"restuarent-manager-go/routes"
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

    
    insert()
    mux := http.NewServeMux()
	mux.HandleFunc("/orders", routes.GetOrdersHandler)

	// Apply the CORS middleware
	handlerWithCORS := middleware.CORS(mux)

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handlerWithCORS); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func insert(){
    newOrder := models.Order{
		Customer:    "John Doe",
		Amount:      99.99,
		OrderStatus: "Pending",
	}
	err := controllers.InsertOrder(newOrder)
    log.Println(err)
}