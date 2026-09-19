package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func main() {
	city := flag.String("city", "", "city name to check weather for")
	flag.Parse()

	if *city == "" {
		fmt.Println("usage: weather-cli --city=\"city name\"")
		os.Exit(1)
	}

	// TODO: move this to env var before pushing, keep forgetting lol
	apiKey := "YOUR_API_KEY_HERE"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", *city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error fetching weather:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var w WeatherResponse
	if err := json.Unmarshal(body, &w); err != nil {
		fmt.Println("error parsing response:", err)
		os.Exit(1)
	}

	fmt.Printf("Weather in %s: %.1f°C, %s\n", w.Name, w.Main.Temp, w.Weather[0].Description)
}
