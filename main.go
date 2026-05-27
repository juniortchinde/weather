package main

import (
	"fmt"
	"weather/json"
)

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

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	jsonStations, _ := json.ExtractJson("./data/weather_data.json")

	st := TransformJsonToModel(jsonStations)
	for _, station := range st {
		fmt.Println("{"+
			"country : ", station.Country, ", "+
			"Altitude : ", station.Altitude, ", ",
			"DeviceModel", station.DeviceModel,
		)
	}

}

func TransformJsonToModel(extractedJson json.Stations) (st []Station) {

	for _, stationJson := range extractedJson.Stations {
		var station Station

		station.Country = countryMap[stationJson.Country]
		station.Altitude = stationJson.Altitude
		station.DeviceModel = stationJson.Device.Model

		for _, observationJson := range stationJson.Obs {

			observation := Observation{
				Temperature: observationJson.Temperature,
				Condition:   observationJson.Conditions,
				Wind: Wind{
					Speed:     observationJson.Wind.Speed,
					Direction: int(observationJson.Wind.Direction),
				},
			}

			station.Observations = append(station.Observations, observation)
			st = append(st, station)
		}
	}

	return
}
