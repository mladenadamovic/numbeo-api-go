package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	BaseURL = "https://www.numbeo.com/api"
)

// Client represents a Numbeo API client
type Client struct {
	APIKey     string
	HTTPClient *http.Client
}

// NewClient creates a new Numbeo API client
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// PriceItem represents a single price item from the API
type PriceItem struct {
	ItemID        int     `json:"item_id"`
	ItemName      string  `json:"item_name"`
	CategoryName  string  `json:"category_name"`
	DataPoints    int     `json:"data_points"`
	AveragePrice  float64 `json:"average_price"`
	LowestPrice   float64 `json:"lowest_price"`
	HighestPrice  float64 `json:"highest_price"`
}

// CityPricesResponse represents the API response for city prices
type CityPricesResponse struct {
	CityName    string      `json:"city_name"`
	Country     string      `json:"country"`
	Currency    string      `json:"currency"`
	Prices      []PriceItem `json:"prices"`
	MonthYear   string      `json:"month_year"`
}

// GetCityPrices fetches prices for a given city and country
func (c *Client) GetCityPrices(city, country string) (*CityPricesResponse, error) {
	// Build the URL with query parameters
	endpoint := fmt.Sprintf("%s/city_prices", BaseURL)
	params := url.Values{}
	params.Add("city", city)
	params.Add("country", country)
	params.Add("api_key", c.APIKey)

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	// Make the HTTP request
	resp, err := c.HTTPClient.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	// Parse the response
	var result CityPricesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}
