package main

import (
	"fmt"
	"weather/json"
	"weather/xml"
)

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

	xmlStations, _ := xml.ExtractXml("./data/weather_data.xml")
	st2 := TransformXmlToModel(xmlStations)
	for _, station := range st2 {
		fmt.Println("{"+
			"country : ", station.Country, ", "+
			"Altitude : ", station.Altitude, ", ",
			"DeviceModel", station.DeviceModel,
		)
	}

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

func TransformXmlToModel(extractedXml xml.Stations) (st2 []Station) {

	for _, stationXml := range extractedXml.Stations {
		var station Station

		station.Country = stationXml.Country
		station.Altitude = stationXml.Coords.Altitude
		station.DeviceModel = stationXml.Hardware.Model

		var observation Observation
		for _, observationXml := range stationXml.Observations {

			var temperature float64
			for _, m := range observationXml.Measures {
				switch m.Type {
				case "temperature":
					temperature = m.Value
				}
			}

			observation = Observation{
				TimeStamp:   observationXml.TimeStamp,
				Temperature: temperature,
				Condition:   "",
				Wind: Wind{
					Speed:     observationXml.Wind.Speed,
					Direction: int(observationXml.Wind.Direction),
				},
			}
		}
		station.Observations = append(station.Observations, observation)
		st2 = append(st2, station)
	}
	return
}
