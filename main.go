package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
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

	BASE_URL := os.Getenv("VISUAL_CROSSING_WEATHER_API_URL")
	API_KEY := os.Getenv("VISUAL_CROSSING_WEATHER_API_KEY")

	if BASE_URL == "" || API_KEY == "" {
		log.Fatal("Missing required environment variables")
	}

	// //////////////////////////////////////////////////////////////////////////////
	// TODO: Make the location, date1, and date2 parameters configureable through CLI
	// //////////////////////////////////////////////////////////////////////////////

	// The location must be an address, partial address, latitude/longitude, or zip code
	location := "93436"

	// These are optional date parameters
	date1 := ""
	date2 := ""

	// Build the URL that will be comsumed by http.Get
	url, err := buildURL(
		BASE_URL,
		location,
		date1,
		date2,
		API_KEY,
	)

    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response status:", resp.Status)

    scanner := bufio.NewScanner(resp.Body)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}