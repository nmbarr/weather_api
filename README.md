# Weather API Client

A Go-based client for querying the [Visual Crossing Weather API](https://www.visualcrossing.com/weather-api). This application fetches weather data for specified locations and date ranges, then saves the formatted JSON responses to timestamped files.

## Features

- **Flexible Date Queries**: Support for single-day, date range, or 15-day forecast queries
- **Location Support**: Query weather by address, zip code, or latitude/longitude
- **Error Handling**: Comprehensive HTTP status code handling with descriptive error messages
- **JSON Formatting**: Pretty-printed JSON output with proper indentation
- **Timestamped Output**: Automatically saves responses to dated files in the `output/` directory
- **Request Timeout**: Built-in 10-second timeout to prevent hanging requests
- **Environment Configuration**: Secure API key management via `.env` file

## Prerequisites

- Go 1.25.4 or later
- Visual Crossing Weather API account and API key ([Sign up here](https://www.visualcrossing.com/weather-api))

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd weather_api
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file from the example:
```bash
cp .env.example .env
```

4. Add your Visual Crossing API credentials to `.env`:
```
VISUAL_CROSSING_WEATHER_API_URL=https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline
VISUAL_CROSSING_WEATHER_API_KEY=your_api_key_here
```

## Usage

Currently, location and date parameters are configured in `main.go` (lines 36-43):

```go
// Modify these values before running
location := []string{"93436"}       // Zip code, address, or lat/long
startDate := "2025-03-25"           // yyyy-MM-dd format (optional)
endDate := "2025-04-20"             // yyyy-MM-dd format (optional)
```

Run the application:
```bash
go run .
```

The weather data will be saved to `output/weather_response_YYYY-MM-DD_HH-MM-SS.json`

### Date Parameter Behavior

- **No dates**: Returns 15-day forecast
- **Start date only**: Returns weather for that single day
- **Start + End dates**: Returns weather for the date range
- **End date only**: Returns an error (start date required)

## Project Structure

```
weather_api/
├── main.go           # Entry point, orchestrates the API request flow
├── api.go            # API URL building, parameter validation, HTTP handling
├── helpers.go        # File I/O operations
├── go.mod            # Go module dependencies
├── .env              # Environment variables (not committed)
├── .env.example      # Environment variable template
├── .gitignore        # Git ignore rules
└── output/           # Generated weather response files
```

## Key Functions

### `api.go`
- `handleDateParams()`: Validates and formats date parameters for the API
- `handleLocationParams()`: Validates location parameters (stub)
- `buildURL()`: Constructs the complete API request URL
- `handleResponse()`: Makes HTTP request and handles errors
- `formatResponse()`: Pretty-prints JSON response

### `helpers.go`
- `writeToFile()`: Writes formatted JSON to timestamped output file

## Error Handling

The application handles the following HTTP error codes:
- **400**: Invalid parameters or malformed request
- **401**: API key issues or subscription problems
- **404**: Invalid API endpoint
- **429**: Rate limit exceeded
- **500**: Internal server error

## Future Enhancements

- [ ] CLI flags for location and date parameters
- [ ] Redis caching layer for API responses
- [ ] Support for multiple location queries
- [ ] Dynamic date period handling
- [ ] URL encoding for location lists
- [ ] Configuration validation and improved error messages

## API Documentation

For detailed API documentation, see the [Visual Crossing Weather API Timeline Documentation](https://www.visualcrossing.com/resources/documentation/weather-api/timeline-weather-api/)

## Dependencies

- [godotenv](https://github.com/joho/godotenv) - Environment variable management

## License

[Add your license here]
