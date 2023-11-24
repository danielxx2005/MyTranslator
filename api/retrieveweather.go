package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

const apiKey = "YOUR_API_KEY"
const apiEndpoint = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/"

type WeatherResponse struct {
	Days []struct {
		Date          string `json:"datetimeStr"`
		Temperature   float64
		Precipitation float64
		// Add more fields as needed
	} `json:"days"`
}

func main() {
	location := "New York, NY"
	startDate := time.Now().Format("2006-01-01")
	endDate := time.Now().Format("2006-01-03")

	client := resty.New()

	resp, err := client.R().
		SetQueryParam("key", apiKey).
		SetQueryParam("location", location).
		SetQueryParam("startDate", startDate).
		SetQueryParam("endDate", endDate).
		SetHeader("Content-Type", "application/json").
		Get(apiEndpoint)

	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode() != 200 {
		log.Fatalf("Error: %v", resp)
	}

	var weatherData WeatherResponse
	err = json.Unmarshal(resp.Body(), &weatherData)
	if err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
	}

	// Print weather data
	fmt.Printf("Weather for %s from %s to %s:\n", location, startDate, endDate)
	for _, day := range weatherData.Days {
		fmt.Printf("Date: %s, Temperature: %.2f°C, Precipitation: %.2fmm\n", day.Date, day.Temperature, day.Precipitation)
		// this is my first comment
	}
}
