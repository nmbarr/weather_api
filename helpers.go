package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

// Write formatted JSON response to output file
func writeToFile(outputDir string, formattedJSON bytes.Buffer) error {

	// Ensure output directory exists before creating file
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("Error creating output directory: %w", err)
	}

	// Write to file with datetime
	fileName := fmt.Sprintf("%s/weather_response_%s.json", outputDir, time.Now().Format("2006-01-02_15-04-05"))
	if err := os.WriteFile(fileName, formattedJSON.Bytes(), 0644); err != nil {
		return fmt.Errorf("Error writing to file: %w", err)
	}

	fmt.Printf("Weather data written to: %s\n", fileName)

	return nil
}