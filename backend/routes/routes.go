package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"restuarent-manager-go/controllers"
)

// GetOrdersHandler handles the /orders route to fetch all orders
func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch orders from the controller
	log.Println("Received request to /orders")
	orders, err := controllers.GetOrders()
	if err != nil {
		http.Error(w, "Failed to fetch orders", http.StatusInternalServerError)
		return
	}

	// Convert orders to JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orders); err != nil {
		log.Println("Failed to encode orders:", err)
		http.Error(w, "Failed to encode orders", http.StatusInternalServerError)
	}
}

// NotFoundHandler handles 404 not found responses
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 not found", http.StatusNotFound)
}
