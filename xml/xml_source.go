package xml

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Stations struct {
	XMLName  xml.Name     `xml:"weather_dataset"`
	Stations []StationXml `xml:"station"`
}
type StationXml struct {
	XMLName xml.Name `xml:"station"`
	// Attributs avec ,attr
	ID      string `xml:"id,attr"`
	Country string `xml:"country,attr"`
	// Élément enfant — pas de ,attr
	Name string `xml:"name"`
	// Élément avec sous-attributs
	Coords struct {
		Lat      float64 `xml:"lat,attr"`
		Lon      float64 `xml:"lon,attr"`
		Altitude int     `xml:"altitude,attr"`
	} `xml:"coordinates"`

	Hardware     hardware      `xml:"hardware"`
	Observations []observation `xml:"observation"`
}

type observation struct {
	TimeStamp string       `xml:"at,attr"`
	Measures  []xmlMeasure `xml:"measure"`
	Wind      wind         `xml:"wind"`
}
type xmlMeasure struct {
	Type  string  `xml:"type,attr"`
	Value float64 `xml:",chardata"`
}
type wind struct {
	Speed     float64 `xml:"speed,attr"`
	Direction float64 `xml:"direction,attr"`
}

type hardware struct {
	Model  string `xml:"model,attr"`
	Vendor string `xml:"vendor,attr"`
	Since  string `xml:"since,attr"`
}

func ExtractXml(path string) (stations Stations, err error) {
	raw, err := os.ReadFile(path)

	if err != nil {
		return
	}
	var stationList Stations

	if err = xml.Unmarshal(raw, &stationList); err != nil {
		fmt.Println("error:", err)
		return stationList, err
	}

	return stationList, nil
}
