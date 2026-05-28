package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type Stations struct {
	Stations []Station `json:"stations"`
}
type Station struct {
	Id       string        `json:"id"`
	Country  string        `json:"country"`
	Altitude int           `json:"altitude_m"`
	Obs      []Observation `json:"observations"`
	Device   Device        `json:"device"`
	Location Coordinate    `json:"location"`
}

type Device struct {
	Model        string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	InstalledOn  string `json:"installed_on"`
}
type Observation struct {
	Timestamp   string  `json:"timestamp"`
	Temperature float64 `json:"temperature_celsius"`
	Conditions  string  `json:"conditions"`
	Wind        Wind    `json:"wind"`
	Notes       *string `json:"notes"`
}

type Wind struct {
	Speed     float64 `json:"speed_kmh"`
	Direction float64 `json:"direction_deg"`
}
type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func ExtractJson(path string) (stations Stations, err error) {
	raw, err := os.ReadFile(path)

	if err != nil {
		return
	}
	var stationList Stations

	if err = json.Unmarshal(raw, &stationList); err != nil {
		fmt.Println("error:", err)
		return stationList, err
	}
	return stationList, nil
}
