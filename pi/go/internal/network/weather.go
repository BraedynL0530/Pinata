package network

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BraedynL0530/Pinata/internal/core"
)

type Weather struct {
	City          string
	Temperature   float64
	Precipitation float64
	Skys          string
	isDay         bool
}

type location struct {
	lat float64 `json:"lat"`
	lon float64 `json:"lon"`
}

func getLatLon(cfg *core.Config) *location {
	if cfg.Extra.Location != "" {
		return nil
	}
	//address := cfg.Extra.Location
	url := "api.openstreetmap.org/"
	resp, err := http.Get(url) // not gonna work prolly
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	return &location{
		lat: 0.0, lon: 0.0,
	}

}

func NewWeather(cfg *core.Config) *Weather {
	return &Weather{
		City:          cfg.Extra.Location,
		Temperature:   0.0, // place holders ill have display N/A
		Precipitation: 0.0,
		Skys:          "",
		isDay:         true,
	}
}

func GetWeather(cfg *core.Config, location *location) *Weather { // add error handling when im done being lazy
	lat := location.lat
	lon := location.lon
	var unit string

	//TODO: add auth

	if cfg.Extra.MetricSystem == true {
		unit = "celsius"
	} else {
		unit = "fahrenheit"
	}

	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.4f&longitude=%.4f&current=temperature_2m,is_day,relative_humidity_2m,precipitation,rain,showers,snowfall&forecast_days=1&temperature_unit=%s", lat, lon, unit)
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return nil
	}

	var weatherResp Weather
	err = json.NewDecoder(resp.Body).Decode(&weatherResp)
	if err != nil {

		return nil
	}

	return &Weather{
		City:          "", //  use getLatLon or seperate script
		Temperature:   weatherResp.Temperature,
		Precipitation: weatherResp.Precipitation,
		Skys:          "",    // if prectiation then not cloudy or read api
		isDay:         false, //time stuff
	}
}

func UpdateWeather(cfg *core.Config, location *location) *Weather {
	return &Weather{}
	//either make this scheduled every hour or so OR remove this func!
}

//pretty good i suppose!
//HOW LONG IS THIS GOING TO TAKEEEE
