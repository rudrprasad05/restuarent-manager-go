package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

// CORS is a middleware function that wraps the existing handler to enable CORS
func CORS(h http.Handler) http.Handler {
	// Create a new CORS handler with options
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow frontend origin
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true, // Allow credentials like cookies or authorization headers
	})
	
	// Return the wrapped handler with CORS enabled
	return c.Handler(h)
}
