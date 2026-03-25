package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

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
	location := []string{"48864"}

	// These are optional date parameters
	// yyyy-MM-dd or yyyy-MM-ddTHH:mm:ss format
	startDate := "2025-03-25"
	// startDate := ""
	endDate := "2025-04-20"
	// endDate := ""

	// Build the URL that will be comsumed by http.Get
	url, err := buildURL(baseURL, location, startDate, endDate, apiKey)
	if err != nil {
		log.Fatal(err)
	}

    resp, err := http.Get(url)
    if err != nil {
        log.Fatal("Error making request: ", err)
    }
    defer resp.Body.Close()

    fmt.Println("Response status:", resp.Status)

    // Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	// Unmarshal into an interface
	var jsonData interface{}
	if err := json.Unmarshal(body, &jsonData); err != nil {
		log.Fatal("Error unmarshaling JSON: ", err)
	}

	// Marshal with indentation
	prettyJSON, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		log.Fatal("Error marshaling JSON: ", err)
	}

	// Write to file with datetime
	// TODO: Need to check id the output dir exists
	fileName := fmt.Sprintf("output/weather_response_%s.json", time.Now().Format("2006-01-02_15-04-05"))
	if err := os.WriteFile(fileName, prettyJSON, 0644); err != nil {
		log.Fatal("Error writing to file: ", err)
	}

	fmt.Printf("Weather data written to: %s\n", fileName)
	// fmt.Println(string(prettyJSON))
}