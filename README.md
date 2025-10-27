# Numbeo API Go Application

A Go web application that fetches and displays cost of living data from the [Numbeo API](https://www.numbeo.com/common/api.jsp) for any city worldwide.

## Features

- Fetch real-time cost of living prices for any city
- Beautiful, responsive web interface
- Organized price data by categories (Restaurants, Markets, Transportation, etc.)
- Display average prices, price ranges, and data points
- Easy-to-use search form

## Prerequisites

- Go 1.16 or higher
- A Numbeo API key (get one from [Numbeo API](https://www.numbeo.com/common/api.jsp))

## Installation

1. Clone the repository:
```bash
git clone https://github.com/mladenadamovic/numbeo-api-go.git
cd numbeo-api-go
```

2. Install dependencies:
```bash
go mod download
```

3. Set up your environment variables:
```bash
cp .env.example .env
# Edit .env and add your Numbeo API key
```

## Configuration

Create a `.env` file in the root directory with the following variables:

```env
NUMBEO_API_KEY=your_api_key_here
PORT=8080
```

Or set them as environment variables:

```bash
export NUMBEO_API_KEY=your_api_key_here
export PORT=8080
```

## Usage

### Running the Application

```bash
go run main.go
```

The application will start on `http://localhost:8080` (or the port specified in your environment variables).

### Building the Application

```bash
go build -o numbeo-api-go
./numbeo-api-go
```

### Using the Application

1. Open your browser and navigate to `http://localhost:8080`
2. Enter a city name (e.g., "San Francisco, CA")
3. Enter a country name (e.g., "United States")
4. Click "Get Prices" to fetch the data
5. View the organized price data by category

## Project Structure

```
numbeo-api-go/
├── api/
│   └── numbeo.go          # Numbeo API client
├── handlers/
│   └── prices.go          # HTTP request handlers
├── templates/
│   └── index.html         # Web page template
├── main.go                # Application entry point
├── go.mod                 # Go module definition
├── .env.example           # Example environment variables
├── .gitignore            # Git ignore rules
└── README.md             # This file
```

## API Reference

The application uses the Numbeo API's `city_prices` endpoint:

```
GET https://www.numbeo.com/api/city_prices?city={city}&country={country}&api_key={api_key}
```

### Response Structure

```json
{
  "city_name": "San Francisco, CA",
  "country": "United States",
  "currency": "USD",
  "month_year": "October 2025",
  "prices": [
    {
      "item_id": 1,
      "item_name": "Meal, Inexpensive Restaurant",
      "category_name": "Restaurants",
      "data_points": 1500,
      "average_price": 25.00,
      "lowest_price": 15.00,
      "highest_price": 35.00
    }
  ]
}
```

## Example Queries

- **San Francisco, CA** / **United States**
- **London** / **United Kingdom**
- **Tokyo** / **Japan**
- **Berlin** / **Germany**
- **Sydney** / **Australia**

## Development

### Running Tests

```bash
go test ./...
```

### Code Structure

- **api/numbeo.go**: Contains the API client that makes HTTP requests to the Numbeo API
- **handlers/prices.go**: HTTP handlers that process requests and render templates
- **templates/index.html**: HTML template with CSS styling for the web interface
- **main.go**: Application entry point that sets up the server

## Error Handling

The application handles various error scenarios:

- Missing API key
- Invalid city or country names
- API request failures
- Network errors
- Invalid API responses

Errors are displayed to the user in a friendly format on the web page.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the MIT License.

## Acknowledgments

- Data provided by [Numbeo.com](https://www.numbeo.com)
- Built with Go standard library and minimal dependencies
