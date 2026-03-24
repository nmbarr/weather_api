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

// Load environment variables from .env
func getEnvironmentVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

// This function validates the date(s) passed to determine how to structure the URL
func checkDateParams(date1 string, date2 string) {}

func main() {

	BASE_URL := getEnvironmentVariable("VISUAL_CROSSING_WEATHER_API_URL")
	API_KEY := getEnvironmentVariable("VISUAL_CROSSING_WEATHER_API_KEY")

	// //////////////////////////////////////////////////////////////////////////////
	// TODO: Make the location, date1, and date2 parameters configureable through CLI
	// //////////////////////////////////////////////////////////////////////////////

	// The location must be an address, partial address, latitude/longitude, or zip code
	location := "93436"

	// These are optional date parameters
	date1 := ""
	date2 := ""

	// http.Get expects the url already built, so build it here
	url := fmt.Sprintf("%s/%s/%s/%s?key=%s", BASE_URL, location, date1, date2, API_KEY)
	fmt.Println(url)

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