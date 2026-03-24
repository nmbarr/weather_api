package main

import (
	"fmt"
)

// This function validates the date(s) passed to determine how to structure the URL
// https://www.visualcrossing.com/resources/documentation/weather-api/timeline-weather-api/#brxe-owrqlf
func checkDateParams(date1, date2 string) {}

// The location must be an address, partial address, latitude/longitude, or zip code
// Can also define multiple locations, so need to be able to handle that
func checkLocationParams(locations []string) {}

// Build the Weather API URL with all validated params
func buildURL(BASE_URL, location, date1, date2, API_KEY string) (string, error) {

	// TODO: Need to error handle if the url is bad

	// http.Get expects the url already built, so build it here
	url := fmt.Sprintf("%s/%s/%s/%s?key=%s", BASE_URL, location, date1, date2, API_KEY)
	return url, nil
}

// Maybe need this. maybe not
func queryAPI(URL string) (string, error) {return "", nil}