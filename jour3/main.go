package main

import (
	"fmt"
	"log"
	"net/http"
	"weather/json"
)

func main() {

	stations, err := LoadFromJson("weather_data.json")
	if err != nil {
		log.Fatal(err)
	}
	store := NewStore()
	for _, s := range stations {
		store.Put(s)
	}
	log.Printf("bootstrap : %d stations chargées", len(stations))
	app := &App{store: store}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	mux.HandleFunc("GET /stations", app.listStations)
	mux.HandleFunc("GET /stations/{id}", app.getStation)
	mux.HandleFunc("POST /stations", app.createStation)
	mux.HandleFunc("PUT /stations/{id}", app.updateStation)
	mux.HandleFunc("DELETE /stations/{id}", app.deleteStation)
	mux.HandleFunc("GET /stations/{id}/observations", app.listObservations)

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

func LoadFromJson(path string) ([]Station, error) {
	jsonStations, err := json.ExtractJson(path)
	if err != nil {
		return nil, err
	}
	return transformJsonToModel(jsonStations), nil
}

func transformJsonToModel(extractedJson json.Stations) (st []Station) {

	for _, stationJson := range extractedJson.Stations {
		var station Station

		station.Id = stationJson.Id
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
			station.Observations = append(station.Observations, observation)
		}

		st = append(st, station)
	}
	return
}
