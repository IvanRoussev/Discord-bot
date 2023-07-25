package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	Name      string  `json:"name"`
	Region    string  `json:"region"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Timezone  string  `json:"tz_id"`
	Localtime string  `json:"localtime"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type CurrentWeather struct {
	LastUpdatedEpoch int64     `json:"last_updated_epoch"`
	LastUpdated      string    `json:"last_updated"`
	TemperatureC     float64   `json:"temp_c"`
	TemperatureF     float64   `json:"temp_f"`
	IsDay            int       `json:"is_day"`
	Condition        Condition `json:"condition"`
	WindMph          float64   `json:"wind_mph"`
	WindKph          float64   `json:"wind_kph"`
	WindDegree       int       `json:"wind_degree"`
	WindDirection    string    `json:"wind_dir"`
	PressureMb       float64   `json:"pressure_mb"`
	PressureIn       float64   `json:"pressure_in"`
	PrecipitationMm  float64   `json:"precip_mm"`
	PrecipitationIn  float64   `json:"precip_in"`
	Humidity         int       `json:"humidity"`
	Cloud            int       `json:"cloud"`
	FeelsLikeC       float64   `json:"feelslike_c"`
	FeelsLikeF       float64   `json:"feelslike_f"`
	VisibilityKm     float64   `json:"vis_km"`
	VisibilityMiles  float64   `json:"vis_miles"`
	UVIndex          float64   `json:"uv"`
	GustMph          float64   `json:"gust_mph"`
	GustKph          float64   `json:"gust_kph"`
}

type WeatherData struct {
	Location Location       `json:"location"`
	Current  CurrentWeather `json:"current"`
}

func getWeatherData(url string) (*WeatherData, error) {

	// if len(url) != 0  {
	// 	fmt.Println("No URL PROVIDED")
	// 	return nil, nil
	// }

	resp, err := http.Get(url)

	if err != nil{
		fmt.Printf("An Error Occurred \n %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP request failed \n %d", resp.StatusCode)
		return nil, err
	}

	var data WeatherData

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Printf("Something occured %s", err)
		return nil, err
	}

	return &data, nil

}



