package main

import (
	"fmt"
	"net/http"
	"weather/json"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	http.ListenAndServe(":8080", mux)
}

var countryMap = map[string]string{
	"France":    "FR",
	"Espagne":   "ES",
	"Portugal":  "PT",
	"Italie":    "IT",
	"Allemagne": "DE",
	"Belgique":  "BE",
	"Pays-Bas":  "NL",
	"Autriche":  "AT",
	"Suisse":    "CH",
	"Danemark":  "DK",
	"Suède":     "SE",
	"Norvège":   "NO",
	"Pologne":   "PL",
	"Tchéquie":  "CZ",
}

func TransformJsonToModel(extractedJson json.Stations) (st []Station) {

	for _, stationJson := range extractedJson.Stations {
		var station Station

		station.Country = countryMap[stationJson.Country]
		station.Altitude = stationJson.Altitude
		station.DeviceModel = stationJson.Device.Model

		var observation Observation
		for _, observationJson := range stationJson.Obs {

			observation = Observation{
				TimeStamp:   observationJson.Timestamp,
				Temperature: observationJson.Temperature,
				Condition:   observationJson.Conditions,
				Wind: Wind{
					Speed:     observationJson.Wind.Speed,
					Direction: int(observationJson.Wind.Direction),
				},
			}
		}
		station.Observations = append(station.Observations, observation)
		st = append(st, station)
	}
	return
}
