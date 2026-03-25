package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// //////////////////////////////////////////////////////////////////////////////
// 1. Check Redis cache
// 2. Redis cache response
// 3. Request Weather API
// 4. Weather API response
// 5. Save cached results to Redis
// //////////////////////////////////////////////////////////////////////////////

func main() {

	// Load .env file
	_ = godotenv.Load() 

	baseURL := os.Getenv("VISUAL_CROSSING_WEATHER_API_URL")
	apiKey := os.Getenv("VISUAL_CROSSING_WEATHER_API_KEY")

	if baseURL == "" || apiKey == "" {
		log.Fatal("Missing required environment variables")
	}

	// //////////////////////////////////////////////////////////////////////////////
	// TODO: Make the location, date1, and date2 parameters configureable through CLI
	// //////////////////////////////////////////////////////////////////////////////

	// The location must be an address, partial address, latitude/longitude, or zip code
	location := []string{"93436"}

	// These are optional date parameters
	// yyyy-MM-dd or yyyy-MM-ddTHH:mm:ss format
	startDate := "2025-03-25"
	// startDate := ""
	endDate := "2025-04-20"
	// endDate := ""

	// Build the URL that will be consumed by handleResponse
	url, err := buildURL(baseURL, location, startDate, endDate, apiKey)
	if err != nil {
		log.Fatal(err)
	}

	// Send request to the API and hande the response
	body, err := handleResponse(url)

	if err != nil {
		log.Fatal(err)
	}

	// Format the JSON response in preparation for writing to file
	formattedJSON, err := formatResponse(body)
	
	if err != nil {
		log.Fatal(err)
	}

	outputDir := "output"

	// Write the formatted JSON response to file
	fileErr := writeToFile(outputDir, formattedJSON)

	if fileErr != nil {
		log.Fatal(fileErr)
	}

	fmt.Println("Successfully requested weather data from API.")
}