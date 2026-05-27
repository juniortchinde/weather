package xml

import "encoding/xml"

type StationXml struct {
	XMLName xml.Name `xml:"station"`

	// Attributs avec ,attr
	ID      string `xml:"id,attr"`
	Country string `xml:"country,attr"`

	// Élément enfant — pas de ,attr
	Name string `xml:"name"`

	// Élément avec sous-attributs
	Coords struct {
		Lat float64 `xml:"lat,attr"`
		Lon float64 `xml:"lon,attr"`
	} `xml:"coordinates"`
}
