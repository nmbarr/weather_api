package main

import (
	"fmt"
)

// This function validates the date(s) passed to determine how to structure the URL
// https://www.visualcrossing.com/resources/documentation/weather-api/timeline-weather-api/#brxe-owrqlf
func handleDateParams(startDate, endDate string) (string, error) {
	
	// If no startDate or endDate is supplied, return an empty string (This returns weather data for the next 15 days)
	if startDate == "" && endDate == "" {
		fmt.Println("No startDate or endDate supplied. Returning weather data for the next 15 days.")
		return "", nil
	}
	
	// If no endDate is supplied, return startDate (This returns just weather data for a single day)
	if endDate == "" {
		fmt.Printf("No endDate supplied. Returning weather data for %s.\n", startDate)
		return startDate, nil
	}

	// If an endDate is supplied but no startDate, return an err
	if endDate != "" && startDate == "" {
		return "", fmt.Errorf("No startDate supplied with the given endDate %s. A startDate must be supplied with an endDate.", endDate)
	}

	// TODO: Handle dynamic date periods

	finalDateParam := fmt.Sprintf("%s/%s", startDate, endDate)
	fmt.Printf("Both startDate and endDate supplied. Passing param: %s\n", finalDateParam)

	return finalDateParam, nil
}

// The location must be an address, partial address, latitude/longitude, or zip code
// Can also define multiple locations, so need to be able to handle that
func handleLocationParams(location string) (error) {return nil}

// If we pass multiple locations, we have to URL encode the list of locations 
func encodeMultipleLocationParams(locations string) (string, error) {return "", nil}

// Build the Weather API URL with all validated params
func buildURL(baseURL string, locationParams []string, startDate, endDate, apiKey string) (string, error) {

	for _, location := range locationParams {
		locationErr := handleLocationParams(location)

		if locationErr != nil {
			return "", fmt.Errorf("Error handling location params: %w", locationErr)
		}
	}

	dateParam, dateErr := handleDateParams(startDate, endDate)
	
	if dateErr != nil {
		return "", fmt.Errorf("Error handling date params: %w", dateErr)
	}
	
	// http.Get expects the url already built, so build it here
	// TODO: Need to error handle if the url is bad
	// TODO: If multiple locations were passed, we need to structure the URL different so need to handle that
	url := fmt.Sprintf("%s/%s/%s?key=%s", baseURL, locationParams, dateParam, apiKey)
	fmt.Println(url)
	return url, nil
}

// Maybe need this. maybe not
func queryAPI(URL string) (string, error) {return "", nil}