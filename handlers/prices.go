package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/mladenadamovic/numbeo-api-go/api"
)

// PricesHandler handles the prices web page
type PricesHandler struct {
	APIClient *api.Client
	Template  *template.Template
}

// NewPricesHandler creates a new PricesHandler
func NewPricesHandler(apiClient *api.Client) (*PricesHandler, error) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		return nil, err
	}

	return &PricesHandler{
		APIClient: apiClient,
		Template:  tmpl,
	}, nil
}

// PageData represents the data passed to the template
type PageData struct {
	City     string
	Country  string
	Response *api.CityPricesResponse
	Error    string
}

// ServeHTTP handles HTTP requests
func (h *PricesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		City:    r.URL.Query().Get("city"),
		Country: r.URL.Query().Get("country"),
	}

	// Set default values for initial load
	if data.City == "" {
		data.City = "San Francisco, CA"
	}
	if data.Country == "" {
		data.Country = "United States"
	}

	// If it's a form submission, fetch the data
	if r.Method == http.MethodGet && r.URL.Query().Get("fetch") == "true" {
		resp, err := h.APIClient.GetCityPrices(data.City, data.Country)
		if err != nil {
			log.Printf("Error fetching prices: %v", err)
			data.Error = err.Error()
		} else {
			data.Response = resp
		}
	}

	// Render the template
	if err := h.Template.Execute(w, data); err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
