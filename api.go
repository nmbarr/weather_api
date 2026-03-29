package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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
	fmt.Println("Both startDate and endDate supplied. Passing param: \n", finalDateParam)

	return finalDateParam, nil
}

// The location must be an address, partial address, latitude/longitude, or zip code
// Can also define multiple locations, so need to be able to handle that
func handleLocationParams(locationParams []string) (string, error) {

	// If only one location was passed, we still need to return it as a string
	if len(locationParams) == 1 {
		return locationParams[0], nil
	}

	return "", nil
}

// If we pass multiple locations, we have to URL encode the list of locations 
func encodeMultipleLocationParams(locations string) (string, error) {return "", nil}

// Build the Weather API URL with all validated params
func buildURL(baseURL string, locationParam string, dateParam string, apiKey string) (string, error) {

	
	// http.Get expects the url already built, so build it here
	// TODO: Need to error handle if the url is bad
	// TODO: If multiple locations were passed, we need to structure the URL different so need to handle that
	url := fmt.Sprintf("%s/%s/%s?key=%s", baseURL, locationParam, dateParam, apiKey)
	fmt.Println(url)
	return url, nil
}

// Send the request to the API and handle potential errors
func handleResponse(url string) ([]byte, error) {
	
	// Create client with timeout to prevent hanging
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error making request: %w", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status: ", resp.Status)

	// Handle non-success status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		switch resp.StatusCode {
		case http.StatusBadRequest:
			// 400 BAD_REQUEST
			return nil, fmt.Errorf("HTTP %d: The format of the API is incorrect or an invalid parameter or combination of parameters was supplied", resp.StatusCode)
		case http.StatusUnauthorized:
			// 401 UNAUTHORIZED
			return nil, fmt.Errorf("HTTP %d: There is a problem with the API key, account or subscription. May also be returned if a feature is requested for which the account does not have access to", resp.StatusCode)
		case http.StatusNotFound:
			// 404 NOT_FOUND
			return nil, fmt.Errorf("HTTP %d: The request cannot be matched to any valid API request endpoint structure", resp.StatusCode)
		case http.StatusTooManyRequests:
			// 429 TOO_MANY_REQUESTS
			return nil, fmt.Errorf("HTTP %d: The account has exceeded their assigned limits", resp.StatusCode)
		case http.StatusInternalServerError:
			// 500 INTERNAL_SERVER_ERROR
			return nil, fmt.Errorf("HTTP %d: A general error has occurred processing the request", resp.StatusCode)
		default:
			// Catch-all for other error codes
			return nil, fmt.Errorf("HTTP %d: Unexpected error status", resp.StatusCode)
		}
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %w", err)
	}

	// if http.Status returns 200 and the response body can be read, return body and no error
	return body, nil
}

// Format JSON with indentation
func formatResponse(body []byte) (bytes.Buffer, error) {

	var formattedJSON bytes.Buffer
	if err := json.Indent(&formattedJSON, body, "", "  "); err != nil {
		return formattedJSON, fmt.Errorf("Error formatting JSON: %w", err)
	}

	return formattedJSON, nil
}
