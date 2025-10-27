package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mladenadamovic/numbeo-api-go/api"
	"github.com/mladenadamovic/numbeo-api-go/handlers"
)

func main() {
	// Get API key from environment variable
	apiKey := os.Getenv("NUMBEO_API_KEY")
	if apiKey == "" {
		log.Fatal("NUMBEO_API_KEY environment variable is required")
	}

	// Get port from environment variable, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create API client
	apiClient := api.NewClient(apiKey)

	// Create prices handler
	pricesHandler, err := handlers.NewPricesHandler(apiClient)
	if err != nil {
		log.Fatalf("Failed to create prices handler: %v", err)
	}

	// Setup routes
	http.Handle("/", pricesHandler)

	// Start server
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Starting server on http://localhost%s", addr)
	log.Printf("Visit http://localhost%s to view the application", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
